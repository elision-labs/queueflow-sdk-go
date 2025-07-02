# QueueFlow Go SDK

The official Go SDK for QueueFlow distributed job queue system.

## Installation

```bash
go get github.com/queueflow/queueflow-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    
    "github.com/queueflow/queueflow-sdk-go"
)

func main() {
    // Create client
    client := queueflow.NewClient("http://localhost:8080", "your-api-key")
    
    // Create a job
    job, err := client.CreateJob(context.Background(), &queueflow.CreateJobRequest{
        TaskName: "process_data",
        Payload: map[string]interface{}{
            "user_id": 123,
            "action": "send_email",
        },
        Config: &queueflow.JobConfig{
            Priority: queueflow.PriorityHigh,
            Retries:  3,
        },
    })
    if err != nil {
        log.Fatalf("Failed to create job: %v", err)
    }
    
    log.Printf("Job created: %s", job.ID)
    
    // Get job status
    status, err := client.GetJob(context.Background(), job.ID)
    if err != nil {
        log.Fatalf("Failed to get job: %v", err)
    }
    
    log.Printf("Job status: %s", status.Status)
}
```

## Features

- ✅ Create and manage jobs
- ✅ Batch job operations  
- ✅ Job status monitoring
- ✅ Workflow support
- ✅ Context cancellation
- ✅ Automatic retries
- ✅ Type-safe API

## API Reference

### Client

```go
// Create a new client
client := queueflow.NewClient(baseURL, apiKey)

// With custom HTTP client
client := queueflow.NewClientWithHTTP(baseURL, apiKey, httpClient)
```

### Jobs

```go
// Create a job
job, err := client.CreateJob(ctx, &queueflow.CreateJobRequest{
    TaskName: "task_name",
    Payload:  map[string]interface{}{"key": "value"},
})

// Get job status
job, err := client.GetJob(ctx, jobID)

// Cancel job
err := client.CancelJob(ctx, jobID)

// List jobs
jobs, err := client.ListJobs(ctx, &queueflow.ListJobsRequest{
    Status: queueflow.JobStatusPending,
    Limit:  10,
})
```

### Batches

```go
// Create batch
batch, err := client.CreateBatch(ctx, &queueflow.CreateBatchRequest{
    Jobs: []queueflow.CreateJobRequest{
        {TaskName: "task1", Payload: data1},
        {TaskName: "task2", Payload: data2},
    },
})

// Get batch status
batch, err := client.GetBatch(ctx, batchID)
```

### Workflows

```go
// Create workflow
workflow, err := client.CreateWorkflow(ctx, &queueflow.CreateWorkflowRequest{
    Name: "data_pipeline",
    Steps: []queueflow.WorkflowStep{
        {
            Name:     "extract",
            TaskName: "extract_data",
            Payload:  extractConfig,
        },
        {
            Name:      "transform", 
            TaskName:  "transform_data",
            DependsOn: []string{"extract"},
        },
    },
})
```

## Configuration

### Job Configuration

```go
config := &queueflow.JobConfig{
    Priority:    queueflow.PriorityHigh,
    Retries:     3,
    Timeout:     time.Minute * 5,
    Delay:       time.Second * 30,
    Queue:       "priority_queue",
}
```

### Client Configuration

```go
client := queueflow.NewClient(baseURL, apiKey)
client.SetTimeout(time.Second * 30)
client.SetRetries(3)
```

## Error Handling

```go
job, err := client.CreateJob(ctx, req)
if err != nil {
    if queueflow.IsNotFoundError(err) {
        // Handle not found
    } else if queueflow.IsValidationError(err) {
        // Handle validation error
    } else {
        // Handle other errors
    }
}
```

## Examples

See the [examples](./examples/) directory for complete examples:

- [Basic Usage](./examples/main.go)
- [Batch Processing](./examples/batch.go) 
- [Workflow Orchestration](./examples/workflow.go)
- [Error Handling](./examples/errors.go)

## Development

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Lint code
golangci-lint run
```

## License

MIT License