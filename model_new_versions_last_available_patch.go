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

// checks if the NewVersionsLastAvailablePatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewVersionsLastAvailablePatch{}

// NewVersionsLastAvailablePatch Information about the most recent available patch update
type NewVersionsLastAvailablePatch struct {
	// Version in the format vX.Y.Z
	Tag *string `json:"tag,omitempty"`
	Description *string `json:"description,omitempty"`
	Title *string `json:"title,omitempty"`
	PublishedDate *string `json:"published_date,omitempty"`
	Semver *NewVersionsLastAvailableMajorSemver `json:"semver,omitempty"`
}

// NewNewVersionsLastAvailablePatch instantiates a new NewVersionsLastAvailablePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewVersionsLastAvailablePatch() *NewVersionsLastAvailablePatch {
	this := NewVersionsLastAvailablePatch{}
	return &this
}

// NewNewVersionsLastAvailablePatchWithDefaults instantiates a new NewVersionsLastAvailablePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewVersionsLastAvailablePatchWithDefaults() *NewVersionsLastAvailablePatch {
	this := NewVersionsLastAvailablePatch{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *NewVersionsLastAvailablePatch) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewVersionsLastAvailablePatch) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *NewVersionsLastAvailablePatch) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *NewVersionsLastAvailablePatch) SetTag(v string) {
	o.Tag = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewVersionsLastAvailablePatch) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewVersionsLastAvailablePatch) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewVersionsLastAvailablePatch) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewVersionsLastAvailablePatch) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NewVersionsLastAvailablePatch) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewVersionsLastAvailablePatch) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NewVersionsLastAvailablePatch) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NewVersionsLastAvailablePatch) SetTitle(v string) {
	o.Title = &v
}

// GetPublishedDate returns the PublishedDate field value if set, zero value otherwise.
func (o *NewVersionsLastAvailablePatch) GetPublishedDate() string {
	if o == nil || IsNil(o.PublishedDate) {
		var ret string
		return ret
	}
	return *o.PublishedDate
}

// GetPublishedDateOk returns a tuple with the PublishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewVersionsLastAvailablePatch) GetPublishedDateOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedDate) {
		return nil, false
	}
	return o.PublishedDate, true
}

// HasPublishedDate returns a boolean if a field has been set.
func (o *NewVersionsLastAvailablePatch) HasPublishedDate() bool {
	if o != nil && !IsNil(o.PublishedDate) {
		return true
	}

	return false
}

// SetPublishedDate gets a reference to the given string and assigns it to the PublishedDate field.
func (o *NewVersionsLastAvailablePatch) SetPublishedDate(v string) {
	o.PublishedDate = &v
}

// GetSemver returns the Semver field value if set, zero value otherwise.
func (o *NewVersionsLastAvailablePatch) GetSemver() NewVersionsLastAvailableMajorSemver {
	if o == nil || IsNil(o.Semver) {
		var ret NewVersionsLastAvailableMajorSemver
		return ret
	}
	return *o.Semver
}

// GetSemverOk returns a tuple with the Semver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewVersionsLastAvailablePatch) GetSemverOk() (*NewVersionsLastAvailableMajorSemver, bool) {
	if o == nil || IsNil(o.Semver) {
		return nil, false
	}
	return o.Semver, true
}

// HasSemver returns a boolean if a field has been set.
func (o *NewVersionsLastAvailablePatch) HasSemver() bool {
	if o != nil && !IsNil(o.Semver) {
		return true
	}

	return false
}

// SetSemver gets a reference to the given NewVersionsLastAvailableMajorSemver and assigns it to the Semver field.
func (o *NewVersionsLastAvailablePatch) SetSemver(v NewVersionsLastAvailableMajorSemver) {
	o.Semver = &v
}

func (o NewVersionsLastAvailablePatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewVersionsLastAvailablePatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.PublishedDate) {
		toSerialize["published_date"] = o.PublishedDate
	}
	if !IsNil(o.Semver) {
		toSerialize["semver"] = o.Semver
	}
	return toSerialize, nil
}

type NullableNewVersionsLastAvailablePatch struct {
	value *NewVersionsLastAvailablePatch
	isSet bool
}

func (v NullableNewVersionsLastAvailablePatch) Get() *NewVersionsLastAvailablePatch {
	return v.value
}

func (v *NullableNewVersionsLastAvailablePatch) Set(val *NewVersionsLastAvailablePatch) {
	v.value = val
	v.isSet = true
}

func (v NullableNewVersionsLastAvailablePatch) IsSet() bool {
	return v.isSet
}

func (v *NullableNewVersionsLastAvailablePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewVersionsLastAvailablePatch(val *NewVersionsLastAvailablePatch) *NullableNewVersionsLastAvailablePatch {
	return &NullableNewVersionsLastAvailablePatch{value: val, isSet: true}
}

func (v NullableNewVersionsLastAvailablePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewVersionsLastAvailablePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


