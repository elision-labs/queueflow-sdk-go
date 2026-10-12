# StepCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to [**[]StepCondition**](StepCondition.md) | True when every member is true. | [optional] 
**Any** | Pointer to [**[]StepCondition**](StepCondition.md) | True when at least one member is true. | [optional] 
**Not** | Pointer to [**NullableStepCondition**](StepCondition.md) |  | [optional] 
**Op** | Pointer to [**NullableConditionOp**](ConditionOp.md) |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewStepCondition

`func NewStepCondition() *StepCondition`

NewStepCondition instantiates a new StepCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepConditionWithDefaults

`func NewStepConditionWithDefaults() *StepCondition`

NewStepConditionWithDefaults instantiates a new StepCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *StepCondition) GetAll() []StepCondition`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *StepCondition) GetAllOk() (*[]StepCondition, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *StepCondition) SetAll(v []StepCondition)`

SetAll sets All field to given value.

### HasAll

`func (o *StepCondition) HasAll() bool`

HasAll returns a boolean if a field has been set.

### SetAllNil

`func (o *StepCondition) SetAllNil(b bool)`

 SetAllNil sets the value for All to be an explicit nil

### UnsetAll
`func (o *StepCondition) UnsetAll()`

UnsetAll ensures that no value is present for All, not even an explicit nil
### GetAny

`func (o *StepCondition) GetAny() []StepCondition`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *StepCondition) GetAnyOk() (*[]StepCondition, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *StepCondition) SetAny(v []StepCondition)`

SetAny sets Any field to given value.

### HasAny

`func (o *StepCondition) HasAny() bool`

HasAny returns a boolean if a field has been set.

### SetAnyNil

`func (o *StepCondition) SetAnyNil(b bool)`

 SetAnyNil sets the value for Any to be an explicit nil

### UnsetAny
`func (o *StepCondition) UnsetAny()`

UnsetAny ensures that no value is present for Any, not even an explicit nil
### GetNot

`func (o *StepCondition) GetNot() StepCondition`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *StepCondition) GetNotOk() (*StepCondition, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *StepCondition) SetNot(v StepCondition)`

SetNot sets Not field to given value.

### HasNot

`func (o *StepCondition) HasNot() bool`

HasNot returns a boolean if a field has been set.

### SetNotNil

`func (o *StepCondition) SetNotNil(b bool)`

 SetNotNil sets the value for Not to be an explicit nil

### UnsetNot
`func (o *StepCondition) UnsetNot()`

UnsetNot ensures that no value is present for Not, not even an explicit nil
### GetOp

`func (o *StepCondition) GetOp() ConditionOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StepCondition) GetOpOk() (*ConditionOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StepCondition) SetOp(v ConditionOp)`

SetOp sets Op field to given value.

### HasOp

`func (o *StepCondition) HasOp() bool`

HasOp returns a boolean if a field has been set.

### SetOpNil

`func (o *StepCondition) SetOpNil(b bool)`

 SetOpNil sets the value for Op to be an explicit nil

### UnsetOp
`func (o *StepCondition) UnsetOp()`

UnsetOp ensures that no value is present for Op, not even an explicit nil
### GetPath

`func (o *StepCondition) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StepCondition) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StepCondition) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StepCondition) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *StepCondition) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *StepCondition) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetValue

`func (o *StepCondition) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StepCondition) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StepCondition) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *StepCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StepCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StepCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


