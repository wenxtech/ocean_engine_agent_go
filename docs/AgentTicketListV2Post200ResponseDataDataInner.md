# AgentTicketListV2ResponseDataDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TicketId** | **int32** |  | 
**TicketSubject** | **NullableString** |  | 
**TicketStatus** | **string** |  | 
**TemplateName** | **string** |  | 
**CreatorName** | **string** |  | 
**CreateTime** | **string** |  | 
**UpdateTime** | **string** |  | 
**CurrentProcessPersons** | **[]string** |  | 
**TaskName** | **string** |  | 
**AdvId** | **NullableInt32** |  | 
**AdvName** | **NullableString** |  | 
**AgentName** | **string** |  | 
**AgentId** | **int32** |  | 
**AgentCompanyName** | **string** |  | 
**AgentCompanyId** | **int32** |  | 
**SalesName** | **string** |  | 
**SalesDepId** | **int32** |  | 
**SalesDepName** | **string** |  | 
**ProcessPersons** | **[]string** |  | 
**TicketStatusName** | **string** |  | 
**TicketApp** | **string** |  | 
**AgentSalesId** | **int32** |  | 
**AgentSalesName** | **string** |  | 
**OptimizerId** | **NullableInt32** |  | 
**OptimizerName** | **NullableString** |  | 
**AthenaOrderId** | **string** |  | 

## Methods

### NewAgentTicketListV2ResponseDataDataInner

`func NewAgentTicketListV2ResponseDataDataInner(ticketId int32, ticketSubject NullableString, ticketStatus string, templateName string, creatorName string, createTime string, updateTime string, currentProcessPersons []string, taskName string, advId NullableInt32, advName NullableString, agentName string, agentId int32, agentCompanyName string, agentCompanyId int32, salesName string, salesDepId int32, salesDepName string, processPersons []string, ticketStatusName string, ticketApp string, agentSalesId int32, agentSalesName string, optimizerId NullableInt32, optimizerName NullableString, athenaOrderId string, ) *AgentTicketListV2ResponseDataDataInner`

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

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.


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


### SetTicketSubjectNil

`func (o *AgentTicketListV2ResponseDataDataInner) SetTicketSubjectNil(b bool)`

 SetTicketSubjectNil sets the value for TicketSubject to be an explicit nil

### UnsetTicketSubject
`func (o *AgentTicketListV2ResponseDataDataInner) UnsetTicketSubject()`

UnsetTicketSubject ensures that no value is present for TicketSubject, not even an explicit nil
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


### GetAdvId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvId() int32`

GetAdvId returns the AdvId field if non-nil, zero value otherwise.

### GetAdvIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAdvIdOk() (*int32, bool)`

GetAdvIdOk returns a tuple with the AdvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAdvId(v int32)`

SetAdvId sets AdvId field to given value.


### SetAdvIdNil

`func (o *AgentTicketListV2ResponseDataDataInner) SetAdvIdNil(b bool)`

 SetAdvIdNil sets the value for AdvId to be an explicit nil

### UnsetAdvId
`func (o *AgentTicketListV2ResponseDataDataInner) UnsetAdvId()`

UnsetAdvId ensures that no value is present for AdvId, not even an explicit nil
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


### SetAdvNameNil

`func (o *AgentTicketListV2ResponseDataDataInner) SetAdvNameNil(b bool)`

 SetAdvNameNil sets the value for AdvName to be an explicit nil

### UnsetAdvName
`func (o *AgentTicketListV2ResponseDataDataInner) UnsetAdvName()`

UnsetAdvName ensures that no value is present for AdvName, not even an explicit nil
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


### GetAgentId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.


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


### GetAgentCompanyId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyId() int32`

GetAgentCompanyId returns the AgentCompanyId field if non-nil, zero value otherwise.

### GetAgentCompanyIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentCompanyIdOk() (*int32, bool)`

GetAgentCompanyIdOk returns a tuple with the AgentCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentCompanyId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentCompanyId(v int32)`

SetAgentCompanyId sets AgentCompanyId field to given value.


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


### GetSalesDepId

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepId() int32`

GetSalesDepId returns the SalesDepId field if non-nil, zero value otherwise.

### GetSalesDepIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetSalesDepIdOk() (*int32, bool)`

GetSalesDepIdOk returns a tuple with the SalesDepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDepId

`func (o *AgentTicketListV2ResponseDataDataInner) SetSalesDepId(v int32)`

SetSalesDepId sets SalesDepId field to given value.


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


### GetAgentSalesId

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesId() int32`

GetAgentSalesId returns the AgentSalesId field if non-nil, zero value otherwise.

### GetAgentSalesIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetAgentSalesIdOk() (*int32, bool)`

GetAgentSalesIdOk returns a tuple with the AgentSalesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSalesId

`func (o *AgentTicketListV2ResponseDataDataInner) SetAgentSalesId(v int32)`

SetAgentSalesId sets AgentSalesId field to given value.


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


### GetOptimizerId

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerId() int32`

GetOptimizerId returns the OptimizerId field if non-nil, zero value otherwise.

### GetOptimizerIdOk

`func (o *AgentTicketListV2ResponseDataDataInner) GetOptimizerIdOk() (*int32, bool)`

GetOptimizerIdOk returns a tuple with the OptimizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizerId

`func (o *AgentTicketListV2ResponseDataDataInner) SetOptimizerId(v int32)`

SetOptimizerId sets OptimizerId field to given value.


### SetOptimizerIdNil

`func (o *AgentTicketListV2ResponseDataDataInner) SetOptimizerIdNil(b bool)`

 SetOptimizerIdNil sets the value for OptimizerId to be an explicit nil

### UnsetOptimizerId
`func (o *AgentTicketListV2ResponseDataDataInner) UnsetOptimizerId()`

UnsetOptimizerId ensures that no value is present for OptimizerId, not even an explicit nil
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


### SetOptimizerNameNil

`func (o *AgentTicketListV2ResponseDataDataInner) SetOptimizerNameNil(b bool)`

 SetOptimizerNameNil sets the value for OptimizerName to be an explicit nil

### UnsetOptimizerName
`func (o *AgentTicketListV2ResponseDataDataInner) UnsetOptimizerName()`

UnsetOptimizerName ensures that no value is present for OptimizerName, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


