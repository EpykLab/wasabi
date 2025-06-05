# APIconfigurationCors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable CORS | [optional] [default to false]
**SourceRoute** | Pointer to **string** | Sources for which the resources will be available. For example &#39;http://client.example.org&#39; | [optional] 
**ExposeHeaders** | Pointer to **string** | Which headers can be exposed as part of the response | [optional] 
**AllowHeaders** | Pointer to **string** | Which HTTP headers can be used during the actual request | [optional] 
**AllowCredentials** | Pointer to **bool** | Browsers will only expose the response to frontend JavaScript code if this is enabled | [optional] [default to false]

## Methods

### NewAPIconfigurationCors

`func NewAPIconfigurationCors() *APIconfigurationCors`

NewAPIconfigurationCors instantiates a new APIconfigurationCors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIconfigurationCorsWithDefaults

`func NewAPIconfigurationCorsWithDefaults() *APIconfigurationCors`

NewAPIconfigurationCorsWithDefaults instantiates a new APIconfigurationCors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *APIconfigurationCors) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *APIconfigurationCors) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *APIconfigurationCors) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *APIconfigurationCors) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSourceRoute

`func (o *APIconfigurationCors) GetSourceRoute() string`

GetSourceRoute returns the SourceRoute field if non-nil, zero value otherwise.

### GetSourceRouteOk

`func (o *APIconfigurationCors) GetSourceRouteOk() (*string, bool)`

GetSourceRouteOk returns a tuple with the SourceRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRoute

`func (o *APIconfigurationCors) SetSourceRoute(v string)`

SetSourceRoute sets SourceRoute field to given value.

### HasSourceRoute

`func (o *APIconfigurationCors) HasSourceRoute() bool`

HasSourceRoute returns a boolean if a field has been set.

### GetExposeHeaders

`func (o *APIconfigurationCors) GetExposeHeaders() string`

GetExposeHeaders returns the ExposeHeaders field if non-nil, zero value otherwise.

### GetExposeHeadersOk

`func (o *APIconfigurationCors) GetExposeHeadersOk() (*string, bool)`

GetExposeHeadersOk returns a tuple with the ExposeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeHeaders

`func (o *APIconfigurationCors) SetExposeHeaders(v string)`

SetExposeHeaders sets ExposeHeaders field to given value.

### HasExposeHeaders

`func (o *APIconfigurationCors) HasExposeHeaders() bool`

HasExposeHeaders returns a boolean if a field has been set.

### GetAllowHeaders

`func (o *APIconfigurationCors) GetAllowHeaders() string`

GetAllowHeaders returns the AllowHeaders field if non-nil, zero value otherwise.

### GetAllowHeadersOk

`func (o *APIconfigurationCors) GetAllowHeadersOk() (*string, bool)`

GetAllowHeadersOk returns a tuple with the AllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHeaders

`func (o *APIconfigurationCors) SetAllowHeaders(v string)`

SetAllowHeaders sets AllowHeaders field to given value.

### HasAllowHeaders

`func (o *APIconfigurationCors) HasAllowHeaders() bool`

HasAllowHeaders returns a boolean if a field has been set.

### GetAllowCredentials

`func (o *APIconfigurationCors) GetAllowCredentials() bool`

GetAllowCredentials returns the AllowCredentials field if non-nil, zero value otherwise.

### GetAllowCredentialsOk

`func (o *APIconfigurationCors) GetAllowCredentialsOk() (*bool, bool)`

GetAllowCredentialsOk returns a tuple with the AllowCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCredentials

`func (o *APIconfigurationCors) SetAllowCredentials(v bool)`

SetAllowCredentials sets AllowCredentials field to given value.

### HasAllowCredentials

`func (o *APIconfigurationCors) HasAllowCredentials() bool`

HasAllowCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


