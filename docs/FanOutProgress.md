# FanOutProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | **int64** |  | 
**Failed** | **int64** | Failed or cancelled children. | 
**Total** | **int64** |  | 

## Methods

### NewFanOutProgress

`func NewFanOutProgress(completed int64, failed int64, total int64, ) *FanOutProgress`

NewFanOutProgress instantiates a new FanOutProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanOutProgressWithDefaults

`func NewFanOutProgressWithDefaults() *FanOutProgress`

NewFanOutProgressWithDefaults instantiates a new FanOutProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *FanOutProgress) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *FanOutProgress) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *FanOutProgress) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.


### GetFailed

`func (o *FanOutProgress) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *FanOutProgress) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *FanOutProgress) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetTotal

`func (o *FanOutProgress) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FanOutProgress) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FanOutProgress) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


