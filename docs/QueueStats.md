# QueueStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldestPendingAgeSecs** | Pointer to **NullableInt64** | Seconds the oldest claimable job has been waiting, &#x60;None&#x60; when nothing is claimable. The queue&#39;s primary health signal: it grows when workers cannot keep up and stays near zero when they can. | [optional] 
**Paused** | Pointer to **bool** | Claims for this tenant&#39;s jobs on the queue are paused (see &#x60;POST /api/v1/queues/{queue}/pause&#x60;). Pending jobs accumulate; running ones finish. In the all-tenants view this is true only when every registered tenant of the queue is paused. | [optional] 
**Pending** | **int64** | Claimable now: &#x60;pending&#x60; or &#x60;retrying&#x60; with &#x60;scheduled_at&#x60; in the past. | 
**Queue** | **string** |  | 
**Running** | **int64** | Currently leased by a worker. | 
**Scheduled** | **int64** | Waiting for a future &#x60;scheduled_at&#x60; (a &#x60;run_at&#x60; job or a backoff retry). | 

## Methods

### NewQueueStats

`func NewQueueStats(pending int64, queue string, running int64, scheduled int64, ) *QueueStats`

NewQueueStats instantiates a new QueueStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueStatsWithDefaults

`func NewQueueStatsWithDefaults() *QueueStats`

NewQueueStatsWithDefaults instantiates a new QueueStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldestPendingAgeSecs

`func (o *QueueStats) GetOldestPendingAgeSecs() int64`

GetOldestPendingAgeSecs returns the OldestPendingAgeSecs field if non-nil, zero value otherwise.

### GetOldestPendingAgeSecsOk

`func (o *QueueStats) GetOldestPendingAgeSecsOk() (*int64, bool)`

GetOldestPendingAgeSecsOk returns a tuple with the OldestPendingAgeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestPendingAgeSecs

`func (o *QueueStats) SetOldestPendingAgeSecs(v int64)`

SetOldestPendingAgeSecs sets OldestPendingAgeSecs field to given value.

### HasOldestPendingAgeSecs

`func (o *QueueStats) HasOldestPendingAgeSecs() bool`

HasOldestPendingAgeSecs returns a boolean if a field has been set.

### SetOldestPendingAgeSecsNil

`func (o *QueueStats) SetOldestPendingAgeSecsNil(b bool)`

 SetOldestPendingAgeSecsNil sets the value for OldestPendingAgeSecs to be an explicit nil

### UnsetOldestPendingAgeSecs
`func (o *QueueStats) UnsetOldestPendingAgeSecs()`

UnsetOldestPendingAgeSecs ensures that no value is present for OldestPendingAgeSecs, not even an explicit nil
### GetPaused

`func (o *QueueStats) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *QueueStats) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *QueueStats) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *QueueStats) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPending

`func (o *QueueStats) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *QueueStats) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *QueueStats) SetPending(v int64)`

SetPending sets Pending field to given value.


### GetQueue

`func (o *QueueStats) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *QueueStats) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *QueueStats) SetQueue(v string)`

SetQueue sets Queue field to given value.


### GetRunning

`func (o *QueueStats) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *QueueStats) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *QueueStats) SetRunning(v int64)`

SetRunning sets Running field to given value.


### GetScheduled

`func (o *QueueStats) GetScheduled() int64`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *QueueStats) GetScheduledOk() (*int64, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *QueueStats) SetScheduled(v int64)`

SetScheduled sets Scheduled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


