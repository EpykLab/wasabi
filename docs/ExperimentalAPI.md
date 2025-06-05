# \ExperimentalAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersExperimentalControllerClearRootcheckDatabase**](ExperimentalAPI.md#ApiControllersExperimentalControllerClearRootcheckDatabase) | **Delete** /experimental/rootcheck | Clear rootcheck results
[**ApiControllersExperimentalControllerClearSyscheckDatabase**](ExperimentalAPI.md#ApiControllersExperimentalControllerClearSyscheckDatabase) | **Delete** /experimental/syscheck | Clear agents FIM results
[**ApiControllersExperimentalControllerGetCisCatResults**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetCisCatResults) | **Get** /experimental/ciscat/results | Get agents CIS-CAT results
[**ApiControllersExperimentalControllerGetHardwareInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetHardwareInfo) | **Get** /experimental/syscollector/hardware | Get agents hardware
[**ApiControllersExperimentalControllerGetHotfixesInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetHotfixesInfo) | **Get** /experimental/syscollector/hotfixes | Get agents hotfixes
[**ApiControllersExperimentalControllerGetNetworkAddressInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetNetworkAddressInfo) | **Get** /experimental/syscollector/netaddr | Get agents netaddr
[**ApiControllersExperimentalControllerGetNetworkInterfaceInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetNetworkInterfaceInfo) | **Get** /experimental/syscollector/netiface | Get agents netiface
[**ApiControllersExperimentalControllerGetNetworkProtocolInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetNetworkProtocolInfo) | **Get** /experimental/syscollector/netproto | Get agents netproto
[**ApiControllersExperimentalControllerGetOsInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetOsInfo) | **Get** /experimental/syscollector/os | Get agents OS
[**ApiControllersExperimentalControllerGetPackagesInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetPackagesInfo) | **Get** /experimental/syscollector/packages | Get agents packages
[**ApiControllersExperimentalControllerGetPortsInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetPortsInfo) | **Get** /experimental/syscollector/ports | Get agents ports
[**ApiControllersExperimentalControllerGetProcessesInfo**](ExperimentalAPI.md#ApiControllersExperimentalControllerGetProcessesInfo) | **Get** /experimental/syscollector/processes | Get agents processes



## ApiControllersExperimentalControllerClearRootcheckDatabase

> ApiControllersMitreControllerGetGroups200Response ApiControllersExperimentalControllerClearRootcheckDatabase(ctx).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Clear rootcheck results



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerClearRootcheckDatabase(context.Background()).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerClearRootcheckDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerClearRootcheckDatabase`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerClearRootcheckDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerClearRootcheckDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
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


## ApiControllersExperimentalControllerClearSyscheckDatabase

> ApiControllersAgentControllerRestartAgentsByNode200Response ApiControllersExperimentalControllerClearSyscheckDatabase(ctx).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Clear agents FIM results



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerClearSyscheckDatabase(context.Background()).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerClearSyscheckDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerClearSyscheckDatabase`: ApiControllersAgentControllerRestartAgentsByNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerClearSyscheckDatabase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerClearSyscheckDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
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


## ApiControllersExperimentalControllerGetCisCatResults

> ApiControllersCiscatControllerGetAgentsCiscatResults200Response ApiControllersExperimentalControllerGetCisCatResults(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Benchmark(benchmark).Profile(profile).Pass(pass).Fail(fail).Error_(error_).Notchecked(notchecked).Unknown(unknown).Score(score).Execute()

Get agents CIS-CAT results



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	benchmark := "benchmark_example" // string | Filter by benchmark type (optional)
	profile := "profile_example" // string | Filter by evaluated profile (optional)
	pass := int32(56) // int32 | Filter by passed checks (optional)
	fail := int32(56) // int32 | Filter by failed checks (optional)
	error_ := int32(56) // int32 | Filter by encountered errors (optional)
	notchecked := int32(56) // int32 | Filter by not checked (optional)
	unknown := int32(56) // int32 | Filter by unknown results (optional)
	score := int32(56) // int32 | Filter by final score (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetCisCatResults(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Benchmark(benchmark).Profile(profile).Pass(pass).Fail(fail).Error_(error_).Notchecked(notchecked).Unknown(unknown).Score(score).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetCisCatResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetCisCatResults`: ApiControllersCiscatControllerGetAgentsCiscatResults200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetCisCatResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetCisCatResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **benchmark** | **string** | Filter by benchmark type |
 **profile** | **string** | Filter by evaluated profile |
 **pass** | **int32** | Filter by passed checks |
 **fail** | **int32** | Filter by failed checks |
 **error_** | **int32** | Filter by encountered errors |
 **notchecked** | **int32** | Filter by not checked |
 **unknown** | **int32** | Filter by unknown results |
 **score** | **int32** | Filter by final score |

### Return type

[**ApiControllersCiscatControllerGetAgentsCiscatResults200Response**](ApiControllersCiscatControllerGetAgentsCiscatResults200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetHardwareInfo

> ApiControllersExperimentalControllerGetHardwareInfo200Response ApiControllersExperimentalControllerGetHardwareInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).RamFree(ramFree).RamTotal(ramTotal).CpuCores(cpuCores).CpuMhz(cpuMhz).CpuName(cpuName).BoardSerial(boardSerial).Execute()

Get agents hardware



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	ramFree := int64(789) // int64 | Filter by ram.free (optional)
	ramTotal := int64(789) // int64 | Filter by ram.total (optional)
	cpuCores := int32(56) // int32 | Filter by cpu.cores (optional)
	cpuMhz := float32(3.4) // float32 | Filter by cpu.mhz (optional)
	cpuName := "cpuName_example" // string | Filter by cpu.name (optional)
	boardSerial := "boardSerial_example" // string | Filter by board_serial (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetHardwareInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).RamFree(ramFree).RamTotal(ramTotal).CpuCores(cpuCores).CpuMhz(cpuMhz).CpuName(cpuName).BoardSerial(boardSerial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetHardwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetHardwareInfo`: ApiControllersExperimentalControllerGetHardwareInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetHardwareInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetHardwareInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **ramFree** | **int64** | Filter by ram.free |
 **ramTotal** | **int64** | Filter by ram.total |
 **cpuCores** | **int32** | Filter by cpu.cores |
 **cpuMhz** | **float32** | Filter by cpu.mhz |
 **cpuName** | **string** | Filter by cpu.name |
 **boardSerial** | **string** | Filter by board_serial |

### Return type

[**ApiControllersExperimentalControllerGetHardwareInfo200Response**](ApiControllersExperimentalControllerGetHardwareInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetHotfixesInfo

> ApiControllersExperimentalControllerGetHotfixesInfo200Response ApiControllersExperimentalControllerGetHotfixesInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Hotfix(hotfix).Execute()

Get agents hotfixes



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	hotfix := "hotfix_example" // string | Filter by hotfix (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetHotfixesInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Hotfix(hotfix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetHotfixesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetHotfixesInfo`: ApiControllersExperimentalControllerGetHotfixesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetHotfixesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetHotfixesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **hotfix** | **string** | Filter by hotfix |

### Return type

[**ApiControllersExperimentalControllerGetHotfixesInfo200Response**](ApiControllersExperimentalControllerGetHotfixesInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetNetworkAddressInfo

> ApiControllersExperimentalControllerGetNetworkAddressInfo200Response ApiControllersExperimentalControllerGetNetworkAddressInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Proto(proto).Address(address).Broadcast(broadcast).Netmask(netmask).Execute()

Get agents netaddr



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	proto := "proto_example" // string | Filter by IP protocol (optional)
	address := "address_example" // string | Filter by IP address (optional)
	broadcast := "broadcast_example" // string | Filter by broadcast direction (optional)
	netmask := "netmask_example" // string | Filter by netmask (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkAddressInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Proto(proto).Address(address).Broadcast(broadcast).Netmask(netmask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkAddressInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetNetworkAddressInfo`: ApiControllersExperimentalControllerGetNetworkAddressInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkAddressInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetNetworkAddressInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **proto** | **string** | Filter by IP protocol |
 **address** | **string** | Filter by IP address |
 **broadcast** | **string** | Filter by broadcast direction |
 **netmask** | **string** | Filter by netmask |

### Return type

[**ApiControllersExperimentalControllerGetNetworkAddressInfo200Response**](ApiControllersExperimentalControllerGetNetworkAddressInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetNetworkInterfaceInfo

> ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response ApiControllersExperimentalControllerGetNetworkInterfaceInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Name(name).Adapter(adapter).Type_(type_).State(state).Mtu(mtu).TxPackets(txPackets).RxPackets(rxPackets).TxBytes(txBytes).RxBytes(rxBytes).TxErrors(txErrors).RxErrors(rxErrors).TxDropped(txDropped).RxDropped(rxDropped).Execute()

Get agents netiface



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	name := "name_example" // string | Filter by name (optional)
	adapter := "adapter_example" // string | Filter by adapter (optional)
	type_ := "type__example" // string | Type of network (optional)
	state := "state_example" // string | Filter by state (optional)
	mtu := int32(56) // int32 | Filter by mtu (optional)
	txPackets := int32(56) // int32 | Filter by tx.packets (optional)
	rxPackets := int32(56) // int32 | Filter by rx.packets (optional)
	txBytes := int32(56) // int32 | Filter by tx.bytes (optional)
	rxBytes := int32(56) // int32 | Filter by rx.bytes (optional)
	txErrors := int32(56) // int32 | Filter by tx.errors (optional)
	rxErrors := int32(56) // int32 | Filter by rx.errors (optional)
	txDropped := int32(56) // int32 | Filter by tx.dropped (optional)
	rxDropped := int32(56) // int32 | Filter by rx.dropped (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkInterfaceInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Name(name).Adapter(adapter).Type_(type_).State(state).Mtu(mtu).TxPackets(txPackets).RxPackets(rxPackets).TxBytes(txBytes).RxBytes(rxBytes).TxErrors(txErrors).RxErrors(rxErrors).TxDropped(txDropped).RxDropped(rxDropped).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkInterfaceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetNetworkInterfaceInfo`: ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetNetworkInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **name** | **string** | Filter by name |
 **adapter** | **string** | Filter by adapter |
 **type_** | **string** | Type of network |
 **state** | **string** | Filter by state |
 **mtu** | **int32** | Filter by mtu |
 **txPackets** | **int32** | Filter by tx.packets |
 **rxPackets** | **int32** | Filter by rx.packets |
 **txBytes** | **int32** | Filter by tx.bytes |
 **rxBytes** | **int32** | Filter by rx.bytes |
 **txErrors** | **int32** | Filter by tx.errors |
 **rxErrors** | **int32** | Filter by rx.errors |
 **txDropped** | **int32** | Filter by tx.dropped |
 **rxDropped** | **int32** | Filter by rx.dropped |

### Return type

[**ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response**](ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetNetworkProtocolInfo

> ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response ApiControllersExperimentalControllerGetNetworkProtocolInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Type_(type_).Gateway(gateway).Dhcp(dhcp).Execute()

Get agents netproto



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	iface := "iface_example" // string | Filter by network interface (optional)
	type_ := "type__example" // string | Type of network (optional)
	gateway := "gateway_example" // string | Filter by network gateway (optional)
	dhcp := openapiclient.DHCPStatus("enabled") // DHCPStatus | Filter by network dhcp (enabled or disabled) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkProtocolInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Type_(type_).Gateway(gateway).Dhcp(dhcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkProtocolInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetNetworkProtocolInfo`: ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkProtocolInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetNetworkProtocolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **iface** | **string** | Filter by network interface |
 **type_** | **string** | Type of network |
 **gateway** | **string** | Filter by network gateway |
 **dhcp** | [**DHCPStatus**](DHCPStatus.md) | Filter by network dhcp (enabled or disabled) |

### Return type

[**ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response**](ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetOsInfo

> ApiControllersExperimentalControllerGetOsInfo200Response ApiControllersExperimentalControllerGetOsInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).OsName(osName).Architecture(architecture).OsVersion(osVersion).Version(version).Release(release).Execute()

Get agents OS



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	architecture := "architecture_example" // string | Filter by architecture (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	release := "release_example" // string | Filter by release (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetOsInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).OsName(osName).Architecture(architecture).OsVersion(osVersion).Version(version).Release(release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetOsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetOsInfo`: ApiControllersExperimentalControllerGetOsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetOsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetOsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **osName** | **string** | Filter by OS name |
 **architecture** | **string** | Filter by architecture |
 **osVersion** | **string** | Filter by OS version |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **release** | **string** | Filter by release |

### Return type

[**ApiControllersExperimentalControllerGetOsInfo200Response**](ApiControllersExperimentalControllerGetOsInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetPackagesInfo

> ApiControllersExperimentalControllerGetPackagesInfo200Response ApiControllersExperimentalControllerGetPackagesInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Vendor(vendor).Name(name).Architecture(architecture).Format(format).Version(version).Execute()

Get agents packages



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	vendor := "vendor_example" // string | Filter by vendor (optional)
	name := "name_example" // string | Filter by name (optional)
	architecture := "architecture_example" // string | Filter by architecture (optional)
	format := "format_example" // string | Filter by file format. For example 'deb' will output deb files (optional)
	version := "version_example" // string | Filter by package version (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetPackagesInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Vendor(vendor).Name(name).Architecture(architecture).Format(format).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetPackagesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetPackagesInfo`: ApiControllersExperimentalControllerGetPackagesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetPackagesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetPackagesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **vendor** | **string** | Filter by vendor |
 **name** | **string** | Filter by name |
 **architecture** | **string** | Filter by architecture |
 **format** | **string** | Filter by file format. For example &#39;deb&#39; will output deb files |
 **version** | **string** | Filter by package version |

### Return type

[**ApiControllersExperimentalControllerGetPackagesInfo200Response**](ApiControllersExperimentalControllerGetPackagesInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetPortsInfo

> ApiControllersExperimentalControllerGetPortsInfo200Response ApiControllersExperimentalControllerGetPortsInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).Protocol(protocol).LocalIp(localIp).LocalPort(localPort).RemoteIp(remoteIp).TxQueue(txQueue).State(state).Process(process).Execute()

Get agents ports



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	pid := "pid_example" // string | Filter by pid (optional)
	protocol := "protocol_example" // string | Filter by protocol (optional)
	localIp := "localIp_example" // string | Filter by Local IP (optional)
	localPort := "localPort_example" // string | Filter by Local Port (optional)
	remoteIp := "remoteIp_example" // string | Filter by Remote IP (optional)
	txQueue := "txQueue_example" // string | Filter by tx_queue (optional)
	state := "state_example" // string | Filter by state (optional)
	process := "process_example" // string | Filter by process name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetPortsInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).Protocol(protocol).LocalIp(localIp).LocalPort(localPort).RemoteIp(remoteIp).TxQueue(txQueue).State(state).Process(process).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetPortsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetPortsInfo`: ApiControllersExperimentalControllerGetPortsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetPortsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetPortsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **pid** | **string** | Filter by pid |
 **protocol** | **string** | Filter by protocol |
 **localIp** | **string** | Filter by Local IP |
 **localPort** | **string** | Filter by Local Port |
 **remoteIp** | **string** | Filter by Remote IP |
 **txQueue** | **string** | Filter by tx_queue |
 **state** | **string** | Filter by state |
 **process** | **string** | Filter by process name |

### Return type

[**ApiControllersExperimentalControllerGetPortsInfo200Response**](ApiControllersExperimentalControllerGetPortsInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersExperimentalControllerGetProcessesInfo

> ApiControllersExperimentalControllerGetProcessesInfo200Response ApiControllersExperimentalControllerGetProcessesInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).State(state).Ppid(ppid).Egroup(egroup).Euser(euser).Fgroup(fgroup).Name(name).Nlwp(nlwp).Pgrp(pgrp).Priority(priority).Rgroup(rgroup).Ruser(ruser).Sgroup(sgroup).Suser(suser).Execute()

Get agents processes



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	pid := "pid_example" // string | Filter by process pid (optional)
	state := "state_example" // string | Filter by process state (optional)
	ppid := "ppid_example" // string | Filter by process parent pid (optional)
	egroup := "egroup_example" // string | Filter by process egroup (optional)
	euser := "euser_example" // string | Filter by process euser (optional)
	fgroup := "fgroup_example" // string | Filter by process fgroup (optional)
	name := "name_example" // string | Filter by process name (optional)
	nlwp := "nlwp_example" // string | Filter by process nlwp (optional)
	pgrp := "pgrp_example" // string | Filter by process pgrp (optional)
	priority := "priority_example" // string | Filter by process priority (optional)
	rgroup := "rgroup_example" // string | Filter by process rgroup (optional)
	ruser := "ruser_example" // string | Filter by process ruser (optional)
	sgroup := "sgroup_example" // string | Filter by process sgroup (optional)
	suser := "suser_example" // string | Filter by process suser (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetProcessesInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).State(state).Ppid(ppid).Egroup(egroup).Euser(euser).Fgroup(fgroup).Name(name).Nlwp(nlwp).Pgrp(pgrp).Priority(priority).Rgroup(rgroup).Ruser(ruser).Sgroup(sgroup).Suser(suser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalAPI.ApiControllersExperimentalControllerGetProcessesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersExperimentalControllerGetProcessesInfo`: ApiControllersExperimentalControllerGetProcessesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `ExperimentalAPI.ApiControllersExperimentalControllerGetProcessesInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersExperimentalControllerGetProcessesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **pid** | **string** | Filter by process pid |
 **state** | **string** | Filter by process state |
 **ppid** | **string** | Filter by process parent pid |
 **egroup** | **string** | Filter by process egroup |
 **euser** | **string** | Filter by process euser |
 **fgroup** | **string** | Filter by process fgroup |
 **name** | **string** | Filter by process name |
 **nlwp** | **string** | Filter by process nlwp |
 **pgrp** | **string** | Filter by process pgrp |
 **priority** | **string** | Filter by process priority |
 **rgroup** | **string** | Filter by process rgroup |
 **ruser** | **string** | Filter by process ruser |
 **sgroup** | **string** | Filter by process sgroup |
 **suser** | **string** | Filter by process suser |

### Return type

[**ApiControllersExperimentalControllerGetProcessesInfo200Response**](ApiControllersExperimentalControllerGetProcessesInfo200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
