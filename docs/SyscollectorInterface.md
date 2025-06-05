# SyscollectorInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC Address of the network interface | [optional] 
**Mtu** | Pointer to **int32** | Network interface&#39;s Maximum Transfer Unit | [optional] 
**Name** | Pointer to **string** | Network interface name | [optional] 
**Rx** | Pointer to [**NetworkInterfaceReceivedPackets**](NetworkInterfaceReceivedPackets.md) |  | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**State** | Pointer to **string** | Network interface state | [optional] 
**Tx** | Pointer to [**NetworkInterfaceSentPackets**](NetworkInterfaceSentPackets.md) |  | [optional] 
**Type** | Pointer to **string** | Network interface type | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorInterface

`func NewSyscollectorInterface() *SyscollectorInterface`

NewSyscollectorInterface instantiates a new SyscollectorInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorInterfaceWithDefaults

`func NewSyscollectorInterfaceWithDefaults() *SyscollectorInterface`

NewSyscollectorInterfaceWithDefaults instantiates a new SyscollectorInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *SyscollectorInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SyscollectorInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SyscollectorInterface) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SyscollectorInterface) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMtu

`func (o *SyscollectorInterface) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *SyscollectorInterface) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *SyscollectorInterface) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *SyscollectorInterface) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *SyscollectorInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyscollectorInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyscollectorInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyscollectorInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRx

`func (o *SyscollectorInterface) GetRx() NetworkInterfaceReceivedPackets`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *SyscollectorInterface) GetRxOk() (*NetworkInterfaceReceivedPackets, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *SyscollectorInterface) SetRx(v NetworkInterfaceReceivedPackets)`

SetRx sets Rx field to given value.

### HasRx

`func (o *SyscollectorInterface) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorInterface) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorInterface) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorInterface) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorInterface) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetState

`func (o *SyscollectorInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SyscollectorInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SyscollectorInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SyscollectorInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTx

`func (o *SyscollectorInterface) GetTx() NetworkInterfaceSentPackets`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *SyscollectorInterface) GetTxOk() (*NetworkInterfaceSentPackets, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *SyscollectorInterface) SetTx(v NetworkInterfaceSentPackets)`

SetTx sets Tx field to given value.

### HasTx

`func (o *SyscollectorInterface) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetType

`func (o *SyscollectorInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyscollectorInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyscollectorInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SyscollectorInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorInterface) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorInterface) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorInterface) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorInterface) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


