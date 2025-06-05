# RuleFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Name of the file | [optional] 
**RelativeDirname** | Pointer to **string** | Folder path where the file is located. This path is relative to the Wazuh installation path | [optional] 
**Status** | Pointer to **string** | Whether the specified ruleset file is enabled or disabled in Wazuh manager configuration | [optional] 

## Methods

### NewRuleFile

`func NewRuleFile() *RuleFile`

NewRuleFile instantiates a new RuleFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleFileWithDefaults

`func NewRuleFileWithDefaults() *RuleFile`

NewRuleFileWithDefaults instantiates a new RuleFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *RuleFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RuleFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RuleFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *RuleFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetRelativeDirname

`func (o *RuleFile) GetRelativeDirname() string`

GetRelativeDirname returns the RelativeDirname field if non-nil, zero value otherwise.

### GetRelativeDirnameOk

`func (o *RuleFile) GetRelativeDirnameOk() (*string, bool)`

GetRelativeDirnameOk returns a tuple with the RelativeDirname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirname

`func (o *RuleFile) SetRelativeDirname(v string)`

SetRelativeDirname sets RelativeDirname field to given value.

### HasRelativeDirname

`func (o *RuleFile) HasRelativeDirname() bool`

HasRelativeDirname returns a boolean if a field has been set.

### GetStatus

`func (o *RuleFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RuleFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RuleFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RuleFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


