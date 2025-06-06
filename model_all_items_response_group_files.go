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

// checks if the AllItemsResponseGroupFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllItemsResponseGroupFiles{}

// AllItemsResponseGroupFiles struct for AllItemsResponseGroupFiles
type AllItemsResponseGroupFiles struct {
	// Number of items that have successfully applied the requested operation
	TotalAffectedItems int32 `json:"total_affected_items"`
	// List of items that have failed applying the requested operation
	FailedItems []SimpleApiError `json:"failed_items"`
	// Number of items that have failed applying the requested operation
	TotalFailedItems int32 `json:"total_failed_items"`
	// Items that successfully applied the API call action
	AffectedItems []GroupFiles `json:"affected_items"`
}

type _AllItemsResponseGroupFiles AllItemsResponseGroupFiles

// NewAllItemsResponseGroupFiles instantiates a new AllItemsResponseGroupFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllItemsResponseGroupFiles(totalAffectedItems int32, failedItems []SimpleApiError, totalFailedItems int32, affectedItems []GroupFiles) *AllItemsResponseGroupFiles {
	this := AllItemsResponseGroupFiles{}
	this.TotalAffectedItems = totalAffectedItems
	this.FailedItems = failedItems
	this.TotalFailedItems = totalFailedItems
	this.AffectedItems = affectedItems
	return &this
}

// NewAllItemsResponseGroupFilesWithDefaults instantiates a new AllItemsResponseGroupFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllItemsResponseGroupFilesWithDefaults() *AllItemsResponseGroupFiles {
	this := AllItemsResponseGroupFiles{}
	return &this
}

// GetTotalAffectedItems returns the TotalAffectedItems field value
func (o *AllItemsResponseGroupFiles) GetTotalAffectedItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAffectedItems
}

// GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field value
// and a boolean to check if the value has been set.
func (o *AllItemsResponseGroupFiles) GetTotalAffectedItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAffectedItems, true
}

// SetTotalAffectedItems sets field value
func (o *AllItemsResponseGroupFiles) SetTotalAffectedItems(v int32) {
	o.TotalAffectedItems = v
}

// GetFailedItems returns the FailedItems field value
func (o *AllItemsResponseGroupFiles) GetFailedItems() []SimpleApiError {
	if o == nil {
		var ret []SimpleApiError
		return ret
	}

	return o.FailedItems
}

// GetFailedItemsOk returns a tuple with the FailedItems field value
// and a boolean to check if the value has been set.
func (o *AllItemsResponseGroupFiles) GetFailedItemsOk() ([]SimpleApiError, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedItems, true
}

// SetFailedItems sets field value
func (o *AllItemsResponseGroupFiles) SetFailedItems(v []SimpleApiError) {
	o.FailedItems = v
}

// GetTotalFailedItems returns the TotalFailedItems field value
func (o *AllItemsResponseGroupFiles) GetTotalFailedItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalFailedItems
}

// GetTotalFailedItemsOk returns a tuple with the TotalFailedItems field value
// and a boolean to check if the value has been set.
func (o *AllItemsResponseGroupFiles) GetTotalFailedItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFailedItems, true
}

// SetTotalFailedItems sets field value
func (o *AllItemsResponseGroupFiles) SetTotalFailedItems(v int32) {
	o.TotalFailedItems = v
}

// GetAffectedItems returns the AffectedItems field value
func (o *AllItemsResponseGroupFiles) GetAffectedItems() []GroupFiles {
	if o == nil {
		var ret []GroupFiles
		return ret
	}

	return o.AffectedItems
}

// GetAffectedItemsOk returns a tuple with the AffectedItems field value
// and a boolean to check if the value has been set.
func (o *AllItemsResponseGroupFiles) GetAffectedItemsOk() ([]GroupFiles, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffectedItems, true
}

// SetAffectedItems sets field value
func (o *AllItemsResponseGroupFiles) SetAffectedItems(v []GroupFiles) {
	o.AffectedItems = v
}

func (o AllItemsResponseGroupFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllItemsResponseGroupFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_affected_items"] = o.TotalAffectedItems
	toSerialize["failed_items"] = o.FailedItems
	toSerialize["total_failed_items"] = o.TotalFailedItems
	toSerialize["affected_items"] = o.AffectedItems
	return toSerialize, nil
}

func (o *AllItemsResponseGroupFiles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_affected_items",
		"failed_items",
		"total_failed_items",
		"affected_items",
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

	varAllItemsResponseGroupFiles := _AllItemsResponseGroupFiles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAllItemsResponseGroupFiles)

	if err != nil {
		return err
	}

	*o = AllItemsResponseGroupFiles(varAllItemsResponseGroupFiles)

	return err
}

type NullableAllItemsResponseGroupFiles struct {
	value *AllItemsResponseGroupFiles
	isSet bool
}

func (v NullableAllItemsResponseGroupFiles) Get() *AllItemsResponseGroupFiles {
	return v.value
}

func (v *NullableAllItemsResponseGroupFiles) Set(val *AllItemsResponseGroupFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableAllItemsResponseGroupFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableAllItemsResponseGroupFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllItemsResponseGroupFiles(val *AllItemsResponseGroupFiles) *NullableAllItemsResponseGroupFiles {
	return &NullableAllItemsResponseGroupFiles{value: val, isSet: true}
}

func (v NullableAllItemsResponseGroupFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllItemsResponseGroupFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


