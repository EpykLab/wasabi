# \CiscatAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersCiscatControllerGetAgentsCiscatResults**](CiscatAPI.md#ApiControllersCiscatControllerGetAgentsCiscatResults) | **Get** /ciscat/{agent_id}/results | Get results



## ApiControllersCiscatControllerGetAgentsCiscatResults

> ApiControllersCiscatControllerGetAgentsCiscatResults200Response ApiControllersCiscatControllerGetAgentsCiscatResults(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Benchmark(benchmark).Profile(profile).Pass(pass).Fail(fail).Error_(error_).Notchecked(notchecked).Unknown(unknown).Score(score).Q(q).Execute()

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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CiscatAPI.ApiControllersCiscatControllerGetAgentsCiscatResults(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Benchmark(benchmark).Profile(profile).Pass(pass).Fail(fail).Error_(error_).Notchecked(notchecked).Unknown(unknown).Score(score).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CiscatAPI.ApiControllersCiscatControllerGetAgentsCiscatResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersCiscatControllerGetAgentsCiscatResults`: ApiControllersCiscatControllerGetAgentsCiscatResults200Response
	fmt.Fprintf(os.Stdout, "Response from `CiscatAPI.ApiControllersCiscatControllerGetAgentsCiscatResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersCiscatControllerGetAgentsCiscatResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
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
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

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
