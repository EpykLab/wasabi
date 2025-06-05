# SCADatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Brief description of what the policy is checking | [optional] 
**EndScan** | Pointer to **time.Time** | When the last scan finished | [optional] 
**Fail** | Pointer to **int32** | Number of failed checks. If this number is higher than 0 the host has a vulnerability | [optional] 
**Name** | Pointer to **string** | Policy name | [optional] 
**Pass** | Pointer to **int32** | Number of passed checks | [optional] 
**PolicyId** | Pointer to **string** | Policy ID | [optional] 
**References** | Pointer to **string** | A link to a documentation page about the policy | [optional] 
**Score** | Pointer to **int32** | Percentage of passed checks | [optional] 
**StartScan** | Pointer to **time.Time** | When the last scan started | [optional] 

## Methods

### NewSCADatabase

`func NewSCADatabase() *SCADatabase`

NewSCADatabase instantiates a new SCADatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCADatabaseWithDefaults

`func NewSCADatabaseWithDefaults() *SCADatabase`

NewSCADatabaseWithDefaults instantiates a new SCADatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SCADatabase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SCADatabase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SCADatabase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SCADatabase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndScan

`func (o *SCADatabase) GetEndScan() time.Time`

GetEndScan returns the EndScan field if non-nil, zero value otherwise.

### GetEndScanOk

`func (o *SCADatabase) GetEndScanOk() (*time.Time, bool)`

GetEndScanOk returns a tuple with the EndScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndScan

`func (o *SCADatabase) SetEndScan(v time.Time)`

SetEndScan sets EndScan field to given value.

### HasEndScan

`func (o *SCADatabase) HasEndScan() bool`

HasEndScan returns a boolean if a field has been set.

### GetFail

`func (o *SCADatabase) GetFail() int32`

GetFail returns the Fail field if non-nil, zero value otherwise.

### GetFailOk

`func (o *SCADatabase) GetFailOk() (*int32, bool)`

GetFailOk returns a tuple with the Fail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFail

`func (o *SCADatabase) SetFail(v int32)`

SetFail sets Fail field to given value.

### HasFail

`func (o *SCADatabase) HasFail() bool`

HasFail returns a boolean if a field has been set.

### GetName

`func (o *SCADatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SCADatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SCADatabase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SCADatabase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPass

`func (o *SCADatabase) GetPass() int32`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *SCADatabase) GetPassOk() (*int32, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *SCADatabase) SetPass(v int32)`

SetPass sets Pass field to given value.

### HasPass

`func (o *SCADatabase) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetPolicyId

`func (o *SCADatabase) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SCADatabase) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SCADatabase) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SCADatabase) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetReferences

`func (o *SCADatabase) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SCADatabase) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SCADatabase) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SCADatabase) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetScore

`func (o *SCADatabase) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SCADatabase) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SCADatabase) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SCADatabase) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStartScan

`func (o *SCADatabase) GetStartScan() time.Time`

GetStartScan returns the StartScan field if non-nil, zero value otherwise.

### GetStartScanOk

`func (o *SCADatabase) GetStartScanOk() (*time.Time, bool)`

GetStartScanOk returns a tuple with the StartScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartScan

`func (o *SCADatabase) SetStartScan(v time.Time)`

SetStartScan sets StartScan field to given value.

### HasStartScan

`func (o *SCADatabase) HasStartScan() bool`

HasStartScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


