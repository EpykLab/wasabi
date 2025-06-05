# Decoder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Decoder name | [optional] 
**Position** | Pointer to **int32** | Position of this decoder in the decoder file. The parent decoder will have position 0, the following defined decoder will have position 1, and so on | [optional] 
**Details** | Pointer to **map[string]interface{}** | Decoder definition fields | [optional] 
**Filename** | Pointer to **string** | Name of the file | [optional] 
**RelativeDirname** | Pointer to **string** | Folder path where the file is located. This path is relative to the Wazuh installation path | [optional] 
**Status** | Pointer to **string** | Whether the specified ruleset file is enabled or disabled in Wazuh manager configuration | [optional] 

## Methods

### NewDecoder

`func NewDecoder() *Decoder`

NewDecoder instantiates a new Decoder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecoderWithDefaults

`func NewDecoderWithDefaults() *Decoder`

NewDecoderWithDefaults instantiates a new Decoder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Decoder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Decoder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Decoder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Decoder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPosition

`func (o *Decoder) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Decoder) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Decoder) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Decoder) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetDetails

`func (o *Decoder) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Decoder) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Decoder) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Decoder) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFilename

`func (o *Decoder) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Decoder) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Decoder) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Decoder) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetRelativeDirname

`func (o *Decoder) GetRelativeDirname() string`

GetRelativeDirname returns the RelativeDirname field if non-nil, zero value otherwise.

### GetRelativeDirnameOk

`func (o *Decoder) GetRelativeDirnameOk() (*string, bool)`

GetRelativeDirnameOk returns a tuple with the RelativeDirname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeDirname

`func (o *Decoder) SetRelativeDirname(v string)`

SetRelativeDirname sets RelativeDirname field to given value.

### HasRelativeDirname

`func (o *Decoder) HasRelativeDirname() bool`

HasRelativeDirname returns a boolean if a field has been set.

### GetStatus

`func (o *Decoder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Decoder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Decoder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Decoder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


