# \GroupsAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersAgentControllerDeleteGroups**](GroupsAPI.md#ApiControllersAgentControllerDeleteGroups) | **Delete** /groups | Delete groups
[**ApiControllersAgentControllerGetAgentsInGroup**](GroupsAPI.md#ApiControllersAgentControllerGetAgentsInGroup) | **Get** /groups/{group_id}/agents | Get agents in a group
[**ApiControllersAgentControllerGetGroupConfig**](GroupsAPI.md#ApiControllersAgentControllerGetGroupConfig) | **Get** /groups/{group_id}/configuration | Get group configuration
[**ApiControllersAgentControllerGetGroupFile**](GroupsAPI.md#ApiControllersAgentControllerGetGroupFile) | **Get** /groups/{group_id}/files/{file_name} | Get a file in group
[**ApiControllersAgentControllerGetGroupFiles**](GroupsAPI.md#ApiControllersAgentControllerGetGroupFiles) | **Get** /groups/{group_id}/files | Get group files
[**ApiControllersAgentControllerGetListGroup**](GroupsAPI.md#ApiControllersAgentControllerGetListGroup) | **Get** /groups | Get groups
[**ApiControllersAgentControllerPostGroup**](GroupsAPI.md#ApiControllersAgentControllerPostGroup) | **Post** /groups | Create a group
[**ApiControllersAgentControllerPutGroupConfig**](GroupsAPI.md#ApiControllersAgentControllerPutGroupConfig) | **Put** /groups/{group_id}/configuration | Update group configuration



## ApiControllersAgentControllerDeleteGroups

> ApiControllersAgentControllerDeleteGroups200Response ApiControllersAgentControllerDeleteGroups(ctx).GroupsList(groupsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete groups



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
	groupsList := []string{"Inner_example"} // []string | List of group IDs (separated by comma), use the keyword 'all' to select all groups
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerDeleteGroups(context.Background()).GroupsList(groupsList).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerDeleteGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerDeleteGroups`: ApiControllersAgentControllerDeleteGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerDeleteGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerDeleteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupsList** | **[]string** | List of group IDs (separated by comma), use the keyword &#39;all&#39; to select all groups |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerDeleteGroups200Response**](ApiControllersAgentControllerDeleteGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetAgentsInGroup

> ApiControllersAgentControllerGetAgents200Response ApiControllersAgentControllerGetAgentsInGroup(ctx, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Status(status).Q(q).Distinct(distinct).Execute()

Get agents in a group



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	status := []string{"Status_example"} // []string | Filter by agent status (use commas to enter multiple statuses) (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerGetAgentsInGroup(context.Background(), groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Status(status).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerGetAgentsInGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentsInGroup`: ApiControllersAgentControllerGetAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerGetAgentsInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentsInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **status** | **[]string** | Filter by agent status (use commas to enter multiple statuses) |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersAgentControllerGetAgents200Response**](ApiControllersAgentControllerGetAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetGroupConfig

> ApiControllersAgentControllerGetGroupConfig200Response ApiControllersAgentControllerGetGroupConfig(ctx, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Execute()

Get group configuration



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerGetGroupConfig(context.Background(), groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerGetGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetGroupConfig`: ApiControllersAgentControllerGetGroupConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerGetGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]

### Return type

[**ApiControllersAgentControllerGetGroupConfig200Response**](ApiControllersAgentControllerGetGroupConfig200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetGroupFile

> ApiControllersAgentControllerGetGroupFile200Response ApiControllersAgentControllerGetGroupFile(ctx, groupId, fileName).Pretty(pretty).WaitForComplete(waitForComplete).Type_(type_).Raw(raw).Execute()

Get a file in group



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	fileName := "fileName_example" // string | Filename
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	type_ := []string{"Type_example"} // []string | Type of file (optional)
	raw := true // bool | Format response in plain text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerGetGroupFile(context.Background(), groupId, fileName).Pretty(pretty).WaitForComplete(waitForComplete).Type_(type_).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerGetGroupFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetGroupFile`: ApiControllersAgentControllerGetGroupFile200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerGetGroupFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |
**fileName** | **string** | Filename |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetGroupFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **type_** | **[]string** | Type of file |
 **raw** | **bool** | Format response in plain text |

### Return type

[**ApiControllersAgentControllerGetGroupFile200Response**](ApiControllersAgentControllerGetGroupFile200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetGroupFiles

> ApiControllersAgentControllerGetGroupFiles200Response ApiControllersAgentControllerGetGroupFiles(ctx, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Hash(hash).Q(q).Select_(select_).Distinct(distinct).Execute()

Get group files



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	hash := "hash_example" // string | Select algorithm to generate the returned checksums (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerGetGroupFiles(context.Background(), groupId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Hash(hash).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerGetGroupFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetGroupFiles`: ApiControllersAgentControllerGetGroupFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerGetGroupFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetGroupFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **hash** | **string** | Select algorithm to generate the returned checksums |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersAgentControllerGetGroupFiles200Response**](ApiControllersAgentControllerGetGroupFiles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetListGroup

> ApiControllersAgentControllerGetListGroup200Response ApiControllersAgentControllerGetListGroup(ctx).Pretty(pretty).WaitForComplete(waitForComplete).GroupsList(groupsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Hash(hash).Q(q).Select_(select_).Distinct(distinct).Execute()

Get groups



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
	groupsList := []string{"Inner_example"} // []string | List of group IDs (separated by comma), all groups selected by default if not specified (optional)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	hash := "hash_example" // string | Select algorithm to generate the returned checksums (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerGetListGroup(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).GroupsList(groupsList).Offset(offset).Limit(limit).Sort(sort).Search(search).Hash(hash).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerGetListGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetListGroup`: ApiControllersAgentControllerGetListGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerGetListGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetListGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **groupsList** | **[]string** | List of group IDs (separated by comma), all groups selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **hash** | **string** | Select algorithm to generate the returned checksums |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersAgentControllerGetListGroup200Response**](ApiControllersAgentControllerGetListGroup200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerPostGroup

> ApiResponse ApiControllersAgentControllerPostGroup(ctx).Pretty(pretty).WaitForComplete(waitForComplete).CreateGroupBody(createGroupBody).Execute()

Create a group



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
	createGroupBody := *openapiclient.NewCreateGroupBody("GroupId_example") // CreateGroupBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerPostGroup(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).CreateGroupBody(createGroupBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerPostGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPostGroup`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerPostGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPostGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **createGroupBody** | [**CreateGroupBody**](CreateGroupBody.md) |  |

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


## ApiControllersAgentControllerPutGroupConfig

> ApiResponse ApiControllersAgentControllerPutGroupConfig(ctx, groupId).ApiControllersAgentControllerPutGroupConfigRequest(apiControllersAgentControllerPutGroupConfigRequest).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Update group configuration



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	apiControllersAgentControllerPutGroupConfigRequest := *openapiclient.NewApiControllersAgentControllerPutGroupConfigRequest() // ApiControllersAgentControllerPutGroupConfigRequest |
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.ApiControllersAgentControllerPutGroupConfig(context.Background(), groupId).ApiControllersAgentControllerPutGroupConfigRequest(apiControllersAgentControllerPutGroupConfigRequest).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ApiControllersAgentControllerPutGroupConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPutGroupConfig`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ApiControllersAgentControllerPutGroupConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPutGroupConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiControllersAgentControllerPutGroupConfigRequest** | [**ApiControllersAgentControllerPutGroupConfigRequest**](ApiControllersAgentControllerPutGroupConfigRequest.md) |  |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
