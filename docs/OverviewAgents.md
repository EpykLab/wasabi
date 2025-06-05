# OverviewAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]OverviewAgentsNodesInner**](OverviewAgentsNodesInner.md) | Active nodes in the cluster | 
**Groups** | [**[]AgentGroup**](AgentGroup.md) | Recount of the number of Wazuh agents group by Wazuh groups | 
**AgentOs** | [**[]OverviewAgentsAgentOsInner**](OverviewAgentsAgentOsInner.md) | Recount of the number of Wazuh agents group by OS | 
**AgentStatus** | [**AgentsSummaryStatus**](AgentsSummaryStatus.md) |  | 
**AgentVersion** | [**[]OverviewAgentsAgentVersionInner**](OverviewAgentsAgentVersionInner.md) | Recount of the number of Wazuh agents group by version | 
**LastRegisteredAgent** | [**[]Agent**](Agent.md) |  | 

## Methods

### NewOverviewAgents

`func NewOverviewAgents(nodes []OverviewAgentsNodesInner, groups []AgentGroup, agentOs []OverviewAgentsAgentOsInner, agentStatus AgentsSummaryStatus, agentVersion []OverviewAgentsAgentVersionInner, lastRegisteredAgent []Agent, ) *OverviewAgents`

NewOverviewAgents instantiates a new OverviewAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewAgentsWithDefaults

`func NewOverviewAgentsWithDefaults() *OverviewAgents`

NewOverviewAgentsWithDefaults instantiates a new OverviewAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *OverviewAgents) GetNodes() []OverviewAgentsNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *OverviewAgents) GetNodesOk() (*[]OverviewAgentsNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *OverviewAgents) SetNodes(v []OverviewAgentsNodesInner)`

SetNodes sets Nodes field to given value.


### GetGroups

`func (o *OverviewAgents) GetGroups() []AgentGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OverviewAgents) GetGroupsOk() (*[]AgentGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OverviewAgents) SetGroups(v []AgentGroup)`

SetGroups sets Groups field to given value.


### GetAgentOs

`func (o *OverviewAgents) GetAgentOs() []OverviewAgentsAgentOsInner`

GetAgentOs returns the AgentOs field if non-nil, zero value otherwise.

### GetAgentOsOk

`func (o *OverviewAgents) GetAgentOsOk() (*[]OverviewAgentsAgentOsInner, bool)`

GetAgentOsOk returns a tuple with the AgentOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentOs

`func (o *OverviewAgents) SetAgentOs(v []OverviewAgentsAgentOsInner)`

SetAgentOs sets AgentOs field to given value.


### GetAgentStatus

`func (o *OverviewAgents) GetAgentStatus() AgentsSummaryStatus`

GetAgentStatus returns the AgentStatus field if non-nil, zero value otherwise.

### GetAgentStatusOk

`func (o *OverviewAgents) GetAgentStatusOk() (*AgentsSummaryStatus, bool)`

GetAgentStatusOk returns a tuple with the AgentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentStatus

`func (o *OverviewAgents) SetAgentStatus(v AgentsSummaryStatus)`

SetAgentStatus sets AgentStatus field to given value.


### GetAgentVersion

`func (o *OverviewAgents) GetAgentVersion() []OverviewAgentsAgentVersionInner`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *OverviewAgents) GetAgentVersionOk() (*[]OverviewAgentsAgentVersionInner, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *OverviewAgents) SetAgentVersion(v []OverviewAgentsAgentVersionInner)`

SetAgentVersion sets AgentVersion field to given value.


### GetLastRegisteredAgent

`func (o *OverviewAgents) GetLastRegisteredAgent() []Agent`

GetLastRegisteredAgent returns the LastRegisteredAgent field if non-nil, zero value otherwise.

### GetLastRegisteredAgentOk

`func (o *OverviewAgents) GetLastRegisteredAgentOk() (*[]Agent, bool)`

GetLastRegisteredAgentOk returns a tuple with the LastRegisteredAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRegisteredAgent

`func (o *OverviewAgents) SetLastRegisteredAgent(v []Agent)`

SetLastRegisteredAgent sets LastRegisteredAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


