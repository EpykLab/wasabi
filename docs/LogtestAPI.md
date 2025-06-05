# \LogtestAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersLogtestControllerEndLogtestSession**](LogtestAPI.md#ApiControllersLogtestControllerEndLogtestSession) | **Delete** /logtest/sessions/{token} | End session
[**ApiControllersLogtestControllerRunLogtestTool**](LogtestAPI.md#ApiControllersLogtestControllerRunLogtestTool) | **Put** /logtest | Run logtest



## ApiControllersLogtestControllerEndLogtestSession

> ApiResponse ApiControllersLogtestControllerEndLogtestSession(ctx, token).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

End session



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
	token := "token_example" // string | Token of the logtest saved session
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogtestAPI.ApiControllersLogtestControllerEndLogtestSession(context.Background(), token).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogtestAPI.ApiControllersLogtestControllerEndLogtestSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersLogtestControllerEndLogtestSession`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `LogtestAPI.ApiControllersLogtestControllerEndLogtestSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the logtest saved session |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersLogtestControllerEndLogtestSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersLogtestControllerRunLogtestTool

> ApiResponse ApiControllersLogtestControllerRunLogtestTool(ctx).LogtestRequest(logtestRequest).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Run logtest



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
	logtestRequest := *openapiclient.NewLogtestRequest("LogFormat_example", "Location_example", "Event_example") // LogtestRequest | Run logtest with the parameters below
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LogtestAPI.ApiControllersLogtestControllerRunLogtestTool(context.Background()).LogtestRequest(logtestRequest).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogtestAPI.ApiControllersLogtestControllerRunLogtestTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersLogtestControllerRunLogtestTool`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `LogtestAPI.ApiControllersLogtestControllerRunLogtestTool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersLogtestControllerRunLogtestToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logtestRequest** | [**LogtestRequest**](LogtestRequest.md) | Run logtest with the parameters below |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
