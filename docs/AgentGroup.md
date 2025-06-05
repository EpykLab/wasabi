# AgentGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Number of agents belonging to that group | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MergedSum** | Pointer to **string** | MD5 checksum of all group shared files merged in a single one (merged.mg) | [optional] 
**ConfigSum** | Pointer to **string** | MD5 checksum of the group configuration file (agent.conf) | [optional] 

## Methods

### NewAgentGroup

`func NewAgentGroup() *AgentGroup`

NewAgentGroup instantiates a new AgentGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupWithDefaults

`func NewAgentGroupWithDefaults() *AgentGroup`

NewAgentGroupWithDefaults instantiates a new AgentGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AgentGroup) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AgentGroup) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AgentGroup) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AgentGroup) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetName

`func (o *AgentGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMergedSum

`func (o *AgentGroup) GetMergedSum() string`

GetMergedSum returns the MergedSum field if non-nil, zero value otherwise.

### GetMergedSumOk

`func (o *AgentGroup) GetMergedSumOk() (*string, bool)`

GetMergedSumOk returns a tuple with the MergedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedSum

`func (o *AgentGroup) SetMergedSum(v string)`

SetMergedSum sets MergedSum field to given value.

### HasMergedSum

`func (o *AgentGroup) HasMergedSum() bool`

HasMergedSum returns a boolean if a field has been set.

### GetConfigSum

`func (o *AgentGroup) GetConfigSum() string`

GetConfigSum returns the ConfigSum field if non-nil, zero value otherwise.

### GetConfigSumOk

`func (o *AgentGroup) GetConfigSumOk() (*string, bool)`

GetConfigSumOk returns a tuple with the ConfigSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSum

`func (o *AgentGroup) SetConfigSum(v string)`

SetConfigSum sets ConfigSum field to given value.

### HasConfigSum

`func (o *AgentGroup) HasConfigSum() bool`

HasConfigSum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


