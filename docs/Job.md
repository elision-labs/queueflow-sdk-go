# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Config** | [**JobConfig**](JobConfig.md) |  | 
**CreatedAt** | **time.Time** |  | 
**DeliveryCount** | Pointer to **int32** | How many times this job has been claimed (delivered to a worker). Greater than &#x60;retry_count + 1&#x60; means a lease expired without a report — i.e. a worker crashed mid-run. | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**IdempotencyKey** | Pointer to **NullableString** | Client-supplied key that makes job creation idempotent per tenant: re-submitting the same key returns the original job instead of creating a duplicate. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**NextRetryAt** | Pointer to **NullableTime** | When this job&#39;s next retry becomes claimable (mirrors &#x60;scheduled_at&#x60; while the job is &#x60;retrying&#x60;; kept for audit/inspection). | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**QueueName** | **string** |  | 
**Result** | Pointer to **map[string]interface{}** |  | [optional] 
**RetryCount** | **int32** |  | 
**ScheduledAt** | **time.Time** | When the job becomes claimable. &#x60;created_at&#x60; for immediate jobs, the requested &#x60;run_at&#x60; for scheduled jobs, and the next backoff instant while retrying — the durable delay lives in the row itself. | 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**TaskName** | **string** |  | 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**WorkflowStepId** | Pointer to **NullableString** | The owning workflow step&#39;s name (steps are addressed by name). | [optional] 

## Methods

### NewJob

`func NewJob(config JobConfig, createdAt time.Time, id string, queueName string, retryCount int32, scheduledAt time.Time, status JobStatus, taskName string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *Job) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Job) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Job) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Job) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Job) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Job) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetConfig

`func (o *Job) GetConfig() JobConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Job) GetConfigOk() (*JobConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Job) SetConfig(v JobConfig)`

SetConfig sets Config field to given value.


### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeliveryCount

`func (o *Job) GetDeliveryCount() int32`

GetDeliveryCount returns the DeliveryCount field if non-nil, zero value otherwise.

### GetDeliveryCountOk

`func (o *Job) GetDeliveryCountOk() (*int32, bool)`

GetDeliveryCountOk returns a tuple with the DeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCount

`func (o *Job) SetDeliveryCount(v int32)`

SetDeliveryCount sets DeliveryCount field to given value.

### HasDeliveryCount

`func (o *Job) HasDeliveryCount() bool`

HasDeliveryCount returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Job) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Job) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Job) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Job) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Job) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Job) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.


### GetIdempotencyKey

`func (o *Job) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Job) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Job) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Job) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### SetIdempotencyKeyNil

`func (o *Job) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *Job) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetMetadata

`func (o *Job) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Job) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Job) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Job) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNextRetryAt

`func (o *Job) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *Job) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *Job) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.

### HasNextRetryAt

`func (o *Job) HasNextRetryAt() bool`

HasNextRetryAt returns a boolean if a field has been set.

### SetNextRetryAtNil

`func (o *Job) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *Job) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil
### GetPayload

`func (o *Job) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Job) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Job) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Job) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetQueueName

`func (o *Job) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *Job) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *Job) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetResult

`func (o *Job) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Job) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Job) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *Job) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRetryCount

`func (o *Job) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Job) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Job) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.


### GetScheduledAt

`func (o *Job) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Job) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Job) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetStartedAt

`func (o *Job) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Job) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Job) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Job) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Job) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Job) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *Job) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetTaskName

`func (o *Job) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *Job) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *Job) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTenantId

`func (o *Job) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Job) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Job) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Job) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Job) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Job) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetWorkflowId

`func (o *Job) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Job) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Job) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *Job) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *Job) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *Job) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetWorkflowStepId

`func (o *Job) GetWorkflowStepId() string`

GetWorkflowStepId returns the WorkflowStepId field if non-nil, zero value otherwise.

### GetWorkflowStepIdOk

`func (o *Job) GetWorkflowStepIdOk() (*string, bool)`

GetWorkflowStepIdOk returns a tuple with the WorkflowStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStepId

`func (o *Job) SetWorkflowStepId(v string)`

SetWorkflowStepId sets WorkflowStepId field to given value.

### HasWorkflowStepId

`func (o *Job) HasWorkflowStepId() bool`

HasWorkflowStepId returns a boolean if a field has been set.

### SetWorkflowStepIdNil

`func (o *Job) SetWorkflowStepIdNil(b bool)`

 SetWorkflowStepIdNil sets the value for WorkflowStepId to be an explicit nil

### UnsetWorkflowStepId
`func (o *Job) UnsetWorkflowStepId()`

UnsetWorkflowStepId ensures that no value is present for WorkflowStepId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


