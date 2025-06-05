# SyscheckDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changes** | Pointer to **int32** | Number of changes applied | [optional] 
**Sha1** | Pointer to **string** | SHA1 checksum of the file | [optional] 
**File** | Pointer to **string** | File name that raised the alert | [optional] 
**Md5** | Pointer to **string** | MD5 checksum of the file | [optional] 
**Inode** | Pointer to **int32** | Inode of the file. Only available in Linux agents | [optional] 
**Uid** | Pointer to **string** | UID of the file | [optional] 
**Date** | Pointer to **time.Time** | Date when the alert was raised | [optional] 
**Perm** | Pointer to **string** | File permissions | [optional] 
**Gname** | Pointer to **string** | Group name. Only available in Linux agents | [optional] 
**Uname** | Pointer to **string** | User name of the file | [optional] 
**Size** | Pointer to **int64** | File size in bytes | [optional] 
**Gid** | Pointer to **string** | GID of the file. Only available in Linux agents | [optional] 
**Mtime** | Pointer to **string** | Last modification date of the file | [optional] 
**Sha256** | Pointer to **string** | SHA256 checksum of the file | [optional] 

## Methods

### NewSyscheckDatabase

`func NewSyscheckDatabase() *SyscheckDatabase`

NewSyscheckDatabase instantiates a new SyscheckDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscheckDatabaseWithDefaults

`func NewSyscheckDatabaseWithDefaults() *SyscheckDatabase`

NewSyscheckDatabaseWithDefaults instantiates a new SyscheckDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanges

`func (o *SyscheckDatabase) GetChanges() int32`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *SyscheckDatabase) GetChangesOk() (*int32, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *SyscheckDatabase) SetChanges(v int32)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *SyscheckDatabase) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetSha1

`func (o *SyscheckDatabase) GetSha1() string`

GetSha1 returns the Sha1 field if non-nil, zero value otherwise.

### GetSha1Ok

`func (o *SyscheckDatabase) GetSha1Ok() (*string, bool)`

GetSha1Ok returns a tuple with the Sha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1

`func (o *SyscheckDatabase) SetSha1(v string)`

SetSha1 sets Sha1 field to given value.

### HasSha1

`func (o *SyscheckDatabase) HasSha1() bool`

HasSha1 returns a boolean if a field has been set.

### GetFile

`func (o *SyscheckDatabase) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SyscheckDatabase) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SyscheckDatabase) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *SyscheckDatabase) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetMd5

`func (o *SyscheckDatabase) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *SyscheckDatabase) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *SyscheckDatabase) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *SyscheckDatabase) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetInode

`func (o *SyscheckDatabase) GetInode() int32`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *SyscheckDatabase) GetInodeOk() (*int32, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *SyscheckDatabase) SetInode(v int32)`

SetInode sets Inode field to given value.

### HasInode

`func (o *SyscheckDatabase) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetUid

`func (o *SyscheckDatabase) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SyscheckDatabase) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SyscheckDatabase) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SyscheckDatabase) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDate

`func (o *SyscheckDatabase) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SyscheckDatabase) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SyscheckDatabase) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *SyscheckDatabase) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPerm

`func (o *SyscheckDatabase) GetPerm() string`

GetPerm returns the Perm field if non-nil, zero value otherwise.

### GetPermOk

`func (o *SyscheckDatabase) GetPermOk() (*string, bool)`

GetPermOk returns a tuple with the Perm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerm

`func (o *SyscheckDatabase) SetPerm(v string)`

SetPerm sets Perm field to given value.

### HasPerm

`func (o *SyscheckDatabase) HasPerm() bool`

HasPerm returns a boolean if a field has been set.

### GetGname

`func (o *SyscheckDatabase) GetGname() string`

GetGname returns the Gname field if non-nil, zero value otherwise.

### GetGnameOk

`func (o *SyscheckDatabase) GetGnameOk() (*string, bool)`

GetGnameOk returns a tuple with the Gname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGname

`func (o *SyscheckDatabase) SetGname(v string)`

SetGname sets Gname field to given value.

### HasGname

`func (o *SyscheckDatabase) HasGname() bool`

HasGname returns a boolean if a field has been set.

### GetUname

`func (o *SyscheckDatabase) GetUname() string`

GetUname returns the Uname field if non-nil, zero value otherwise.

### GetUnameOk

`func (o *SyscheckDatabase) GetUnameOk() (*string, bool)`

GetUnameOk returns a tuple with the Uname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUname

`func (o *SyscheckDatabase) SetUname(v string)`

SetUname sets Uname field to given value.

### HasUname

`func (o *SyscheckDatabase) HasUname() bool`

HasUname returns a boolean if a field has been set.

### GetSize

`func (o *SyscheckDatabase) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SyscheckDatabase) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SyscheckDatabase) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SyscheckDatabase) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetGid

`func (o *SyscheckDatabase) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *SyscheckDatabase) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *SyscheckDatabase) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *SyscheckDatabase) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetMtime

`func (o *SyscheckDatabase) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *SyscheckDatabase) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *SyscheckDatabase) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *SyscheckDatabase) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetSha256

`func (o *SyscheckDatabase) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *SyscheckDatabase) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *SyscheckDatabase) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *SyscheckDatabase) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


