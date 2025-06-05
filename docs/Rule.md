# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Name of the file | [optional] 
**RelativeDirname** | Pointer to **string** | Folder path where the file is located. This path is relative to the Wazuh installation path | [optional] 
**Status** | Pointer to **string** | Whether the specified ruleset file is enabled or disabled in Wazuh manager configuration | [optional] 
**Description** | Pointer to **string** | Rule description. This description is shown when an alert matching the rule is raised | [optional] 
**Details** | Pointer to **map[string]interface{}** | Rule definition details | [optional] 
**Gdpr** | Pointer to **[]string** | GDPR checks the rule is checking | [optional] 
**Gpg13** | Pointer to **[]string** | GPG13 checks the rule is checking | [optional] 
**Groups** | Pointer to **[]string** | Groups the rule belongs to | [optional] 
**Hipaa** | Pointer to **[]string** | HIPAA checks the rule is checking | [optional] 
**Id** | Pointer to **int32** | Rule ID | [optional] 
**Level** | Pointer to **int32** | Rule level | [optional] 
**Nist80053** | Pointer to **[]string** | NIST-800-53 checks the rule is checking | [optional] 
**Tsc** | Pointer to **[]string** | TSC checks the rule is checking | [optional] 
**Pci** | Pointer to **[]string** | PCI DSS checks the rule is checking | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *Rule) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Rule) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Rule) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Rule) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetRelativeDirname

`func (o *Rule) GetRelativeDirname() string`

GetRelativeDirname returns the RelativeDirname field if non-nil, zero value otherwise.

### GetRelativeDirnameOk

`func (o *Rule) GetRelativeDirnameOk() (*string, bool)`

GetRelativeDirnameOk returns a tuple with the RelativeDirname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirname

`func (o *Rule) SetRelativeDirname(v string)`

SetRelativeDirname sets RelativeDirname field to given value.

### HasRelativeDirname

`func (o *Rule) HasRelativeDirname() bool`

HasRelativeDirname returns a boolean if a field has been set.

### GetStatus

`func (o *Rule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Rule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Rule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Rule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *Rule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Rule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Rule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *Rule) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Rule) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Rule) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Rule) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGdpr

`func (o *Rule) GetGdpr() []string`

GetGdpr returns the Gdpr field if non-nil, zero value otherwise.

### GetGdprOk

`func (o *Rule) GetGdprOk() (*[]string, bool)`

GetGdprOk returns a tuple with the Gdpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGdpr

`func (o *Rule) SetGdpr(v []string)`

SetGdpr sets Gdpr field to given value.

### HasGdpr

`func (o *Rule) HasGdpr() bool`

HasGdpr returns a boolean if a field has been set.

### GetGpg13

`func (o *Rule) GetGpg13() []string`

GetGpg13 returns the Gpg13 field if non-nil, zero value otherwise.

### GetGpg13Ok

`func (o *Rule) GetGpg13Ok() (*[]string, bool)`

GetGpg13Ok returns a tuple with the Gpg13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpg13

`func (o *Rule) SetGpg13(v []string)`

SetGpg13 sets Gpg13 field to given value.

### HasGpg13

`func (o *Rule) HasGpg13() bool`

HasGpg13 returns a boolean if a field has been set.

### GetGroups

`func (o *Rule) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Rule) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Rule) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Rule) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHipaa

`func (o *Rule) GetHipaa() []string`

GetHipaa returns the Hipaa field if non-nil, zero value otherwise.

### GetHipaaOk

`func (o *Rule) GetHipaaOk() (*[]string, bool)`

GetHipaaOk returns a tuple with the Hipaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHipaa

`func (o *Rule) SetHipaa(v []string)`

SetHipaa sets Hipaa field to given value.

### HasHipaa

`func (o *Rule) HasHipaa() bool`

HasHipaa returns a boolean if a field has been set.

### GetId

`func (o *Rule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *Rule) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Rule) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Rule) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Rule) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetNist80053

`func (o *Rule) GetNist80053() []string`

GetNist80053 returns the Nist80053 field if non-nil, zero value otherwise.

### GetNist80053Ok

`func (o *Rule) GetNist80053Ok() (*[]string, bool)`

GetNist80053Ok returns a tuple with the Nist80053 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNist80053

`func (o *Rule) SetNist80053(v []string)`

SetNist80053 sets Nist80053 field to given value.

### HasNist80053

`func (o *Rule) HasNist80053() bool`

HasNist80053 returns a boolean if a field has been set.

### GetTsc

`func (o *Rule) GetTsc() []string`

GetTsc returns the Tsc field if non-nil, zero value otherwise.

### GetTscOk

`func (o *Rule) GetTscOk() (*[]string, bool)`

GetTscOk returns a tuple with the Tsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsc

`func (o *Rule) SetTsc(v []string)`

SetTsc sets Tsc field to given value.

### HasTsc

`func (o *Rule) HasTsc() bool`

HasTsc returns a boolean if a field has been set.

### GetPci

`func (o *Rule) GetPci() []string`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *Rule) GetPciOk() (*[]string, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *Rule) SetPci(v []string)`

SetPci sets Pci field to given value.

### HasPci

`func (o *Rule) HasPci() bool`

HasPci returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


