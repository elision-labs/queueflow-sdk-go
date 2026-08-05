# CreateBatchJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]CreateJobRequest**](CreateJobRequest.md) |  | 

## Methods

### NewCreateBatchJobsRequest

`func NewCreateBatchJobsRequest(jobs []CreateJobRequest, ) *CreateBatchJobsRequest`

NewCreateBatchJobsRequest instantiates a new CreateBatchJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchJobsRequestWithDefaults

`func NewCreateBatchJobsRequestWithDefaults() *CreateBatchJobsRequest`

NewCreateBatchJobsRequestWithDefaults instantiates a new CreateBatchJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *CreateBatchJobsRequest) GetJobs() []CreateJobRequest`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CreateBatchJobsRequest) GetJobsOk() (*[]CreateJobRequest, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CreateBatchJobsRequest) SetJobs(v []CreateJobRequest)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


