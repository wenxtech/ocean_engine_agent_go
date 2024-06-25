# AgentTicketListV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Msg** | **string** |  | 
**Data** | [**AgentTicketListV2ResponseData**](AgentTicketListV2ResponseData.md) |  | 

## Methods

### NewAgentTicketListV2Response

`func NewAgentTicketListV2Response(code int32, msg string, data AgentTicketListV2ResponseData, ) *AgentTicketListV2Response`

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

`func (o *AgentTicketListV2Response) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AgentTicketListV2Response) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AgentTicketListV2Response) SetCode(v int32)`

SetCode sets Code field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


