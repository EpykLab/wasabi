# SCAChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of what is being checked | [optional] 
**Directory** | Pointer to **string** | Analyzed directories | [optional] 
**File** | Pointer to **string** | Analyzed file path | [optional] 
**Id** | Pointer to **int32** | Policy check ID. A policy contains multiple checks | [optional] 
**PolicyId** | Pointer to **string** | Scanned policy ID | [optional] 
**Process** | Pointer to **string** | Check whether a process is running or not. It&#39;s only returned when the checked process is running | [optional] 
**Rationale** | Pointer to **string** | Explain why this check is necessary | [optional] 
**References** | Pointer to **string** | A link to a documentation page about the check | [optional] 
**Registry** | Pointer to **string** | Analyzed registry | [optional] 
**Remediation** | Pointer to **string** | Explain how to fix the check, this field is very useful in case the check failed | [optional] 
**Result** | Pointer to **string** | Check result | [optional] 
**Title** | Pointer to **string** | A brief description of what is being checked | [optional] 
**Condition** | Pointer to **string** | Specify how rule results are aggregated in order to calculate the final value of a check | [optional] 

## Methods

### NewSCAChecks

`func NewSCAChecks() *SCAChecks`

NewSCAChecks instantiates a new SCAChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCAChecksWithDefaults

`func NewSCAChecksWithDefaults() *SCAChecks`

NewSCAChecksWithDefaults instantiates a new SCAChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SCAChecks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SCAChecks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SCAChecks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SCAChecks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectory

`func (o *SCAChecks) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *SCAChecks) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *SCAChecks) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *SCAChecks) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetFile

`func (o *SCAChecks) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SCAChecks) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SCAChecks) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SCAChecks) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetId

`func (o *SCAChecks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SCAChecks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SCAChecks) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SCAChecks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyId

`func (o *SCAChecks) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SCAChecks) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SCAChecks) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SCAChecks) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetProcess

`func (o *SCAChecks) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *SCAChecks) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *SCAChecks) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *SCAChecks) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetRationale

`func (o *SCAChecks) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *SCAChecks) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *SCAChecks) SetRationale(v string)`

SetRationale sets Rationale field to given value.

### HasRationale

`func (o *SCAChecks) HasRationale() bool`

HasRationale returns a boolean if a field has been set.

### GetReferences

`func (o *SCAChecks) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SCAChecks) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SCAChecks) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SCAChecks) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRegistry

`func (o *SCAChecks) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *SCAChecks) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *SCAChecks) SetRegistry(v string)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *SCAChecks) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetRemediation

`func (o *SCAChecks) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *SCAChecks) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *SCAChecks) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *SCAChecks) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetResult

`func (o *SCAChecks) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SCAChecks) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SCAChecks) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *SCAChecks) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTitle

`func (o *SCAChecks) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SCAChecks) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SCAChecks) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SCAChecks) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCondition

`func (o *SCAChecks) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SCAChecks) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SCAChecks) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SCAChecks) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


