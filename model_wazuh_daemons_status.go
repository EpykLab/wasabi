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

// checks if the WazuhDaemonsStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WazuhDaemonsStatus{}

// WazuhDaemonsStatus struct for WazuhDaemonsStatus
type WazuhDaemonsStatus struct {
	WazuhAgentlessd *DaemonStatus `json:"wazuh-agentlessd,omitempty"`
	WazuhAnalysisd *DaemonStatus `json:"wazuh-analysisd,omitempty"`
	WazuhAuthd *DaemonStatus `json:"wazuh-authd,omitempty"`
	WazuhCsyslogd *DaemonStatus `json:"wazuh-csyslogd,omitempty"`
	WazuhDbd *DaemonStatus `json:"wazuh-dbd,omitempty"`
	WazuhExecd *DaemonStatus `json:"wazuh-execd,omitempty"`
	WazuhIntegratord *DaemonStatus `json:"wazuh-integratord,omitempty"`
	WazuhLogcollector *DaemonStatus `json:"wazuh-logcollector,omitempty"`
	WazuhMaild *DaemonStatus `json:"wazuh-maild,omitempty"`
	WazuhMonitord *DaemonStatus `json:"wazuh-monitord,omitempty"`
	WazuhRemoted *DaemonStatus `json:"wazuh-remoted,omitempty"`
	WazuhReportd *DaemonStatus `json:"wazuh-reportd,omitempty"`
	WazuhSyscheckd *DaemonStatus `json:"wazuh-syscheckd,omitempty"`
	WazuhApid *DaemonStatus `json:"wazuh-apid,omitempty"`
	WazuhClusterd *DaemonStatus `json:"wazuh-clusterd,omitempty"`
	WazuhDb *DaemonStatus `json:"wazuh-db,omitempty"`
	WazuhModulesd *DaemonStatus `json:"wazuh-modulesd,omitempty"`
}

// NewWazuhDaemonsStatus instantiates a new WazuhDaemonsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWazuhDaemonsStatus() *WazuhDaemonsStatus {
	this := WazuhDaemonsStatus{}
	return &this
}

// NewWazuhDaemonsStatusWithDefaults instantiates a new WazuhDaemonsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWazuhDaemonsStatusWithDefaults() *WazuhDaemonsStatus {
	this := WazuhDaemonsStatus{}
	return &this
}

// GetWazuhAgentlessd returns the WazuhAgentlessd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhAgentlessd() DaemonStatus {
	if o == nil || IsNil(o.WazuhAgentlessd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhAgentlessd
}

// GetWazuhAgentlessdOk returns a tuple with the WazuhAgentlessd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhAgentlessdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhAgentlessd) {
		return nil, false
	}
	return o.WazuhAgentlessd, true
}

// HasWazuhAgentlessd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhAgentlessd() bool {
	if o != nil && !IsNil(o.WazuhAgentlessd) {
		return true
	}

	return false
}

// SetWazuhAgentlessd gets a reference to the given DaemonStatus and assigns it to the WazuhAgentlessd field.
func (o *WazuhDaemonsStatus) SetWazuhAgentlessd(v DaemonStatus) {
	o.WazuhAgentlessd = &v
}

// GetWazuhAnalysisd returns the WazuhAnalysisd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhAnalysisd() DaemonStatus {
	if o == nil || IsNil(o.WazuhAnalysisd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhAnalysisd
}

// GetWazuhAnalysisdOk returns a tuple with the WazuhAnalysisd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhAnalysisdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhAnalysisd) {
		return nil, false
	}
	return o.WazuhAnalysisd, true
}

// HasWazuhAnalysisd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhAnalysisd() bool {
	if o != nil && !IsNil(o.WazuhAnalysisd) {
		return true
	}

	return false
}

// SetWazuhAnalysisd gets a reference to the given DaemonStatus and assigns it to the WazuhAnalysisd field.
func (o *WazuhDaemonsStatus) SetWazuhAnalysisd(v DaemonStatus) {
	o.WazuhAnalysisd = &v
}

// GetWazuhAuthd returns the WazuhAuthd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhAuthd() DaemonStatus {
	if o == nil || IsNil(o.WazuhAuthd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhAuthd
}

// GetWazuhAuthdOk returns a tuple with the WazuhAuthd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhAuthdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhAuthd) {
		return nil, false
	}
	return o.WazuhAuthd, true
}

// HasWazuhAuthd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhAuthd() bool {
	if o != nil && !IsNil(o.WazuhAuthd) {
		return true
	}

	return false
}

// SetWazuhAuthd gets a reference to the given DaemonStatus and assigns it to the WazuhAuthd field.
func (o *WazuhDaemonsStatus) SetWazuhAuthd(v DaemonStatus) {
	o.WazuhAuthd = &v
}

// GetWazuhCsyslogd returns the WazuhCsyslogd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhCsyslogd() DaemonStatus {
	if o == nil || IsNil(o.WazuhCsyslogd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhCsyslogd
}

// GetWazuhCsyslogdOk returns a tuple with the WazuhCsyslogd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhCsyslogdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhCsyslogd) {
		return nil, false
	}
	return o.WazuhCsyslogd, true
}

// HasWazuhCsyslogd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhCsyslogd() bool {
	if o != nil && !IsNil(o.WazuhCsyslogd) {
		return true
	}

	return false
}

// SetWazuhCsyslogd gets a reference to the given DaemonStatus and assigns it to the WazuhCsyslogd field.
func (o *WazuhDaemonsStatus) SetWazuhCsyslogd(v DaemonStatus) {
	o.WazuhCsyslogd = &v
}

// GetWazuhDbd returns the WazuhDbd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhDbd() DaemonStatus {
	if o == nil || IsNil(o.WazuhDbd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhDbd
}

// GetWazuhDbdOk returns a tuple with the WazuhDbd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhDbdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhDbd) {
		return nil, false
	}
	return o.WazuhDbd, true
}

// HasWazuhDbd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhDbd() bool {
	if o != nil && !IsNil(o.WazuhDbd) {
		return true
	}

	return false
}

// SetWazuhDbd gets a reference to the given DaemonStatus and assigns it to the WazuhDbd field.
func (o *WazuhDaemonsStatus) SetWazuhDbd(v DaemonStatus) {
	o.WazuhDbd = &v
}

// GetWazuhExecd returns the WazuhExecd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhExecd() DaemonStatus {
	if o == nil || IsNil(o.WazuhExecd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhExecd
}

// GetWazuhExecdOk returns a tuple with the WazuhExecd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhExecdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhExecd) {
		return nil, false
	}
	return o.WazuhExecd, true
}

// HasWazuhExecd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhExecd() bool {
	if o != nil && !IsNil(o.WazuhExecd) {
		return true
	}

	return false
}

// SetWazuhExecd gets a reference to the given DaemonStatus and assigns it to the WazuhExecd field.
func (o *WazuhDaemonsStatus) SetWazuhExecd(v DaemonStatus) {
	o.WazuhExecd = &v
}

// GetWazuhIntegratord returns the WazuhIntegratord field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhIntegratord() DaemonStatus {
	if o == nil || IsNil(o.WazuhIntegratord) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhIntegratord
}

// GetWazuhIntegratordOk returns a tuple with the WazuhIntegratord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhIntegratordOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhIntegratord) {
		return nil, false
	}
	return o.WazuhIntegratord, true
}

// HasWazuhIntegratord returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhIntegratord() bool {
	if o != nil && !IsNil(o.WazuhIntegratord) {
		return true
	}

	return false
}

// SetWazuhIntegratord gets a reference to the given DaemonStatus and assigns it to the WazuhIntegratord field.
func (o *WazuhDaemonsStatus) SetWazuhIntegratord(v DaemonStatus) {
	o.WazuhIntegratord = &v
}

// GetWazuhLogcollector returns the WazuhLogcollector field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhLogcollector() DaemonStatus {
	if o == nil || IsNil(o.WazuhLogcollector) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhLogcollector
}

// GetWazuhLogcollectorOk returns a tuple with the WazuhLogcollector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhLogcollectorOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhLogcollector) {
		return nil, false
	}
	return o.WazuhLogcollector, true
}

// HasWazuhLogcollector returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhLogcollector() bool {
	if o != nil && !IsNil(o.WazuhLogcollector) {
		return true
	}

	return false
}

// SetWazuhLogcollector gets a reference to the given DaemonStatus and assigns it to the WazuhLogcollector field.
func (o *WazuhDaemonsStatus) SetWazuhLogcollector(v DaemonStatus) {
	o.WazuhLogcollector = &v
}

// GetWazuhMaild returns the WazuhMaild field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhMaild() DaemonStatus {
	if o == nil || IsNil(o.WazuhMaild) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhMaild
}

// GetWazuhMaildOk returns a tuple with the WazuhMaild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhMaildOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhMaild) {
		return nil, false
	}
	return o.WazuhMaild, true
}

// HasWazuhMaild returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhMaild() bool {
	if o != nil && !IsNil(o.WazuhMaild) {
		return true
	}

	return false
}

// SetWazuhMaild gets a reference to the given DaemonStatus and assigns it to the WazuhMaild field.
func (o *WazuhDaemonsStatus) SetWazuhMaild(v DaemonStatus) {
	o.WazuhMaild = &v
}

// GetWazuhMonitord returns the WazuhMonitord field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhMonitord() DaemonStatus {
	if o == nil || IsNil(o.WazuhMonitord) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhMonitord
}

// GetWazuhMonitordOk returns a tuple with the WazuhMonitord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhMonitordOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhMonitord) {
		return nil, false
	}
	return o.WazuhMonitord, true
}

// HasWazuhMonitord returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhMonitord() bool {
	if o != nil && !IsNil(o.WazuhMonitord) {
		return true
	}

	return false
}

// SetWazuhMonitord gets a reference to the given DaemonStatus and assigns it to the WazuhMonitord field.
func (o *WazuhDaemonsStatus) SetWazuhMonitord(v DaemonStatus) {
	o.WazuhMonitord = &v
}

// GetWazuhRemoted returns the WazuhRemoted field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhRemoted() DaemonStatus {
	if o == nil || IsNil(o.WazuhRemoted) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhRemoted
}

// GetWazuhRemotedOk returns a tuple with the WazuhRemoted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhRemotedOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhRemoted) {
		return nil, false
	}
	return o.WazuhRemoted, true
}

// HasWazuhRemoted returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhRemoted() bool {
	if o != nil && !IsNil(o.WazuhRemoted) {
		return true
	}

	return false
}

// SetWazuhRemoted gets a reference to the given DaemonStatus and assigns it to the WazuhRemoted field.
func (o *WazuhDaemonsStatus) SetWazuhRemoted(v DaemonStatus) {
	o.WazuhRemoted = &v
}

// GetWazuhReportd returns the WazuhReportd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhReportd() DaemonStatus {
	if o == nil || IsNil(o.WazuhReportd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhReportd
}

// GetWazuhReportdOk returns a tuple with the WazuhReportd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhReportdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhReportd) {
		return nil, false
	}
	return o.WazuhReportd, true
}

// HasWazuhReportd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhReportd() bool {
	if o != nil && !IsNil(o.WazuhReportd) {
		return true
	}

	return false
}

// SetWazuhReportd gets a reference to the given DaemonStatus and assigns it to the WazuhReportd field.
func (o *WazuhDaemonsStatus) SetWazuhReportd(v DaemonStatus) {
	o.WazuhReportd = &v
}

// GetWazuhSyscheckd returns the WazuhSyscheckd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhSyscheckd() DaemonStatus {
	if o == nil || IsNil(o.WazuhSyscheckd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhSyscheckd
}

// GetWazuhSyscheckdOk returns a tuple with the WazuhSyscheckd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhSyscheckdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhSyscheckd) {
		return nil, false
	}
	return o.WazuhSyscheckd, true
}

// HasWazuhSyscheckd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhSyscheckd() bool {
	if o != nil && !IsNil(o.WazuhSyscheckd) {
		return true
	}

	return false
}

// SetWazuhSyscheckd gets a reference to the given DaemonStatus and assigns it to the WazuhSyscheckd field.
func (o *WazuhDaemonsStatus) SetWazuhSyscheckd(v DaemonStatus) {
	o.WazuhSyscheckd = &v
}

// GetWazuhApid returns the WazuhApid field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhApid() DaemonStatus {
	if o == nil || IsNil(o.WazuhApid) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhApid
}

// GetWazuhApidOk returns a tuple with the WazuhApid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhApidOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhApid) {
		return nil, false
	}
	return o.WazuhApid, true
}

// HasWazuhApid returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhApid() bool {
	if o != nil && !IsNil(o.WazuhApid) {
		return true
	}

	return false
}

// SetWazuhApid gets a reference to the given DaemonStatus and assigns it to the WazuhApid field.
func (o *WazuhDaemonsStatus) SetWazuhApid(v DaemonStatus) {
	o.WazuhApid = &v
}

// GetWazuhClusterd returns the WazuhClusterd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhClusterd() DaemonStatus {
	if o == nil || IsNil(o.WazuhClusterd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhClusterd
}

// GetWazuhClusterdOk returns a tuple with the WazuhClusterd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhClusterdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhClusterd) {
		return nil, false
	}
	return o.WazuhClusterd, true
}

// HasWazuhClusterd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhClusterd() bool {
	if o != nil && !IsNil(o.WazuhClusterd) {
		return true
	}

	return false
}

// SetWazuhClusterd gets a reference to the given DaemonStatus and assigns it to the WazuhClusterd field.
func (o *WazuhDaemonsStatus) SetWazuhClusterd(v DaemonStatus) {
	o.WazuhClusterd = &v
}

// GetWazuhDb returns the WazuhDb field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhDb() DaemonStatus {
	if o == nil || IsNil(o.WazuhDb) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhDb
}

// GetWazuhDbOk returns a tuple with the WazuhDb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhDbOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhDb) {
		return nil, false
	}
	return o.WazuhDb, true
}

// HasWazuhDb returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhDb() bool {
	if o != nil && !IsNil(o.WazuhDb) {
		return true
	}

	return false
}

// SetWazuhDb gets a reference to the given DaemonStatus and assigns it to the WazuhDb field.
func (o *WazuhDaemonsStatus) SetWazuhDb(v DaemonStatus) {
	o.WazuhDb = &v
}

// GetWazuhModulesd returns the WazuhModulesd field value if set, zero value otherwise.
func (o *WazuhDaemonsStatus) GetWazuhModulesd() DaemonStatus {
	if o == nil || IsNil(o.WazuhModulesd) {
		var ret DaemonStatus
		return ret
	}
	return *o.WazuhModulesd
}

// GetWazuhModulesdOk returns a tuple with the WazuhModulesd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WazuhDaemonsStatus) GetWazuhModulesdOk() (*DaemonStatus, bool) {
	if o == nil || IsNil(o.WazuhModulesd) {
		return nil, false
	}
	return o.WazuhModulesd, true
}

// HasWazuhModulesd returns a boolean if a field has been set.
func (o *WazuhDaemonsStatus) HasWazuhModulesd() bool {
	if o != nil && !IsNil(o.WazuhModulesd) {
		return true
	}

	return false
}

// SetWazuhModulesd gets a reference to the given DaemonStatus and assigns it to the WazuhModulesd field.
func (o *WazuhDaemonsStatus) SetWazuhModulesd(v DaemonStatus) {
	o.WazuhModulesd = &v
}

func (o WazuhDaemonsStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WazuhDaemonsStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WazuhAgentlessd) {
		toSerialize["wazuh-agentlessd"] = o.WazuhAgentlessd
	}
	if !IsNil(o.WazuhAnalysisd) {
		toSerialize["wazuh-analysisd"] = o.WazuhAnalysisd
	}
	if !IsNil(o.WazuhAuthd) {
		toSerialize["wazuh-authd"] = o.WazuhAuthd
	}
	if !IsNil(o.WazuhCsyslogd) {
		toSerialize["wazuh-csyslogd"] = o.WazuhCsyslogd
	}
	if !IsNil(o.WazuhDbd) {
		toSerialize["wazuh-dbd"] = o.WazuhDbd
	}
	if !IsNil(o.WazuhExecd) {
		toSerialize["wazuh-execd"] = o.WazuhExecd
	}
	if !IsNil(o.WazuhIntegratord) {
		toSerialize["wazuh-integratord"] = o.WazuhIntegratord
	}
	if !IsNil(o.WazuhLogcollector) {
		toSerialize["wazuh-logcollector"] = o.WazuhLogcollector
	}
	if !IsNil(o.WazuhMaild) {
		toSerialize["wazuh-maild"] = o.WazuhMaild
	}
	if !IsNil(o.WazuhMonitord) {
		toSerialize["wazuh-monitord"] = o.WazuhMonitord
	}
	if !IsNil(o.WazuhRemoted) {
		toSerialize["wazuh-remoted"] = o.WazuhRemoted
	}
	if !IsNil(o.WazuhReportd) {
		toSerialize["wazuh-reportd"] = o.WazuhReportd
	}
	if !IsNil(o.WazuhSyscheckd) {
		toSerialize["wazuh-syscheckd"] = o.WazuhSyscheckd
	}
	if !IsNil(o.WazuhApid) {
		toSerialize["wazuh-apid"] = o.WazuhApid
	}
	if !IsNil(o.WazuhClusterd) {
		toSerialize["wazuh-clusterd"] = o.WazuhClusterd
	}
	if !IsNil(o.WazuhDb) {
		toSerialize["wazuh-db"] = o.WazuhDb
	}
	if !IsNil(o.WazuhModulesd) {
		toSerialize["wazuh-modulesd"] = o.WazuhModulesd
	}
	return toSerialize, nil
}

type NullableWazuhDaemonsStatus struct {
	value *WazuhDaemonsStatus
	isSet bool
}

func (v NullableWazuhDaemonsStatus) Get() *WazuhDaemonsStatus {
	return v.value
}

func (v *NullableWazuhDaemonsStatus) Set(val *WazuhDaemonsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWazuhDaemonsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWazuhDaemonsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWazuhDaemonsStatus(val *WazuhDaemonsStatus) *NullableWazuhDaemonsStatus {
	return &NullableWazuhDaemonsStatus{value: val, isSet: true}
}

func (v NullableWazuhDaemonsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWazuhDaemonsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


