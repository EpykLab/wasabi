# APIconfigurationAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLoginAttempts** | Pointer to **int32** | Maximum number of login attempts in {block_time} seconds | [optional] 
**BlockTime** | Pointer to **int32** | Blocking time for IPs that have exceeded {max_login_attempts}. Time counts from the first attempt | [optional] 
**MaxRequestPerMinute** | Pointer to **int32** | Maximum number of requests allowed per minute | [optional] 

## Methods

### NewAPIconfigurationAccess

`func NewAPIconfigurationAccess() *APIconfigurationAccess`

NewAPIconfigurationAccess instantiates a new APIconfigurationAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIconfigurationAccessWithDefaults

`func NewAPIconfigurationAccessWithDefaults() *APIconfigurationAccess`

NewAPIconfigurationAccessWithDefaults instantiates a new APIconfigurationAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLoginAttempts

`func (o *APIconfigurationAccess) GetMaxLoginAttempts() int32`

GetMaxLoginAttempts returns the MaxLoginAttempts field if non-nil, zero value otherwise.

### GetMaxLoginAttemptsOk

`func (o *APIconfigurationAccess) GetMaxLoginAttemptsOk() (*int32, bool)`

GetMaxLoginAttemptsOk returns a tuple with the MaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoginAttempts

`func (o *APIconfigurationAccess) SetMaxLoginAttempts(v int32)`

SetMaxLoginAttempts sets MaxLoginAttempts field to given value.

### HasMaxLoginAttempts

`func (o *APIconfigurationAccess) HasMaxLoginAttempts() bool`

HasMaxLoginAttempts returns a boolean if a field has been set.

### GetBlockTime

`func (o *APIconfigurationAccess) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *APIconfigurationAccess) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *APIconfigurationAccess) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *APIconfigurationAccess) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetMaxRequestPerMinute

`func (o *APIconfigurationAccess) GetMaxRequestPerMinute() int32`

GetMaxRequestPerMinute returns the MaxRequestPerMinute field if non-nil, zero value otherwise.

### GetMaxRequestPerMinuteOk

`func (o *APIconfigurationAccess) GetMaxRequestPerMinuteOk() (*int32, bool)`

GetMaxRequestPerMinuteOk returns a tuple with the MaxRequestPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestPerMinute

`func (o *APIconfigurationAccess) SetMaxRequestPerMinute(v int32)`

SetMaxRequestPerMinute sets MaxRequestPerMinute field to given value.

### HasMaxRequestPerMinute

`func (o *APIconfigurationAccess) HasMaxRequestPerMinute() bool`

HasMaxRequestPerMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


