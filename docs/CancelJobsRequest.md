# CancelJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAfter** | Pointer to **NullableTime** | Only jobs created at or after this instant (RFC 3339, inclusive). | [optional] 
**CreatedBefore** | Pointer to **NullableTime** | Only jobs created strictly before this instant (RFC 3339, exclusive). | [optional] 
**Limit** | Pointer to **NullableInt64** | Cancel at most this many (oldest first). Default 100; the maximum is the server&#39;s batch limit (&#x60;--max-batch&#x60;, 1000 by default). | [optional] 
**Queue** | Pointer to **NullableString** | Only jobs on this queue. | [optional] 
**Status** | Pointer to **NullableString** | Only jobs in this state: &#x60;pending&#x60;, &#x60;retrying&#x60; or &#x60;running&#x60; (all three when omitted). Terminal states are rejected: history is never rewritten. | [optional] 
**TaskName** | Pointer to **NullableString** | Only jobs for this task. | [optional] 

## Methods

### NewCancelJobsRequest

`func NewCancelJobsRequest() *CancelJobsRequest`

NewCancelJobsRequest instantiates a new CancelJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelJobsRequestWithDefaults

`func NewCancelJobsRequestWithDefaults() *CancelJobsRequest`

NewCancelJobsRequestWithDefaults instantiates a new CancelJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAfter

`func (o *CancelJobsRequest) GetCreatedAfter() time.Time`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *CancelJobsRequest) GetCreatedAfterOk() (*time.Time, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *CancelJobsRequest) SetCreatedAfter(v time.Time)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *CancelJobsRequest) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### SetCreatedAfterNil

`func (o *CancelJobsRequest) SetCreatedAfterNil(b bool)`

 SetCreatedAfterNil sets the value for CreatedAfter to be an explicit nil

### UnsetCreatedAfter
`func (o *CancelJobsRequest) UnsetCreatedAfter()`

UnsetCreatedAfter ensures that no value is present for CreatedAfter, not even an explicit nil
### GetCreatedBefore

`func (o *CancelJobsRequest) GetCreatedBefore() time.Time`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *CancelJobsRequest) GetCreatedBeforeOk() (*time.Time, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *CancelJobsRequest) SetCreatedBefore(v time.Time)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *CancelJobsRequest) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### SetCreatedBeforeNil

`func (o *CancelJobsRequest) SetCreatedBeforeNil(b bool)`

 SetCreatedBeforeNil sets the value for CreatedBefore to be an explicit nil

### UnsetCreatedBefore
`func (o *CancelJobsRequest) UnsetCreatedBefore()`

UnsetCreatedBefore ensures that no value is present for CreatedBefore, not even an explicit nil
### GetLimit

`func (o *CancelJobsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CancelJobsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CancelJobsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CancelJobsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *CancelJobsRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *CancelJobsRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetQueue

`func (o *CancelJobsRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CancelJobsRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CancelJobsRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CancelJobsRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *CancelJobsRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *CancelJobsRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *CancelJobsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelJobsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelJobsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelJobsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CancelJobsRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CancelJobsRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskName

`func (o *CancelJobsRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *CancelJobsRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *CancelJobsRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *CancelJobsRequest) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### SetTaskNameNil

`func (o *CancelJobsRequest) SetTaskNameNil(b bool)`

 SetTaskNameNil sets the value for TaskName to be an explicit nil

### UnsetTaskName
`func (o *CancelJobsRequest) UnsetTaskName()`

UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


