# DeadLetter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Id** | **int64** |  | 
**JobId** | **string** |  | 
**QueueName** | Pointer to **NullableString** |  | [optional] 
**Reason** | **string** | Why the job dead-lettered: &#x60;max_attempts_exceeded&#x60;, &#x60;non_retryable&#x60;, or &#x60;handler_not_found&#x60;. | 
**ReplayJobId** | Pointer to **NullableString** | The fresh job created by the replay. | [optional] 
**ReplayedAt** | Pointer to **NullableTime** | Set once this entry has been replayed; a dead letter replays at most once. | [optional] 
**TaskName** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeadLetter

`func NewDeadLetter(createdAt time.Time, id int64, jobId string, reason string, ) *DeadLetter`

NewDeadLetter instantiates a new DeadLetter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeadLetterWithDefaults

`func NewDeadLetterWithDefaults() *DeadLetter`

NewDeadLetterWithDefaults instantiates a new DeadLetter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DeadLetter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeadLetter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeadLetter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorMessage

`func (o *DeadLetter) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DeadLetter) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DeadLetter) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DeadLetter) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *DeadLetter) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *DeadLetter) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *DeadLetter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeadLetter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeadLetter) SetId(v int64)`

SetId sets Id field to given value.


### GetJobId

`func (o *DeadLetter) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DeadLetter) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DeadLetter) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetQueueName

`func (o *DeadLetter) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *DeadLetter) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *DeadLetter) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *DeadLetter) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### SetQueueNameNil

`func (o *DeadLetter) SetQueueNameNil(b bool)`

 SetQueueNameNil sets the value for QueueName to be an explicit nil

### UnsetQueueName
`func (o *DeadLetter) UnsetQueueName()`

UnsetQueueName ensures that no value is present for QueueName, not even an explicit nil
### GetReason

`func (o *DeadLetter) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DeadLetter) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DeadLetter) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReplayJobId

`func (o *DeadLetter) GetReplayJobId() string`

GetReplayJobId returns the ReplayJobId field if non-nil, zero value otherwise.

### GetReplayJobIdOk

`func (o *DeadLetter) GetReplayJobIdOk() (*string, bool)`

GetReplayJobIdOk returns a tuple with the ReplayJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayJobId

`func (o *DeadLetter) SetReplayJobId(v string)`

SetReplayJobId sets ReplayJobId field to given value.

### HasReplayJobId

`func (o *DeadLetter) HasReplayJobId() bool`

HasReplayJobId returns a boolean if a field has been set.

### SetReplayJobIdNil

`func (o *DeadLetter) SetReplayJobIdNil(b bool)`

 SetReplayJobIdNil sets the value for ReplayJobId to be an explicit nil

### UnsetReplayJobId
`func (o *DeadLetter) UnsetReplayJobId()`

UnsetReplayJobId ensures that no value is present for ReplayJobId, not even an explicit nil
### GetReplayedAt

`func (o *DeadLetter) GetReplayedAt() time.Time`

GetReplayedAt returns the ReplayedAt field if non-nil, zero value otherwise.

### GetReplayedAtOk

`func (o *DeadLetter) GetReplayedAtOk() (*time.Time, bool)`

GetReplayedAtOk returns a tuple with the ReplayedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayedAt

`func (o *DeadLetter) SetReplayedAt(v time.Time)`

SetReplayedAt sets ReplayedAt field to given value.

### HasReplayedAt

`func (o *DeadLetter) HasReplayedAt() bool`

HasReplayedAt returns a boolean if a field has been set.

### SetReplayedAtNil

`func (o *DeadLetter) SetReplayedAtNil(b bool)`

 SetReplayedAtNil sets the value for ReplayedAt to be an explicit nil

### UnsetReplayedAt
`func (o *DeadLetter) UnsetReplayedAt()`

UnsetReplayedAt ensures that no value is present for ReplayedAt, not even an explicit nil
### GetTaskName

`func (o *DeadLetter) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *DeadLetter) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *DeadLetter) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *DeadLetter) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *DeadLetter) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *DeadLetter) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
### GetTenantId

`func (o *DeadLetter) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DeadLetter) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DeadLetter) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DeadLetter) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *DeadLetter) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *DeadLetter) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


