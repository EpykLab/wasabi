# \AgentsAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersAgentControllerAddAgent**](AgentsAPI.md#ApiControllersAgentControllerAddAgent) | **Post** /agents | Add agent
[**ApiControllersAgentControllerDeleteAgents**](AgentsAPI.md#ApiControllersAgentControllerDeleteAgents) | **Delete** /agents | Delete agents
[**ApiControllersAgentControllerDeleteMultipleAgentSingleGroup**](AgentsAPI.md#ApiControllersAgentControllerDeleteMultipleAgentSingleGroup) | **Delete** /agents/group | Remove agents from group
[**ApiControllersAgentControllerDeleteSingleAgentMultipleGroups**](AgentsAPI.md#ApiControllersAgentControllerDeleteSingleAgentMultipleGroups) | **Delete** /agents/{agent_id}/group | Remove agent from groups
[**ApiControllersAgentControllerDeleteSingleAgentSingleGroup**](AgentsAPI.md#ApiControllersAgentControllerDeleteSingleAgentSingleGroup) | **Delete** /agents/{agent_id}/group/{group_id} | Remove agent from group
[**ApiControllersAgentControllerGetAgentConfig**](AgentsAPI.md#ApiControllersAgentControllerGetAgentConfig) | **Get** /agents/{agent_id}/config/{component}/{configuration} | Get active configuration
[**ApiControllersAgentControllerGetAgentFields**](AgentsAPI.md#ApiControllersAgentControllerGetAgentFields) | **Get** /agents/stats/distinct | List agents distinct
[**ApiControllersAgentControllerGetAgentKey**](AgentsAPI.md#ApiControllersAgentControllerGetAgentKey) | **Get** /agents/{agent_id}/key | Get key
[**ApiControllersAgentControllerGetAgentNoGroup**](AgentsAPI.md#ApiControllersAgentControllerGetAgentNoGroup) | **Get** /agents/no_group | List agents without group
[**ApiControllersAgentControllerGetAgentOutdated**](AgentsAPI.md#ApiControllersAgentControllerGetAgentOutdated) | **Get** /agents/outdated | List outdated agents
[**ApiControllersAgentControllerGetAgentSummaryOs**](AgentsAPI.md#ApiControllersAgentControllerGetAgentSummaryOs) | **Get** /agents/summary/os | Summarize agents OS
[**ApiControllersAgentControllerGetAgentSummaryStatus**](AgentsAPI.md#ApiControllersAgentControllerGetAgentSummaryStatus) | **Get** /agents/summary/status | Summarize agents status
[**ApiControllersAgentControllerGetAgentUninstallPermission**](AgentsAPI.md#ApiControllersAgentControllerGetAgentUninstallPermission) | **Get** /agents/uninstall | Check user&#39;s permission to uninstall agents
[**ApiControllersAgentControllerGetAgentUpgrade**](AgentsAPI.md#ApiControllersAgentControllerGetAgentUpgrade) | **Get** /agents/upgrade_result | Get upgrade results
[**ApiControllersAgentControllerGetAgents**](AgentsAPI.md#ApiControllersAgentControllerGetAgents) | **Get** /agents | List agents
[**ApiControllersAgentControllerGetComponentStats**](AgentsAPI.md#ApiControllersAgentControllerGetComponentStats) | **Get** /agents/{agent_id}/stats/{component} | Get agent&#39;s component stats
[**ApiControllersAgentControllerGetDaemonStats**](AgentsAPI.md#ApiControllersAgentControllerGetDaemonStats) | **Get** /agents/{agent_id}/daemons/stats | Get Wazuh daemon stats from an agent
[**ApiControllersAgentControllerGetSyncAgent**](AgentsAPI.md#ApiControllersAgentControllerGetSyncAgent) | **Get** /agents/{agent_id}/group/is_sync | Get configuration sync status
[**ApiControllersAgentControllerInsertAgent**](AgentsAPI.md#ApiControllersAgentControllerInsertAgent) | **Post** /agents/insert | Add agent full
[**ApiControllersAgentControllerPostNewAgent**](AgentsAPI.md#ApiControllersAgentControllerPostNewAgent) | **Post** /agents/insert/quick | Add agent quick
[**ApiControllersAgentControllerPutAgentSingleGroup**](AgentsAPI.md#ApiControllersAgentControllerPutAgentSingleGroup) | **Put** /agents/{agent_id}/group/{group_id} | Assign agent to group
[**ApiControllersAgentControllerPutMultipleAgentSingleGroup**](AgentsAPI.md#ApiControllersAgentControllerPutMultipleAgentSingleGroup) | **Put** /agents/group | Assign agents to group
[**ApiControllersAgentControllerPutUpgradeAgents**](AgentsAPI.md#ApiControllersAgentControllerPutUpgradeAgents) | **Put** /agents/upgrade | Upgrade agents
[**ApiControllersAgentControllerPutUpgradeCustomAgents**](AgentsAPI.md#ApiControllersAgentControllerPutUpgradeCustomAgents) | **Put** /agents/upgrade_custom | Upgrade agents custom
[**ApiControllersAgentControllerReconnectAgents**](AgentsAPI.md#ApiControllersAgentControllerReconnectAgents) | **Put** /agents/reconnect | Force reconnect agents
[**ApiControllersAgentControllerRestartAgent**](AgentsAPI.md#ApiControllersAgentControllerRestartAgent) | **Put** /agents/{agent_id}/restart | Restart agent
[**ApiControllersAgentControllerRestartAgents**](AgentsAPI.md#ApiControllersAgentControllerRestartAgents) | **Put** /agents/restart | Restart agents
[**ApiControllersAgentControllerRestartAgentsByGroup**](AgentsAPI.md#ApiControllersAgentControllerRestartAgentsByGroup) | **Put** /agents/group/{group_id}/restart | Restart agents in group
[**ApiControllersAgentControllerRestartAgentsByNode**](AgentsAPI.md#ApiControllersAgentControllerRestartAgentsByNode) | **Put** /agents/node/{node_id}/restart | Restart agents in node



## ApiControllersAgentControllerAddAgent

> ApiControllersAgentControllerAddAgent200Response ApiControllersAgentControllerAddAgent(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentAddBody(agentAddBody).Execute()

Add agent



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
	agentAddBody := *openapiclient.NewAgentAddBody("Name_example") // AgentAddBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerAddAgent(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentAddBody(agentAddBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerAddAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerAddAgent`: ApiControllersAgentControllerAddAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerAddAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerAddAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentAddBody** | [**AgentAddBody**](AgentAddBody.md) |  |

### Return type

[**ApiControllersAgentControllerAddAgent200Response**](ApiControllersAgentControllerAddAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerDeleteAgents

> ApiControllersAgentControllerDeleteAgents200Response ApiControllersAgentControllerDeleteAgents(ctx).AgentsList(agentsList).Status(status).Pretty(pretty).WaitForComplete(waitForComplete).Purge(purge).OlderThan(olderThan).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()

Delete agents



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	status := []string{"Status_example"} // []string | Filter by agent status (use commas to enter multiple statuses)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	purge := true // bool | Permanently delete an agent from the key store (optional) (default to false)
	olderThan := "olderThan_example" // string | Consider only agents whose last keep alive is older than the specified time frame. For never_connected agents, register date is considered instead of last keep alive. For example, `7d`, `10s` and `10` are valid values. When no time unit is specified, seconds are assumed. Use 0s to select all agents (optional) (default to "7d")
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	osPlatform := "osPlatform_example" // string | Filter by OS platform (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	manager := "manager_example" // string | Filter by manager hostname where agents are connected to (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	group := "group_example" // string | Filter by group of agents (optional)
	nodeName := "nodeName_example" // string | Filter by node name (optional)
	name := "name_example" // string | Filter by name (optional)
	ip := "ip_example" // string | Filter by the IP used by the agent to communicate with the manager. If it's not available, it will have the same value as registerIP (optional)
	registerIP := "registerIP_example" // string | Filter by the IP used when registering the agent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerDeleteAgents(context.Background()).AgentsList(agentsList).Status(status).Pretty(pretty).WaitForComplete(waitForComplete).Purge(purge).OlderThan(olderThan).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerDeleteAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerDeleteAgents`: ApiControllersAgentControllerDeleteAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerDeleteAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerDeleteAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
 **status** | **[]string** | Filter by agent status (use commas to enter multiple statuses) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **purge** | **bool** | Permanently delete an agent from the key store | [default to false]
 **olderThan** | **string** | Consider only agents whose last keep alive is older than the specified time frame. For never_connected agents, register date is considered instead of last keep alive. For example, &#x60;7d&#x60;, &#x60;10s&#x60; and &#x60;10&#x60; are valid values. When no time unit is specified, seconds are assumed. Use 0s to select all agents | [default to &quot;7d&quot;]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **osPlatform** | **string** | Filter by OS platform |
 **osVersion** | **string** | Filter by OS version |
 **osName** | **string** | Filter by OS name |
 **manager** | **string** | Filter by manager hostname where agents are connected to |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **group** | **string** | Filter by group of agents |
 **nodeName** | **string** | Filter by node name |
 **name** | **string** | Filter by name |
 **ip** | **string** | Filter by the IP used by the agent to communicate with the manager. If it&#39;s not available, it will have the same value as registerIP |
 **registerIP** | **string** | Filter by the IP used when registering the agent |

### Return type

[**ApiControllersAgentControllerDeleteAgents200Response**](ApiControllersAgentControllerDeleteAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerDeleteMultipleAgentSingleGroup

> ApiControllersAgentControllerDeleteAgents200Response ApiControllersAgentControllerDeleteMultipleAgentSingleGroup(ctx).AgentsList(agentsList).GroupId(groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Remove agents from group



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerDeleteMultipleAgentSingleGroup(context.Background()).AgentsList(agentsList).GroupId(groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerDeleteMultipleAgentSingleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerDeleteMultipleAgentSingleGroup`: ApiControllersAgentControllerDeleteAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerDeleteMultipleAgentSingleGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerDeleteMultipleAgentSingleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
 **groupId** | **string** | Group ID. (Name of the group) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerDeleteAgents200Response**](ApiControllersAgentControllerDeleteAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerDeleteSingleAgentMultipleGroups

> ApiControllersAgentControllerDeleteSingleAgentMultipleGroups200Response ApiControllersAgentControllerDeleteSingleAgentMultipleGroups(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).GroupsList(groupsList).Execute()

Remove agent from groups



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
	groupsList := []string{"Inner_example"} // []string | List of group IDs (separated by comma), all groups selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentMultipleGroups(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).GroupsList(groupsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentMultipleGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerDeleteSingleAgentMultipleGroups`: ApiControllersAgentControllerDeleteSingleAgentMultipleGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentMultipleGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerDeleteSingleAgentMultipleGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **groupsList** | **[]string** | List of group IDs (separated by comma), all groups selected by default if not specified |

### Return type

[**ApiControllersAgentControllerDeleteSingleAgentMultipleGroups200Response**](ApiControllersAgentControllerDeleteSingleAgentMultipleGroups200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerDeleteSingleAgentSingleGroup

> ApiResponse ApiControllersAgentControllerDeleteSingleAgentSingleGroup(ctx, agentId, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Remove agent from group



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentSingleGroup(context.Background(), agentId, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentSingleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerDeleteSingleAgentSingleGroup`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerDeleteSingleAgentSingleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerDeleteSingleAgentSingleGroupRequest struct via the builder pattern


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


## ApiControllersAgentControllerGetAgentConfig

> ApiControllersAgentControllerGetAgentConfig200Response ApiControllersAgentControllerGetAgentConfig(ctx, agentId, component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

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
	agentId := "agentId_example" // string | Agent ID. All possible values from 000 onwards
	component := "component_example" // string | Selected agent's component
	configuration := "configuration_example" // string | <p>Selected agent's configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:</p> <table class=\"table table-striped table-bordered\"> <thead> <tr> <th>Component</th> <th>Configuration</th> <th>Tag</th> </tr> </thead> <tbody> <tr> <td>agent</td> <td>client</td> <td><code>&lt;client&gt;</code></td> </tr> <tr> <td>agent</td> <td>buffer</td> <td><code>&lt;client_buffer&gt;</code></td> </tr> <tr> <td>agent</td> <td>labels</td> <td><code>&lt;labels&gt;</code></td> </tr> <tr> <td>agent</td> <td>internal</td> <td><code>&lt;agent&gt;</code>, <code>&lt;monitord&gt;</code>, <code>&lt;remoted&gt;</code></td> </tr> <tr> <td>agent</td> <td>anti_tampering</td> <td><code>&lt;anti_tampering&gt;</code></td> </tr> <tr> <td>agentless</td> <td>agentless</td> <td><code>&lt;agentless&gt;</code></td> </tr> <tr> <td>analysis</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>analysis</td> <td>active_response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>analysis</td> <td>alerts</td> <td><code>&lt;alerts&gt;</code></td> </tr> <tr> <td>analysis</td> <td>command</td> <td><code>&lt;command&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rules</td> <td><code>&lt;rule&gt;</code></td> </tr> <tr> <td>analysis</td> <td>decoders</td> <td><code>&lt;decoder&gt;</code></td> </tr> <tr> <td>analysis</td> <td>internal</td> <td><code>&lt;analysisd&gt;</code></td> </tr> <tr> <td>analysis</td> <td>rule_test</td> <td><code>&lt;rule_test&gt;</code></td> </tr> <tr> <td>auth</td> <td>auth</td> <td><code>&lt;auth&gt;</code></td> </tr> <tr> <td>com</td> <td>active-response</td> <td><code>&lt;active-response&gt;</code></td> </tr> <tr> <td>com</td> <td>logging</td> <td><code>&lt;logging&gt;</code></td> </tr> <tr> <td>com</td> <td>internal</td> <td><code>&lt;execd&gt;</code></td> </tr> <tr> <td>com</td> <td>cluster</td> <td><code>&lt;cluster&gt;</code></td> </tr> <tr> <td>csyslog</td> <td>csyslog</td> <td><code>&lt;csyslog_output&gt;</code></td> </tr> <tr> <td>integrator</td> <td>integration</td> <td><code>&lt;integration&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>localfile</td> <td><code>&lt;localfile&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>socket</td> <td><code>&lt;socket&gt;</code></td> </tr> <tr> <td>logcollector</td> <td>internal</td> <td><code>&lt;logcollector&gt;</code></td> </tr> <tr> <td>mail</td> <td>global</td> <td><code>&lt;global&gt;&lt;email...&gt;</code></td> </tr> <tr> <td>mail</td> <td>alerts</td> <td><code>&lt;email_alerts&gt;</code></td> </tr> <tr> <td>mail</td> <td>internal</td> <td><code>&lt;maild&gt;</code></td> </tr> <tr> <td>monitor</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>monitor</td> <td>internal</td> <td><code>&lt;monitord&gt;</code></td> </tr> <tr> <td>monitor</td> <td>reports</td> <td><code>&lt;reports&gt;</code></td> </tr> <tr> <td>request</td> <td>global</td> <td><code>&lt;global&gt;</code></td> </tr> <tr> <td>request</td> <td>remote</td> <td><code>&lt;remote&gt;</code></td> </tr> <tr> <td>request</td> <td>internal</td> <td><code>&lt;remoted&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>syscheck</td> <td><code>&lt;syscheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>rootcheck</td> <td><code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>syscheck</td> <td>internal</td> <td><code>&lt;syscheck&gt;</code>, <code>&lt;rootcheck&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>internal</td> <td><code>&lt;wazuh_db&gt;</code></td> </tr> <tr> <td>wazuh-db</td> <td>wdb</td> <td><code>&lt;wdb&gt;</code></td> </tr> <tr> <td>wmodules</td> <td>wmodules</td> <td><code>&lt;wodle&gt;</code></td> </tr> </tbody> </table>
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentConfig(context.Background(), agentId, component, configuration).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentConfig`: ApiControllersAgentControllerGetAgentConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |
**component** | **string** | Selected agent&#39;s component |
**configuration** | **string** | &lt;p&gt;Selected agent&#39;s configuration to read. The configuration to read depends on the selected component. The following table shows all available combinations of component and configuration values:&lt;/p&gt; &lt;table class&#x3D;\&quot;table table-striped table-bordered\&quot;&gt; &lt;thead&gt; &lt;tr&gt; &lt;th&gt;Component&lt;/th&gt; &lt;th&gt;Configuration&lt;/th&gt; &lt;th&gt;Tag&lt;/th&gt; &lt;/tr&gt; &lt;/thead&gt; &lt;tbody&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;client&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;buffer&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;client_buffer&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;labels&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;labels&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agent&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agent&lt;/td&gt; &lt;td&gt;anti_tampering&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;anti_tampering&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;agentless&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;agentless&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;active_response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;command&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;command&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;decoders&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;decoder&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;analysisd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;analysis&lt;/td&gt; &lt;td&gt;rule_test&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rule_test&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;auth&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;auth&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;active-response&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;active-response&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;logging&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logging&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;execd&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;com&lt;/td&gt; &lt;td&gt;cluster&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;cluster&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;csyslog&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;csyslog_output&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;integrator&lt;/td&gt; &lt;td&gt;integration&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;integration&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;localfile&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;localfile&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;socket&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;socket&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;logcollector&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;logcollector&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&amp;lt;email...&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;alerts&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;email_alerts&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;mail&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;maild&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;monitord&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;monitor&lt;/td&gt; &lt;td&gt;reports&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;reports&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;global&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;global&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;remote&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remote&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;request&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;remoted&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;rootcheck&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;syscheck&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;syscheck&amp;gt;&lt;/code&gt;, &lt;code&gt;&amp;lt;rootcheck&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;internal&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wazuh_db&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wazuh-db&lt;/td&gt; &lt;td&gt;wdb&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wdb&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;wmodules&lt;/td&gt; &lt;td&gt;&lt;code&gt;&amp;lt;wodle&amp;gt;&lt;/code&gt;&lt;/td&gt; &lt;/tr&gt; &lt;/tbody&gt; &lt;/table&gt;  |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerGetAgentConfig200Response**](ApiControllersAgentControllerGetAgentConfig200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetAgentFields

> ApiControllersAgentControllerGetAgentFields200Response ApiControllersAgentControllerGetAgentFields(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Fields(fields).Offset(offset).Limit(limit).Sort(sort).Search(search).Q(q).Execute()

List agents distinct



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
	fields := []string{"Inner_example"} // []string | List of fields affecting the operation (optional)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentFields(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Fields(fields).Offset(offset).Limit(limit).Sort(sort).Search(search).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentFields`: ApiControllersAgentControllerGetAgentFields200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **fields** | **[]string** | List of fields affecting the operation |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

### Return type

[**ApiControllersAgentControllerGetAgentFields200Response**](ApiControllersAgentControllerGetAgentFields200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetAgentKey

> ApiControllersAgentControllerGetAgentKey200Response ApiControllersAgentControllerGetAgentKey(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentKey(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentKey`: ApiControllersAgentControllerGetAgentKey200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerGetAgentKey200Response**](ApiControllersAgentControllerGetAgentKey200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetAgentNoGroup

> ApiControllersAgentControllerGetAgents200Response ApiControllersAgentControllerGetAgentNoGroup(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Execute()

List agents without group



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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentNoGroup(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentNoGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentNoGroup`: ApiControllersAgentControllerGetAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentNoGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentNoGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

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


## ApiControllersAgentControllerGetAgentOutdated

> ApiControllersAgentControllerGetAgents200Response ApiControllersAgentControllerGetAgentOutdated(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Execute()

List outdated agents



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
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentOutdated(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Offset(offset).Limit(limit).Sort(sort).Search(search).Select_(select_).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentOutdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentOutdated`: ApiControllersAgentControllerGetAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentOutdated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentOutdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |

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


## ApiControllersAgentControllerGetAgentSummaryOs

> ApiResponse ApiControllersAgentControllerGetAgentSummaryOs(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Summarize agents OS



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
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentSummaryOs(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentSummaryOs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentSummaryOs`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentSummaryOs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentSummaryOsRequest struct via the builder pattern


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


## ApiControllersAgentControllerGetAgentSummaryStatus

> ApiControllersAgentControllerGetAgentSummaryStatus200Response ApiControllersAgentControllerGetAgentSummaryStatus(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Summarize agents status



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
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentSummaryStatus(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentSummaryStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentSummaryStatus`: ApiControllersAgentControllerGetAgentSummaryStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentSummaryStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentSummaryStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerGetAgentSummaryStatus200Response**](ApiControllersAgentControllerGetAgentSummaryStatus200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetAgentUninstallPermission

> ApiResponse ApiControllersAgentControllerGetAgentUninstallPermission(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Check user's permission to uninstall agents



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
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentUninstallPermission(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentUninstallPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentUninstallPermission`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentUninstallPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentUninstallPermissionRequest struct via the builder pattern


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


## ApiControllersAgentControllerGetAgentUpgrade

> ApiResponse ApiControllersAgentControllerGetAgentUpgrade(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()

Get upgrade results



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	osPlatform := "osPlatform_example" // string | Filter by OS platform (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	manager := "manager_example" // string | Filter by manager hostname where agents are connected to (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	group := "group_example" // string | Filter by group of agents (optional)
	nodeName := "nodeName_example" // string | Filter by node name (optional)
	name := "name_example" // string | Filter by name (optional)
	ip := "ip_example" // string | Filter by the IP used by the agent to communicate with the manager. If it's not available, it will have the same value as registerIP (optional)
	registerIP := "registerIP_example" // string | Filter by the IP used when registering the agent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgentUpgrade(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgentUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgentUpgrade`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgentUpgrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **osPlatform** | **string** | Filter by OS platform |
 **osVersion** | **string** | Filter by OS version |
 **osName** | **string** | Filter by OS name |
 **manager** | **string** | Filter by manager hostname where agents are connected to |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **group** | **string** | Filter by group of agents |
 **nodeName** | **string** | Filter by node name |
 **name** | **string** | Filter by name |
 **ip** | **string** | Filter by the IP used by the agent to communicate with the manager. If it&#39;s not available, it will have the same value as registerIP |
 **registerIP** | **string** | Filter by the IP used when registering the agent |

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


## ApiControllersAgentControllerGetAgents

> ApiControllersAgentControllerGetAgents200Response ApiControllersAgentControllerGetAgents(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Status(status).Q(q).OlderThan(olderThan).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).GroupConfigStatus(groupConfigStatus).Distinct(distinct).Execute()

List agents



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	status := []string{"Status_example"} // []string | Filter by agent status (use commas to enter multiple statuses) (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	olderThan := "olderThan_example" // string | Filter out agents whose time lapse from last keep alive signal is longer than specified. Time in seconds, ‘[n_days]d’, ‘[n_hours]h’, ‘[n_minutes]m’ or ‘[n_seconds]s’. For never_connected agents, uses the register date. For example, `7d`, `10s` and `10` are valid values. If no time unit is specified, seconds are used (optional)
	osPlatform := "osPlatform_example" // string | Filter by OS platform (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	manager := "manager_example" // string | Filter by manager hostname where agents are connected to (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	group := "group_example" // string | Filter by group of agents (optional)
	nodeName := "nodeName_example" // string | Filter by node name (optional)
	name := "name_example" // string | Filter by name (optional)
	ip := "ip_example" // string | Filter by the IP used by the agent to communicate with the manager. If it's not available, it will have the same value as registerIP (optional)
	registerIP := "registerIP_example" // string | Filter by the IP used when registering the agent (optional)
	groupConfigStatus := "groupConfigStatus_example" // string | Agent groups configuration sync status (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetAgents(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Offset(offset).Limit(limit).Select_(select_).Sort(sort).Search(search).Status(status).Q(q).OlderThan(olderThan).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).GroupConfigStatus(groupConfigStatus).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetAgents`: ApiControllersAgentControllerGetAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **status** | **[]string** | Filter by agent status (use commas to enter multiple statuses) |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **olderThan** | **string** | Filter out agents whose time lapse from last keep alive signal is longer than specified. Time in seconds, ‘[n_days]d’, ‘[n_hours]h’, ‘[n_minutes]m’ or ‘[n_seconds]s’. For never_connected agents, uses the register date. For example, &#x60;7d&#x60;, &#x60;10s&#x60; and &#x60;10&#x60; are valid values. If no time unit is specified, seconds are used |
 **osPlatform** | **string** | Filter by OS platform |
 **osVersion** | **string** | Filter by OS version |
 **osName** | **string** | Filter by OS name |
 **manager** | **string** | Filter by manager hostname where agents are connected to |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **group** | **string** | Filter by group of agents |
 **nodeName** | **string** | Filter by node name |
 **name** | **string** | Filter by name |
 **ip** | **string** | Filter by the IP used by the agent to communicate with the manager. If it&#39;s not available, it will have the same value as registerIP |
 **registerIP** | **string** | Filter by the IP used when registering the agent |
 **groupConfigStatus** | **string** | Agent groups configuration sync status |
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


## ApiControllersAgentControllerGetComponentStats

> ApiResponse ApiControllersAgentControllerGetComponentStats(ctx, agentId, component).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get agent's component stats



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
	component := "component_example" // string | Selected component stats
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetComponentStats(context.Background(), agentId, component).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetComponentStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetComponentStats`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetComponentStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |
**component** | **string** | Selected component stats |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetComponentStatsRequest struct via the builder pattern


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


## ApiControllersAgentControllerGetDaemonStats

> ApiControllersAgentControllerGetDaemonStats200Response ApiControllersAgentControllerGetDaemonStats(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()

Get Wazuh daemon stats from an agent



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
	daemonsList := []string{"DaemonsList_example"} // []string | List of daemon names (separated by comma), all daemons selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetDaemonStats(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).DaemonsList(daemonsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetDaemonStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetDaemonStats`: ApiControllersAgentControllerGetDaemonStats200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetDaemonStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetDaemonStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **daemonsList** | **[]string** | List of daemon names (separated by comma), all daemons selected by default if not specified |

### Return type

[**ApiControllersAgentControllerGetDaemonStats200Response**](ApiControllersAgentControllerGetDaemonStats200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerGetSyncAgent

> ApiControllersAgentControllerGetSyncAgent200Response ApiControllersAgentControllerGetSyncAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get configuration sync status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerGetSyncAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerGetSyncAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerGetSyncAgent`: ApiControllersAgentControllerGetSyncAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerGetSyncAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerGetSyncAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerGetSyncAgent200Response**](ApiControllersAgentControllerGetSyncAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerInsertAgent

> ApiControllersAgentControllerAddAgent200Response ApiControllersAgentControllerInsertAgent(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentInsertBody(agentInsertBody).Execute()

Add agent full



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
	agentInsertBody := *openapiclient.NewAgentInsertBody("Name_example") // AgentInsertBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerInsertAgent(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentInsertBody(agentInsertBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerInsertAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerInsertAgent`: ApiControllersAgentControllerAddAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerInsertAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerInsertAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentInsertBody** | [**AgentInsertBody**](AgentInsertBody.md) |  |

### Return type

[**ApiControllersAgentControllerAddAgent200Response**](ApiControllersAgentControllerAddAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerPostNewAgent

> ApiControllersAgentControllerAddAgent200Response ApiControllersAgentControllerPostNewAgent(ctx).AgentName(agentName).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Add agent quick



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
	agentName := "agentName_example" // string | Agent name. The special characters allowed are: '-','_','.'
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerPostNewAgent(context.Background()).AgentName(agentName).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerPostNewAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPostNewAgent`: ApiControllersAgentControllerAddAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerPostNewAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPostNewAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentName** | **string** | Agent name. The special characters allowed are: &#39;-&#39;,&#39;_&#39;,&#39;.&#39;  |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerAddAgent200Response**](ApiControllersAgentControllerAddAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerPutAgentSingleGroup

> ApiResponse ApiControllersAgentControllerPutAgentSingleGroup(ctx, agentId, groupId).Pretty(pretty).WaitForComplete(waitForComplete).ForceSingleGroup(forceSingleGroup).Execute()

Assign agent to group



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
	groupId := "groupId_example" // string | Group ID. (Name of the group)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	forceSingleGroup := true // bool | Removes the agent from all groups to which it belongs and assigns it to the specified group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerPutAgentSingleGroup(context.Background(), agentId, groupId).Pretty(pretty).WaitForComplete(waitForComplete).ForceSingleGroup(forceSingleGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerPutAgentSingleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPutAgentSingleGroup`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerPutAgentSingleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPutAgentSingleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **forceSingleGroup** | **bool** | Removes the agent from all groups to which it belongs and assigns it to the specified group |

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


## ApiControllersAgentControllerPutMultipleAgentSingleGroup

> ApiControllersAgentControllerDeleteAgents200Response ApiControllersAgentControllerPutMultipleAgentSingleGroup(ctx).GroupId(groupId).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).ForceSingleGroup(forceSingleGroup).Execute()

Assign agents to group



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)
	forceSingleGroup := true // bool | Removes the agent from all groups to which it belongs and assigns it to the specified group (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerPutMultipleAgentSingleGroup(context.Background()).GroupId(groupId).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).ForceSingleGroup(forceSingleGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerPutMultipleAgentSingleGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPutMultipleAgentSingleGroup`: ApiControllersAgentControllerDeleteAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerPutMultipleAgentSingleGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPutMultipleAgentSingleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **string** | Group ID. (Name of the group) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |
 **forceSingleGroup** | **bool** | Removes the agent from all groups to which it belongs and assigns it to the specified group |

### Return type

[**ApiControllersAgentControllerDeleteAgents200Response**](ApiControllersAgentControllerDeleteAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerPutUpgradeAgents

> ApiResponse ApiControllersAgentControllerPutUpgradeAgents(ctx).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).WpkRepo(wpkRepo).UpgradeVersion(upgradeVersion).UseHttp(useHttp).Force(force).PackageType(packageType).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()

Upgrade agents



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	wpkRepo := "wpkRepo_example" // string | WPK repository (optional)
	upgradeVersion := "upgradeVersion_example" // string | Wazuh version to upgrade to (optional)
	useHttp := true // bool | Use http protocol. If it's false use https. By default the value is set to false (optional) (default to false)
	force := true // bool | Force upgrade (optional) (default to false)
	packageType := "packageType_example" // string | Select package type to use. By default, the manager infers this. (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	osPlatform := "osPlatform_example" // string | Filter by OS platform (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	manager := "manager_example" // string | Filter by manager hostname where agents are connected to (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	group := "group_example" // string | Filter by group of agents (optional)
	nodeName := "nodeName_example" // string | Filter by node name (optional)
	name := "name_example" // string | Filter by name (optional)
	ip := "ip_example" // string | Filter by the IP used by the agent to communicate with the manager. If it's not available, it will have the same value as registerIP (optional)
	registerIP := "registerIP_example" // string | Filter by the IP used when registering the agent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerPutUpgradeAgents(context.Background()).AgentsList(agentsList).Pretty(pretty).WaitForComplete(waitForComplete).WpkRepo(wpkRepo).UpgradeVersion(upgradeVersion).UseHttp(useHttp).Force(force).PackageType(packageType).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerPutUpgradeAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPutUpgradeAgents`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerPutUpgradeAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPutUpgradeAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **wpkRepo** | **string** | WPK repository |
 **upgradeVersion** | **string** | Wazuh version to upgrade to |
 **useHttp** | **bool** | Use http protocol. If it&#39;s false use https. By default the value is set to false | [default to false]
 **force** | **bool** | Force upgrade | [default to false]
 **packageType** | **string** | Select package type to use. By default, the manager infers this. |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **osPlatform** | **string** | Filter by OS platform |
 **osVersion** | **string** | Filter by OS version |
 **osName** | **string** | Filter by OS name |
 **manager** | **string** | Filter by manager hostname where agents are connected to |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **group** | **string** | Filter by group of agents |
 **nodeName** | **string** | Filter by node name |
 **name** | **string** | Filter by name |
 **ip** | **string** | Filter by the IP used by the agent to communicate with the manager. If it&#39;s not available, it will have the same value as registerIP |
 **registerIP** | **string** | Filter by the IP used when registering the agent |

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


## ApiControllersAgentControllerPutUpgradeCustomAgents

> ApiResponse ApiControllersAgentControllerPutUpgradeCustomAgents(ctx).AgentsList(agentsList).FilePath(filePath).Pretty(pretty).WaitForComplete(waitForComplete).Installer(installer).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()

Upgrade agents custom



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), use the keyword `all` to select all agents
	filePath := "filePath_example" // string | Full path to the WPK file. The file must be on a folder on the Wazuh's installation directory (by default, <code>/var/ossec</code>)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	installer := "installer_example" // string | Installation script. Default is <code>upgrade.sh</code> or <code>upgrade.bat</code> for windows agents (optional)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	osPlatform := "osPlatform_example" // string | Filter by OS platform (optional)
	osVersion := "osVersion_example" // string | Filter by OS version (optional)
	osName := "osName_example" // string | Filter by OS name (optional)
	manager := "manager_example" // string | Filter by manager hostname where agents are connected to (optional)
	version := "version_example" // string | Filter by agents version using one of the following formats: 'X.Y.Z', 'vX.Y.Z', 'wazuh X.Y.Z' or 'wazuh vX.Y.Z'. For example: '4.4.0' (optional)
	group := "group_example" // string | Filter by group of agents (optional)
	nodeName := "nodeName_example" // string | Filter by node name (optional)
	name := "name_example" // string | Filter by name (optional)
	ip := "ip_example" // string | Filter by the IP used by the agent to communicate with the manager. If it's not available, it will have the same value as registerIP (optional)
	registerIP := "registerIP_example" // string | Filter by the IP used when registering the agent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerPutUpgradeCustomAgents(context.Background()).AgentsList(agentsList).FilePath(filePath).Pretty(pretty).WaitForComplete(waitForComplete).Installer(installer).Q(q).OsPlatform(osPlatform).OsVersion(osVersion).OsName(osName).Manager(manager).Version(version).Group(group).NodeName(nodeName).Name(name).Ip(ip).RegisterIP(registerIP).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerPutUpgradeCustomAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerPutUpgradeCustomAgents`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerPutUpgradeCustomAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerPutUpgradeCustomAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agentsList** | **[]string** | List of agent IDs (separated by comma), use the keyword &#x60;all&#x60; to select all agents |
 **filePath** | **string** | Full path to the WPK file. The file must be on a folder on the Wazuh&#39;s installation directory (by default, &lt;code&gt;/var/ossec&lt;/code&gt;) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **installer** | **string** | Installation script. Default is &lt;code&gt;upgrade.sh&lt;/code&gt; or &lt;code&gt;upgrade.bat&lt;/code&gt; for windows agents |
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **osPlatform** | **string** | Filter by OS platform |
 **osVersion** | **string** | Filter by OS version |
 **osName** | **string** | Filter by OS name |
 **manager** | **string** | Filter by manager hostname where agents are connected to |
 **version** | **string** | Filter by agents version using one of the following formats: &#39;X.Y.Z&#39;, &#39;vX.Y.Z&#39;, &#39;wazuh X.Y.Z&#39; or &#39;wazuh vX.Y.Z&#39;. For example: &#39;4.4.0&#39; |
 **group** | **string** | Filter by group of agents |
 **nodeName** | **string** | Filter by node name |
 **name** | **string** | Filter by name |
 **ip** | **string** | Filter by the IP used by the agent to communicate with the manager. If it&#39;s not available, it will have the same value as registerIP |
 **registerIP** | **string** | Filter by the IP used when registering the agent |

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


## ApiControllersAgentControllerReconnectAgents

> ApiControllersAgentControllerReconnectAgents200Response ApiControllersAgentControllerReconnectAgents(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()

Force reconnect agents



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerReconnectAgents(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerReconnectAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerReconnectAgents`: ApiControllersAgentControllerReconnectAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerReconnectAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerReconnectAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |

### Return type

[**ApiControllersAgentControllerReconnectAgents200Response**](ApiControllersAgentControllerReconnectAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerRestartAgent

> ApiControllersAgentControllerRestartAgent200Response ApiControllersAgentControllerRestartAgent(ctx, agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Restart agent



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerRestartAgent(context.Background(), agentId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerRestartAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerRestartAgent`: ApiControllersAgentControllerRestartAgent200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerRestartAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** | Agent ID. All possible values from 000 onwards |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerRestartAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerRestartAgent200Response**](ApiControllersAgentControllerRestartAgent200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerRestartAgents

> ApiControllersAgentControllerRestartAgents200Response ApiControllersAgentControllerRestartAgents(ctx).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()

Restart agents



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
	agentsList := []string{"Inner_example"} // []string | List of agent IDs (separated by comma), all agents selected by default if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerRestartAgents(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).AgentsList(agentsList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerRestartAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerRestartAgents`: ApiControllersAgentControllerRestartAgents200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerRestartAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerRestartAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **agentsList** | **[]string** | List of agent IDs (separated by comma), all agents selected by default if not specified |

### Return type

[**ApiControllersAgentControllerRestartAgents200Response**](ApiControllersAgentControllerRestartAgents200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerRestartAgentsByGroup

> ApiControllersAgentControllerRestartAgentsByGroup200Response ApiControllersAgentControllerRestartAgentsByGroup(ctx, groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Restart agents in group



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerRestartAgentsByGroup(context.Background(), groupId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerRestartAgentsByGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerRestartAgentsByGroup`: ApiControllersAgentControllerRestartAgentsByGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerRestartAgentsByGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Group ID. (Name of the group) |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerRestartAgentsByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerRestartAgentsByGroup200Response**](ApiControllersAgentControllerRestartAgentsByGroup200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersAgentControllerRestartAgentsByNode

> ApiControllersAgentControllerRestartAgentsByNode200Response ApiControllersAgentControllerRestartAgentsByNode(ctx, nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Restart agents in node



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
	resp, r, err := apiClient.AgentsAPI.ApiControllersAgentControllerRestartAgentsByNode(context.Background(), nodeId).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentsAPI.ApiControllersAgentControllerRestartAgentsByNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersAgentControllerRestartAgentsByNode`: ApiControllersAgentControllerRestartAgentsByNode200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentsAPI.ApiControllersAgentControllerRestartAgentsByNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | Cluster node name |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersAgentControllerRestartAgentsByNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersAgentControllerRestartAgentsByNode200Response**](ApiControllersAgentControllerRestartAgentsByNode200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
