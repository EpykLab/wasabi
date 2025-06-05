# \TasksAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersTaskControllerGetTasksStatus**](TasksAPI.md#ApiControllersTaskControllerGetTasksStatus) | **Get** /tasks/status | List tasks



## ApiControllersTaskControllerGetTasksStatus

> ApiResponse ApiControllersTaskControllerGetTasksStatus(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Q(q).Search(search).Select_(select_).Sort(sort).AgentsList(agentsList).TasksList(tasksList).Command(command).Node(node).Module(module).Status(status).Execute()

List tasks



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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	tasksList := []string{"Inner_example"} // []string | List of task IDs (separated by comma) (optional)
	command := "command_example" // string | Filter by command (optional)
	node := "node_example" // string | Show results filtered by node (optional)
	module := "module_example" // string | Show results filtered by module (optional)
	status := "status_example" // string | Filter by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ApiControllersTaskControllerGetTasksStatus(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Q(q).Search(search).Select_(select_).Sort(sort).AgentsList(agentsList).TasksList(tasksList).Command(command).Node(node).Module(module).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ApiControllersTaskControllerGetTasksStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersTaskControllerGetTasksStatus`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ApiControllersTaskControllerGetTasksStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersTaskControllerGetTasksStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **tasksList** | **[]string** | List of task IDs (separated by comma) |
 **command** | **string** | Filter by command |
 **node** | **string** | Show results filtered by node |
 **module** | **string** | Show results filtered by module |
 **status** | **string** | Filter by status |

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
