# \SyscheckAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersSyscheckControllerDeleteSyscheckAgent**](SyscheckAPI.md#ApiControllersSyscheckControllerDeleteSyscheckAgent) | **Delete** /syscheck/{agent_id} | Clear results
[**ApiControllersSyscheckControllerGetLastScanAgent**](SyscheckAPI.md#ApiControllersSyscheckControllerGetLastScanAgent) | **Get** /syscheck/{agent_id}/last_scan | Get last scan datetime
[**ApiControllersSyscheckControllerGetSyscheckAgent**](SyscheckAPI.md#ApiControllersSyscheckControllerGetSyscheckAgent) | **Get** /syscheck/{agent_id} | Get results
[**ApiControllersSyscheckControllerPutSyscheck**](SyscheckAPI.md#ApiControllersSyscheckControllerPutSyscheck) | **Put** /syscheck | Run scan



## ApiControllersSyscheckControllerDeleteSyscheckAgent

> ApiControllersMitreControllerGetGroups200Response ApiControllersSyscheckControllerDeleteSyscheckAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

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
	resp, r, err := apiClient.SyscheckAPI.ApiControllersSyscheckControllerDeleteSyscheckAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscheckAPI.ApiControllersSyscheckControllerDeleteSyscheckAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscheckControllerDeleteSyscheckAgent`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscheckAPI.ApiControllersSyscheckControllerDeleteSyscheckAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscheckControllerDeleteSyscheckAgentRequest struct via the builder pattern


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


## ApiControllersSyscheckControllerGetLastScanAgent

> ApiControllersSyscheckControllerGetLastScanAgent200Response ApiControllersSyscheckControllerGetLastScanAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

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
	resp, r, err := apiClient.SyscheckAPI.ApiControllersSyscheckControllerGetLastScanAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscheckAPI.ApiControllersSyscheckControllerGetLastScanAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscheckControllerGetLastScanAgent`: ApiControllersSyscheckControllerGetLastScanAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscheckAPI.ApiControllersSyscheckControllerGetLastScanAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscheckControllerGetLastScanAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSyscheckControllerGetLastScanAgent200Response**](ApiControllersSyscheckControllerGetLastScanAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSyscheckControllerGetSyscheckAgent

> ApiControllersSyscheckControllerGetSyscheckAgent200Response ApiControllersSyscheckControllerGetSyscheckAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).File(file).Arch(arch).ValueName(valueName).ValueType(valueType).Type_(type_).Summary(summary).Md5(md5).Sha1(sha1).Sha256(sha256).Hash(hash).Distinct(distinct).Q(q).Execute()

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
	file := "file_example" // string | Filter by full path (optional)
	arch := "arch_example" // string | Filter by architecture (optional)
	valueName := "valueName_example" // string | Filter by value name (optional)
	valueType := "valueType_example" // string | Filter by value type (optional)
	type_ := "type__example" // string | Filter by file type. Registry_key and registry_value types are only available in Windows agents (optional)
	summary := true // bool | Return a summary grouping by filename (optional) (default to false)
	md5 := "md5_example" // string | Filter files with the specified MD5 checksum (optional)
	sha1 := "sha1_example" // string | Filter files with the specified SHA1 checksum (optional)
	sha256 := "sha256_example" // string | Filter files with the specified SHA256 checksum (optional)
	hash := "hash_example" // string | Filter files with the specified hash (md5, sha256 or sha1) (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscheckAPI.ApiControllersSyscheckControllerGetSyscheckAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).File(file).Arch(arch).ValueName(valueName).ValueType(valueType).Type_(type_).Summary(summary).Md5(md5).Sha1(sha1).Sha256(sha256).Hash(hash).Distinct(distinct).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscheckAPI.ApiControllersSyscheckControllerGetSyscheckAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscheckControllerGetSyscheckAgent`: ApiControllersSyscheckControllerGetSyscheckAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscheckAPI.ApiControllersSyscheckControllerGetSyscheckAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscheckControllerGetSyscheckAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **file** | **string** | Filter by full path |
 **arch** | **string** | Filter by architecture |
 **valueName** | **string** | Filter by value name |
 **valueType** | **string** | Filter by value type |
 **type_** | **string** | Filter by file type. Registry_key and registry_value types are only available in Windows agents |
 **summary** | **bool** | Return a summary grouping by filename | [default to false]
 **md5** | **string** | Filter files with the specified MD5 checksum |
 **sha1** | **string** | Filter files with the specified SHA1 checksum |
 **sha256** | **string** | Filter files with the specified SHA256 checksum |
 **hash** | **string** | Filter files with the specified hash (md5, sha256 or sha1) |
 **distinct** | **bool** | Look for distinct values. | [default to false]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

### Return type

[**ApiControllersSyscheckControllerGetSyscheckAgent200Response**](ApiControllersSyscheckControllerGetSyscheckAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSyscheckControllerPutSyscheck

> ApiControllersAgentControllerRestartAgentsByNode200Response ApiControllersSyscheckControllerPutSyscheck(ctx).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscheckAPI.ApiControllersSyscheckControllerPutSyscheck(context.Background()).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscheckAPI.ApiControllersSyscheckControllerPutSyscheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscheckControllerPutSyscheck`: ApiControllersAgentControllerRestartAgentsByNode200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscheckAPI.ApiControllersSyscheckControllerPutSyscheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscheckControllerPutSyscheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerRestartAgentsByNode200Response**](ApiControllersAgentControllerRestartAgentsByNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
