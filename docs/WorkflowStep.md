# WorkflowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJobConfig**](JobConfig.md) |  | [optional] 
**DependsOn** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** |  | 
**OnFailure** | Pointer to [**OnFailure**](OnFailure.md) |  | [optional] 
**OnSuccess** | Pointer to [**OnSuccess**](OnSuccess.md) |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskName** | **string** |  | 

## Methods

### NewWorkflowStep

`func NewWorkflowStep(name string, taskName string, ) *WorkflowStep`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


