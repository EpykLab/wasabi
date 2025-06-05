# WazuhInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | Wazuh installation path | [optional] 
**Version** | Pointer to **string** | Wazuh version | [optional] 
**Type** | Pointer to **string** | Wazuh installation type | [optional] 
**MaxAgents** | Pointer to **string** | Maximum number of agents that can be registered. | [optional] 
**OpensslSupport** | Pointer to **string** |  | [optional] 
**TzOffset** | Pointer to **string** |  | [optional] 
**TzName** | Pointer to **string** |  | [optional] 

## Methods

### NewWazuhInfo

`func NewWazuhInfo() *WazuhInfo`

NewWazuhInfo instantiates a new WazuhInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhInfoWithDefaults

`func NewWazuhInfoWithDefaults() *WazuhInfo`

NewWazuhInfoWithDefaults instantiates a new WazuhInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *WazuhInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WazuhInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WazuhInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WazuhInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVersion

`func (o *WazuhInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WazuhInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WazuhInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WazuhInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *WazuhInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WazuhInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WazuhInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WazuhInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaxAgents

`func (o *WazuhInfo) GetMaxAgents() string`

GetMaxAgents returns the MaxAgents field if non-nil, zero value otherwise.

### GetMaxAgentsOk

`func (o *WazuhInfo) GetMaxAgentsOk() (*string, bool)`

GetMaxAgentsOk returns a tuple with the MaxAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgents

`func (o *WazuhInfo) SetMaxAgents(v string)`

SetMaxAgents sets MaxAgents field to given value.

### HasMaxAgents

`func (o *WazuhInfo) HasMaxAgents() bool`

HasMaxAgents returns a boolean if a field has been set.

### GetOpensslSupport

`func (o *WazuhInfo) GetOpensslSupport() string`

GetOpensslSupport returns the OpensslSupport field if non-nil, zero value otherwise.

### GetOpensslSupportOk

`func (o *WazuhInfo) GetOpensslSupportOk() (*string, bool)`

GetOpensslSupportOk returns a tuple with the OpensslSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpensslSupport

`func (o *WazuhInfo) SetOpensslSupport(v string)`

SetOpensslSupport sets OpensslSupport field to given value.

### HasOpensslSupport

`func (o *WazuhInfo) HasOpensslSupport() bool`

HasOpensslSupport returns a boolean if a field has been set.

### GetTzOffset

`func (o *WazuhInfo) GetTzOffset() string`

GetTzOffset returns the TzOffset field if non-nil, zero value otherwise.

### GetTzOffsetOk

`func (o *WazuhInfo) GetTzOffsetOk() (*string, bool)`

GetTzOffsetOk returns a tuple with the TzOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzOffset

`func (o *WazuhInfo) SetTzOffset(v string)`

SetTzOffset sets TzOffset field to given value.

### HasTzOffset

`func (o *WazuhInfo) HasTzOffset() bool`

HasTzOffset returns a boolean if a field has been set.

### GetTzName

`func (o *WazuhInfo) GetTzName() string`

GetTzName returns the TzName field if non-nil, zero value otherwise.

### GetTzNameOk

`func (o *WazuhInfo) GetTzNameOk() (*string, bool)`

GetTzNameOk returns a tuple with the TzName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzName

`func (o *WazuhInfo) SetTzName(v string)`

SetTzName sets TzName field to given value.

### HasTzName

`func (o *WazuhInfo) HasTzName() bool`

HasTzName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


