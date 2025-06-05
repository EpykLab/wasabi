# RulesetFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Name of the file | [optional] 
**RelativeDirname** | Pointer to **string** | Folder path where the file is located. This path is relative to the Wazuh installation path | [optional] 

## Methods

### NewRulesetFile

`func NewRulesetFile() *RulesetFile`

NewRulesetFile instantiates a new RulesetFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetFileWithDefaults

`func NewRulesetFileWithDefaults() *RulesetFile`

NewRulesetFileWithDefaults instantiates a new RulesetFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *RulesetFile) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RulesetFile) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RulesetFile) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *RulesetFile) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetRelativeDirname

`func (o *RulesetFile) GetRelativeDirname() string`

GetRelativeDirname returns the RelativeDirname field if non-nil, zero value otherwise.

### GetRelativeDirnameOk

`func (o *RulesetFile) GetRelativeDirnameOk() (*string, bool)`

GetRelativeDirnameOk returns a tuple with the RelativeDirname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirname

`func (o *RulesetFile) SetRelativeDirname(v string)`

SetRelativeDirname sets RelativeDirname field to given value.

### HasRelativeDirname

`func (o *RulesetFile) HasRelativeDirname() bool`

HasRelativeDirname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


