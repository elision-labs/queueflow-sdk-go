# WorkflowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJobConfig**](JobConfig.md) |  | [optional] 
**DependsOn** | Pointer to **[]string** |  | [optional] 
**FanOut** | Pointer to [**NullableFanOut**](FanOut.md) | Run the task once per element of an array in the context. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**OnFailure** | Pointer to [**OnFailure**](OnFailure.md) |  | [optional] 
**OnSuccess** | Pointer to [**OnSuccess**](OnSuccess.md) |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Queue** | Pointer to **NullableString** | Queue for this step&#39;s job(s); the server&#39;s default queue when absent. | [optional] 
**TaskName** | Pointer to **string** | The task handler to run. May be omitted for a sub-workflow step. | [optional] 
**When** | Pointer to [**NullableStepCondition**](StepCondition.md) | Run only when this condition holds against the workflow context at the moment the step becomes ready; otherwise the step is skipped with reason &#x60;condition not met&#x60; and its dependents proceed. | [optional] 
**Workflow** | Pointer to [**NullableCreateWorkflowRequest**](CreateWorkflowRequest.md) | Run an inline child workflow instead of a task. The child&#39;s context is its own declared context plus this step&#39;s payload, with the parent&#39;s context under &#x60;_parent&#x60;; the step&#39;s result is the child&#39;s final context. | [optional] 

## Methods

### NewWorkflowStep

`func NewWorkflowStep(name string, ) *WorkflowStep`

NewWorkflowStep instantiates a new WorkflowStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepWithDefaults

`func NewWorkflowStepWithDefaults() *WorkflowStep`

NewWorkflowStepWithDefaults instantiates a new WorkflowStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *WorkflowStep) GetConfig() JobConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkflowStep) GetConfigOk() (*JobConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkflowStep) SetConfig(v JobConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkflowStep) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *WorkflowStep) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *WorkflowStep) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDependsOn

`func (o *WorkflowStep) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *WorkflowStep) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *WorkflowStep) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *WorkflowStep) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetFanOut

`func (o *WorkflowStep) GetFanOut() FanOut`

GetFanOut returns the FanOut field if non-nil, zero value otherwise.

### GetFanOutOk

`func (o *WorkflowStep) GetFanOutOk() (*FanOut, bool)`

GetFanOutOk returns a tuple with the FanOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanOut

`func (o *WorkflowStep) SetFanOut(v FanOut)`

SetFanOut sets FanOut field to given value.

### HasFanOut

`func (o *WorkflowStep) HasFanOut() bool`

HasFanOut returns a boolean if a field has been set.

### SetFanOutNil

`func (o *WorkflowStep) SetFanOutNil(b bool)`

 SetFanOutNil sets the value for FanOut to be an explicit nil

### UnsetFanOut
`func (o *WorkflowStep) UnsetFanOut()`

UnsetFanOut ensures that no value is present for FanOut, not even an explicit nil
### GetMetadata

`func (o *WorkflowStep) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WorkflowStep) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WorkflowStep) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WorkflowStep) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *WorkflowStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowStep) SetName(v string)`

SetName sets Name field to given value.


### GetOnFailure

`func (o *WorkflowStep) GetOnFailure() OnFailure`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *WorkflowStep) GetOnFailureOk() (*OnFailure, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *WorkflowStep) SetOnFailure(v OnFailure)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *WorkflowStep) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnSuccess

`func (o *WorkflowStep) GetOnSuccess() OnSuccess`

GetOnSuccess returns the OnSuccess field if non-nil, zero value otherwise.

### GetOnSuccessOk

`func (o *WorkflowStep) GetOnSuccessOk() (*OnSuccess, bool)`

GetOnSuccessOk returns a tuple with the OnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnSuccess

`func (o *WorkflowStep) SetOnSuccess(v OnSuccess)`

SetOnSuccess sets OnSuccess field to given value.

### HasOnSuccess

`func (o *WorkflowStep) HasOnSuccess() bool`

HasOnSuccess returns a boolean if a field has been set.

### GetPayload

`func (o *WorkflowStep) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WorkflowStep) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WorkflowStep) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WorkflowStep) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetQueue

`func (o *WorkflowStep) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *WorkflowStep) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *WorkflowStep) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *WorkflowStep) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *WorkflowStep) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *WorkflowStep) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTaskName

`func (o *WorkflowStep) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *WorkflowStep) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *WorkflowStep) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *WorkflowStep) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetWhen

`func (o *WorkflowStep) GetWhen() StepCondition`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *WorkflowStep) GetWhenOk() (*StepCondition, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *WorkflowStep) SetWhen(v StepCondition)`

SetWhen sets When field to given value.

### HasWhen

`func (o *WorkflowStep) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### SetWhenNil

`func (o *WorkflowStep) SetWhenNil(b bool)`

 SetWhenNil sets the value for When to be an explicit nil

### UnsetWhen
`func (o *WorkflowStep) UnsetWhen()`

UnsetWhen ensures that no value is present for When, not even an explicit nil
### GetWorkflow

`func (o *WorkflowStep) GetWorkflow() CreateWorkflowRequest`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowStep) GetWorkflowOk() (*CreateWorkflowRequest, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowStep) SetWorkflow(v CreateWorkflowRequest)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WorkflowStep) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *WorkflowStep) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *WorkflowStep) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


