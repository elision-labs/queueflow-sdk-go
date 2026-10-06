# \WorkerAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteJob**](WorkerAPI.md#CompleteJob) | **Post** /api/v1/jobs/{id}/complete | 
[**FailJob**](WorkerAPI.md#FailJob) | **Post** /api/v1/jobs/{id}/fail | 
[**HeartbeatJob**](WorkerAPI.md#HeartbeatJob) | **Post** /api/v1/jobs/{id}/heartbeat | 
[**LeaseJobs**](WorkerAPI.md#LeaseJobs) | **Post** /api/v1/queues/{queue}/lease | 



## CompleteJob

> CompleteJob(ctx, id).CompleteJobRequest(completeJobRequest).Execute()



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
	id := "id_example" // string | Job id
	completeJobRequest := *openapiclient.NewCompleteJobRequest("LeaseToken_example") // CompleteJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkerAPI.CompleteJob(context.Background(), id).CompleteJobRequest(completeJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkerAPI.CompleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCompleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completeJobRequest** | [**CompleteJobRequest**](CompleteJobRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FailJob

> FailJob(ctx, id).FailJobRequest(failJobRequest).Execute()



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
	id := "id_example" // string | Job id
	failJobRequest := *openapiclient.NewFailJobRequest("Error_example", "LeaseToken_example") // FailJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkerAPI.FailJob(context.Background(), id).FailJobRequest(failJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkerAPI.FailJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFailJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failJobRequest** | [**FailJobRequest**](FailJobRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeartbeatJob

> HeartbeatResponse HeartbeatJob(ctx, id).HeartbeatRequest(heartbeatRequest).Execute()



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
	id := "id_example" // string | Job id
	heartbeatRequest := *openapiclient.NewHeartbeatRequest(int32(123), "LeaseToken_example") // HeartbeatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkerAPI.HeartbeatJob(context.Background(), id).HeartbeatRequest(heartbeatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkerAPI.HeartbeatJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeartbeatJob`: HeartbeatResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkerAPI.HeartbeatJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeartbeatJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **heartbeatRequest** | [**HeartbeatRequest**](HeartbeatRequest.md) |  | 

### Return type

[**HeartbeatResponse**](HeartbeatResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaseJobs

> LeaseJobsResponse LeaseJobs(ctx, queue).LeaseJobsRequest(leaseJobsRequest).Execute()



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
	queue := "queue_example" // string | Queue to lease from
	leaseJobsRequest := *openapiclient.NewLeaseJobsRequest() // LeaseJobsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkerAPI.LeaseJobs(context.Background(), queue).LeaseJobsRequest(leaseJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkerAPI.LeaseJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LeaseJobs`: LeaseJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkerAPI.LeaseJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queue** | **string** | Queue to lease from | 

### Other Parameters

Other parameters are passed through a pointer to a apiLeaseJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leaseJobsRequest** | [**LeaseJobsRequest**](LeaseJobsRequest.md) |  | 

### Return type

[**LeaseJobsResponse**](LeaseJobsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

