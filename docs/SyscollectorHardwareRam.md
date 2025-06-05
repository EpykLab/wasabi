# SyscollectorHardwareRam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | Pointer to **int32** | Current free RAM memory | [optional] 
**Total** | Pointer to **int32** | Total RAM memory | [optional] 
**Usage** | Pointer to **int32** | RAM memory currently used | [optional] 

## Methods

### NewSyscollectorHardwareRam

`func NewSyscollectorHardwareRam() *SyscollectorHardwareRam`

NewSyscollectorHardwareRam instantiates a new SyscollectorHardwareRam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorHardwareRamWithDefaults

`func NewSyscollectorHardwareRamWithDefaults() *SyscollectorHardwareRam`

NewSyscollectorHardwareRamWithDefaults instantiates a new SyscollectorHardwareRam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *SyscollectorHardwareRam) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *SyscollectorHardwareRam) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *SyscollectorHardwareRam) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *SyscollectorHardwareRam) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetTotal

`func (o *SyscollectorHardwareRam) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SyscollectorHardwareRam) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SyscollectorHardwareRam) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SyscollectorHardwareRam) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsage

`func (o *SyscollectorHardwareRam) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SyscollectorHardwareRam) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SyscollectorHardwareRam) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SyscollectorHardwareRam) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


