# FailJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Human-readable failure reason. | 
**LeaseToken** | **string** | The lease token returned by the lease call. | 
**Retryable** | Pointer to **bool** | Whether the engine may retry (subject to the job&#39;s max_retries). Defaults to true; send false for permanent failures (e.g. bad input). | [optional] 

## Methods

### NewFailJobRequest

`func NewFailJobRequest(error_ string, leaseToken string, ) *FailJobRequest`

NewFailJobRequest instantiates a new FailJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailJobRequestWithDefaults

`func NewFailJobRequestWithDefaults() *FailJobRequest`

NewFailJobRequestWithDefaults instantiates a new FailJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FailJobRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FailJobRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FailJobRequest) SetError(v string)`

SetError sets Error field to given value.


### GetLeaseToken

`func (o *FailJobRequest) GetLeaseToken() string`

GetLeaseToken returns the LeaseToken field if non-nil, zero value otherwise.

### GetLeaseTokenOk

`func (o *FailJobRequest) GetLeaseTokenOk() (*string, bool)`

GetLeaseTokenOk returns a tuple with the LeaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseToken

`func (o *FailJobRequest) SetLeaseToken(v string)`

SetLeaseToken sets LeaseToken field to given value.


### GetRetryable

`func (o *FailJobRequest) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *FailJobRequest) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *FailJobRequest) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *FailJobRequest) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


