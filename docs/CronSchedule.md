# CronSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJobConfig**](JobConfig.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**CronExpr** | **string** |  | 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**LastEnqueuedAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** | Unique per tenant. | 
**NextRunAt** | **time.Time** | The next instant this schedule fires. Missed occurrences (server down) collapse into at most one catch-up firing. | 
**Payload** | Pointer to **map[string]interface{}** |  | [optional] 
**QueueName** | Pointer to **NullableString** | Queue for the enqueued jobs (the engine default when absent). | [optional] 
**TaskName** | **string** |  | 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCronSchedule

`func NewCronSchedule(createdAt time.Time, cronExpr string, enabled bool, id string, name string, nextRunAt time.Time, taskName string, ) *CronSchedule`

NewCronSchedule instantiates a new CronSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronScheduleWithDefaults

`func NewCronScheduleWithDefaults() *CronSchedule`

NewCronScheduleWithDefaults instantiates a new CronSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CronSchedule) GetConfig() JobConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CronSchedule) GetConfigOk() (*JobConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CronSchedule) SetConfig(v JobConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CronSchedule) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *CronSchedule) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *CronSchedule) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetCreatedAt

`func (o *CronSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CronSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CronSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCronExpr

`func (o *CronSchedule) GetCronExpr() string`

GetCronExpr returns the CronExpr field if non-nil, zero value otherwise.

### GetCronExprOk

`func (o *CronSchedule) GetCronExprOk() (*string, bool)`

GetCronExprOk returns a tuple with the CronExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpr

`func (o *CronSchedule) SetCronExpr(v string)`

SetCronExpr sets CronExpr field to given value.


### GetEnabled

`func (o *CronSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CronSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CronSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *CronSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CronSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CronSchedule) SetId(v string)`

SetId sets Id field to given value.


### GetLastEnqueuedAt

`func (o *CronSchedule) GetLastEnqueuedAt() time.Time`

GetLastEnqueuedAt returns the LastEnqueuedAt field if non-nil, zero value otherwise.

### GetLastEnqueuedAtOk

`func (o *CronSchedule) GetLastEnqueuedAtOk() (*time.Time, bool)`

GetLastEnqueuedAtOk returns a tuple with the LastEnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEnqueuedAt

`func (o *CronSchedule) SetLastEnqueuedAt(v time.Time)`

SetLastEnqueuedAt sets LastEnqueuedAt field to given value.

### HasLastEnqueuedAt

`func (o *CronSchedule) HasLastEnqueuedAt() bool`

HasLastEnqueuedAt returns a boolean if a field has been set.

### SetLastEnqueuedAtNil

`func (o *CronSchedule) SetLastEnqueuedAtNil(b bool)`

 SetLastEnqueuedAtNil sets the value for LastEnqueuedAt to be an explicit nil

### UnsetLastEnqueuedAt
`func (o *CronSchedule) UnsetLastEnqueuedAt()`

UnsetLastEnqueuedAt ensures that no value is present for LastEnqueuedAt, not even an explicit nil
### GetName

`func (o *CronSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CronSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CronSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetNextRunAt

`func (o *CronSchedule) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *CronSchedule) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *CronSchedule) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.


### GetPayload

`func (o *CronSchedule) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CronSchedule) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CronSchedule) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CronSchedule) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetQueueName

`func (o *CronSchedule) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *CronSchedule) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *CronSchedule) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *CronSchedule) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### SetQueueNameNil

`func (o *CronSchedule) SetQueueNameNil(b bool)`

 SetQueueNameNil sets the value for QueueName to be an explicit nil

### UnsetQueueName
`func (o *CronSchedule) UnsetQueueName()`

UnsetQueueName ensures that no value is present for QueueName, not even an explicit nil
### GetTaskName

`func (o *CronSchedule) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *CronSchedule) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *CronSchedule) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTenantId

`func (o *CronSchedule) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CronSchedule) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CronSchedule) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CronSchedule) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CronSchedule) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CronSchedule) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


