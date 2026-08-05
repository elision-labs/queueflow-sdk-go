# ListJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Jobs** | [**[]Job**](Job.md) |  | 
**Limit** | **int64** |  | 
**Offset** | **int64** |  | 
**Total** | Pointer to **NullableInt64** | Exact total match count. Only present when the request set &#x60;include_total&#x3D;true&#x60;; computing it costs a full count over the filtered set, so it is opt-in. | [optional] 

## Methods

### NewListJobsResponse

`func NewListJobsResponse(hasMore bool, jobs []Job, limit int64, offset int64, ) *ListJobsResponse`

NewListJobsResponse instantiates a new ListJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobsResponseWithDefaults

`func NewListJobsResponseWithDefaults() *ListJobsResponse`

NewListJobsResponseWithDefaults instantiates a new ListJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListJobsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListJobsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListJobsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetJobs

`func (o *ListJobsResponse) GetJobs() []Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListJobsResponse) GetJobsOk() (*[]Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListJobsResponse) SetJobs(v []Job)`

SetJobs sets Jobs field to given value.


### GetLimit

`func (o *ListJobsResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListJobsResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListJobsResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListJobsResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListJobsResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListJobsResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListJobsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListJobsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListJobsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListJobsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListJobsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListJobsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


