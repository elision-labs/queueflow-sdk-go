/*
Tests for the hand-written facade.

The generated core is covered by codegen and the spec drift guard; these
exercise only the ergonomics injected on top of it: DAG validation that must
fail before any round-trip, the create-then-fetch pairing, and the polling loop
that turns a sequence of responses into a terminal record or a timeout. The
transport is a stdlib httptest server, so nothing here needs a live QueueFlow.
*/

package queueflow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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
	return fmt.Sprintf(`{
		"id":%q,
		"task_name":"echo",
		"status":%q,
		"queue_name":"default",
		"retry_count":0,
		"created_at":"2026-01-01T00:00:00Z",
		"scheduled_at":"2026-01-01T00:00:00Z",
		"config":{"max_retries":3,"priority":0,"retry_delay_secs":1,"retry_max_delay_secs":60,"timeout_secs":300}
	}`, id, status)
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

// --- CreateJob -------------------------------------------------------------

// CreateJob is a create-then-fetch pair: the POST returns only an id, so the
// facade must follow up with a GET and hand back the full record.
func TestCreateJobPostsThenFetchesTheRecord(t *testing.T) {
	var posted map[string]any
	var gotAuth string
	var paths []string

	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.Method+" "+r.URL.Path)
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(r.Body).Decode(&posted); err != nil {
				t.Errorf("decoding request body: %v", err)
			}
			fmt.Fprint(w, `{"job_id":"job-1"}`)
			return
		}
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_PENDING))
	})

	job, err := qf.CreateJob("echo", map[string]interface{}{"hi": 1})
	if err != nil {
		t.Fatalf("CreateJob: %v", err)
	}
	if job.Id != "job-1" {
		t.Errorf("Id = %q, want %q", job.Id, "job-1")
	}
	if posted["task_name"] != "echo" {
		t.Errorf("task_name = %v, want echo", posted["task_name"])
	}
	if payload, ok := posted["payload"].(map[string]any); !ok || payload["hi"] != float64(1) {
		t.Errorf("payload = %v, want {hi: 1}", posted["payload"])
	}
	if gotAuth != "Bearer test-token" {
		t.Errorf("Authorization = %q, want %q", gotAuth, "Bearer test-token")
	}
	if len(paths) != 2 || !strings.HasPrefix(paths[0], "POST") || !strings.HasPrefix(paths[1], "GET") {
		t.Errorf("request sequence = %v, want a POST then a GET", paths)
	}
}

// A nil payload must stay absent rather than serializing as null, which the
// server would reject as a type error on an object field.
func TestCreateJobOmitsNilPayload(t *testing.T) {
	var posted map[string]any
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			json.NewDecoder(r.Body).Decode(&posted)
			fmt.Fprint(w, `{"job_id":"job-1"}`)
			return
		}
		fmt.Fprint(w, jobJSON("job-1", JOBSTATUS_PENDING))
	})

	if _, err := qf.CreateJob("echo", nil); err != nil {
		t.Fatalf("CreateJob: %v", err)
	}
	if v, present := posted["payload"]; present && v == nil {
		t.Error("payload serialized as null; want the key omitted")
	}
}

func TestCreateJobPropagatesServerError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"unknown task"}`)
	})

	job, err := qf.CreateJob("nope", nil)
	if err == nil {
		t.Fatalf("expected an error, got job %+v", job)
	}
	if job != nil {
		t.Errorf("job = %+v, want nil alongside the error", job)
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

	job, err := qf.WaitForJob("job-1", 5*time.Second, time.Millisecond)
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
			job, err := qf.WaitForJob("job-1", time.Second, time.Millisecond)
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
	job, err := qf.WaitForJob("job-1", 50*time.Millisecond, 10*time.Millisecond)
	if err == nil {
		t.Fatalf("expected a timeout, got job %+v", job)
	}
	if !strings.Contains(err.Error(), "did not finish within") {
		t.Errorf("error = %q, want a timeout message", err)
	}
	// The loop must give up near the deadline rather than sleeping past it.
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("took %s to time out on a 50ms deadline", elapsed)
	}
}

func TestWaitForJobPropagatesFetchError(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	if _, err := qf.WaitForJob("job-1", time.Second, time.Millisecond); err == nil {
		t.Fatal("expected the fetch error to propagate, got nil")
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
	if _, err := qf.CreateWorkflow(cyclic); err == nil {
		t.Fatal("expected a validation error for a cyclic DAG")
	}
	if got := atomic.LoadInt32(&called); got != 0 {
		t.Errorf("sent %d requests, want 0 for a locally-invalid DAG", got)
	}
}

func TestCreateWorkflowPostsThenFetchesTheRecord(t *testing.T) {
	var paths []string
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.Method)
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			fmt.Fprint(w, `{"workflow_id":"wf-1"}`)
			return
		}
		fmt.Fprint(w, workflowJSON("wf-1", WORKFLOWSTATUS_CREATED))
	})

	got, err := qf.CreateWorkflow(NewWorkflowBuilder("etl").Step("a", "task"))
	if err != nil {
		t.Fatalf("CreateWorkflow: %v", err)
	}
	if got.Id != "wf-1" {
		t.Errorf("Id = %q, want %q", got.Id, "wf-1")
	}
	if len(paths) != 2 || paths[0] != http.MethodPost || paths[1] != http.MethodGet {
		t.Errorf("request sequence = %v, want [POST GET]", paths)
	}
}

// partially_failed is terminal for a workflow but has no job equivalent, so it
// is the case most likely to be missed in the terminal set.
func TestWaitForWorkflowTreatsPartiallyFailedAsTerminal(t *testing.T) {
	qf := newTestQueueFlow(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, workflowJSON("wf-1", WORKFLOWSTATUS_PARTIALLY_FAILED))
	})

	got, err := qf.WaitForWorkflow("wf-1", time.Second, time.Millisecond)
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

	if _, err := qf.WaitForWorkflow("wf-1", 50*time.Millisecond, 10*time.Millisecond); err == nil {
		t.Fatal("expected a timeout while the workflow is still running")
	}
}
