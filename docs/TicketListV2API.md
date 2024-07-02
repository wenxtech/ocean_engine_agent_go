# \TicketListV2API

All URIs are relative to *https://agent.oceanengine.com/agent*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketListV2Post**](TicketListV2API.md#TicketListV2Post) | **Post** /ticket/list/V2 | 工单列表



## TicketListV2Post

> TicketListV2Response TicketListV2Post(ctx).Page(page).TicketApp(ticketApp).Size(size).TicketTabType(ticketTabType).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).TicketStatus(ticketStatus).Execute()

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
	page := int64(789) // int64 | 
	ticketApp := "ticketApp_example" // string | 
	size := int64(789) // int64 | 
	ticketTabType := "ticketTabType_example" // string | 
	createTimeStart := "createTimeStart_example" // string |  (optional)
	createTimeEnd := "createTimeEnd_example" // string |  (optional)
	ticketStatus := "ticketStatus_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketListV2API.TicketListV2Post(context.Background()).Page(page).TicketApp(ticketApp).Size(size).TicketTabType(ticketTabType).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).TicketStatus(ticketStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketListV2API.TicketListV2Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TicketListV2Post`: TicketListV2Response
	fmt.Fprintf(os.Stdout, "Response from `TicketListV2API.TicketListV2Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketListV2PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **ticketApp** | **string** |  | 
 **size** | **int64** |  | 
 **ticketTabType** | **string** |  | 
 **createTimeStart** | **string** |  | 
 **createTimeEnd** | **string** |  | 
 **ticketStatus** | **string** |  | 

### Return type

[**TicketListV2Response**](TicketListV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

