# HeartbeatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendSecs** | **int32** | New lease duration in seconds, measured from now (1..&#x3D;3600). | 
**LeaseToken** | **string** | The lease token returned by the lease call. | 

## Methods

### NewHeartbeatRequest

`func NewHeartbeatRequest(extendSecs int32, leaseToken string, ) *HeartbeatRequest`

NewHeartbeatRequest instantiates a new HeartbeatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeartbeatRequestWithDefaults

`func NewHeartbeatRequestWithDefaults() *HeartbeatRequest`

NewHeartbeatRequestWithDefaults instantiates a new HeartbeatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendSecs

`func (o *HeartbeatRequest) GetExtendSecs() int32`

GetExtendSecs returns the ExtendSecs field if non-nil, zero value otherwise.

### GetExtendSecsOk

`func (o *HeartbeatRequest) GetExtendSecsOk() (*int32, bool)`

GetExtendSecsOk returns a tuple with the ExtendSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendSecs

`func (o *HeartbeatRequest) SetExtendSecs(v int32)`

SetExtendSecs sets ExtendSecs field to given value.


### GetLeaseToken

`func (o *HeartbeatRequest) GetLeaseToken() string`

GetLeaseToken returns the LeaseToken field if non-nil, zero value otherwise.

### GetLeaseTokenOk

`func (o *HeartbeatRequest) GetLeaseTokenOk() (*string, bool)`

GetLeaseTokenOk returns a tuple with the LeaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseToken

`func (o *HeartbeatRequest) SetLeaseToken(v string)`

SetLeaseToken sets LeaseToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


