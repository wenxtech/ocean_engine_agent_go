# TicketListV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**TicketListV2ResponseData**](TicketListV2ResponseData.md) |  | [optional] 

## Methods

### NewTicketListV2Response

`func NewTicketListV2Response() *TicketListV2Response`

NewTicketListV2Response instantiates a new TicketListV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketListV2ResponseWithDefaults

`func NewTicketListV2ResponseWithDefaults() *TicketListV2Response`

NewTicketListV2ResponseWithDefaults instantiates a new TicketListV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TicketListV2Response) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TicketListV2Response) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TicketListV2Response) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *TicketListV2Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *TicketListV2Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TicketListV2Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TicketListV2Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TicketListV2Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *TicketListV2Response) GetData() TicketListV2ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TicketListV2Response) GetDataOk() (*TicketListV2ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TicketListV2Response) SetData(v TicketListV2ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TicketListV2Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


