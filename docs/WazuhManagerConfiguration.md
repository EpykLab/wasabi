# WazuhManagerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveResponse** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Agentless** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Alerts** | Pointer to **map[string]interface{}** |  | [optional] 
**Auth** | Pointer to **map[string]interface{}** |  | [optional] 
**Cluster** | Pointer to **map[string]interface{}** |  | [optional] 
**Command** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DatabaseOutput** | Pointer to **map[string]interface{}** |  | [optional] 
**EmailAlerts** | Pointer to **map[string]interface{}** |  | [optional] 
**GcpPubsub** | Pointer to **map[string]interface{}** |  | [optional] 
**Global** | Pointer to **map[string]interface{}** |  | [optional] 
**Integration** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Localfile** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Logging** | Pointer to **map[string]interface{}** |  | [optional] 
**Remote** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Reports** | Pointer to **map[string]interface{}** |  | [optional] 
**Rootcheck** | Pointer to **map[string]interface{}** |  | [optional] 
**Ruleset** | Pointer to **map[string]interface{}** |  | [optional] 
**Sca** | Pointer to **map[string]interface{}** |  | [optional] 
**Socket** | Pointer to **map[string]interface{}** |  | [optional] 
**Syscheck** | Pointer to **map[string]interface{}** |  | [optional] 
**SyslogOutput** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AwsS3** | Pointer to **map[string]interface{}** |  | [optional] 
**AzureLogs** | Pointer to **map[string]interface{}** |  | [optional] 
**CisCat** | Pointer to **map[string]interface{}** |  | [optional] 
**DockerListener** | Pointer to **map[string]interface{}** |  | [optional] 
**OpenScap** | Pointer to **map[string]interface{}** |  | [optional] 
**Osquery** | Pointer to **map[string]interface{}** |  | [optional] 
**Syscollector** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWazuhManagerConfiguration

`func NewWazuhManagerConfiguration() *WazuhManagerConfiguration`

NewWazuhManagerConfiguration instantiates a new WazuhManagerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhManagerConfigurationWithDefaults

`func NewWazuhManagerConfigurationWithDefaults() *WazuhManagerConfiguration`

NewWazuhManagerConfigurationWithDefaults instantiates a new WazuhManagerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveResponse

`func (o *WazuhManagerConfiguration) GetActiveResponse() []map[string]interface{}`

GetActiveResponse returns the ActiveResponse field if non-nil, zero value otherwise.

### GetActiveResponseOk

`func (o *WazuhManagerConfiguration) GetActiveResponseOk() (*[]map[string]interface{}, bool)`

GetActiveResponseOk returns a tuple with the ActiveResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveResponse

`func (o *WazuhManagerConfiguration) SetActiveResponse(v []map[string]interface{})`

SetActiveResponse sets ActiveResponse field to given value.

### HasActiveResponse

`func (o *WazuhManagerConfiguration) HasActiveResponse() bool`

HasActiveResponse returns a boolean if a field has been set.

### GetAgentless

`func (o *WazuhManagerConfiguration) GetAgentless() []map[string]interface{}`

GetAgentless returns the Agentless field if non-nil, zero value otherwise.

### GetAgentlessOk

`func (o *WazuhManagerConfiguration) GetAgentlessOk() (*[]map[string]interface{}, bool)`

GetAgentlessOk returns a tuple with the Agentless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentless

`func (o *WazuhManagerConfiguration) SetAgentless(v []map[string]interface{})`

SetAgentless sets Agentless field to given value.

### HasAgentless

`func (o *WazuhManagerConfiguration) HasAgentless() bool`

HasAgentless returns a boolean if a field has been set.

### GetAlerts

`func (o *WazuhManagerConfiguration) GetAlerts() map[string]interface{}`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *WazuhManagerConfiguration) GetAlertsOk() (*map[string]interface{}, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *WazuhManagerConfiguration) SetAlerts(v map[string]interface{})`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *WazuhManagerConfiguration) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetAuth

`func (o *WazuhManagerConfiguration) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *WazuhManagerConfiguration) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *WazuhManagerConfiguration) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *WazuhManagerConfiguration) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCluster

`func (o *WazuhManagerConfiguration) GetCluster() map[string]interface{}`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *WazuhManagerConfiguration) GetClusterOk() (*map[string]interface{}, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *WazuhManagerConfiguration) SetCluster(v map[string]interface{})`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *WazuhManagerConfiguration) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCommand

`func (o *WazuhManagerConfiguration) GetCommand() []map[string]interface{}`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WazuhManagerConfiguration) GetCommandOk() (*[]map[string]interface{}, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WazuhManagerConfiguration) SetCommand(v []map[string]interface{})`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WazuhManagerConfiguration) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDatabaseOutput

`func (o *WazuhManagerConfiguration) GetDatabaseOutput() map[string]interface{}`

GetDatabaseOutput returns the DatabaseOutput field if non-nil, zero value otherwise.

### GetDatabaseOutputOk

`func (o *WazuhManagerConfiguration) GetDatabaseOutputOk() (*map[string]interface{}, bool)`

GetDatabaseOutputOk returns a tuple with the DatabaseOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOutput

`func (o *WazuhManagerConfiguration) SetDatabaseOutput(v map[string]interface{})`

SetDatabaseOutput sets DatabaseOutput field to given value.

### HasDatabaseOutput

`func (o *WazuhManagerConfiguration) HasDatabaseOutput() bool`

HasDatabaseOutput returns a boolean if a field has been set.

### GetEmailAlerts

`func (o *WazuhManagerConfiguration) GetEmailAlerts() map[string]interface{}`

GetEmailAlerts returns the EmailAlerts field if non-nil, zero value otherwise.

### GetEmailAlertsOk

`func (o *WazuhManagerConfiguration) GetEmailAlertsOk() (*map[string]interface{}, bool)`

GetEmailAlertsOk returns a tuple with the EmailAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAlerts

`func (o *WazuhManagerConfiguration) SetEmailAlerts(v map[string]interface{})`

SetEmailAlerts sets EmailAlerts field to given value.

### HasEmailAlerts

`func (o *WazuhManagerConfiguration) HasEmailAlerts() bool`

HasEmailAlerts returns a boolean if a field has been set.

### GetGcpPubsub

`func (o *WazuhManagerConfiguration) GetGcpPubsub() map[string]interface{}`

GetGcpPubsub returns the GcpPubsub field if non-nil, zero value otherwise.

### GetGcpPubsubOk

`func (o *WazuhManagerConfiguration) GetGcpPubsubOk() (*map[string]interface{}, bool)`

GetGcpPubsubOk returns a tuple with the GcpPubsub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPubsub

`func (o *WazuhManagerConfiguration) SetGcpPubsub(v map[string]interface{})`

SetGcpPubsub sets GcpPubsub field to given value.

### HasGcpPubsub

`func (o *WazuhManagerConfiguration) HasGcpPubsub() bool`

HasGcpPubsub returns a boolean if a field has been set.

### GetGlobal

`func (o *WazuhManagerConfiguration) GetGlobal() map[string]interface{}`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *WazuhManagerConfiguration) GetGlobalOk() (*map[string]interface{}, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *WazuhManagerConfiguration) SetGlobal(v map[string]interface{})`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *WazuhManagerConfiguration) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetIntegration

`func (o *WazuhManagerConfiguration) GetIntegration() []map[string]interface{}`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *WazuhManagerConfiguration) GetIntegrationOk() (*[]map[string]interface{}, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *WazuhManagerConfiguration) SetIntegration(v []map[string]interface{})`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *WazuhManagerConfiguration) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetLabels

`func (o *WazuhManagerConfiguration) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WazuhManagerConfiguration) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WazuhManagerConfiguration) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WazuhManagerConfiguration) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocalfile

`func (o *WazuhManagerConfiguration) GetLocalfile() []map[string]interface{}`

GetLocalfile returns the Localfile field if non-nil, zero value otherwise.

### GetLocalfileOk

`func (o *WazuhManagerConfiguration) GetLocalfileOk() (*[]map[string]interface{}, bool)`

GetLocalfileOk returns a tuple with the Localfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalfile

`func (o *WazuhManagerConfiguration) SetLocalfile(v []map[string]interface{})`

SetLocalfile sets Localfile field to given value.

### HasLocalfile

`func (o *WazuhManagerConfiguration) HasLocalfile() bool`

HasLocalfile returns a boolean if a field has been set.

### GetLogging

`func (o *WazuhManagerConfiguration) GetLogging() map[string]interface{}`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *WazuhManagerConfiguration) GetLoggingOk() (*map[string]interface{}, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *WazuhManagerConfiguration) SetLogging(v map[string]interface{})`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *WazuhManagerConfiguration) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetRemote

`func (o *WazuhManagerConfiguration) GetRemote() []map[string]interface{}`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *WazuhManagerConfiguration) GetRemoteOk() (*[]map[string]interface{}, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *WazuhManagerConfiguration) SetRemote(v []map[string]interface{})`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *WazuhManagerConfiguration) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetReports

`func (o *WazuhManagerConfiguration) GetReports() map[string]interface{}`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *WazuhManagerConfiguration) GetReportsOk() (*map[string]interface{}, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *WazuhManagerConfiguration) SetReports(v map[string]interface{})`

SetReports sets Reports field to given value.

### HasReports

`func (o *WazuhManagerConfiguration) HasReports() bool`

HasReports returns a boolean if a field has been set.

### GetRootcheck

`func (o *WazuhManagerConfiguration) GetRootcheck() map[string]interface{}`

GetRootcheck returns the Rootcheck field if non-nil, zero value otherwise.

### GetRootcheckOk

`func (o *WazuhManagerConfiguration) GetRootcheckOk() (*map[string]interface{}, bool)`

GetRootcheckOk returns a tuple with the Rootcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheck

`func (o *WazuhManagerConfiguration) SetRootcheck(v map[string]interface{})`

SetRootcheck sets Rootcheck field to given value.

### HasRootcheck

`func (o *WazuhManagerConfiguration) HasRootcheck() bool`

HasRootcheck returns a boolean if a field has been set.

### GetRuleset

`func (o *WazuhManagerConfiguration) GetRuleset() map[string]interface{}`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *WazuhManagerConfiguration) GetRulesetOk() (*map[string]interface{}, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *WazuhManagerConfiguration) SetRuleset(v map[string]interface{})`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *WazuhManagerConfiguration) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### GetSca

`func (o *WazuhManagerConfiguration) GetSca() map[string]interface{}`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhManagerConfiguration) GetScaOk() (*map[string]interface{}, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhManagerConfiguration) SetSca(v map[string]interface{})`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhManagerConfiguration) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetSocket

`func (o *WazuhManagerConfiguration) GetSocket() map[string]interface{}`

GetSocket returns the Socket field if non-nil, zero value otherwise.

### GetSocketOk

`func (o *WazuhManagerConfiguration) GetSocketOk() (*map[string]interface{}, bool)`

GetSocketOk returns a tuple with the Socket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocket

`func (o *WazuhManagerConfiguration) SetSocket(v map[string]interface{})`

SetSocket sets Socket field to given value.

### HasSocket

`func (o *WazuhManagerConfiguration) HasSocket() bool`

HasSocket returns a boolean if a field has been set.

### GetSyscheck

`func (o *WazuhManagerConfiguration) GetSyscheck() map[string]interface{}`

GetSyscheck returns the Syscheck field if non-nil, zero value otherwise.

### GetSyscheckOk

`func (o *WazuhManagerConfiguration) GetSyscheckOk() (*map[string]interface{}, bool)`

GetSyscheckOk returns a tuple with the Syscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheck

`func (o *WazuhManagerConfiguration) SetSyscheck(v map[string]interface{})`

SetSyscheck sets Syscheck field to given value.

### HasSyscheck

`func (o *WazuhManagerConfiguration) HasSyscheck() bool`

HasSyscheck returns a boolean if a field has been set.

### GetSyslogOutput

`func (o *WazuhManagerConfiguration) GetSyslogOutput() []map[string]interface{}`

GetSyslogOutput returns the SyslogOutput field if non-nil, zero value otherwise.

### GetSyslogOutputOk

`func (o *WazuhManagerConfiguration) GetSyslogOutputOk() (*[]map[string]interface{}, bool)`

GetSyslogOutputOk returns a tuple with the SyslogOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogOutput

`func (o *WazuhManagerConfiguration) SetSyslogOutput(v []map[string]interface{})`

SetSyslogOutput sets SyslogOutput field to given value.

### HasSyslogOutput

`func (o *WazuhManagerConfiguration) HasSyslogOutput() bool`

HasSyslogOutput returns a boolean if a field has been set.

### GetAwsS3

`func (o *WazuhManagerConfiguration) GetAwsS3() map[string]interface{}`

GetAwsS3 returns the AwsS3 field if non-nil, zero value otherwise.

### GetAwsS3Ok

`func (o *WazuhManagerConfiguration) GetAwsS3Ok() (*map[string]interface{}, bool)`

GetAwsS3Ok returns a tuple with the AwsS3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3

`func (o *WazuhManagerConfiguration) SetAwsS3(v map[string]interface{})`

SetAwsS3 sets AwsS3 field to given value.

### HasAwsS3

`func (o *WazuhManagerConfiguration) HasAwsS3() bool`

HasAwsS3 returns a boolean if a field has been set.

### GetAzureLogs

`func (o *WazuhManagerConfiguration) GetAzureLogs() map[string]interface{}`

GetAzureLogs returns the AzureLogs field if non-nil, zero value otherwise.

### GetAzureLogsOk

`func (o *WazuhManagerConfiguration) GetAzureLogsOk() (*map[string]interface{}, bool)`

GetAzureLogsOk returns a tuple with the AzureLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLogs

`func (o *WazuhManagerConfiguration) SetAzureLogs(v map[string]interface{})`

SetAzureLogs sets AzureLogs field to given value.

### HasAzureLogs

`func (o *WazuhManagerConfiguration) HasAzureLogs() bool`

HasAzureLogs returns a boolean if a field has been set.

### GetCisCat

`func (o *WazuhManagerConfiguration) GetCisCat() map[string]interface{}`

GetCisCat returns the CisCat field if non-nil, zero value otherwise.

### GetCisCatOk

`func (o *WazuhManagerConfiguration) GetCisCatOk() (*map[string]interface{}, bool)`

GetCisCatOk returns a tuple with the CisCat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisCat

`func (o *WazuhManagerConfiguration) SetCisCat(v map[string]interface{})`

SetCisCat sets CisCat field to given value.

### HasCisCat

`func (o *WazuhManagerConfiguration) HasCisCat() bool`

HasCisCat returns a boolean if a field has been set.

### GetDockerListener

`func (o *WazuhManagerConfiguration) GetDockerListener() map[string]interface{}`

GetDockerListener returns the DockerListener field if non-nil, zero value otherwise.

### GetDockerListenerOk

`func (o *WazuhManagerConfiguration) GetDockerListenerOk() (*map[string]interface{}, bool)`

GetDockerListenerOk returns a tuple with the DockerListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerListener

`func (o *WazuhManagerConfiguration) SetDockerListener(v map[string]interface{})`

SetDockerListener sets DockerListener field to given value.

### HasDockerListener

`func (o *WazuhManagerConfiguration) HasDockerListener() bool`

HasDockerListener returns a boolean if a field has been set.

### GetOpenScap

`func (o *WazuhManagerConfiguration) GetOpenScap() map[string]interface{}`

GetOpenScap returns the OpenScap field if non-nil, zero value otherwise.

### GetOpenScapOk

`func (o *WazuhManagerConfiguration) GetOpenScapOk() (*map[string]interface{}, bool)`

GetOpenScapOk returns a tuple with the OpenScap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenScap

`func (o *WazuhManagerConfiguration) SetOpenScap(v map[string]interface{})`

SetOpenScap sets OpenScap field to given value.

### HasOpenScap

`func (o *WazuhManagerConfiguration) HasOpenScap() bool`

HasOpenScap returns a boolean if a field has been set.

### GetOsquery

`func (o *WazuhManagerConfiguration) GetOsquery() map[string]interface{}`

GetOsquery returns the Osquery field if non-nil, zero value otherwise.

### GetOsqueryOk

`func (o *WazuhManagerConfiguration) GetOsqueryOk() (*map[string]interface{}, bool)`

GetOsqueryOk returns a tuple with the Osquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsquery

`func (o *WazuhManagerConfiguration) SetOsquery(v map[string]interface{})`

SetOsquery sets Osquery field to given value.

### HasOsquery

`func (o *WazuhManagerConfiguration) HasOsquery() bool`

HasOsquery returns a boolean if a field has been set.

### GetSyscollector

`func (o *WazuhManagerConfiguration) GetSyscollector() map[string]interface{}`

GetSyscollector returns the Syscollector field if non-nil, zero value otherwise.

### GetSyscollectorOk

`func (o *WazuhManagerConfiguration) GetSyscollectorOk() (*map[string]interface{}, bool)`

GetSyscollectorOk returns a tuple with the Syscollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollector

`func (o *WazuhManagerConfiguration) SetSyscollector(v map[string]interface{})`

SetSyscollector sets Syscollector field to given value.

### HasSyscollector

`func (o *WazuhManagerConfiguration) HasSyscollector() bool`

HasSyscollector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


