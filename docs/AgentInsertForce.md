# AgentInsertForce

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable force option | [optional] [default to true]
**DisconnectedTime** | Pointer to [**AgentInsertForceDisconnectedTime**](AgentInsertForceDisconnectedTime.md) |  | [optional] 
**AfterRegistrationTime** | Pointer to **string** | Time the agent must has been registered to force the insertion. Time in seconds, ‘[n_days]d’, ‘[n_hours]h’, ‘[n_minutes]m’ or ‘[n_seconds]s’. For example, &#x60;7d&#x60;, &#x60;10s&#x60; and &#x60;10&#x60; are valid values. If no time unit is specified, seconds are used | [optional] [default to "1h"]

## Methods

### NewAgentInsertForce

`func NewAgentInsertForce() *AgentInsertForce`

NewAgentInsertForce instantiates a new AgentInsertForce object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInsertForceWithDefaults

`func NewAgentInsertForceWithDefaults() *AgentInsertForce`

NewAgentInsertForceWithDefaults instantiates a new AgentInsertForce object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgentInsertForce) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentInsertForce) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentInsertForce) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentInsertForce) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDisconnectedTime

`func (o *AgentInsertForce) GetDisconnectedTime() AgentInsertForceDisconnectedTime`

GetDisconnectedTime returns the DisconnectedTime field if non-nil, zero value otherwise.

### GetDisconnectedTimeOk

`func (o *AgentInsertForce) GetDisconnectedTimeOk() (*AgentInsertForceDisconnectedTime, bool)`

GetDisconnectedTimeOk returns a tuple with the DisconnectedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedTime

`func (o *AgentInsertForce) SetDisconnectedTime(v AgentInsertForceDisconnectedTime)`

SetDisconnectedTime sets DisconnectedTime field to given value.

### HasDisconnectedTime

`func (o *AgentInsertForce) HasDisconnectedTime() bool`

HasDisconnectedTime returns a boolean if a field has been set.

### GetAfterRegistrationTime

`func (o *AgentInsertForce) GetAfterRegistrationTime() string`

GetAfterRegistrationTime returns the AfterRegistrationTime field if non-nil, zero value otherwise.

### GetAfterRegistrationTimeOk

`func (o *AgentInsertForce) GetAfterRegistrationTimeOk() (*string, bool)`

GetAfterRegistrationTimeOk returns a tuple with the AfterRegistrationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterRegistrationTime

`func (o *AgentInsertForce) SetAfterRegistrationTime(v string)`

SetAfterRegistrationTime sets AfterRegistrationTime field to given value.

### HasAfterRegistrationTime

`func (o *AgentInsertForce) HasAfterRegistrationTime() bool`

HasAfterRegistrationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


