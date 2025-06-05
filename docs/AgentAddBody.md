# AgentAddBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Agent name | 
**Ip** | Pointer to **string** | If this is not included, the API will get the IP automatically. Allowed values: IP, IP/NET, ANY | [optional] 

## Methods

### NewAgentAddBody

`func NewAgentAddBody(name string, ) *AgentAddBody`

NewAgentAddBody instantiates a new AgentAddBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAddBodyWithDefaults

`func NewAgentAddBodyWithDefaults() *AgentAddBody`

NewAgentAddBodyWithDefaults instantiates a new AgentAddBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AgentAddBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentAddBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentAddBody) SetName(v string)`

SetName sets Name field to given value.


### GetIp

`func (o *AgentAddBody) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AgentAddBody) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AgentAddBody) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AgentAddBody) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


