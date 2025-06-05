# UsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | User&#39;s id | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AllowRunAs** | Pointer to **bool** | Flag to enable the user to log in using authorization context | [optional] 
**Roles** | Pointer to **[]string** | User&#39;s roles | [optional] 

## Methods

### NewUsersResponse

`func NewUsersResponse() *UsersResponse`

NewUsersResponse instantiates a new UsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponseWithDefaults

`func NewUsersResponseWithDefaults() *UsersResponse`

NewUsersResponseWithDefaults instantiates a new UsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UsersResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UsersResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsersResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsersResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsersResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAllowRunAs

`func (o *UsersResponse) GetAllowRunAs() bool`

GetAllowRunAs returns the AllowRunAs field if non-nil, zero value otherwise.

### GetAllowRunAsOk

`func (o *UsersResponse) GetAllowRunAsOk() (*bool, bool)`

GetAllowRunAsOk returns a tuple with the AllowRunAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRunAs

`func (o *UsersResponse) SetAllowRunAs(v bool)`

SetAllowRunAs sets AllowRunAs field to given value.

### HasAllowRunAs

`func (o *UsersResponse) HasAllowRunAs() bool`

HasAllowRunAs returns a boolean if a field has been set.

### GetRoles

`func (o *UsersResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UsersResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UsersResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UsersResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


