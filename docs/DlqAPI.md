# \DlqAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeadLetter**](DlqAPI.md#GetDeadLetter) | **Get** /api/v1/dlq/{id} | 
[**ListDeadLetters**](DlqAPI.md#ListDeadLetters) | **Get** /api/v1/dlq | 
[**ReplayDeadLetter**](DlqAPI.md#ReplayDeadLetter) | **Post** /api/v1/dlq/{id}/replay | 



## GetDeadLetter

> DeadLetter GetDeadLetter(ctx, id).Execute()



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
	id := int64(789) // int64 | Dead letter id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlqAPI.GetDeadLetter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlqAPI.GetDeadLetter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeadLetter`: DeadLetter
	fmt.Fprintf(os.Stdout, "Response from `DlqAPI.GetDeadLetter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Dead letter id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeadLetterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeadLetter**](DeadLetter.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeadLetters

> ListDeadLettersResponse ListDeadLetters(ctx).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()



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
	resp, r, err := apiClient.DlqAPI.ListDeadLetters(context.Background()).Status(status).Queue(queue).Limit(limit).Offset(offset).OrderBy(orderBy).IncludeTotal(includeTotal).Cursor(cursor).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlqAPI.ListDeadLetters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeadLetters`: ListDeadLettersResponse
	fmt.Fprintf(os.Stdout, "Response from `DlqAPI.ListDeadLetters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeadLettersRequest struct via the builder pattern


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

[**ListDeadLettersResponse**](ListDeadLettersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayDeadLetter

> ReplayDeadLetterResponse ReplayDeadLetter(ctx, id).Execute()



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
	id := int64(789) // int64 | Dead letter id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlqAPI.ReplayDeadLetter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlqAPI.ReplayDeadLetter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayDeadLetter`: ReplayDeadLetterResponse
	fmt.Fprintf(os.Stdout, "Response from `DlqAPI.ReplayDeadLetter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Dead letter id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayDeadLetterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplayDeadLetterResponse**](ReplayDeadLetterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

