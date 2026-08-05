# ListDeadLettersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeadLetters** | [**[]DeadLetter**](DeadLetter.md) |  | 
**HasMore** | **bool** |  | 
**Limit** | **int64** |  | 
**Offset** | **int64** |  | 
**Total** | Pointer to **NullableInt64** | Exact total match count; only present when &#x60;include_total&#x3D;true&#x60;. | [optional] 

## Methods

### NewListDeadLettersResponse

`func NewListDeadLettersResponse(deadLetters []DeadLetter, hasMore bool, limit int64, offset int64, ) *ListDeadLettersResponse`

NewListDeadLettersResponse instantiates a new ListDeadLettersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeadLettersResponseWithDefaults

`func NewListDeadLettersResponseWithDefaults() *ListDeadLettersResponse`

NewListDeadLettersResponseWithDefaults instantiates a new ListDeadLettersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadLetters

`func (o *ListDeadLettersResponse) GetDeadLetters() []DeadLetter`

GetDeadLetters returns the DeadLetters field if non-nil, zero value otherwise.

### GetDeadLettersOk

`func (o *ListDeadLettersResponse) GetDeadLettersOk() (*[]DeadLetter, bool)`

GetDeadLettersOk returns a tuple with the DeadLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetters

`func (o *ListDeadLettersResponse) SetDeadLetters(v []DeadLetter)`

SetDeadLetters sets DeadLetters field to given value.


### GetHasMore

`func (o *ListDeadLettersResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListDeadLettersResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListDeadLettersResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLimit

`func (o *ListDeadLettersResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListDeadLettersResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListDeadLettersResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListDeadLettersResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListDeadLettersResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListDeadLettersResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListDeadLettersResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListDeadLettersResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListDeadLettersResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListDeadLettersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListDeadLettersResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListDeadLettersResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


