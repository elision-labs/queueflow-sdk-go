# LeasedJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**Job**](Job.md) |  | 
**LeaseToken** | **string** | Opaque, unguessable proof of lease ownership, regenerated on every claim. Pass it back on heartbeat/complete/fail; a stale token (the lease expired and the job was reclaimed) is rejected. | 

## Methods

### NewLeasedJob

`func NewLeasedJob(job Job, leaseToken string, ) *LeasedJob`

NewLeasedJob instantiates a new LeasedJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeasedJobWithDefaults

`func NewLeasedJobWithDefaults() *LeasedJob`

NewLeasedJobWithDefaults instantiates a new LeasedJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *LeasedJob) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *LeasedJob) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *LeasedJob) SetJob(v Job)`

SetJob sets Job field to given value.


### GetLeaseToken

`func (o *LeasedJob) GetLeaseToken() string`

GetLeaseToken returns the LeaseToken field if non-nil, zero value otherwise.

### GetLeaseTokenOk

`func (o *LeasedJob) GetLeaseTokenOk() (*string, bool)`

GetLeaseTokenOk returns a tuple with the LeaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseToken

`func (o *LeasedJob) SetLeaseToken(v string)`

SetLeaseToken sets LeaseToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


