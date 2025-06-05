# \RootcheckAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersRootcheckControllerDeleteRootcheck**](RootcheckAPI.md#ApiControllersRootcheckControllerDeleteRootcheck) | **Delete** /rootcheck/{agent_id} | Clear results
[**ApiControllersRootcheckControllerGetLastScanAgent**](RootcheckAPI.md#ApiControllersRootcheckControllerGetLastScanAgent) | **Get** /rootcheck/{agent_id}/last_scan | Get last scan datetime
[**ApiControllersRootcheckControllerGetRootcheckAgent**](RootcheckAPI.md#ApiControllersRootcheckControllerGetRootcheckAgent) | **Get** /rootcheck/{agent_id} | Get results
[**ApiControllersRootcheckControllerPutRootcheck**](RootcheckAPI.md#ApiControllersRootcheckControllerPutRootcheck) | **Put** /rootcheck | Run scan



## ApiControllersRootcheckControllerDeleteRootcheck

> ApiControllersMitreControllerGetGroups200Response ApiControllersRootcheckControllerDeleteRootcheck(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Clear results



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
	agentId := "agentId_example" // string | Agent ID. All possible values from 000 onwards
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RootcheckAPI.ApiControllersRootcheckControllerDeleteRootcheck(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RootcheckAPI.ApiControllersRootcheckControllerDeleteRootcheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRootcheckControllerDeleteRootcheck`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RootcheckAPI.ApiControllersRootcheckControllerDeleteRootcheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRootcheckControllerDeleteRootcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersRootcheckControllerGetLastScanAgent

> ApiControllersMitreControllerGetGroups200Response ApiControllersRootcheckControllerGetLastScanAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get last scan datetime



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
	agentId := "agentId_example" // string | Agent ID. All possible values from 000 onwards
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RootcheckAPI.ApiControllersRootcheckControllerGetLastScanAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RootcheckAPI.ApiControllersRootcheckControllerGetLastScanAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRootcheckControllerGetLastScanAgent`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RootcheckAPI.ApiControllersRootcheckControllerGetLastScanAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRootcheckControllerGetLastScanAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersRootcheckControllerGetRootcheckAgent

> ApiControllersMitreControllerGetGroups200Response ApiControllersRootcheckControllerGetRootcheckAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Status(status).PciDss(pciDss).Cis(cis).Execute()

Get results



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
	agentId := "agentId_example" // string | Agent ID. All possible values from 000 onwards
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)
	status := "status_example" // string | Filter by status (optional)
	pciDss := "pciDss_example" // string | Filter by PCI_DSS requirement name (optional)
	cis := "cis_example" // string | Filter by CIS requirement (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RootcheckAPI.ApiControllersRootcheckControllerGetRootcheckAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Status(status).PciDss(pciDss).Cis(cis).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RootcheckAPI.ApiControllersRootcheckControllerGetRootcheckAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRootcheckControllerGetRootcheckAgent`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RootcheckAPI.ApiControllersRootcheckControllerGetRootcheckAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRootcheckControllerGetRootcheckAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]
 **status** | **string** | Filter by status |
 **pciDss** | **string** | Filter by PCI_DSS requirement name |
 **cis** | **string** | Filter by CIS requirement |

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersRootcheckControllerPutRootcheck

> ApiControllersMitreControllerGetGroups200Response ApiControllersRootcheckControllerPutRootcheck(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()

Run scan



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RootcheckAPI.ApiControllersRootcheckControllerPutRootcheck(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RootcheckAPI.ApiControllersRootcheckControllerPutRootcheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRootcheckControllerPutRootcheck`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RootcheckAPI.ApiControllersRootcheckControllerPutRootcheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRootcheckControllerPutRootcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
