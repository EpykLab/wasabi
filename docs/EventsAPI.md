# \EventsAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersEventControllerForwardEvent**](EventsAPI.md#ApiControllersEventControllerForwardEvent) | **Post** /events | Ingest events



## ApiControllersEventControllerForwardEvent

> ApiResponse ApiControllersEventControllerForwardEvent(ctx).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersEventControllerForwardEventRequest(apiControllersEventControllerForwardEventRequest).Execute()

Ingest events



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
	apiControllersEventControllerForwardEventRequest := *openapiclient.NewApiControllersEventControllerForwardEventRequest([]string{"Events_example"}) // ApiControllersEventControllerForwardEventRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.ApiControllersEventControllerForwardEvent(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersEventControllerForwardEventRequest(apiControllersEventControllerForwardEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.ApiControllersEventControllerForwardEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersEventControllerForwardEvent`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.ApiControllersEventControllerForwardEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersEventControllerForwardEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **apiControllersEventControllerForwardEventRequest** | [**ApiControllersEventControllerForwardEventRequest**](ApiControllersEventControllerForwardEventRequest.md) |  |

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
