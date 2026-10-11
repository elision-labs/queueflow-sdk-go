# WhoAmIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | &#x60;read_write&#x60;, or &#x60;read_only&#x60; for a &#x60;:ro&#x60; key or a &#x60;scope: read&#x60; JWT. | 
**Tenant** | **string** | The tenant every request with this credential is scoped to. | 

## Methods

### NewWhoAmIResponse

`func NewWhoAmIResponse(access string, tenant string, ) *WhoAmIResponse`

NewWhoAmIResponse instantiates a new WhoAmIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoAmIResponseWithDefaults

`func NewWhoAmIResponseWithDefaults() *WhoAmIResponse`

NewWhoAmIResponseWithDefaults instantiates a new WhoAmIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *WhoAmIResponse) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *WhoAmIResponse) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *WhoAmIResponse) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetTenant

`func (o *WhoAmIResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WhoAmIResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WhoAmIResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


