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

// checks if the APIconfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIconfiguration{}

// APIconfiguration struct for APIconfiguration
type APIconfiguration struct {
	Access *APIconfigurationAccess `json:"access,omitempty"`
	Logs *APIconfigurationLogs `json:"logs,omitempty"`
	Cors *APIconfigurationCors `json:"cors,omitempty"`
	// Enable features under development
	ExperimentalFeatures *bool `json:"experimental_features,omitempty"`
	// Number of processes dedicated to processing authentication requests
	AuthenticationPoolSize *int32 `json:"authentication_pool_size,omitempty"`
}

// NewAPIconfiguration instantiates a new APIconfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIconfiguration() *APIconfiguration {
	this := APIconfiguration{}
	var experimentalFeatures bool = false
	this.ExperimentalFeatures = &experimentalFeatures
	var authenticationPoolSize int32 = 2
	this.AuthenticationPoolSize = &authenticationPoolSize
	return &this
}

// NewAPIconfigurationWithDefaults instantiates a new APIconfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIconfigurationWithDefaults() *APIconfiguration {
	this := APIconfiguration{}
	var experimentalFeatures bool = false
	this.ExperimentalFeatures = &experimentalFeatures
	var authenticationPoolSize int32 = 2
	this.AuthenticationPoolSize = &authenticationPoolSize
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *APIconfiguration) GetAccess() APIconfigurationAccess {
	if o == nil || IsNil(o.Access) {
		var ret APIconfigurationAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfiguration) GetAccessOk() (*APIconfigurationAccess, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *APIconfiguration) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given APIconfigurationAccess and assigns it to the Access field.
func (o *APIconfiguration) SetAccess(v APIconfigurationAccess) {
	o.Access = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *APIconfiguration) GetLogs() APIconfigurationLogs {
	if o == nil || IsNil(o.Logs) {
		var ret APIconfigurationLogs
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfiguration) GetLogsOk() (*APIconfigurationLogs, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *APIconfiguration) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given APIconfigurationLogs and assigns it to the Logs field.
func (o *APIconfiguration) SetLogs(v APIconfigurationLogs) {
	o.Logs = &v
}

// GetCors returns the Cors field value if set, zero value otherwise.
func (o *APIconfiguration) GetCors() APIconfigurationCors {
	if o == nil || IsNil(o.Cors) {
		var ret APIconfigurationCors
		return ret
	}
	return *o.Cors
}

// GetCorsOk returns a tuple with the Cors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfiguration) GetCorsOk() (*APIconfigurationCors, bool) {
	if o == nil || IsNil(o.Cors) {
		return nil, false
	}
	return o.Cors, true
}

// HasCors returns a boolean if a field has been set.
func (o *APIconfiguration) HasCors() bool {
	if o != nil && !IsNil(o.Cors) {
		return true
	}

	return false
}

// SetCors gets a reference to the given APIconfigurationCors and assigns it to the Cors field.
func (o *APIconfiguration) SetCors(v APIconfigurationCors) {
	o.Cors = &v
}

// GetExperimentalFeatures returns the ExperimentalFeatures field value if set, zero value otherwise.
func (o *APIconfiguration) GetExperimentalFeatures() bool {
	if o == nil || IsNil(o.ExperimentalFeatures) {
		var ret bool
		return ret
	}
	return *o.ExperimentalFeatures
}

// GetExperimentalFeaturesOk returns a tuple with the ExperimentalFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfiguration) GetExperimentalFeaturesOk() (*bool, bool) {
	if o == nil || IsNil(o.ExperimentalFeatures) {
		return nil, false
	}
	return o.ExperimentalFeatures, true
}

// HasExperimentalFeatures returns a boolean if a field has been set.
func (o *APIconfiguration) HasExperimentalFeatures() bool {
	if o != nil && !IsNil(o.ExperimentalFeatures) {
		return true
	}

	return false
}

// SetExperimentalFeatures gets a reference to the given bool and assigns it to the ExperimentalFeatures field.
func (o *APIconfiguration) SetExperimentalFeatures(v bool) {
	o.ExperimentalFeatures = &v
}

// GetAuthenticationPoolSize returns the AuthenticationPoolSize field value if set, zero value otherwise.
func (o *APIconfiguration) GetAuthenticationPoolSize() int32 {
	if o == nil || IsNil(o.AuthenticationPoolSize) {
		var ret int32
		return ret
	}
	return *o.AuthenticationPoolSize
}

// GetAuthenticationPoolSizeOk returns a tuple with the AuthenticationPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIconfiguration) GetAuthenticationPoolSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.AuthenticationPoolSize) {
		return nil, false
	}
	return o.AuthenticationPoolSize, true
}

// HasAuthenticationPoolSize returns a boolean if a field has been set.
func (o *APIconfiguration) HasAuthenticationPoolSize() bool {
	if o != nil && !IsNil(o.AuthenticationPoolSize) {
		return true
	}

	return false
}

// SetAuthenticationPoolSize gets a reference to the given int32 and assigns it to the AuthenticationPoolSize field.
func (o *APIconfiguration) SetAuthenticationPoolSize(v int32) {
	o.AuthenticationPoolSize = &v
}

func (o APIconfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIconfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Cors) {
		toSerialize["cors"] = o.Cors
	}
	if !IsNil(o.ExperimentalFeatures) {
		toSerialize["experimental_features"] = o.ExperimentalFeatures
	}
	if !IsNil(o.AuthenticationPoolSize) {
		toSerialize["authentication_pool_size"] = o.AuthenticationPoolSize
	}
	return toSerialize, nil
}

type NullableAPIconfiguration struct {
	value *APIconfiguration
	isSet bool
}

func (v NullableAPIconfiguration) Get() *APIconfiguration {
	return v.value
}

func (v *NullableAPIconfiguration) Set(val *APIconfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIconfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIconfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIconfiguration(val *APIconfiguration) *NullableAPIconfiguration {
	return &NullableAPIconfiguration{value: val, isSet: true}
}

func (v NullableAPIconfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIconfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


