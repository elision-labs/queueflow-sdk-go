# QueuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queues** | [**[]QueueStats**](QueueStats.md) | One entry per queue that currently has non-terminal jobs, sorted by name. | 

## Methods

### NewQueuesResponse

`func NewQueuesResponse(queues []QueueStats, ) *QueuesResponse`

NewQueuesResponse instantiates a new QueuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuesResponseWithDefaults

`func NewQueuesResponseWithDefaults() *QueuesResponse`

NewQueuesResponseWithDefaults instantiates a new QueuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueues

`func (o *QueuesResponse) GetQueues() []QueueStats`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *QueuesResponse) GetQueuesOk() (*[]QueueStats, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *QueuesResponse) SetQueues(v []QueueStats)`

SetQueues sets Queues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


