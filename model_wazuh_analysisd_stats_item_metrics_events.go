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

// checks if the WazuhAnalysisdStatsItemMetricsEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhAnalysisdStatsItemMetricsEvents{}

// WazuhAnalysisdStatsItemMetricsEvents struct for WazuhAnalysisdStatsItemMetricsEvents
type WazuhAnalysisdStatsItemMetricsEvents struct {
	// Total processed events (analyzed by rules)
	Processed *int32 `json:"processed,omitempty"`
	// Total received events from agents and local modules
	Received *int32 `json:"received,omitempty"`
	ReceivedBreakdown *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown `json:"received_breakdown,omitempty"`
	WrittenBreakdown *WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown `json:"written_breakdown,omitempty"`
}

// NewWazuhAnalysisdStatsItemMetricsEvents instantiates a new WazuhAnalysisdStatsItemMetricsEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhAnalysisdStatsItemMetricsEvents() *WazuhAnalysisdStatsItemMetricsEvents {
	this := WazuhAnalysisdStatsItemMetricsEvents{}
	return &this
}

// NewWazuhAnalysisdStatsItemMetricsEventsWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhAnalysisdStatsItemMetricsEventsWithDefaults() *WazuhAnalysisdStatsItemMetricsEvents {
	this := WazuhAnalysisdStatsItemMetricsEvents{}
	return &this
}

// GetProcessed returns the Processed field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetProcessed() int32 {
	if o == nil || IsNil(o.Processed) {
		var ret int32
		return ret
	}
	return *o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetProcessedOk() (*int32, bool) {
	if o == nil || IsNil(o.Processed) {
		return nil, false
	}
	return o.Processed, true
}

// HasProcessed returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) HasProcessed() bool {
	if o != nil && !IsNil(o.Processed) {
		return true
	}

	return false
}

// SetProcessed gets a reference to the given int32 and assigns it to the Processed field.
func (o *WazuhAnalysisdStatsItemMetricsEvents) SetProcessed(v int32) {
	o.Processed = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceived() int32 {
	if o == nil || IsNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *WazuhAnalysisdStatsItemMetricsEvents) SetReceived(v int32) {
	o.Received = &v
}

// GetReceivedBreakdown returns the ReceivedBreakdown field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown {
	if o == nil || IsNil(o.ReceivedBreakdown) {
		var ret WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown
		return ret
	}
	return *o.ReceivedBreakdown
}

// GetReceivedBreakdownOk returns a tuple with the ReceivedBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown, bool) {
	if o == nil || IsNil(o.ReceivedBreakdown) {
		return nil, false
	}
	return o.ReceivedBreakdown, true
}

// HasReceivedBreakdown returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) HasReceivedBreakdown() bool {
	if o != nil && !IsNil(o.ReceivedBreakdown) {
		return true
	}

	return false
}

// SetReceivedBreakdown gets a reference to the given WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown and assigns it to the ReceivedBreakdown field.
func (o *WazuhAnalysisdStatsItemMetricsEvents) SetReceivedBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown) {
	o.ReceivedBreakdown = &v
}

// GetWrittenBreakdown returns the WrittenBreakdown field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetWrittenBreakdown() WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown {
	if o == nil || IsNil(o.WrittenBreakdown) {
		var ret WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown
		return ret
	}
	return *o.WrittenBreakdown
}

// GetWrittenBreakdownOk returns a tuple with the WrittenBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) GetWrittenBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown, bool) {
	if o == nil || IsNil(o.WrittenBreakdown) {
		return nil, false
	}
	return o.WrittenBreakdown, true
}

// HasWrittenBreakdown returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEvents) HasWrittenBreakdown() bool {
	if o != nil && !IsNil(o.WrittenBreakdown) {
		return true
	}

	return false
}

// SetWrittenBreakdown gets a reference to the given WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown and assigns it to the WrittenBreakdown field.
func (o *WazuhAnalysisdStatsItemMetricsEvents) SetWrittenBreakdown(v WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown) {
	o.WrittenBreakdown = &v
}

func (o WazuhAnalysisdStatsItemMetricsEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhAnalysisdStatsItemMetricsEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Processed) {
		toSerialize["processed"] = o.Processed
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.ReceivedBreakdown) {
		toSerialize["received_breakdown"] = o.ReceivedBreakdown
	}
	if !IsNil(o.WrittenBreakdown) {
		toSerialize["written_breakdown"] = o.WrittenBreakdown
	}
	return toSerialize, nil
}

type NullableWazuhAnalysisdStatsItemMetricsEvents struct {
	value *WazuhAnalysisdStatsItemMetricsEvents
	isSet bool
}

func (v NullableWazuhAnalysisdStatsItemMetricsEvents) Get() *WazuhAnalysisdStatsItemMetricsEvents {
	return v.value
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEvents) Set(val *WazuhAnalysisdStatsItemMetricsEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhAnalysisdStatsItemMetricsEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhAnalysisdStatsItemMetricsEvents(val *WazuhAnalysisdStatsItemMetricsEvents) *NullableWazuhAnalysisdStatsItemMetricsEvents {
	return &NullableWazuhAnalysisdStatsItemMetricsEvents{value: val, isSet: true}
}

func (v NullableWazuhAnalysisdStatsItemMetricsEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


