# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Context** | Pointer to **map[string]interface{}** | Accumulated step results, keyed by step name. Passed to downstream steps under the &#x60;_context&#x60; payload key. | [optional] 
**CreatedAt** | **time.Time** |  | 
**DeadlineAt** | Pointer to **NullableTime** | When the janitor fails the workflow if it is still running (set from &#x60;timeout_secs&#x60;; pushed forward by a retry). | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Why the workflow ended &#x60;failed&#x60; or &#x60;cancelled&#x60;, when known (an upstream failure, a timeout, a user cancel). | [optional] 
**Id** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**ParentStepName** | Pointer to **NullableString** | The parent&#39;s step that spawned this child. | [optional] 
**ParentWorkflowId** | Pointer to **NullableString** | Set on a child workflow spawned by a sub-workflow step. | [optional] 
**RetryCount** | Pointer to **int32** | How many times &#x60;POST /workflows/{id}/retry&#x60; re-ran this workflow. | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | [**WorkflowStatus**](WorkflowStatus.md) |  | 
**Steps** | [**[]WorkflowStep**](WorkflowStep.md) |  | 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**TimeoutSecs** | Pointer to **NullableInt64** | The whole-workflow time budget it was created with, if any. | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow(createdAt time.Time, id string, name string, status WorkflowStatus, steps []WorkflowStep, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *Workflow) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Workflow) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Workflow) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Workflow) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Workflow) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Workflow) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetContext

`func (o *Workflow) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Workflow) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Workflow) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *Workflow) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Workflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeadlineAt

`func (o *Workflow) GetDeadlineAt() time.Time`

GetDeadlineAt returns the DeadlineAt field if non-nil, zero value otherwise.

### GetDeadlineAtOk

`func (o *Workflow) GetDeadlineAtOk() (*time.Time, bool)`

GetDeadlineAtOk returns a tuple with the DeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineAt

`func (o *Workflow) SetDeadlineAt(v time.Time)`

SetDeadlineAt sets DeadlineAt field to given value.

### HasDeadlineAt

`func (o *Workflow) HasDeadlineAt() bool`

HasDeadlineAt returns a boolean if a field has been set.

### SetDeadlineAtNil

`func (o *Workflow) SetDeadlineAtNil(b bool)`

 SetDeadlineAtNil sets the value for DeadlineAt to be an explicit nil

### UnsetDeadlineAt
`func (o *Workflow) UnsetDeadlineAt()`

UnsetDeadlineAt ensures that no value is present for DeadlineAt, not even an explicit nil
### GetErrorMessage

`func (o *Workflow) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Workflow) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Workflow) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Workflow) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Workflow) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Workflow) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *Workflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *Workflow) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Workflow) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Workflow) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Workflow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.


### GetParentStepName

`func (o *Workflow) GetParentStepName() string`

GetParentStepName returns the ParentStepName field if non-nil, zero value otherwise.

### GetParentStepNameOk

`func (o *Workflow) GetParentStepNameOk() (*string, bool)`

GetParentStepNameOk returns a tuple with the ParentStepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStepName

`func (o *Workflow) SetParentStepName(v string)`

SetParentStepName sets ParentStepName field to given value.

### HasParentStepName

`func (o *Workflow) HasParentStepName() bool`

HasParentStepName returns a boolean if a field has been set.

### SetParentStepNameNil

`func (o *Workflow) SetParentStepNameNil(b bool)`

 SetParentStepNameNil sets the value for ParentStepName to be an explicit nil

### UnsetParentStepName
`func (o *Workflow) UnsetParentStepName()`

UnsetParentStepName ensures that no value is present for ParentStepName, not even an explicit nil
### GetParentWorkflowId

`func (o *Workflow) GetParentWorkflowId() string`

GetParentWorkflowId returns the ParentWorkflowId field if non-nil, zero value otherwise.

### GetParentWorkflowIdOk

`func (o *Workflow) GetParentWorkflowIdOk() (*string, bool)`

GetParentWorkflowIdOk returns a tuple with the ParentWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowId

`func (o *Workflow) SetParentWorkflowId(v string)`

SetParentWorkflowId sets ParentWorkflowId field to given value.

### HasParentWorkflowId

`func (o *Workflow) HasParentWorkflowId() bool`

HasParentWorkflowId returns a boolean if a field has been set.

### SetParentWorkflowIdNil

`func (o *Workflow) SetParentWorkflowIdNil(b bool)`

 SetParentWorkflowIdNil sets the value for ParentWorkflowId to be an explicit nil

### UnsetParentWorkflowId
`func (o *Workflow) UnsetParentWorkflowId()`

UnsetParentWorkflowId ensures that no value is present for ParentWorkflowId, not even an explicit nil
### GetRetryCount

`func (o *Workflow) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *Workflow) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *Workflow) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *Workflow) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetStartedAt

`func (o *Workflow) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Workflow) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Workflow) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Workflow) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Workflow) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Workflow) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *Workflow) GetStatus() WorkflowStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Workflow) GetStatusOk() (*WorkflowStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Workflow) SetStatus(v WorkflowStatus)`

SetStatus sets Status field to given value.


### GetSteps

`func (o *Workflow) GetSteps() []WorkflowStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Workflow) GetStepsOk() (*[]WorkflowStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Workflow) SetSteps(v []WorkflowStep)`

SetSteps sets Steps field to given value.


### GetTenantId

`func (o *Workflow) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Workflow) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Workflow) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Workflow) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *Workflow) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *Workflow) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetTimeoutSecs

`func (o *Workflow) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *Workflow) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *Workflow) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *Workflow) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### SetTimeoutSecsNil

`func (o *Workflow) SetTimeoutSecsNil(b bool)`

 SetTimeoutSecsNil sets the value for TimeoutSecs to be an explicit nil

### UnsetTimeoutSecs
`func (o *Workflow) UnsetTimeoutSecs()`

UnsetTimeoutSecs ensures that no value is present for TimeoutSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


