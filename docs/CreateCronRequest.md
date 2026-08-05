# CreateCronRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJobConfig**](JobConfig.md) |  | [optional] 
**CronExpr** | **string** | 5-field crontab (UTC); 6/7 fields with leading seconds also accepted. | 
**Name** | **string** | Unique per tenant. | 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**TaskName** | **string** |  | 

## Methods

### NewCreateCronRequest

`func NewCreateCronRequest(cronExpr string, name string, taskName string, ) *CreateCronRequest`

NewCreateCronRequest instantiates a new CreateCronRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCronRequestWithDefaults

`func NewCreateCronRequestWithDefaults() *CreateCronRequest`

NewCreateCronRequestWithDefaults instantiates a new CreateCronRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateCronRequest) GetConfig() JobConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateCronRequest) GetConfigOk() (*JobConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateCronRequest) SetConfig(v JobConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateCronRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *CreateCronRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *CreateCronRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetCronExpr

`func (o *CreateCronRequest) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *CreateCronRequest) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *CreateCronRequest) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.


### GetName

`func (o *CreateCronRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCronRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCronRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *CreateCronRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateCronRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateCronRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CreateCronRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetQueue

`func (o *CreateCronRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CreateCronRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CreateCronRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CreateCronRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *CreateCronRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *CreateCronRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetTaskName

`func (o *CreateCronRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *CreateCronRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *CreateCronRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


