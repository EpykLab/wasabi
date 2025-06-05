# \MITREAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersMitreControllerGetGroups**](MITREAPI.md#ApiControllersMitreControllerGetGroups) | **Get** /mitre/groups | Get MITRE groups
[**ApiControllersMitreControllerGetMetadata**](MITREAPI.md#ApiControllersMitreControllerGetMetadata) | **Get** /mitre/metadata | Get MITRE metadata
[**ApiControllersMitreControllerGetMitigations**](MITREAPI.md#ApiControllersMitreControllerGetMitigations) | **Get** /mitre/mitigations | Get MITRE mitigations
[**ApiControllersMitreControllerGetReferences**](MITREAPI.md#ApiControllersMitreControllerGetReferences) | **Get** /mitre/references | Get MITRE references
[**ApiControllersMitreControllerGetSoftware**](MITREAPI.md#ApiControllersMitreControllerGetSoftware) | **Get** /mitre/software | Get MITRE software
[**ApiControllersMitreControllerGetTactics**](MITREAPI.md#ApiControllersMitreControllerGetTactics) | **Get** /mitre/tactics | Get MITRE tactics
[**ApiControllersMitreControllerGetTechniques**](MITREAPI.md#ApiControllersMitreControllerGetTechniques) | **Get** /mitre/techniques | Get MITRE techniques



## ApiControllersMitreControllerGetGroups

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetGroups(ctx).GroupIds(groupIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get MITRE groups



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
	groupIds := []string{"Inner_example"} // []string | List of MITRE's group IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetGroups(context.Background()).GroupIds(groupIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetGroups`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupIds** | **[]string** | List of MITRE&#39;s group IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersMitreControllerGetMetadata

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetMetadata(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get MITRE metadata



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
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetMetadata(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetMetadata`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetMetadataRequest struct via the builder pattern


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


## ApiControllersMitreControllerGetMitigations

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetMitigations(ctx).MitigationIds(mitigationIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get MITRE mitigations



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
	mitigationIds := []string{"Inner_example"} // []string | List of MITRE's mitigations IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetMitigations(context.Background()).MitigationIds(mitigationIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetMitigations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetMitigations`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetMitigations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetMitigationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mitigationIds** | **[]string** | List of MITRE&#39;s mitigations IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersMitreControllerGetReferences

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetReferences(ctx).ReferenceIds(referenceIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Execute()

Get MITRE references



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
	referenceIds := []string{"Inner_example"} // []string | List of MITRE's references IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetReferences(context.Background()).ReferenceIds(referenceIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetReferences`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetReferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceIds** | **[]string** | List of MITRE&#39;s references IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

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


## ApiControllersMitreControllerGetSoftware

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetSoftware(ctx).SoftwareIds(softwareIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get MITRE software



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
	softwareIds := []string{"Inner_example"} // []string | List of MITRE's software IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetSoftware(context.Background()).SoftwareIds(softwareIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetSoftware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetSoftware`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetSoftware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetSoftwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **softwareIds** | **[]string** | List of MITRE&#39;s software IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersMitreControllerGetTactics

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetTactics(ctx).TacticIds(tacticIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get MITRE tactics



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
	tacticIds := []string{"Inner_example"} // []string | List of MITRE's tactics IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetTactics(context.Background()).TacticIds(tacticIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetTactics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetTactics`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetTactics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetTacticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tacticIds** | **[]string** | List of MITRE&#39;s tactics IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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


## ApiControllersMitreControllerGetTechniques

> ApiControllersMitreControllerGetGroups200Response ApiControllersMitreControllerGetTechniques(ctx).TechniqueIds(techniqueIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()

Get MITRE techniques



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
	techniqueIds := []string{"Inner_example"} // []string | List of MITRE's techniques IDs (separated by comma) (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MITREAPI.ApiControllersMitreControllerGetTechniques(context.Background()).TechniqueIds(techniqueIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MITREAPI.ApiControllersMitreControllerGetTechniques``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersMitreControllerGetTechniques`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `MITREAPI.ApiControllersMitreControllerGetTechniques`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersMitreControllerGetTechniquesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **techniqueIds** | **[]string** | List of MITRE&#39;s techniques IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

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
