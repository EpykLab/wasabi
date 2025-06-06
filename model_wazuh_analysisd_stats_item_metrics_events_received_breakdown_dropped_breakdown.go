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

// checks if the WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown{}

// WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown struct for WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown
type WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown struct {
	// Events discarded from agentd because the queue was full
	Agent *int32 `json:"agent,omitempty"`
	// Events discarded from agentlessd because the queue was full
	Agentless *int32 `json:"agentless,omitempty"`
	// Synchronization events discarded because the queue was full
	Dbsync *int32 `json:"dbsync,omitempty"`
	IntegrationsBreakdown *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown `json:"integrations_breakdown,omitempty"`
	ModulesBreakdown *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown `json:"modules_breakdown,omitempty"`
	// Events discarded from monitord because the queue was full
	Monitor *int32 `json:"monitor,omitempty"`
	// Events discarded from remoted because the queue was full
	Remote *int32 `json:"remote,omitempty"`
	// Events discarded from syslog remoted because the queue was full
	Syslog *int32 `json:"syslog,omitempty"`
}

// NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown {
	this := WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown{}
	return &this
}

// NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown {
	this := WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown{}
	return &this
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgent() int32 {
	if o == nil || IsNil(o.Agent) {
		var ret int32
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentOk() (*int32, bool) {
	if o == nil || IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasAgent() bool {
	if o != nil && !IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given int32 and assigns it to the Agent field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetAgent(v int32) {
	o.Agent = &v
}

// GetAgentless returns the Agentless field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentless() int32 {
	if o == nil || IsNil(o.Agentless) {
		var ret int32
		return ret
	}
	return *o.Agentless
}

// GetAgentlessOk returns a tuple with the Agentless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentlessOk() (*int32, bool) {
	if o == nil || IsNil(o.Agentless) {
		return nil, false
	}
	return o.Agentless, true
}

// HasAgentless returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasAgentless() bool {
	if o != nil && !IsNil(o.Agentless) {
		return true
	}

	return false
}

// SetAgentless gets a reference to the given int32 and assigns it to the Agentless field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetAgentless(v int32) {
	o.Agentless = &v
}

// GetDbsync returns the Dbsync field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetDbsync() int32 {
	if o == nil || IsNil(o.Dbsync) {
		var ret int32
		return ret
	}
	return *o.Dbsync
}

// GetDbsyncOk returns a tuple with the Dbsync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetDbsyncOk() (*int32, bool) {
	if o == nil || IsNil(o.Dbsync) {
		return nil, false
	}
	return o.Dbsync, true
}

// HasDbsync returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasDbsync() bool {
	if o != nil && !IsNil(o.Dbsync) {
		return true
	}

	return false
}

// SetDbsync gets a reference to the given int32 and assigns it to the Dbsync field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetDbsync(v int32) {
	o.Dbsync = &v
}

// GetIntegrationsBreakdown returns the IntegrationsBreakdown field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetIntegrationsBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown {
	if o == nil || IsNil(o.IntegrationsBreakdown) {
		var ret WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown
		return ret
	}
	return *o.IntegrationsBreakdown
}

// GetIntegrationsBreakdownOk returns a tuple with the IntegrationsBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetIntegrationsBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown, bool) {
	if o == nil || IsNil(o.IntegrationsBreakdown) {
		return nil, false
	}
	return o.IntegrationsBreakdown, true
}

// HasIntegrationsBreakdown returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasIntegrationsBreakdown() bool {
	if o != nil && !IsNil(o.IntegrationsBreakdown) {
		return true
	}

	return false
}

// SetIntegrationsBreakdown gets a reference to the given WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown and assigns it to the IntegrationsBreakdown field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetIntegrationsBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown) {
	o.IntegrationsBreakdown = &v
}

// GetModulesBreakdown returns the ModulesBreakdown field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetModulesBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown {
	if o == nil || IsNil(o.ModulesBreakdown) {
		var ret WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown
		return ret
	}
	return *o.ModulesBreakdown
}

// GetModulesBreakdownOk returns a tuple with the ModulesBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetModulesBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown, bool) {
	if o == nil || IsNil(o.ModulesBreakdown) {
		return nil, false
	}
	return o.ModulesBreakdown, true
}

// HasModulesBreakdown returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasModulesBreakdown() bool {
	if o != nil && !IsNil(o.ModulesBreakdown) {
		return true
	}

	return false
}

// SetModulesBreakdown gets a reference to the given WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown and assigns it to the ModulesBreakdown field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetModulesBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) {
	o.ModulesBreakdown = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetMonitor() int32 {
	if o == nil || IsNil(o.Monitor) {
		var ret int32
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetMonitorOk() (*int32, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given int32 and assigns it to the Monitor field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetMonitor(v int32) {
	o.Monitor = &v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetRemote() int32 {
	if o == nil || IsNil(o.Remote) {
		var ret int32
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetRemoteOk() (*int32, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given int32 and assigns it to the Remote field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetRemote(v int32) {
	o.Remote = &v
}

// GetSyslog returns the Syslog field value if set, zero value otherwise.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetSyslog() int32 {
	if o == nil || IsNil(o.Syslog) {
		var ret int32
		return ret
	}
	return *o.Syslog
}

// GetSyslogOk returns a tuple with the Syslog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetSyslogOk() (*int32, bool) {
	if o == nil || IsNil(o.Syslog) {
		return nil, false
	}
	return o.Syslog, true
}

// HasSyslog returns a boolean if a field has been set.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasSyslog() bool {
	if o != nil && !IsNil(o.Syslog) {
		return true
	}

	return false
}

// SetSyslog gets a reference to the given int32 and assigns it to the Syslog field.
func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetSyslog(v int32) {
	o.Syslog = &v
}

func (o WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agent) {
		toSerialize["agent"] = o.Agent
	}
	if !IsNil(o.Agentless) {
		toSerialize["agentless"] = o.Agentless
	}
	if !IsNil(o.Dbsync) {
		toSerialize["dbsync"] = o.Dbsync
	}
	if !IsNil(o.IntegrationsBreakdown) {
		toSerialize["integrations_breakdown"] = o.IntegrationsBreakdown
	}
	if !IsNil(o.ModulesBreakdown) {
		toSerialize["modules_breakdown"] = o.ModulesBreakdown
	}
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if !IsNil(o.Remote) {
		toSerialize["remote"] = o.Remote
	}
	if !IsNil(o.Syslog) {
		toSerialize["syslog"] = o.Syslog
	}
	return toSerialize, nil
}

type NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown struct {
	value *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown
	isSet bool
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) Get() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown {
	return v.value
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) Set(val *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown(val *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown {
	return &NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown{value: val, isSet: true}
}

func (v NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


