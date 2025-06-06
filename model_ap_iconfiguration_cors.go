/*
Wazuh API REST

The Wazuh API is an open-source RESTful API that allows for interaction with the Wazuh manager from a web browser, command line tools like cURL or any script or program that can make web requests. The Wazuh WUI relies on this heavily and Wazuh’s goal is to accommodate complete remote management of the Wazuh infrastructure via the Wazuh WUI. Use the Wazuh API to easily perform everyday actions like adding an agent, restarting the manager(s) or agent(s) or looking up syscheck details.  # Authentication  Wazuh API endpoints require authentication to be used. Therefore, all calls must include a JSON Web Token. JWT is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. Perform a call with `basicAuth` to `POST /security/user/authenticate` and obtain a JWT token to run any endpoint.  JWT tokens have a default duration of 900 seconds. To change this value, you must perform a call with a valid JWT token to `PUT /security/config`. After this change, you will need to get a new JWT token as all previously issued tokens are revoked when any change is performed on security configuration.  Login with USER and PASSWORD:  `curl -u <USER>:<PASSWORD> -k -X POST \"https://<HOST_IP>:55000/security/user/authenticate\"` ```json {     \"data\": {         \"token\": \"<YOUR_JWT_TOKEN>\"     },     \"error\": 0 } ``` Use the token from the previous response to perform any endpoint request:  `curl -k -X <METHOD> \"https://<HOST_IP>:55000/<ENDPOINT>\" -H  \"Authorization: Bearer <YOUR_JWT_TOKEN>\"`  Change the token base duration:  `curl -k -X PUT \"https://<HOST_IP>:55000/security/config\" -H \"Authorization: Bearer <YOUR_JWT_TOKEN>\" -d '{\"auth_token_exp_timeout\": <NEW_EXPIRE_TIME_IN_SECONDS>}'`  <SecurityDefinitions /> 

API version: 4.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the APIconfigurationCors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIconfigurationCors{}

// APIconfigurationCors struct for APIconfigurationCors
type APIconfigurationCors struct {
	// Enable CORS
	Enabled *bool `json:"enabled,omitempty"`
	// Sources for which the resources will be available. For example 'http://client.example.org'
	SourceRoute *string `json:"source_route,omitempty"`
	// Which headers can be exposed as part of the response
	ExposeHeaders *string `json:"expose_headers,omitempty"`
	// Which HTTP headers can be used during the actual request
	AllowHeaders *string `json:"allow_headers,omitempty"`
	// Browsers will only expose the response to frontend JavaScript code if this is enabled
	AllowCredentials *bool `json:"allow_credentials,omitempty"`
}

// NewAPIconfigurationCors instantiates a new APIconfigurationCors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIconfigurationCors() *APIconfigurationCors {
	this := APIconfigurationCors{}
	var enabled bool = false
	this.Enabled = &enabled
	var allowCredentials bool = false
	this.AllowCredentials = &allowCredentials
	return &this
}

// NewAPIconfigurationCorsWithDefaults instantiates a new APIconfigurationCors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIconfigurationCorsWithDefaults() *APIconfigurationCors {
	this := APIconfigurationCors{}
	var enabled bool = false
	this.Enabled = &enabled
	var allowCredentials bool = false
	this.AllowCredentials = &allowCredentials
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *APIconfigurationCors) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfigurationCors) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *APIconfigurationCors) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *APIconfigurationCors) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSourceRoute returns the SourceRoute field value if set, zero value otherwise.
func (o *APIconfigurationCors) GetSourceRoute() string {
	if o == nil || IsNil(o.SourceRoute) {
		var ret string
		return ret
	}
	return *o.SourceRoute
}

// GetSourceRouteOk returns a tuple with the SourceRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfigurationCors) GetSourceRouteOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRoute) {
		return nil, false
	}
	return o.SourceRoute, true
}

// HasSourceRoute returns a boolean if a field has been set.
func (o *APIconfigurationCors) HasSourceRoute() bool {
	if o != nil && !IsNil(o.SourceRoute) {
		return true
	}

	return false
}

// SetSourceRoute gets a reference to the given string and assigns it to the SourceRoute field.
func (o *APIconfigurationCors) SetSourceRoute(v string) {
	o.SourceRoute = &v
}

// GetExposeHeaders returns the ExposeHeaders field value if set, zero value otherwise.
func (o *APIconfigurationCors) GetExposeHeaders() string {
	if o == nil || IsNil(o.ExposeHeaders) {
		var ret string
		return ret
	}
	return *o.ExposeHeaders
}

// GetExposeHeadersOk returns a tuple with the ExposeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfigurationCors) GetExposeHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.ExposeHeaders) {
		return nil, false
	}
	return o.ExposeHeaders, true
}

// HasExposeHeaders returns a boolean if a field has been set.
func (o *APIconfigurationCors) HasExposeHeaders() bool {
	if o != nil && !IsNil(o.ExposeHeaders) {
		return true
	}

	return false
}

// SetExposeHeaders gets a reference to the given string and assigns it to the ExposeHeaders field.
func (o *APIconfigurationCors) SetExposeHeaders(v string) {
	o.ExposeHeaders = &v
}

// GetAllowHeaders returns the AllowHeaders field value if set, zero value otherwise.
func (o *APIconfigurationCors) GetAllowHeaders() string {
	if o == nil || IsNil(o.AllowHeaders) {
		var ret string
		return ret
	}
	return *o.AllowHeaders
}

// GetAllowHeadersOk returns a tuple with the AllowHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfigurationCors) GetAllowHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AllowHeaders) {
		return nil, false
	}
	return o.AllowHeaders, true
}

// HasAllowHeaders returns a boolean if a field has been set.
func (o *APIconfigurationCors) HasAllowHeaders() bool {
	if o != nil && !IsNil(o.AllowHeaders) {
		return true
	}

	return false
}

// SetAllowHeaders gets a reference to the given string and assigns it to the AllowHeaders field.
func (o *APIconfigurationCors) SetAllowHeaders(v string) {
	o.AllowHeaders = &v
}

// GetAllowCredentials returns the AllowCredentials field value if set, zero value otherwise.
func (o *APIconfigurationCors) GetAllowCredentials() bool {
	if o == nil || IsNil(o.AllowCredentials) {
		var ret bool
		return ret
	}
	return *o.AllowCredentials
}

// GetAllowCredentialsOk returns a tuple with the AllowCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfigurationCors) GetAllowCredentialsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCredentials) {
		return nil, false
	}
	return o.AllowCredentials, true
}

// HasAllowCredentials returns a boolean if a field has been set.
func (o *APIconfigurationCors) HasAllowCredentials() bool {
	if o != nil && !IsNil(o.AllowCredentials) {
		return true
	}

	return false
}

// SetAllowCredentials gets a reference to the given bool and assigns it to the AllowCredentials field.
func (o *APIconfigurationCors) SetAllowCredentials(v bool) {
	o.AllowCredentials = &v
}

func (o APIconfigurationCors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIconfigurationCors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.SourceRoute) {
		toSerialize["source_route"] = o.SourceRoute
	}
	if !IsNil(o.ExposeHeaders) {
		toSerialize["expose_headers"] = o.ExposeHeaders
	}
	if !IsNil(o.AllowHeaders) {
		toSerialize["allow_headers"] = o.AllowHeaders
	}
	if !IsNil(o.AllowCredentials) {
		toSerialize["allow_credentials"] = o.AllowCredentials
	}
	return toSerialize, nil
}

type NullableAPIconfigurationCors struct {
	value *APIconfigurationCors
	isSet bool
}

func (v NullableAPIconfigurationCors) Get() *APIconfigurationCors {
	return v.value
}

func (v *NullableAPIconfigurationCors) Set(val *APIconfigurationCors) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIconfigurationCors) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIconfigurationCors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIconfigurationCors(val *APIconfigurationCors) *NullableAPIconfigurationCors {
	return &NullableAPIconfigurationCors{value: val, isSet: true}
}

func (v NullableAPIconfigurationCors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIconfigurationCors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


