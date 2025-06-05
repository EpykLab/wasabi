# ClusterLocalInfoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | Pointer to **string** | Node name | [optional] 
**Cluster** | Pointer to **string** | Cluster name the node belongs to | [optional] 
**Type** | Pointer to **string** | Node type | [optional] 

## Methods

### NewClusterLocalInfoData

`func NewClusterLocalInfoData() *ClusterLocalInfoData`

NewClusterLocalInfoData instantiates a new ClusterLocalInfoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLocalInfoDataWithDefaults

`func NewClusterLocalInfoDataWithDefaults() *ClusterLocalInfoData`

NewClusterLocalInfoDataWithDefaults instantiates a new ClusterLocalInfoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *ClusterLocalInfoData) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ClusterLocalInfoData) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ClusterLocalInfoData) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ClusterLocalInfoData) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetCluster

`func (o *ClusterLocalInfoData) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ClusterLocalInfoData) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ClusterLocalInfoData) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ClusterLocalInfoData) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetType

`func (o *ClusterLocalInfoData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterLocalInfoData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterLocalInfoData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterLocalInfoData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


