# NodeHealthcheckNameInfoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Node type | [optional] 
**Version** | Pointer to **string** | Wazuh version installed in the node | [optional] 
**Ip** | Pointer to **string** | IP the node is using to communicate with other nodes in the cluster | [optional] 
**TotalActiveAgents** | Pointer to **int32** | Number of agents currently reporting to that node | [optional] 

## Methods

### NewNodeHealthcheckNameInfoInfo

`func NewNodeHealthcheckNameInfoInfo() *NodeHealthcheckNameInfoInfo`

NewNodeHealthcheckNameInfoInfo instantiates a new NodeHealthcheckNameInfoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHealthcheckNameInfoInfoWithDefaults

`func NewNodeHealthcheckNameInfoInfoWithDefaults() *NodeHealthcheckNameInfoInfo`

NewNodeHealthcheckNameInfoInfoWithDefaults instantiates a new NodeHealthcheckNameInfoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NodeHealthcheckNameInfoInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeHealthcheckNameInfoInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeHealthcheckNameInfoInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeHealthcheckNameInfoInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *NodeHealthcheckNameInfoInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeHealthcheckNameInfoInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeHealthcheckNameInfoInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeHealthcheckNameInfoInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIp

`func (o *NodeHealthcheckNameInfoInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeHealthcheckNameInfoInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeHealthcheckNameInfoInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NodeHealthcheckNameInfoInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetTotalActiveAgents

`func (o *NodeHealthcheckNameInfoInfo) GetTotalActiveAgents() int32`

GetTotalActiveAgents returns the TotalActiveAgents field if non-nil, zero value otherwise.

### GetTotalActiveAgentsOk

`func (o *NodeHealthcheckNameInfoInfo) GetTotalActiveAgentsOk() (*int32, bool)`

GetTotalActiveAgentsOk returns a tuple with the TotalActiveAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActiveAgents

`func (o *NodeHealthcheckNameInfoInfo) SetTotalActiveAgents(v int32)`

SetTotalActiveAgents sets TotalActiveAgents field to given value.

### HasTotalActiveAgents

`func (o *NodeHealthcheckNameInfoInfo) HasTotalActiveAgents() bool`

HasTotalActiveAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


