# \WorkflowsAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWorkflow**](WorkflowsAPI.md#CancelWorkflow) | **Post** /api/v1/workflows/{id}/cancel | 
[**CreateWorkflow**](WorkflowsAPI.md#CreateWorkflow) | **Post** /api/v1/workflows | 
[**GetWorkflow**](WorkflowsAPI.md#GetWorkflow) | **Get** /api/v1/workflows/{id} | 
[**GetWorkflowDiagram**](WorkflowsAPI.md#GetWorkflowDiagram) | **Get** /api/v1/workflows/{id}/diagram | 
[**ListWorkflows**](WorkflowsAPI.md#ListWorkflows) | **Get** /api/v1/workflows | 



## CancelWorkflow

> CancelWorkflow(ctx, id).Execute()



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
	id := "id_example" // string | Workflow id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.CancelWorkflow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CancelWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkflowRequest struct via the builder pattern


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


## CreateWorkflow

> CreateWorkflowResponse CreateWorkflow(ctx).CreateWorkflowRequest(createWorkflowRequest).Execute()



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
	createWorkflowRequest := *openapiclient.NewCreateWorkflowRequest("Name_example", []openapiclient.WorkflowStep{*openapiclient.NewWorkflowStep("Name_example", "TaskName_example")}) // CreateWorkflowRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.CreateWorkflow(context.Background()).CreateWorkflowRequest(createWorkflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflow`: CreateWorkflowResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRequest** | [**CreateWorkflowRequest**](CreateWorkflowRequest.md) |  | 

### Return type

[**CreateWorkflowResponse**](CreateWorkflowResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> Workflow GetWorkflow(ctx, id).Execute()



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
	id := "id_example" // string | Workflow id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflow`: Workflow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workflow**](Workflow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowDiagram

> WorkflowDiagramResponse GetWorkflowDiagram(ctx, id).Execute()



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
	id := "id_example" // string | Workflow id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowDiagram(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowDiagram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowDiagram`: WorkflowDiagramResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowDiagram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowDiagramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowDiagramResponse**](WorkflowDiagramResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> ListWorkflowsResponse ListWorkflows(ctx).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).Execute()



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
	status := "status_example" // string | Filter by status (e.g. `pending`, `completed`). (optional)
	queue := "queue_example" // string | Filter by queue name (jobs only). (optional)
	limit := int64(789) // int64 | Page size, 1..=100 (default 50). (optional)
	offset := int64(789) // int64 | Number of records to skip (default 0). (optional)
	orderBy := "orderBy_example" // string | `created_at ASC` or `created_at DESC` (default DESC). (optional)
	includeTotal := true // bool | Include the exact `total` count in the response (default false; the count is an extra full scan over the filtered set). (optional)
	cursor := "cursor_example" // string | Opaque keyset cursor from a previous page's `next_cursor`. When set, `offset` is ignored and listing continues where that page ended. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflows(context.Background()).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflows`: ListWorkflowsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by status (e.g. &#x60;pending&#x60;, &#x60;completed&#x60;). | 
 **queue** | **string** | Filter by queue name (jobs only). | 
 **limit** | **int64** | Page size, 1..&#x3D;100 (default 50). | 
 **offset** | **int64** | Number of records to skip (default 0). | 
 **orderBy** | **string** | &#x60;created_at ASC&#x60; or &#x60;created_at DESC&#x60; (default DESC). | 
 **includeTotal** | **bool** | Include the exact &#x60;total&#x60; count in the response (default false; the count is an extra full scan over the filtered set). | 
 **cursor** | **string** | Opaque keyset cursor from a previous page&#39;s &#x60;next_cursor&#x60;. When set, &#x60;offset&#x60; is ignored and listing continues where that page ended. | 

### Return type

[**ListWorkflowsResponse**](ListWorkflowsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

