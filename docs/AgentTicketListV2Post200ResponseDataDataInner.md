# AgentTicketListV2ResponseDataDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | Pointer to **int64** |  | [optional] 
**TicketSubject** | Pointer to **string** |  | [optional] 
**TicketStatus** | Pointer to **string** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **string** |  | [optional] 
**CurrentProcessPersons** | Pointer to **[]string** |  | [optional] 
**TaskName** | Pointer to **string** |  | [optional] 
**AdvId** | Pointer to **int64** |  | [optional] 
**AdvName** | Pointer to **string** |  | [optional] 
**AgentName** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **int64** |  | [optional] 
**AgentCompanyName** | Pointer to **string** |  | [optional] 
**AgentCompanyId** | Pointer to **int64** |  | [optional] 
**SalesName** | Pointer to **string** |  | [optional] 
**SalesDepId** | Pointer to **int64** |  | [optional] 
**SalesDepName** | Pointer to **string** |  | [optional] 
**ProcessPersons** | Pointer to **[]string** |  | [optional] 
**TicketStatusName** | Pointer to **string** |  | [optional] 
**TicketApp** | Pointer to **string** |  | [optional] 
**AgentSalesId** | Pointer to **int64** |  | [optional] 
**AgentSalesName** | Pointer to **string** |  | [optional] 
**OptimizerId** | Pointer to **int64** |  | [optional] 
**OptimizerName** | Pointer to **string** |  | [optional] 
**AthenaOrderId** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentTicketListV2ResponseDataDataInner

`func NewAgentTicketListV2ResponseDataDataInner() *AgentTicketListV2ResponseDataDataInner`

NewAgentTicketListV2ResponseDataDataInner instantiates a new AgentTicketListV2ResponseDataDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTicketListV2ResponseDataDataInnerWithDefaults

`func NewAgentTicketListV2ResponseDataDataInnerWithDefaults() *AgentTicketListV2ResponseDataDataInner`

NewAgentTicketListV2ResponseDataDataInnerWithDefaults instantiates a new AgentTicketListV2ResponseDataDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTicketId

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketId() int64`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketIdOk() (*int64, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketId(v int64)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *AgentTicketListV2ResponseDataDataInner) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### GetTicketSubject

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketSubject() string`

GetTicketSubject returns the TicketSubject field if non-nil, zero value otherwise.

### GetTicketSubjectOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketSubjectOk() (*string, bool)`

GetTicketSubjectOk returns a tuple with the TicketSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSubject

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketSubject(v string)`

SetTicketSubject sets TicketSubject field to given value.

### HasTicketSubject

`func (o *AgentTicketListV2ResponseDataDataInner) HasTicketSubject() bool`

HasTicketSubject returns a boolean if a field has been set.

### GetTicketStatus

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketStatus() string`

GetTicketStatus returns the TicketStatus field if non-nil, zero value otherwise.

### GetTicketStatusOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketStatusOk() (*string, bool)`

GetTicketStatusOk returns a tuple with the TicketStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatus

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketStatus(v string)`

SetTicketStatus sets TicketStatus field to given value.

### HasTicketStatus

`func (o *AgentTicketListV2ResponseDataDataInner) HasTicketStatus() bool`

HasTicketStatus returns a boolean if a field has been set.

### GetTemplateName

`func (o *AgentTicketListV2ResponseDataDataInner) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *AgentTicketListV2ResponseDataDataInner) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *AgentTicketListV2ResponseDataDataInner) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetCreatorName

`func (o *AgentTicketListV2ResponseDataDataInner) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *AgentTicketListV2ResponseDataDataInner) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *AgentTicketListV2ResponseDataDataInner) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetCreateTime

`func (o *AgentTicketListV2ResponseDataDataInner) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AgentTicketListV2ResponseDataDataInner) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AgentTicketListV2ResponseDataDataInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AgentTicketListV2ResponseDataDataInner) GetUpdateTime() string`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetUpdateTimeOk() (*string, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AgentTicketListV2ResponseDataDataInner) SetUpdateTime(v string)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AgentTicketListV2ResponseDataDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetCurrentProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) GetCurrentProcessPersons() []string`

GetCurrentProcessPersons returns the CurrentProcessPersons field if non-nil, zero value otherwise.

### GetCurrentProcessPersonsOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetCurrentProcessPersonsOk() (*[]string, bool)`

GetCurrentProcessPersonsOk returns a tuple with the CurrentProcessPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) SetCurrentProcessPersons(v []string)`

SetCurrentProcessPersons sets CurrentProcessPersons field to given value.

### HasCurrentProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) HasCurrentProcessPersons() bool`

HasCurrentProcessPersons returns a boolean if a field has been set.

### GetTaskName

`func (o *AgentTicketListV2ResponseDataDataInner) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *AgentTicketListV2ResponseDataDataInner) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *AgentTicketListV2ResponseDataDataInner) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetAdvId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvId() int64`

GetAdvId returns the AdvId field if non-nil, zero value otherwise.

### GetAdvIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvIdOk() (*int64, bool)`

GetAdvIdOk returns a tuple with the AdvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAdvId(v int64)`

SetAdvId sets AdvId field to given value.

### HasAdvId

`func (o *AgentTicketListV2ResponseDataDataInner) HasAdvId() bool`

HasAdvId returns a boolean if a field has been set.

### GetAdvName

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvName() string`

GetAdvName returns the AdvName field if non-nil, zero value otherwise.

### GetAdvNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvNameOk() (*string, bool)`

GetAdvNameOk returns a tuple with the AdvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvName

`func (o *AgentTicketListV2ResponseDataDataInner) SetAdvName(v string)`

SetAdvName sets AdvName field to given value.

### HasAdvName

`func (o *AgentTicketListV2ResponseDataDataInner) HasAdvName() bool`

HasAdvName returns a boolean if a field has been set.

### GetAgentName

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetAgentId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentId() int64`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentIdOk() (*int64, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentId(v int64)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentCompanyName

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyName() string`

GetAgentCompanyName returns the AgentCompanyName field if non-nil, zero value otherwise.

### GetAgentCompanyNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyNameOk() (*string, bool)`

GetAgentCompanyNameOk returns a tuple with the AgentCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCompanyName

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentCompanyName(v string)`

SetAgentCompanyName sets AgentCompanyName field to given value.

### HasAgentCompanyName

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentCompanyName() bool`

HasAgentCompanyName returns a boolean if a field has been set.

### GetAgentCompanyId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyId() int64`

GetAgentCompanyId returns the AgentCompanyId field if non-nil, zero value otherwise.

### GetAgentCompanyIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyIdOk() (*int64, bool)`

GetAgentCompanyIdOk returns a tuple with the AgentCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCompanyId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentCompanyId(v int64)`

SetAgentCompanyId sets AgentCompanyId field to given value.

### HasAgentCompanyId

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentCompanyId() bool`

HasAgentCompanyId returns a boolean if a field has been set.

### GetSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesName() string`

GetSalesName returns the SalesName field if non-nil, zero value otherwise.

### GetSalesNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesNameOk() (*string, bool)`

GetSalesNameOk returns a tuple with the SalesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) SetSalesName(v string)`

SetSalesName sets SalesName field to given value.

### HasSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) HasSalesName() bool`

HasSalesName returns a boolean if a field has been set.

### GetSalesDepId

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepId() int64`

GetSalesDepId returns the SalesDepId field if non-nil, zero value otherwise.

### GetSalesDepIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepIdOk() (*int64, bool)`

GetSalesDepIdOk returns a tuple with the SalesDepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDepId

`func (o *AgentTicketListV2ResponseDataDataInner) SetSalesDepId(v int64)`

SetSalesDepId sets SalesDepId field to given value.

### HasSalesDepId

`func (o *AgentTicketListV2ResponseDataDataInner) HasSalesDepId() bool`

HasSalesDepId returns a boolean if a field has been set.

### GetSalesDepName

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepName() string`

GetSalesDepName returns the SalesDepName field if non-nil, zero value otherwise.

### GetSalesDepNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepNameOk() (*string, bool)`

GetSalesDepNameOk returns a tuple with the SalesDepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDepName

`func (o *AgentTicketListV2ResponseDataDataInner) SetSalesDepName(v string)`

SetSalesDepName sets SalesDepName field to given value.

### HasSalesDepName

`func (o *AgentTicketListV2ResponseDataDataInner) HasSalesDepName() bool`

HasSalesDepName returns a boolean if a field has been set.

### GetProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) GetProcessPersons() []string`

GetProcessPersons returns the ProcessPersons field if non-nil, zero value otherwise.

### GetProcessPersonsOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetProcessPersonsOk() (*[]string, bool)`

GetProcessPersonsOk returns a tuple with the ProcessPersons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) SetProcessPersons(v []string)`

SetProcessPersons sets ProcessPersons field to given value.

### HasProcessPersons

`func (o *AgentTicketListV2ResponseDataDataInner) HasProcessPersons() bool`

HasProcessPersons returns a boolean if a field has been set.

### GetTicketStatusName

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketStatusName() string`

GetTicketStatusName returns the TicketStatusName field if non-nil, zero value otherwise.

### GetTicketStatusNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketStatusNameOk() (*string, bool)`

GetTicketStatusNameOk returns a tuple with the TicketStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatusName

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketStatusName(v string)`

SetTicketStatusName sets TicketStatusName field to given value.

### HasTicketStatusName

`func (o *AgentTicketListV2ResponseDataDataInner) HasTicketStatusName() bool`

HasTicketStatusName returns a boolean if a field has been set.

### GetTicketApp

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketApp() string`

GetTicketApp returns the TicketApp field if non-nil, zero value otherwise.

### GetTicketAppOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketAppOk() (*string, bool)`

GetTicketAppOk returns a tuple with the TicketApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketApp

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketApp(v string)`

SetTicketApp sets TicketApp field to given value.

### HasTicketApp

`func (o *AgentTicketListV2ResponseDataDataInner) HasTicketApp() bool`

HasTicketApp returns a boolean if a field has been set.

### GetAgentSalesId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesId() int64`

GetAgentSalesId returns the AgentSalesId field if non-nil, zero value otherwise.

### GetAgentSalesIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesIdOk() (*int64, bool)`

GetAgentSalesIdOk returns a tuple with the AgentSalesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSalesId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentSalesId(v int64)`

SetAgentSalesId sets AgentSalesId field to given value.

### HasAgentSalesId

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentSalesId() bool`

HasAgentSalesId returns a boolean if a field has been set.

### GetAgentSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesName() string`

GetAgentSalesName returns the AgentSalesName field if non-nil, zero value otherwise.

### GetAgentSalesNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesNameOk() (*string, bool)`

GetAgentSalesNameOk returns a tuple with the AgentSalesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentSalesName(v string)`

SetAgentSalesName sets AgentSalesName field to given value.

### HasAgentSalesName

`func (o *AgentTicketListV2ResponseDataDataInner) HasAgentSalesName() bool`

HasAgentSalesName returns a boolean if a field has been set.

### GetOptimizerId

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerId() int64`

GetOptimizerId returns the OptimizerId field if non-nil, zero value otherwise.

### GetOptimizerIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerIdOk() (*int64, bool)`

GetOptimizerIdOk returns a tuple with the OptimizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizerId

`func (o *AgentTicketListV2ResponseDataDataInner) SetOptimizerId(v int64)`

SetOptimizerId sets OptimizerId field to given value.

### HasOptimizerId

`func (o *AgentTicketListV2ResponseDataDataInner) HasOptimizerId() bool`

HasOptimizerId returns a boolean if a field has been set.

### GetOptimizerName

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerName() string`

GetOptimizerName returns the OptimizerName field if non-nil, zero value otherwise.

### GetOptimizerNameOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerNameOk() (*string, bool)`

GetOptimizerNameOk returns a tuple with the OptimizerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizerName

`func (o *AgentTicketListV2ResponseDataDataInner) SetOptimizerName(v string)`

SetOptimizerName sets OptimizerName field to given value.

### HasOptimizerName

`func (o *AgentTicketListV2ResponseDataDataInner) HasOptimizerName() bool`

HasOptimizerName returns a boolean if a field has been set.

### GetAthenaOrderId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAthenaOrderId() string`

GetAthenaOrderId returns the AthenaOrderId field if non-nil, zero value otherwise.

### GetAthenaOrderIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAthenaOrderIdOk() (*string, bool)`

GetAthenaOrderIdOk returns a tuple with the AthenaOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAthenaOrderId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAthenaOrderId(v string)`

SetAthenaOrderId sets AthenaOrderId field to given value.

### HasAthenaOrderId

`func (o *AgentTicketListV2ResponseDataDataInner) HasAthenaOrderId() bool`

HasAthenaOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


