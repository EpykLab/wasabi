# \APIInfoAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersDefaultControllerDefaultInfo**](APIInfoAPI.md#ApiControllersDefaultControllerDefaultInfo) | **Get** / | Get API info



## ApiControllersDefaultControllerDefaultInfo

> ApiControllersDefaultControllerDefaultInfo200Response ApiControllersDefaultControllerDefaultInfo(ctx).Pretty(pretty).Execute()

Get API info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIInfoAPI.ApiControllersDefaultControllerDefaultInfo(context.Background()).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIInfoAPI.ApiControllersDefaultControllerDefaultInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDefaultControllerDefaultInfo`: ApiControllersDefaultControllerDefaultInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `APIInfoAPI.ApiControllersDefaultControllerDefaultInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDefaultControllerDefaultInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]

### Return type

[**ApiControllersDefaultControllerDefaultInfo200Response**](ApiControllersDefaultControllerDefaultInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
