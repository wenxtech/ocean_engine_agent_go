# \AgentTicketListV2API

All URIs are relative to *https://agent.oceanengine.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentTicketListV2Post**](AgentTicketListV2API.md#AgentTicketListV2Post) | **Post** /agent/ticket/list/V2 | 工单列表



## AgentTicketListV2Post

> AgentTicketListV2Response AgentTicketListV2Post(ctx).Page(page).TicketApp(ticketApp).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Size(size).TicketTabType(ticketTabType).Execute()

工单列表



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wenxtech/ocean_engine_agent_go"
)

func main() {
	page := "page_example" // string |  (optional)
	ticketApp := "ticketApp_example" // string |  (optional)
	createTimeStart := "createTimeStart_example" // string |  (optional)
	createTimeEnd := "createTimeEnd_example" // string |  (optional)
	size := "size_example" // string |  (optional)
	ticketTabType := "ticketTabType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentTicketListV2API.AgentTicketListV2Post(context.Background()).Page(page).TicketApp(ticketApp).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Size(size).TicketTabType(ticketTabType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentTicketListV2API.AgentTicketListV2Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentTicketListV2Post`: AgentTicketListV2Response
	fmt.Fprintf(os.Stdout, "Response from `AgentTicketListV2API.AgentTicketListV2Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentTicketListV2PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** |  | 
 **ticketApp** | **string** |  | 
 **createTimeStart** | **string** |  | 
 **createTimeEnd** | **string** |  | 
 **size** | **string** |  | 
 **ticketTabType** | **string** |  | 

### Return type

[**AgentTicketListV2Response**](AgentTicketListV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

