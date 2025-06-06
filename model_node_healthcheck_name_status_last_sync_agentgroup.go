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

// checks if the NodeHealthcheckNameStatusLastSyncAgentgroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeHealthcheckNameStatusLastSyncAgentgroup{}

// NodeHealthcheckNameStatusLastSyncAgentgroup struct for NodeHealthcheckNameStatusLastSyncAgentgroup
type NodeHealthcheckNameStatusLastSyncAgentgroup struct {
	DateStart *string `json:"date_start,omitempty"`
	DateEnd *string `json:"date_end,omitempty"`
	NSyncedChunks *int32 `json:"n_synced_chunks,omitempty"`
}

// NewNodeHealthcheckNameStatusLastSyncAgentgroup instantiates a new NodeHealthcheckNameStatusLastSyncAgentgroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeHealthcheckNameStatusLastSyncAgentgroup() *NodeHealthcheckNameStatusLastSyncAgentgroup {
	this := NodeHealthcheckNameStatusLastSyncAgentgroup{}
	return &this
}

// NewNodeHealthcheckNameStatusLastSyncAgentgroupWithDefaults instantiates a new NodeHealthcheckNameStatusLastSyncAgentgroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeHealthcheckNameStatusLastSyncAgentgroupWithDefaults() *NodeHealthcheckNameStatusLastSyncAgentgroup {
	this := NodeHealthcheckNameStatusLastSyncAgentgroup{}
	return &this
}

// GetDateStart returns the DateStart field value if set, zero value otherwise.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetDateStart() string {
	if o == nil || IsNil(o.DateStart) {
		var ret string
		return ret
	}
	return *o.DateStart
}

// GetDateStartOk returns a tuple with the DateStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetDateStartOk() (*string, bool) {
	if o == nil || IsNil(o.DateStart) {
		return nil, false
	}
	return o.DateStart, true
}

// HasDateStart returns a boolean if a field has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) HasDateStart() bool {
	if o != nil && !IsNil(o.DateStart) {
		return true
	}

	return false
}

// SetDateStart gets a reference to the given string and assigns it to the DateStart field.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) SetDateStart(v string) {
	o.DateStart = &v
}

// GetDateEnd returns the DateEnd field value if set, zero value otherwise.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetDateEnd() string {
	if o == nil || IsNil(o.DateEnd) {
		var ret string
		return ret
	}
	return *o.DateEnd
}

// GetDateEndOk returns a tuple with the DateEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetDateEndOk() (*string, bool) {
	if o == nil || IsNil(o.DateEnd) {
		return nil, false
	}
	return o.DateEnd, true
}

// HasDateEnd returns a boolean if a field has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) HasDateEnd() bool {
	if o != nil && !IsNil(o.DateEnd) {
		return true
	}

	return false
}

// SetDateEnd gets a reference to the given string and assigns it to the DateEnd field.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) SetDateEnd(v string) {
	o.DateEnd = &v
}

// GetNSyncedChunks returns the NSyncedChunks field value if set, zero value otherwise.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetNSyncedChunks() int32 {
	if o == nil || IsNil(o.NSyncedChunks) {
		var ret int32
		return ret
	}
	return *o.NSyncedChunks
}

// GetNSyncedChunksOk returns a tuple with the NSyncedChunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) GetNSyncedChunksOk() (*int32, bool) {
	if o == nil || IsNil(o.NSyncedChunks) {
		return nil, false
	}
	return o.NSyncedChunks, true
}

// HasNSyncedChunks returns a boolean if a field has been set.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) HasNSyncedChunks() bool {
	if o != nil && !IsNil(o.NSyncedChunks) {
		return true
	}

	return false
}

// SetNSyncedChunks gets a reference to the given int32 and assigns it to the NSyncedChunks field.
func (o *NodeHealthcheckNameStatusLastSyncAgentgroup) SetNSyncedChunks(v int32) {
	o.NSyncedChunks = &v
}

func (o NodeHealthcheckNameStatusLastSyncAgentgroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeHealthcheckNameStatusLastSyncAgentgroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateStart) {
		toSerialize["date_start"] = o.DateStart
	}
	if !IsNil(o.DateEnd) {
		toSerialize["date_end"] = o.DateEnd
	}
	if !IsNil(o.NSyncedChunks) {
		toSerialize["n_synced_chunks"] = o.NSyncedChunks
	}
	return toSerialize, nil
}

type NullableNodeHealthcheckNameStatusLastSyncAgentgroup struct {
	value *NodeHealthcheckNameStatusLastSyncAgentgroup
	isSet bool
}

func (v NullableNodeHealthcheckNameStatusLastSyncAgentgroup) Get() *NodeHealthcheckNameStatusLastSyncAgentgroup {
	return v.value
}

func (v *NullableNodeHealthcheckNameStatusLastSyncAgentgroup) Set(val *NodeHealthcheckNameStatusLastSyncAgentgroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeHealthcheckNameStatusLastSyncAgentgroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeHealthcheckNameStatusLastSyncAgentgroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeHealthcheckNameStatusLastSyncAgentgroup(val *NodeHealthcheckNameStatusLastSyncAgentgroup) *NullableNodeHealthcheckNameStatusLastSyncAgentgroup {
	return &NullableNodeHealthcheckNameStatusLastSyncAgentgroup{value: val, isSet: true}
}

func (v NullableNodeHealthcheckNameStatusLastSyncAgentgroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeHealthcheckNameStatusLastSyncAgentgroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


