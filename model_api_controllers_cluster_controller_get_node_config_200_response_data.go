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
	"gopkg.in/validator.v2"
)

// ApiControllersClusterControllerGetNodeConfig200ResponseData - struct for ApiControllersClusterControllerGetNodeConfig200ResponseData
type ApiControllersClusterControllerGetNodeConfig200ResponseData struct {
	ArrayOfString *[]string
	MapmapOfStringAny *map[string]interface{}
}

// []stringAsApiControllersClusterControllerGetNodeConfig200ResponseData is a convenience function that returns []string wrapped in ApiControllersClusterControllerGetNodeConfig200ResponseData
func ArrayOfStringAsApiControllersClusterControllerGetNodeConfig200ResponseData(v *[]string) ApiControllersClusterControllerGetNodeConfig200ResponseData {
	return ApiControllersClusterControllerGetNodeConfig200ResponseData{
		ArrayOfString: v,
	}
}

// map[string]interface{}AsApiControllersClusterControllerGetNodeConfig200ResponseData is a convenience function that returns map[string]interface{} wrapped in ApiControllersClusterControllerGetNodeConfig200ResponseData
func MapmapOfStringAnyAsApiControllersClusterControllerGetNodeConfig200ResponseData(v *map[string]interface{}) ApiControllersClusterControllerGetNodeConfig200ResponseData {
	return ApiControllersClusterControllerGetNodeConfig200ResponseData{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ApiControllersClusterControllerGetNodeConfig200ResponseData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			if err = validator.Validate(dst.ArrayOfString); err != nil {
				dst.ArrayOfString = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ApiControllersClusterControllerGetNodeConfig200ResponseData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ApiControllersClusterControllerGetNodeConfig200ResponseData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ApiControllersClusterControllerGetNodeConfig200ResponseData) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ApiControllersClusterControllerGetNodeConfig200ResponseData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ApiControllersClusterControllerGetNodeConfig200ResponseData) GetActualInstanceValue() (interface{}) {
	if obj.ArrayOfString != nil {
		return *obj.ArrayOfString
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableApiControllersClusterControllerGetNodeConfig200ResponseData struct {
	value *ApiControllersClusterControllerGetNodeConfig200ResponseData
	isSet bool
}

func (v NullableApiControllersClusterControllerGetNodeConfig200ResponseData) Get() *ApiControllersClusterControllerGetNodeConfig200ResponseData {
	return v.value
}

func (v *NullableApiControllersClusterControllerGetNodeConfig200ResponseData) Set(val *ApiControllersClusterControllerGetNodeConfig200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiControllersClusterControllerGetNodeConfig200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiControllersClusterControllerGetNodeConfig200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiControllersClusterControllerGetNodeConfig200ResponseData(val *ApiControllersClusterControllerGetNodeConfig200ResponseData) *NullableApiControllersClusterControllerGetNodeConfig200ResponseData {
	return &NullableApiControllersClusterControllerGetNodeConfig200ResponseData{value: val, isSet: true}
}

func (v NullableApiControllersClusterControllerGetNodeConfig200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiControllersClusterControllerGetNodeConfig200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


