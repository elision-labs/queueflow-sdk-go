# ListCronsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crons** | [**[]CronSchedule**](CronSchedule.md) |  | 
**HasMore** | **bool** |  | 
**Limit** | **int64** |  | 
**NextCursor** | Pointer to **NullableString** | Opaque keyset cursor for the next page (present when &#x60;has_more&#x60;). Pass it back as &#x60;cursor&#x60; to continue where this page ended; cheaper than deep OFFSET paging. | [optional] 
**Offset** | **int64** |  | 
**Total** | Pointer to **NullableInt64** | Exact total match count; only present when &#x60;include_total&#x3D;true&#x60;. | [optional] 

## Methods

### NewListCronsResponse

`func NewListCronsResponse(crons []CronSchedule, hasMore bool, limit int64, offset int64, ) *ListCronsResponse`

NewListCronsResponse instantiates a new ListCronsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCronsResponseWithDefaults

`func NewListCronsResponseWithDefaults() *ListCronsResponse`

NewListCronsResponseWithDefaults instantiates a new ListCronsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrons

`func (o *ListCronsResponse) GetCrons() []CronSchedule`

GetCrons returns the Crons field if non-nil, zero value otherwise.

### GetCronsOk

`func (o *ListCronsResponse) GetCronsOk() (*[]CronSchedule, bool)`

GetCronsOk returns a tuple with the Crons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrons

`func (o *ListCronsResponse) SetCrons(v []CronSchedule)`

SetCrons sets Crons field to given value.


### GetHasMore

`func (o *ListCronsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListCronsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListCronsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLimit

`func (o *ListCronsResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCronsResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCronsResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetNextCursor

`func (o *ListCronsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListCronsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListCronsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListCronsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListCronsResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListCronsResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetOffset

`func (o *ListCronsResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListCronsResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListCronsResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListCronsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListCronsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListCronsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListCronsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListCronsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListCronsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


