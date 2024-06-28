/*
巨量方舟

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
    "fmt"
)
// TicketStatus the model 'TicketStatus'
type TicketStatus string

// List of ticketStatus
const (
	UN_HANDLE_TicketStatus = "UN_HANDLE"
	HANDLING_TicketStatus = "HANDLING"
	HANDLED_TicketStatus = "HANDLED"
	REJECT_TicketStatus = "REJECT"
	WITHDRAW_TicketStatus = "WITHDRAW"
)

// All allowed values of TicketStatus enum
var AllowedTicketStatusEnumValues = []TicketStatus{
	"UN_HANDLE",
	"HANDLING",
	"HANDLED",
	"REJECT",
	"WITHDRAW",
}

// NewTicketStatusFromValue returns a pointer to a valid TicketStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTicketStatusFromValue(v string) (*TicketStatus, error) {
	ev := TicketStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TicketStatus: valid values are %v", v, AllowedTicketStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TicketStatus) IsValid() bool {
	for _, existing := range AllowedTicketStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ticketStatus value
func (v TicketStatus) Ptr() *TicketStatus {
	return &v
}

