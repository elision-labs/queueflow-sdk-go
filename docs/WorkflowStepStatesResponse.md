# WorkflowStepStatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Steps** | [**[]WorkflowStepState**](WorkflowStepState.md) | One entry per step, in declaration order. | 

## Methods

### NewWorkflowStepStatesResponse

`func NewWorkflowStepStatesResponse(steps []WorkflowStepState, ) *WorkflowStepStatesResponse`

NewWorkflowStepStatesResponse instantiates a new WorkflowStepStatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStepStatesResponseWithDefaults

`func NewWorkflowStepStatesResponseWithDefaults() *WorkflowStepStatesResponse`

NewWorkflowStepStatesResponseWithDefaults instantiates a new WorkflowStepStatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteps

`func (o *WorkflowStepStatesResponse) GetSteps() []WorkflowStepState`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowStepStatesResponse) GetStepsOk() (*[]WorkflowStepState, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowStepStatesResponse) SetSteps(v []WorkflowStepState)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


