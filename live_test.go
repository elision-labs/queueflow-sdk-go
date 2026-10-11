/*
Live conformance test: runs the facade against a REAL QueueFlow server.

Skipped unless QUEUEFLOW_URL is set. The server must run in `--mode all` with
the built-in `echo` handler registered (the default `queueflow serve`).

	QUEUEFLOW_URL           base URL, e.g. http://localhost:8000
	QUEUEFLOW_TOKEN         tenant bearer token (default "dev")
	QUEUEFLOW_WORKER_TOKEN  worker-protocol token (defaults to QUEUEFLOW_TOKEN,
	                        which only works when the server has no --worker-token)

Run with `go test -run Live ./...`. Plain `go test ./...` collects it too, but
it skips itself, so CI stays offline.
*/

package queueflow

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

var liveWait = WaitOptions{Timeout: 60 * time.Second, Interval: 250 * time.Millisecond}

// liveClient returns a client for the live server, or skips the test.
func liveClient(t *testing.T) *QueueFlow {
	t.Helper()
	url := os.Getenv("QUEUEFLOW_URL")
	if url == "" {
		t.Skip("QUEUEFLOW_URL is not set")
	}
	token := os.Getenv("QUEUEFLOW_TOKEN")
	if token == "" {
		token = "dev"
	}
	qf, err := NewQueueFlowWithOptions(Options{
		BaseURL:     url,
		Token:       token,
		WorkerToken: os.Getenv("QUEUEFLOW_WORKER_TOKEN"),
	})
	if err != nil {
		t.Fatal(err)
	}
	return qf
}

func liveSuffix(t *testing.T) string {
	t.Helper()
	b := make([]byte, 3)
	if _, err := rand.Read(b); err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x-%s", time.Now().Unix(), hex.EncodeToString(b))
}

// canonical renders a JSON-ish value with sorted keys so maps decoded from the
// wire (float64 numbers) compare equal to literals.
func canonical(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func errorMessage(job *Job) string {
	if m := job.ErrorMessage.Get(); m != nil {
		return *m
	}
	return ""
}

func TestLiveServerReachableAndEchoRegistered(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	health, err := qf.Health(ctx)
	if err != nil {
		t.Fatalf("Health: %v", err)
	}
	if health.Status == "" {
		t.Fatal("Health returned an empty status")
	}
	if _, err := qf.Ready(ctx); err != nil {
		t.Fatalf("Ready: %v", err)
	}
	tasks, err := qf.Tasks(ctx)
	if err != nil {
		t.Fatalf("Tasks: %v", err)
	}
	for _, name := range tasks {
		if name == "echo" {
			return
		}
	}
	t.Fatalf("echo handler not registered: %v", tasks)
}

func TestLiveEchoJobCompletesAndEchoesPayload(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	payload := map[string]interface{}{"hello": "world", "n": 42, "nested": map[string]interface{}{"ok": true}}
	key := "sdk-live-" + liveSuffix(t)

	id, err := qf.CreateJobWithOptions(ctx, "echo", CreateJobOptions{Payload: payload, IdempotencyKey: key})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}

	fetched, err := qf.GetJob(ctx, id)
	if err != nil {
		t.Fatalf("GetJob: %v", err)
	}
	if fetched.TaskName != "echo" {
		t.Fatalf("task_name = %q, want echo", fetched.TaskName)
	}
	if got := fetched.IdempotencyKey.Get(); got == nil || *got != key {
		t.Fatalf("idempotency_key = %v, want %q", got, key)
	}

	done, err := qf.WaitForJob(ctx, id, liveWait)
	if err != nil {
		t.Fatalf("WaitForJob: %v", err)
	}
	if done.Status != JOBSTATUS_COMPLETED {
		t.Fatalf("status = %q, want completed (error: %s)", done.Status, errorMessage(done))
	}
	// The built-in echo handler returns the payload plus `echoed: true`.
	want := map[string]interface{}{"hello": "world", "n": 42, "nested": map[string]interface{}{"ok": true}, "echoed": true}
	if got, exp := canonical(t, done.Result), canonical(t, want); got != exp {
		t.Fatalf("result = %s, want %s", got, exp)
	}

	// Re-submitting the same idempotency key returns the original job.
	again, err := qf.CreateJobWithOptions(ctx, "echo", CreateJobOptions{Payload: payload, IdempotencyKey: key})
	if err != nil {
		t.Fatalf("CreateJobWithOptions (replay): %v", err)
	}
	if again != id {
		t.Fatalf("idempotent replay returned job %q, want %q", again, id)
	}

	// The job shows up in a filtered list.
	listed, err := qf.ListJobs(ctx, ListOptions{Status: "completed", Limit: 100, OrderBy: "created_at DESC"})
	if err != nil {
		t.Fatalf("ListJobs: %v", err)
	}
	found := false
	for _, j := range listed.Jobs {
		if j.Id == id {
			found = true
		}
	}
	if !found {
		t.Fatalf("job %s missing from the completed list", id)
	}
}

func TestLiveTwoStepEchoWorkflowCompletes(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	dag := NewWorkflowBuilder("sdk-live-wf-"+liveSuffix(t)).
		StepWithOptions("first", "echo", StepOptions{Payload: map[string]interface{}{"step": 1}}).
		StepWithOptions("second", "echo", StepOptions{After: []string{"first"}, Payload: map[string]interface{}{"step": 2}})
	id, err := qf.CreateWorkflow(ctx, dag)
	if err != nil {
		t.Fatalf("CreateWorkflow: %v", err)
	}
	wf, err := qf.GetWorkflow(ctx, id)
	if err != nil {
		t.Fatalf("GetWorkflow: %v", err)
	}
	if len(wf.Steps) != 2 {
		t.Fatalf("steps = %d, want 2", len(wf.Steps))
	}

	finished, err := qf.WaitForWorkflow(ctx, id, liveWait)
	if err != nil {
		t.Fatalf("WaitForWorkflow: %v", err)
	}
	if finished.Status != WORKFLOWSTATUS_COMPLETED {
		t.Fatalf("workflow status = %q, want completed", finished.Status)
	}

	states, err := qf.WorkflowSteps(ctx, id)
	if err != nil {
		t.Fatalf("WorkflowSteps: %v", err)
	}
	if len(states) != 2 {
		t.Fatalf("step states = %d, want 2", len(states))
	}
	for _, s := range states {
		if s.Status != STEPSTATUS_COMPLETED {
			t.Fatalf("step %q status = %q, want completed", s.Name, s.Status)
		}
	}

	diagram, err := qf.WorkflowDiagram(ctx, id)
	if err != nil {
		t.Fatalf("WorkflowDiagram: %v", err)
	}
	if diagram.Diagram == "" {
		t.Fatal("WorkflowDiagram returned an empty diagram")
	}
}

func TestLiveCronCreateListPauseResumeDelete(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	name := "sdk-live-cron-" + liveSuffix(t)

	// Fires once a year; the schedule never triggers during the test.
	id, err := qf.CreateCron(ctx, name, "0 0 1 1 *", "echo", map[string]interface{}{"from": "cron"})
	if err != nil {
		t.Fatalf("CreateCron: %v", err)
	}

	cron, err := qf.GetCron(ctx, id)
	if err != nil {
		t.Fatalf("GetCron: %v", err)
	}
	if cron.Name != name || !cron.Enabled {
		t.Fatalf("cron = %+v, want name %q enabled", cron, name)
	}

	listed, err := qf.ListCrons(ctx, ListOptions{Limit: 100})
	if err != nil {
		t.Fatalf("ListCrons: %v", err)
	}
	found := false
	for _, c := range listed.Crons {
		if c.Id == id {
			found = true
		}
	}
	if !found {
		t.Fatal("new cron missing from list")
	}

	if err := qf.PauseCron(ctx, id); err != nil {
		t.Fatalf("PauseCron: %v", err)
	}
	if paused, err := qf.GetCron(ctx, id); err != nil || paused.Enabled {
		t.Fatalf("after pause: err=%v enabled=%v", err, paused != nil && paused.Enabled)
	}

	if err := qf.ResumeCron(ctx, id); err != nil {
		t.Fatalf("ResumeCron: %v", err)
	}
	if resumed, err := qf.GetCron(ctx, id); err != nil || !resumed.Enabled {
		t.Fatalf("after resume: err=%v enabled=%v", err, resumed != nil && resumed.Enabled)
	}

	if err := qf.DeleteCron(ctx, id); err != nil {
		t.Fatalf("DeleteCron: %v", err)
	}
	if _, err := qf.GetCron(ctx, id); StatusCode(err) != 404 {
		t.Fatalf("GetCron after delete: %v, want a 404 APIError", err)
	}
}

func TestLiveDeadLetterListQueuesAndStats(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	dlq, err := qf.ListDeadLetters(ctx, ListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	if dlq.DeadLetters == nil {
		t.Fatal("dead_letters is nil, want a (possibly empty) list")
	}

	stats, err := qf.Stats(ctx)
	if err != nil {
		t.Fatalf("Stats: %v", err)
	}
	if stats.JobsCreated < 1 {
		t.Fatalf("jobs_created = %d, want >= 1 (this suite created jobs)", stats.JobsCreated)
	}

	// A far-future job keeps its queue in the backlog view while we look.
	queue := "sdk-live-backlog-" + liveSuffix(t)
	runAt := time.Now().Add(24 * time.Hour)
	id, err := qf.CreateJobWithOptions(ctx, "echo", CreateJobOptions{Queue: queue, RunAt: &runAt})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}
	defer func() { _ = qf.CancelJob(ctx, id) }()
	queues, err := qf.Queues(ctx)
	if err != nil {
		t.Fatalf("Queues: %v", err)
	}
	for _, qs := range queues {
		if qs.Queue == queue {
			if qs.Scheduled != 1 {
				t.Fatalf("queue %s scheduled = %d, want 1", queue, qs.Scheduled)
			}
			return
		}
	}
	t.Fatalf("queue %s missing from Queues(): %+v", queue, queues)
}

// The worker runtime: enqueue on a queue the server's in-process workers never
// poll, with a task name only this test's worker knows, so the job can only be
// completed by Run.
func TestLiveWorkerRuntimeLeasesAndCompletesAJob(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	suffix := liveSuffix(t)
	queue := "sdk-live-q-" + suffix
	task := "sdk-live-task-" + suffix
	payload := map[string]interface{}{"order": 7}

	id, err := qf.CreateJobWithOptions(ctx, task, CreateJobOptions{Payload: payload, Queue: queue})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}
	job, err := qf.GetJob(ctx, id)
	if err != nil {
		t.Fatalf("GetJob: %v", err)
	}
	if job.QueueName != queue || job.Status != JOBSTATUS_PENDING {
		t.Fatalf("job = queue %q status %q, want %q pending", job.QueueName, job.Status, queue)
	}

	var mu sync.Mutex
	var handled []string
	handlers := map[string]Handler{
		task: func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			mu.Lock()
			handled = append(handled, job.Id)
			mu.Unlock()
			out := map[string]interface{}{"handled": true}
			for k, v := range job.Payload {
				out[k] = v
			}
			return out, nil
		},
	}
	workerCtx, stop := context.WithCancel(ctx)
	errc := make(chan error, 1)
	go func() {
		errc <- qf.Run(workerCtx, queue, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 2, Logf: t.Logf})
	}()

	done, waitErr := qf.WaitForJob(ctx, id, liveWait)
	stop()
	if err := <-errc; err != nil {
		t.Fatalf("Run: %v", err)
	}
	if waitErr != nil {
		t.Fatalf("WaitForJob: %v", waitErr)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(handled) != 1 || handled[0] != id {
		t.Fatalf("handled = %v, want [%s]", handled, id)
	}
	if done.Status != JOBSTATUS_COMPLETED {
		t.Fatalf("status = %q, want completed (error: %s)", done.Status, errorMessage(done))
	}
	want := map[string]interface{}{"order": 7, "handled": true}
	if got, exp := canonical(t, done.Result), canonical(t, want); got != exp {
		t.Fatalf("result = %s, want %s", got, exp)
	}
}

// A NonRetryable error dead-letters immediately; the entry can be read and
// replayed exactly once.
func TestLiveWorkerNonRetryableFailureDeadLetters(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	suffix := liveSuffix(t)
	queue := "sdk-live-dlq-q-" + suffix
	task := "sdk-live-dlq-task-" + suffix

	id, err := qf.CreateJobWithOptions(ctx, task, CreateJobOptions{Queue: queue, MaxRetries: PtrInt32(5)})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}
	handlers := map[string]Handler{
		task: func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			return nil, NonRetryable(fmt.Errorf("bad input for %s", job.Id))
		},
	}
	workerCtx, stop := context.WithCancel(ctx)
	errc := make(chan error, 1)
	go func() {
		errc <- qf.Run(workerCtx, queue, handlers, WorkerOptions{LeaseSecs: 10, WaitSecs: 2, Logf: t.Logf})
	}()
	done, waitErr := qf.WaitForJob(ctx, id, liveWait)
	stop()
	if err := <-errc; err != nil {
		t.Fatalf("Run: %v", err)
	}
	if waitErr != nil {
		t.Fatalf("WaitForJob: %v", waitErr)
	}
	if done.Status != JOBSTATUS_FAILED || done.RetryCount != 0 {
		t.Fatalf("status = %q retry_count = %d, want failed with no retries", done.Status, done.RetryCount)
	}

	dlq, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	if len(dlq.DeadLetters) != 1 || dlq.DeadLetters[0].JobId != id {
		t.Fatalf("dead letters for %s = %+v, want exactly the failed job", queue, dlq.DeadLetters)
	}
	entry, err := qf.GetDeadLetter(ctx, dlq.DeadLetters[0].Id)
	if err != nil {
		t.Fatalf("GetDeadLetter: %v", err)
	}
	replayed, err := qf.ReplayDeadLetter(ctx, entry.Id)
	if err != nil {
		t.Fatalf("ReplayDeadLetter: %v", err)
	}
	if replayed == "" || replayed == id {
		t.Fatalf("replay returned %q, want a fresh job id", replayed)
	}
	defer func() { _ = qf.CancelJob(ctx, replayed) }()
	if _, err := qf.ReplayDeadLetter(ctx, entry.Id); StatusCode(err) != 409 {
		t.Fatalf("second replay: %v, want a 409 APIError", err)
	}
}
