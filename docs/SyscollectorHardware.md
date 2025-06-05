# SyscollectorHardware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardSerial** | Pointer to **string** | Motherboard serial number. This value will be empty in virtual machines | [optional] 
**Cpu** | Pointer to [**SyscollectorHardwareCpu**](SyscollectorHardwareCpu.md) |  | [optional] 
**Ram** | Pointer to [**SyscollectorHardwareRam**](SyscollectorHardwareRam.md) |  | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorHardware

`func NewSyscollectorHardware() *SyscollectorHardware`

NewSyscollectorHardware instantiates a new SyscollectorHardware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorHardwareWithDefaults

`func NewSyscollectorHardwareWithDefaults() *SyscollectorHardware`

NewSyscollectorHardwareWithDefaults instantiates a new SyscollectorHardware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardSerial

`func (o *SyscollectorHardware) GetBoardSerial() string`

GetBoardSerial returns the BoardSerial field if non-nil, zero value otherwise.

### GetBoardSerialOk

`func (o *SyscollectorHardware) GetBoardSerialOk() (*string, bool)`

GetBoardSerialOk returns a tuple with the BoardSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardSerial

`func (o *SyscollectorHardware) SetBoardSerial(v string)`

SetBoardSerial sets BoardSerial field to given value.

### HasBoardSerial

`func (o *SyscollectorHardware) HasBoardSerial() bool`

HasBoardSerial returns a boolean if a field has been set.

### GetCpu

`func (o *SyscollectorHardware) GetCpu() SyscollectorHardwareCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *SyscollectorHardware) GetCpuOk() (*SyscollectorHardwareCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *SyscollectorHardware) SetCpu(v SyscollectorHardwareCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *SyscollectorHardware) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetRam

`func (o *SyscollectorHardware) GetRam() SyscollectorHardwareRam`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *SyscollectorHardware) GetRamOk() (*SyscollectorHardwareRam, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *SyscollectorHardware) SetRam(v SyscollectorHardwareRam)`

SetRam sets Ram field to given value.

### HasRam

`func (o *SyscollectorHardware) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorHardware) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorHardware) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorHardware) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorHardware) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorHardware) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorHardware) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorHardware) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorHardware) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


