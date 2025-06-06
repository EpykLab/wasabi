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

// checks if the WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown{}

// WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown struct for WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown
type WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown struct {
	DecodedBreakdown *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown `json:"decoded_breakdown,omitempty"`
}

// NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown {
	this := WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown{}
	return &this
}

// NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownWithDefaults() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown {
	this := WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown{}
	return &this
}

// GetDecodedBreakdown returns the DecodedBreakdown field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) GetDecodedBreakdown() WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown {
	if o == nil || IsNil(o.DecodedBreakdown) {
		var ret WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown
		return ret
	}
	return *o.DecodedBreakdown
}

// GetDecodedBreakdownOk returns a tuple with the DecodedBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) GetDecodedBreakdownOk() (*WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown, bool) {
	if o == nil || IsNil(o.DecodedBreakdown) {
		return nil, false
	}
	return o.DecodedBreakdown, true
}

// HasDecodedBreakdown returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) HasDecodedBreakdown() bool {
	if o != nil && !IsNil(o.DecodedBreakdown) {
		return true
	}

	return false
}

// SetDecodedBreakdown gets a reference to the given WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown and assigns it to the DecodedBreakdown field.
func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) SetDecodedBreakdown(v WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) {
	o.DecodedBreakdown = &v
}

func (o WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DecodedBreakdown) {
		toSerialize["decoded_breakdown"] = o.DecodedBreakdown
	}
	return toSerialize, nil
}

type NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown struct {
	value *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown
	isSet bool
}

func (v NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) Get() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown {
	return v.value
}

func (v *NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) Set(val *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown(val *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) *NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown {
	return &NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown{value: val, isSet: true}
}

func (v NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


