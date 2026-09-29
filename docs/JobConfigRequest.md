# JobConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JitterFactor** | Pointer to **NullableFloat64** | Retry-delay jitter in &#x60;0.0..&#x3D;1.0&#x60; (e.g. &#x60;0.1&#x60; &#x3D; +/-10%). | [optional] 
**MaxRetries** | Pointer to **NullableInt32** |  | [optional] 
**Priority** | Pointer to **NullableInt32** | Higher is claimed first within a queue (ties: oldest first). | [optional] 
**Queue** | Pointer to **NullableString** | Override the destination queue. | [optional] 
**RetryBackoff** | Pointer to [**NullableBackoffStrategy**](BackoffStrategy.md) | How retry delays grow between attempts (default exponential). | [optional] 
**RetryDelaySecs** | Pointer to **NullableInt64** | Base retry delay, in seconds. | [optional] 
**RetryMaxDelaySecs** | Pointer to **NullableInt64** | Upper bound on any computed retry delay, in seconds. | [optional] 
**Timeout** | Pointer to **NullableInt64** | Per-attempt timeout, in seconds. | [optional] 

## Methods

### NewJobConfigRequest

`func NewJobConfigRequest() *JobConfigRequest`

NewJobConfigRequest instantiates a new JobConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigRequestWithDefaults

`func NewJobConfigRequestWithDefaults() *JobConfigRequest`

NewJobConfigRequestWithDefaults instantiates a new JobConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJitterFactor

`func (o *JobConfigRequest) GetJitterFactor() float64`

GetJitterFactor returns the JitterFactor field if non-nil, zero value otherwise.

### GetJitterFactorOk

`func (o *JobConfigRequest) GetJitterFactorOk() (*float64, bool)`

GetJitterFactorOk returns a tuple with the JitterFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitterFactor

`func (o *JobConfigRequest) SetJitterFactor(v float64)`

SetJitterFactor sets JitterFactor field to given value.

### HasJitterFactor

`func (o *JobConfigRequest) HasJitterFactor() bool`

HasJitterFactor returns a boolean if a field has been set.

### SetJitterFactorNil

`func (o *JobConfigRequest) SetJitterFactorNil(b bool)`

 SetJitterFactorNil sets the value for JitterFactor to be an explicit nil

### UnsetJitterFactor
`func (o *JobConfigRequest) UnsetJitterFactor()`

UnsetJitterFactor ensures that no value is present for JitterFactor, not even an explicit nil
### GetMaxRetries

`func (o *JobConfigRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *JobConfigRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *JobConfigRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *JobConfigRequest) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *JobConfigRequest) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *JobConfigRequest) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetPriority

`func (o *JobConfigRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *JobConfigRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *JobConfigRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *JobConfigRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *JobConfigRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *JobConfigRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetQueue

`func (o *JobConfigRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *JobConfigRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *JobConfigRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *JobConfigRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *JobConfigRequest) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *JobConfigRequest) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRetryBackoff

`func (o *JobConfigRequest) GetRetryBackoff() BackoffStrategy`

GetRetryBackoff returns the RetryBackoff field if non-nil, zero value otherwise.

### GetRetryBackoffOk

`func (o *JobConfigRequest) GetRetryBackoffOk() (*BackoffStrategy, bool)`

GetRetryBackoffOk returns a tuple with the RetryBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryBackoff

`func (o *JobConfigRequest) SetRetryBackoff(v BackoffStrategy)`

SetRetryBackoff sets RetryBackoff field to given value.

### HasRetryBackoff

`func (o *JobConfigRequest) HasRetryBackoff() bool`

HasRetryBackoff returns a boolean if a field has been set.

### SetRetryBackoffNil

`func (o *JobConfigRequest) SetRetryBackoffNil(b bool)`

 SetRetryBackoffNil sets the value for RetryBackoff to be an explicit nil

### UnsetRetryBackoff
`func (o *JobConfigRequest) UnsetRetryBackoff()`

UnsetRetryBackoff ensures that no value is present for RetryBackoff, not even an explicit nil
### GetRetryDelaySecs

`func (o *JobConfigRequest) GetRetryDelaySecs() int64`

GetRetryDelaySecs returns the RetryDelaySecs field if non-nil, zero value otherwise.

### GetRetryDelaySecsOk

`func (o *JobConfigRequest) GetRetryDelaySecsOk() (*int64, bool)`

GetRetryDelaySecsOk returns a tuple with the RetryDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySecs

`func (o *JobConfigRequest) SetRetryDelaySecs(v int64)`

SetRetryDelaySecs sets RetryDelaySecs field to given value.

### HasRetryDelaySecs

`func (o *JobConfigRequest) HasRetryDelaySecs() bool`

HasRetryDelaySecs returns a boolean if a field has been set.

### SetRetryDelaySecsNil

`func (o *JobConfigRequest) SetRetryDelaySecsNil(b bool)`

 SetRetryDelaySecsNil sets the value for RetryDelaySecs to be an explicit nil

### UnsetRetryDelaySecs
`func (o *JobConfigRequest) UnsetRetryDelaySecs()`

UnsetRetryDelaySecs ensures that no value is present for RetryDelaySecs, not even an explicit nil
### GetRetryMaxDelaySecs

`func (o *JobConfigRequest) GetRetryMaxDelaySecs() int64`

GetRetryMaxDelaySecs returns the RetryMaxDelaySecs field if non-nil, zero value otherwise.

### GetRetryMaxDelaySecsOk

`func (o *JobConfigRequest) GetRetryMaxDelaySecsOk() (*int64, bool)`

GetRetryMaxDelaySecsOk returns a tuple with the RetryMaxDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryMaxDelaySecs

`func (o *JobConfigRequest) SetRetryMaxDelaySecs(v int64)`

SetRetryMaxDelaySecs sets RetryMaxDelaySecs field to given value.

### HasRetryMaxDelaySecs

`func (o *JobConfigRequest) HasRetryMaxDelaySecs() bool`

HasRetryMaxDelaySecs returns a boolean if a field has been set.

### SetRetryMaxDelaySecsNil

`func (o *JobConfigRequest) SetRetryMaxDelaySecsNil(b bool)`

 SetRetryMaxDelaySecsNil sets the value for RetryMaxDelaySecs to be an explicit nil

### UnsetRetryMaxDelaySecs
`func (o *JobConfigRequest) UnsetRetryMaxDelaySecs()`

UnsetRetryMaxDelaySecs ensures that no value is present for RetryMaxDelaySecs, not even an explicit nil
### GetTimeout

`func (o *JobConfigRequest) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *JobConfigRequest) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *JobConfigRequest) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *JobConfigRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *JobConfigRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *JobConfigRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


