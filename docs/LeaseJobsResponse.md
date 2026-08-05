# LeaseJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]LeasedJob**](LeasedJob.md) |  | 

## Methods

### NewLeaseJobsResponse

`func NewLeaseJobsResponse(jobs []LeasedJob, ) *LeaseJobsResponse`

NewLeaseJobsResponse instantiates a new LeaseJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseJobsResponseWithDefaults

`func NewLeaseJobsResponseWithDefaults() *LeaseJobsResponse`

NewLeaseJobsResponseWithDefaults instantiates a new LeaseJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *LeaseJobsResponse) GetJobs() []LeasedJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *LeaseJobsResponse) GetJobsOk() (*[]LeasedJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *LeaseJobsResponse) SetJobs(v []LeasedJob)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


