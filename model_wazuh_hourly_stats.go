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

// checks if the WazuhHourlyStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhHourlyStats{}

// WazuhHourlyStats struct for WazuhHourlyStats
type WazuhHourlyStats struct {
	// Array containing the number of alerts for every hour
	Averages []int32 `json:"averages,omitempty"`
	Interactions *int32 `json:"interactions,omitempty"`
}

// NewWazuhHourlyStats instantiates a new WazuhHourlyStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhHourlyStats() *WazuhHourlyStats {
	this := WazuhHourlyStats{}
	return &this
}

// NewWazuhHourlyStatsWithDefaults instantiates a new WazuhHourlyStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhHourlyStatsWithDefaults() *WazuhHourlyStats {
	this := WazuhHourlyStats{}
	return &this
}

// GetAverages returns the Averages field value if set, zero value otherwise.
func (o *WazuhHourlyStats) GetAverages() []int32 {
	if o == nil || IsNil(o.Averages) {
		var ret []int32
		return ret
	}
	return o.Averages
}

// GetAveragesOk returns a tuple with the Averages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhHourlyStats) GetAveragesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Averages) {
		return nil, false
	}
	return o.Averages, true
}

// HasAverages returns a boolean if a field has been set.
func (o *WazuhHourlyStats) HasAverages() bool {
	if o != nil && !IsNil(o.Averages) {
		return true
	}

	return false
}

// SetAverages gets a reference to the given []int32 and assigns it to the Averages field.
func (o *WazuhHourlyStats) SetAverages(v []int32) {
	o.Averages = v
}

// GetInteractions returns the Interactions field value if set, zero value otherwise.
func (o *WazuhHourlyStats) GetInteractions() int32 {
	if o == nil || IsNil(o.Interactions) {
		var ret int32
		return ret
	}
	return *o.Interactions
}

// GetInteractionsOk returns a tuple with the Interactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhHourlyStats) GetInteractionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Interactions) {
		return nil, false
	}
	return o.Interactions, true
}

// HasInteractions returns a boolean if a field has been set.
func (o *WazuhHourlyStats) HasInteractions() bool {
	if o != nil && !IsNil(o.Interactions) {
		return true
	}

	return false
}

// SetInteractions gets a reference to the given int32 and assigns it to the Interactions field.
func (o *WazuhHourlyStats) SetInteractions(v int32) {
	o.Interactions = &v
}

func (o WazuhHourlyStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhHourlyStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Averages) {
		toSerialize["averages"] = o.Averages
	}
	if !IsNil(o.Interactions) {
		toSerialize["interactions"] = o.Interactions
	}
	return toSerialize, nil
}

type NullableWazuhHourlyStats struct {
	value *WazuhHourlyStats
	isSet bool
}

func (v NullableWazuhHourlyStats) Get() *WazuhHourlyStats {
	return v.value
}

func (v *NullableWazuhHourlyStats) Set(val *WazuhHourlyStats) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhHourlyStats) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhHourlyStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhHourlyStats(val *WazuhHourlyStats) *NullableWazuhHourlyStats {
	return &NullableWazuhHourlyStats{value: val, isSet: true}
}

func (v NullableWazuhHourlyStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhHourlyStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


