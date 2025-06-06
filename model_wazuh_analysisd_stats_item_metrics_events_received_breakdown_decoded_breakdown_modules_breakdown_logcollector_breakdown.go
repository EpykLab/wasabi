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

// checks if the WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown{}

// WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown struct for WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown
type WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown struct {
	// EventChannel events collected by logcollector
	Eventchannel *int32 `json:"eventchannel,omitempty"`
	// EventLog events collected by logcollector
	Eventlog *int32 `json:"eventlog,omitempty"`
	// MacOS events collected by logcollector
	Macos *int32 `json:"macos,omitempty"`
	// Other events collected by logcollector
	Others *int32 `json:"others,omitempty"`
}

// NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown {
	this := WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown{}
	return &this
}

// NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown {
	this := WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown{}
	return &this
}

// GetEventchannel returns the Eventchannel field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetEventchannel() int32 {
	if o == nil || IsNil(o.Eventchannel) {
		var ret int32
		return ret
	}
	return *o.Eventchannel
}

// GetEventchannelOk returns a tuple with the Eventchannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetEventchannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Eventchannel) {
		return nil, false
	}
	return o.Eventchannel, true
}

// HasEventchannel returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) HasEventchannel() bool {
	if o != nil && !IsNil(o.Eventchannel) {
		return true
	}

	return false
}

// SetEventchannel gets a reference to the given int32 and assigns it to the Eventchannel field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) SetEventchannel(v int32) {
	o.Eventchannel = &v
}

// GetEventlog returns the Eventlog field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetEventlog() int32 {
	if o == nil || IsNil(o.Eventlog) {
		var ret int32
		return ret
	}
	return *o.Eventlog
}

// GetEventlogOk returns a tuple with the Eventlog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetEventlogOk() (*int32, bool) {
	if o == nil || IsNil(o.Eventlog) {
		return nil, false
	}
	return o.Eventlog, true
}

// HasEventlog returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) HasEventlog() bool {
	if o != nil && !IsNil(o.Eventlog) {
		return true
	}

	return false
}

// SetEventlog gets a reference to the given int32 and assigns it to the Eventlog field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) SetEventlog(v int32) {
	o.Eventlog = &v
}

// GetMacos returns the Macos field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetMacos() int32 {
	if o == nil || IsNil(o.Macos) {
		var ret int32
		return ret
	}
	return *o.Macos
}

// GetMacosOk returns a tuple with the Macos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetMacosOk() (*int32, bool) {
	if o == nil || IsNil(o.Macos) {
		return nil, false
	}
	return o.Macos, true
}

// HasMacos returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) HasMacos() bool {
	if o != nil && !IsNil(o.Macos) {
		return true
	}

	return false
}

// SetMacos gets a reference to the given int32 and assigns it to the Macos field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) SetMacos(v int32) {
	o.Macos = &v
}

// GetOthers returns the Others field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetOthers() int32 {
	if o == nil || IsNil(o.Others) {
		var ret int32
		return ret
	}
	return *o.Others
}

// GetOthersOk returns a tuple with the Others field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) GetOthersOk() (*int32, bool) {
	if o == nil || IsNil(o.Others) {
		return nil, false
	}
	return o.Others, true
}

// HasOthers returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) HasOthers() bool {
	if o != nil && !IsNil(o.Others) {
		return true
	}

	return false
}

// SetOthers gets a reference to the given int32 and assigns it to the Others field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) SetOthers(v int32) {
	o.Others = &v
}

func (o WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Eventchannel) {
		toSerialize["eventchannel"] = o.Eventchannel
	}
	if !IsNil(o.Eventlog) {
		toSerialize["eventlog"] = o.Eventlog
	}
	if !IsNil(o.Macos) {
		toSerialize["macos"] = o.Macos
	}
	if !IsNil(o.Others) {
		toSerialize["others"] = o.Others
	}
	return toSerialize, nil
}

type NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown struct {
	value *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown
	isSet bool
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) Get() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown {
	return v.value
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) Set(val *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown(val *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown {
	return &NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown{value: val, isSet: true}
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


