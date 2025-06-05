# \RulesAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersRuleControllerDeleteFile**](RulesAPI.md#ApiControllersRuleControllerDeleteFile) | **Delete** /rules/files/{filename} | Delete rules file
[**ApiControllersRuleControllerGetFile**](RulesAPI.md#ApiControllersRuleControllerGetFile) | **Get** /rules/files/{filename} | Get rules file content
[**ApiControllersRuleControllerGetRules**](RulesAPI.md#ApiControllersRuleControllerGetRules) | **Get** /rules | List rules
[**ApiControllersRuleControllerGetRulesFiles**](RulesAPI.md#ApiControllersRuleControllerGetRulesFiles) | **Get** /rules/files | Get files
[**ApiControllersRuleControllerGetRulesGroups**](RulesAPI.md#ApiControllersRuleControllerGetRulesGroups) | **Get** /rules/groups | Get groups
[**ApiControllersRuleControllerGetRulesRequirement**](RulesAPI.md#ApiControllersRuleControllerGetRulesRequirement) | **Get** /rules/requirement/{requirement} | Get requirements
[**ApiControllersRuleControllerPutFile**](RulesAPI.md#ApiControllersRuleControllerPutFile) | **Put** /rules/files/{filename} | Update rules file



## ApiControllersRuleControllerDeleteFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersRuleControllerDeleteFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).RelativeDirname(relativeDirname).Execute()

Delete rules file



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
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerDeleteFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerDeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerDeleteFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerDeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerDeleteFileRequest struct via the builder pattern


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


## ApiControllersRuleControllerGetFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersRuleControllerGetFile(ctx, filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).RelativeDirname(relativeDirname).Execute()

Get rules file content



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
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerGetFile(context.Background(), filename).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerGetFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerGetFileRequest struct via the builder pattern


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


## ApiControllersRuleControllerGetRules

> ApiControllersRuleControllerGetRules200Response ApiControllersRuleControllerGetRules(ctx).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Status(status).Group(group).Level(level).Filename(filename).RelativeDirname(relativeDirname).PciDss(pciDss).Gdpr(gdpr).Gpg13(gpg13).Hipaa(hipaa).Nist80053(nist80053).Tsc(tsc).Mitre(mitre).Distinct(distinct).Execute()

List rules



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
	ruleIds := []int32{int32(123)} // []int32 | List of rule IDs (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	status := "status_example" // string | Filter by list status. Use commas to enter multiple statuses (optional)
	group := "group_example" // string | Filter by rule group (optional)
	level := "level_example" // string | Filter by rule level. Can be a single level (4) or an interval (2-4) (optional)
	filename := []string{"Inner_example"} // []string | Filter by filename (optional)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	pciDss := "pciDss_example" // string | Filter by PCI_DSS requirement name (optional)
	gdpr := "gdpr_example" // string | Filter by GDPR requirement (optional)
	gpg13 := "gpg13_example" // string | Filter by GPG13 requirement (optional)
	hipaa := "hipaa_example" // string | Filter by HIPAA requirement (optional)
	nist80053 := "nist80053_example" // string | Filter by NIST-800-53 requirement (optional)
	tsc := "tsc_example" // string | Filters by TSC requirement (optional)
	mitre := "mitre_example" // string | Filters by MITRE technique ID (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerGetRules(context.Background()).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Status(status).Group(group).Level(level).Filename(filename).RelativeDirname(relativeDirname).PciDss(pciDss).Gdpr(gdpr).Gpg13(gpg13).Hipaa(hipaa).Nist80053(nist80053).Tsc(tsc).Mitre(mitre).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerGetRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerGetRules`: ApiControllersRuleControllerGetRules200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerGetRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleIds** | **[]int32** | List of rule IDs |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **status** | **string** | Filter by list status. Use commas to enter multiple statuses |
 **group** | **string** | Filter by rule group |
 **level** | **string** | Filter by rule level. Can be a single level (4) or an interval (2-4) |
 **filename** | **[]string** | Filter by filename |
 **relativeDirname** | **string** | Filter by relative directory name |
 **pciDss** | **string** | Filter by PCI_DSS requirement name |
 **gdpr** | **string** | Filter by GDPR requirement |
 **gpg13** | **string** | Filter by GPG13 requirement |
 **hipaa** | **string** | Filter by HIPAA requirement |
 **nist80053** | **string** | Filter by NIST-800-53 requirement |
 **tsc** | **string** | Filters by TSC requirement |
 **mitre** | **string** | Filters by MITRE technique ID |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersRuleControllerGetRules200Response**](ApiControllersRuleControllerGetRules200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersRuleControllerGetRulesFiles

> ApiControllersRuleControllerGetRulesFiles200Response ApiControllersRuleControllerGetRulesFiles(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Status(status).Q(q).Select_(select_).Distinct(distinct).Execute()

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
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)
	filename := []string{"Inner_example"} // []string | Filter by filename of one or more rule or decoder files. (optional)
	status := "status_example" // string | Filter by list status. Use commas to enter multiple statuses (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerGetRulesFiles(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).RelativeDirname(relativeDirname).Filename(filename).Status(status).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerGetRulesFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerGetRulesFiles`: ApiControllersRuleControllerGetRulesFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerGetRulesFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerGetRulesFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **relativeDirname** | **string** | Filter by relative directory name |
 **filename** | **[]string** | Filter by filename of one or more rule or decoder files. |
 **status** | **string** | Filter by list status. Use commas to enter multiple statuses |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersRuleControllerGetRulesFiles200Response**](ApiControllersRuleControllerGetRulesFiles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersRuleControllerGetRulesGroups

> ApiControllersMitreControllerGetGroups200Response ApiControllersRuleControllerGetRulesGroups(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Execute()

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
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerGetRulesGroups(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerGetRulesGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerGetRulesGroups`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerGetRulesGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerGetRulesGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
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


## ApiControllersRuleControllerGetRulesRequirement

> ApiControllersMitreControllerGetGroups200Response ApiControllersRuleControllerGetRulesRequirement(ctx, requirement).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Execute()

Get requirements



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
	requirement := "requirement_example" // string |
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerGetRulesRequirement(context.Background(), requirement).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerGetRulesRequirement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerGetRulesRequirement`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerGetRulesRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirement** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerGetRulesRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
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


## ApiControllersRuleControllerPutFile

> ApiControllersMitreControllerGetGroups200Response ApiControllersRuleControllerPutFile(ctx, filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).RelativeDirname(relativeDirname).Execute()

Update rules file



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
	body := os.NewFile(1234, "some_file") // *os.File | Content of the rule to be uploaded
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	overwrite := true // bool | If set to false, an exception will be raised when updating contents of an already existing filename (optional) (default to false)
	relativeDirname := "relativeDirname_example" // string | Filter by relative directory name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RulesAPI.ApiControllersRuleControllerPutFile(context.Background(), filename).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Overwrite(overwrite).RelativeDirname(relativeDirname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RulesAPI.ApiControllersRuleControllerPutFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersRuleControllerPutFile`: ApiControllersMitreControllerGetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `RulesAPI.ApiControllersRuleControllerPutFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename (rule or decoder) to download/upload/edit file. |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersRuleControllerPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Content of the rule to be uploaded |
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
