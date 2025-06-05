# GroupFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | File name | [optional] 
**Hash** | Pointer to **string** | File content hash | [optional] 

## Methods

### NewGroupFiles

`func NewGroupFiles() *GroupFiles`

NewGroupFiles instantiates a new GroupFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupFilesWithDefaults

`func NewGroupFilesWithDefaults() *GroupFiles`

NewGroupFilesWithDefaults instantiates a new GroupFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *GroupFiles) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GroupFiles) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GroupFiles) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *GroupFiles) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetHash

`func (o *GroupFiles) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GroupFiles) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GroupFiles) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GroupFiles) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


