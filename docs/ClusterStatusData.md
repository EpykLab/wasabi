# ClusterStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **string** | Whether the cluster is enabled in the Wazuh configuration | [optional] 
**Running** | Pointer to **string** | Whether the cluster daemon is running | [optional] 

## Methods

### NewClusterStatusData

`func NewClusterStatusData() *ClusterStatusData`

NewClusterStatusData instantiates a new ClusterStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusDataWithDefaults

`func NewClusterStatusDataWithDefaults() *ClusterStatusData`

NewClusterStatusDataWithDefaults instantiates a new ClusterStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ClusterStatusData) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterStatusData) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterStatusData) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterStatusData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRunning

`func (o *ClusterStatusData) GetRunning() string`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ClusterStatusData) GetRunningOk() (*string, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ClusterStatusData) SetRunning(v string)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ClusterStatusData) HasRunning() bool`

HasRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


