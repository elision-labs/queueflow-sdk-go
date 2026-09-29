# WorkflowStepState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **NullableString** | The job executing this step, once one has been scheduled. | [optional] 
**Name** | **string** | The step&#39;s name (its address within the workflow). | 
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


