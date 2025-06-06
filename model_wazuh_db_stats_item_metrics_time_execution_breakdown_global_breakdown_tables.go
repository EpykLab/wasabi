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

// checks if the WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables{}

// WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables Time taken by all queries per operation in tables (milliseconds)
type WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables struct {
	Agent *WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent `json:"agent,omitempty"`
	Belongs *WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs `json:"belongs,omitempty"`
	Group *WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup `json:"group,omitempty"`
	Labels *WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels `json:"labels,omitempty"`
}

// NewWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables instantiates a new WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables() *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables {
	this := WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables{}
	return &this
}

// NewWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTablesWithDefaults instantiates a new WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTablesWithDefaults() *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables {
	this := WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables{}
	return &this
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetAgent() WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent {
	if o == nil || IsNil(o.Agent) {
		var ret WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetAgentOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent, bool) {
	if o == nil || IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) HasAgent() bool {
	if o != nil && !IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent and assigns it to the Agent field.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) SetAgent(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesAgent) {
	o.Agent = &v
}

// GetBelongs returns the Belongs field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetBelongs() WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs {
	if o == nil || IsNil(o.Belongs) {
		var ret WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs
		return ret
	}
	return *o.Belongs
}

// GetBelongsOk returns a tuple with the Belongs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetBelongsOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs, bool) {
	if o == nil || IsNil(o.Belongs) {
		return nil, false
	}
	return o.Belongs, true
}

// HasBelongs returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) HasBelongs() bool {
	if o != nil && !IsNil(o.Belongs) {
		return true
	}

	return false
}

// SetBelongs gets a reference to the given WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs and assigns it to the Belongs field.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) SetBelongs(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesBelongs) {
	o.Belongs = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetGroup() WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup {
	if o == nil || IsNil(o.Group) {
		var ret WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetGroupOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup and assigns it to the Group field.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) SetGroup(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesGroup) {
	o.Group = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetLabels() WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels {
	if o == nil || IsNil(o.Labels) {
		var ret WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) GetLabelsOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels and assigns it to the Labels field.
func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) SetLabels(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdownTablesLabels) {
	o.Labels = &v
}

func (o WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agent) {
		toSerialize["agent"] = o.Agent
	}
	if !IsNil(o.Belongs) {
		toSerialize["belongs"] = o.Belongs
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables struct {
	value *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables
	isSet bool
}

func (v NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) Get() *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables {
	return v.value
}

func (v *NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) Set(val *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables(val *WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) *NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables {
	return &NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables{value: val, isSet: true}
}

func (v NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdownTables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


