/*
Wazuh API REST

The Wazuh API is an open-source RESTful API that allows for interaction with the Wazuh manager from a web browser, command line tools like cURL or any script or program that can make web requests. The Wazuh WUI relies on this heavily and Wazuh’s goal is to accommodate complete remote management of the Wazuh infrastructure via the Wazuh WUI. Use the Wazuh API to easily perform everyday actions like adding an agent, restarting the manager(s) or agent(s) or looking up syscheck details.  # Authentication  Wazuh API endpoints require authentication to be used. Therefore, all calls must include a JSON Web Token. JWT is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. Perform a call with `basicAuth` to `POST /security/user/authenticate` and obtain a JWT token to run any endpoint.  JWT tokens have a default duration of 900 seconds. To change this value, you must perform a call with a valid JWT token to `PUT /security/config`. After this change, you will need to get a new JWT token as all previously issued tokens are revoked when any change is performed on security configuration.  Login with USER and PASSWORD:  `curl -u <USER>:<PASSWORD> -k -X POST \"https://<HOST_IP>:55000/security/user/authenticate\"` ```json {     \"data\": {         \"token\": \"<YOUR_JWT_TOKEN>\"     },     \"error\": 0 } ``` Use the token from the previous response to perform any endpoint request:  `curl -k -X <METHOD> \"https://<HOST_IP>:55000/<ENDPOINT>\" -H  \"Authorization: Bearer <YOUR_JWT_TOKEN>\"`  Change the token base duration:  `curl -k -X PUT \"https://<HOST_IP>:55000/security/config\" -H \"Authorization: Bearer <YOUR_JWT_TOKEN>\" -d '{\"auth_token_exp_timeout\": <NEW_EXPIRE_TIME_IN_SECONDS>}'`  <SecurityDefinitions /> 

API version: 4.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DHCPStatus DHCP status
type DHCPStatus string

// List of DHCPStatus
const (
	ENABLED DHCPStatus = "enabled"
	DISABLED DHCPStatus = "disabled"
	UNKNOWN DHCPStatus = "unknown"
	BOOTP DHCPStatus = "BOOTP"
)

// All allowed values of DHCPStatus enum
var AllowedDHCPStatusEnumValues = []DHCPStatus{
	"enabled",
	"disabled",
	"unknown",
	"BOOTP",
}

func (v *DHCPStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DHCPStatus(value)
	for _, existing := range AllowedDHCPStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DHCPStatus", value)
}

// NewDHCPStatusFromValue returns a pointer to a valid DHCPStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDHCPStatusFromValue(v string) (*DHCPStatus, error) {
	ev := DHCPStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DHCPStatus: valid values are %v", v, AllowedDHCPStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DHCPStatus) IsValid() bool {
	for _, existing := range AllowedDHCPStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DHCPStatus value
func (v DHCPStatus) Ptr() *DHCPStatus {
	return &v
}

type NullableDHCPStatus struct {
	value *DHCPStatus
	isSet bool
}

func (v NullableDHCPStatus) Get() *DHCPStatus {
	return v.value
}

func (v *NullableDHCPStatus) Set(val *DHCPStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPStatus(val *DHCPStatus) *NullableDHCPStatus {
	return &NullableDHCPStatus{value: val, isSet: true}
}

func (v NullableDHCPStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

