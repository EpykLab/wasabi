# WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **int32** | Events coming from AWS module | [optional] 
**Azure** | Pointer to **int32** | Events coming from Azure module | [optional] 
**Ciscat** | Pointer to **int32** | Events coming from CIS-CAT module | [optional] 
**Command** | Pointer to **int32** | Events coming from command module | [optional] 
**Docker** | Pointer to **int32** | Events coming from Docker module | [optional] 
**Gcp** | Pointer to **int32** | Events coming from GCP module | [optional] 
**Github** | Pointer to **int32** | Events coming from GitHub module | [optional] 
**LogcollectorBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown.md) |  | [optional] 
**Office365** | Pointer to **int32** | Events coming from Office365 module | [optional] 
**MsGraph** | Pointer to **int32** | Events coming from ms-graph module | [optional] 
**Oscap** | Pointer to **int32** | Events coming from OSCAP module | [optional] 
**Osquery** | Pointer to **int32** | Events coming from OSQuery module | [optional] 
**Rootcheck** | Pointer to **int32** | Events coming from rootcheck (syscheckd) | [optional] 
**Sca** | Pointer to **int32** | Events coming from SCA module | [optional] 
**Syscheck** | Pointer to **int32** | Events coming from syscheckd | [optional] 
**Syscollector** | Pointer to **int32** | Events coming from syscollector module | [optional] 
**Upgrade** | Pointer to **int32** | Events coming from upgrade agent module (upgrade responses) | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAws() int32`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAwsOk() (*int32, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetAws(v int32)`

SetAws sets Aws field to given value.

### HasAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAzure() int32`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetAzureOk() (*int32, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetAzure(v int32)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCiscat() int32`

GetCiscat returns the Ciscat field if non-nil, zero value otherwise.

### GetCiscatOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCiscatOk() (*int32, bool)`

GetCiscatOk returns a tuple with the Ciscat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetCiscat(v int32)`

SetCiscat sets Ciscat field to given value.

### HasCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasCiscat() bool`

HasCiscat returns a boolean if a field has been set.

### GetCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCommand() int32`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetCommandOk() (*int32, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetCommand(v int32)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetDocker() int32`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetDockerOk() (*int32, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetDocker(v int32)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGcp() int32`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGcpOk() (*int32, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetGcp(v int32)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGithub() int32`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetGithubOk() (*int32, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetGithub(v int32)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetLogcollectorBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown`

GetLogcollectorBreakdown returns the LogcollectorBreakdown field if non-nil, zero value otherwise.

### GetLogcollectorBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetLogcollectorBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown, bool)`

GetLogcollectorBreakdownOk returns a tuple with the LogcollectorBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetLogcollectorBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdownLogcollectorBreakdown)`

SetLogcollectorBreakdown sets LogcollectorBreakdown field to given value.

### HasLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasLogcollectorBreakdown() bool`

HasLogcollectorBreakdown returns a boolean if a field has been set.

### GetOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOffice365() int32`

GetOffice365 returns the Office365 field if non-nil, zero value otherwise.

### GetOffice365Ok

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOffice365Ok() (*int32, bool)`

GetOffice365Ok returns a tuple with the Office365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOffice365(v int32)`

SetOffice365 sets Office365 field to given value.

### HasOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOffice365() bool`

HasOffice365 returns a boolean if a field has been set.

### GetMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetMsGraph() int32`

GetMsGraph returns the MsGraph field if non-nil, zero value otherwise.

### GetMsGraphOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetMsGraphOk() (*int32, bool)`

GetMsGraphOk returns a tuple with the MsGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetMsGraph(v int32)`

SetMsGraph sets MsGraph field to given value.

### HasMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasMsGraph() bool`

HasMsGraph returns a boolean if a field has been set.

### GetOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOscap() int32`

GetOscap returns the Oscap field if non-nil, zero value otherwise.

### GetOscapOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOscapOk() (*int32, bool)`

GetOscapOk returns a tuple with the Oscap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOscap(v int32)`

SetOscap sets Oscap field to given value.

### HasOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOscap() bool`

HasOscap returns a boolean if a field has been set.

### GetOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOsquery() int32`

GetOsquery returns the Osquery field if non-nil, zero value otherwise.

### GetOsqueryOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetOsqueryOk() (*int32, bool)`

GetOsqueryOk returns a tuple with the Osquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetOsquery(v int32)`

SetOsquery sets Osquery field to given value.

### HasOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasOsquery() bool`

HasOsquery returns a boolean if a field has been set.

### GetRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetRootcheck() int32`

GetRootcheck returns the Rootcheck field if non-nil, zero value otherwise.

### GetRootcheckOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetRootcheckOk() (*int32, bool)`

GetRootcheckOk returns a tuple with the Rootcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetRootcheck(v int32)`

SetRootcheck sets Rootcheck field to given value.

### HasRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasRootcheck() bool`

HasRootcheck returns a boolean if a field has been set.

### GetSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSca() int32`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetScaOk() (*int32, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSca(v int32)`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscheck() int32`

GetSyscheck returns the Syscheck field if non-nil, zero value otherwise.

### GetSyscheckOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscheckOk() (*int32, bool)`

GetSyscheckOk returns a tuple with the Syscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSyscheck(v int32)`

SetSyscheck sets Syscheck field to given value.

### HasSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSyscheck() bool`

HasSyscheck returns a boolean if a field has been set.

### GetSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscollector() int32`

GetSyscollector returns the Syscollector field if non-nil, zero value otherwise.

### GetSyscollectorOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetSyscollectorOk() (*int32, bool)`

GetSyscollectorOk returns a tuple with the Syscollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetSyscollector(v int32)`

SetSyscollector sets Syscollector field to given value.

### HasSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasSyscollector() bool`

HasSyscollector returns a boolean if a field has been set.

### GetUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetUpgrade() int32`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) GetUpgradeOk() (*int32, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) SetUpgrade(v int32)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


