# SecurityConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTokenExpTimeout** | Pointer to **int32** | Time in seconds until the token expires | [optional] 
**RbacMode** | Pointer to **string** | RBAC mode (white/black) | [optional] 

## Methods

### NewSecurityConfiguration

`func NewSecurityConfiguration() *SecurityConfiguration`

NewSecurityConfiguration instantiates a new SecurityConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityConfigurationWithDefaults

`func NewSecurityConfigurationWithDefaults() *SecurityConfiguration`

NewSecurityConfigurationWithDefaults instantiates a new SecurityConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTokenExpTimeout

`func (o *SecurityConfiguration) GetAuthTokenExpTimeout() int32`

GetAuthTokenExpTimeout returns the AuthTokenExpTimeout field if non-nil, zero value otherwise.

### GetAuthTokenExpTimeoutOk

`func (o *SecurityConfiguration) GetAuthTokenExpTimeoutOk() (*int32, bool)`

GetAuthTokenExpTimeoutOk returns a tuple with the AuthTokenExpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTokenExpTimeout

`func (o *SecurityConfiguration) SetAuthTokenExpTimeout(v int32)`

SetAuthTokenExpTimeout sets AuthTokenExpTimeout field to given value.

### HasAuthTokenExpTimeout

`func (o *SecurityConfiguration) HasAuthTokenExpTimeout() bool`

HasAuthTokenExpTimeout returns a boolean if a field has been set.

### GetRbacMode

`func (o *SecurityConfiguration) GetRbacMode() string`

GetRbacMode returns the RbacMode field if non-nil, zero value otherwise.

### GetRbacModeOk

`func (o *SecurityConfiguration) GetRbacModeOk() (*string, bool)`

GetRbacModeOk returns a tuple with the RbacMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacMode

`func (o *SecurityConfiguration) SetRbacMode(v string)`

SetRbacMode sets RbacMode field to given value.

### HasRbacMode

`func (o *SecurityConfiguration) HasRbacMode() bool`

HasRbacMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


