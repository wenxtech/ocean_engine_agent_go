# TicketListV2ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]TicketListV2ResponseDataDataInner**](TicketListV2ResponseDataDataInner.md) |  | [optional] 

## Methods

### NewTicketListV2ResponseData

`func NewTicketListV2ResponseData() *TicketListV2ResponseData`

NewTicketListV2ResponseData instantiates a new TicketListV2ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketListV2ResponseDataWithDefaults

`func NewTicketListV2ResponseDataWithDefaults() *TicketListV2ResponseData`

NewTicketListV2ResponseDataWithDefaults instantiates a new TicketListV2ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *TicketListV2ResponseData) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TicketListV2ResponseData) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TicketListV2ResponseData) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TicketListV2ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *TicketListV2ResponseData) GetData() []TicketListV2ResponseDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TicketListV2ResponseData) GetDataOk() (*[]TicketListV2ResponseDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TicketListV2ResponseData) SetData(v []TicketListV2ResponseDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *TicketListV2ResponseData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


