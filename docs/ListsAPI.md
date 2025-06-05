# \ListsAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersCdbListControllerDeleteFile**](ListsAPI.md#ApiControllersCdbListControllerDeleteFile) | **Delete** /lists/files/{filename} | Delete CDB list file
[**ApiControllersCdbListControllerGetFile**](ListsAPI.md#ApiControllersCdbListControllerGetFile) | **Get** /lists/files/{filename} | Get CDB list file content
[**ApiControllersCdbListControllerGetLists**](ListsAPI.md#ApiControllersCdbListControllerGetLists) | **Get** /lists | Get CDB lists info
[**ApiControllersCdbListControllerGetListsFiles**](ListsAPI.md#ApiControllersCdbListControllerGetListsFiles) | **Get** /lists/files | Get CDB lists files
[**ApiControllersCdbListControllerPutFile**](ListsAPI.md#ApiControllersCdbListControllerPutFile) | **Put** /lists/files/{filename} | Update CDB list file



## ApiControllersCdbListControllerDeleteFile

> ApiControllersClusterControllerUpdateConfiguration200Response ApiControllersCdbListControllerDeleteFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete CDB list file



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
	filename := "filename_example" // string | Filename (CDB list) to get/edit/delete.
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ApiControllersCdbListControllerDeleteFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ApiControllersCdbListControllerDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCdbListControllerDeleteFile`: ApiControllersClusterControllerUpdateConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ApiControllersCdbListControllerDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (CDB list) to get/edit/delete. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCdbListControllerDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersClusterControllerUpdateConfiguration200Response**](ApiControllersClusterControllerUpdateConfiguration200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersCdbListControllerGetFile

> ApiControllersCdbListControllerGetFile200Response ApiControllersCdbListControllerGetFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Execute()

Get CDB list file content



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
	filename := "filename_example" // string | Filename (CDB list) to get/edit/delete.
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	raw := true // bool | Format response in plain text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ApiControllersCdbListControllerGetFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ApiControllersCdbListControllerGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCdbListControllerGetFile`: ApiControllersCdbListControllerGetFile200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ApiControllersCdbListControllerGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (CDB list) to get/edit/delete. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCdbListControllerGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **raw** | **bool** | Format response in plain text |

### Return type

[**ApiControllersCdbListControllerGetFile200Response**](ApiControllersCdbListControllerGetFile200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersCdbListControllerGetLists

> ApiControllersCdbListControllerGetLists200Response ApiControllersCdbListControllerGetLists(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Q(q).Distinct(distinct).Execute()

Get CDB lists info



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
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	filename := []string{"Inner_example"} // []string | Filter by filename (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ApiControllersCdbListControllerGetLists(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ApiControllersCdbListControllerGetLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCdbListControllerGetLists`: ApiControllersCdbListControllerGetLists200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ApiControllersCdbListControllerGetLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCdbListControllerGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **relativeDirname** | **string** | Filter by relative directory name |
 **filename** | **[]string** | Filter by filename |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersCdbListControllerGetLists200Response**](ApiControllersCdbListControllerGetLists200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersCdbListControllerGetListsFiles

> ApiControllersCdbListControllerGetFile200Response ApiControllersCdbListControllerGetListsFiles(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Execute()

Get CDB lists files



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
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	filename := []string{"Inner_example"} // []string | Filter by filename (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ApiControllersCdbListControllerGetListsFiles(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ApiControllersCdbListControllerGetListsFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCdbListControllerGetListsFiles`: ApiControllersCdbListControllerGetFile200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ApiControllersCdbListControllerGetListsFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCdbListControllerGetListsFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **relativeDirname** | **string** | Filter by relative directory name |
 **filename** | **[]string** | Filter by filename |

### Return type

[**ApiControllersCdbListControllerGetFile200Response**](ApiControllersCdbListControllerGetFile200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersCdbListControllerPutFile

> ApiControllersClusterControllerUpdateConfiguration200Response ApiControllersCdbListControllerPutFile(ctx, filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).Execute()

Update CDB list file



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
	filename := "filename_example" // string | Filename (CDB list) to get/edit/delete.
	body := os.NewFile(1234, "some_file") // *os.File | Content of the file to be uploaded
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	overwrite := true // bool | If set to false, an exception will be raised when updating contents of an already existing filename (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ApiControllersCdbListControllerPutFile(context.Background(), filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ApiControllersCdbListControllerPutFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCdbListControllerPutFile`: ApiControllersClusterControllerUpdateConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ApiControllersCdbListControllerPutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (CDB list) to get/edit/delete. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCdbListControllerPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Content of the file to be uploaded |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **overwrite** | **bool** | If set to false, an exception will be raised when updating contents of an already existing filename | [default to false]

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
