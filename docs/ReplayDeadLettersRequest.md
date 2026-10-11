# ReplayDeadLettersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAfter** | Pointer to **NullableTime** | Only entries created at or after this instant (RFC 3339, inclusive). | [optional] 
**CreatedBefore** | Pointer to **NullableTime** | Only entries created strictly before this instant (RFC 3339, exclusive). | [optional] 
**Limit** | Pointer to **NullableInt64** | Consider at most this many entries (oldest first). Default 100; the maximum is the server&#39;s batch limit (&#x60;--max-batch&#x60;). | [optional] 
**Queue** | Pointer to **NullableString** | Only dead letters from this queue. | [optional] 
**TaskName** | Pointer to **NullableString** | Only dead letters for this task. | [optional] 

## Methods

### NewReplayDeadLettersRequest

`func NewReplayDeadLettersRequest() *ReplayDeadLettersRequest`

NewReplayDeadLettersRequest instantiates a new ReplayDeadLettersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayDeadLettersRequestWithDefaults

`func NewReplayDeadLettersRequestWithDefaults() *ReplayDeadLettersRequest`

NewReplayDeadLettersRequestWithDefaults instantiates a new ReplayDeadLettersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAfter

`func (o *ReplayDeadLettersRequest) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *ReplayDeadLettersRequest) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *ReplayDeadLettersRequest) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *ReplayDeadLettersRequest) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### SetCreatedAfterNil

`func (o *ReplayDeadLettersRequest) SetCreatedAfterNil(b bool)`

 SetCreatedAfterNil sets the value for CreatedAfter to be an explicit nil

### UnsetCreatedAfter
`func (o *ReplayDeadLettersRequest) UnsetCreatedAfter()`

UnsetCreatedAfter ensures that no value is present for CreatedAfter, not even an explicit nil
### GetCreatedBefore

`func (o *ReplayDeadLettersRequest) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *ReplayDeadLettersRequest) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *ReplayDeadLettersRequest) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *ReplayDeadLettersRequest) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### SetCreatedBeforeNil

`func (o *ReplayDeadLettersRequest) SetCreatedBeforeNil(b bool)`

 SetCreatedBeforeNil sets the value for CreatedBefore to be an explicit nil

### UnsetCreatedBefore
`func (o *ReplayDeadLettersRequest) UnsetCreatedBefore()`

UnsetCreatedBefore ensures that no value is present for CreatedBefore, not even an explicit nil
### GetLimit

`func (o *ReplayDeadLettersRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReplayDeadLettersRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReplayDeadLettersRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReplayDeadLettersRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ReplayDeadLettersRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ReplayDeadLettersRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetQueue

`func (o *ReplayDeadLettersRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *ReplayDeadLettersRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *ReplayDeadLettersRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *ReplayDeadLettersRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *ReplayDeadLettersRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *ReplayDeadLettersRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTaskName

`func (o *ReplayDeadLettersRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *ReplayDeadLettersRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *ReplayDeadLettersRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *ReplayDeadLettersRequest) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *ReplayDeadLettersRequest) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *ReplayDeadLettersRequest) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


