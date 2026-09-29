# JobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JitterFactor** | Pointer to **NullableFloat64** | Optional jitter in &#x60;0.0..&#x3D;1.0&#x60;. &#x60;0.1&#x60; &#x3D;&gt; +/-10% randomization of each retry delay, which spreads out thundering-herd retries. | [optional] 
**MaxRetries** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** | Higher is claimed first within a queue; ties break on &#x60;scheduled_at&#x60;, then &#x60;created_at&#x60;. | [optional] 
**RetryBackoff** | Pointer to [**BackoffStrategy**](BackoffStrategy.md) |  | [optional] 
**RetryDelaySecs** | Pointer to **int64** |  | [optional] 
**RetryMaxDelaySecs** | Pointer to **int64** |  | [optional] 
**TimeoutSecs** | Pointer to **int64** |  | [optional] 

## Methods

### NewJobConfig

`func NewJobConfig() *JobConfig`

NewJobConfig instantiates a new JobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigWithDefaults

`func NewJobConfigWithDefaults() *JobConfig`

NewJobConfigWithDefaults instantiates a new JobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJitterFactor

`func (o *JobConfig) GetJitterFactor() float64`

GetJitterFactor returns the JitterFactor field if non-nil, zero value otherwise.

### GetJitterFactorOk

`func (o *JobConfig) GetJitterFactorOk() (*float64, bool)`

GetJitterFactorOk returns a tuple with the JitterFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterFactor

`func (o *JobConfig) SetJitterFactor(v float64)`

SetJitterFactor sets JitterFactor field to given value.

### HasJitterFactor

`func (o *JobConfig) HasJitterFactor() bool`

HasJitterFactor returns a boolean if a field has been set.

### SetJitterFactorNil

`func (o *JobConfig) SetJitterFactorNil(b bool)`

 SetJitterFactorNil sets the value for JitterFactor to be an explicit nil

### UnsetJitterFactor
`func (o *JobConfig) UnsetJitterFactor()`

UnsetJitterFactor ensures that no value is present for JitterFactor, not even an explicit nil
### GetMaxRetries

`func (o *JobConfig) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *JobConfig) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *JobConfig) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *JobConfig) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetPriority

`func (o *JobConfig) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobConfig) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobConfig) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetryBackoff

`func (o *JobConfig) GetRetryBackoff() BackoffStrategy`

GetRetryBackoff returns the RetryBackoff field if non-nil, zero value otherwise.

### GetRetryBackoffOk

`func (o *JobConfig) GetRetryBackoffOk() (*BackoffStrategy, bool)`

GetRetryBackoffOk returns a tuple with the RetryBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryBackoff

`func (o *JobConfig) SetRetryBackoff(v BackoffStrategy)`

SetRetryBackoff sets RetryBackoff field to given value.

### HasRetryBackoff

`func (o *JobConfig) HasRetryBackoff() bool`

HasRetryBackoff returns a boolean if a field has been set.

### GetRetryDelaySecs

`func (o *JobConfig) GetRetryDelaySecs() int64`

GetRetryDelaySecs returns the RetryDelaySecs field if non-nil, zero value otherwise.

### GetRetryDelaySecsOk

`func (o *JobConfig) GetRetryDelaySecsOk() (*int64, bool)`

GetRetryDelaySecsOk returns a tuple with the RetryDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySecs

`func (o *JobConfig) SetRetryDelaySecs(v int64)`

SetRetryDelaySecs sets RetryDelaySecs field to given value.

### HasRetryDelaySecs

`func (o *JobConfig) HasRetryDelaySecs() bool`

HasRetryDelaySecs returns a boolean if a field has been set.

### GetRetryMaxDelaySecs

`func (o *JobConfig) GetRetryMaxDelaySecs() int64`

GetRetryMaxDelaySecs returns the RetryMaxDelaySecs field if non-nil, zero value otherwise.

### GetRetryMaxDelaySecsOk

`func (o *JobConfig) GetRetryMaxDelaySecsOk() (*int64, bool)`

GetRetryMaxDelaySecsOk returns a tuple with the RetryMaxDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMaxDelaySecs

`func (o *JobConfig) SetRetryMaxDelaySecs(v int64)`

SetRetryMaxDelaySecs sets RetryMaxDelaySecs field to given value.

### HasRetryMaxDelaySecs

`func (o *JobConfig) HasRetryMaxDelaySecs() bool`

HasRetryMaxDelaySecs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *JobConfig) GetTimeoutSecs() int64`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *JobConfig) GetTimeoutSecsOk() (*int64, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *JobConfig) SetTimeoutSecs(v int64)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *JobConfig) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


