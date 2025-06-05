# RemotePortInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Bind IP | [optional] 
**Port** | Pointer to **int32** | Port used | [optional] 

## Methods

### NewRemotePortInfo

`func NewRemotePortInfo() *RemotePortInfo`

NewRemotePortInfo instantiates a new RemotePortInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemotePortInfoWithDefaults

`func NewRemotePortInfoWithDefaults() *RemotePortInfo`

NewRemotePortInfoWithDefaults instantiates a new RemotePortInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *RemotePortInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RemotePortInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RemotePortInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RemotePortInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *RemotePortInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemotePortInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemotePortInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemotePortInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


