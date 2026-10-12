# FanOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemKey** | Pointer to **NullableString** | Payload key the element is passed under (default &#x60;item&#x60;). | [optional] 
**MaxItems** | Pointer to **NullableInt32** | Upper bound on elements; more than this fails the step before any child runs. Default [&#x60;limits::MAX_FAN_OUT&#x60;]. | [optional] 
**Path** | **string** | Dotted path into the context, for example &#x60;fetch.images&#x60;. | 

## Methods

### NewFanOut

`func NewFanOut(path string, ) *FanOut`

NewFanOut instantiates a new FanOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanOutWithDefaults

`func NewFanOutWithDefaults() *FanOut`

NewFanOutWithDefaults instantiates a new FanOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemKey

`func (o *FanOut) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *FanOut) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *FanOut) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *FanOut) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### SetItemKeyNil

`func (o *FanOut) SetItemKeyNil(b bool)`

 SetItemKeyNil sets the value for ItemKey to be an explicit nil

### UnsetItemKey
`func (o *FanOut) UnsetItemKey()`

UnsetItemKey ensures that no value is present for ItemKey, not even an explicit nil
### GetMaxItems

`func (o *FanOut) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *FanOut) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *FanOut) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *FanOut) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *FanOut) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *FanOut) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetPath

`func (o *FanOut) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FanOut) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FanOut) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


