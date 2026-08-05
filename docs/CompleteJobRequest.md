# CompleteJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseToken** | **string** | The lease token returned by the lease call. | 
**Result** | Pointer to **map[string]interface{}** | Handler result, recorded on the job and merged into workflow context. | [optional] 

## Methods

### NewCompleteJobRequest

`func NewCompleteJobRequest(leaseToken string, ) *CompleteJobRequest`

NewCompleteJobRequest instantiates a new CompleteJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteJobRequestWithDefaults

`func NewCompleteJobRequestWithDefaults() *CompleteJobRequest`

NewCompleteJobRequestWithDefaults instantiates a new CompleteJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseToken

`func (o *CompleteJobRequest) GetLeaseToken() string`

GetLeaseToken returns the LeaseToken field if non-nil, zero value otherwise.

### GetLeaseTokenOk

`func (o *CompleteJobRequest) GetLeaseTokenOk() (*string, bool)`

GetLeaseTokenOk returns a tuple with the LeaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseToken

`func (o *CompleteJobRequest) SetLeaseToken(v string)`

SetLeaseToken sets LeaseToken field to given value.


### GetResult

`func (o *CompleteJobRequest) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CompleteJobRequest) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CompleteJobRequest) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *CompleteJobRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


