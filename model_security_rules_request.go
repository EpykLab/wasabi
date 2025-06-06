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

// checks if the SecurityRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityRulesRequest{}

// SecurityRulesRequest struct for SecurityRulesRequest
type SecurityRulesRequest struct {
	// Rule name
	Name string `json:"name"`
	// Rule body
	Rule map[string]interface{} `json:"rule"`
}

type _SecurityRulesRequest SecurityRulesRequest

// NewSecurityRulesRequest instantiates a new SecurityRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityRulesRequest(name string, rule map[string]interface{}) *SecurityRulesRequest {
	this := SecurityRulesRequest{}
	this.Name = name
	this.Rule = rule
	return &this
}

// NewSecurityRulesRequestWithDefaults instantiates a new SecurityRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityRulesRequestWithDefaults() *SecurityRulesRequest {
	this := SecurityRulesRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SecurityRulesRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityRulesRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityRulesRequest) SetName(v string) {
	o.Name = v
}

// GetRule returns the Rule field value
func (o *SecurityRulesRequest) GetRule() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *SecurityRulesRequest) GetRuleOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Rule, true
}

// SetRule sets field value
func (o *SecurityRulesRequest) SetRule(v map[string]interface{}) {
	o.Rule = v
}

func (o SecurityRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["rule"] = o.Rule
	return toSerialize, nil
}

func (o *SecurityRulesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"rule",
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

	varSecurityRulesRequest := _SecurityRulesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSecurityRulesRequest)

	if err != nil {
		return err
	}

	*o = SecurityRulesRequest(varSecurityRulesRequest)

	return err
}

type NullableSecurityRulesRequest struct {
	value *SecurityRulesRequest
	isSet bool
}

func (v NullableSecurityRulesRequest) Get() *SecurityRulesRequest {
	return v.value
}

func (v *NullableSecurityRulesRequest) Set(val *SecurityRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityRulesRequest(val *SecurityRulesRequest) *NullableSecurityRulesRequest {
	return &NullableSecurityRulesRequest{value: val, isSet: true}
}

func (v NullableSecurityRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


