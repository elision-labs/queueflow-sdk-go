# CreateBatchJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**JobIds** | **[]string** |  | 

## Methods

### NewCreateBatchJobsResponse

`func NewCreateBatchJobsResponse(count int32, jobIds []string, ) *CreateBatchJobsResponse`

NewCreateBatchJobsResponse instantiates a new CreateBatchJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchJobsResponseWithDefaults

`func NewCreateBatchJobsResponseWithDefaults() *CreateBatchJobsResponse`

NewCreateBatchJobsResponseWithDefaults instantiates a new CreateBatchJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CreateBatchJobsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateBatchJobsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateBatchJobsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetJobIds

`func (o *CreateBatchJobsResponse) GetJobIds() []string`

GetJobIds returns the JobIds field if non-nil, zero value otherwise.

### GetJobIdsOk

`func (o *CreateBatchJobsResponse) GetJobIdsOk() (*[]string, bool)`

GetJobIdsOk returns a tuple with the JobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIds

`func (o *CreateBatchJobsResponse) SetJobIds(v []string)`

SetJobIds sets JobIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


