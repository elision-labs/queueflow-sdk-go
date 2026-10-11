# CancelJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | **int64** | How many jobs this call cancelled. Less than &#x60;limit&#x60; means the matching set is exhausted (jobs finishing concurrently are left as they ended, never overwritten). | 

## Methods

### NewCancelJobsResponse

`func NewCancelJobsResponse(cancelled int64, ) *CancelJobsResponse`

NewCancelJobsResponse instantiates a new CancelJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelJobsResponseWithDefaults

`func NewCancelJobsResponseWithDefaults() *CancelJobsResponse`

NewCancelJobsResponseWithDefaults instantiates a new CancelJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *CancelJobsResponse) GetCancelled() int64`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *CancelJobsResponse) GetCancelledOk() (*int64, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *CancelJobsResponse) SetCancelled(v int64)`

SetCancelled sets Cancelled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


