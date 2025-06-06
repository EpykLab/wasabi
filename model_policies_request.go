/*
Wazuh API REST

The Wazuh API is an open-source RESTful API that allows for interaction with the Wazuh manager from a web browser, command line tools like cURL or any script or program that can make web requests. The Wazuh WUI relies on this heavily and Wazuh’s goal is to accommodate complete remote management of the Wazuh infrastructure via the Wazuh WUI. Use the Wazuh API to easily perform everyday actions like adding an agent, restarting the manager(s) or agent(s) or looking up syscheck details.  # Authentication  Wazuh API endpoints require authentication to be used. Therefore, all calls must include a JSON Web Token. JWT is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. Perform a call with `basicAuth` to `POST /security/user/authenticate` and obtain a JWT token to run any endpoint.  JWT tokens have a default duration of 900 seconds. To change this value, you must perform a call with a valid JWT token to `PUT /security/config`. After this change, you will need to get a new JWT token as all previously issued tokens are revoked when any change is performed on security configuration.  Login with USER and PASSWORD:  `curl -u <USER>:<PASSWORD> -k -X POST \"https://<HOST_IP>:55000/security/user/authenticate\"` ```json {     \"data\": {         \"token\": \"<YOUR_JWT_TOKEN>\"     },     \"error\": 0 } ``` Use the token from the previous response to perform any endpoint request:  `curl -k -X <METHOD> \"https://<HOST_IP>:55000/<ENDPOINT>\" -H  \"Authorization: Bearer <YOUR_JWT_TOKEN>\"`  Change the token base duration:  `curl -k -X PUT \"https://<HOST_IP>:55000/security/config\" -H \"Authorization: Bearer <YOUR_JWT_TOKEN>\" -d '{\"auth_token_exp_timeout\": <NEW_EXPIRE_TIME_IN_SECONDS>}'`  <SecurityDefinitions /> 

API version: 4.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PoliciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoliciesRequest{}

// PoliciesRequest struct for PoliciesRequest
type PoliciesRequest struct {
	// Policy name
	Name string `json:"name"`
	Policy PoliciesRequestPolicy `json:"policy"`
}

type _PoliciesRequest PoliciesRequest

// NewPoliciesRequest instantiates a new PoliciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoliciesRequest(name string, policy PoliciesRequestPolicy) *PoliciesRequest {
	this := PoliciesRequest{}
	this.Name = name
	this.Policy = policy
	return &this
}

// NewPoliciesRequestWithDefaults instantiates a new PoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoliciesRequestWithDefaults() *PoliciesRequest {
	this := PoliciesRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PoliciesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PoliciesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PoliciesRequest) SetName(v string) {
	o.Name = v
}

// GetPolicy returns the Policy field value
func (o *PoliciesRequest) GetPolicy() PoliciesRequestPolicy {
	if o == nil {
		var ret PoliciesRequestPolicy
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *PoliciesRequest) GetPolicyOk() (*PoliciesRequestPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *PoliciesRequest) SetPolicy(v PoliciesRequestPolicy) {
	o.Policy = v
}

func (o PoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoliciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

func (o *PoliciesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"policy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPoliciesRequest := _PoliciesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoliciesRequest)

	if err != nil {
		return err
	}

	*o = PoliciesRequest(varPoliciesRequest)

	return err
}

type NullablePoliciesRequest struct {
	value *PoliciesRequest
	isSet bool
}

func (v NullablePoliciesRequest) Get() *PoliciesRequest {
	return v.value
}

func (v *NullablePoliciesRequest) Set(val *PoliciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePoliciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePoliciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoliciesRequest(val *PoliciesRequest) *NullablePoliciesRequest {
	return &NullablePoliciesRequest{value: val, isSet: true}
}

func (v NullablePoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoliciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


