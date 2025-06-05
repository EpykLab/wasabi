# \ClusterAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersClusterControllerGetApiConfig**](ClusterAPI.md#ApiControllersClusterControllerGetApiConfig) | **Get** /cluster/api/config | Get nodes API config
[**ApiControllersClusterControllerGetClusterNode**](ClusterAPI.md#ApiControllersClusterControllerGetClusterNode) | **Get** /cluster/local/info | Get local node info
[**ApiControllersClusterControllerGetClusterNodes**](ClusterAPI.md#ApiControllersClusterControllerGetClusterNodes) | **Get** /cluster/nodes | Get nodes info
[**ApiControllersClusterControllerGetConfValidation**](ClusterAPI.md#ApiControllersClusterControllerGetConfValidation) | **Get** /cluster/configuration/validation | Check nodes config
[**ApiControllersClusterControllerGetConfig**](ClusterAPI.md#ApiControllersClusterControllerGetConfig) | **Get** /cluster/local/config | Get local node config
[**ApiControllersClusterControllerGetConfigurationNode**](ClusterAPI.md#ApiControllersClusterControllerGetConfigurationNode) | **Get** /cluster/{node_id}/configuration | Get node config
[**ApiControllersClusterControllerGetDaemonStatsNode**](ClusterAPI.md#ApiControllersClusterControllerGetDaemonStatsNode) | **Get** /cluster/{node_id}/daemons/stats | Get Wazuh daemon stats from a cluster node
[**ApiControllersClusterControllerGetHealthcheck**](ClusterAPI.md#ApiControllersClusterControllerGetHealthcheck) | **Get** /cluster/healthcheck | Get nodes healthcheck
[**ApiControllersClusterControllerGetInfoNode**](ClusterAPI.md#ApiControllersClusterControllerGetInfoNode) | **Get** /cluster/{node_id}/info | Get node info
[**ApiControllersClusterControllerGetLogNode**](ClusterAPI.md#ApiControllersClusterControllerGetLogNode) | **Get** /cluster/{node_id}/logs | Get node logs
[**ApiControllersClusterControllerGetLogSummaryNode**](ClusterAPI.md#ApiControllersClusterControllerGetLogSummaryNode) | **Get** /cluster/{node_id}/logs/summary | Get node logs summary
[**ApiControllersClusterControllerGetNodeConfig**](ClusterAPI.md#ApiControllersClusterControllerGetNodeConfig) | **Get** /cluster/{node_id}/configuration/{component}/{configuration} | Get node active configuration
[**ApiControllersClusterControllerGetNodesRulesetSyncStatus**](ClusterAPI.md#ApiControllersClusterControllerGetNodesRulesetSyncStatus) | **Get** /cluster/ruleset/synchronization | Get cluster nodes ruleset synchronization status
[**ApiControllersClusterControllerGetStatsAnalysisdNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatsAnalysisdNode) | **Get** /cluster/{node_id}/stats/analysisd | Get node stats analysisd
[**ApiControllersClusterControllerGetStatsHourlyNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatsHourlyNode) | **Get** /cluster/{node_id}/stats/hourly | Get node stats hour
[**ApiControllersClusterControllerGetStatsNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatsNode) | **Get** /cluster/{node_id}/stats | Get node stats
[**ApiControllersClusterControllerGetStatsRemotedNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatsRemotedNode) | **Get** /cluster/{node_id}/stats/remoted | Get node stats remoted
[**ApiControllersClusterControllerGetStatsWeeklyNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatsWeeklyNode) | **Get** /cluster/{node_id}/stats/weekly | Get node stats week
[**ApiControllersClusterControllerGetStatus**](ClusterAPI.md#ApiControllersClusterControllerGetStatus) | **Get** /cluster/status | Get cluster status
[**ApiControllersClusterControllerGetStatusNode**](ClusterAPI.md#ApiControllersClusterControllerGetStatusNode) | **Get** /cluster/{node_id}/status | Get node status
[**ApiControllersClusterControllerPutRestart**](ClusterAPI.md#ApiControllersClusterControllerPutRestart) | **Put** /cluster/restart | Restart nodes
[**ApiControllersClusterControllerUpdateConfiguration**](ClusterAPI.md#ApiControllersClusterControllerUpdateConfiguration) | **Put** /cluster/{node_id}/configuration | Update node configuration



## ApiControllersClusterControllerGetApiConfig

> map[string]interface{} ApiControllersClusterControllerGetApiConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()

Get nodes API config



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
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetApiConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetApiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetApiConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetApiConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetApiConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetClusterNode

> ApiControllersClusterControllerGetClusterNode200Response ApiControllersClusterControllerGetClusterNode(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get local node info



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
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetClusterNode(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetClusterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetClusterNode`: ApiControllersClusterControllerGetClusterNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetClusterNode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetClusterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetClusterNode200Response**](ApiControllersClusterControllerGetClusterNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetClusterNodes

> ApiControllersClusterControllerGetClusterNodes200Response ApiControllersClusterControllerGetClusterNodes(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Type_(type_).NodesList(nodesList).Q(q).Distinct(distinct).Execute()

Get nodes info



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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	type_ := "type__example" // string | Filter by node type (optional)
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetClusterNodes(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Type_(type_).NodesList(nodesList).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetClusterNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetClusterNodes`: ApiControllersClusterControllerGetClusterNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetClusterNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **type_** | **string** | Filter by node type |
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersClusterControllerGetClusterNodes200Response**](ApiControllersClusterControllerGetClusterNodes200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetConfValidation

> ApiControllersClusterControllerGetConfValidation200Response ApiControllersClusterControllerGetConfValidation(ctx).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()

Check nodes config



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
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetConfValidation(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetConfValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetConfValidation`: ApiControllersClusterControllerGetConfValidation200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetConfValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetConfValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |

### Return type

[**ApiControllersClusterControllerGetConfValidation200Response**](ApiControllersClusterControllerGetConfValidation200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetConfig

> ApiControllersClusterControllerGetConfig200Response ApiControllersClusterControllerGetConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get local node config



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
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetConfig`: ApiControllersClusterControllerGetConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetConfig200Response**](ApiControllersClusterControllerGetConfig200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetConfigurationNode

> ApiControllersClusterControllerGetConfigurationNode200Response ApiControllersClusterControllerGetConfigurationNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Section(section).Field(field).Execute()

Get node config



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	raw := true // bool | Format response in plain text (optional)
	section := "section_example" // string | Indicates the wazuh configuration section (optional)
	field := "field_example" // string | Indicate a section child. E.g, fields for *ruleset* section are: decoder_dir, rule_dir, etc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetConfigurationNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Section(section).Field(field).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetConfigurationNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetConfigurationNode`: ApiControllersClusterControllerGetConfigurationNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetConfigurationNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetConfigurationNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **raw** | **bool** | Format response in plain text |
 **section** | **string** | Indicates the wazuh configuration section |
 **field** | **string** | Indicate a section child. E.g, fields for *ruleset* section are: decoder_dir, rule_dir, etc |

### Return type

[**ApiControllersClusterControllerGetConfigurationNode200Response**](ApiControllersClusterControllerGetConfigurationNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetDaemonStatsNode

> ApiControllersClusterControllerGetDaemonStatsNode200Response ApiControllersClusterControllerGetDaemonStatsNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()

Get Wazuh daemon stats from a cluster node



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	daemonsList := []string{"DaemonsList_example"} // []string | List of daemon names (separated by comma), all daemons selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetDaemonStatsNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetDaemonStatsNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetDaemonStatsNode`: ApiControllersClusterControllerGetDaemonStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetDaemonStatsNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetDaemonStatsNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **daemonsList** | **[]string** | List of daemon names (separated by comma), all daemons selected by default if not specified |

### Return type

[**ApiControllersClusterControllerGetDaemonStatsNode200Response**](ApiControllersClusterControllerGetDaemonStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetHealthcheck

> ApiControllersClusterControllerGetHealthcheck200Response ApiControllersClusterControllerGetHealthcheck(ctx).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()

Get nodes healthcheck



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
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetHealthcheck(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetHealthcheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetHealthcheck`: ApiControllersClusterControllerGetHealthcheck200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetHealthcheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetHealthcheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |

### Return type

[**ApiControllersClusterControllerGetHealthcheck200Response**](ApiControllersClusterControllerGetHealthcheck200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetInfoNode

> ApiControllersClusterControllerGetInfoNode200Response ApiControllersClusterControllerGetInfoNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node info



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetInfoNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetInfoNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetInfoNode`: ApiControllersClusterControllerGetInfoNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetInfoNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetInfoNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetInfoNode200Response**](ApiControllersClusterControllerGetInfoNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetLogNode

> ApiControllersClusterControllerGetLogNode200Response ApiControllersClusterControllerGetLogNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Tag(tag).Level(level).Q(q).Select_(select_).Distinct(distinct).Execute()

Get node logs



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of lines to return. (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	tag := "tag_example" // string | Wazuh component that logged the event (optional)
	level := "level_example" // string | Filter by log level (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetLogNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Tag(tag).Level(level).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetLogNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetLogNode`: ApiControllersClusterControllerGetLogNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetLogNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetLogNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of lines to return. | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **tag** | **string** | Wazuh component that logged the event |
 **level** | **string** | Filter by log level |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersClusterControllerGetLogNode200Response**](ApiControllersClusterControllerGetLogNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetLogSummaryNode

> ApiControllersClusterControllerGetLogSummaryNode200Response ApiControllersClusterControllerGetLogSummaryNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node logs summary



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetLogSummaryNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetLogSummaryNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetLogSummaryNode`: ApiControllersClusterControllerGetLogSummaryNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetLogSummaryNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetLogSummaryNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetLogSummaryNode200Response**](ApiControllersClusterControllerGetLogSummaryNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetNodeConfig

> ApiControllersClusterControllerGetNodeConfig200Response ApiControllersClusterControllerGetNodeConfig(ctx, nodeId, component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node active configuration



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
	nodeId := "nodeId_example" // string | Cluster node name
	component := "component_example" // string | Selected agent's component
	configuration := "configuration_example" // string | <p>Selected agent's configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:</p> <table class=\"table table-striped table-bordered\"> <thead> <tr> <th>Component</th> <th>Configuration</th> <th>Tag</th> </tr> </thead> <tbody> <tr> <td>agent</td> <td>client</td> <td><code>&lt;client&gt;</code></td> </tr> <tr> <td>agent</td> <td>buffer</td> <td><code>&lt;client_buffer&gt;</code></td> </tr> <tr> <td>agent</td> <td>labels</td> <td><code>&lt;labels&gt;</code></td> </tr> <tr> <td>agent</td> <td>internal</td> <td><code>&lt;agent&gt;</code>, <code>&lt;monitord&gt;</code>, <code>&lt;remoted&gt;</code></td> </tr> <tr> <td>agent</td> <td>anti_tampering</td> <td><code>&lt;anti_tampering&gt;</code></td> </tr> <tr> <td>agentless</td> <td>agentless</td> <td><code>&lt;agentless&gt;</code></td> </tr> <tr> <td>analysis</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>analysis</td> <td>active_response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>analysis</td> <td>alerts</td> <td><code>&lt;alerts&gt;</code></td> </tr> <tr> <td>analysis</td> <td>command</td> <td><code>&lt;command&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rules</td> <td><code>&lt;rule&gt;</code></td> </tr> <tr> <td>analysis</td> <td>decoders</td> <td><code>&lt;decoder&gt;</code></td> </tr> <tr> <td>analysis</td> <td>internal</td> <td><code>&lt;analysisd&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rule_test</td> <td><code>&lt;rule_test&gt;</code></td> </tr> <tr> <td>auth</td> <td>auth</td> <td><code>&lt;auth&gt;</code></td> </tr> <tr> <td>com</td> <td>active-response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>com</td> <td>logging</td> <td><code>&lt;logging&gt;</code></td> </tr> <tr> <td>com</td> <td>internal</td> <td><code>&lt;execd&gt;</code></td> </tr> <tr> <td>com</td> <td>cluster</td> <td><code>&lt;cluster&gt;</code></td> </tr> <tr> <td>csyslog</td> <td>csyslog</td> <td><code>&lt;csyslog_output&gt;</code></td> </tr> <tr> <td>integrator</td> <td>integration</td> <td><code>&lt;integration&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>localfile</td> <td><code>&lt;localfile&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>socket</td> <td><code>&lt;socket&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>internal</td> <td><code>&lt;logcollector&gt;</code></td> </tr> <tr> <td>mail</td> <td>global</td> <td><code>&lt;global&gt;&lt;email...&gt;</code></td> </tr> <tr> <td>mail</td> <td>alerts</td> <td><code>&lt;email_alerts&gt;</code></td> </tr> <tr> <td>mail</td> <td>internal</td> <td><code>&lt;maild&gt;</code></td> </tr> <tr> <td>monitor</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>monitor</td> <td>internal</td> <td><code>&lt;monitord&gt;</code></td> </tr> <tr> <td>monitor</td> <td>reports</td> <td><code>&lt;reports&gt;</code></td> </tr> <tr> <td>request</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>request</td> <td>remote</td> <td><code>&lt;remote&gt;</code></td> </tr> <tr> <td>request</td> <td>internal</td> <td><code>&lt;remoted&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>syscheck</td> <td><code>&lt;syscheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>rootcheck</td> <td><code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>internal</td> <td><code>&lt;syscheck&gt;</code>, <code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>internal</td> <td><code>&lt;wazuh_db&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>wdb</td> <td><code>&lt;wdb&gt;</code></td> </tr> <tr> <td>wmodules</td> <td>wmodules</td> <td><code>&lt;wodle&gt;</code></td> </tr> </tbody> </table>
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetNodeConfig(context.Background(), nodeId, component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetNodeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetNodeConfig`: ApiControllersClusterControllerGetNodeConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetNodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |
**component** | **string** | Selected agent&#39;s component |
**configuration** | **string** | &lt;p&gt;Selected agent&#39;s configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:&lt;/p&gt; &lt;table class&#x3D;\&quot;table table-striped table-bordered\&quot;&gt; &lt;thead&gt; &lt;tr&gt; &lt;th&gt;Component&lt;/th&gt; &lt;th&gt;Configuration&lt;/th&gt; &lt;th&gt;Tag&lt;/th&gt; &lt;/tr&gt; &lt;/thead&gt; &lt;tbody&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;client&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;buffer&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client_buffer&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;labels&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;labels&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agent&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;anti_tampering&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;anti_tampering&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agentless&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;active_response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;command&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;command&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;decoders&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;decoder&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;analysisd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rule_test&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule_test&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;auth&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;active-response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;logging&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logging&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;execd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;cluster&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;cluster&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;csyslog_output&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;integrator&lt;/td&gt; &lt;td&gt;integration&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;integration&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;localfile&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;localfile&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;socket&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;socket&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logcollector&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&amp;lt;email...&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;email_alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;maild&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;reports&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;reports&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;remote&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remote&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;rootcheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wazuh_db&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;wdb&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wdb&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wodle&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;/tbody&gt; &lt;/table&gt;  |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetNodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetNodeConfig200Response**](ApiControllersClusterControllerGetNodeConfig200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetNodesRulesetSyncStatus

> ApiControllersClusterControllerGetNodesRulesetSyncStatus200Response ApiControllersClusterControllerGetNodesRulesetSyncStatus(ctx).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()

Get cluster nodes ruleset synchronization status



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
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetNodesRulesetSyncStatus(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetNodesRulesetSyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetNodesRulesetSyncStatus`: ApiControllersClusterControllerGetNodesRulesetSyncStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetNodesRulesetSyncStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetNodesRulesetSyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |

### Return type

[**ApiControllersClusterControllerGetNodesRulesetSyncStatus200Response**](ApiControllersClusterControllerGetNodesRulesetSyncStatus200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatsAnalysisdNode

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersClusterControllerGetStatsAnalysisdNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node stats analysisd



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatsAnalysisdNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatsAnalysisdNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatsAnalysisdNode`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatsAnalysisdNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatsAnalysisdNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatsNode200Response**](ApiControllersClusterControllerGetStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatsHourlyNode

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersClusterControllerGetStatsHourlyNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node stats hour



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatsHourlyNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatsHourlyNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatsHourlyNode`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatsHourlyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatsHourlyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatsNode200Response**](ApiControllersClusterControllerGetStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatsNode

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersClusterControllerGetStatsNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Date(date).Execute()

Get node stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	date := time.Now() // string | Date to obtain statistical information from. Format YYYY-MM-DD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatsNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatsNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatsNode`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatsNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatsNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **date** | **string** | Date to obtain statistical information from. Format YYYY-MM-DD |

### Return type

[**ApiControllersClusterControllerGetStatsNode200Response**](ApiControllersClusterControllerGetStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatsRemotedNode

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersClusterControllerGetStatsRemotedNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node stats remoted



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatsRemotedNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatsRemotedNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatsRemotedNode`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatsRemotedNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatsRemotedNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatsNode200Response**](ApiControllersClusterControllerGetStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatsWeeklyNode

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersClusterControllerGetStatsWeeklyNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node stats week



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatsWeeklyNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatsWeeklyNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatsWeeklyNode`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatsWeeklyNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatsWeeklyNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatsNode200Response**](ApiControllersClusterControllerGetStatsNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatus

> ApiControllersClusterControllerGetStatus200Response ApiControllersClusterControllerGetStatus(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get cluster status



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
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatus(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatus`: ApiControllersClusterControllerGetStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatus200Response**](ApiControllersClusterControllerGetStatus200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerGetStatusNode

> ApiControllersClusterControllerGetStatusNode200Response ApiControllersClusterControllerGetStatusNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get node status



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
	nodeId := "nodeId_example" // string | Cluster node name
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerGetStatusNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerGetStatusNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerGetStatusNode`: ApiControllersClusterControllerGetStatusNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerGetStatusNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerGetStatusNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerGetStatusNode200Response**](ApiControllersClusterControllerGetStatusNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerPutRestart

> ApiControllersClusterControllerPutRestart200Response ApiControllersClusterControllerPutRestart(ctx).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()

Restart nodes



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
	nodesList := []string{"Inner_example"} // []string | List of node IDs (separated by comma), all nodes selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerPutRestart(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).NodesList(nodesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerPutRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerPutRestart`: ApiControllersClusterControllerPutRestart200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerPutRestart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerPutRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **nodesList** | **[]string** | List of node IDs (separated by comma), all nodes selected by default if not specified |

### Return type

[**ApiControllersClusterControllerPutRestart200Response**](ApiControllersClusterControllerPutRestart200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersClusterControllerUpdateConfiguration

> ApiControllersClusterControllerUpdateConfiguration200Response ApiControllersClusterControllerUpdateConfiguration(ctx, nodeId).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Update node configuration



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
	nodeId := "nodeId_example" // string | Cluster node name
	body := os.NewFile(1234, "some_file") // *os.File | Content of the ossec.conf to be uploaded
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.ApiControllersClusterControllerUpdateConfiguration(context.Background(), nodeId).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.ApiControllersClusterControllerUpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersClusterControllerUpdateConfiguration`: ApiControllersClusterControllerUpdateConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.ApiControllersClusterControllerUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersClusterControllerUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Content of the ossec.conf to be uploaded |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerUpdateConfiguration200Response**](ApiControllersClusterControllerUpdateConfiguration200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
