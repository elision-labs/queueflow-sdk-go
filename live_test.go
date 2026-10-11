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

func TestLiveWhoAmIReportsTheTenant(t *testing.T) {
	qf := liveClient(t)
	me, err := qf.WhoAmI(context.Background())
	if err != nil {
		t.Fatalf("WhoAmI: %v", err)
	}
	if me.Tenant == "" {
		t.Fatal("WhoAmI returned an empty tenant")
	}
	// The rest of this suite creates jobs, so the configured token must be a
	// read-write one.
	if me.Access != "read_write" {
		t.Fatalf("access = %q, want read_write", me.Access)
	}
}

// Bulk cancel is scoped to this run's queue (the server is shared), filtered
// by task and status, and refuses terminal statuses.
func TestLiveCancelJobsByQueueTaskAndStatus(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	suffix := liveSuffix(t)
	queue := "go-live-cancel-" + suffix
	taskA := "go-live-cancel-a-" + suffix
	taskB := "go-live-cancel-b-" + suffix

	// Far-future jobs stay pending (nothing claims them) until cancelled.
	runAt := time.Now().Add(24 * time.Hour)
	var ids []string
	for _, task := range []string{taskA, taskA, taskB} {
		id, err := qf.CreateJobWithOptions(ctx, task, CreateJobOptions{Queue: queue, RunAt: &runAt})
		if err != nil {
			t.Fatalf("CreateJobWithOptions(%s): %v", task, err)
		}
		ids = append(ids, id)
	}
	defer func() { _, _ = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue}) }()

	listed, err := qf.ListJobs(ctx, ListOptions{Queue: queue, TaskName: taskA, Limit: 10})
	if err != nil {
		t.Fatalf("ListJobs: %v", err)
	}
	if len(listed.Jobs) != 2 {
		t.Fatalf("ListJobs(task %s) = %d jobs, want 2", taskA, len(listed.Jobs))
	}

	if _, err := qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue, Status: "completed"}); StatusCode(err) != 400 {
		t.Fatalf("CancelJobs(status completed): %v, want a 400 APIError", err)
	}

	// A status filter that matches nothing on this queue cancels nothing.
	n, err := qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue, Status: "running"})
	if err != nil {
		t.Fatalf("CancelJobs(running): %v", err)
	}
	if n != 0 {
		t.Fatalf("CancelJobs(running) = %d, want 0", n)
	}

	n, err = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue, TaskName: taskA, Limit: 1})
	if err != nil {
		t.Fatalf("CancelJobs(task A, limit 1): %v", err)
	}
	if n != 1 {
		t.Fatalf("CancelJobs(task A, limit 1) = %d, want 1", n)
	}
	// Exactly one task-A job went; task B was never a candidate.
	status := func(id string) JobStatus {
		job, err := qf.GetJob(ctx, id)
		if err != nil {
			t.Fatalf("GetJob(%s): %v", id, err)
		}
		return job.Status
	}
	cancelledA := 0
	for _, id := range ids[:2] {
		if status(id) == JOBSTATUS_CANCELLED {
			cancelledA++
		}
	}
	if cancelledA != 1 {
		t.Fatalf("task-A jobs cancelled = %d, want 1", cancelledA)
	}
	if got := status(ids[2]); got != JOBSTATUS_PENDING {
		t.Fatalf("task-B job status = %q, want pending", got)
	}

	n, err = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue, Status: "pending"})
	if err != nil {
		t.Fatalf("CancelJobs(pending): %v", err)
	}
	if n != 2 {
		t.Fatalf("CancelJobs(pending) = %d, want the remaining 2", n)
	}
	for _, id := range ids {
		if got := status(id); got != JOBSTATUS_CANCELLED {
			t.Fatalf("job %s status = %q, want cancelled", id, got)
		}
	}
	n, err = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue})
	if err != nil {
		t.Fatalf("CancelJobs(exhausted): %v", err)
	}
	if n != 0 {
		t.Fatalf("CancelJobs on an exhausted queue = %d, want 0", n)
	}
}

// Bulk replay re-runs every unreplayed dead letter on this run's queue once;
// the replayed flag then separates them in the dead-letter list.
func TestLiveReplayDeadLettersByQueue(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	suffix := liveSuffix(t)
	queue := "go-live-replay-" + suffix
	task := "go-live-replay-task-" + suffix

	var ids []string
	for i := 0; i < 2; i++ {
		id, err := qf.CreateJobWithOptions(ctx, task, CreateJobOptions{Queue: queue, Payload: map[string]interface{}{"i": i}})
		if err != nil {
			t.Fatalf("CreateJobWithOptions: %v", err)
		}
		ids = append(ids, id)
	}
	// Whatever the replay enqueues on this queue is cleaned up.
	defer func() { _, _ = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue}) }()

	handlers := map[string]Handler{
		task: func(ctx context.Context, job *Job) (map[string]interface{}, error) {
			return nil, NonRetryable(fmt.Errorf("poison %s", job.Id))
		},
	}
	workerCtx, stop := context.WithCancel(ctx)
	errc := make(chan error, 1)
	go func() {
		errc <- qf.Run(workerCtx, queue, handlers, WorkerOptions{Concurrency: 2, LeaseSecs: 10, WaitSecs: 2, Logf: t.Logf})
	}()
	for _, id := range ids {
		done, err := qf.WaitForJob(ctx, id, liveWait)
		if err != nil {
			stop()
			t.Fatalf("WaitForJob(%s): %v", id, err)
		}
		if done.Status != JOBSTATUS_FAILED {
			stop()
			t.Fatalf("job %s status = %q, want failed", id, done.Status)
		}
	}
	stop()
	if err := <-errc; err != nil {
		t.Fatalf("Run: %v", err)
	}

	notReplayed := false
	pending, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, Replayed: &notReplayed, Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters(replayed=false): %v", err)
	}
	if len(pending.DeadLetters) != 2 {
		t.Fatalf("unreplayed dead letters on %s = %d, want 2", queue, len(pending.DeadLetters))
	}
	byTask, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, TaskName: task, Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters(task): %v", err)
	}
	if len(byTask.DeadLetters) != 2 {
		t.Fatalf("dead letters for task %s = %d, want 2", task, len(byTask.DeadLetters))
	}
	other, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, TaskName: task + "-none", Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters(other task): %v", err)
	}
	if len(other.DeadLetters) != 0 {
		t.Fatalf("dead letters for an unknown task = %d, want 0", len(other.DeadLetters))
	}

	out, err := qf.ReplayDeadLetters(ctx, ReplayDeadLettersFilter{Queue: queue, TaskName: task})
	if err != nil {
		t.Fatalf("ReplayDeadLetters: %v", err)
	}
	if out.Replayed != 2 || len(out.Results) != 2 {
		t.Fatalf("ReplayDeadLetters = %+v, want 2 replayed", out)
	}
	original := map[string]bool{ids[0]: true, ids[1]: true}
	for _, r := range out.Results {
		newID := r.GetJobId()
		if newID == "" || original[newID] || r.Error.IsSet() {
			t.Fatalf("result %+v, want a fresh job id and no error", r)
		}
	}

	// Already-replayed entries are not considered again.
	again, err := qf.ReplayDeadLetters(ctx, ReplayDeadLettersFilter{Queue: queue})
	if err != nil {
		t.Fatalf("ReplayDeadLetters (second): %v", err)
	}
	if again.Replayed != 0 || len(again.Results) != 0 {
		t.Fatalf("second ReplayDeadLetters = %+v, want nothing", again)
	}
	replayed := true
	doneList, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, Replayed: &replayed, Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters(replayed=true): %v", err)
	}
	if len(doneList.DeadLetters) != 2 {
		t.Fatalf("replayed dead letters on %s = %d, want 2", queue, len(doneList.DeadLetters))
	}
	left, err := qf.ListDeadLetters(ctx, ListOptions{Queue: queue, Replayed: &notReplayed, Limit: 10})
	if err != nil {
		t.Fatalf("ListDeadLetters(replayed=false, after): %v", err)
	}
	if len(left.DeadLetters) != 0 {
		t.Fatalf("unreplayed dead letters after replay = %d, want 0", len(left.DeadLetters))
	}
}

// A paused queue hands out no leases and shows paused in the backlog view;
// resuming makes the job claimable again.
func TestLivePauseAndResumeQueueGateLeases(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	suffix := liveSuffix(t)
	queue := "go-live-pause-" + suffix
	task := "go-live-pause-task-" + suffix

	id, err := qf.CreateJobWithOptions(ctx, task, CreateJobOptions{Queue: queue})
	if err != nil {
		t.Fatalf("CreateJobWithOptions: %v", err)
	}
	defer func() {
		_, _ = qf.ResumeQueue(ctx, queue)
		_, _ = qf.CancelJobs(ctx, CancelJobsFilter{Queue: queue})
	}()

	paused, err := qf.PauseQueue(ctx, queue)
	if err != nil {
		t.Fatalf("PauseQueue: %v", err)
	}
	if !paused.Paused || paused.Queue != queue || paused.Pending != 1 || paused.Running != 0 {
		t.Fatalf("PauseQueue = %+v, want paused with 1 pending, 0 running", paused)
	}
	// Idempotent.
	if again, err := qf.PauseQueue(ctx, queue); err != nil || !again.Paused {
		t.Fatalf("second PauseQueue: %+v, %v", again, err)
	}

	leased, err := qf.LeaseJobs(ctx, queue, LeaseOptions{MaxJobs: 1, LeaseSecs: 10, WaitSecs: 1})
	if err != nil {
		t.Fatalf("LeaseJobs while paused: %v", err)
	}
	if len(leased) != 0 {
		t.Fatalf("leased %d jobs from a paused queue, want 0", len(leased))
	}

	queues, err := qf.Queues(ctx)
	if err != nil {
		t.Fatalf("Queues: %v", err)
	}
	seen := false
	for _, qs := range queues {
		if qs.Queue == queue {
			seen = true
			if !qs.GetPaused() {
				t.Fatalf("Queues() reports %s as not paused: %+v", queue, qs)
			}
		}
	}
	if !seen {
		t.Fatalf("queue %s missing from Queues()", queue)
	}

	resumed, err := qf.ResumeQueue(ctx, queue)
	if err != nil {
		t.Fatalf("ResumeQueue: %v", err)
	}
	if resumed.Paused || resumed.Pending != 1 {
		t.Fatalf("ResumeQueue = %+v, want not paused with 1 pending", resumed)
	}
	leased, err = qf.LeaseJobs(ctx, queue, LeaseOptions{MaxJobs: 1, LeaseSecs: 10, WaitSecs: 5})
	if err != nil {
		t.Fatalf("LeaseJobs after resume: %v", err)
	}
	if len(leased) != 1 || leased[0].Job.Id != id {
		t.Fatalf("leased = %+v, want job %s", leased, id)
	}
	if err := qf.CompleteJob(ctx, &leased[0], map[string]interface{}{"ok": true}); err != nil {
		t.Fatalf("CompleteJob: %v", err)
	}
	done, err := qf.GetJob(ctx, id)
	if err != nil {
		t.Fatalf("GetJob: %v", err)
	}
	if done.Status != JOBSTATUS_COMPLETED {
		t.Fatalf("status = %q, want completed", done.Status)
	}
}
