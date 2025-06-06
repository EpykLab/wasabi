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

// checks if the WazuhRemotedStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhRemotedStats{}

// WazuhRemotedStats struct for WazuhRemotedStats
type WazuhRemotedStats struct {
	// Number of control messages received from all agents during the last five seconds
	CtrlMsgCount *float32 `json:"ctrl_msg_count,omitempty"`
	// Number of discarded events received from agents during the last five seconds
	DiscardedCount *float32 `json:"discarded_count,omitempty"`
	// Number of events sent to analysisd during the last five seconds
	EvtCount *float32 `json:"evt_count,omitempty"`
	// Number of sent bytes to the agents during the last five seconds
	SentBytes *float32 `json:"sent_bytes,omitempty"`
	// Usage of the queue to storage events from agents
	QueueSize *float32 `json:"queue_size,omitempty"`
	// Number of received bytes from all agents during the last five seconds
	RecvBytes *float32 `json:"recv_bytes,omitempty"`
	// Number of TCP active sessions during the last five seconds
	TcpSessions *float32 `json:"tcp_sessions,omitempty"`
	// Total queue size to store events from agents
	TotalQueueSize *float32 `json:"total_queue_size,omitempty"`
}

// NewWazuhRemotedStats instantiates a new WazuhRemotedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhRemotedStats() *WazuhRemotedStats {
	this := WazuhRemotedStats{}
	return &this
}

// NewWazuhRemotedStatsWithDefaults instantiates a new WazuhRemotedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhRemotedStatsWithDefaults() *WazuhRemotedStats {
	this := WazuhRemotedStats{}
	return &this
}

// GetCtrlMsgCount returns the CtrlMsgCount field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetCtrlMsgCount() float32 {
	if o == nil || IsNil(o.CtrlMsgCount) {
		var ret float32
		return ret
	}
	return *o.CtrlMsgCount
}

// GetCtrlMsgCountOk returns a tuple with the CtrlMsgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetCtrlMsgCountOk() (*float32, bool) {
	if o == nil || IsNil(o.CtrlMsgCount) {
		return nil, false
	}
	return o.CtrlMsgCount, true
}

// HasCtrlMsgCount returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasCtrlMsgCount() bool {
	if o != nil && !IsNil(o.CtrlMsgCount) {
		return true
	}

	return false
}

// SetCtrlMsgCount gets a reference to the given float32 and assigns it to the CtrlMsgCount field.
func (o *WazuhRemotedStats) SetCtrlMsgCount(v float32) {
	o.CtrlMsgCount = &v
}

// GetDiscardedCount returns the DiscardedCount field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetDiscardedCount() float32 {
	if o == nil || IsNil(o.DiscardedCount) {
		var ret float32
		return ret
	}
	return *o.DiscardedCount
}

// GetDiscardedCountOk returns a tuple with the DiscardedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetDiscardedCountOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscardedCount) {
		return nil, false
	}
	return o.DiscardedCount, true
}

// HasDiscardedCount returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasDiscardedCount() bool {
	if o != nil && !IsNil(o.DiscardedCount) {
		return true
	}

	return false
}

// SetDiscardedCount gets a reference to the given float32 and assigns it to the DiscardedCount field.
func (o *WazuhRemotedStats) SetDiscardedCount(v float32) {
	o.DiscardedCount = &v
}

// GetEvtCount returns the EvtCount field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetEvtCount() float32 {
	if o == nil || IsNil(o.EvtCount) {
		var ret float32
		return ret
	}
	return *o.EvtCount
}

// GetEvtCountOk returns a tuple with the EvtCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetEvtCountOk() (*float32, bool) {
	if o == nil || IsNil(o.EvtCount) {
		return nil, false
	}
	return o.EvtCount, true
}

// HasEvtCount returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasEvtCount() bool {
	if o != nil && !IsNil(o.EvtCount) {
		return true
	}

	return false
}

// SetEvtCount gets a reference to the given float32 and assigns it to the EvtCount field.
func (o *WazuhRemotedStats) SetEvtCount(v float32) {
	o.EvtCount = &v
}

// GetSentBytes returns the SentBytes field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetSentBytes() float32 {
	if o == nil || IsNil(o.SentBytes) {
		var ret float32
		return ret
	}
	return *o.SentBytes
}

// GetSentBytesOk returns a tuple with the SentBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetSentBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.SentBytes) {
		return nil, false
	}
	return o.SentBytes, true
}

// HasSentBytes returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasSentBytes() bool {
	if o != nil && !IsNil(o.SentBytes) {
		return true
	}

	return false
}

// SetSentBytes gets a reference to the given float32 and assigns it to the SentBytes field.
func (o *WazuhRemotedStats) SetSentBytes(v float32) {
	o.SentBytes = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetQueueSize() float32 {
	if o == nil || IsNil(o.QueueSize) {
		var ret float32
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetQueueSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given float32 and assigns it to the QueueSize field.
func (o *WazuhRemotedStats) SetQueueSize(v float32) {
	o.QueueSize = &v
}

// GetRecvBytes returns the RecvBytes field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetRecvBytes() float32 {
	if o == nil || IsNil(o.RecvBytes) {
		var ret float32
		return ret
	}
	return *o.RecvBytes
}

// GetRecvBytesOk returns a tuple with the RecvBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetRecvBytesOk() (*float32, bool) {
	if o == nil || IsNil(o.RecvBytes) {
		return nil, false
	}
	return o.RecvBytes, true
}

// HasRecvBytes returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasRecvBytes() bool {
	if o != nil && !IsNil(o.RecvBytes) {
		return true
	}

	return false
}

// SetRecvBytes gets a reference to the given float32 and assigns it to the RecvBytes field.
func (o *WazuhRemotedStats) SetRecvBytes(v float32) {
	o.RecvBytes = &v
}

// GetTcpSessions returns the TcpSessions field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetTcpSessions() float32 {
	if o == nil || IsNil(o.TcpSessions) {
		var ret float32
		return ret
	}
	return *o.TcpSessions
}

// GetTcpSessionsOk returns a tuple with the TcpSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetTcpSessionsOk() (*float32, bool) {
	if o == nil || IsNil(o.TcpSessions) {
		return nil, false
	}
	return o.TcpSessions, true
}

// HasTcpSessions returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasTcpSessions() bool {
	if o != nil && !IsNil(o.TcpSessions) {
		return true
	}

	return false
}

// SetTcpSessions gets a reference to the given float32 and assigns it to the TcpSessions field.
func (o *WazuhRemotedStats) SetTcpSessions(v float32) {
	o.TcpSessions = &v
}

// GetTotalQueueSize returns the TotalQueueSize field value if set, zero value otherwise.
func (o *WazuhRemotedStats) GetTotalQueueSize() float32 {
	if o == nil || IsNil(o.TotalQueueSize) {
		var ret float32
		return ret
	}
	return *o.TotalQueueSize
}

// GetTotalQueueSizeOk returns a tuple with the TotalQueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhRemotedStats) GetTotalQueueSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalQueueSize) {
		return nil, false
	}
	return o.TotalQueueSize, true
}

// HasTotalQueueSize returns a boolean if a field has been set.
func (o *WazuhRemotedStats) HasTotalQueueSize() bool {
	if o != nil && !IsNil(o.TotalQueueSize) {
		return true
	}

	return false
}

// SetTotalQueueSize gets a reference to the given float32 and assigns it to the TotalQueueSize field.
func (o *WazuhRemotedStats) SetTotalQueueSize(v float32) {
	o.TotalQueueSize = &v
}

func (o WazuhRemotedStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhRemotedStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CtrlMsgCount) {
		toSerialize["ctrl_msg_count"] = o.CtrlMsgCount
	}
	if !IsNil(o.DiscardedCount) {
		toSerialize["discarded_count"] = o.DiscardedCount
	}
	if !IsNil(o.EvtCount) {
		toSerialize["evt_count"] = o.EvtCount
	}
	if !IsNil(o.SentBytes) {
		toSerialize["sent_bytes"] = o.SentBytes
	}
	if !IsNil(o.QueueSize) {
		toSerialize["queue_size"] = o.QueueSize
	}
	if !IsNil(o.RecvBytes) {
		toSerialize["recv_bytes"] = o.RecvBytes
	}
	if !IsNil(o.TcpSessions) {
		toSerialize["tcp_sessions"] = o.TcpSessions
	}
	if !IsNil(o.TotalQueueSize) {
		toSerialize["total_queue_size"] = o.TotalQueueSize
	}
	return toSerialize, nil
}

type NullableWazuhRemotedStats struct {
	value *WazuhRemotedStats
	isSet bool
}

func (v NullableWazuhRemotedStats) Get() *WazuhRemotedStats {
	return v.value
}

func (v *NullableWazuhRemotedStats) Set(val *WazuhRemotedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhRemotedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhRemotedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhRemotedStats(val *WazuhRemotedStats) *NullableWazuhRemotedStats {
	return &NullableWazuhRemotedStats{value: val, isSet: true}
}

func (v NullableWazuhRemotedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhRemotedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


