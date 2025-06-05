# ClusterLocalConfigData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster name | [optional] 
**NodeName** | Pointer to **string** | Node name | [optional] 
**NodeType** | Pointer to **string** | Node type | [optional] 
**Key** | Pointer to **string** | Cluster key used to encrypt messages | [optional] 
**Port** | Pointer to **int32** | Port used by the **master** node to communicate with workers | [optional] 
**BindAddr** | Pointer to **string** | Network interface used by the **master** to listen to incoming connections | [optional] 
**Nodes** | Pointer to **[]string** | List of cluster master nodes. This list is used by **worker** nodes to connect to the master | [optional] 
**Hidden** | Pointer to **string** | Whether to hide the cluster information in the alerts | [optional] 
**Disabled** | Pointer to **bool** | Whether the cluster is enabled or not | [optional] 

## Methods

### NewClusterLocalConfigData

`func NewClusterLocalConfigData() *ClusterLocalConfigData`

NewClusterLocalConfigData instantiates a new ClusterLocalConfigData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLocalConfigDataWithDefaults

`func NewClusterLocalConfigDataWithDefaults() *ClusterLocalConfigData`

NewClusterLocalConfigDataWithDefaults instantiates a new ClusterLocalConfigData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterLocalConfigData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLocalConfigData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLocalConfigData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLocalConfigData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeName

`func (o *ClusterLocalConfigData) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *ClusterLocalConfigData) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *ClusterLocalConfigData) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *ClusterLocalConfigData) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeType

`func (o *ClusterLocalConfigData) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ClusterLocalConfigData) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ClusterLocalConfigData) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ClusterLocalConfigData) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetKey

`func (o *ClusterLocalConfigData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ClusterLocalConfigData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ClusterLocalConfigData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ClusterLocalConfigData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPort

`func (o *ClusterLocalConfigData) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClusterLocalConfigData) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClusterLocalConfigData) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ClusterLocalConfigData) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetBindAddr

`func (o *ClusterLocalConfigData) GetBindAddr() string`

GetBindAddr returns the BindAddr field if non-nil, zero value otherwise.

### GetBindAddrOk

`func (o *ClusterLocalConfigData) GetBindAddrOk() (*string, bool)`

GetBindAddrOk returns a tuple with the BindAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAddr

`func (o *ClusterLocalConfigData) SetBindAddr(v string)`

SetBindAddr sets BindAddr field to given value.

### HasBindAddr

`func (o *ClusterLocalConfigData) HasBindAddr() bool`

HasBindAddr returns a boolean if a field has been set.

### GetNodes

`func (o *ClusterLocalConfigData) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterLocalConfigData) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterLocalConfigData) SetNodes(v []string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterLocalConfigData) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetHidden

`func (o *ClusterLocalConfigData) GetHidden() string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ClusterLocalConfigData) GetHiddenOk() (*string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ClusterLocalConfigData) SetHidden(v string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ClusterLocalConfigData) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetDisabled

`func (o *ClusterLocalConfigData) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ClusterLocalConfigData) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ClusterLocalConfigData) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ClusterLocalConfigData) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


