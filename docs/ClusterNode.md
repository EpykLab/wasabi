# ClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Node type | [optional] 
**Version** | Pointer to **string** | Wazuh version installed in the node | [optional] 
**Ip** | Pointer to **string** | IP the node is using to communicate with other nodes in the cluster | [optional] 
**Name** | Pointer to **string** | Node ID | [optional] 

## Methods

### NewClusterNode

`func NewClusterNode() *ClusterNode`

NewClusterNode instantiates a new ClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeWithDefaults

`func NewClusterNodeWithDefaults() *ClusterNode`

NewClusterNodeWithDefaults instantiates a new ClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterNode) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterNode) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIp

`func (o *ClusterNode) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClusterNode) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClusterNode) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClusterNode) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetName

`func (o *ClusterNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterNode) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


