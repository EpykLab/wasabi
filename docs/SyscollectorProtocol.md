# SyscollectorProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcp** | Pointer to [**DHCPStatus**](DHCPStatus.md) |  | [optional] 
**Gateway** | Pointer to **string** | Gateway IP | [optional] 
**Iface** | Pointer to **string** | Network interface name | [optional] 
**ScanId** | Pointer to **int64** | Scan ID | [optional] 
**Type** | Pointer to **string** | Protocol of the interface data | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorProtocol

`func NewSyscollectorProtocol() *SyscollectorProtocol`

NewSyscollectorProtocol instantiates a new SyscollectorProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorProtocolWithDefaults

`func NewSyscollectorProtocolWithDefaults() *SyscollectorProtocol`

NewSyscollectorProtocolWithDefaults instantiates a new SyscollectorProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcp

`func (o *SyscollectorProtocol) GetDhcp() DHCPStatus`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *SyscollectorProtocol) GetDhcpOk() (*DHCPStatus, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *SyscollectorProtocol) SetDhcp(v DHCPStatus)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *SyscollectorProtocol) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetGateway

`func (o *SyscollectorProtocol) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SyscollectorProtocol) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SyscollectorProtocol) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *SyscollectorProtocol) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIface

`func (o *SyscollectorProtocol) GetIface() string`

GetIface returns the Iface field if non-nil, zero value otherwise.

### GetIfaceOk

`func (o *SyscollectorProtocol) GetIfaceOk() (*string, bool)`

GetIfaceOk returns a tuple with the Iface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIface

`func (o *SyscollectorProtocol) SetIface(v string)`

SetIface sets Iface field to given value.

### HasIface

`func (o *SyscollectorProtocol) HasIface() bool`

HasIface returns a boolean if a field has been set.

### GetScanId

`func (o *SyscollectorProtocol) GetScanId() int64`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *SyscollectorProtocol) GetScanIdOk() (*int64, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *SyscollectorProtocol) SetScanId(v int64)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *SyscollectorProtocol) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetType

`func (o *SyscollectorProtocol) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyscollectorProtocol) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyscollectorProtocol) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SyscollectorProtocol) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorProtocol) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorProtocol) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorProtocol) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorProtocol) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


