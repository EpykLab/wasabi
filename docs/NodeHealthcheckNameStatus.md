# NodeHealthcheckNameStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastKeepAlive** | Pointer to **string** |  | [optional] 
**LastSyncAgentinfo** | Pointer to [**NodeHealthcheckNameStatusLastSyncAgentinfo**](NodeHealthcheckNameStatusLastSyncAgentinfo.md) |  | [optional] 
**LastSyncAgentgroup** | Pointer to [**NodeHealthcheckNameStatusLastSyncAgentgroup**](NodeHealthcheckNameStatusLastSyncAgentgroup.md) |  | [optional] 
**LastSyncFullAgentgroup** | Pointer to [**NodeHealthcheckNameStatusLastSyncAgentgroup**](NodeHealthcheckNameStatusLastSyncAgentgroup.md) |  | [optional] 
**LastSyncIntegrity** | Pointer to [**NodeHealthcheckNameStatusLastSyncIntegrity**](NodeHealthcheckNameStatusLastSyncIntegrity.md) |  | [optional] 
**SyncAgentInfoFree** | Pointer to **bool** |  | [optional] 
**SyncIntegrityFree** | Pointer to **bool** |  | [optional] 

## Methods

### NewNodeHealthcheckNameStatus

`func NewNodeHealthcheckNameStatus() *NodeHealthcheckNameStatus`

NewNodeHealthcheckNameStatus instantiates a new NodeHealthcheckNameStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHealthcheckNameStatusWithDefaults

`func NewNodeHealthcheckNameStatusWithDefaults() *NodeHealthcheckNameStatus`

NewNodeHealthcheckNameStatusWithDefaults instantiates a new NodeHealthcheckNameStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastKeepAlive

`func (o *NodeHealthcheckNameStatus) GetLastKeepAlive() string`

GetLastKeepAlive returns the LastKeepAlive field if non-nil, zero value otherwise.

### GetLastKeepAliveOk

`func (o *NodeHealthcheckNameStatus) GetLastKeepAliveOk() (*string, bool)`

GetLastKeepAliveOk returns a tuple with the LastKeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKeepAlive

`func (o *NodeHealthcheckNameStatus) SetLastKeepAlive(v string)`

SetLastKeepAlive sets LastKeepAlive field to given value.

### HasLastKeepAlive

`func (o *NodeHealthcheckNameStatus) HasLastKeepAlive() bool`

HasLastKeepAlive returns a boolean if a field has been set.

### GetLastSyncAgentinfo

`func (o *NodeHealthcheckNameStatus) GetLastSyncAgentinfo() NodeHealthcheckNameStatusLastSyncAgentinfo`

GetLastSyncAgentinfo returns the LastSyncAgentinfo field if non-nil, zero value otherwise.

### GetLastSyncAgentinfoOk

`func (o *NodeHealthcheckNameStatus) GetLastSyncAgentinfoOk() (*NodeHealthcheckNameStatusLastSyncAgentinfo, bool)`

GetLastSyncAgentinfoOk returns a tuple with the LastSyncAgentinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAgentinfo

`func (o *NodeHealthcheckNameStatus) SetLastSyncAgentinfo(v NodeHealthcheckNameStatusLastSyncAgentinfo)`

SetLastSyncAgentinfo sets LastSyncAgentinfo field to given value.

### HasLastSyncAgentinfo

`func (o *NodeHealthcheckNameStatus) HasLastSyncAgentinfo() bool`

HasLastSyncAgentinfo returns a boolean if a field has been set.

### GetLastSyncAgentgroup

`func (o *NodeHealthcheckNameStatus) GetLastSyncAgentgroup() NodeHealthcheckNameStatusLastSyncAgentgroup`

GetLastSyncAgentgroup returns the LastSyncAgentgroup field if non-nil, zero value otherwise.

### GetLastSyncAgentgroupOk

`func (o *NodeHealthcheckNameStatus) GetLastSyncAgentgroupOk() (*NodeHealthcheckNameStatusLastSyncAgentgroup, bool)`

GetLastSyncAgentgroupOk returns a tuple with the LastSyncAgentgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAgentgroup

`func (o *NodeHealthcheckNameStatus) SetLastSyncAgentgroup(v NodeHealthcheckNameStatusLastSyncAgentgroup)`

SetLastSyncAgentgroup sets LastSyncAgentgroup field to given value.

### HasLastSyncAgentgroup

`func (o *NodeHealthcheckNameStatus) HasLastSyncAgentgroup() bool`

HasLastSyncAgentgroup returns a boolean if a field has been set.

### GetLastSyncFullAgentgroup

`func (o *NodeHealthcheckNameStatus) GetLastSyncFullAgentgroup() NodeHealthcheckNameStatusLastSyncAgentgroup`

GetLastSyncFullAgentgroup returns the LastSyncFullAgentgroup field if non-nil, zero value otherwise.

### GetLastSyncFullAgentgroupOk

`func (o *NodeHealthcheckNameStatus) GetLastSyncFullAgentgroupOk() (*NodeHealthcheckNameStatusLastSyncAgentgroup, bool)`

GetLastSyncFullAgentgroupOk returns a tuple with the LastSyncFullAgentgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncFullAgentgroup

`func (o *NodeHealthcheckNameStatus) SetLastSyncFullAgentgroup(v NodeHealthcheckNameStatusLastSyncAgentgroup)`

SetLastSyncFullAgentgroup sets LastSyncFullAgentgroup field to given value.

### HasLastSyncFullAgentgroup

`func (o *NodeHealthcheckNameStatus) HasLastSyncFullAgentgroup() bool`

HasLastSyncFullAgentgroup returns a boolean if a field has been set.

### GetLastSyncIntegrity

`func (o *NodeHealthcheckNameStatus) GetLastSyncIntegrity() NodeHealthcheckNameStatusLastSyncIntegrity`

GetLastSyncIntegrity returns the LastSyncIntegrity field if non-nil, zero value otherwise.

### GetLastSyncIntegrityOk

`func (o *NodeHealthcheckNameStatus) GetLastSyncIntegrityOk() (*NodeHealthcheckNameStatusLastSyncIntegrity, bool)`

GetLastSyncIntegrityOk returns a tuple with the LastSyncIntegrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncIntegrity

`func (o *NodeHealthcheckNameStatus) SetLastSyncIntegrity(v NodeHealthcheckNameStatusLastSyncIntegrity)`

SetLastSyncIntegrity sets LastSyncIntegrity field to given value.

### HasLastSyncIntegrity

`func (o *NodeHealthcheckNameStatus) HasLastSyncIntegrity() bool`

HasLastSyncIntegrity returns a boolean if a field has been set.

### GetSyncAgentInfoFree

`func (o *NodeHealthcheckNameStatus) GetSyncAgentInfoFree() bool`

GetSyncAgentInfoFree returns the SyncAgentInfoFree field if non-nil, zero value otherwise.

### GetSyncAgentInfoFreeOk

`func (o *NodeHealthcheckNameStatus) GetSyncAgentInfoFreeOk() (*bool, bool)`

GetSyncAgentInfoFreeOk returns a tuple with the SyncAgentInfoFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAgentInfoFree

`func (o *NodeHealthcheckNameStatus) SetSyncAgentInfoFree(v bool)`

SetSyncAgentInfoFree sets SyncAgentInfoFree field to given value.

### HasSyncAgentInfoFree

`func (o *NodeHealthcheckNameStatus) HasSyncAgentInfoFree() bool`

HasSyncAgentInfoFree returns a boolean if a field has been set.

### GetSyncIntegrityFree

`func (o *NodeHealthcheckNameStatus) GetSyncIntegrityFree() bool`

GetSyncIntegrityFree returns the SyncIntegrityFree field if non-nil, zero value otherwise.

### GetSyncIntegrityFreeOk

`func (o *NodeHealthcheckNameStatus) GetSyncIntegrityFreeOk() (*bool, bool)`

GetSyncIntegrityFreeOk returns a tuple with the SyncIntegrityFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncIntegrityFree

`func (o *NodeHealthcheckNameStatus) SetSyncIntegrityFree(v bool)`

SetSyncIntegrityFree sets SyncIntegrityFree field to given value.

### HasSyncIntegrityFree

`func (o *NodeHealthcheckNameStatus) HasSyncIntegrityFree() bool`

HasSyncIntegrityFree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


