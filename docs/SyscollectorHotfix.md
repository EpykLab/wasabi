# SyscollectorHotfix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**Hotfix** | Pointer to **string** | Hotfixes for windows agents | [optional] 

## Methods

### NewSyscollectorHotfix

`func NewSyscollectorHotfix() *SyscollectorHotfix`

NewSyscollectorHotfix instantiates a new SyscollectorHotfix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorHotfixWithDefaults

`func NewSyscollectorHotfixWithDefaults() *SyscollectorHotfix`

NewSyscollectorHotfixWithDefaults instantiates a new SyscollectorHotfix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScan

`func (o *SyscollectorHotfix) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorHotfix) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorHotfix) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorHotfix) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetHotfix

`func (o *SyscollectorHotfix) GetHotfix() string`

GetHotfix returns the Hotfix field if non-nil, zero value otherwise.

### GetHotfixOk

`func (o *SyscollectorHotfix) GetHotfixOk() (*string, bool)`

GetHotfixOk returns a tuple with the Hotfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotfix

`func (o *SyscollectorHotfix) SetHotfix(v string)`

SetHotfix sets Hotfix field to given value.

### HasHotfix

`func (o *SyscollectorHotfix) HasHotfix() bool`

HasHotfix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


