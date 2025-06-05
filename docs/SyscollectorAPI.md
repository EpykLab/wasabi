# \SyscollectorAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersSyscollectorControllerGetHardwareInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetHardwareInfo) | **Get** /syscollector/{agent_id}/hardware | Get agent hardware
[**ApiControllersSyscollectorControllerGetHotfixInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetHotfixInfo) | **Get** /syscollector/{agent_id}/hotfixes | Get agent hotfixes
[**ApiControllersSyscollectorControllerGetNetworkAddressInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetNetworkAddressInfo) | **Get** /syscollector/{agent_id}/netaddr | Get agent netaddr
[**ApiControllersSyscollectorControllerGetNetworkInterfaceInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetNetworkInterfaceInfo) | **Get** /syscollector/{agent_id}/netiface | Get agent netiface
[**ApiControllersSyscollectorControllerGetNetworkProtocolInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetNetworkProtocolInfo) | **Get** /syscollector/{agent_id}/netproto | Get agent netproto
[**ApiControllersSyscollectorControllerGetOsInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetOsInfo) | **Get** /syscollector/{agent_id}/os | Get agent OS
[**ApiControllersSyscollectorControllerGetPackagesInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetPackagesInfo) | **Get** /syscollector/{agent_id}/packages | Get agent packages
[**ApiControllersSyscollectorControllerGetPortsInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetPortsInfo) | **Get** /syscollector/{agent_id}/ports | Get agent ports
[**ApiControllersSyscollectorControllerGetProcessesInfo**](SyscollectorAPI.md#ApiControllersSyscollectorControllerGetProcessesInfo) | **Get** /syscollector/{agent_id}/processes | Get agent processes



## ApiControllersSyscollectorControllerGetHardwareInfo

> ApiControllersExperimentalControllerGetHardwareInfo200Response ApiControllersSyscollectorControllerGetHardwareInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Select_(select_).Execute()

Get agent hardware



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
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetHardwareInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetHardwareInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetHardwareInfo`: ApiControllersExperimentalControllerGetHardwareInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetHardwareInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetHardwareInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |

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


## ApiControllersSyscollectorControllerGetHotfixInfo

> ApiControllersExperimentalControllerGetHotfixesInfo200Response ApiControllersSyscollectorControllerGetHotfixInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Hotfix(hotfix).Q(q).Distinct(distinct).Execute()

Get agent hotfixes



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
	hotfix := "hotfix_example" // string | Filter by hotfix (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetHotfixInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Hotfix(hotfix).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetHotfixInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetHotfixInfo`: ApiControllersExperimentalControllerGetHotfixesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetHotfixInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetHotfixInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **hotfix** | **string** | Filter by hotfix |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetNetworkAddressInfo

> ApiControllersExperimentalControllerGetNetworkAddressInfo200Response ApiControllersSyscollectorControllerGetNetworkAddressInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Proto(proto).Address(address).Broadcast(broadcast).Netmask(netmask).Q(q).Distinct(distinct).Execute()

Get agent netaddr



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
	iface := "iface_example" // string | Filter by network interface (optional)
	proto := "proto_example" // string | Filter by IP protocol (optional)
	address := "address_example" // string | Filter by IP address (optional)
	broadcast := "broadcast_example" // string | Filter by broadcast direction (optional)
	netmask := "netmask_example" // string | Filter by netmask (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkAddressInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Proto(proto).Address(address).Broadcast(broadcast).Netmask(netmask).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkAddressInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetNetworkAddressInfo`: ApiControllersExperimentalControllerGetNetworkAddressInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkAddressInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetNetworkAddressInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **iface** | **string** | Filter by network interface |
 **proto** | **string** | Filter by IP protocol |
 **address** | **string** | Filter by IP address |
 **broadcast** | **string** | Filter by broadcast direction |
 **netmask** | **string** | Filter by netmask |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetNetworkInterfaceInfo

> ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response ApiControllersSyscollectorControllerGetNetworkInterfaceInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Name(name).Adapter(adapter).Type_(type_).State(state).Mtu(mtu).TxPackets(txPackets).RxPackets(rxPackets).TxBytes(txBytes).RxBytes(rxBytes).TxErrors(txErrors).RxErrors(rxErrors).TxDropped(txDropped).RxDropped(rxDropped).Q(q).Distinct(distinct).Execute()

Get agent netiface



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
	name := "name_example" // string | Filter by name (optional)
	adapter := "adapter_example" // string | Filter by adapter (optional)
	type_ := "type__example" // string | Type of interface (optional)
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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkInterfaceInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Name(name).Adapter(adapter).Type_(type_).State(state).Mtu(mtu).TxPackets(txPackets).RxPackets(rxPackets).TxBytes(txBytes).RxBytes(rxBytes).TxErrors(txErrors).RxErrors(rxErrors).TxDropped(txDropped).RxDropped(rxDropped).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkInterfaceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetNetworkInterfaceInfo`: ApiControllersExperimentalControllerGetNetworkInterfaceInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetNetworkInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **name** | **string** | Filter by name |
 **adapter** | **string** | Filter by adapter |
 **type_** | **string** | Type of interface |
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
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetNetworkProtocolInfo

> ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response ApiControllersSyscollectorControllerGetNetworkProtocolInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Type_(type_).Gateway(gateway).Dhcp(dhcp).Q(q).Distinct(distinct).Execute()

Get agent netproto



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
	iface := "iface_example" // string | Filter by network interface (optional)
	type_ := "type__example" // string | Type of network (optional)
	gateway := "gateway_example" // string | Filter by network gateway (optional)
	dhcp := openapiclient.DHCPStatus("enabled") // DHCPStatus | Filter by network dhcp (enabled or disabled) (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkProtocolInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Iface(iface).Type_(type_).Gateway(gateway).Dhcp(dhcp).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkProtocolInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetNetworkProtocolInfo`: ApiControllersExperimentalControllerGetNetworkProtocolInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkProtocolInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetNetworkProtocolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **iface** | **string** | Filter by network interface |
 **type_** | **string** | Type of network |
 **gateway** | **string** | Filter by network gateway |
 **dhcp** | [**DHCPStatus**](DHCPStatus.md) | Filter by network dhcp (enabled or disabled) |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetOsInfo

> ApiControllersExperimentalControllerGetOsInfo200Response ApiControllersSyscollectorControllerGetOsInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Select_(select_).Execute()

Get agent OS



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
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetOsInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetOsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetOsInfo`: ApiControllersExperimentalControllerGetOsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetOsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetOsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |

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


## ApiControllersSyscollectorControllerGetPackagesInfo

> ApiControllersExperimentalControllerGetPackagesInfo200Response ApiControllersSyscollectorControllerGetPackagesInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Vendor(vendor).Name(name).Architecture(architecture).Format(format).Version(version).Q(q).Distinct(distinct).Execute()

Get agent packages



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
	vendor := "vendor_example" // string | Filter by vendor (optional)
	name := "name_example" // string | Filter by name (optional)
	architecture := "architecture_example" // string | Filter by architecture (optional)
	format := "format_example" // string | Filter by file format. For example 'deb' will output deb files (optional)
	version := "version_example" // string | Filter by package version (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetPackagesInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Vendor(vendor).Name(name).Architecture(architecture).Format(format).Version(version).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetPackagesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetPackagesInfo`: ApiControllersExperimentalControllerGetPackagesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetPackagesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetPackagesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
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
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetPortsInfo

> ApiControllersExperimentalControllerGetPortsInfo200Response ApiControllersSyscollectorControllerGetPortsInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).Protocol(protocol).LocalIp(localIp).LocalPort(localPort).RemoteIp(remoteIp).TxQueue(txQueue).State(state).Process(process).Q(q).Distinct(distinct).Execute()

Get agent ports



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
	pid := "pid_example" // string | Filter by pid (optional)
	protocol := "protocol_example" // string | Filter by protocol (optional)
	localIp := "localIp_example" // string | Filter by Local IP (optional)
	localPort := "localPort_example" // string | Filter by Local Port (optional)
	remoteIp := "remoteIp_example" // string | Filter by Remote IP (optional)
	txQueue := "txQueue_example" // string | Filter by tx_queue (optional)
	state := "state_example" // string | Filter by state (optional)
	process := "process_example" // string | Filter by process name (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetPortsInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).Protocol(protocol).LocalIp(localIp).LocalPort(localPort).RemoteIp(remoteIp).TxQueue(txQueue).State(state).Process(process).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetPortsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetPortsInfo`: ApiControllersExperimentalControllerGetPortsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetPortsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetPortsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
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
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersSyscollectorControllerGetProcessesInfo

> ApiControllersExperimentalControllerGetProcessesInfo200Response ApiControllersSyscollectorControllerGetProcessesInfo(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).State(state).Ppid(ppid).Egroup(egroup).Euser(euser).Fgroup(fgroup).Name(name).Nlwp(nlwp).Pgrp(pgrp).Priority(priority).Rgroup(rgroup).Ruser(ruser).Sgroup(sgroup).Suser(suser).Q(q).Distinct(distinct).Execute()

Get agent processes



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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetProcessesInfo(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Pid(pid).State(state).Ppid(ppid).Egroup(egroup).Euser(euser).Fgroup(fgroup).Name(name).Nlwp(nlwp).Pgrp(pgrp).Priority(priority).Rgroup(rgroup).Ruser(ruser).Sgroup(sgroup).Suser(suser).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyscollectorAPI.ApiControllersSyscollectorControllerGetProcessesInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSyscollectorControllerGetProcessesInfo`: ApiControllersExperimentalControllerGetProcessesInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `SyscollectorAPI.ApiControllersSyscollectorControllerGetProcessesInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSyscollectorControllerGetProcessesInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
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
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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
