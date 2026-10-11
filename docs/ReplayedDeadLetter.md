# ReplayedDeadLetter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeadLetterId** | **int64** |  | 
**Error** | Pointer to **NullableString** | Why it did not, for example the original job was purged by retention. | [optional] 
**JobId** | Pointer to **NullableString** | The fresh job, when the replay succeeded. | [optional] 

## Methods

### NewReplayedDeadLetter

`func NewReplayedDeadLetter(deadLetterId int64, ) *ReplayedDeadLetter`

NewReplayedDeadLetter instantiates a new ReplayedDeadLetter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayedDeadLetterWithDefaults

`func NewReplayedDeadLetterWithDefaults() *ReplayedDeadLetter`

NewReplayedDeadLetterWithDefaults instantiates a new ReplayedDeadLetter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadLetterId

`func (o *ReplayedDeadLetter) GetDeadLetterId() int64`

GetDeadLetterId returns the DeadLetterId field if non-nil, zero value otherwise.

### GetDeadLetterIdOk

`func (o *ReplayedDeadLetter) GetDeadLetterIdOk() (*int64, bool)`

GetDeadLetterIdOk returns a tuple with the DeadLetterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterId

`func (o *ReplayedDeadLetter) SetDeadLetterId(v int64)`

SetDeadLetterId sets DeadLetterId field to given value.


### GetError

`func (o *ReplayedDeadLetter) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReplayedDeadLetter) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReplayedDeadLetter) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ReplayedDeadLetter) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ReplayedDeadLetter) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ReplayedDeadLetter) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetJobId

`func (o *ReplayedDeadLetter) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ReplayedDeadLetter) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ReplayedDeadLetter) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ReplayedDeadLetter) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *ReplayedDeadLetter) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *ReplayedDeadLetter) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


