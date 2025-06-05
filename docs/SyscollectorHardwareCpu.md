# SyscollectorHardwareCpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cores** | Pointer to **int32** | Number of cores the CPU has | [optional] 
**Mhz** | Pointer to **float32** | Mhz the CPU has | [optional] 
**Name** | Pointer to **string** | CPU name | [optional] 

## Methods

### NewSyscollectorHardwareCpu

`func NewSyscollectorHardwareCpu() *SyscollectorHardwareCpu`

NewSyscollectorHardwareCpu instantiates a new SyscollectorHardwareCpu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorHardwareCpuWithDefaults

`func NewSyscollectorHardwareCpuWithDefaults() *SyscollectorHardwareCpu`

NewSyscollectorHardwareCpuWithDefaults instantiates a new SyscollectorHardwareCpu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *SyscollectorHardwareCpu) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *SyscollectorHardwareCpu) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *SyscollectorHardwareCpu) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *SyscollectorHardwareCpu) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetMhz

`func (o *SyscollectorHardwareCpu) GetMhz() float32`

GetMhz returns the Mhz field if non-nil, zero value otherwise.

### GetMhzOk

`func (o *SyscollectorHardwareCpu) GetMhzOk() (*float32, bool)`

GetMhzOk returns a tuple with the Mhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMhz

`func (o *SyscollectorHardwareCpu) SetMhz(v float32)`

SetMhz sets Mhz field to given value.

### HasMhz

`func (o *SyscollectorHardwareCpu) HasMhz() bool`

HasMhz returns a boolean if a field has been set.

### GetName

`func (o *SyscollectorHardwareCpu) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyscollectorHardwareCpu) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyscollectorHardwareCpu) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyscollectorHardwareCpu) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


