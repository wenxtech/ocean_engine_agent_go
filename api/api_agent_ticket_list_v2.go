/*
巨量方舟

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
    . "github.com/wenxtech/ocean_engine_agent_go/models"
)


// AgentTicketListV2APIService AgentTicketListV2API service
type AgentTicketListV2APIService service

type ApiAgentTicketListV2PostRequest struct {
	ctx context.Context
	ApiService *AgentTicketListV2APIService
	page *int64
	ticketApp *string
	size *int64
	ticketTabType *string
	createTimeStart *string
	createTimeEnd *string
	ticketStatus *string
}

func (r ApiAgentTicketListV2PostRequest) Page(page int64) ApiAgentTicketListV2PostRequest {
	r.page = &page
	return r
}

func (r ApiAgentTicketListV2PostRequest) TicketApp(ticketApp string) ApiAgentTicketListV2PostRequest {
	r.ticketApp = &ticketApp
	return r
}

func (r ApiAgentTicketListV2PostRequest) Size(size int64) ApiAgentTicketListV2PostRequest {
	r.size = &size
	return r
}

func (r ApiAgentTicketListV2PostRequest) TicketTabType(ticketTabType string) ApiAgentTicketListV2PostRequest {
	r.ticketTabType = &ticketTabType
	return r
}

func (r ApiAgentTicketListV2PostRequest) CreateTimeStart(createTimeStart string) ApiAgentTicketListV2PostRequest {
	r.createTimeStart = &createTimeStart
	return r
}

func (r ApiAgentTicketListV2PostRequest) CreateTimeEnd(createTimeEnd string) ApiAgentTicketListV2PostRequest {
	r.createTimeEnd = &createTimeEnd
	return r
}

func (r ApiAgentTicketListV2PostRequest) TicketStatus(ticketStatus string) ApiAgentTicketListV2PostRequest {
	r.ticketStatus = &ticketStatus
	return r
}

func (r ApiAgentTicketListV2PostRequest) Execute() (*AgentTicketListV2Response, *http.Response, error) {
	return r.ApiService.AgentTicketListV2PostExecute(r)
}

func (r ApiAgentTicketListV2PostRequest) WithLog(enable bool) ApiAgentTicketListV2PostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, ContextEnableLog, true)
	}
	return r
}

/*
AgentTicketListV2Post 工单列表



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAgentTicketListV2PostRequest
*/
func (a *AgentTicketListV2APIService) Post(ctx context.Context) ApiAgentTicketListV2PostRequest {
	return ApiAgentTicketListV2PostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AgentTicketListV2Response
func (a *AgentTicketListV2APIService) AgentTicketListV2PostExecute(r ApiAgentTicketListV2PostRequest) (*AgentTicketListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []FormFile
		localVarReturnValue  *AgentTicketListV2Response
	)
	
	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath, err := a.client.Cfg.ServerURLWithContext(r.ctx, "AgentTicketListV2APIService.AgentTicketListV2Post")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agent/ticket/list/V2"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if r.ticketApp == nil {
		return localVarReturnValue, nil, reportError("ticketApp is required and must be specified")
	}
	if r.size == nil {
		return localVarReturnValue, nil, reportError("size is required and must be specified")
	}
	if r.ticketTabType == nil {
		return localVarReturnValue, nil, reportError("ticketTabType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "page", r.page, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "ticketApp", r.ticketApp, "")
	if r.createTimeStart != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "createTimeStart", r.createTimeStart, "")
	}
	if r.createTimeEnd != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "createTimeEnd", r.createTimeEnd, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "size", r.size, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "ticketTabType", r.ticketTabType, "")
	if r.ticketStatus != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ticketStatus", r.ticketStatus, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
