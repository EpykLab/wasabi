# APIconfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**APIconfigurationAccess**](APIconfigurationAccess.md) |  | [optional] 
**Logs** | Pointer to [**APIconfigurationLogs**](APIconfigurationLogs.md) |  | [optional] 
**Cors** | Pointer to [**APIconfigurationCors**](APIconfigurationCors.md) |  | [optional] 
**ExperimentalFeatures** | Pointer to **bool** | Enable features under development | [optional] [default to false]
**AuthenticationPoolSize** | Pointer to **int32** | Number of processes dedicated to processing authentication requests | [optional] [default to 2]

## Methods

### NewAPIconfiguration

`func NewAPIconfiguration() *APIconfiguration`

NewAPIconfiguration instantiates a new APIconfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIconfigurationWithDefaults

`func NewAPIconfigurationWithDefaults() *APIconfiguration`

NewAPIconfigurationWithDefaults instantiates a new APIconfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *APIconfiguration) GetAccess() APIconfigurationAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *APIconfiguration) GetAccessOk() (*APIconfigurationAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *APIconfiguration) SetAccess(v APIconfigurationAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *APIconfiguration) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetLogs

`func (o *APIconfiguration) GetLogs() APIconfigurationLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *APIconfiguration) GetLogsOk() (*APIconfigurationLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *APIconfiguration) SetLogs(v APIconfigurationLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *APIconfiguration) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetCors

`func (o *APIconfiguration) GetCors() APIconfigurationCors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *APIconfiguration) GetCorsOk() (*APIconfigurationCors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *APIconfiguration) SetCors(v APIconfigurationCors)`

SetCors sets Cors field to given value.

### HasCors

`func (o *APIconfiguration) HasCors() bool`

HasCors returns a boolean if a field has been set.

### GetExperimentalFeatures

`func (o *APIconfiguration) GetExperimentalFeatures() bool`

GetExperimentalFeatures returns the ExperimentalFeatures field if non-nil, zero value otherwise.

### GetExperimentalFeaturesOk

`func (o *APIconfiguration) GetExperimentalFeaturesOk() (*bool, bool)`

GetExperimentalFeaturesOk returns a tuple with the ExperimentalFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentalFeatures

`func (o *APIconfiguration) SetExperimentalFeatures(v bool)`

SetExperimentalFeatures sets ExperimentalFeatures field to given value.

### HasExperimentalFeatures

`func (o *APIconfiguration) HasExperimentalFeatures() bool`

HasExperimentalFeatures returns a boolean if a field has been set.

### GetAuthenticationPoolSize

`func (o *APIconfiguration) GetAuthenticationPoolSize() int32`

GetAuthenticationPoolSize returns the AuthenticationPoolSize field if non-nil, zero value otherwise.

### GetAuthenticationPoolSizeOk

`func (o *APIconfiguration) GetAuthenticationPoolSizeOk() (*int32, bool)`

GetAuthenticationPoolSizeOk returns a tuple with the AuthenticationPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPoolSize

`func (o *APIconfiguration) SetAuthenticationPoolSize(v int32)`

SetAuthenticationPoolSize sets AuthenticationPoolSize field to given value.

### HasAuthenticationPoolSize

`func (o *APIconfiguration) HasAuthenticationPoolSize() bool`

HasAuthenticationPoolSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


