# WorkflowDiagramResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diagram** | **string** | The diagram document (Mermaid &#x60;graph TD&#x60;). | 
**Format** | **string** | Diagram source format. Always &#x60;mermaid&#x60; today. | 

## Methods

### NewWorkflowDiagramResponse

`func NewWorkflowDiagramResponse(diagram string, format string, ) *WorkflowDiagramResponse`

NewWorkflowDiagramResponse instantiates a new WorkflowDiagramResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowDiagramResponseWithDefaults

`func NewWorkflowDiagramResponseWithDefaults() *WorkflowDiagramResponse`

NewWorkflowDiagramResponseWithDefaults instantiates a new WorkflowDiagramResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiagram

`func (o *WorkflowDiagramResponse) GetDiagram() string`

GetDiagram returns the Diagram field if non-nil, zero value otherwise.

### GetDiagramOk

`func (o *WorkflowDiagramResponse) GetDiagramOk() (*string, bool)`

GetDiagramOk returns a tuple with the Diagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagram

`func (o *WorkflowDiagramResponse) SetDiagram(v string)`

SetDiagram sets Diagram field to given value.


### GetFormat

`func (o *WorkflowDiagramResponse) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WorkflowDiagramResponse) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WorkflowDiagramResponse) SetFormat(v string)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


