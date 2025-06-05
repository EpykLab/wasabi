# WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to **int32** | Events discarded from AWS module because the queue was full | [optional] 
**Azure** | Pointer to **int32** | Events discarded from Azure module because the queue was full | [optional] 
**Ciscat** | Pointer to **int32** | Events discarded from CIS-CAT module because the queue was full | [optional] 
**Command** | Pointer to **int32** | Events discarded from command module because the queue was full | [optional] 
**Docker** | Pointer to **int32** | Events discarded from Docker module because the queue was full | [optional] 
**Gcp** | Pointer to **int32** | Events discarded from GCP module because the queue was full | [optional] 
**Github** | Pointer to **int32** | Events discarded from GitHub module because the queue was full | [optional] 
**LogcollectorBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownLogcollectorBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownLogcollectorBreakdown.md) |  | [optional] 
**Office365** | Pointer to **int32** | Events discarded from Office365 module because the queue was full | [optional] 
**MsGraph** | Pointer to **int32** | Events discarded from ms-graph module because the queue was full | [optional] 
**Oscap** | Pointer to **int32** | Events discarded from OSCAP module because the queue was full | [optional] 
**Osquery** | Pointer to **int32** | Events discarded from OSQuery module because the queue was full | [optional] 
**Rootcheck** | Pointer to **int32** | Events discarded from rootcheck (syscheckd) because the queue was full | [optional] 
**Sca** | Pointer to **int32** | Events discarded from SCA module because the queue was full | [optional] 
**Syscheck** | Pointer to **int32** | Events discarded from syscheckd because the queue was full | [optional] 
**Syscollector** | Pointer to **int32** | Events discarded from syscollector module because the queue was full | [optional] 
**Upgrade** | Pointer to **int32** | Events discarded from upgrade agent module because the queue was full | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetAws() int32`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetAwsOk() (*int32, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetAws(v int32)`

SetAws sets Aws field to given value.

### HasAws

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasAws() bool`

HasAws returns a boolean if a field has been set.

### GetAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetAzure() int32`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetAzureOk() (*int32, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetAzure(v int32)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### GetCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetCiscat() int32`

GetCiscat returns the Ciscat field if non-nil, zero value otherwise.

### GetCiscatOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetCiscatOk() (*int32, bool)`

GetCiscatOk returns a tuple with the Ciscat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetCiscat(v int32)`

SetCiscat sets Ciscat field to given value.

### HasCiscat

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasCiscat() bool`

HasCiscat returns a boolean if a field has been set.

### GetCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetCommand() int32`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetCommandOk() (*int32, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetCommand(v int32)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetDocker() int32`

GetDocker returns the Docker field if non-nil, zero value otherwise.

### GetDockerOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetDockerOk() (*int32, bool)`

GetDockerOk returns a tuple with the Docker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetDocker(v int32)`

SetDocker sets Docker field to given value.

### HasDocker

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasDocker() bool`

HasDocker returns a boolean if a field has been set.

### GetGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetGcp() int32`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetGcpOk() (*int32, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetGcp(v int32)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

### GetGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetGithub() int32`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetGithubOk() (*int32, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetGithub(v int32)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetLogcollectorBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownLogcollectorBreakdown`

GetLogcollectorBreakdown returns the LogcollectorBreakdown field if non-nil, zero value otherwise.

### GetLogcollectorBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetLogcollectorBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownLogcollectorBreakdown, bool)`

GetLogcollectorBreakdownOk returns a tuple with the LogcollectorBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetLogcollectorBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdownLogcollectorBreakdown)`

SetLogcollectorBreakdown sets LogcollectorBreakdown field to given value.

### HasLogcollectorBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasLogcollectorBreakdown() bool`

HasLogcollectorBreakdown returns a boolean if a field has been set.

### GetOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOffice365() int32`

GetOffice365 returns the Office365 field if non-nil, zero value otherwise.

### GetOffice365Ok

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOffice365Ok() (*int32, bool)`

GetOffice365Ok returns a tuple with the Office365 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetOffice365(v int32)`

SetOffice365 sets Office365 field to given value.

### HasOffice365

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasOffice365() bool`

HasOffice365 returns a boolean if a field has been set.

### GetMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetMsGraph() int32`

GetMsGraph returns the MsGraph field if non-nil, zero value otherwise.

### GetMsGraphOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetMsGraphOk() (*int32, bool)`

GetMsGraphOk returns a tuple with the MsGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetMsGraph(v int32)`

SetMsGraph sets MsGraph field to given value.

### HasMsGraph

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasMsGraph() bool`

HasMsGraph returns a boolean if a field has been set.

### GetOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOscap() int32`

GetOscap returns the Oscap field if non-nil, zero value otherwise.

### GetOscapOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOscapOk() (*int32, bool)`

GetOscapOk returns a tuple with the Oscap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetOscap(v int32)`

SetOscap sets Oscap field to given value.

### HasOscap

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasOscap() bool`

HasOscap returns a boolean if a field has been set.

### GetOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOsquery() int32`

GetOsquery returns the Osquery field if non-nil, zero value otherwise.

### GetOsqueryOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetOsqueryOk() (*int32, bool)`

GetOsqueryOk returns a tuple with the Osquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetOsquery(v int32)`

SetOsquery sets Osquery field to given value.

### HasOsquery

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasOsquery() bool`

HasOsquery returns a boolean if a field has been set.

### GetRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetRootcheck() int32`

GetRootcheck returns the Rootcheck field if non-nil, zero value otherwise.

### GetRootcheckOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetRootcheckOk() (*int32, bool)`

GetRootcheckOk returns a tuple with the Rootcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetRootcheck(v int32)`

SetRootcheck sets Rootcheck field to given value.

### HasRootcheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasRootcheck() bool`

HasRootcheck returns a boolean if a field has been set.

### GetSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetSca() int32`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetScaOk() (*int32, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetSca(v int32)`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetSyscheck() int32`

GetSyscheck returns the Syscheck field if non-nil, zero value otherwise.

### GetSyscheckOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetSyscheckOk() (*int32, bool)`

GetSyscheckOk returns a tuple with the Syscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetSyscheck(v int32)`

SetSyscheck sets Syscheck field to given value.

### HasSyscheck

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasSyscheck() bool`

HasSyscheck returns a boolean if a field has been set.

### GetSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetSyscollector() int32`

GetSyscollector returns the Syscollector field if non-nil, zero value otherwise.

### GetSyscollectorOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetSyscollectorOk() (*int32, bool)`

GetSyscollectorOk returns a tuple with the Syscollector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetSyscollector(v int32)`

SetSyscollector sets Syscollector field to given value.

### HasSyscollector

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasSyscollector() bool`

HasSyscollector returns a boolean if a field has been set.

### GetUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetUpgrade() int32`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) GetUpgradeOk() (*int32, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) SetUpgrade(v int32)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


