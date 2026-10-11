# ReplayDeadLettersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replayed** | **int64** | How many of &#x60;results&#x60; produced a job. | 
**Results** | [**[]ReplayedDeadLetter**](ReplayedDeadLetter.md) | Entries that were replayed (&#x60;job_id&#x60; set) or could not be (&#x60;error&#x60; set), oldest first. Entries already replayed are not listed. | 

## Methods

### NewReplayDeadLettersResponse

`func NewReplayDeadLettersResponse(replayed int64, results []ReplayedDeadLetter, ) *ReplayDeadLettersResponse`

NewReplayDeadLettersResponse instantiates a new ReplayDeadLettersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayDeadLettersResponseWithDefaults

`func NewReplayDeadLettersResponseWithDefaults() *ReplayDeadLettersResponse`

NewReplayDeadLettersResponseWithDefaults instantiates a new ReplayDeadLettersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplayed

`func (o *ReplayDeadLettersResponse) GetReplayed() int64`

GetReplayed returns the Replayed field if non-nil, zero value otherwise.

### GetReplayedOk

`func (o *ReplayDeadLettersResponse) GetReplayedOk() (*int64, bool)`

GetReplayedOk returns a tuple with the Replayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayed

`func (o *ReplayDeadLettersResponse) SetReplayed(v int64)`

SetReplayed sets Replayed field to given value.


### GetResults

`func (o *ReplayDeadLettersResponse) GetResults() []ReplayedDeadLetter`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ReplayDeadLettersResponse) GetResultsOk() (*[]ReplayedDeadLetter, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ReplayDeadLettersResponse) SetResults(v []ReplayedDeadLetter)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


