# AgentTicketListV2ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Data** | [**[]AgentTicketListV2ResponseDataDataInner**](AgentTicketListV2ResponseDataDataInner.md) |  | 

## Methods

### NewAgentTicketListV2ResponseData

`func NewAgentTicketListV2ResponseData(total int32, data []AgentTicketListV2ResponseDataDataInner, ) *AgentTicketListV2ResponseData`

NewAgentTicketListV2ResponseData instantiates a new AgentTicketListV2ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTicketListV2ResponseDataWithDefaults

`func NewAgentTicketListV2ResponseDataWithDefaults() *AgentTicketListV2ResponseData`

NewAgentTicketListV2ResponseDataWithDefaults instantiates a new AgentTicketListV2ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AgentTicketListV2ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AgentTicketListV2ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AgentTicketListV2ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetData

`func (o *AgentTicketListV2ResponseData) GetData() []AgentTicketListV2ResponseDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentTicketListV2ResponseData) GetDataOk() (*[]AgentTicketListV2ResponseDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentTicketListV2ResponseData) SetData(v []AgentTicketListV2ResponseDataDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


