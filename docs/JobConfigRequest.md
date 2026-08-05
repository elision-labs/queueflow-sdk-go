# JobConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRetries** | Pointer to **NullableInt32** |  | [optional] 
**Priority** | Pointer to **NullableInt32** | Higher is claimed first within a queue (ties: oldest first). | [optional] 
**Queue** | Pointer to **NullableString** | Override the destination queue. | [optional] 
**Timeout** | Pointer to **NullableInt64** | Per-attempt timeout, in seconds. | [optional] 

## Methods

### NewJobConfigRequest

`func NewJobConfigRequest() *JobConfigRequest`

NewJobConfigRequest instantiates a new JobConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigRequestWithDefaults

`func NewJobConfigRequestWithDefaults() *JobConfigRequest`

NewJobConfigRequestWithDefaults instantiates a new JobConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxRetries

`func (o *JobConfigRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *JobConfigRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *JobConfigRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *JobConfigRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *JobConfigRequest) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *JobConfigRequest) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPriority

`func (o *JobConfigRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobConfigRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobConfigRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobConfigRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *JobConfigRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *JobConfigRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQueue

`func (o *JobConfigRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *JobConfigRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *JobConfigRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *JobConfigRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *JobConfigRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *JobConfigRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTimeout

`func (o *JobConfigRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *JobConfigRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *JobConfigRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *JobConfigRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *JobConfigRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *JobConfigRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


