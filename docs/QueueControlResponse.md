# QueueControlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paused** | **bool** | Claims of the caller&#39;s jobs on this queue are paused. | 
**Pending** | **int64** | Claimable jobs of the caller waiting on this queue. | 
**Queue** | **string** |  | 
**Running** | **int64** | Jobs of the caller still leased by workers. A paused queue is drained once this reaches zero. | 

## Methods

### NewQueueControlResponse

`func NewQueueControlResponse(paused bool, pending int64, queue string, running int64, ) *QueueControlResponse`

NewQueueControlResponse instantiates a new QueueControlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueControlResponseWithDefaults

`func NewQueueControlResponseWithDefaults() *QueueControlResponse`

NewQueueControlResponseWithDefaults instantiates a new QueueControlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaused

`func (o *QueueControlResponse) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *QueueControlResponse) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *QueueControlResponse) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetPending

`func (o *QueueControlResponse) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *QueueControlResponse) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *QueueControlResponse) SetPending(v int64)`

SetPending sets Pending field to given value.


### GetQueue

`func (o *QueueControlResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *QueueControlResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *QueueControlResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.


### GetRunning

`func (o *QueueControlResponse) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *QueueControlResponse) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *QueueControlResponse) SetRunning(v int64)`

SetRunning sets Running field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


