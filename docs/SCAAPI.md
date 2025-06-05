# \SCAAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersScaControllerGetScaAgent**](SCAAPI.md#ApiControllersScaControllerGetScaAgent) | **Get** /sca/{agent_id} | Get results
[**ApiControllersScaControllerGetScaChecks**](SCAAPI.md#ApiControllersScaControllerGetScaChecks) | **Get** /sca/{agent_id}/checks/{policy_id} | Get policy checks



## ApiControllersScaControllerGetScaAgent

> ApiControllersScaControllerGetScaAgent200Response ApiControllersScaControllerGetScaAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Name(name).Description(description).References(references).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

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
	name := "name_example" // string | Filter by policy name (optional)
	description := "description_example" // string | Filter by policy description (optional)
	references := "references_example" // string | Filter by references (optional)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SCAAPI.ApiControllersScaControllerGetScaAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Name(name).Description(description).References(references).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCAAPI.ApiControllersScaControllerGetScaAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersScaControllerGetScaAgent`: ApiControllersScaControllerGetScaAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `SCAAPI.ApiControllersScaControllerGetScaAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersScaControllerGetScaAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **name** | **string** | Filter by policy name |
 **description** | **string** | Filter by policy description |
 **references** | **string** | Filter by references |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersScaControllerGetScaAgent200Response**](ApiControllersScaControllerGetScaAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersScaControllerGetScaChecks

> ApiControllersScaControllerGetScaChecks200Response ApiControllersScaControllerGetScaChecks(ctx, agentId, policyId).Pretty(pretty).WaitForComplete(waitForComplete).Title(title).Description(description).Rationale(rationale).Remediation(remediation).Command(command).Reason(reason).File(file).Process(process).Directory(directory).Registry(registry).References(references).Result(result).Condition(condition).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get policy checks



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
	policyId := "policyId_example" // string | Filter by policy id
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	title := "title_example" // string | Filter by title (optional)
	description := "description_example" // string | Filter by policy description (optional)
	rationale := "rationale_example" // string | Filter by rationale (optional)
	remediation := "remediation_example" // string | Filter by remediation (optional)
	command := "command_example" // string | Filter by command (optional)
	reason := "reason_example" // string | Filter by reason (optional)
	file := "file_example" // string | Filter by full path (optional)
	process := "process_example" // string | Filter by process name (optional)
	directory := "directory_example" // string | Filter by directory (optional)
	registry := "registry_example" // string | Filter by registry (optional)
	references := "references_example" // string | Filter by references (optional)
	result := "result_example" // string | Filter by result (optional)
	condition := "condition_example" // string | Filter by condition (optional)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SCAAPI.ApiControllersScaControllerGetScaChecks(context.Background(), agentId, policyId).Pretty(pretty).WaitForComplete(waitForComplete).Title(title).Description(description).Rationale(rationale).Remediation(remediation).Command(command).Reason(reason).File(file).Process(process).Directory(directory).Registry(registry).References(references).Result(result).Condition(condition).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SCAAPI.ApiControllersScaControllerGetScaChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersScaControllerGetScaChecks`: ApiControllersScaControllerGetScaChecks200Response
	fmt.Fprintf(os.Stdout, "Response from `SCAAPI.ApiControllersScaControllerGetScaChecks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |
**policyId** | **string** | Filter by policy id |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersScaControllerGetScaChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **title** | **string** | Filter by title |
 **description** | **string** | Filter by policy description |
 **rationale** | **string** | Filter by rationale |
 **remediation** | **string** | Filter by remediation |
 **command** | **string** | Filter by command |
 **reason** | **string** | Filter by reason |
 **file** | **string** | Filter by full path |
 **process** | **string** | Filter by process name |
 **directory** | **string** | Filter by directory |
 **registry** | **string** | Filter by registry |
 **references** | **string** | Filter by references |
 **result** | **string** | Filter by result |
 **condition** | **string** | Filter by condition |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersScaControllerGetScaChecks200Response**](ApiControllersScaControllerGetScaChecks200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
