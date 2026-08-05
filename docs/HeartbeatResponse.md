# HeartbeatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**JobStatus**](JobStatus.md) | The job&#39;s current status. &#x60;running&#x60; means the lease was extended; anything else (&#x60;cancelled&#x60;, &#x60;completed&#x60;, ...) means it was not, and the worker should stop working on the job. | 

## Methods

### NewHeartbeatResponse

`func NewHeartbeatResponse(status JobStatus, ) *HeartbeatResponse`

NewHeartbeatResponse instantiates a new HeartbeatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartbeatResponseWithDefaults

`func NewHeartbeatResponseWithDefaults() *HeartbeatResponse`

NewHeartbeatResponseWithDefaults instantiates a new HeartbeatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HeartbeatResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HeartbeatResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HeartbeatResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


