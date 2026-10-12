# WorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Status** | [**WorkflowStatus**](WorkflowStatus.md) |  | 
**Steps** | [**[]WorkflowStepState**](WorkflowStepState.md) |  | 
**WorkflowId** | **string** |  | 

## Methods

### NewWorkflowEvent

`func NewWorkflowEvent(status WorkflowStatus, steps []WorkflowStepState, workflowId string, ) *WorkflowEvent`

NewWorkflowEvent instantiates a new WorkflowEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEventWithDefaults

`func NewWorkflowEventWithDefaults() *WorkflowEvent`

NewWorkflowEventWithDefaults instantiates a new WorkflowEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *WorkflowEvent) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WorkflowEvent) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WorkflowEvent) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WorkflowEvent) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *WorkflowEvent) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WorkflowEvent) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatus

`func (o *WorkflowEvent) GetStatus() WorkflowStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowEvent) GetStatusOk() (*WorkflowStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowEvent) SetStatus(v WorkflowStatus)`

SetStatus sets Status field to given value.


### GetSteps

`func (o *WorkflowEvent) GetSteps() []WorkflowStepState`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowEvent) GetStepsOk() (*[]WorkflowStepState, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowEvent) SetSteps(v []WorkflowStepState)`

SetSteps sets Steps field to given value.


### GetWorkflowId

`func (o *WorkflowEvent) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowEvent) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowEvent) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


