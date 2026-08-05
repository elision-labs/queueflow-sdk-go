# CreateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**Steps** | [**[]WorkflowStep**](WorkflowStep.md) |  | 

## Methods

### NewCreateWorkflowRequest

`func NewCreateWorkflowRequest(name string, steps []WorkflowStep, ) *CreateWorkflowRequest`

NewCreateWorkflowRequest instantiates a new CreateWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowRequestWithDefaults

`func NewCreateWorkflowRequestWithDefaults() *CreateWorkflowRequest`

NewCreateWorkflowRequestWithDefaults instantiates a new CreateWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CreateWorkflowRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateWorkflowRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateWorkflowRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateWorkflowRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateWorkflowRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateWorkflowRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateWorkflowRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateWorkflowRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSteps

`func (o *CreateWorkflowRequest) GetSteps() []WorkflowStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateWorkflowRequest) GetStepsOk() (*[]WorkflowStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateWorkflowRequest) SetSteps(v []WorkflowStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


