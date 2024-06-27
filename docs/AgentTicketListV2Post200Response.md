# AgentTicketListV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**AgentTicketListV2ResponseData**](AgentTicketListV2ResponseData.md) |  | [optional] 

## Methods

### NewAgentTicketListV2Response

`func NewAgentTicketListV2Response() *AgentTicketListV2Response`

NewAgentTicketListV2Response instantiates a new AgentTicketListV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTicketListV2ResponseWithDefaults

`func NewAgentTicketListV2ResponseWithDefaults() *AgentTicketListV2Response`

NewAgentTicketListV2ResponseWithDefaults instantiates a new AgentTicketListV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AgentTicketListV2Response) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AgentTicketListV2Response) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AgentTicketListV2Response) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *AgentTicketListV2Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *AgentTicketListV2Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AgentTicketListV2Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AgentTicketListV2Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AgentTicketListV2Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AgentTicketListV2Response) GetData() AgentTicketListV2ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AgentTicketListV2Response) GetDataOk() (*AgentTicketListV2ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AgentTicketListV2Response) SetData(v AgentTicketListV2ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *AgentTicketListV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


