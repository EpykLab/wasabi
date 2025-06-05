# WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | Events coming from agentd | [optional] 
**Agentless** | Pointer to **int32** | Events coming from agentlessd | [optional] 
**Dbsync** | Pointer to **int32** | Synchronization events | [optional] 
**IntegrationsBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown.md) |  | [optional] 
**ModulesBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown.md) |  | [optional] 
**Monitor** | Pointer to **int32** | Events coming from monitord | [optional] 
**Remote** | Pointer to **int32** | Events coming from remoted | [optional] 
**Syslog** | Pointer to **int32** | Events coming from syslog remoted | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgentless() int32`

GetAgentless returns the Agentless field if non-nil, zero value otherwise.

### GetAgentlessOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgentlessOk() (*int32, bool)`

GetAgentlessOk returns a tuple with the Agentless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetAgentless(v int32)`

SetAgentless sets Agentless field to given value.

### HasAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasAgentless() bool`

HasAgentless returns a boolean if a field has been set.

### GetDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetDbsync() int32`

GetDbsync returns the Dbsync field if non-nil, zero value otherwise.

### GetDbsyncOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetDbsyncOk() (*int32, bool)`

GetDbsyncOk returns a tuple with the Dbsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetDbsync(v int32)`

SetDbsync sets Dbsync field to given value.

### HasDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasDbsync() bool`

HasDbsync returns a boolean if a field has been set.

### GetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetIntegrationsBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown`

GetIntegrationsBreakdown returns the IntegrationsBreakdown field if non-nil, zero value otherwise.

### GetIntegrationsBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetIntegrationsBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown, bool)`

GetIntegrationsBreakdownOk returns a tuple with the IntegrationsBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetIntegrationsBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown)`

SetIntegrationsBreakdown sets IntegrationsBreakdown field to given value.

### HasIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasIntegrationsBreakdown() bool`

HasIntegrationsBreakdown returns a boolean if a field has been set.

### GetModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetModulesBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

GetModulesBreakdown returns the ModulesBreakdown field if non-nil, zero value otherwise.

### GetModulesBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetModulesBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown, bool)`

GetModulesBreakdownOk returns a tuple with the ModulesBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetModulesBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown)`

SetModulesBreakdown sets ModulesBreakdown field to given value.

### HasModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasModulesBreakdown() bool`

HasModulesBreakdown returns a boolean if a field has been set.

### GetMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetMonitor() int32`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetMonitorOk() (*int32, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetMonitor(v int32)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetRemote() int32`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetRemoteOk() (*int32, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetRemote(v int32)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetSyslog() int32`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) GetSyslogOk() (*int32, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) SetSyslog(v int32)`

SetSyslog sets Syslog field to given value.

### HasSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDecodedBreakdown) HasSyslog() bool`

HasSyslog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


