# AgentDistinct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Wazuh version the agent has installed | [optional] 
**Id** | Pointer to **string** | Agent ID | [optional] 
**Name** | Pointer to **string** | Agent name used at the registration process | [optional] 
**Status** | Pointer to [**AgentStatus**](AgentStatus.md) |  | [optional] 
**ConfigSum** | Pointer to **string** | MD5 checksum of the group configuration file (agent.conf) | [optional] 
**Group** | Pointer to **[]string** | List of groups the agent belongs to | [optional] 
**MergedSum** | Pointer to **string** | MD5 checksum of all group shared files merged in a single one (merged.mg) | [optional] 
**Ip** | Pointer to **string** | IP where the agent communicates with the manager. If the manager can&#39;t get this information, it will be the same as registerIP field | [optional] 
**RegisterIP** | Pointer to **string** | IP used at agent the registration process | [optional] 
**Manager** | Pointer to **string** | Hostname of the manager the agent is reporting to | [optional] 
**NodeName** | Pointer to **string** | ID of the node the agent is reporting to | [optional] 
**DateAdd** | Pointer to **string** | Date when the agent was registered | [optional] 
**LastKeepAlive** | Pointer to **string** | Date when the last keepalive was received from the agent | [optional] 
**Os** | Pointer to [**AgentOs**](AgentOs.md) |  | [optional] 
**StatusCode** | Pointer to **int32** | Agent connection status code | [optional] [default to 0]
**GroupConfigStatus** | Pointer to **string** | Agent groups configuration sync status | [optional] 
**Count** | Pointer to **int32** | Number of agents with the specified unique fields | [optional] 

## Methods

### NewAgentDistinct

`func NewAgentDistinct() *AgentDistinct`

NewAgentDistinct instantiates a new AgentDistinct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentDistinctWithDefaults

`func NewAgentDistinctWithDefaults() *AgentDistinct`

NewAgentDistinctWithDefaults instantiates a new AgentDistinct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AgentDistinct) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentDistinct) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentDistinct) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AgentDistinct) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *AgentDistinct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentDistinct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentDistinct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentDistinct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentDistinct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentDistinct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentDistinct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentDistinct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AgentDistinct) GetStatus() AgentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentDistinct) GetStatusOk() (*AgentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentDistinct) SetStatus(v AgentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentDistinct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConfigSum

`func (o *AgentDistinct) GetConfigSum() string`

GetConfigSum returns the ConfigSum field if non-nil, zero value otherwise.

### GetConfigSumOk

`func (o *AgentDistinct) GetConfigSumOk() (*string, bool)`

GetConfigSumOk returns a tuple with the ConfigSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSum

`func (o *AgentDistinct) SetConfigSum(v string)`

SetConfigSum sets ConfigSum field to given value.

### HasConfigSum

`func (o *AgentDistinct) HasConfigSum() bool`

HasConfigSum returns a boolean if a field has been set.

### GetGroup

`func (o *AgentDistinct) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AgentDistinct) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AgentDistinct) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AgentDistinct) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMergedSum

`func (o *AgentDistinct) GetMergedSum() string`

GetMergedSum returns the MergedSum field if non-nil, zero value otherwise.

### GetMergedSumOk

`func (o *AgentDistinct) GetMergedSumOk() (*string, bool)`

GetMergedSumOk returns a tuple with the MergedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedSum

`func (o *AgentDistinct) SetMergedSum(v string)`

SetMergedSum sets MergedSum field to given value.

### HasMergedSum

`func (o *AgentDistinct) HasMergedSum() bool`

HasMergedSum returns a boolean if a field has been set.

### GetIp

`func (o *AgentDistinct) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AgentDistinct) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AgentDistinct) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AgentDistinct) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRegisterIP

`func (o *AgentDistinct) GetRegisterIP() string`

GetRegisterIP returns the RegisterIP field if non-nil, zero value otherwise.

### GetRegisterIPOk

`func (o *AgentDistinct) GetRegisterIPOk() (*string, bool)`

GetRegisterIPOk returns a tuple with the RegisterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterIP

`func (o *AgentDistinct) SetRegisterIP(v string)`

SetRegisterIP sets RegisterIP field to given value.

### HasRegisterIP

`func (o *AgentDistinct) HasRegisterIP() bool`

HasRegisterIP returns a boolean if a field has been set.

### GetManager

`func (o *AgentDistinct) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *AgentDistinct) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *AgentDistinct) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *AgentDistinct) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetNodeName

`func (o *AgentDistinct) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *AgentDistinct) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *AgentDistinct) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *AgentDistinct) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetDateAdd

`func (o *AgentDistinct) GetDateAdd() string`

GetDateAdd returns the DateAdd field if non-nil, zero value otherwise.

### GetDateAddOk

`func (o *AgentDistinct) GetDateAddOk() (*string, bool)`

GetDateAddOk returns a tuple with the DateAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdd

`func (o *AgentDistinct) SetDateAdd(v string)`

SetDateAdd sets DateAdd field to given value.

### HasDateAdd

`func (o *AgentDistinct) HasDateAdd() bool`

HasDateAdd returns a boolean if a field has been set.

### GetLastKeepAlive

`func (o *AgentDistinct) GetLastKeepAlive() string`

GetLastKeepAlive returns the LastKeepAlive field if non-nil, zero value otherwise.

### GetLastKeepAliveOk

`func (o *AgentDistinct) GetLastKeepAliveOk() (*string, bool)`

GetLastKeepAliveOk returns a tuple with the LastKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKeepAlive

`func (o *AgentDistinct) SetLastKeepAlive(v string)`

SetLastKeepAlive sets LastKeepAlive field to given value.

### HasLastKeepAlive

`func (o *AgentDistinct) HasLastKeepAlive() bool`

HasLastKeepAlive returns a boolean if a field has been set.

### GetOs

`func (o *AgentDistinct) GetOs() AgentOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *AgentDistinct) GetOsOk() (*AgentOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *AgentDistinct) SetOs(v AgentOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *AgentDistinct) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetStatusCode

`func (o *AgentDistinct) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AgentDistinct) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AgentDistinct) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AgentDistinct) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetGroupConfigStatus

`func (o *AgentDistinct) GetGroupConfigStatus() string`

GetGroupConfigStatus returns the GroupConfigStatus field if non-nil, zero value otherwise.

### GetGroupConfigStatusOk

`func (o *AgentDistinct) GetGroupConfigStatusOk() (*string, bool)`

GetGroupConfigStatusOk returns a tuple with the GroupConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupConfigStatus

`func (o *AgentDistinct) SetGroupConfigStatus(v string)`

SetGroupConfigStatus sets GroupConfigStatus field to given value.

### HasGroupConfigStatus

`func (o *AgentDistinct) HasGroupConfigStatus() bool`

HasGroupConfigStatus returns a boolean if a field has been set.

### GetCount

`func (o *AgentDistinct) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentDistinct) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentDistinct) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AgentDistinct) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


