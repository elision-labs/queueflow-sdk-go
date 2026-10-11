# queueflow-sdk-go

Go client for [QueueFlow](https://queueflow.dev), a PostgreSQL-native distributed job queue and
workflow engine.

- **Typed end-to-end**: request/response models mirror the server's OpenAPI 3.1 spec.
- **Ergonomic**: one method per endpoint on `QueueFlow`, `WaitForJob`/`WaitForWorkflow` pollers,
  and a workflow builder with local DAG validation.
- **Worker runtime**: `Run` leases jobs, runs your handlers concurrently, heartbeats them, and
  reports outcomes, with the server's retry policy and dead-letter queue doing the rest.
- **Typed errors**: every non-2xx response is an `*APIError` carrying the status and body.
- **Thin facade over a generated core**: the facade (`facade.go`) adds what codegen cannot; the
  generated `APIClient` and its `*APIService` groups stay available for everything else.

Module path: `github.com/elision-labs/queueflow-sdk-go`, package `queueflow`. Go 1.18 or newer.

## Install

```bash
go get github.com/elision-labs/queueflow-sdk-go
```

## Quick start

```go
package main

import (
	"context"
	"log"

	queueflow "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	ctx := context.Background()
	qf := queueflow.NewQueueFlow("http://localhost:8000", "dev")

	// Enqueue a job and wait for the result.
	id, err := qf.CreateJob(ctx, "echo", map[string]interface{}{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	done, err := qf.WaitForJob(ctx, id, queueflow.WaitOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(done.Status, done.Result) // completed map[echoed:true hello:world]

	// Declare and run a DAG workflow.
	dag := queueflow.NewWorkflowBuilder("etl").
		Step("extract", "echo").
		Step("transform", "echo", "extract").
		Step("load", "echo", "transform")
	wfID, err := qf.CreateWorkflow(ctx, dag)
	if err != nil {
		log.Fatal(err)
	}
	finished, err := qf.WaitForWorkflow(ctx, wfID, queueflow.WaitOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(finished.Status)
}
```

`echo` is a handler built into the server; it returns the payload plus `"echoed": true`.

## API

### Client

```go
qf := queueflow.NewQueueFlow(baseURL, token)

// or, with a separate worker token and a custom timeout:
qf, err := queueflow.NewQueueFlowWithOptions(queueflow.Options{
	BaseURL:     "http://localhost:8000",
	Token:       "dev",            // tenant token (required)
	WorkerToken: "worker-secret",  // for lease/heartbeat/complete/fail and Run (default: Token)
	Timeout:     30 * time.Second, // per-request HTTP timeout (default 30s)
	HTTPClient:  nil,              // optional base *http.Client (transport, proxies)
})
```

Every method takes a `context.Context` first. The per-request timeout applies to every call
except the lease long-poll, which is bounded by `wait_secs + 35s` instead.

| Field | Description |
| --- | --- |
| `qf.Client` | The generated `*APIClient` with the tenant token (`.JobsAPI`, `.WorkflowsAPI`, `.CronAPI`, `.DlqAPI`, `.SystemAPI`, `.HealthAPI`). |
| `qf.WorkerClient` | The generated `*APIClient` with the worker token (`.WorkerAPI`). |

### Jobs

| Method | Description |
| --- | --- |
| `CreateJob(ctx, task, payload) (string, error)` | Enqueue with the server defaults; returns the job id. |
| `CreateJobWithOptions(ctx, task, CreateJobOptions) (string, error)` | Enqueue with config, idempotency key and `RunAt`; returns the job id. |
| `GetJob(ctx, id) (*Job, error)` | Fetch a job. |
| `ListJobs(ctx, ListOptions) (*ListJobsResponse, error)` | List jobs (`Status`, `Queue`, `Limit`, `Offset`, `OrderBy`, `IncludeTotal`, `Cursor`, `CreatedAfter`, `CreatedBefore`). |
| `CancelJob(ctx, id) error` | Cancel a job. |
| `WaitForJob(ctx, id, WaitOptions) (*Job, error)` | Poll until `completed` / `failed` / `cancelled`. |

`CreateJobOptions` is `{Payload, Queue, Priority, MaxRetries, Timeout, RetryBackoff,
RetryDelaySecs, RetryMaxDelaySecs, JitterFactor, IdempotencyKey, RunAt}`. Pointer fields use
the generated helpers (`queueflow.PtrInt32(3)`, `PtrInt64`, `PtrFloat64`); unset fields are not
sent, so the server defaults apply. The same `IdempotencyKey` returns the original job id.

`WaitOptions` is `{Timeout, Interval time.Duration}`; the zero value means 60s and 500ms. Past
the deadline the error wraps `queueflow.ErrTimeout`; a cancelled context returns `ctx.Err()`.

List responses carry `NextCursor` when there are more pages; pass it back as `Cursor` for keyset
pagination (cheaper than deep `Offset`).

### Workflows

| Method | Description |
| --- | --- |
| `CreateWorkflow(ctx, *WorkflowBuilder) (string, error)` | Validate the DAG locally, create it, return the workflow id. |
| `GetWorkflow(ctx, id)` / `ListWorkflows(ctx, ListOptions)` / `CancelWorkflow(ctx, id)` | Fetch / list / cancel. |
| `WorkflowDiagram(ctx, id) (*WorkflowDiagramResponse, error)` | Mermaid (`graph TD`) diagram of the DAG. |
| `WorkflowSteps(ctx, id) ([]WorkflowStepState, error)` | Runtime status of every step, with the job id executing it. |
| `WaitForWorkflow(ctx, id, WaitOptions) (*Workflow, error)` | Poll until a terminal workflow state. |

### Workflow builder

```go
dag := queueflow.NewWorkflowBuilder("order_123").
	Step("validate", "validate_order").
	StepWithOptions("pay", "process_payment", queueflow.StepOptions{
		After:     []string{"validate"},
		Payload:   map[string]interface{}{"amount": 42},
		OnFailure: queueflow.ONFAILURE_HALT,
		Config:    &queueflow.JobConfig{MaxRetries: queueflow.PtrInt32(5)},
	}).
	Step("ship", "create_shipment", "pay").
	Context(map[string]interface{}{"source": "web"})
body, err := dag.Build() // duplicate names, dangling deps and cycles fail here, before any request
```

`Step(name, task string, after ...string)` adds a step gated on the named `after` steps;
`StepWithOptions` adds `Payload`, `Config`, `OnFailure`, `OnSuccess` and `Metadata`.
`Context` and `Metadata` seed the workflow-level fields. `CreateWorkflow` calls `Build()` for
you.

### Worker

Run task handlers in this process against a remote QueueFlow server:

```go
handlers := map[string]queueflow.Handler{
	"send-email": func(ctx context.Context, job *queueflow.Job) (map[string]interface{}, error) {
		// ctx is cancelled if the job's lease is lost (cancelled mid-run or
		// reclaimed) and when the drain timeout passes after a stop.
		if job.Payload["to"] == nil {
			return nil, queueflow.NonRetryable(errors.New("no recipient")) // dead-letters immediately
		}
		if err := sendEmail(ctx, job.Payload); err != nil {
			return nil, err // retryable: the server applies the job's retry policy
		}
		return map[string]interface{}{"sent": true}, nil
	},
}

ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
defer stop()
err := qf.Run(ctx, "default", handlers, queueflow.WorkerOptions{
	Concurrency:  4,                // jobs in flight at once (default 1)
	LeaseSecs:    30,               // heartbeat at half this interval (default 30)
	WaitSecs:     20,               // server-side long poll when idle (default 20, max 30)
	DrainTimeout: 30 * time.Second, // wait for in-flight handlers after ctx is cancelled (default 30s)
})
```

`Run` leases up to `Concurrency` jobs per call (one per free slot), runs each handler in its
own goroutine, heartbeats every job at half the lease interval, and reports `complete` or
`fail`. A heartbeat that shows the lease is lost (HTTP 409, or any status other than
`running`) cancels the handler's context and suppresses the report; the server already owns the
outcome. Jobs whose task has no handler are failed non-retryably, and a panicking handler is
reported as a retryable failure.

When `ctx` is cancelled, `Run` abandons the in-flight lease long-poll, waits up to
`DrainTimeout` for running handlers to finish and report, then returns `nil`. It returns the
`*APIError` when the lease call is rejected with 401 or 403 (a wrong or missing worker token
cannot heal by retrying); other lease errors back off one second and are logged on the first
failure and every 30th after that (`OnError` replaces the log line; `Logf` replaces
`log.Printf`).

Delivery is at-least-once, so handlers must be idempotent. Servers started with
`--worker-token` require `Options.WorkerToken`.

The raw protocol is also exposed: `LeaseJobs(ctx, queue, LeaseOptions)`,
`HeartbeatJob(ctx, lease, extendSecs)`, `CompleteJob(ctx, lease, result)`,
`FailJob(ctx, lease, message, retryable)`.

### Cron

| Method | Description |
| --- | --- |
| `CreateCron(ctx, name, cronExpr, task, payload) (string, error)` | Register a recurring enqueue (5-field crontab, UTC); returns the schedule id. |
| `CreateCronWithOptions(ctx, name, cronExpr, task, CreateCronOptions) (string, error)` | Also set `Queue` and a `Config` override. |
| `GetCron(ctx, id)` / `ListCrons(ctx, ListOptions)` / `DeleteCron(ctx, id)` | Fetch / list / delete. |
| `PauseCron(ctx, id)` / `ResumeCron(ctx, id)` | Stop firings / resume at the next future occurrence. |

### Dead letters

| Method | Description |
| --- | --- |
| `ListDeadLetters(ctx, ListOptions)` / `GetDeadLetter(ctx, id int64)` | Inspect terminally-failed jobs. |
| `ReplayDeadLetter(ctx, id int64) (string, error)` | Re-run one as a fresh job; returns the new job id (at most once; a second replay is a 409). |

### System

| Method | Description |
| --- | --- |
| `Stats(ctx) (*StatsSnapshot, error)` | Engine counters. |
| `Queues(ctx) ([]QueueStats, error)` | Live per-queue backlog: `Pending`, `Scheduled`, `Running`, `OldestPendingAgeSecs`. |
| `Tasks(ctx) ([]string, error)` | Registered task handler names. |
| `Health(ctx)` / `Ready(ctx)` | `GET /health` / `GET /ready`. |

### Everything else: the generated client

Anything the facade does not wrap (batch enqueue, raw request bodies) is one call away on
`qf.Client`. Every operation is a request builder ending in `.Execute()`, which returns
`(result, *http.Response, error)`:

```go
batch, _, err := qf.Client.JobsAPI.CreateBatchJobs(ctx).
	CreateBatchJobsRequest(queueflow.CreateBatchJobsRequest{Jobs: reqs}).
	Execute()
```

Per-endpoint and per-model reference for the generated layer lives in [`docs/`](./docs).

## Authentication

Every request carries `Authorization: Bearer <token>`. Two credentials exist:

- The **tenant token** (`Token`) authenticates the job, workflow, cron, DLQ and system routes.
  On a server started without `--api-keys`, any non-empty token is accepted (development mode).
- The **worker token** (`WorkerToken`) authenticates the worker-protocol routes (lease,
  heartbeat, complete, fail), used by `Run`. Configure it on the server with `--worker-token` /
  `QUEUEFLOW_WORKER_TOKEN`. It is a separate secret, never a tenant token. It defaults to
  `Token`, which only works on a server running without one.

## Error handling

Facade methods return `*APIError` for non-2xx responses, with the status and body attached;
transport failures and context cancellation are returned unwrapped.

```go
job, err := qf.GetJob(ctx, id)
if err != nil {
	var apiErr *queueflow.APIError
	switch {
	case errors.As(err, &apiErr) && apiErr.Status == http.StatusNotFound:
		// no such job
	case errors.As(err, &apiErr):
		log.Printf("%s failed: HTTP %d: %s", apiErr.Op, apiErr.Status, apiErr.Body)
	case errors.Is(err, context.Canceled):
		// caller gave up
	default:
		return err // connection refused, timeout, decode error
	}
}
```

`queueflow.StatusCode(err)` returns the status of an `*APIError` anywhere in the chain (0
otherwise). `WaitForJob` and `WaitForWorkflow` wrap `queueflow.ErrTimeout` past their deadline.
`Build()` and `CreateWorkflow` return a plain error for a structurally invalid DAG before any
request is made. Return `queueflow.NonRetryable(err)` from a worker handler to dead-letter the
job immediately.

## Known limitation: `StreamJobEvents`

`JobsAPI.StreamJobEvents` cannot consume the server's SSE stream: it buffers the whole response
until the stream closes (terminal status or the 15-minute cap) and returns it as one string. Use
`WaitForJob` (polling) instead, or a hand-rolled SSE consumer over `GET /api/v1/jobs/{id}/events`.

## Conformance tests

`live_test.go` runs the facade against a real QueueFlow server (started with the default
`queueflow serve`, i.e. `--mode all`, so the built-in `echo` handler is registered). It creates
and waits on an `echo` job, checks idempotent re-creation and list filters, runs a two-step
workflow, exercises the cron, dead-letter, queue backlog and stats endpoints, and runs `Run`
against a dedicated queue: one scenario completes a job with a handler only the test knows,
another dead-letters a job with `NonRetryable` and replays it. The tests skip themselves unless
`QUEUEFLOW_URL` is set, so plain `go test ./...` and CI stay offline.

```bash
QUEUEFLOW_URL=http://localhost:8000 \
QUEUEFLOW_TOKEN=dev \
QUEUEFLOW_WORKER_TOKEN=worker-secret \
go test -count=1 -run Live ./...
```

`QUEUEFLOW_TOKEN` is the tenant token (default `dev`). `QUEUEFLOW_WORKER_TOKEN` is the server's
`--worker-token`; it defaults to the tenant token, which only works when the server runs without
one.

## Architecture

This package is a thin hand-written **facade** over a **generated core**:

```
queueflow-sdk-go/
├── facade.go            the hand-written facade (QueueFlow, Run, WorkflowBuilder, APIError)
├── facade_test.go       offline facade tests (httptest server)
├── live_test.go         live conformance tests (need QUEUEFLOW_URL)
├── api_*.go             generated per-tag services (JobsAPI, WorkflowsAPI, WorkerAPI, ...)
├── model_*.go           generated models
├── client.go, configuration.go, response.go, utils.go   generated transport
├── docs/                generated per-endpoint and per-model reference
└── test/                generated per-service test stubs (skipped)
```

The core is generated by openapi-generator from the server's
[OpenAPI spec](https://github.com/elision-labs/queueflow-core/blob/main/spec/openapi.yaml)
(`scripts/generate-sdks.sh go` in queueflow-core). The facade is injected at generation time from
`sdk-templates/go/facade.mustache` in that repo, so edit the template, not `facade.go`.

### Development

```bash
go vet ./...
go test ./...
```

## Links

- Documentation: [docs.queueflow.dev](https://docs.queueflow.dev)
- Server: [github.com/elision-labs/queueflow-core](https://github.com/elision-labs/queueflow-core)
- This SDK: [github.com/elision-labs/queueflow-sdk-go](https://github.com/elision-labs/queueflow-sdk-go)

## License

[MIT](./LICENSE)
