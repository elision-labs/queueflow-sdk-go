# CreateJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJobConfigRequest**](JobConfigRequest.md) |  | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Arbitrary JSON object passed to the handler. | [optional] 
**RunAt** | Pointer to **NullableTime** | Don&#39;t run before this instant (RFC 3339). The job is created immediately but stays invisible to workers until then. | [optional] 
**TaskName** | **string** | The registered task handler to invoke. | 

## Methods

### NewCreateJobRequest

`func NewCreateJobRequest(taskName string, ) *CreateJobRequest`

NewCreateJobRequest instantiates a new CreateJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobRequestWithDefaults

`func NewCreateJobRequestWithDefaults() *CreateJobRequest`

NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateJobRequest) GetConfig() JobConfigRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateJobRequest) GetConfigOk() (*JobConfigRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateJobRequest) SetConfig(v JobConfigRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateJobRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *CreateJobRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *CreateJobRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetPayload

`func (o *CreateJobRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CreateJobRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CreateJobRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CreateJobRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRunAt

`func (o *CreateJobRequest) GetRunAt() time.Time`

GetRunAt returns the RunAt field if non-nil, zero value otherwise.

### GetRunAtOk

`func (o *CreateJobRequest) GetRunAtOk() (*time.Time, bool)`

GetRunAtOk returns a tuple with the RunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAt

`func (o *CreateJobRequest) SetRunAt(v time.Time)`

SetRunAt sets RunAt field to given value.

### HasRunAt

`func (o *CreateJobRequest) HasRunAt() bool`

HasRunAt returns a boolean if a field has been set.

### SetRunAtNil

`func (o *CreateJobRequest) SetRunAtNil(b bool)`

 SetRunAtNil sets the value for RunAt to be an explicit nil

### UnsetRunAt
`func (o *CreateJobRequest) UnsetRunAt()`

UnsetRunAt ensures that no value is present for RunAt, not even an explicit nil
### GetTaskName

`func (o *CreateJobRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *CreateJobRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *CreateJobRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


