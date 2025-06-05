# NetworkInterfaceSentPackets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Number of bytes in the network interface | [optional] 
**Dropped** | Pointer to **int32** | Number of dropped packages in the network interface | [optional] 
**Error** | Pointer to **int32** | Number of packages containing any error in the network interface | [optional] 
**Packets** | Pointer to **int32** | Number of packages in the network interface | [optional] 

## Methods

### NewNetworkInterfaceSentPackets

`func NewNetworkInterfaceSentPackets() *NetworkInterfaceSentPackets`

NewNetworkInterfaceSentPackets instantiates a new NetworkInterfaceSentPackets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceSentPacketsWithDefaults

`func NewNetworkInterfaceSentPacketsWithDefaults() *NetworkInterfaceSentPackets`

NewNetworkInterfaceSentPacketsWithDefaults instantiates a new NetworkInterfaceSentPackets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *NetworkInterfaceSentPackets) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *NetworkInterfaceSentPackets) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *NetworkInterfaceSentPackets) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *NetworkInterfaceSentPackets) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetDropped

`func (o *NetworkInterfaceSentPackets) GetDropped() int32`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *NetworkInterfaceSentPackets) GetDroppedOk() (*int32, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *NetworkInterfaceSentPackets) SetDropped(v int32)`

SetDropped sets Dropped field to given value.

### HasDropped

`func (o *NetworkInterfaceSentPackets) HasDropped() bool`

HasDropped returns a boolean if a field has been set.

### GetError

`func (o *NetworkInterfaceSentPackets) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkInterfaceSentPackets) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkInterfaceSentPackets) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkInterfaceSentPackets) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPackets

`func (o *NetworkInterfaceSentPackets) GetPackets() int32`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *NetworkInterfaceSentPackets) GetPacketsOk() (*int32, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *NetworkInterfaceSentPackets) SetPackets(v int32)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *NetworkInterfaceSentPackets) HasPackets() bool`

HasPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


