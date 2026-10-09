# queueflow-sdk-go

Go client for [QueueFlow](https://queueflow.dev), a PostgreSQL-native distributed job queue and
workflow engine.

- **Typed end-to-end**: request/response models mirror the server's OpenAPI 3.1 spec.
- **Ergonomic**: `NewQueueFlow`, `CreateJob`, `WaitForJob`/`WaitForWorkflow` pollers, and a
  workflow builder with local DAG validation.
- **Thin facade over a generated core**: the facade (`facade.go`) adds what codegen cannot; the
  generated `APIClient` and its `*APIService` groups stay available for everything else.
- **No worker runtime**: this SDK exposes the worker-protocol endpoints as raw calls only
  (see [Worker protocol](#worker-protocol)).

Module path: `github.com/elision-labs/queueflow-sdk-go`, package `queueflow`. Go 1.18 or newer.

## Install

```bash
go get github.com/elision-labs/queueflow-sdk-go
```

## Quick start

```go
package main

import (
	"log"
	"time"

	queueflow "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	qf := queueflow.NewQueueFlow("http://localhost:8000", "dev")

	// Enqueue a job and wait for the result.
	job, err := qf.CreateJob("echo", map[string]interface{}{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	done, err := qf.WaitForJob(job.Id, 60*time.Second, 500*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(done.Status, done.Result) // completed map[echoed:true hello:world]

	// Declare and run a DAG workflow.
	dag := queueflow.NewWorkflowBuilder("etl").
		Step("extract", "echo").
		Step("transform", "echo", "extract").
		Step("load", "echo", "transform")
	wf, err := qf.CreateWorkflow(dag)
	if err != nil {
		log.Fatal(err)
	}
	finished, err := qf.WaitForWorkflow(wf.Id, 60*time.Second, 500*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(finished.Status)
}
```

`echo` is a handler built into the server; it returns the payload plus `"echoed": true`.

## API

### Client

`NewQueueFlow(baseURL, token string) *QueueFlow` configures the generated client with
`Authorization: Bearer <token>` and exposes:

| Method | Description |
| --- | --- |
| `.Client` | The generated `*APIClient` (`.Client.JobsAPI`, `.Client.WorkflowsAPI`, `.Client.WorkerAPI`, `.Client.CronAPI`, `.Client.DlqAPI`, `.Client.SystemAPI`, `.Client.HealthAPI`). |
| `CreateJob(task string, payload map[string]interface{}) (*Job, error)` | Enqueue a job and return the created record. |
| `WaitForJob(jobID string, timeout, interval time.Duration) (*Job, error)` | Poll until `completed` / `failed` / `cancelled`; errors past `timeout`. |
| `CreateWorkflow(b *WorkflowBuilder) (*Workflow, error)` | Validate the DAG locally, create it, return the record. |
| `WaitForWorkflow(workflowID string, timeout, interval time.Duration) (*Workflow, error)` | Poll until a terminal workflow state. |
| `CreateCron(name, cronExpr, task string, payload map[string]interface{}) (string, error)` | Register a recurring enqueue (5-field crontab, UTC); returns the schedule id. |
| `ReplayDeadLetter(id int64) (*Job, error)` | Re-run a dead-lettered job as a fresh job (at most once; a second replay is a 409). |

### Workflow builder

```go
dag := queueflow.NewWorkflowBuilder("order_123").
	Step("validate", "validate_order").
	Step("pay", "process_payment", "validate").
	Step("ship", "create_shipment", "pay")
body, err := dag.Build() // duplicate names, dangling deps and cycles fail here, before any request
```

`Step(name, task string, after ...string)` adds a step gated on the named `after` steps.
`CreateWorkflow` calls `Build()` for you.

### Everything else: the generated client

Anything the facade does not wrap is one call away on `qf.Client`. Every operation is a request
builder ending in `.Execute()`, which returns `(result, *http.Response, error)`:

```go
ctx := context.Background()

// Idempotent create: the same key returns the original job instead of a duplicate.
created, _, err := qf.Client.JobsAPI.CreateJob(ctx).
	CreateJobRequest(queueflow.CreateJobRequest{TaskName: "echo", Payload: payload}).
	IdempotencyKey("order-123").
	Execute()

// Lists: filters, paging, keyset cursors.
jobs, _, err := qf.Client.JobsAPI.ListJobs(ctx).Status("failed").Queue("billing").Limit(50).Execute()
crons, _, err := qf.Client.CronAPI.ListCrons(ctx).Limit(100).Execute()
dlq, _, err := qf.Client.DlqAPI.ListDeadLetters(ctx).Limit(10).Execute()

// Cron lifecycle.
_, err = qf.Client.CronAPI.PauseCron(ctx, cronID).Execute()
_, err = qf.Client.CronAPI.ResumeCron(ctx, cronID).Execute()
_, err = qf.Client.CronAPI.DeleteCron(ctx, cronID).Execute()

// Workflow progress and diagram.
states, _, err := qf.Client.WorkflowsAPI.GetWorkflowStepStates(ctx, wfID).Execute()
diagram, _, err := qf.Client.WorkflowsAPI.GetWorkflowDiagram(ctx, wfID).Execute()

// Engine introspection.
stats, _, err := qf.Client.SystemAPI.GetStats(ctx).Execute()
tasks, _, err := qf.Client.SystemAPI.ListTasks(ctx).Execute()
```

List responses carry `NextCursor` when there are more pages; pass it back with `.Cursor(...)` for
keyset pagination (cheaper than deep `.Offset(...)`). Per-endpoint and per-model reference for
the generated layer lives in [`docs/`](./docs).

## Authentication

Every request carries `Authorization: Bearer <token>`. Two credentials exist:

- The **tenant token** (the `token` passed to `NewQueueFlow`) authenticates the job, workflow,
  cron, DLQ and system routes. On a server started without `--api-keys`, any non-empty token is
  accepted (development mode).
- The **worker token** authenticates the worker-protocol routes (lease, heartbeat, complete,
  fail). Configure it on the server with `--worker-token` / `QUEUEFLOW_WORKER_TOKEN`. It is a
  separate secret, never a tenant token. Build a second client for it:

  ```go
  cfg := queueflow.NewConfiguration()
  cfg.Servers = queueflow.ServerConfigurations{{URL: baseURL}}
  cfg.AddDefaultHeader("Authorization", "Bearer "+workerToken)
  worker := queueflow.NewAPIClient(cfg).WorkerAPI
  ```

## Worker protocol

This SDK ships **no worker runtime**. `WorkerAPI` (`LeaseJobs`, `HeartbeatJob`, `CompleteJob`,
`FailJob`) is exposed as raw calls; if you build a loop on top, three rules keep the
at-least-once contract honest:

1. Worker routes authenticate with the worker token, not a tenant token (see above).
2. Heartbeat every in-flight job at roughly half its lease interval. A heartbeat whose `status`
   is anything other than `running` (or an HTTP 409) means the server owns the outcome: abandon
   the handler and report nothing. Never process a leased batch sequentially without
   heartbeating the jobs still waiting; their leases expire and the server redelivers them.
3. Delivery is at-least-once, so handlers must be idempotent. Report permanent failures with
   `Retryable: false` so they dead-letter immediately instead of burning retries.

The Node (`@queueflow/sdk`) and Python (`queueflow`) SDKs include a worker runtime if you need
one off the shelf.

## Error handling

Generated calls return a `*GenericOpenAPIError` for non-2xx responses; the `*http.Response` is
returned alongside, so the status code is always available:

```go
job, resp, err := qf.Client.JobsAPI.GetJob(ctx, id).Execute()
if err != nil {
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		// no such job
	}
	var apiErr *queueflow.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		log.Printf("%s: %s", apiErr.Error(), apiErr.Body())
	}
	return err
}
```

`WaitForJob` and `WaitForWorkflow` return a plain error past their deadline; `Build()` and
`CreateWorkflow` return an error for a structurally invalid DAG before any request is made.

## Known limitation: `StreamJobEvents`

`JobsAPI.StreamJobEvents` cannot consume the server's SSE stream: it buffers the whole response
until the stream closes (terminal status or the 15-minute cap) and returns it as one string. Use
`WaitForJob` (polling) instead, or a hand-rolled SSE consumer over `GET /api/v1/jobs/{id}/events`.

## Conformance tests

`live_test.go` runs the facade against a real QueueFlow server (started with the default
`queueflow serve`, i.e. `--mode all`, so the built-in `echo` handler is registered). It creates
and waits on an `echo` job, checks idempotent re-creation, runs a two-step workflow, exercises the
cron and dead-letter endpoints, and reads stats. The tests skip themselves unless `QUEUEFLOW_URL`
is set, so plain `go test ./...` and CI stay offline.

```bash
QUEUEFLOW_URL=http://localhost:8000 \
QUEUEFLOW_TOKEN=dev \
QUEUEFLOW_WORKER_TOKEN=worker-secret \
go test -run Live ./...
```

`QUEUEFLOW_TOKEN` is the tenant token (default `dev`). `QUEUEFLOW_WORKER_TOKEN` is read for
parity with the other SDKs but unused here, since this SDK has no worker runtime.

## Architecture

This package is a thin hand-written **facade** over a **generated core**:

```
queueflow-sdk-go/
├── facade.go            the hand-written facade (QueueFlow, WorkflowBuilder)
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
