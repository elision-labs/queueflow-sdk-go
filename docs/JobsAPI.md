# \JobsAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobsAPI.md#CancelJob) | **Post** /api/v1/jobs/{id}/cancel | 
[**CreateBatchJobs**](JobsAPI.md#CreateBatchJobs) | **Post** /api/v1/jobs/batch | 
[**CreateJob**](JobsAPI.md#CreateJob) | **Post** /api/v1/jobs | 
[**GetJob**](JobsAPI.md#GetJob) | **Get** /api/v1/jobs/{id} | 
[**ListJobs**](JobsAPI.md#ListJobs) | **Get** /api/v1/jobs | 
[**StreamJobEvents**](JobsAPI.md#StreamJobEvents) | **Get** /api/v1/jobs/{id}/events | Stream a job&#39;s status transitions as Server-Sent Events until it reaches a terminal state. Lets clients await completion without polling the REST endpoint themselves.



## CancelJob

> CancelJob(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	id := "id_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JobsAPI.CancelJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.CancelJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchJobs

> CreateBatchJobsResponse CreateBatchJobs(ctx).CreateBatchJobsRequest(createBatchJobsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	createBatchJobsRequest := *openapiclient.NewCreateBatchJobsRequest([]openapiclient.CreateJobRequest{*openapiclient.NewCreateJobRequest("TaskName_example")}) // CreateBatchJobsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.CreateBatchJobs(context.Background()).CreateBatchJobsRequest(createBatchJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.CreateBatchJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchJobs`: CreateBatchJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.CreateBatchJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchJobsRequest** | [**CreateBatchJobsRequest**](CreateBatchJobsRequest.md) |  | 

### Return type

[**CreateBatchJobsResponse**](CreateBatchJobsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJob

> CreateJobResponse CreateJob(ctx).CreateJobRequest(createJobRequest).IdempotencyKey(idempotencyKey).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	createJobRequest := *openapiclient.NewCreateJobRequest("TaskName_example") // CreateJobRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Optional client-supplied key making this create idempotent per tenant: retrying with the same key returns the original job instead of creating a duplicate. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.CreateJob(context.Background()).CreateJobRequest(createJobRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.CreateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJob`: CreateJobResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createJobRequest** | [**CreateJobRequest**](CreateJobRequest.md) |  | 
 **idempotencyKey** | **string** | Optional client-supplied key making this create idempotent per tenant: retrying with the same key returns the original job instead of creating a duplicate. | 

### Return type

[**CreateJobResponse**](CreateJobResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> Job GetJob(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	id := "id_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.GetJob(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobs

> ListJobsResponse ListJobs(ctx).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	status := "status_example" // string | Filter by status (e.g. `pending`, `completed`). (optional)
	queue := "queue_example" // string | Filter by queue name (jobs only). (optional)
	limit := int64(789) // int64 | Page size, 1..=100 (default 50). (optional)
	offset := int64(789) // int64 | Number of records to skip (default 0). (optional)
	orderBy := "orderBy_example" // string | `created_at ASC` or `created_at DESC` (default DESC). (optional)
	includeTotal := true // bool | Include the exact `total` count in the response (default false; the count is an extra full scan over the filtered set). (optional)
	cursor := "cursor_example" // string | Opaque keyset cursor from a previous page's `next_cursor`. When set, `offset` is ignored and listing continues where that page ended. (optional)
	createdAfter := time.Now() // time.Time | Only rows created at or after this instant (RFC 3339, inclusive). With `created_before` this forms the half-open range `[after, before)` — the natural shape for walking history period by period. (optional)
	createdBefore := time.Now() // time.Time | Only rows created strictly before this instant (RFC 3339, exclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.ListJobs(context.Background()).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.ListJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListJobs`: ListJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.ListJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by status (e.g. &#x60;pending&#x60;, &#x60;completed&#x60;). | 
 **queue** | **string** | Filter by queue name (jobs only). | 
 **limit** | **int64** | Page size, 1..&#x3D;100 (default 50). | 
 **offset** | **int64** | Number of records to skip (default 0). | 
 **orderBy** | **string** | &#x60;created_at ASC&#x60; or &#x60;created_at DESC&#x60; (default DESC). | 
 **includeTotal** | **bool** | Include the exact &#x60;total&#x60; count in the response (default false; the count is an extra full scan over the filtered set). | 
 **cursor** | **string** | Opaque keyset cursor from a previous page&#39;s &#x60;next_cursor&#x60;. When set, &#x60;offset&#x60; is ignored and listing continues where that page ended. | 
 **createdAfter** | **time.Time** | Only rows created at or after this instant (RFC 3339, inclusive). With &#x60;created_before&#x60; this forms the half-open range &#x60;[after, before)&#x60; — the natural shape for walking history period by period. | 
 **createdBefore** | **time.Time** | Only rows created strictly before this instant (RFC 3339, exclusive). | 

### Return type

[**ListJobsResponse**](ListJobsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamJobEvents

> string StreamJobEvents(ctx, id).Execute()

Stream a job's status transitions as Server-Sent Events until it reaches a terminal state. Lets clients await completion without polling the REST endpoint themselves.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/queue-flow/sdk-go"
)

func main() {
	id := "id_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobsAPI.StreamJobEvents(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.StreamJobEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamJobEvents`: string
	fmt.Fprintf(os.Stdout, "Response from `JobsAPI.StreamJobEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamJobEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

