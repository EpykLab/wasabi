# SyscollectorOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | OS architecture | [optional] 
**Hostname** | Pointer to **string** | Machine&#39;s hostname | [optional] 
**Os** | Pointer to [**SyscollectorOSOs**](SyscollectorOSOs.md) |  | [optional] 
**Release** | Pointer to **string** | Release name | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**Sysname** | Pointer to **string** | System name | [optional] 
**Version** | Pointer to **string** | Release version | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorOS

`func NewSyscollectorOS() *SyscollectorOS`

NewSyscollectorOS instantiates a new SyscollectorOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorOSWithDefaults

`func NewSyscollectorOSWithDefaults() *SyscollectorOS`

NewSyscollectorOSWithDefaults instantiates a new SyscollectorOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *SyscollectorOS) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *SyscollectorOS) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *SyscollectorOS) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *SyscollectorOS) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetHostname

`func (o *SyscollectorOS) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyscollectorOS) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyscollectorOS) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyscollectorOS) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOs

`func (o *SyscollectorOS) GetOs() SyscollectorOSOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *SyscollectorOS) GetOsOk() (*SyscollectorOSOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *SyscollectorOS) SetOs(v SyscollectorOSOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *SyscollectorOS) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetRelease

`func (o *SyscollectorOS) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SyscollectorOS) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SyscollectorOS) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *SyscollectorOS) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorOS) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorOS) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorOS) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorOS) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetSysname

`func (o *SyscollectorOS) GetSysname() string`

GetSysname returns the Sysname field if non-nil, zero value otherwise.

### GetSysnameOk

`func (o *SyscollectorOS) GetSysnameOk() (*string, bool)`

GetSysnameOk returns a tuple with the Sysname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysname

`func (o *SyscollectorOS) SetSysname(v string)`

SetSysname sets Sysname field to given value.

### HasSysname

`func (o *SyscollectorOS) HasSysname() bool`

HasSysname returns a boolean if a field has been set.

### GetVersion

`func (o *SyscollectorOS) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SyscollectorOS) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SyscollectorOS) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SyscollectorOS) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorOS) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorOS) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorOS) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorOS) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


