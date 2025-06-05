# \ManagerAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersManagerControllerCheckAvailableVersion**](ManagerAPI.md#ApiControllersManagerControllerCheckAvailableVersion) | **Get** /manager/version/check | Check available updates
[**ApiControllersManagerControllerGetApiConfig**](ManagerAPI.md#ApiControllersManagerControllerGetApiConfig) | **Get** /manager/api/config | Get API config
[**ApiControllersManagerControllerGetConfValidation**](ManagerAPI.md#ApiControllersManagerControllerGetConfValidation) | **Get** /manager/configuration/validation | Check config
[**ApiControllersManagerControllerGetConfiguration**](ManagerAPI.md#ApiControllersManagerControllerGetConfiguration) | **Get** /manager/configuration | Get configuration
[**ApiControllersManagerControllerGetDaemonStats**](ManagerAPI.md#ApiControllersManagerControllerGetDaemonStats) | **Get** /manager/daemons/stats | Get Wazuh daemon stats
[**ApiControllersManagerControllerGetInfo**](ManagerAPI.md#ApiControllersManagerControllerGetInfo) | **Get** /manager/info | Get information
[**ApiControllersManagerControllerGetLog**](ManagerAPI.md#ApiControllersManagerControllerGetLog) | **Get** /manager/logs | Get logs
[**ApiControllersManagerControllerGetLogSummary**](ManagerAPI.md#ApiControllersManagerControllerGetLogSummary) | **Get** /manager/logs/summary | Get logs summary
[**ApiControllersManagerControllerGetManagerConfigOndemand**](ManagerAPI.md#ApiControllersManagerControllerGetManagerConfigOndemand) | **Get** /manager/configuration/{component}/{configuration} | Get active configuration
[**ApiControllersManagerControllerGetStats**](ManagerAPI.md#ApiControllersManagerControllerGetStats) | **Get** /manager/stats | Get stats
[**ApiControllersManagerControllerGetStatsAnalysisd**](ManagerAPI.md#ApiControllersManagerControllerGetStatsAnalysisd) | **Get** /manager/stats/analysisd | Get stats analysisd
[**ApiControllersManagerControllerGetStatsHourly**](ManagerAPI.md#ApiControllersManagerControllerGetStatsHourly) | **Get** /manager/stats/hourly | Get stats hour
[**ApiControllersManagerControllerGetStatsRemoted**](ManagerAPI.md#ApiControllersManagerControllerGetStatsRemoted) | **Get** /manager/stats/remoted | Get stats remoted
[**ApiControllersManagerControllerGetStatsWeekly**](ManagerAPI.md#ApiControllersManagerControllerGetStatsWeekly) | **Get** /manager/stats/weekly | Get stats week
[**ApiControllersManagerControllerGetStatus**](ManagerAPI.md#ApiControllersManagerControllerGetStatus) | **Get** /manager/status | Get status
[**ApiControllersManagerControllerPutRestart**](ManagerAPI.md#ApiControllersManagerControllerPutRestart) | **Put** /manager/restart | Restart manager
[**ApiControllersManagerControllerUpdateConfiguration**](ManagerAPI.md#ApiControllersManagerControllerUpdateConfiguration) | **Put** /manager/configuration | Update Wazuh configuration



## ApiControllersManagerControllerCheckAvailableVersion

> ApiControllersManagerControllerCheckAvailableVersion200Response ApiControllersManagerControllerCheckAvailableVersion(ctx).Pretty(pretty).ForceQuery(forceQuery).Execute()

Check available updates



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
	forceQuery := true // bool | Force query to CTI service (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerCheckAvailableVersion(context.Background()).Pretty(pretty).ForceQuery(forceQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerCheckAvailableVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerCheckAvailableVersion`: ApiControllersManagerControllerCheckAvailableVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerCheckAvailableVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerCheckAvailableVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **forceQuery** | **bool** | Force query to CTI service | [default to false]

### Return type

[**ApiControllersManagerControllerCheckAvailableVersion200Response**](ApiControllersManagerControllerCheckAvailableVersion200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersManagerControllerGetApiConfig

> ApiResponse ApiControllersManagerControllerGetApiConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get API config



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetApiConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetApiConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetApiConfig`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetApiConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetApiConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

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


## ApiControllersManagerControllerGetConfValidation

> ConfigurationValidation ApiControllersManagerControllerGetConfValidation(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Check config



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetConfValidation(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetConfValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetConfValidation`: ConfigurationValidation
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetConfValidation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetConfValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ConfigurationValidation**](ConfigurationValidation.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersManagerControllerGetConfiguration

> ApiControllersClusterControllerGetConfigurationNode200Response ApiControllersManagerControllerGetConfiguration(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Section(section).Field(field).Distinct(distinct).Execute()

Get configuration



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
	raw := true // bool | Format response in plain text (optional)
	section := "section_example" // string | Indicates the wazuh configuration section (optional)
	field := "field_example" // string | Indicate a section child. E.g, fields for *ruleset* section are: decoder_dir, rule_dir, etc (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetConfiguration(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Raw(raw).Section(section).Field(field).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetConfiguration`: ApiControllersClusterControllerGetConfigurationNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **raw** | **bool** | Format response in plain text |
 **section** | **string** | Indicates the wazuh configuration section |
 **field** | **string** | Indicate a section child. E.g, fields for *ruleset* section are: decoder_dir, rule_dir, etc |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersClusterControllerGetConfigurationNode200Response**](ApiControllersClusterControllerGetConfigurationNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersManagerControllerGetDaemonStats

> ApiControllersClusterControllerGetDaemonStatsNode200Response ApiControllersManagerControllerGetDaemonStats(ctx).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()

Get Wazuh daemon stats



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
	daemonsList := []string{"DaemonsList_example"} // []string | List of daemon names (separated by comma), all daemons selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetDaemonStats(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetDaemonStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetDaemonStats`: ApiControllersClusterControllerGetDaemonStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetDaemonStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetDaemonStatsRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetInfo

> ApiControllersClusterControllerGetInfoNode200Response ApiControllersManagerControllerGetInfo(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get information



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetInfo(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetInfo`: ApiControllersClusterControllerGetInfoNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetInfoRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetLog

> ApiControllersClusterControllerGetLogNode200Response ApiControllersManagerControllerGetLog(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Tag(tag).Level(level).Q(q).Select_(select_).Distinct(distinct).Execute()

Get logs



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetLog(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Tag(tag).Level(level).Q(q).Select_(select_).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetLog`: ApiControllersClusterControllerGetLogNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetLogRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetLogSummary

> ApiControllersClusterControllerGetLogSummaryNode200Response ApiControllersManagerControllerGetLogSummary(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get logs summary



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetLogSummary(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetLogSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetLogSummary`: ApiControllersClusterControllerGetLogSummaryNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetLogSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetLogSummaryRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetManagerConfigOndemand

> ApiResponse ApiControllersManagerControllerGetManagerConfigOndemand(ctx, component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get active configuration



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
	component := "component_example" // string | Selected agent's component
	configuration := "configuration_example" // string | <p>Selected agent's configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:</p> <table class=\"table table-striped table-bordered\"> <thead> <tr> <th>Component</th> <th>Configuration</th> <th>Tag</th> </tr> </thead> <tbody> <tr> <td>agent</td> <td>client</td> <td><code>&lt;client&gt;</code></td> </tr> <tr> <td>agent</td> <td>buffer</td> <td><code>&lt;client_buffer&gt;</code></td> </tr> <tr> <td>agent</td> <td>labels</td> <td><code>&lt;labels&gt;</code></td> </tr> <tr> <td>agent</td> <td>internal</td> <td><code>&lt;agent&gt;</code>, <code>&lt;monitord&gt;</code>, <code>&lt;remoted&gt;</code></td> </tr> <tr> <td>agent</td> <td>anti_tampering</td> <td><code>&lt;anti_tampering&gt;</code></td> </tr> <tr> <td>agentless</td> <td>agentless</td> <td><code>&lt;agentless&gt;</code></td> </tr> <tr> <td>analysis</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>analysis</td> <td>active_response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>analysis</td> <td>alerts</td> <td><code>&lt;alerts&gt;</code></td> </tr> <tr> <td>analysis</td> <td>command</td> <td><code>&lt;command&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rules</td> <td><code>&lt;rule&gt;</code></td> </tr> <tr> <td>analysis</td> <td>decoders</td> <td><code>&lt;decoder&gt;</code></td> </tr> <tr> <td>analysis</td> <td>internal</td> <td><code>&lt;analysisd&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rule_test</td> <td><code>&lt;rule_test&gt;</code></td> </tr> <tr> <td>auth</td> <td>auth</td> <td><code>&lt;auth&gt;</code></td> </tr> <tr> <td>com</td> <td>active-response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>com</td> <td>logging</td> <td><code>&lt;logging&gt;</code></td> </tr> <tr> <td>com</td> <td>internal</td> <td><code>&lt;execd&gt;</code></td> </tr> <tr> <td>com</td> <td>cluster</td> <td><code>&lt;cluster&gt;</code></td> </tr> <tr> <td>csyslog</td> <td>csyslog</td> <td><code>&lt;csyslog_output&gt;</code></td> </tr> <tr> <td>integrator</td> <td>integration</td> <td><code>&lt;integration&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>localfile</td> <td><code>&lt;localfile&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>socket</td> <td><code>&lt;socket&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>internal</td> <td><code>&lt;logcollector&gt;</code></td> </tr> <tr> <td>mail</td> <td>global</td> <td><code>&lt;global&gt;&lt;email...&gt;</code></td> </tr> <tr> <td>mail</td> <td>alerts</td> <td><code>&lt;email_alerts&gt;</code></td> </tr> <tr> <td>mail</td> <td>internal</td> <td><code>&lt;maild&gt;</code></td> </tr> <tr> <td>monitor</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>monitor</td> <td>internal</td> <td><code>&lt;monitord&gt;</code></td> </tr> <tr> <td>monitor</td> <td>reports</td> <td><code>&lt;reports&gt;</code></td> </tr> <tr> <td>request</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>request</td> <td>remote</td> <td><code>&lt;remote&gt;</code></td> </tr> <tr> <td>request</td> <td>internal</td> <td><code>&lt;remoted&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>syscheck</td> <td><code>&lt;syscheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>rootcheck</td> <td><code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>internal</td> <td><code>&lt;syscheck&gt;</code>, <code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>internal</td> <td><code>&lt;wazuh_db&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>wdb</td> <td><code>&lt;wdb&gt;</code></td> </tr> <tr> <td>wmodules</td> <td>wmodules</td> <td><code>&lt;wodle&gt;</code></td> </tr> </tbody> </table>
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetManagerConfigOndemand(context.Background(), component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetManagerConfigOndemand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetManagerConfigOndemand`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetManagerConfigOndemand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**component** | **string** | Selected agent&#39;s component |
**configuration** | **string** | &lt;p&gt;Selected agent&#39;s configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:&lt;/p&gt; &lt;table class&#x3D;\&quot;table table-striped table-bordered\&quot;&gt; &lt;thead&gt; &lt;tr&gt; &lt;th&gt;Component&lt;/th&gt; &lt;th&gt;Configuration&lt;/th&gt; &lt;th&gt;Tag&lt;/th&gt; &lt;/tr&gt; &lt;/thead&gt; &lt;tbody&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;client&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;buffer&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client_buffer&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;labels&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;labels&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agent&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;anti_tampering&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;anti_tampering&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agentless&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;active_response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;command&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;command&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;decoders&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;decoder&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;analysisd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rule_test&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule_test&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;auth&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;active-response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;logging&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logging&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;execd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;cluster&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;cluster&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;csyslog_output&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;integrator&lt;/td&gt; &lt;td&gt;integration&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;integration&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;localfile&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;localfile&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;socket&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;socket&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logcollector&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&amp;lt;email...&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;email_alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;maild&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;reports&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;reports&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;remote&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remote&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;rootcheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wazuh_db&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;wdb&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wdb&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wodle&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;/tbody&gt; &lt;/table&gt;  |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetManagerConfigOndemandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

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


## ApiControllersManagerControllerGetStats

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersManagerControllerGetStats(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Date(date).Execute()

Get stats



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
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	date := time.Now() // string | Date to obtain statistical information from. Format YYYY-MM-DD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStats(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStats`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatsRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetStatsAnalysisd

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersManagerControllerGetStatsAnalysisd(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get stats analysisd



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStatsAnalysisd(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStatsAnalysisd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStatsAnalysisd`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStatsAnalysisd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatsAnalysisdRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetStatsHourly

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersManagerControllerGetStatsHourly(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get stats hour



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStatsHourly(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStatsHourly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStatsHourly`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStatsHourly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatsHourlyRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetStatsRemoted

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersManagerControllerGetStatsRemoted(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get stats remoted



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStatsRemoted(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStatsRemoted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStatsRemoted`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStatsRemoted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatsRemotedRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetStatsWeekly

> ApiControllersClusterControllerGetStatsNode200Response ApiControllersManagerControllerGetStatsWeekly(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get stats week



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStatsWeekly(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStatsWeekly``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStatsWeekly`: ApiControllersClusterControllerGetStatsNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStatsWeekly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatsWeeklyRequest struct via the builder pattern


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


## ApiControllersManagerControllerGetStatus

> ApiControllersClusterControllerGetStatusNode200Response ApiControllersManagerControllerGetStatus(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get status



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerGetStatus(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerGetStatus`: ApiControllersClusterControllerGetStatusNode200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerGetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerGetStatusRequest struct via the builder pattern


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


## ApiControllersManagerControllerPutRestart

> ApiResponse ApiControllersManagerControllerPutRestart(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Restart manager



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
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerPutRestart(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerPutRestart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerPutRestart`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerPutRestart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerPutRestartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

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


## ApiControllersManagerControllerUpdateConfiguration

> ApiControllersClusterControllerUpdateConfiguration200Response ApiControllersManagerControllerUpdateConfiguration(ctx).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Update Wazuh configuration



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
	body := os.NewFile(1234, "some_file") // *os.File | Content of the ossec.conf to be uploaded
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagerAPI.ApiControllersManagerControllerUpdateConfiguration(context.Background()).Body(body).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagerAPI.ApiControllersManagerControllerUpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersManagerControllerUpdateConfiguration`: ApiControllersClusterControllerUpdateConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagerAPI.ApiControllersManagerControllerUpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersManagerControllerUpdateConfigurationRequest struct via the builder pattern


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
