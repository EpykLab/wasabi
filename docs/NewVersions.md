# NewVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCheckDate** | Pointer to **string** | Datetime of the last query to the CTI service | [optional] 
**CurrentVersion** | Pointer to **string** | Current version in the format vX.Y.Z | [optional] 
**UpdateCheck** | Pointer to **bool** | Flag that indicates if the service is enabled | [optional] 
**LastAvailableMajor** | Pointer to [**NewVersionsLastAvailableMajor**](NewVersionsLastAvailableMajor.md) |  | [optional] 
**LastAvailableMinor** | Pointer to [**NewVersionsLastAvailableMinor**](NewVersionsLastAvailableMinor.md) |  | [optional] 
**LastAvailablePatch** | Pointer to [**NewVersionsLastAvailablePatch**](NewVersionsLastAvailablePatch.md) |  | [optional] 
**Uuid** | Pointer to **string** | Identifier of the Wazuh instance | [optional] 

## Methods

### NewNewVersions

`func NewNewVersions() *NewVersions`

NewNewVersions instantiates a new NewVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewVersionsWithDefaults

`func NewNewVersionsWithDefaults() *NewVersions`

NewNewVersionsWithDefaults instantiates a new NewVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCheckDate

`func (o *NewVersions) GetLastCheckDate() string`

GetLastCheckDate returns the LastCheckDate field if non-nil, zero value otherwise.

### GetLastCheckDateOk

`func (o *NewVersions) GetLastCheckDateOk() (*string, bool)`

GetLastCheckDateOk returns a tuple with the LastCheckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckDate

`func (o *NewVersions) SetLastCheckDate(v string)`

SetLastCheckDate sets LastCheckDate field to given value.

### HasLastCheckDate

`func (o *NewVersions) HasLastCheckDate() bool`

HasLastCheckDate returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *NewVersions) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *NewVersions) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *NewVersions) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *NewVersions) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetUpdateCheck

`func (o *NewVersions) GetUpdateCheck() bool`

GetUpdateCheck returns the UpdateCheck field if non-nil, zero value otherwise.

### GetUpdateCheckOk

`func (o *NewVersions) GetUpdateCheckOk() (*bool, bool)`

GetUpdateCheckOk returns a tuple with the UpdateCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCheck

`func (o *NewVersions) SetUpdateCheck(v bool)`

SetUpdateCheck sets UpdateCheck field to given value.

### HasUpdateCheck

`func (o *NewVersions) HasUpdateCheck() bool`

HasUpdateCheck returns a boolean if a field has been set.

### GetLastAvailableMajor

`func (o *NewVersions) GetLastAvailableMajor() NewVersionsLastAvailableMajor`

GetLastAvailableMajor returns the LastAvailableMajor field if non-nil, zero value otherwise.

### GetLastAvailableMajorOk

`func (o *NewVersions) GetLastAvailableMajorOk() (*NewVersionsLastAvailableMajor, bool)`

GetLastAvailableMajorOk returns a tuple with the LastAvailableMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableMajor

`func (o *NewVersions) SetLastAvailableMajor(v NewVersionsLastAvailableMajor)`

SetLastAvailableMajor sets LastAvailableMajor field to given value.

### HasLastAvailableMajor

`func (o *NewVersions) HasLastAvailableMajor() bool`

HasLastAvailableMajor returns a boolean if a field has been set.

### GetLastAvailableMinor

`func (o *NewVersions) GetLastAvailableMinor() NewVersionsLastAvailableMinor`

GetLastAvailableMinor returns the LastAvailableMinor field if non-nil, zero value otherwise.

### GetLastAvailableMinorOk

`func (o *NewVersions) GetLastAvailableMinorOk() (*NewVersionsLastAvailableMinor, bool)`

GetLastAvailableMinorOk returns a tuple with the LastAvailableMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailableMinor

`func (o *NewVersions) SetLastAvailableMinor(v NewVersionsLastAvailableMinor)`

SetLastAvailableMinor sets LastAvailableMinor field to given value.

### HasLastAvailableMinor

`func (o *NewVersions) HasLastAvailableMinor() bool`

HasLastAvailableMinor returns a boolean if a field has been set.

### GetLastAvailablePatch

`func (o *NewVersions) GetLastAvailablePatch() NewVersionsLastAvailablePatch`

GetLastAvailablePatch returns the LastAvailablePatch field if non-nil, zero value otherwise.

### GetLastAvailablePatchOk

`func (o *NewVersions) GetLastAvailablePatchOk() (*NewVersionsLastAvailablePatch, bool)`

GetLastAvailablePatchOk returns a tuple with the LastAvailablePatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAvailablePatch

`func (o *NewVersions) SetLastAvailablePatch(v NewVersionsLastAvailablePatch)`

SetLastAvailablePatch sets LastAvailablePatch field to given value.

### HasLastAvailablePatch

`func (o *NewVersions) HasLastAvailablePatch() bool`

HasLastAvailablePatch returns a boolean if a field has been set.

### GetUuid

`func (o *NewVersions) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NewVersions) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NewVersions) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NewVersions) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


