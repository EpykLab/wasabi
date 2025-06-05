# \ActiveResponseAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersActiveResponseControllerRunCommand**](ActiveResponseAPI.md#ApiControllersActiveResponseControllerRunCommand) | **Put** /active-response | Run command



## ApiControllersActiveResponseControllerRunCommand

> ApiResponse ApiControllersActiveResponseControllerRunCommand(ctx).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).ActiveResponseBody(activeResponseBody).Execute()

Run command



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	activeResponseBody := *openapiclient.NewActiveResponseBody("Command_example") // ActiveResponseBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActiveResponseAPI.ApiControllersActiveResponseControllerRunCommand(context.Background()).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).ActiveResponseBody(activeResponseBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActiveResponseAPI.ApiControllersActiveResponseControllerRunCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersActiveResponseControllerRunCommand`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ActiveResponseAPI.ApiControllersActiveResponseControllerRunCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersActiveResponseControllerRunCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **activeResponseBody** | [**ActiveResponseBody**](ActiveResponseBody.md) |  |

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
