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

// checks if the WazuhDBStatsItemMetricsQueries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhDBStatsItemMetricsQueries{}

// WazuhDBStatsItemMetricsQueries struct for WazuhDBStatsItemMetricsQueries
type WazuhDBStatsItemMetricsQueries struct {
	// Total number of queries through WazuhDB socket
	Received *int32 `json:"received,omitempty"`
	ReceivedBreakdown *WazuhDBStatsItemMetricsQueriesReceivedBreakdown `json:"received_breakdown,omitempty"`
}

// NewWazuhDBStatsItemMetricsQueries instantiates a new WazuhDBStatsItemMetricsQueries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhDBStatsItemMetricsQueries() *WazuhDBStatsItemMetricsQueries {
	this := WazuhDBStatsItemMetricsQueries{}
	return &this
}

// NewWazuhDBStatsItemMetricsQueriesWithDefaults instantiates a new WazuhDBStatsItemMetricsQueries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhDBStatsItemMetricsQueriesWithDefaults() *WazuhDBStatsItemMetricsQueries {
	this := WazuhDBStatsItemMetricsQueries{}
	return &this
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsQueries) GetReceived() int32 {
	if o == nil || IsNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsQueries) GetReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsQueries) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *WazuhDBStatsItemMetricsQueries) SetReceived(v int32) {
	o.Received = &v
}

// GetReceivedBreakdown returns the ReceivedBreakdown field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsQueries) GetReceivedBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdown {
	if o == nil || IsNil(o.ReceivedBreakdown) {
		var ret WazuhDBStatsItemMetricsQueriesReceivedBreakdown
		return ret
	}
	return *o.ReceivedBreakdown
}

// GetReceivedBreakdownOk returns a tuple with the ReceivedBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsQueries) GetReceivedBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdown, bool) {
	if o == nil || IsNil(o.ReceivedBreakdown) {
		return nil, false
	}
	return o.ReceivedBreakdown, true
}

// HasReceivedBreakdown returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsQueries) HasReceivedBreakdown() bool {
	if o != nil && !IsNil(o.ReceivedBreakdown) {
		return true
	}

	return false
}

// SetReceivedBreakdown gets a reference to the given WazuhDBStatsItemMetricsQueriesReceivedBreakdown and assigns it to the ReceivedBreakdown field.
func (o *WazuhDBStatsItemMetricsQueries) SetReceivedBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdown) {
	o.ReceivedBreakdown = &v
}

func (o WazuhDBStatsItemMetricsQueries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhDBStatsItemMetricsQueries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.ReceivedBreakdown) {
		toSerialize["received_breakdown"] = o.ReceivedBreakdown
	}
	return toSerialize, nil
}

type NullableWazuhDBStatsItemMetricsQueries struct {
	value *WazuhDBStatsItemMetricsQueries
	isSet bool
}

func (v NullableWazuhDBStatsItemMetricsQueries) Get() *WazuhDBStatsItemMetricsQueries {
	return v.value
}

func (v *NullableWazuhDBStatsItemMetricsQueries) Set(val *WazuhDBStatsItemMetricsQueries) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhDBStatsItemMetricsQueries) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhDBStatsItemMetricsQueries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhDBStatsItemMetricsQueries(val *WazuhDBStatsItemMetricsQueries) *NullableWazuhDBStatsItemMetricsQueries {
	return &NullableWazuhDBStatsItemMetricsQueries{value: val, isSet: true}
}

func (v NullableWazuhDBStatsItemMetricsQueries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhDBStatsItemMetricsQueries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


