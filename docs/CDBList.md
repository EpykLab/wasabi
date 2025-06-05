# CDBList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Name of the file | [optional] 
**RelativeDirname** | Pointer to **string** | Folder path where the file is located. This path is relative to the Wazuh installation path | [optional] 
**Items** | Pointer to [**[]CDBListPair**](CDBListPair.md) |  | [optional] 

## Methods

### NewCDBList

`func NewCDBList() *CDBList`

NewCDBList instantiates a new CDBList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDBListWithDefaults

`func NewCDBListWithDefaults() *CDBList`

NewCDBListWithDefaults instantiates a new CDBList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *CDBList) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CDBList) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CDBList) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CDBList) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetRelativeDirname

`func (o *CDBList) GetRelativeDirname() string`

GetRelativeDirname returns the RelativeDirname field if non-nil, zero value otherwise.

### GetRelativeDirnameOk

`func (o *CDBList) GetRelativeDirnameOk() (*string, bool)`

GetRelativeDirnameOk returns a tuple with the RelativeDirname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirname

`func (o *CDBList) SetRelativeDirname(v string)`

SetRelativeDirname sets RelativeDirname field to given value.

### HasRelativeDirname

`func (o *CDBList) HasRelativeDirname() bool`

HasRelativeDirname returns a boolean if a field has been set.

### GetItems

`func (o *CDBList) GetItems() []CDBListPair`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CDBList) GetItemsOk() (*[]CDBListPair, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CDBList) SetItems(v []CDBListPair)`

SetItems sets Items field to given value.

### HasItems

`func (o *CDBList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


