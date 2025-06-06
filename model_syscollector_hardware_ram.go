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

// checks if the SyscollectorHardwareRam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyscollectorHardwareRam{}

// SyscollectorHardwareRam struct for SyscollectorHardwareRam
type SyscollectorHardwareRam struct {
	// Current free RAM memory
	Free *int32 `json:"free,omitempty"`
	// Total RAM memory
	Total *int32 `json:"total,omitempty"`
	// RAM memory currently used
	Usage *int32 `json:"usage,omitempty"`
}

// NewSyscollectorHardwareRam instantiates a new SyscollectorHardwareRam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyscollectorHardwareRam() *SyscollectorHardwareRam {
	this := SyscollectorHardwareRam{}
	return &this
}

// NewSyscollectorHardwareRamWithDefaults instantiates a new SyscollectorHardwareRam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyscollectorHardwareRamWithDefaults() *SyscollectorHardwareRam {
	this := SyscollectorHardwareRam{}
	return &this
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *SyscollectorHardwareRam) GetFree() int32 {
	if o == nil || IsNil(o.Free) {
		var ret int32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyscollectorHardwareRam) GetFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.Free) {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *SyscollectorHardwareRam) HasFree() bool {
	if o != nil && !IsNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given int32 and assigns it to the Free field.
func (o *SyscollectorHardwareRam) SetFree(v int32) {
	o.Free = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SyscollectorHardwareRam) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyscollectorHardwareRam) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SyscollectorHardwareRam) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SyscollectorHardwareRam) SetTotal(v int32) {
	o.Total = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SyscollectorHardwareRam) GetUsage() int32 {
	if o == nil || IsNil(o.Usage) {
		var ret int32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyscollectorHardwareRam) GetUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SyscollectorHardwareRam) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int32 and assigns it to the Usage field.
func (o *SyscollectorHardwareRam) SetUsage(v int32) {
	o.Usage = &v
}

func (o SyscollectorHardwareRam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyscollectorHardwareRam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableSyscollectorHardwareRam struct {
	value *SyscollectorHardwareRam
	isSet bool
}

func (v NullableSyscollectorHardwareRam) Get() *SyscollectorHardwareRam {
	return v.value
}

func (v *NullableSyscollectorHardwareRam) Set(val *SyscollectorHardwareRam) {
	v.value = val
	v.isSet = true
}

func (v NullableSyscollectorHardwareRam) IsSet() bool {
	return v.isSet
}

func (v *NullableSyscollectorHardwareRam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyscollectorHardwareRam(val *SyscollectorHardwareRam) *NullableSyscollectorHardwareRam {
	return &NullableSyscollectorHardwareRam{value: val, isSet: true}
}

func (v NullableSyscollectorHardwareRam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyscollectorHardwareRam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


