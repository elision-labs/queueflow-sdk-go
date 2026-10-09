/*
Live conformance test: runs the facade against a REAL QueueFlow server.

Skipped unless QUEUEFLOW_URL is set. The server must run in `--mode all` with
the built-in `echo` handler registered (the default `queueflow serve`).

	QUEUEFLOW_URL           base URL, e.g. http://localhost:8000
	QUEUEFLOW_TOKEN         tenant bearer token (default "dev")
	QUEUEFLOW_WORKER_TOKEN  worker-protocol token (read for parity with the other
	                        SDKs; this SDK ships no worker runtime, so it is unused)

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
	"testing"
	"time"
)

const (
	liveWait = 60 * time.Second
	livePoll = 250 * time.Millisecond
)

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
	return NewQueueFlow(url, token)
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

func TestLiveServerReachableAndEchoRegistered(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	health, _, err := qf.Client.HealthAPI.GetHealth(ctx).Execute()
	if err != nil {
		t.Fatalf("GetHealth: %v", err)
	}
	if health.Status == "" {
		t.Fatal("GetHealth returned an empty status")
	}
	tasks, _, err := qf.Client.SystemAPI.ListTasks(ctx).Execute()
	if err != nil {
		t.Fatalf("ListTasks: %v", err)
	}
	for _, name := range tasks.Tasks {
		if name == "echo" {
			return
		}
	}
	t.Fatalf("echo handler not registered: %v", tasks.Tasks)
}

func TestLiveEchoJobCompletesAndEchoesPayload(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	payload := map[string]interface{}{"hello": "world", "n": 42, "nested": map[string]interface{}{"ok": true}}
	key := "sdk-live-" + liveSuffix(t)

	// The facade's CreateJob has no idempotency-key parameter; use the
	// generated client for the create and the facade for everything else.
	req := CreateJobRequest{TaskName: "echo", Payload: payload}
	created, _, err := qf.Client.JobsAPI.CreateJob(ctx).CreateJobRequest(req).IdempotencyKey(key).Execute()
	if err != nil {
		t.Fatalf("CreateJob: %v", err)
	}

	fetched, _, err := qf.Client.JobsAPI.GetJob(ctx, created.JobId).Execute()
	if err != nil {
		t.Fatalf("GetJob: %v", err)
	}
	if fetched.TaskName != "echo" {
		t.Fatalf("task_name = %q, want echo", fetched.TaskName)
	}
	if got := fetched.IdempotencyKey.Get(); got == nil || *got != key {
		t.Fatalf("idempotency_key = %v, want %q", got, key)
	}

	done, err := qf.WaitForJob(created.JobId, liveWait, livePoll)
	if err != nil {
		t.Fatalf("WaitForJob: %v", err)
	}
	if done.Status != "completed" {
		msg := ""
		if m := done.ErrorMessage.Get(); m != nil {
			msg = *m
		}
		t.Fatalf("status = %q, want completed (error: %s)", done.Status, msg)
	}
	// The built-in echo handler returns the payload plus `echoed: true`.
	want := map[string]interface{}{"hello": "world", "n": 42, "nested": map[string]interface{}{"ok": true}, "echoed": true}
	if got, exp := canonical(t, done.Result), canonical(t, want); got != exp {
		t.Fatalf("result = %s, want %s", got, exp)
	}

	// Re-submitting the same idempotency key returns the original job.
	again, _, err := qf.Client.JobsAPI.CreateJob(ctx).CreateJobRequest(req).IdempotencyKey(key).Execute()
	if err != nil {
		t.Fatalf("CreateJob (replay): %v", err)
	}
	if again.JobId != created.JobId {
		t.Fatalf("idempotent replay returned job %q, want %q", again.JobId, created.JobId)
	}
}

func TestLiveTwoStepEchoWorkflowCompletes(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	dag := NewWorkflowBuilder("sdk-live-wf-"+liveSuffix(t)).
		Step("first", "echo").
		Step("second", "echo", "first")
	wf, err := qf.CreateWorkflow(dag)
	if err != nil {
		t.Fatalf("CreateWorkflow: %v", err)
	}
	if len(wf.Steps) != 2 {
		t.Fatalf("steps = %d, want 2", len(wf.Steps))
	}

	finished, err := qf.WaitForWorkflow(wf.Id, liveWait, livePoll)
	if err != nil {
		t.Fatalf("WaitForWorkflow: %v", err)
	}
	if finished.Status != "completed" {
		t.Fatalf("workflow status = %q, want completed", finished.Status)
	}

	states, _, err := qf.Client.WorkflowsAPI.GetWorkflowStepStates(ctx, wf.Id).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowStepStates: %v", err)
	}
	if len(states.Steps) != 2 {
		t.Fatalf("step states = %d, want 2", len(states.Steps))
	}
	for _, s := range states.Steps {
		if s.Status != "completed" {
			t.Fatalf("step %q status = %q, want completed", s.Name, s.Status)
		}
	}
}

func TestLiveCronCreateListPauseResumeDelete(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()
	name := "sdk-live-cron-" + liveSuffix(t)

	// Fires once a year; the schedule never triggers during the test.
	id, err := qf.CreateCron(name, "0 0 1 1 *", "echo", map[string]interface{}{"from": "cron"})
	if err != nil {
		t.Fatalf("CreateCron: %v", err)
	}

	cron, _, err := qf.Client.CronAPI.GetCron(ctx, id).Execute()
	if err != nil {
		t.Fatalf("GetCron: %v", err)
	}
	if cron.Name != name || !cron.Enabled {
		t.Fatalf("cron = %+v, want name %q enabled", cron, name)
	}

	listed, _, err := qf.Client.CronAPI.ListCrons(ctx).Limit(100).Execute()
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

	if _, err := qf.Client.CronAPI.PauseCron(ctx, id).Execute(); err != nil {
		t.Fatalf("PauseCron: %v", err)
	}
	paused, _, err := qf.Client.CronAPI.GetCron(ctx, id).Execute()
	if err != nil {
		t.Fatalf("GetCron after pause: %v", err)
	}
	if paused.Enabled {
		t.Fatal("after pause: cron still enabled")
	}

	if _, err := qf.Client.CronAPI.ResumeCron(ctx, id).Execute(); err != nil {
		t.Fatalf("ResumeCron: %v", err)
	}
	resumed, _, err := qf.Client.CronAPI.GetCron(ctx, id).Execute()
	if err != nil {
		t.Fatalf("GetCron after resume: %v", err)
	}
	if !resumed.Enabled {
		t.Fatal("after resume: cron still disabled")
	}

	if _, err := qf.Client.CronAPI.DeleteCron(ctx, id).Execute(); err != nil {
		t.Fatalf("DeleteCron: %v", err)
	}
	_, resp, err := qf.Client.CronAPI.GetCron(ctx, id).Execute()
	if err == nil || resp == nil || resp.StatusCode != 404 {
		t.Fatalf("GetCron after delete: err=%v resp=%v, want 404", err, resp)
	}
}

func TestLiveDeadLetterListAndStats(t *testing.T) {
	qf := liveClient(t)
	ctx := context.Background()

	dlq, _, err := qf.Client.DlqAPI.ListDeadLetters(ctx).Limit(10).Execute()
	if err != nil {
		t.Fatalf("ListDeadLetters: %v", err)
	}
	if dlq.DeadLetters == nil {
		t.Fatal("dead_letters is nil, want a (possibly empty) list")
	}

	stats, _, err := qf.Client.SystemAPI.GetStats(ctx).Execute()
	if err != nil {
		t.Fatalf("GetStats: %v", err)
	}
	if stats.JobsCreated < 1 {
		t.Fatalf("jobs_created = %d, want >= 1 (this suite created jobs)", stats.JobsCreated)
	}
}
