/*
巨量方舟

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// TicketListV2ResponseDataDataInner struct for TicketListV2ResponseDataDataInner
type TicketListV2ResponseDataDataInner struct {
	TicketId *int64 `json:"ticketId,omitempty"`
	TicketSubject *string `json:"ticketSubject,omitempty"`
	TicketStatus *string `json:"ticketStatus,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	CreatorName *string `json:"creatorName,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty"`
	TaskName *string `json:"taskName,omitempty"`
	AdvId *int64 `json:"advId,omitempty"`
	AdvName *string `json:"advName,omitempty"`
	AgentName *string `json:"agentName,omitempty"`
	AgentId *int64 `json:"agentId,omitempty"`
	AgentCompanyName *string `json:"agentCompanyName,omitempty"`
	AgentCompanyId *int64 `json:"agentCompanyId,omitempty"`
	SalesName *string `json:"salesName,omitempty"`
	SalesDepId *int64 `json:"salesDepId,omitempty"`
	SalesDepName *string `json:"salesDepName,omitempty"`
	TicketStatusName *string `json:"ticketStatusName,omitempty"`
	TicketApp *string `json:"ticketApp,omitempty"`
	AgentSalesId *int64 `json:"agentSalesId,omitempty"`
	AthenaOrderId *string `json:"athenaOrderId,omitempty"`
}


