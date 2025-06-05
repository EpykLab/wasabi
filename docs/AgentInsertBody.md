# AgentInsertBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Agent ID | [optional] 
**Key** | Pointer to **string** | Key to use when communicating with the manager. The agent must have the same key on its &#x60;client.keys&#x60; file | [optional] 
**Name** | **string** | Agent name | 
**Ip** | Pointer to **string** | If this is not included, the API will get the IP automatically. Allowed values: IP, IP/NET, ANY | [optional] 
**Force** | Pointer to [**AgentInsertForce**](AgentInsertForce.md) |  | [optional] 

## Methods

### NewAgentInsertBody

`func NewAgentInsertBody(name string, ) *AgentInsertBody`

NewAgentInsertBody instantiates a new AgentInsertBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInsertBodyWithDefaults

`func NewAgentInsertBodyWithDefaults() *AgentInsertBody`

NewAgentInsertBodyWithDefaults instantiates a new AgentInsertBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentInsertBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentInsertBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentInsertBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentInsertBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AgentInsertBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AgentInsertBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AgentInsertBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AgentInsertBody) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AgentInsertBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentInsertBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentInsertBody) SetName(v string)`

SetName sets Name field to given value.


### GetIp

`func (o *AgentInsertBody) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AgentInsertBody) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AgentInsertBody) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AgentInsertBody) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetForce

`func (o *AgentInsertBody) GetForce() AgentInsertForce`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AgentInsertBody) GetForceOk() (*AgentInsertForce, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AgentInsertBody) SetForce(v AgentInsertForce)`

SetForce sets Force field to given value.

### HasForce

`func (o *AgentInsertBody) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


