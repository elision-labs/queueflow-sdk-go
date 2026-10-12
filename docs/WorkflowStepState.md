# WorkflowStepState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildWorkflowId** | Pointer to **NullableString** | The child workflow of a sub-workflow step, once spawned. | [optional] 
**CompletedAt** | Pointer to **NullableTime** | When the step reached a terminal status. | [optional] 
**Error** | Pointer to **NullableString** | Why the step failed, was skipped (&#x60;condition not met&#x60; for a false &#x60;when&#x60;) or was cancelled. | [optional] 
**FanOut** | Pointer to [**NullableFanOutProgress**](FanOutProgress.md) | Progress of a fan-out step&#39;s children, once spawned. | [optional] 
**JobId** | Pointer to **NullableString** | The job executing this step, once one has been scheduled (the first child, for a fan-out). | [optional] 
**Name** | **string** | The step&#39;s name (its address within the workflow). | 
**StartedAt** | Pointer to **NullableTime** | When the step was scheduled (its job or child workflow created). | [optional] 
**Status** | [**StepStatus**](StepStatus.md) |  | 

## Methods

### NewWorkflowStepState

`func NewWorkflowStepState(name string, status StepStatus, ) *WorkflowStepState`

NewWorkflowStepState instantiates a new WorkflowStepState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepStateWithDefaults

`func NewWorkflowStepStateWithDefaults() *WorkflowStepState`

NewWorkflowStepStateWithDefaults instantiates a new WorkflowStepState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildWorkflowId

`func (o *WorkflowStepState) GetChildWorkflowId() string`

GetChildWorkflowId returns the ChildWorkflowId field if non-nil, zero value otherwise.

### GetChildWorkflowIdOk

`func (o *WorkflowStepState) GetChildWorkflowIdOk() (*string, bool)`

GetChildWorkflowIdOk returns a tuple with the ChildWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildWorkflowId

`func (o *WorkflowStepState) SetChildWorkflowId(v string)`

SetChildWorkflowId sets ChildWorkflowId field to given value.

### HasChildWorkflowId

`func (o *WorkflowStepState) HasChildWorkflowId() bool`

HasChildWorkflowId returns a boolean if a field has been set.

### SetChildWorkflowIdNil

`func (o *WorkflowStepState) SetChildWorkflowIdNil(b bool)`

 SetChildWorkflowIdNil sets the value for ChildWorkflowId to be an explicit nil

### UnsetChildWorkflowId
`func (o *WorkflowStepState) UnsetChildWorkflowId()`

UnsetChildWorkflowId ensures that no value is present for ChildWorkflowId, not even an explicit nil
### GetCompletedAt

`func (o *WorkflowStepState) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *WorkflowStepState) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *WorkflowStepState) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *WorkflowStepState) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *WorkflowStepState) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *WorkflowStepState) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetError

`func (o *WorkflowStepState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkflowStepState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkflowStepState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WorkflowStepState) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WorkflowStepState) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WorkflowStepState) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFanOut

`func (o *WorkflowStepState) GetFanOut() FanOutProgress`

GetFanOut returns the FanOut field if non-nil, zero value otherwise.

### GetFanOutOk

`func (o *WorkflowStepState) GetFanOutOk() (*FanOutProgress, bool)`

GetFanOutOk returns a tuple with the FanOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanOut

`func (o *WorkflowStepState) SetFanOut(v FanOutProgress)`

SetFanOut sets FanOut field to given value.

### HasFanOut

`func (o *WorkflowStepState) HasFanOut() bool`

HasFanOut returns a boolean if a field has been set.

### SetFanOutNil

`func (o *WorkflowStepState) SetFanOutNil(b bool)`

 SetFanOutNil sets the value for FanOut to be an explicit nil

### UnsetFanOut
`func (o *WorkflowStepState) UnsetFanOut()`

UnsetFanOut ensures that no value is present for FanOut, not even an explicit nil
### GetJobId

`func (o *WorkflowStepState) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *WorkflowStepState) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *WorkflowStepState) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *WorkflowStepState) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *WorkflowStepState) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *WorkflowStepState) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetName

`func (o *WorkflowStepState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStepState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStepState) SetName(v string)`

SetName sets Name field to given value.


### GetStartedAt

`func (o *WorkflowStepState) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WorkflowStepState) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WorkflowStepState) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WorkflowStepState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *WorkflowStepState) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *WorkflowStepState) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *WorkflowStepState) GetStatus() StepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowStepState) GetStatusOk() (*StepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowStepState) SetStatus(v StepStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


