# SyscollectorPackages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | Package architecture | [optional] 
**Description** | Pointer to **string** | Brief package description | [optional] 
**Format** | Pointer to **string** | Package format | [optional] 
**Multiarch** | Pointer to **string** | Whether the package has multi architecture support | [optional] 
**Name** | Pointer to **string** | Package name | [optional] 
**Priority** | Pointer to **string** | Package priority | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**Section** | Pointer to **string** | Package section | [optional] 
**Source** | Pointer to **string** | Source section | [optional] 
**Size** | Pointer to **int32** | Installed package size in bytes | [optional] 
**Vendor** | Pointer to **string** | Vendor name | [optional] 
**Version** | Pointer to **string** | Release version installed | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorPackages

`func NewSyscollectorPackages() *SyscollectorPackages`

NewSyscollectorPackages instantiates a new SyscollectorPackages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorPackagesWithDefaults

`func NewSyscollectorPackagesWithDefaults() *SyscollectorPackages`

NewSyscollectorPackagesWithDefaults instantiates a new SyscollectorPackages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *SyscollectorPackages) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *SyscollectorPackages) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *SyscollectorPackages) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *SyscollectorPackages) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetDescription

`func (o *SyscollectorPackages) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SyscollectorPackages) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SyscollectorPackages) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SyscollectorPackages) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFormat

`func (o *SyscollectorPackages) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SyscollectorPackages) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SyscollectorPackages) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SyscollectorPackages) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMultiarch

`func (o *SyscollectorPackages) GetMultiarch() string`

GetMultiarch returns the Multiarch field if non-nil, zero value otherwise.

### GetMultiarchOk

`func (o *SyscollectorPackages) GetMultiarchOk() (*string, bool)`

GetMultiarchOk returns a tuple with the Multiarch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiarch

`func (o *SyscollectorPackages) SetMultiarch(v string)`

SetMultiarch sets Multiarch field to given value.

### HasMultiarch

`func (o *SyscollectorPackages) HasMultiarch() bool`

HasMultiarch returns a boolean if a field has been set.

### GetName

`func (o *SyscollectorPackages) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyscollectorPackages) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyscollectorPackages) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyscollectorPackages) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *SyscollectorPackages) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SyscollectorPackages) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SyscollectorPackages) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SyscollectorPackages) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorPackages) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorPackages) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorPackages) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorPackages) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetSection

`func (o *SyscollectorPackages) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *SyscollectorPackages) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *SyscollectorPackages) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *SyscollectorPackages) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetSource

`func (o *SyscollectorPackages) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyscollectorPackages) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyscollectorPackages) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SyscollectorPackages) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSize

`func (o *SyscollectorPackages) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SyscollectorPackages) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SyscollectorPackages) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SyscollectorPackages) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVendor

`func (o *SyscollectorPackages) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *SyscollectorPackages) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *SyscollectorPackages) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *SyscollectorPackages) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *SyscollectorPackages) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SyscollectorPackages) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SyscollectorPackages) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SyscollectorPackages) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorPackages) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorPackages) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorPackages) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorPackages) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


