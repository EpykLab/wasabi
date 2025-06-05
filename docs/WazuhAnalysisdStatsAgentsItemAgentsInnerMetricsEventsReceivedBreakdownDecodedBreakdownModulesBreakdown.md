# WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **int32** | Events coming from AWS module (this agent) | [optional] 
**Azure** | Pointer to **int32** | Events coming from Azure module (this agent) | [optional] 
**Ciscat** | Pointer to **int32** | Events coming from CIS-CAT module (this agent) | [optional] 
**Command** | Pointer to **int32** | Events coming from command module (this agent) | [optional] 
**Docker** | Pointer to **int32** | Events coming from Docker module (this agent) | [optional] 
**Gcp** | Pointer to **int32** | Events coming from GCP module (this agent) | [optional] 
**Github** | Pointer to **int32** | Events coming from GitHub module (this agent) | [optional] 
**LogcollectorBreakdown** | Pointer to [**WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown**](WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown.md) |  | [optional] 
**Office365** | Pointer to **int32** | Events coming from the Office365 module (this agent) | [optional] 
**MsGraph** | Pointer to **int32** | Events coming from the ms-graph module (this agent) | [optional] 
**Oscap** | Pointer to **int32** | Events coming from the OSCAP module (this agent) | [optional] 
**Osquery** | Pointer to **int32** | Events coming from the OSQuery module (this agent) | [optional] 
**Rootcheck** | Pointer to **int32** | Events coming from the rootcheck (syscheckd) (this agent) | [optional] 
**Sca** | Pointer to **int32** | Events coming from the SCA module (this agent) | [optional] 
**Syscheck** | Pointer to **int32** | Events coming from the syscheckd (this agent) | [optional] 
**Syscollector** | Pointer to **int32** | Events coming from the syscollector module (this agent) | [optional] 
**Upgrade** | Pointer to **int32** | Events coming from the upgrade agent module (this agent) | [optional] 
**Vulnerability** | Pointer to **int32** | Events coming from the vulnerability detector module (this agent) | [optional] 

## Methods

### NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown

`func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults

`func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAws() int32`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAwsOk() (*int32, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetAws(v int32)`

SetAws sets Aws field to given value.

### HasAws

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAzure() int32`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAzureOk() (*int32, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetAzure(v int32)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetCiscat

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCiscat() int32`

GetCiscat returns the Ciscat field if non-nil, zero value otherwise.

### GetCiscatOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCiscatOk() (*int32, bool)`

GetCiscatOk returns a tuple with the Ciscat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscat

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetCiscat(v int32)`

SetCiscat sets Ciscat field to given value.

### HasCiscat

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasCiscat() bool`

HasCiscat returns a boolean if a field has been set.

### GetCommand

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCommand() int32`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCommandOk() (*int32, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetCommand(v int32)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDocker

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetDocker() int32`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetDockerOk() (*int32, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetDocker(v int32)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGcp

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGcp() int32`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGcpOk() (*int32, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetGcp(v int32)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetGithub

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGithub() int32`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGithubOk() (*int32, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetGithub(v int32)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetLogcollectorBreakdown() WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown`

GetLogcollectorBreakdown returns the LogcollectorBreakdown field if non-nil, zero value otherwise.

### GetLogcollectorBreakdownOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetLogcollectorBreakdownOk() (*WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown, bool)`

GetLogcollectorBreakdownOk returns a tuple with the LogcollectorBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetLogcollectorBreakdown(v WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown)`

SetLogcollectorBreakdown sets LogcollectorBreakdown field to given value.

### HasLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasLogcollectorBreakdown() bool`

HasLogcollectorBreakdown returns a boolean if a field has been set.

### GetOffice365

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOffice365() int32`

GetOffice365 returns the Office365 field if non-nil, zero value otherwise.

### GetOffice365Ok

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOffice365Ok() (*int32, bool)`

GetOffice365Ok returns a tuple with the Office365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOffice365(v int32)`

SetOffice365 sets Office365 field to given value.

### HasOffice365

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOffice365() bool`

HasOffice365 returns a boolean if a field has been set.

### GetMsGraph

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetMsGraph() int32`

GetMsGraph returns the MsGraph field if non-nil, zero value otherwise.

### GetMsGraphOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetMsGraphOk() (*int32, bool)`

GetMsGraphOk returns a tuple with the MsGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGraph

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetMsGraph(v int32)`

SetMsGraph sets MsGraph field to given value.

### HasMsGraph

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasMsGraph() bool`

HasMsGraph returns a boolean if a field has been set.

### GetOscap

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOscap() int32`

GetOscap returns the Oscap field if non-nil, zero value otherwise.

### GetOscapOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOscapOk() (*int32, bool)`

GetOscapOk returns a tuple with the Oscap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOscap

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOscap(v int32)`

SetOscap sets Oscap field to given value.

### HasOscap

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOscap() bool`

HasOscap returns a boolean if a field has been set.

### GetOsquery

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOsquery() int32`

GetOsquery returns the Osquery field if non-nil, zero value otherwise.

### GetOsqueryOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOsqueryOk() (*int32, bool)`

GetOsqueryOk returns a tuple with the Osquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsquery

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOsquery(v int32)`

SetOsquery sets Osquery field to given value.

### HasOsquery

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOsquery() bool`

HasOsquery returns a boolean if a field has been set.

### GetRootcheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetRootcheck() int32`

GetRootcheck returns the Rootcheck field if non-nil, zero value otherwise.

### GetRootcheckOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetRootcheckOk() (*int32, bool)`

GetRootcheckOk returns a tuple with the Rootcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetRootcheck(v int32)`

SetRootcheck sets Rootcheck field to given value.

### HasRootcheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasRootcheck() bool`

HasRootcheck returns a boolean if a field has been set.

### GetSca

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSca() int32`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetScaOk() (*int32, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSca(v int32)`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetSyscheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscheck() int32`

GetSyscheck returns the Syscheck field if non-nil, zero value otherwise.

### GetSyscheckOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscheckOk() (*int32, bool)`

GetSyscheckOk returns a tuple with the Syscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSyscheck(v int32)`

SetSyscheck sets Syscheck field to given value.

### HasSyscheck

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSyscheck() bool`

HasSyscheck returns a boolean if a field has been set.

### GetSyscollector

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscollector() int32`

GetSyscollector returns the Syscollector field if non-nil, zero value otherwise.

### GetSyscollectorOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscollectorOk() (*int32, bool)`

GetSyscollectorOk returns a tuple with the Syscollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollector

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSyscollector(v int32)`

SetSyscollector sets Syscollector field to given value.

### HasSyscollector

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSyscollector() bool`

HasSyscollector returns a boolean if a field has been set.

### GetUpgrade

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetUpgrade() int32`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetUpgradeOk() (*int32, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetUpgrade(v int32)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetVulnerability

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetVulnerability() int32`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetVulnerabilityOk() (*int32, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetVulnerability(v int32)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


