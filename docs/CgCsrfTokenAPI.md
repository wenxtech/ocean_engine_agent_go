# \CgCsrfTokenAPI

All URIs are relative to *https://agent.oceanengine.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CgCsrfTokenGet**](CgCsrfTokenAPI.md#CgCsrfTokenGet) | **Get** /cg-csrf-token | 获取CgCsrfToken



## CgCsrfTokenGet

> CgCsrfTokenResponse CgCsrfTokenGet(ctx).Execute()

获取CgCsrfToken



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CgCsrfTokenAPI.CgCsrfTokenGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CgCsrfTokenAPI.CgCsrfTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CgCsrfTokenGet`: CgCsrfTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `CgCsrfTokenAPI.CgCsrfTokenGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCgCsrfTokenGetRequest struct via the builder pattern


### Return type

[**CgCsrfTokenResponse**](CgCsrfTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

