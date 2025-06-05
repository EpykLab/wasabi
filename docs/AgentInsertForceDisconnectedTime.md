# AgentInsertForceDisconnectedTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable force disconnected_time option | [optional] [default to true]
**Value** | Pointer to **string** | Time the agent must has been disconnected to force the insertion. Time in seconds, ‘[n_days]d’, ‘[n_hours]h’, ‘[n_minutes]m’ or ‘[n_seconds]s’. For example, &#x60;7d&#x60;, &#x60;10s&#x60; and &#x60;10&#x60; are valid values. If no time unit is specified, seconds are used | [optional] [default to "1h"]

## Methods

### NewAgentInsertForceDisconnectedTime

`func NewAgentInsertForceDisconnectedTime() *AgentInsertForceDisconnectedTime`

NewAgentInsertForceDisconnectedTime instantiates a new AgentInsertForceDisconnectedTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInsertForceDisconnectedTimeWithDefaults

`func NewAgentInsertForceDisconnectedTimeWithDefaults() *AgentInsertForceDisconnectedTime`

NewAgentInsertForceDisconnectedTimeWithDefaults instantiates a new AgentInsertForceDisconnectedTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AgentInsertForceDisconnectedTime) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentInsertForceDisconnectedTime) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentInsertForceDisconnectedTime) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentInsertForceDisconnectedTime) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetValue

`func (o *AgentInsertForceDisconnectedTime) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AgentInsertForceDisconnectedTime) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AgentInsertForceDisconnectedTime) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AgentInsertForceDisconnectedTime) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


