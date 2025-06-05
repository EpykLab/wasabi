# \OverviewAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersOverviewControllerGetOverviewAgents**](OverviewAPI.md#ApiControllersOverviewControllerGetOverviewAgents) | **Get** /overview/agents | Get agents overview



## ApiControllersOverviewControllerGetOverviewAgents

> ApiControllersOverviewControllerGetOverviewAgents200Response ApiControllersOverviewControllerGetOverviewAgents(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get agents overview



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
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverviewAPI.ApiControllersOverviewControllerGetOverviewAgents(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewAPI.ApiControllersOverviewControllerGetOverviewAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersOverviewControllerGetOverviewAgents`: ApiControllersOverviewControllerGetOverviewAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `OverviewAPI.ApiControllersOverviewControllerGetOverviewAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersOverviewControllerGetOverviewAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersOverviewControllerGetOverviewAgents200Response**](ApiControllersOverviewControllerGetOverviewAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
