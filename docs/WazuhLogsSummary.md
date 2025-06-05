# WazuhLogsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexerConnector** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhAgentlessd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhAnalysisd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhAuthd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhCsyslogd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhDbd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhExecd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhIntegratord** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhMaild** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhMonitord** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhLogcollector** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhRemoted** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhReportd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**Rootcheck** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhSyscheckd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**Sca** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhDb** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesd** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdAgentUpgrade** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdAwsS3** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdAzureLogs** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdCiscat** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdControl** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdCommand** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdContentManager** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdDatabase** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdDockerListener** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdDownload** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdOscap** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdOsquery** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdSyscollector** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdVulnerabilityScanner** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 
**WazuhModulesdTaskManager** | Pointer to [**LogSummary**](LogSummary.md) |  | [optional] 

## Methods

### NewWazuhLogsSummary

`func NewWazuhLogsSummary() *WazuhLogsSummary`

NewWazuhLogsSummary instantiates a new WazuhLogsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhLogsSummaryWithDefaults

`func NewWazuhLogsSummaryWithDefaults() *WazuhLogsSummary`

NewWazuhLogsSummaryWithDefaults instantiates a new WazuhLogsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexerConnector

`func (o *WazuhLogsSummary) GetIndexerConnector() LogSummary`

GetIndexerConnector returns the IndexerConnector field if non-nil, zero value otherwise.

### GetIndexerConnectorOk

`func (o *WazuhLogsSummary) GetIndexerConnectorOk() (*LogSummary, bool)`

GetIndexerConnectorOk returns a tuple with the IndexerConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerConnector

`func (o *WazuhLogsSummary) SetIndexerConnector(v LogSummary)`

SetIndexerConnector sets IndexerConnector field to given value.

### HasIndexerConnector

`func (o *WazuhLogsSummary) HasIndexerConnector() bool`

HasIndexerConnector returns a boolean if a field has been set.

### GetWazuhAgentlessd

`func (o *WazuhLogsSummary) GetWazuhAgentlessd() LogSummary`

GetWazuhAgentlessd returns the WazuhAgentlessd field if non-nil, zero value otherwise.

### GetWazuhAgentlessdOk

`func (o *WazuhLogsSummary) GetWazuhAgentlessdOk() (*LogSummary, bool)`

GetWazuhAgentlessdOk returns a tuple with the WazuhAgentlessd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAgentlessd

`func (o *WazuhLogsSummary) SetWazuhAgentlessd(v LogSummary)`

SetWazuhAgentlessd sets WazuhAgentlessd field to given value.

### HasWazuhAgentlessd

`func (o *WazuhLogsSummary) HasWazuhAgentlessd() bool`

HasWazuhAgentlessd returns a boolean if a field has been set.

### GetWazuhAnalysisd

`func (o *WazuhLogsSummary) GetWazuhAnalysisd() LogSummary`

GetWazuhAnalysisd returns the WazuhAnalysisd field if non-nil, zero value otherwise.

### GetWazuhAnalysisdOk

`func (o *WazuhLogsSummary) GetWazuhAnalysisdOk() (*LogSummary, bool)`

GetWazuhAnalysisdOk returns a tuple with the WazuhAnalysisd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAnalysisd

`func (o *WazuhLogsSummary) SetWazuhAnalysisd(v LogSummary)`

SetWazuhAnalysisd sets WazuhAnalysisd field to given value.

### HasWazuhAnalysisd

`func (o *WazuhLogsSummary) HasWazuhAnalysisd() bool`

HasWazuhAnalysisd returns a boolean if a field has been set.

### GetWazuhAuthd

`func (o *WazuhLogsSummary) GetWazuhAuthd() LogSummary`

GetWazuhAuthd returns the WazuhAuthd field if non-nil, zero value otherwise.

### GetWazuhAuthdOk

`func (o *WazuhLogsSummary) GetWazuhAuthdOk() (*LogSummary, bool)`

GetWazuhAuthdOk returns a tuple with the WazuhAuthd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhAuthd

`func (o *WazuhLogsSummary) SetWazuhAuthd(v LogSummary)`

SetWazuhAuthd sets WazuhAuthd field to given value.

### HasWazuhAuthd

`func (o *WazuhLogsSummary) HasWazuhAuthd() bool`

HasWazuhAuthd returns a boolean if a field has been set.

### GetWazuhCsyslogd

`func (o *WazuhLogsSummary) GetWazuhCsyslogd() LogSummary`

GetWazuhCsyslogd returns the WazuhCsyslogd field if non-nil, zero value otherwise.

### GetWazuhCsyslogdOk

`func (o *WazuhLogsSummary) GetWazuhCsyslogdOk() (*LogSummary, bool)`

GetWazuhCsyslogdOk returns a tuple with the WazuhCsyslogd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhCsyslogd

`func (o *WazuhLogsSummary) SetWazuhCsyslogd(v LogSummary)`

SetWazuhCsyslogd sets WazuhCsyslogd field to given value.

### HasWazuhCsyslogd

`func (o *WazuhLogsSummary) HasWazuhCsyslogd() bool`

HasWazuhCsyslogd returns a boolean if a field has been set.

### GetWazuhDbd

`func (o *WazuhLogsSummary) GetWazuhDbd() LogSummary`

GetWazuhDbd returns the WazuhDbd field if non-nil, zero value otherwise.

### GetWazuhDbdOk

`func (o *WazuhLogsSummary) GetWazuhDbdOk() (*LogSummary, bool)`

GetWazuhDbdOk returns a tuple with the WazuhDbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhDbd

`func (o *WazuhLogsSummary) SetWazuhDbd(v LogSummary)`

SetWazuhDbd sets WazuhDbd field to given value.

### HasWazuhDbd

`func (o *WazuhLogsSummary) HasWazuhDbd() bool`

HasWazuhDbd returns a boolean if a field has been set.

### GetWazuhExecd

`func (o *WazuhLogsSummary) GetWazuhExecd() LogSummary`

GetWazuhExecd returns the WazuhExecd field if non-nil, zero value otherwise.

### GetWazuhExecdOk

`func (o *WazuhLogsSummary) GetWazuhExecdOk() (*LogSummary, bool)`

GetWazuhExecdOk returns a tuple with the WazuhExecd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhExecd

`func (o *WazuhLogsSummary) SetWazuhExecd(v LogSummary)`

SetWazuhExecd sets WazuhExecd field to given value.

### HasWazuhExecd

`func (o *WazuhLogsSummary) HasWazuhExecd() bool`

HasWazuhExecd returns a boolean if a field has been set.

### GetWazuhIntegratord

`func (o *WazuhLogsSummary) GetWazuhIntegratord() LogSummary`

GetWazuhIntegratord returns the WazuhIntegratord field if non-nil, zero value otherwise.

### GetWazuhIntegratordOk

`func (o *WazuhLogsSummary) GetWazuhIntegratordOk() (*LogSummary, bool)`

GetWazuhIntegratordOk returns a tuple with the WazuhIntegratord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhIntegratord

`func (o *WazuhLogsSummary) SetWazuhIntegratord(v LogSummary)`

SetWazuhIntegratord sets WazuhIntegratord field to given value.

### HasWazuhIntegratord

`func (o *WazuhLogsSummary) HasWazuhIntegratord() bool`

HasWazuhIntegratord returns a boolean if a field has been set.

### GetWazuhMaild

`func (o *WazuhLogsSummary) GetWazuhMaild() LogSummary`

GetWazuhMaild returns the WazuhMaild field if non-nil, zero value otherwise.

### GetWazuhMaildOk

`func (o *WazuhLogsSummary) GetWazuhMaildOk() (*LogSummary, bool)`

GetWazuhMaildOk returns a tuple with the WazuhMaild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhMaild

`func (o *WazuhLogsSummary) SetWazuhMaild(v LogSummary)`

SetWazuhMaild sets WazuhMaild field to given value.

### HasWazuhMaild

`func (o *WazuhLogsSummary) HasWazuhMaild() bool`

HasWazuhMaild returns a boolean if a field has been set.

### GetWazuhMonitord

`func (o *WazuhLogsSummary) GetWazuhMonitord() LogSummary`

GetWazuhMonitord returns the WazuhMonitord field if non-nil, zero value otherwise.

### GetWazuhMonitordOk

`func (o *WazuhLogsSummary) GetWazuhMonitordOk() (*LogSummary, bool)`

GetWazuhMonitordOk returns a tuple with the WazuhMonitord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhMonitord

`func (o *WazuhLogsSummary) SetWazuhMonitord(v LogSummary)`

SetWazuhMonitord sets WazuhMonitord field to given value.

### HasWazuhMonitord

`func (o *WazuhLogsSummary) HasWazuhMonitord() bool`

HasWazuhMonitord returns a boolean if a field has been set.

### GetWazuhLogcollector

`func (o *WazuhLogsSummary) GetWazuhLogcollector() LogSummary`

GetWazuhLogcollector returns the WazuhLogcollector field if non-nil, zero value otherwise.

### GetWazuhLogcollectorOk

`func (o *WazuhLogsSummary) GetWazuhLogcollectorOk() (*LogSummary, bool)`

GetWazuhLogcollectorOk returns a tuple with the WazuhLogcollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhLogcollector

`func (o *WazuhLogsSummary) SetWazuhLogcollector(v LogSummary)`

SetWazuhLogcollector sets WazuhLogcollector field to given value.

### HasWazuhLogcollector

`func (o *WazuhLogsSummary) HasWazuhLogcollector() bool`

HasWazuhLogcollector returns a boolean if a field has been set.

### GetWazuhRemoted

`func (o *WazuhLogsSummary) GetWazuhRemoted() LogSummary`

GetWazuhRemoted returns the WazuhRemoted field if non-nil, zero value otherwise.

### GetWazuhRemotedOk

`func (o *WazuhLogsSummary) GetWazuhRemotedOk() (*LogSummary, bool)`

GetWazuhRemotedOk returns a tuple with the WazuhRemoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhRemoted

`func (o *WazuhLogsSummary) SetWazuhRemoted(v LogSummary)`

SetWazuhRemoted sets WazuhRemoted field to given value.

### HasWazuhRemoted

`func (o *WazuhLogsSummary) HasWazuhRemoted() bool`

HasWazuhRemoted returns a boolean if a field has been set.

### GetWazuhReportd

`func (o *WazuhLogsSummary) GetWazuhReportd() LogSummary`

GetWazuhReportd returns the WazuhReportd field if non-nil, zero value otherwise.

### GetWazuhReportdOk

`func (o *WazuhLogsSummary) GetWazuhReportdOk() (*LogSummary, bool)`

GetWazuhReportdOk returns a tuple with the WazuhReportd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhReportd

`func (o *WazuhLogsSummary) SetWazuhReportd(v LogSummary)`

SetWazuhReportd sets WazuhReportd field to given value.

### HasWazuhReportd

`func (o *WazuhLogsSummary) HasWazuhReportd() bool`

HasWazuhReportd returns a boolean if a field has been set.

### GetRootcheck

`func (o *WazuhLogsSummary) GetRootcheck() LogSummary`

GetRootcheck returns the Rootcheck field if non-nil, zero value otherwise.

### GetRootcheckOk

`func (o *WazuhLogsSummary) GetRootcheckOk() (*LogSummary, bool)`

GetRootcheckOk returns a tuple with the Rootcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheck

`func (o *WazuhLogsSummary) SetRootcheck(v LogSummary)`

SetRootcheck sets Rootcheck field to given value.

### HasRootcheck

`func (o *WazuhLogsSummary) HasRootcheck() bool`

HasRootcheck returns a boolean if a field has been set.

### GetWazuhSyscheckd

`func (o *WazuhLogsSummary) GetWazuhSyscheckd() LogSummary`

GetWazuhSyscheckd returns the WazuhSyscheckd field if non-nil, zero value otherwise.

### GetWazuhSyscheckdOk

`func (o *WazuhLogsSummary) GetWazuhSyscheckdOk() (*LogSummary, bool)`

GetWazuhSyscheckdOk returns a tuple with the WazuhSyscheckd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhSyscheckd

`func (o *WazuhLogsSummary) SetWazuhSyscheckd(v LogSummary)`

SetWazuhSyscheckd sets WazuhSyscheckd field to given value.

### HasWazuhSyscheckd

`func (o *WazuhLogsSummary) HasWazuhSyscheckd() bool`

HasWazuhSyscheckd returns a boolean if a field has been set.

### GetSca

`func (o *WazuhLogsSummary) GetSca() LogSummary`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhLogsSummary) GetScaOk() (*LogSummary, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhLogsSummary) SetSca(v LogSummary)`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhLogsSummary) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetWazuhDb

`func (o *WazuhLogsSummary) GetWazuhDb() LogSummary`

GetWazuhDb returns the WazuhDb field if non-nil, zero value otherwise.

### GetWazuhDbOk

`func (o *WazuhLogsSummary) GetWazuhDbOk() (*LogSummary, bool)`

GetWazuhDbOk returns a tuple with the WazuhDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhDb

`func (o *WazuhLogsSummary) SetWazuhDb(v LogSummary)`

SetWazuhDb sets WazuhDb field to given value.

### HasWazuhDb

`func (o *WazuhLogsSummary) HasWazuhDb() bool`

HasWazuhDb returns a boolean if a field has been set.

### GetWazuhModulesd

`func (o *WazuhLogsSummary) GetWazuhModulesd() LogSummary`

GetWazuhModulesd returns the WazuhModulesd field if non-nil, zero value otherwise.

### GetWazuhModulesdOk

`func (o *WazuhLogsSummary) GetWazuhModulesdOk() (*LogSummary, bool)`

GetWazuhModulesdOk returns a tuple with the WazuhModulesd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesd

`func (o *WazuhLogsSummary) SetWazuhModulesd(v LogSummary)`

SetWazuhModulesd sets WazuhModulesd field to given value.

### HasWazuhModulesd

`func (o *WazuhLogsSummary) HasWazuhModulesd() bool`

HasWazuhModulesd returns a boolean if a field has been set.

### GetWazuhModulesdAgentUpgrade

`func (o *WazuhLogsSummary) GetWazuhModulesdAgentUpgrade() LogSummary`

GetWazuhModulesdAgentUpgrade returns the WazuhModulesdAgentUpgrade field if non-nil, zero value otherwise.

### GetWazuhModulesdAgentUpgradeOk

`func (o *WazuhLogsSummary) GetWazuhModulesdAgentUpgradeOk() (*LogSummary, bool)`

GetWazuhModulesdAgentUpgradeOk returns a tuple with the WazuhModulesdAgentUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdAgentUpgrade

`func (o *WazuhLogsSummary) SetWazuhModulesdAgentUpgrade(v LogSummary)`

SetWazuhModulesdAgentUpgrade sets WazuhModulesdAgentUpgrade field to given value.

### HasWazuhModulesdAgentUpgrade

`func (o *WazuhLogsSummary) HasWazuhModulesdAgentUpgrade() bool`

HasWazuhModulesdAgentUpgrade returns a boolean if a field has been set.

### GetWazuhModulesdAwsS3

`func (o *WazuhLogsSummary) GetWazuhModulesdAwsS3() LogSummary`

GetWazuhModulesdAwsS3 returns the WazuhModulesdAwsS3 field if non-nil, zero value otherwise.

### GetWazuhModulesdAwsS3Ok

`func (o *WazuhLogsSummary) GetWazuhModulesdAwsS3Ok() (*LogSummary, bool)`

GetWazuhModulesdAwsS3Ok returns a tuple with the WazuhModulesdAwsS3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdAwsS3

`func (o *WazuhLogsSummary) SetWazuhModulesdAwsS3(v LogSummary)`

SetWazuhModulesdAwsS3 sets WazuhModulesdAwsS3 field to given value.

### HasWazuhModulesdAwsS3

`func (o *WazuhLogsSummary) HasWazuhModulesdAwsS3() bool`

HasWazuhModulesdAwsS3 returns a boolean if a field has been set.

### GetWazuhModulesdAzureLogs

`func (o *WazuhLogsSummary) GetWazuhModulesdAzureLogs() LogSummary`

GetWazuhModulesdAzureLogs returns the WazuhModulesdAzureLogs field if non-nil, zero value otherwise.

### GetWazuhModulesdAzureLogsOk

`func (o *WazuhLogsSummary) GetWazuhModulesdAzureLogsOk() (*LogSummary, bool)`

GetWazuhModulesdAzureLogsOk returns a tuple with the WazuhModulesdAzureLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdAzureLogs

`func (o *WazuhLogsSummary) SetWazuhModulesdAzureLogs(v LogSummary)`

SetWazuhModulesdAzureLogs sets WazuhModulesdAzureLogs field to given value.

### HasWazuhModulesdAzureLogs

`func (o *WazuhLogsSummary) HasWazuhModulesdAzureLogs() bool`

HasWazuhModulesdAzureLogs returns a boolean if a field has been set.

### GetWazuhModulesdCiscat

`func (o *WazuhLogsSummary) GetWazuhModulesdCiscat() LogSummary`

GetWazuhModulesdCiscat returns the WazuhModulesdCiscat field if non-nil, zero value otherwise.

### GetWazuhModulesdCiscatOk

`func (o *WazuhLogsSummary) GetWazuhModulesdCiscatOk() (*LogSummary, bool)`

GetWazuhModulesdCiscatOk returns a tuple with the WazuhModulesdCiscat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdCiscat

`func (o *WazuhLogsSummary) SetWazuhModulesdCiscat(v LogSummary)`

SetWazuhModulesdCiscat sets WazuhModulesdCiscat field to given value.

### HasWazuhModulesdCiscat

`func (o *WazuhLogsSummary) HasWazuhModulesdCiscat() bool`

HasWazuhModulesdCiscat returns a boolean if a field has been set.

### GetWazuhModulesdControl

`func (o *WazuhLogsSummary) GetWazuhModulesdControl() LogSummary`

GetWazuhModulesdControl returns the WazuhModulesdControl field if non-nil, zero value otherwise.

### GetWazuhModulesdControlOk

`func (o *WazuhLogsSummary) GetWazuhModulesdControlOk() (*LogSummary, bool)`

GetWazuhModulesdControlOk returns a tuple with the WazuhModulesdControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdControl

`func (o *WazuhLogsSummary) SetWazuhModulesdControl(v LogSummary)`

SetWazuhModulesdControl sets WazuhModulesdControl field to given value.

### HasWazuhModulesdControl

`func (o *WazuhLogsSummary) HasWazuhModulesdControl() bool`

HasWazuhModulesdControl returns a boolean if a field has been set.

### GetWazuhModulesdCommand

`func (o *WazuhLogsSummary) GetWazuhModulesdCommand() LogSummary`

GetWazuhModulesdCommand returns the WazuhModulesdCommand field if non-nil, zero value otherwise.

### GetWazuhModulesdCommandOk

`func (o *WazuhLogsSummary) GetWazuhModulesdCommandOk() (*LogSummary, bool)`

GetWazuhModulesdCommandOk returns a tuple with the WazuhModulesdCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdCommand

`func (o *WazuhLogsSummary) SetWazuhModulesdCommand(v LogSummary)`

SetWazuhModulesdCommand sets WazuhModulesdCommand field to given value.

### HasWazuhModulesdCommand

`func (o *WazuhLogsSummary) HasWazuhModulesdCommand() bool`

HasWazuhModulesdCommand returns a boolean if a field has been set.

### GetWazuhModulesdContentManager

`func (o *WazuhLogsSummary) GetWazuhModulesdContentManager() LogSummary`

GetWazuhModulesdContentManager returns the WazuhModulesdContentManager field if non-nil, zero value otherwise.

### GetWazuhModulesdContentManagerOk

`func (o *WazuhLogsSummary) GetWazuhModulesdContentManagerOk() (*LogSummary, bool)`

GetWazuhModulesdContentManagerOk returns a tuple with the WazuhModulesdContentManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdContentManager

`func (o *WazuhLogsSummary) SetWazuhModulesdContentManager(v LogSummary)`

SetWazuhModulesdContentManager sets WazuhModulesdContentManager field to given value.

### HasWazuhModulesdContentManager

`func (o *WazuhLogsSummary) HasWazuhModulesdContentManager() bool`

HasWazuhModulesdContentManager returns a boolean if a field has been set.

### GetWazuhModulesdDatabase

`func (o *WazuhLogsSummary) GetWazuhModulesdDatabase() LogSummary`

GetWazuhModulesdDatabase returns the WazuhModulesdDatabase field if non-nil, zero value otherwise.

### GetWazuhModulesdDatabaseOk

`func (o *WazuhLogsSummary) GetWazuhModulesdDatabaseOk() (*LogSummary, bool)`

GetWazuhModulesdDatabaseOk returns a tuple with the WazuhModulesdDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdDatabase

`func (o *WazuhLogsSummary) SetWazuhModulesdDatabase(v LogSummary)`

SetWazuhModulesdDatabase sets WazuhModulesdDatabase field to given value.

### HasWazuhModulesdDatabase

`func (o *WazuhLogsSummary) HasWazuhModulesdDatabase() bool`

HasWazuhModulesdDatabase returns a boolean if a field has been set.

### GetWazuhModulesdDockerListener

`func (o *WazuhLogsSummary) GetWazuhModulesdDockerListener() LogSummary`

GetWazuhModulesdDockerListener returns the WazuhModulesdDockerListener field if non-nil, zero value otherwise.

### GetWazuhModulesdDockerListenerOk

`func (o *WazuhLogsSummary) GetWazuhModulesdDockerListenerOk() (*LogSummary, bool)`

GetWazuhModulesdDockerListenerOk returns a tuple with the WazuhModulesdDockerListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdDockerListener

`func (o *WazuhLogsSummary) SetWazuhModulesdDockerListener(v LogSummary)`

SetWazuhModulesdDockerListener sets WazuhModulesdDockerListener field to given value.

### HasWazuhModulesdDockerListener

`func (o *WazuhLogsSummary) HasWazuhModulesdDockerListener() bool`

HasWazuhModulesdDockerListener returns a boolean if a field has been set.

### GetWazuhModulesdDownload

`func (o *WazuhLogsSummary) GetWazuhModulesdDownload() LogSummary`

GetWazuhModulesdDownload returns the WazuhModulesdDownload field if non-nil, zero value otherwise.

### GetWazuhModulesdDownloadOk

`func (o *WazuhLogsSummary) GetWazuhModulesdDownloadOk() (*LogSummary, bool)`

GetWazuhModulesdDownloadOk returns a tuple with the WazuhModulesdDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdDownload

`func (o *WazuhLogsSummary) SetWazuhModulesdDownload(v LogSummary)`

SetWazuhModulesdDownload sets WazuhModulesdDownload field to given value.

### HasWazuhModulesdDownload

`func (o *WazuhLogsSummary) HasWazuhModulesdDownload() bool`

HasWazuhModulesdDownload returns a boolean if a field has been set.

### GetWazuhModulesdOscap

`func (o *WazuhLogsSummary) GetWazuhModulesdOscap() LogSummary`

GetWazuhModulesdOscap returns the WazuhModulesdOscap field if non-nil, zero value otherwise.

### GetWazuhModulesdOscapOk

`func (o *WazuhLogsSummary) GetWazuhModulesdOscapOk() (*LogSummary, bool)`

GetWazuhModulesdOscapOk returns a tuple with the WazuhModulesdOscap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdOscap

`func (o *WazuhLogsSummary) SetWazuhModulesdOscap(v LogSummary)`

SetWazuhModulesdOscap sets WazuhModulesdOscap field to given value.

### HasWazuhModulesdOscap

`func (o *WazuhLogsSummary) HasWazuhModulesdOscap() bool`

HasWazuhModulesdOscap returns a boolean if a field has been set.

### GetWazuhModulesdOsquery

`func (o *WazuhLogsSummary) GetWazuhModulesdOsquery() LogSummary`

GetWazuhModulesdOsquery returns the WazuhModulesdOsquery field if non-nil, zero value otherwise.

### GetWazuhModulesdOsqueryOk

`func (o *WazuhLogsSummary) GetWazuhModulesdOsqueryOk() (*LogSummary, bool)`

GetWazuhModulesdOsqueryOk returns a tuple with the WazuhModulesdOsquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdOsquery

`func (o *WazuhLogsSummary) SetWazuhModulesdOsquery(v LogSummary)`

SetWazuhModulesdOsquery sets WazuhModulesdOsquery field to given value.

### HasWazuhModulesdOsquery

`func (o *WazuhLogsSummary) HasWazuhModulesdOsquery() bool`

HasWazuhModulesdOsquery returns a boolean if a field has been set.

### GetWazuhModulesdSyscollector

`func (o *WazuhLogsSummary) GetWazuhModulesdSyscollector() LogSummary`

GetWazuhModulesdSyscollector returns the WazuhModulesdSyscollector field if non-nil, zero value otherwise.

### GetWazuhModulesdSyscollectorOk

`func (o *WazuhLogsSummary) GetWazuhModulesdSyscollectorOk() (*LogSummary, bool)`

GetWazuhModulesdSyscollectorOk returns a tuple with the WazuhModulesdSyscollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdSyscollector

`func (o *WazuhLogsSummary) SetWazuhModulesdSyscollector(v LogSummary)`

SetWazuhModulesdSyscollector sets WazuhModulesdSyscollector field to given value.

### HasWazuhModulesdSyscollector

`func (o *WazuhLogsSummary) HasWazuhModulesdSyscollector() bool`

HasWazuhModulesdSyscollector returns a boolean if a field has been set.

### GetWazuhModulesdVulnerabilityScanner

`func (o *WazuhLogsSummary) GetWazuhModulesdVulnerabilityScanner() LogSummary`

GetWazuhModulesdVulnerabilityScanner returns the WazuhModulesdVulnerabilityScanner field if non-nil, zero value otherwise.

### GetWazuhModulesdVulnerabilityScannerOk

`func (o *WazuhLogsSummary) GetWazuhModulesdVulnerabilityScannerOk() (*LogSummary, bool)`

GetWazuhModulesdVulnerabilityScannerOk returns a tuple with the WazuhModulesdVulnerabilityScanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdVulnerabilityScanner

`func (o *WazuhLogsSummary) SetWazuhModulesdVulnerabilityScanner(v LogSummary)`

SetWazuhModulesdVulnerabilityScanner sets WazuhModulesdVulnerabilityScanner field to given value.

### HasWazuhModulesdVulnerabilityScanner

`func (o *WazuhLogsSummary) HasWazuhModulesdVulnerabilityScanner() bool`

HasWazuhModulesdVulnerabilityScanner returns a boolean if a field has been set.

### GetWazuhModulesdTaskManager

`func (o *WazuhLogsSummary) GetWazuhModulesdTaskManager() LogSummary`

GetWazuhModulesdTaskManager returns the WazuhModulesdTaskManager field if non-nil, zero value otherwise.

### GetWazuhModulesdTaskManagerOk

`func (o *WazuhLogsSummary) GetWazuhModulesdTaskManagerOk() (*LogSummary, bool)`

GetWazuhModulesdTaskManagerOk returns a tuple with the WazuhModulesdTaskManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhModulesdTaskManager

`func (o *WazuhLogsSummary) SetWazuhModulesdTaskManager(v LogSummary)`

SetWazuhModulesdTaskManager sets WazuhModulesdTaskManager field to given value.

### HasWazuhModulesdTaskManager

`func (o *WazuhLogsSummary) HasWazuhModulesdTaskManager() bool`

HasWazuhModulesdTaskManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


