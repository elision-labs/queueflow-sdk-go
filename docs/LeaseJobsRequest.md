# LeaseJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseSecs** | Pointer to **NullableInt32** | Lease duration in seconds (1..&#x3D;3600, default 30). Heartbeat to extend. | [optional] 
**MaxJobs** | Pointer to **NullableInt32** | Maximum jobs to lease in one call (1..&#x3D;100, default 1). | [optional] 
**WaitSecs** | Pointer to **NullableInt32** | Long-poll wait when the queue is empty, in seconds (0..&#x3D;30, default 0). | [optional] 

## Methods

### NewLeaseJobsRequest

`func NewLeaseJobsRequest() *LeaseJobsRequest`

NewLeaseJobsRequest instantiates a new LeaseJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseJobsRequestWithDefaults

`func NewLeaseJobsRequestWithDefaults() *LeaseJobsRequest`

NewLeaseJobsRequestWithDefaults instantiates a new LeaseJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseSecs

`func (o *LeaseJobsRequest) GetLeaseSecs() int32`

GetLeaseSecs returns the LeaseSecs field if non-nil, zero value otherwise.

### GetLeaseSecsOk

`func (o *LeaseJobsRequest) GetLeaseSecsOk() (*int32, bool)`

GetLeaseSecsOk returns a tuple with the LeaseSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseSecs

`func (o *LeaseJobsRequest) SetLeaseSecs(v int32)`

SetLeaseSecs sets LeaseSecs field to given value.

### HasLeaseSecs

`func (o *LeaseJobsRequest) HasLeaseSecs() bool`

HasLeaseSecs returns a boolean if a field has been set.

### SetLeaseSecsNil

`func (o *LeaseJobsRequest) SetLeaseSecsNil(b bool)`

 SetLeaseSecsNil sets the value for LeaseSecs to be an explicit nil

### UnsetLeaseSecs
`func (o *LeaseJobsRequest) UnsetLeaseSecs()`

UnsetLeaseSecs ensures that no value is present for LeaseSecs, not even an explicit nil
### GetMaxJobs

`func (o *LeaseJobsRequest) GetMaxJobs() int32`

GetMaxJobs returns the MaxJobs field if non-nil, zero value otherwise.

### GetMaxJobsOk

`func (o *LeaseJobsRequest) GetMaxJobsOk() (*int32, bool)`

GetMaxJobsOk returns a tuple with the MaxJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobs

`func (o *LeaseJobsRequest) SetMaxJobs(v int32)`

SetMaxJobs sets MaxJobs field to given value.

### HasMaxJobs

`func (o *LeaseJobsRequest) HasMaxJobs() bool`

HasMaxJobs returns a boolean if a field has been set.

### SetMaxJobsNil

`func (o *LeaseJobsRequest) SetMaxJobsNil(b bool)`

 SetMaxJobsNil sets the value for MaxJobs to be an explicit nil

### UnsetMaxJobs
`func (o *LeaseJobsRequest) UnsetMaxJobs()`

UnsetMaxJobs ensures that no value is present for MaxJobs, not even an explicit nil
### GetWaitSecs

`func (o *LeaseJobsRequest) GetWaitSecs() int32`

GetWaitSecs returns the WaitSecs field if non-nil, zero value otherwise.

### GetWaitSecsOk

`func (o *LeaseJobsRequest) GetWaitSecsOk() (*int32, bool)`

GetWaitSecsOk returns a tuple with the WaitSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitSecs

`func (o *LeaseJobsRequest) SetWaitSecs(v int32)`

SetWaitSecs sets WaitSecs field to given value.

### HasWaitSecs

`func (o *LeaseJobsRequest) HasWaitSecs() bool`

HasWaitSecs returns a boolean if a field has been set.

### SetWaitSecsNil

`func (o *LeaseJobsRequest) SetWaitSecsNil(b bool)`

 SetWaitSecsNil sets the value for WaitSecs to be an explicit nil

### UnsetWaitSecs
`func (o *LeaseJobsRequest) UnsetWaitSecs()`

UnsetWaitSecs ensures that no value is present for WaitSecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


