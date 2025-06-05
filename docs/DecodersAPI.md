# \DecodersAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersDecoderControllerDeleteFile**](DecodersAPI.md#ApiControllersDecoderControllerDeleteFile) | **Delete** /decoders/files/{filename} | Delete decoders file
[**ApiControllersDecoderControllerGetDecoders**](DecodersAPI.md#ApiControllersDecoderControllerGetDecoders) | **Get** /decoders | List decoders
[**ApiControllersDecoderControllerGetDecodersFiles**](DecodersAPI.md#ApiControllersDecoderControllerGetDecodersFiles) | **Get** /decoders/files | Get files
[**ApiControllersDecoderControllerGetDecodersParents**](DecodersAPI.md#ApiControllersDecoderControllerGetDecodersParents) | **Get** /decoders/parents | Get parent decoders
[**ApiControllersDecoderControllerGetFile**](DecodersAPI.md#ApiControllersDecoderControllerGetFile) | **Get** /decoders/files/{filename} | Get decoders file content
[**ApiControllersDecoderControllerPutFile**](DecodersAPI.md#ApiControllersDecoderControllerPutFile) | **Put** /decoders/files/{filename} | Update decoders file



## ApiControllersDecoderControllerDeleteFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersDecoderControllerDeleteFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).RelativeDirname(relativeDirname).Execute()

Delete decoders file



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
	filename := "filename_example" // string | Filename (rule or decoder) to download/upload/edit file.
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerDeleteFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerDeleteFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **relativeDirname** | **string** | Filter by relative directory name |

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


## ApiControllersDecoderControllerGetDecoders

> ApiControllersDecoderControllerGetDecoders200Response ApiControllersDecoderControllerGetDecoders(ctx).DecoderNames(decoderNames).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Filename(filename).RelativeDirname(relativeDirname).Status(status).Distinct(distinct).Execute()

List decoders



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
	decoderNames := []string{"Inner_example"} // []string | Decoder name (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	filename := []string{"Inner_example"} // []string | Filter by filename (optional)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	status := "status_example" // string | Filter by list status. Use commas to enter multiple statuses (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerGetDecoders(context.Background()).DecoderNames(decoderNames).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Filename(filename).RelativeDirname(relativeDirname).Status(status).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerGetDecoders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerGetDecoders`: ApiControllersDecoderControllerGetDecoders200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerGetDecoders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerGetDecodersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decoderNames** | **[]string** | Decoder name |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **filename** | **[]string** | Filter by filename |
 **relativeDirname** | **string** | Filter by relative directory name |
 **status** | **string** | Filter by list status. Use commas to enter multiple statuses |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersDecoderControllerGetDecoders200Response**](ApiControllersDecoderControllerGetDecoders200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersDecoderControllerGetDecodersFiles

> ApiControllersDecoderControllerGetDecodersFiles200Response ApiControllersDecoderControllerGetDecodersFiles(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Filename(filename).RelativeDirname(relativeDirname).Status(status).Q(q).Select_(select_).Distinct(distinct).Execute()

Get files



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
	filename := []string{"Inner_example"} // []string | Filter by filename of one or more rule or decoder files. (optional)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	status := "status_example" // string | Filter by list status. Use commas to enter multiple statuses (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerGetDecodersFiles(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Filename(filename).RelativeDirname(relativeDirname).Status(status).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerGetDecodersFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerGetDecodersFiles`: ApiControllersDecoderControllerGetDecodersFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerGetDecodersFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerGetDecodersFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **filename** | **[]string** | Filter by filename of one or more rule or decoder files. |
 **relativeDirname** | **string** | Filter by relative directory name |
 **status** | **string** | Filter by list status. Use commas to enter multiple statuses |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersDecoderControllerGetDecodersFiles200Response**](ApiControllersDecoderControllerGetDecodersFiles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersDecoderControllerGetDecodersParents

> ApiControllersMitreControllerGetGroups200Response ApiControllersDecoderControllerGetDecodersParents(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Execute()

Get parent decoders



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerGetDecodersParents(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerGetDecodersParents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerGetDecodersParents`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerGetDecodersParents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerGetDecodersParentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |

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


## ApiControllersDecoderControllerGetFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersDecoderControllerGetFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).RelativeDirname(relativeDirname).Execute()

Get decoders file content



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
	filename := "filename_example" // string | Filename (rule or decoder) to download/upload/edit file.
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	raw := true // bool | Format response in plain text (optional)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerGetFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerGetFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **raw** | **bool** | Format response in plain text |
 **relativeDirname** | **string** | Filter by relative directory name |

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersDecoderControllerPutFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersDecoderControllerPutFile(ctx, filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).RelativeDirname(relativeDirname).Execute()

Update decoders file



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
	filename := "filename_example" // string | Filename (rule or decoder) to download/upload/edit file.
	body := os.NewFile(1234, "some_file") // *os.File | Content of the decoder to be uploaded
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	overwrite := true // bool | If set to false, an exception will be raised when updating contents of an already existing filename (optional) (default to false)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecodersAPI.ApiControllersDecoderControllerPutFile(context.Background(), filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecodersAPI.ApiControllersDecoderControllerPutFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersDecoderControllerPutFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DecodersAPI.ApiControllersDecoderControllerPutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersDecoderControllerPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Content of the decoder to be uploaded |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **overwrite** | **bool** | If set to false, an exception will be raised when updating contents of an already existing filename | [default to false]
 **relativeDirname** | **string** | Filter by relative directory name |

### Return type

[**ApiControllersMitreControllerGetGroups200Response**](ApiControllersMitreControllerGetGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
