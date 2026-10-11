/*
Tests for the hand-written facade.

The generated core is covered by codegen and the spec drift guard; these
exercise only the ergonomics injected on top of it: DAG validation that must
fail before any round-trip, request shaping (options, headers, filters), the
polling loop that turns a sequence of responses into a terminal record or a
timeout, typed errors, the timeout split between regular and lease calls, and
the worker runtime (lease, heartbeat, complete, fail, lost lease, stop and
drain). The transport is a stdlib httptest server, so nothing here needs a
live QueueFlow.
*/

package queueflow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// newTestQueueFlow points a real client at a local test server, so the request
// path, auth header, and JSON decoding are all genuinely exercised.
func newTestQueueFlow(t *testing.T, handler http.HandlerFunc) *QueueFlow {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return NewQueueFlow(srv.URL, "test-token")
}

// The generated core decodes with DisallowUnknownFields and requires every
// non-omitempty field, so these fixtures carry the full required set rather
// than a convenient subset.
func jobJSON(id string, status JobStatus) string {
	return jobJSONTask(id, "echo", status)
}

func jobJSONTask(id, task string, status JobStatus) string {
	return fmt.Sprintf(`{
		"id":%q,
		"task_name":%q,
		"status":%q,
		"queue_name":"default",
		"payload":{"n":1},
		"retry_count":0,
		"created_at":"2026-01-01T00:00:00Z",
		"scheduled_at":"2026-01-01T00:00:00Z",
		"config":{"max_retries":3,"priority":0,"retry_delay_secs":1,"retry_max_delay_secs":60,"timeout_secs":300}
	}`, id, task, status)
}

func workflowJSON(id string, status WorkflowStatus) string {
	return fmt.Sprintf(`{
		"id":%q,
		"name":"wf",
		"status":%q,
		"steps":[{"name":"a","task_name":"task"}],
		"created_at":"2026-01-01T00:00:00Z"
	}`, id, status)
}

func decodeBody(t *testing.T, r *http.Request) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil && err != io.EOF {
		t.Errorf("decoding request body: %v", err)
	}
	return m
}

// --- WorkflowBuilder validation -------------------------------------------

// Every one of these must be caught locally: the point of the builder is that a
// malformed DAG never reaches the network.
func TestWorkflowBuilderRejectsInvalidDAGs(t *testing.T) {
	cases := []struct {
		name    string
		build   func() *WorkflowBuilder
		wantErr string
	}{
		{
			name:    "empty name",
			build:   func() *WorkflowBuilder { return NewWorkflowBuilder("").Step("a", "task") },
			wantErr: "workflow name is required",
		},
		{
			name:    "no steps",
			build:   func() *WorkflowBuilder { return NewWorkflowBuilder("etl") },
			wantErr: `workflow "etl" has no steps`,
		},
		{
			name: "duplicate step name",
			build: func() *WorkflowBuilder {
				return NewWorkflowBuilder("etl").Step("a", "task").Step("a", "other")
			},
			wantErr: `duplicate step name "a"`,
		},
		{
			name: "dangling dependency",
			build: func() *WorkflowBuilder {
				return NewWorkflowBuilder("etl").Step("a", "task", "ghost")
			},
			wantErr: `step "a" depends on unknown step "ghost"`,
		},
		{
			name: "direct cycle",
			build: func() *WorkflowBuilder {
				return NewWorkflowBuilder("etl").Step("a", "task", "b").Step("b", "task", "a")
			},
			wantErr: "dependency cycle",
		},
		{
			name: "self cycle",
			build: func() *WorkflowBuilder {
				return NewWorkflowBuilder("etl").Step("a", "task", "a")
			},
			wantErr: "dependency cycle",
		},
		{
			name: "long cycle",
			build: func() *WorkflowBuilder {
				return NewWorkflowBuilder("etl").
					Step("a", "task", "c").
					Step("b", "task", "a").
					Step("c", "task", "b")
			},
			wantErr: "dependency cycle",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body, err := tc.build().Build()
			if err == nil {
				t.Fatalf("expected an error, got a request body: %+v", body)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error %q does not contain %q", err, tc.wantErr)
			}
		})
	}
}

// A diamond is the interesting valid case: `load` is reachable by two paths, so
// a cycle checker that confuses "already visited" with "currently visiting"
// would wrongly reject it.
func TestWorkflowBuilderAcceptsDiamond(t *testing.T) {
	body, err := NewWorkflowBuilder("etl").
		Step("extract", "fetch").
		Step("left", "normalize", "extract").
		Step("right", "enrich", "extract").
		Step("load", "upsert", "left", "right").
		Build()
	if err != nil {
		t.Fatalf("valid diamond DAG rejected: %v", err)
	}
	if body.Name != "etl" {
		t.Errorf("Name = %q, want %q", body.Name, "etl")
	}
	if len(body.Steps) != 4 {
		t.Fatalf("Steps = %d, want 4", len(body.Steps))
	}
	if got := body.Steps[3].DependsOn; len(got) != 2 || got[0] != "left" || got[1] != "right" {
		t.Errorf("load.DependsOn = %v, want [left right]", got)
	}
	// A step declared without dependencies must not carry an empty slice that
	// would serialize as `"depends_on": []` and imply an explicit no-op gate.
	if body.Steps[0].DependsOn != nil {
		t.Errorf("extract.DependsOn = %v, want nil", body.Steps[0].DependsOn)
	}
}

// StepWithOptions must put every per-step option into the wire shape, and the
// builder-level Context/Metadata must merge across calls.
func TestWorkflowBuilderStepOptionsAndContext(t *testing.T) {
	cfg := JobConfig{MaxRetries: PtrInt32(5)}
	body, err := NewWorkflowBuilder("etl").
		StepWithOptions("extract", "fetch", StepOptions{Payload: map[string]interface{}{"n": 1}}).
		StepWithOptions("load", "upsert", StepOptions{
			After:     []string{"extract"},
			Config:    &cfg,
			OnFailure: ONFAILURE_CONTINUE,
			OnSuccess: ONSUCCESS_CONTINUE,
			Metadata:  map[string]interface{}{"owner": "data"},
		}).
		Context(map[string]interface{}{"a": 1}).
		Context(map[string]interface{}{"b": 2}).
		Metadata(map[string]interface{}{"source": "test"}).
		Build()
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatal(err)
	}
	steps := got["steps"].([]any)
	first := steps[0].(map[string]any)
	if first["payload"].(map[string]any)["n"] != float64(1) {
		t.Errorf("extract.payload = %v", first["payload"])
	}
	if _, present := first["on_failure"]; present {
		t.Errorf("extract.on_failure should be omitted when unset, got %v", first["on_failure"])
	}
	second := steps[1].(map[string]any)
	if second["on_failure"] != "continue" || second["on_success"] != "continue" {
		t.Errorf("load policies = %v / %v", second["on_failure"], second["on_success"])
	}
	if second["config"].(map[string]any)["max_retries"] != float64(5) {
		t.Errorf("load.config = %v", second["config"])
	}
	if second["metadata"].(map[string]any)["owner"] != "data" {
		t.Errorf("load.metadata = %v", second["metadata"])
	}
	if deps := second["depends_on"].([]any); len(deps) != 1 || deps[0] != "extract" {
		t.Errorf("load.depends_on = %v", deps)
	}
	ctx := got["context"].(map[string]any)
	if ctx["a"] != float64(1) || ctx["b"] != float64(2) {
		t.Errorf("context = %v, want both merged keys", ctx)
	}
	if got["metadata"].(map[string]any)["source"] != "test" {
		t.Errorf("metadata = %v", got["metadata"])
	}
}

// --- Client construction ---------------------------------------------------

func TestNewQueueFlowWithOptionsValidates(t *testing.T) {
	if _, err := NewQueueFlowWithOptions(Options{Token: "t"}); err == nil {
		t.Error("expected an error for a missing BaseURL")
	}
	if _, err := NewQueueFlowWithOptions(Options{BaseURL: "http://x"}); err == nil {
		t.Error("expected an error for a missing Token")
	}
	qf, err := NewQueueFlowWithOptions(Options{BaseURL: "http://x/", Token: "t", Timeout: 5 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	if qf.Timeout() != 5*time.Second {
		t.Errorf("Timeout = %s, want 5s", qf.Timeout())
	}
	if NewQueueFlow("http://x", "t").Timeout() != DefaultTimeout {
		t.Error("default timeout not applied")
	}
}

// Tenant routes carry the tenant token, worker routes the worker token.
func TestWorkerRoutesUseTheWorkerToken(t *testing.T) {
	var mu sync.Mutex
	auth := map[string]string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		auth[r.URL.Path] = r.Header.Get("Authorization")
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/lease") {
			fmt.Fprint(w, `{"jobs":[]}`)
			return
		}
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_PENDING))
	}))
	t.Cleanup(srv.Close)
	qf, err := NewQueueFlowWithOptions(Options{BaseURL: srv.URL + "/", Token: "tenant", WorkerToken: "worker"})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if _, err := qf.GetJob(ctx, "job-1"); err != nil {
		t.Fatal(err)
	}
	if _, err := qf.LeaseJobs(ctx, "q", LeaseOptions{}); err != nil {
		t.Fatal(err)
	}
	if auth["/api/v1/jobs/job-1"] != "Bearer tenant" {
		t.Errorf("tenant route auth = %q", auth["/api/v1/jobs/job-1"])
	}
	if auth["/api/v1/queues/q/lease"] != "Bearer worker" {
		t.Errorf("worker route auth = %q", auth["/api/v1/queues/q/lease"])
	}
}

// --- CreateJob -------------------------------------------------------------

// CreateJob returns the id the server handed back and sends only what was set.
func TestCreateJobPostsAndReturnsTheId(t *testing.T) {
	var posted map[string]any
	var gotAuth, gotKey string

	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/jobs" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		gotAuth = r.Header.Get("Authorization")
		gotKey = r.Header.Get("Idempotency-Key")
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"job_id":"job-1"}`)
	})

	id, err := qf.CreateJob(context.Background(), "echo", map[string]interface{}{"hi": 1})
	if err != nil {
		t.Fatalf("CreateJob: %v", err)
	}
	if id != "job-1" {
		t.Errorf("id = %q, want %q", id, "job-1")
	}
	if posted["task_name"] != "echo" {
		t.Errorf("task_name = %v, want echo", posted["task_name"])
	}
	if payload, ok := posted["payload"].(map[string]any); !ok || payload["hi"] != float64(1) {
		t.Errorf("payload = %v, want {hi: 1}", posted["payload"])
	}
	for _, k := range []string{"config", "run_at"} {
		if _, present := posted[k]; present {
			t.Errorf("%s should be omitted when unset, got %v", k, posted[k])
		}
	}
	if gotKey != "" {
		t.Errorf("Idempotency-Key = %q, want none", gotKey)
	}
	if gotAuth != "Bearer test-token" {
		t.Errorf("Authorization = %q, want %q", gotAuth, "Bearer test-token")
	}
}

// Every option must land in the right wire field, and the key in the header.
func TestCreateJobWithOptionsSendsEveryField(t *testing.T) {
	var posted map[string]any
	var gotKey string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.Header.Get("Idempotency-Key")
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"job_id":"job-1"}`)
	})

	runAt := time.Date(2030, 1, 2, 3, 4, 5, 0, time.UTC)
	_, err := qf.CreateJobWithOptions(context.Background(), "echo", CreateJobOptions{
		Payload:           map[string]interface{}{"hi": 1},
		Queue:             "billing",
		Priority:          PtrInt32(7),
		MaxRetries:        PtrInt32(2),
		Timeout:           PtrInt64(45),
		RetryBackoff:      BACKOFFSTRATEGY_LINEAR,
		RetryDelaySecs:    PtrInt64(3),
		RetryMaxDelaySecs: PtrInt64(90),
		JitterFactor:      PtrFloat64(0.25),
		IdempotencyKey:    "order-1",
		RunAt:             &runAt,
	})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}
	if gotKey != "order-1" {
		t.Errorf("Idempotency-Key = %q, want order-1", gotKey)
	}
	cfg, ok := posted["config"].(map[string]any)
	if !ok {
		t.Fatalf("config = %v, want an object", posted["config"])
	}
	want := map[string]any{
		"queue": "billing", "priority": float64(7), "max_retries": float64(2), "timeout": float64(45),
		"retry_backoff": "linear", "retry_delay_secs": float64(3), "retry_max_delay_secs": float64(90),
		"jitter_factor": 0.25,
	}
	for k, v := range want {
		if cfg[k] != v {
			t.Errorf("config.%s = %v, want %v", k, cfg[k], v)
		}
	}
	if !strings.HasPrefix(fmt.Sprint(posted["run_at"]), "2030-01-02T03:04:05") {
		t.Errorf("run_at = %v, want RFC 3339 of the given instant", posted["run_at"])
	}
}

// A nil payload must stay absent rather than serializing as null, which the
// server would reject as a type error on an object field.
func TestCreateJobOmitsNilPayload(t *testing.T) {
	var posted map[string]any
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"job_id":"job-1"}`)
	})

	if _, err := qf.CreateJob(context.Background(), "echo", nil); err != nil {
		t.Fatalf("CreateJob: %v", err)
	}
	if v, present := posted["payload"]; present && v == nil {
		t.Error("payload serialized as null; want the key omitted")
	}
}

// Non-2xx responses become *APIError with the status and body attached.
func TestServerErrorsBecomeAPIError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":"no such job"}`)
	})

	job, err := qf.GetJob(context.Background(), "missing")
	if err == nil {
		t.Fatalf("expected an error, got job %+v", job)
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error %T is not an *APIError: %v", err, err)
	}
	if apiErr.Status != 404 || StatusCode(err) != 404 {
		t.Errorf("Status = %d, StatusCode = %d, want 404", apiErr.Status, StatusCode(err))
	}
	if apiErr.Op != "GetJob" || !strings.Contains(string(apiErr.Body), "no such job") {
		t.Errorf("Op = %q, Body = %q", apiErr.Op, apiErr.Body)
	}
	if !strings.Contains(err.Error(), "404") || !strings.Contains(err.Error(), "no such job") {
		t.Errorf("Error() = %q", err)
	}
	var g *GenericOpenAPIError
	if !errors.As(err, &g) {
		t.Error("the generated error should remain reachable through Unwrap")
	}
}

// Transport failures (here: a cancelled context) are not APIErrors.
func TestContextCancellationIsNotAnAPIError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	})
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := qf.GetJob(ctx, "job-1")
	if err == nil {
		t.Fatal("expected a deadline error")
	}
	if StatusCode(err) != 0 {
		t.Errorf("StatusCode = %d, want 0 for a transport error", StatusCode(err))
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want it to wrap context.DeadlineExceeded", err)
	}
}

// The per-request timeout applies to regular calls but not to the lease long
// poll, which gets wait_secs plus 35 seconds instead.
func TestLeaseCallOutlivesTheRequestTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/lease") {
			fmt.Fprint(w, `{"jobs":[]}`)
			return
		}
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_PENDING))
	}))
	t.Cleanup(srv.Close)
	qf, err := NewQueueFlowWithOptions(Options{BaseURL: srv.URL, Token: "t", Timeout: 50 * time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := qf.GetJob(context.Background(), "job-1"); err == nil {
		t.Error("GetJob should have hit the 50ms timeout")
	}
	if _, err := qf.LeaseJobs(context.Background(), "q", LeaseOptions{WaitSecs: 1}); err != nil {
		t.Errorf("LeaseJobs should not be bound by the 50ms timeout: %v", err)
	}
}

// --- Lists -----------------------------------------------------------------

// Set filters are forwarded as query parameters; unset ones are omitted.
func TestListJobsForwardsFilters(t *testing.T) {
	var query string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		query = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"jobs":[],"limit":50,"offset":0,"has_more":false}`)
	})
	after := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	replayed := false
	_, err := qf.ListJobs(context.Background(), ListOptions{Status: "failed", Queue: "billing", TaskName: "charge", Limit: 50, Cursor: "abc", CreatedAfter: &after, Replayed: &replayed})
	if err != nil {
		t.Fatalf("ListJobs: %v", err)
	}
	for _, want := range []string{"status=failed", "queue=billing", "task_name=charge", "limit=50", "cursor=abc", "created_after=2026-01-01"} {
		if !strings.Contains(query, want) {
			t.Errorf("query %q lacks %q", query, want)
		}
	}
	// Replayed is a dead-letter filter; the job list never sends it.
	for _, absent := range []string{"offset=", "order_by=", "include_total=", "created_before=", "replayed="} {
		if strings.Contains(query, absent) {
			t.Errorf("query %q should not carry %q", query, absent)
		}
	}
}

// The dead-letter list forwards TaskName and both values of Replayed (false
// is a real filter, so it cannot be dropped as a zero value).
func TestListDeadLettersForwardsTaskNameAndReplayed(t *testing.T) {
	var queries []string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/dlq" {
			t.Errorf("path = %s, want /api/v1/dlq", r.URL.Path)
		}
		queries = append(queries, r.URL.RawQuery)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"dead_letters":[],"limit":10,"offset":0,"has_more":false}`)
	})
	ctx := context.Background()
	notReplayed := false
	if _, err := qf.ListDeadLetters(ctx, ListOptions{Queue: "billing", TaskName: "charge", Replayed: &notReplayed, Limit: 10}); err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	for _, want := range []string{"queue=billing", "task_name=charge", "replayed=false", "limit=10"} {
		if !strings.Contains(queries[0], want) {
			t.Errorf("query %q lacks %q", queries[0], want)
		}
	}
	yes := true
	if _, err := qf.ListDeadLetters(ctx, ListOptions{Replayed: &yes}); err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	if queries[1] != "replayed=true" {
		t.Errorf("query = %q, want exactly replayed=true", queries[1])
	}
	if _, err := qf.ListDeadLetters(ctx, ListOptions{}); err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	if queries[2] != "" {
		t.Errorf("query = %q, want no parameters for the zero ListOptions", queries[2])
	}
}

// --- Bulk cancel and replay ------------------------------------------------

// Set filter fields are posted under their wire names, unset ones are omitted,
// and the cancelled count is unwrapped.
func TestCancelJobsPostsTheFilterAndReturnsTheCount(t *testing.T) {
	var posted map[string]any
	var method, path string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"cancelled":7}`)
	})
	before := time.Date(2026, 3, 1, 12, 0, 0, 0, time.FixedZone("plus2", 2*3600))
	n, err := qf.CancelJobs(context.Background(), CancelJobsFilter{
		Queue:         "billing",
		TaskName:      "charge",
		Status:        "pending",
		CreatedBefore: &before,
		Limit:         25,
	})
	if err != nil {
		t.Fatalf("CancelJobs: %v", err)
	}
	if n != 7 {
		t.Errorf("cancelled = %d, want 7", n)
	}
	if method != http.MethodPost || path != "/api/v1/jobs/cancel" {
		t.Errorf("%s %s, want POST /api/v1/jobs/cancel", method, path)
	}
	if posted["queue"] != "billing" || posted["task_name"] != "charge" || posted["status"] != "pending" {
		t.Errorf("posted = %v", posted)
	}
	if posted["limit"] != float64(25) {
		t.Errorf("limit = %v, want 25", posted["limit"])
	}
	// Times are sent in UTC.
	if posted["created_before"] != "2026-03-01T10:00:00Z" {
		t.Errorf("created_before = %v, want 2026-03-01T10:00:00Z", posted["created_before"])
	}
	if _, ok := posted["created_after"]; ok {
		t.Errorf("created_after should be omitted, posted = %v", posted)
	}
}

// The zero filter is a valid request (cancel the oldest 100) and sends an
// empty object, never nulls.
func TestCancelJobsZeroFilterSendsAnEmptyBody(t *testing.T) {
	var raw string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		raw = strings.TrimSpace(string(b))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"cancelled":0}`)
	})
	n, err := qf.CancelJobs(context.Background(), CancelJobsFilter{})
	if err != nil {
		t.Fatalf("CancelJobs: %v", err)
	}
	if n != 0 {
		t.Errorf("cancelled = %d, want 0", n)
	}
	if raw != "{}" {
		t.Errorf("body = %q, want {}", raw)
	}
}

// A terminal status is refused by the server; the 400 surfaces as an APIError
// carrying the operation name.
func TestCancelJobsRejectedStatusIsAnAPIError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"status must be pending, retrying or running"}`)
	})
	_, err := qf.CancelJobs(context.Background(), CancelJobsFilter{Status: "completed"})
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error %T is not an *APIError: %v", err, err)
	}
	if apiErr.Status != 400 || apiErr.Op != "CancelJobs" || !strings.Contains(string(apiErr.Body), "pending, retrying or running") {
		t.Errorf("apiErr = %+v", apiErr)
	}
}

func TestReplayDeadLettersPostsTheFilterAndDecodesResults(t *testing.T) {
	var posted map[string]any
	var method, path string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"replayed":1,"results":[{"dead_letter_id":11,"job_id":"job-new"},{"dead_letter_id":12,"error":"original job purged"}]}`)
	})
	after := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	out, err := qf.ReplayDeadLetters(context.Background(), ReplayDeadLettersFilter{Queue: "billing", TaskName: "charge", CreatedAfter: &after, Limit: 2})
	if err != nil {
		t.Fatalf("ReplayDeadLetters: %v", err)
	}
	if method != http.MethodPost || path != "/api/v1/dlq/replay" {
		t.Errorf("%s %s, want POST /api/v1/dlq/replay", method, path)
	}
	if posted["queue"] != "billing" || posted["task_name"] != "charge" || posted["limit"] != float64(2) || posted["created_after"] != "2026-01-01T00:00:00Z" {
		t.Errorf("posted = %v", posted)
	}
	if _, ok := posted["created_before"]; ok {
		t.Errorf("created_before should be omitted, posted = %v", posted)
	}
	if out.Replayed != 1 || len(out.Results) != 2 {
		t.Fatalf("response = %+v", out)
	}
	ok, failed := out.Results[0], out.Results[1]
	if ok.DeadLetterId != 11 || ok.GetJobId() != "job-new" || ok.Error.IsSet() {
		t.Errorf("results[0] = %+v", ok)
	}
	if failed.DeadLetterId != 12 || failed.GetError() != "original job purged" || failed.JobId.IsSet() {
		t.Errorf("results[1] = %+v", failed)
	}
}

// --- Queue control and identity --------------------------------------------

func TestPauseAndResumeQueuePostToTheQueuePath(t *testing.T) {
	var calls []string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		calls = append(calls, r.Method+" "+r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		paused := strings.HasSuffix(r.URL.Path, "/pause")
		fmt.Fprintf(w, `{"queue":"billing","paused":%t,"running":2,"pending":5}`, paused)
	})
	ctx := context.Background()
	paused, err := qf.PauseQueue(ctx, "billing")
	if err != nil {
		t.Fatalf("PauseQueue: %v", err)
	}
	if !paused.Paused || paused.Queue != "billing" || paused.Running != 2 || paused.Pending != 5 {
		t.Errorf("paused = %+v", paused)
	}
	resumed, err := qf.ResumeQueue(ctx, "billing")
	if err != nil {
		t.Fatalf("ResumeQueue: %v", err)
	}
	if resumed.Paused {
		t.Errorf("resumed = %+v, want paused=false", resumed)
	}
	want := []string{"POST /api/v1/queues/billing/pause", "POST /api/v1/queues/billing/resume"}
	if strings.Join(calls, ",") != strings.Join(want, ",") {
		t.Errorf("calls = %v, want %v", calls, want)
	}
}

// Queue names are path segments and must be escaped on the way out.
func TestPauseQueueEscapesTheQueueName(t *testing.T) {
	var rawPath string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		rawPath = r.URL.EscapedPath()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"queue":"a/b c","paused":true,"running":0,"pending":0}`)
	})
	if _, err := qf.PauseQueue(context.Background(), "a/b c"); err != nil {
		t.Fatalf("PauseQueue: %v", err)
	}
	if rawPath != "/api/v1/queues/a%2Fb%20c/pause" {
		t.Errorf("path = %q", rawPath)
	}
}

func TestWhoAmIReturnsTenantAndAccess(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/v1/whoami" {
			t.Errorf("%s %s, want GET /api/v1/whoami", r.Method, r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Authorization = %q", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"tenant":"acme","access":"read_only"}`)
	})
	me, err := qf.WhoAmI(context.Background())
	if err != nil {
		t.Fatalf("WhoAmI: %v", err)
	}
	if me.Tenant != "acme" || me.Access != "read_only" {
		t.Errorf("whoami = %+v", me)
	}
}

func TestQueuesAndTasksUnwrapTheirEnvelopes(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/queues":
			fmt.Fprint(w, `{"queues":[{"queue":"default","pending":2,"scheduled":0,"running":1,"oldest_pending_age_secs":12,"paused":true}]}`)
		case "/api/v1/tasks":
			fmt.Fprint(w, `{"tasks":["echo","sleep"]}`)
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
		}
	})
	queues, err := qf.Queues(context.Background())
	if err != nil {
		t.Fatalf("Queues: %v", err)
	}
	if len(queues) != 1 || queues[0].Queue != "default" || queues[0].Pending != 2 || !queues[0].GetPaused() {
		t.Errorf("queues = %+v", queues)
	}
	tasks, err := qf.Tasks(context.Background())
	if err != nil {
		t.Fatalf("Tasks: %v", err)
	}
	if len(tasks) != 2 || tasks[1] != "sleep" {
		t.Errorf("tasks = %v", tasks)
	}
}

// --- Cron ------------------------------------------------------------------

func TestCreateCronWithOptionsSendsQueueAndPayload(t *testing.T) {
	var posted map[string]any
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		posted = decodeBody(t, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"cron_id":"c1"}`)
	})
	id, err := qf.CreateCronWithOptions(context.Background(), "nightly", "0 3 * * *", "report", CreateCronOptions{
		Payload: map[string]interface{}{"from": "cron"},
		Queue:   "reports",
	})
	if err != nil {
		t.Fatalf("CreateCronWithOptions: %v", err)
	}
	if id != "c1" {
		t.Errorf("id = %q", id)
	}
	if posted["name"] != "nightly" || posted["cron_expr"] != "0 3 * * *" || posted["task_name"] != "report" {
		t.Errorf("posted = %v", posted)
	}
	if posted["queue"] != "reports" || posted["payload"].(map[string]any)["from"] != "cron" {
		t.Errorf("queue/payload = %v / %v", posted["queue"], posted["payload"])
	}
}

// --- WaitForJob ------------------------------------------------------------

// The poll loop must keep going through non-terminal states and stop at the
// first terminal one, returning that record.
func TestWaitForJobPollsUntilTerminal(t *testing.T) {
	var calls int32
	statuses := []JobStatus{JOBSTATUS_PENDING, JOBSTATUS_RUNNING, JOBSTATUS_COMPLETED}

	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		idx := int(n) - 1
		if idx >= len(statuses) {
			idx = len(statuses) - 1
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jobJSON("job-1", statuses[idx]))
	})

	job, err := qf.WaitForJob(context.Background(), "job-1", WaitOptions{Timeout: 5 * time.Second, Interval: time.Millisecond})
	if err != nil {
		t.Fatalf("WaitForJob: %v", err)
	}
	if job.Status != JOBSTATUS_COMPLETED {
		t.Errorf("Status = %q, want completed", job.Status)
	}
	if got := atomic.LoadInt32(&calls); got != 3 {
		t.Errorf("polled %d times, want exactly 3 (stop on first terminal)", got)
	}
}

// failed and cancelled are terminal too: waiting must return them, not treat a
// failure as "keep polling until the deadline".
func TestWaitForJobReturnsFailureStates(t *testing.T) {
	for _, status := range []JobStatus{JOBSTATUS_FAILED, JOBSTATUS_CANCELLED} {
		t.Run(string(status), func(t *testing.T) {
			qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, jobJSON("job-1", status))
			})
			job, err := qf.WaitForJob(context.Background(), "job-1", WaitOptions{Timeout: time.Second, Interval: time.Millisecond})
			if err != nil {
				t.Fatalf("WaitForJob(%s): %v", status, err)
			}
			if job.Status != status {
				t.Errorf("Status = %q, want %q", job.Status, status)
			}
		})
	}
}

// retrying is explicitly NOT terminal; a job still being retried must not be
// mistaken for a finished one.
func TestWaitForJobTimesOutOnNonTerminalStatus(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_RETRYING))
	})

	start := time.Now()
	job, err := qf.WaitForJob(context.Background(), "job-1", WaitOptions{Timeout: 50 * time.Millisecond, Interval: 10 * time.Millisecond})
	if err == nil {
		t.Fatalf("expected a timeout, got job %+v", job)
	}
	if !errors.Is(err, ErrTimeout) {
		t.Errorf("error = %q, want it to wrap ErrTimeout", err)
	}
	// The loop must give up near the deadline rather than sleeping past it.
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("took %s to time out on a 50ms deadline", elapsed)
	}
}

// A cancelled context ends the wait with ctx.Err(), not a timeout.
func TestWaitForJobHonoursContext(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_RUNNING))
	})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	_, err := qf.WaitForJob(ctx, "job-1", WaitOptions{Timeout: 10 * time.Second, Interval: 5 * time.Millisecond})
	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
}

func TestWaitForJobPropagatesFetchError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := qf.WaitForJob(context.Background(), "job-1", WaitOptions{Timeout: time.Second, Interval: time.Millisecond}); StatusCode(err) != 500 {
		t.Fatalf("expected the 500 to propagate, got %v", err)
	}
}

// --- Workflows -------------------------------------------------------------

// An invalid DAG must fail inside CreateWorkflow before any request is sent.
func TestCreateWorkflowValidatesBeforeSending(t *testing.T) {
	var called int32
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&called, 1)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"workflow_id":"wf-1"}`)
	})

	cyclic := NewWorkflowBuilder("etl").Step("a", "task", "b").Step("b", "task", "a")
	if _, err := qf.CreateWorkflow(context.Background(), cyclic); err == nil {
		t.Fatal("expected a validation error for a cyclic DAG")
	}
	if got := atomic.LoadInt32(&called); got != 0 {
		t.Errorf("sent %d requests, want 0 for a locally-invalid DAG", got)
	}
}

func TestCreateWorkflowReturnsTheId(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/api/v1/workflows" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"workflow_id":"wf-1"}`)
	})

	id, err := qf.CreateWorkflow(context.Background(), NewWorkflowBuilder("etl").Step("a", "task"))
	if err != nil {
		t.Fatalf("CreateWorkflow: %v", err)
	}
	if id != "wf-1" {
		t.Errorf("id = %q, want %q", id, "wf-1")
	}
}

func TestWorkflowStepsUnwrapsTheEnvelope(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"steps":[{"name":"a","status":"completed","job_id":"job-1"},{"name":"b","status":"pending"}]}`)
	})
	steps, err := qf.WorkflowSteps(context.Background(), "wf-1")
	if err != nil {
		t.Fatalf("WorkflowSteps: %v", err)
	}
	if len(steps) != 2 || steps[0].Status != STEPSTATUS_COMPLETED || steps[1].Name != "b" {
		t.Errorf("steps = %+v", steps)
	}
}

// partially_failed is terminal for a workflow but has no job equivalent, so it
// is the case most likely to be missed in the terminal set.
func TestWaitForWorkflowTreatsPartiallyFailedAsTerminal(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, workflowJSON("wf-1", WORKFLOWSTATUS_PARTIALLY_FAILED))
	})

	got, err := qf.WaitForWorkflow(context.Background(), "wf-1", WaitOptions{Timeout: time.Second, Interval: time.Millisecond})
	if err != nil {
		t.Fatalf("WaitForWorkflow: %v", err)
	}
	if got.Status != WORKFLOWSTATUS_PARTIALLY_FAILED {
		t.Errorf("Status = %q, want partially_failed", got.Status)
	}
}

func TestWaitForWorkflowTimesOutWhileRunning(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, workflowJSON("wf-1", WORKFLOWSTATUS_RUNNING))
	})

	if _, err := qf.WaitForWorkflow(context.Background(), "wf-1", WaitOptions{Timeout: 50 * time.Millisecond, Interval: 10 * time.Millisecond}); !errors.Is(err, ErrTimeout) {
		t.Fatalf("expected ErrTimeout while the workflow is still running, got %v", err)
	}
}

// --- Worker runtime --------------------------------------------------------

// fakeBroker is a scripted worker-protocol server: it hands out the queued
// leases (honouring max_jobs), then returns empty leases, and records every
// heartbeat/complete/fail body.
type fakeBroker struct {
	t  *testing.T
	mu sync.Mutex

	pending      []string // job JSON bodies still to be leased
	leaseStatus  int      // non-zero: every lease call fails with this status
	leaseFailN   int32    // fail this many lease calls with 500, then serve
	heartbeat    JobStatus
	heartbeat409 bool

	leaseBodies []map[string]any
	heartbeats  []map[string]any
	completes   map[string]map[string]any
	fails       map[string]map[string]any
	auth        map[string]bool
}

func newFakeBroker(t *testing.T) *fakeBroker {
	return &fakeBroker{
		t:         t,
		heartbeat: JOBSTATUS_RUNNING,
		completes: map[string]map[string]any{},
		fails:     map[string]map[string]any{},
		auth:      map[string]bool{},
	}
}

func (b *fakeBroker) enqueue(id, task string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.pending = append(b.pending, jobJSONTask(id, task, JOBSTATUS_RUNNING))
}

func (b *fakeBroker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.auth[r.Header.Get("Authorization")] = true
	w.Header().Set("Content-Type", "application/json")
	path := r.URL.Path
	switch {
	case strings.HasSuffix(path, "/lease"):
		if b.leaseStatus != 0 {
			w.WriteHeader(b.leaseStatus)
			fmt.Fprint(w, `{"error":"nope"}`)
			return
		}
		if b.leaseFailN > 0 {
			b.leaseFailN--
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `{"error":"flaky"}`)
			return
		}
		body := decodeBody(b.t, r)
		b.leaseBodies = append(b.leaseBodies, body)
		max := 1
		if v, ok := body["max_jobs"].(float64); ok {
			max = int(v)
		}
		var out []string
		for len(b.pending) > 0 && len(out) < max {
			out = append(out, fmt.Sprintf(`{"job":%s,"lease_token":"tok"}`, b.pending[0]))
			b.pending = b.pending[1:]
		}
		if len(out) == 0 {
			// Stand in for the server's long poll so an idle worker does not
			// spin against the test server.
			b.mu.Unlock()
			time.Sleep(20 * time.Millisecond)
			b.mu.Lock()
		}
		fmt.Fprintf(w, `{"jobs":[%s]}`, strings.Join(out, ","))
	case strings.HasSuffix(path, "/heartbeat"):
		b.heartbeats = append(b.heartbeats, decodeBody(b.t, r))
		if b.heartbeat409 {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprint(w, `{"error":"lease lost"}`)
			return
		}
		fmt.Fprintf(w, `{"status":%q}`, b.heartbeat)
	case strings.HasSuffix(path, "/complete"):
		b.completes[jobIDFromPath(path)] = decodeBody(b.t, r)
		w.WriteHeader(http.StatusNoContent)
	case strings.HasSuffix(path, "/fail"):
		b.fails[jobIDFromPath(path)] = decodeBody(b.t, r)
		w.WriteHeader(http.StatusNoContent)
	default:
		b.t.Errorf("unexpected request %s %s", r.Method, path)
		w.WriteHeader(http.StatusNotFound)
	}
}

func jobIDFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-2]
}

func (b *fakeBroker) snapshot() (completes, fails map[string]map[string]any, heartbeats, leases []map[string]any) {
	b.mu.Lock()
	defer b.mu.Unlock()
	completes = map[string]map[string]any{}
	for k, v := range b.completes {
		completes[k] = v
	}
	fails = map[string]map[string]any{}
	for k, v := range b.fails {
		fails[k] = v
	}
	return completes, fails, append([]map[string]any(nil), b.heartbeats...), append([]map[string]any(nil), b.leaseBodies...)
}

// runWorker starts Run against the broker and returns a stop function that
// cancels it and returns Run's error.
func runWorker(t *testing.T, b *fakeBroker, handlers map[string]Handler, opts WorkerOptions) (*QueueFlow, func() error) {
	t.Helper()
	srv := httptest.NewServer(b)
	t.Cleanup(srv.Close)
	qf, err := NewQueueFlowWithOptions(Options{BaseURL: srv.URL, Token: "tenant", WorkerToken: "worker-token"})
	if err != nil {
		t.Fatal(err)
	}
	if opts.Logf == nil {
		opts.Logf = t.Logf
	}
	ctx, cancel := context.WithCancel(context.Background())
	errc := make(chan error, 1)
	go func() { errc <- qf.Run(ctx, "q", handlers, opts) }()
	stop := func() error {
		cancel()
		select {
		case err := <-errc:
			return err
		case <-time.After(10 * time.Second):
			t.Fatal("Run did not return within 10s of cancellation")
			return nil
		}
	}
	t.Cleanup(func() { cancel() })
	return qf, stop
}

// waitUntil polls cond for up to 5 seconds.
func waitUntil(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for %s", what)
}

// The happy path: lease with the worker token and the configured lease/wait,
// run the handler, report its result with the lease token.
func TestRunLeasesRunsAndCompletes(t *testing.T) {
	b := newFakeBroker(t)
	b.enqueue("job-1", "send-email")
	var seen []string
	var mu sync.Mutex
	handlers := map[string]Handler{
		"send-email": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			mu.Lock()
			seen = append(seen, job.Id)
			mu.Unlock()
			return map[string]interface{}{"sent": true, "n": job.Payload["n"]}, nil
		},
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 2})
	waitUntil(t, "completion", func() bool { c, _, _, _ := b.snapshot(); return len(c) == 1 })
	if err := stop(); err != nil {
		t.Fatalf("Run returned %v, want nil after cancellation", err)
	}

	completes, fails, _, leases := b.snapshot()
	if len(fails) != 0 {
		t.Errorf("unexpected fail reports: %v", fails)
	}
	body := completes["job-1"]
	if body["lease_token"] != "tok" {
		t.Errorf("complete.lease_token = %v", body["lease_token"])
	}
	if res := body["result"].(map[string]any); res["sent"] != true || res["n"] != float64(1) {
		t.Errorf("complete.result = %v", res)
	}
	if len(leases) == 0 {
		t.Fatal("no lease call recorded")
	}
	if leases[0]["max_jobs"] != float64(1) || leases[0]["lease_secs"] != float64(10) || leases[0]["wait_secs"] != float64(2) {
		t.Errorf("lease body = %v, want max_jobs 1, lease_secs 10, wait_secs 2", leases[0])
	}
	mu.Lock()
	defer mu.Unlock()
	if len(seen) != 1 || seen[0] != "job-1" {
		t.Errorf("handled = %v", seen)
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if !b.auth["Bearer worker-token"] || b.auth["Bearer tenant"] {
		t.Errorf("worker routes must use the worker token only; saw %v", b.auth)
	}
}

// A plain error is a retryable failure; a NonRetryableError (even wrapped)
// dead-letters; an unknown task fails non-retryably without running anything.
func TestRunReportsFailures(t *testing.T) {
	b := newFakeBroker(t)
	b.enqueue("job-retry", "flaky")
	b.enqueue("job-perm", "bad-input")
	b.enqueue("job-wrapped", "bad-wrapped")
	b.enqueue("job-ghost", "ghost")
	handlers := map[string]Handler{
		"flaky": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			return nil, errors.New("upstream 503")
		},
		"bad-input": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			return nil, NonRetryable(errors.New("malformed order"))
		},
		"bad-wrapped": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			return nil, fmt.Errorf("validating: %w", NonRetryable(errors.New("bad id")))
		},
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 1})
	waitUntil(t, "four fail reports", func() bool { _, f, _, _ := b.snapshot(); return len(f) == 4 })
	if err := stop(); err != nil {
		t.Fatal(err)
	}

	completes, fails, _, _ := b.snapshot()
	if len(completes) != 0 {
		t.Errorf("unexpected completions: %v", completes)
	}
	check := func(id string, retryable bool, msg string) {
		f := fails[id]
		if f["lease_token"] != "tok" {
			t.Errorf("%s: lease_token = %v", id, f["lease_token"])
		}
		if f["retryable"] != retryable {
			t.Errorf("%s: retryable = %v, want %v", id, f["retryable"], retryable)
		}
		if !strings.Contains(fmt.Sprint(f["error"]), msg) {
			t.Errorf("%s: error = %v, want it to contain %q", id, f["error"], msg)
		}
	}
	check("job-retry", true, "upstream 503")
	check("job-perm", false, "malformed order")
	check("job-wrapped", false, "bad id")
	check("job-ghost", false, "no handler for task")
}

// A panicking handler is a retryable failure, not a dead worker.
func TestRunRecoversHandlerPanics(t *testing.T) {
	b := newFakeBroker(t)
	b.enqueue("job-1", "boom")
	handlers := map[string]Handler{
		"boom": func(ctx context.Context, job *Job) (map[string]interface{}, error) { panic("kaboom") },
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 1})
	waitUntil(t, "fail report", func() bool { _, f, _, _ := b.snapshot(); return len(f) == 1 })
	if err := stop(); err != nil {
		t.Fatal(err)
	}
	_, fails, _, _ := b.snapshot()
	if f := fails["job-1"]; f["retryable"] != true || !strings.Contains(fmt.Sprint(f["error"]), "kaboom") {
		t.Errorf("fail = %v", f)
	}
}

// Heartbeats go out at half the lease interval with the lease token and the
// lease length; a lost lease (409, or a non-running status) cancels the
// handler's context and suppresses the outcome report.
func TestRunHeartbeatsAndAbandonsLostLeases(t *testing.T) {
	for _, mode := range []string{"409", "cancelled"} {
		t.Run(mode, func(t *testing.T) {
			b := newFakeBroker(t)
			if mode == "409" {
				b.heartbeat409 = true
			} else {
				b.heartbeat = JOBSTATUS_CANCELLED
			}
			b.enqueue("job-1", "slow")
			var ctxErr error
			done := make(chan struct{})
			handlers := map[string]Handler{
				"slow": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
					<-ctx.Done()
					ctxErr = ctx.Err()
					close(done)
					return map[string]interface{}{"late": true}, nil
				},
			}
			// LeaseSecs 2 => heartbeat every second.
			_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 2, WaitSecs: 1})
			select {
			case <-done:
			case <-time.After(5 * time.Second):
				t.Fatal("handler context was not cancelled after the lost-lease heartbeat")
			}
			if err := stop(); err != nil {
				t.Fatal(err)
			}
			if ctxErr != context.Canceled {
				t.Errorf("handler ctx.Err() = %v, want Canceled", ctxErr)
			}
			completes, fails, heartbeats, _ := b.snapshot()
			if len(heartbeats) == 0 {
				t.Fatal("no heartbeat recorded")
			}
			if heartbeats[0]["lease_token"] != "tok" || heartbeats[0]["extend_secs"] != float64(2) {
				t.Errorf("heartbeat body = %v", heartbeats[0])
			}
			if len(completes) != 0 || len(fails) != 0 {
				t.Errorf("nothing may be reported after a lost lease; got completes=%v fails=%v", completes, fails)
			}
		})
	}
}

// Cancelling Run's context stops leasing but lets in-flight handlers finish
// and report within the drain window.
func TestRunStopDrainsInFlightJobs(t *testing.T) {
	b := newFakeBroker(t)
	b.enqueue("job-1", "slow")
	started := make(chan struct{})
	release := make(chan struct{})
	handlers := map[string]Handler{
		"slow": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			close(started)
			<-release
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			return map[string]interface{}{"drained": true}, nil
		},
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 1, DrainTimeout: 5 * time.Second})
	<-started
	errc := make(chan error, 1)
	go func() { errc <- stop() }()
	select {
	case err := <-errc:
		t.Fatalf("Run returned (%v) before the in-flight handler finished", err)
	case <-time.After(200 * time.Millisecond):
	}
	close(release)
	if err := <-errc; err != nil {
		t.Fatal(err)
	}
	completes, _, _, _ := b.snapshot()
	if completes["job-1"] == nil || completes["job-1"]["result"].(map[string]any)["drained"] != true {
		t.Errorf("drained job was not reported complete: %v", completes)
	}
}

// Past the drain timeout the handler's context is cancelled and Run returns;
// the job is left to the server (nothing reported).
func TestRunDrainTimeoutCancelsHandlers(t *testing.T) {
	b := newFakeBroker(t)
	b.enqueue("job-1", "stuck")
	started := make(chan struct{})
	ctxDone := make(chan struct{})
	handlers := map[string]Handler{
		"stuck": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			close(started)
			<-ctx.Done()
			close(ctxDone)
			return nil, ctx.Err()
		},
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 1, DrainTimeout: 100 * time.Millisecond})
	<-started
	begin := time.Now()
	if err := stop(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(begin); elapsed > 3*time.Second {
		t.Errorf("Run took %s to give up on a 100ms drain timeout", elapsed)
	}
	select {
	case <-ctxDone:
	case <-time.After(time.Second):
		t.Fatal("handler context was not cancelled at the drain timeout")
	}
	time.Sleep(50 * time.Millisecond) // let any (wrong) report land
	completes, fails, _, _ := b.snapshot()
	if len(completes)+len(fails) != 0 {
		t.Errorf("a drain-cancelled job must not be reported; got completes=%v fails=%v", completes, fails)
	}
}

// 401/403 on lease are fatal: Run returns the error instead of retrying.
func TestRunReturnsOnAuthFailure(t *testing.T) {
	for _, status := range []int{http.StatusUnauthorized, http.StatusForbidden} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			b := newFakeBroker(t)
			b.leaseStatus = status
			srv := httptest.NewServer(b)
			t.Cleanup(srv.Close)
			qf := NewQueueFlow(srv.URL, "t")
			errc := make(chan error, 1)
			go func() { errc <- qf.Run(context.Background(), "q", map[string]Handler{}, WorkerOptions{Logf: t.Logf}) }()
			select {
			case err := <-errc:
				if StatusCode(err) != status {
					t.Errorf("Run returned %v, want an APIError with status %d", err, status)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("Run kept retrying a fatal auth error")
			}
		})
	}
}

// Transient lease errors back off one second, are reported through OnError,
// and do not stop the loop.
func TestRunBacksOffOnLeaseErrors(t *testing.T) {
	b := newFakeBroker(t)
	b.leaseFailN = 2
	b.enqueue("job-1", "echo")
	var reported int32
	handlers := map[string]Handler{
		"echo": func(ctx context.Context, job *Job) (map[string]interface{}, error) { return nil, nil },
	}
	begin := time.Now()
	_, stop := runWorker(t, b, handlers, WorkerOptions{WaitSecs: 1, OnError: func(err error) {
		atomic.AddInt32(&reported, 1)
		if StatusCode(err) != 500 {
			t.Errorf("OnError got %v, want a 500", err)
		}
	}})
	waitUntil(t, "completion after backoff", func() bool { c, _, _, _ := b.snapshot(); return len(c) == 1 })
	if err := stop(); err != nil {
		t.Fatal(err)
	}
	if got := atomic.LoadInt32(&reported); got != 2 {
		t.Errorf("OnError called %d times, want 2", got)
	}
	if elapsed := time.Since(begin); elapsed < 2*time.Second {
		t.Errorf("two failures should have cost two 1s backoffs, took %s", elapsed)
	}
}

// With Concurrency N the lease asks for N jobs and runs them at the same time.
func TestRunLeasesAndRunsConcurrently(t *testing.T) {
	b := newFakeBroker(t)
	for i := 1; i <= 3; i++ {
		b.enqueue(fmt.Sprintf("job-%d", i), "barrier")
	}
	var wg sync.WaitGroup
	wg.Add(3)
	allStarted := make(chan struct{})
	go func() {
		wg.Wait()
		close(allStarted)
	}()
	handlers := map[string]Handler{
		"barrier": func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			wg.Done()
			// Only returns once every job is in flight at the same time.
			select {
			case <-allStarted:
				return map[string]interface{}{}, nil
			case <-time.After(3 * time.Second):
				return nil, errors.New("the other handlers never started: not concurrent")
			}
		},
	}
	_, stop := runWorker(t, b, handlers, WorkerOptions{Concurrency: 3, LeaseSecs: 10, WaitSecs: 1})
	waitUntil(t, "three completions", func() bool { c, _, _, _ := b.snapshot(); return len(c) == 3 })
	if err := stop(); err != nil {
		t.Fatal(err)
	}
	_, fails, _, leases := b.snapshot()
	if len(fails) != 0 {
		t.Errorf("handlers did not overlap: %v", fails)
	}
	if leases[0]["max_jobs"] != float64(3) {
		t.Errorf("first lease max_jobs = %v, want 3 (all slots free)", leases[0]["max_jobs"])
	}
}
