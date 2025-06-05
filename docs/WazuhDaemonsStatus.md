# WazuhDaemonsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WazuhAgentlessd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhAnalysisd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhAuthd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhCsyslogd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhDbd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhExecd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhIntegratord** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhLogcollector** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhMaild** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhMonitord** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhRemoted** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhReportd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhSyscheckd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhApid** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhClusterd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhDb** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 
**WazuhModulesd** | Pointer to [**DaemonStatus**](DaemonStatus.md) |  | [optional] 

## Methods

### NewWazuhDaemonsStatus

`func NewWazuhDaemonsStatus() *WazuhDaemonsStatus`

NewWazuhDaemonsStatus instantiates a new WazuhDaemonsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhDaemonsStatusWithDefaults

`func NewWazuhDaemonsStatusWithDefaults() *WazuhDaemonsStatus`

NewWazuhDaemonsStatusWithDefaults instantiates a new WazuhDaemonsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWazuhAgentlessd

`func (o *WazuhDaemonsStatus) GetWazuhAgentlessd() DaemonStatus`

GetWazuhAgentlessd returns the WazuhAgentlessd field if non-nil, zero value otherwise.

### GetWazuhAgentlessdOk

`func (o *WazuhDaemonsStatus) GetWazuhAgentlessdOk() (*DaemonStatus, bool)`

GetWazuhAgentlessdOk returns a tuple with the WazuhAgentlessd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAgentlessd

`func (o *WazuhDaemonsStatus) SetWazuhAgentlessd(v DaemonStatus)`

SetWazuhAgentlessd sets WazuhAgentlessd field to given value.

### HasWazuhAgentlessd

`func (o *WazuhDaemonsStatus) HasWazuhAgentlessd() bool`

HasWazuhAgentlessd returns a boolean if a field has been set.

### GetWazuhAnalysisd

`func (o *WazuhDaemonsStatus) GetWazuhAnalysisd() DaemonStatus`

GetWazuhAnalysisd returns the WazuhAnalysisd field if non-nil, zero value otherwise.

### GetWazuhAnalysisdOk

`func (o *WazuhDaemonsStatus) GetWazuhAnalysisdOk() (*DaemonStatus, bool)`

GetWazuhAnalysisdOk returns a tuple with the WazuhAnalysisd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAnalysisd

`func (o *WazuhDaemonsStatus) SetWazuhAnalysisd(v DaemonStatus)`

SetWazuhAnalysisd sets WazuhAnalysisd field to given value.

### HasWazuhAnalysisd

`func (o *WazuhDaemonsStatus) HasWazuhAnalysisd() bool`

HasWazuhAnalysisd returns a boolean if a field has been set.

### GetWazuhAuthd

`func (o *WazuhDaemonsStatus) GetWazuhAuthd() DaemonStatus`

GetWazuhAuthd returns the WazuhAuthd field if non-nil, zero value otherwise.

### GetWazuhAuthdOk

`func (o *WazuhDaemonsStatus) GetWazuhAuthdOk() (*DaemonStatus, bool)`

GetWazuhAuthdOk returns a tuple with the WazuhAuthd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAuthd

`func (o *WazuhDaemonsStatus) SetWazuhAuthd(v DaemonStatus)`

SetWazuhAuthd sets WazuhAuthd field to given value.

### HasWazuhAuthd

`func (o *WazuhDaemonsStatus) HasWazuhAuthd() bool`

HasWazuhAuthd returns a boolean if a field has been set.

### GetWazuhCsyslogd

`func (o *WazuhDaemonsStatus) GetWazuhCsyslogd() DaemonStatus`

GetWazuhCsyslogd returns the WazuhCsyslogd field if non-nil, zero value otherwise.

### GetWazuhCsyslogdOk

`func (o *WazuhDaemonsStatus) GetWazuhCsyslogdOk() (*DaemonStatus, bool)`

GetWazuhCsyslogdOk returns a tuple with the WazuhCsyslogd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhCsyslogd

`func (o *WazuhDaemonsStatus) SetWazuhCsyslogd(v DaemonStatus)`

SetWazuhCsyslogd sets WazuhCsyslogd field to given value.

### HasWazuhCsyslogd

`func (o *WazuhDaemonsStatus) HasWazuhCsyslogd() bool`

HasWazuhCsyslogd returns a boolean if a field has been set.

### GetWazuhDbd

`func (o *WazuhDaemonsStatus) GetWazuhDbd() DaemonStatus`

GetWazuhDbd returns the WazuhDbd field if non-nil, zero value otherwise.

### GetWazuhDbdOk

`func (o *WazuhDaemonsStatus) GetWazuhDbdOk() (*DaemonStatus, bool)`

GetWazuhDbdOk returns a tuple with the WazuhDbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhDbd

`func (o *WazuhDaemonsStatus) SetWazuhDbd(v DaemonStatus)`

SetWazuhDbd sets WazuhDbd field to given value.

### HasWazuhDbd

`func (o *WazuhDaemonsStatus) HasWazuhDbd() bool`

HasWazuhDbd returns a boolean if a field has been set.

### GetWazuhExecd

`func (o *WazuhDaemonsStatus) GetWazuhExecd() DaemonStatus`

GetWazuhExecd returns the WazuhExecd field if non-nil, zero value otherwise.

### GetWazuhExecdOk

`func (o *WazuhDaemonsStatus) GetWazuhExecdOk() (*DaemonStatus, bool)`

GetWazuhExecdOk returns a tuple with the WazuhExecd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhExecd

`func (o *WazuhDaemonsStatus) SetWazuhExecd(v DaemonStatus)`

SetWazuhExecd sets WazuhExecd field to given value.

### HasWazuhExecd

`func (o *WazuhDaemonsStatus) HasWazuhExecd() bool`

HasWazuhExecd returns a boolean if a field has been set.

### GetWazuhIntegratord

`func (o *WazuhDaemonsStatus) GetWazuhIntegratord() DaemonStatus`

GetWazuhIntegratord returns the WazuhIntegratord field if non-nil, zero value otherwise.

### GetWazuhIntegratordOk

`func (o *WazuhDaemonsStatus) GetWazuhIntegratordOk() (*DaemonStatus, bool)`

GetWazuhIntegratordOk returns a tuple with the WazuhIntegratord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhIntegratord

`func (o *WazuhDaemonsStatus) SetWazuhIntegratord(v DaemonStatus)`

SetWazuhIntegratord sets WazuhIntegratord field to given value.

### HasWazuhIntegratord

`func (o *WazuhDaemonsStatus) HasWazuhIntegratord() bool`

HasWazuhIntegratord returns a boolean if a field has been set.

### GetWazuhLogcollector

`func (o *WazuhDaemonsStatus) GetWazuhLogcollector() DaemonStatus`

GetWazuhLogcollector returns the WazuhLogcollector field if non-nil, zero value otherwise.

### GetWazuhLogcollectorOk

`func (o *WazuhDaemonsStatus) GetWazuhLogcollectorOk() (*DaemonStatus, bool)`

GetWazuhLogcollectorOk returns a tuple with the WazuhLogcollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhLogcollector

`func (o *WazuhDaemonsStatus) SetWazuhLogcollector(v DaemonStatus)`

SetWazuhLogcollector sets WazuhLogcollector field to given value.

### HasWazuhLogcollector

`func (o *WazuhDaemonsStatus) HasWazuhLogcollector() bool`

HasWazuhLogcollector returns a boolean if a field has been set.

### GetWazuhMaild

`func (o *WazuhDaemonsStatus) GetWazuhMaild() DaemonStatus`

GetWazuhMaild returns the WazuhMaild field if non-nil, zero value otherwise.

### GetWazuhMaildOk

`func (o *WazuhDaemonsStatus) GetWazuhMaildOk() (*DaemonStatus, bool)`

GetWazuhMaildOk returns a tuple with the WazuhMaild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhMaild

`func (o *WazuhDaemonsStatus) SetWazuhMaild(v DaemonStatus)`

SetWazuhMaild sets WazuhMaild field to given value.

### HasWazuhMaild

`func (o *WazuhDaemonsStatus) HasWazuhMaild() bool`

HasWazuhMaild returns a boolean if a field has been set.

### GetWazuhMonitord

`func (o *WazuhDaemonsStatus) GetWazuhMonitord() DaemonStatus`

GetWazuhMonitord returns the WazuhMonitord field if non-nil, zero value otherwise.

### GetWazuhMonitordOk

`func (o *WazuhDaemonsStatus) GetWazuhMonitordOk() (*DaemonStatus, bool)`

GetWazuhMonitordOk returns a tuple with the WazuhMonitord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhMonitord

`func (o *WazuhDaemonsStatus) SetWazuhMonitord(v DaemonStatus)`

SetWazuhMonitord sets WazuhMonitord field to given value.

### HasWazuhMonitord

`func (o *WazuhDaemonsStatus) HasWazuhMonitord() bool`

HasWazuhMonitord returns a boolean if a field has been set.

### GetWazuhRemoted

`func (o *WazuhDaemonsStatus) GetWazuhRemoted() DaemonStatus`

GetWazuhRemoted returns the WazuhRemoted field if non-nil, zero value otherwise.

### GetWazuhRemotedOk

`func (o *WazuhDaemonsStatus) GetWazuhRemotedOk() (*DaemonStatus, bool)`

GetWazuhRemotedOk returns a tuple with the WazuhRemoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhRemoted

`func (o *WazuhDaemonsStatus) SetWazuhRemoted(v DaemonStatus)`

SetWazuhRemoted sets WazuhRemoted field to given value.

### HasWazuhRemoted

`func (o *WazuhDaemonsStatus) HasWazuhRemoted() bool`

HasWazuhRemoted returns a boolean if a field has been set.

### GetWazuhReportd

`func (o *WazuhDaemonsStatus) GetWazuhReportd() DaemonStatus`

GetWazuhReportd returns the WazuhReportd field if non-nil, zero value otherwise.

### GetWazuhReportdOk

`func (o *WazuhDaemonsStatus) GetWazuhReportdOk() (*DaemonStatus, bool)`

GetWazuhReportdOk returns a tuple with the WazuhReportd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhReportd

`func (o *WazuhDaemonsStatus) SetWazuhReportd(v DaemonStatus)`

SetWazuhReportd sets WazuhReportd field to given value.

### HasWazuhReportd

`func (o *WazuhDaemonsStatus) HasWazuhReportd() bool`

HasWazuhReportd returns a boolean if a field has been set.

### GetWazuhSyscheckd

`func (o *WazuhDaemonsStatus) GetWazuhSyscheckd() DaemonStatus`

GetWazuhSyscheckd returns the WazuhSyscheckd field if non-nil, zero value otherwise.

### GetWazuhSyscheckdOk

`func (o *WazuhDaemonsStatus) GetWazuhSyscheckdOk() (*DaemonStatus, bool)`

GetWazuhSyscheckdOk returns a tuple with the WazuhSyscheckd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhSyscheckd

`func (o *WazuhDaemonsStatus) SetWazuhSyscheckd(v DaemonStatus)`

SetWazuhSyscheckd sets WazuhSyscheckd field to given value.

### HasWazuhSyscheckd

`func (o *WazuhDaemonsStatus) HasWazuhSyscheckd() bool`

HasWazuhSyscheckd returns a boolean if a field has been set.

### GetWazuhApid

`func (o *WazuhDaemonsStatus) GetWazuhApid() DaemonStatus`

GetWazuhApid returns the WazuhApid field if non-nil, zero value otherwise.

### GetWazuhApidOk

`func (o *WazuhDaemonsStatus) GetWazuhApidOk() (*DaemonStatus, bool)`

GetWazuhApidOk returns a tuple with the WazuhApid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhApid

`func (o *WazuhDaemonsStatus) SetWazuhApid(v DaemonStatus)`

SetWazuhApid sets WazuhApid field to given value.

### HasWazuhApid

`func (o *WazuhDaemonsStatus) HasWazuhApid() bool`

HasWazuhApid returns a boolean if a field has been set.

### GetWazuhClusterd

`func (o *WazuhDaemonsStatus) GetWazuhClusterd() DaemonStatus`

GetWazuhClusterd returns the WazuhClusterd field if non-nil, zero value otherwise.

### GetWazuhClusterdOk

`func (o *WazuhDaemonsStatus) GetWazuhClusterdOk() (*DaemonStatus, bool)`

GetWazuhClusterdOk returns a tuple with the WazuhClusterd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhClusterd

`func (o *WazuhDaemonsStatus) SetWazuhClusterd(v DaemonStatus)`

SetWazuhClusterd sets WazuhClusterd field to given value.

### HasWazuhClusterd

`func (o *WazuhDaemonsStatus) HasWazuhClusterd() bool`

HasWazuhClusterd returns a boolean if a field has been set.

### GetWazuhDb

`func (o *WazuhDaemonsStatus) GetWazuhDb() DaemonStatus`

GetWazuhDb returns the WazuhDb field if non-nil, zero value otherwise.

### GetWazuhDbOk

`func (o *WazuhDaemonsStatus) GetWazuhDbOk() (*DaemonStatus, bool)`

GetWazuhDbOk returns a tuple with the WazuhDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhDb

`func (o *WazuhDaemonsStatus) SetWazuhDb(v DaemonStatus)`

SetWazuhDb sets WazuhDb field to given value.

### HasWazuhDb

`func (o *WazuhDaemonsStatus) HasWazuhDb() bool`

HasWazuhDb returns a boolean if a field has been set.

### GetWazuhModulesd

`func (o *WazuhDaemonsStatus) GetWazuhModulesd() DaemonStatus`

GetWazuhModulesd returns the WazuhModulesd field if non-nil, zero value otherwise.

### GetWazuhModulesdOk

`func (o *WazuhDaemonsStatus) GetWazuhModulesdOk() (*DaemonStatus, bool)`

GetWazuhModulesdOk returns a tuple with the WazuhModulesd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesd

`func (o *WazuhDaemonsStatus) SetWazuhModulesd(v DaemonStatus)`

SetWazuhModulesd sets WazuhModulesd field to given value.

### HasWazuhModulesd

`func (o *WazuhDaemonsStatus) HasWazuhModulesd() bool`

HasWazuhModulesd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


