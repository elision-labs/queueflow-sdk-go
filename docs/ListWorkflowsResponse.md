# ListWorkflowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Limit** | **int64** |  | 
**Offset** | **int64** |  | 
**Total** | Pointer to **NullableInt64** | Exact total match count; only present when &#x60;include_total&#x3D;true&#x60;. | [optional] 
**Workflows** | [**[]Workflow**](Workflow.md) |  | 

## Methods

### NewListWorkflowsResponse

`func NewListWorkflowsResponse(hasMore bool, limit int64, offset int64, workflows []Workflow, ) *ListWorkflowsResponse`

NewListWorkflowsResponse instantiates a new ListWorkflowsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowsResponseWithDefaults

`func NewListWorkflowsResponseWithDefaults() *ListWorkflowsResponse`

NewListWorkflowsResponseWithDefaults instantiates a new ListWorkflowsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListWorkflowsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListWorkflowsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListWorkflowsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLimit

`func (o *ListWorkflowsResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListWorkflowsResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListWorkflowsResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *ListWorkflowsResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListWorkflowsResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListWorkflowsResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *ListWorkflowsResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListWorkflowsResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListWorkflowsResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListWorkflowsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ListWorkflowsResponse) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ListWorkflowsResponse) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetWorkflows

`func (o *ListWorkflowsResponse) GetWorkflows() []Workflow`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ListWorkflowsResponse) GetWorkflowsOk() (*[]Workflow, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ListWorkflowsResponse) SetWorkflows(v []Workflow)`

SetWorkflows sets Workflows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


