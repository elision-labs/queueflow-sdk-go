# \CronAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCron**](CronAPI.md#CreateCron) | **Post** /api/v1/cron | 
[**DeleteCron**](CronAPI.md#DeleteCron) | **Delete** /api/v1/cron/{id} | 
[**GetCron**](CronAPI.md#GetCron) | **Get** /api/v1/cron/{id} | 
[**ListCrons**](CronAPI.md#ListCrons) | **Get** /api/v1/cron | 
[**PauseCron**](CronAPI.md#PauseCron) | **Post** /api/v1/cron/{id}/pause | 
[**ResumeCron**](CronAPI.md#ResumeCron) | **Post** /api/v1/cron/{id}/resume | 



## CreateCron

> CreateCronResponse CreateCron(ctx).CreateCronRequest(createCronRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	createCronRequest := *openapiclient.NewCreateCronRequest("CronExpr_example", "Name_example", "TaskName_example") // CreateCronRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.CreateCron(context.Background()).CreateCronRequest(createCronRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.CreateCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCron`: CreateCronResponse
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.CreateCron`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCronRequest** | [**CreateCronRequest**](CreateCronRequest.md) |  | 

### Return type

[**CreateCronResponse**](CreateCronResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCron

> DeleteCron(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	id := "id_example" // string | Cron schedule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronAPI.DeleteCron(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.DeleteCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cron schedule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCronRequest struct via the builder pattern


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


## GetCron

> CronSchedule GetCron(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	id := "id_example" // string | Cron schedule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CronAPI.GetCron(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.GetCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCron`: CronSchedule
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.GetCron`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cron schedule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CronSchedule**](CronSchedule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCrons

> ListCronsResponse ListCrons(ctx).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
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
	resp, r, err := apiClient.CronAPI.ListCrons(context.Background()).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.ListCrons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCrons`: ListCronsResponse
	fmt.Fprintf(os.Stdout, "Response from `CronAPI.ListCrons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCronsRequest struct via the builder pattern


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

[**ListCronsResponse**](ListCronsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseCron

> PauseCron(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	id := "id_example" // string | Cron schedule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronAPI.PauseCron(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.PauseCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cron schedule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseCronRequest struct via the builder pattern


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


## ResumeCron

> ResumeCron(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/elision-labs/queueflow-sdk-go"
)

func main() {
	id := "id_example" // string | Cron schedule id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CronAPI.ResumeCron(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CronAPI.ResumeCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Cron schedule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCronRequest struct via the builder pattern


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

