# WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | Events discarded from agentd because the queue was full | [optional] 
**Agentless** | Pointer to **int32** | Events discarded from agentlessd because the queue was full | [optional] 
**Dbsync** | Pointer to **int32** | Synchronization events discarded because the queue was full | [optional] 
**IntegrationsBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown.md) |  | [optional] 
**ModulesBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown.md) |  | [optional] 
**Monitor** | Pointer to **int32** | Events discarded from monitord because the queue was full | [optional] 
**Remote** | Pointer to **int32** | Events discarded from remoted because the queue was full | [optional] 
**Syslog** | Pointer to **int32** | Events discarded from syslog remoted because the queue was full | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownWithDefaults() *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown`

NewWazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentless() int32`

GetAgentless returns the Agentless field if non-nil, zero value otherwise.

### GetAgentlessOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetAgentlessOk() (*int32, bool)`

GetAgentlessOk returns a tuple with the Agentless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetAgentless(v int32)`

SetAgentless sets Agentless field to given value.

### HasAgentless

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasAgentless() bool`

HasAgentless returns a boolean if a field has been set.

### GetDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetDbsync() int32`

GetDbsync returns the Dbsync field if non-nil, zero value otherwise.

### GetDbsyncOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetDbsyncOk() (*int32, bool)`

GetDbsyncOk returns a tuple with the Dbsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetDbsync(v int32)`

SetDbsync sets Dbsync field to given value.

### HasDbsync

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasDbsync() bool`

HasDbsync returns a boolean if a field has been set.

### GetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetIntegrationsBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown`

GetIntegrationsBreakdown returns the IntegrationsBreakdown field if non-nil, zero value otherwise.

### GetIntegrationsBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetIntegrationsBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown, bool)`

GetIntegrationsBreakdownOk returns a tuple with the IntegrationsBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetIntegrationsBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownIntegrationsBreakdown)`

SetIntegrationsBreakdown sets IntegrationsBreakdown field to given value.

### HasIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasIntegrationsBreakdown() bool`

HasIntegrationsBreakdown returns a boolean if a field has been set.

### GetModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetModulesBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown`

GetModulesBreakdown returns the ModulesBreakdown field if non-nil, zero value otherwise.

### GetModulesBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetModulesBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown, bool)`

GetModulesBreakdownOk returns a tuple with the ModulesBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetModulesBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdownModulesBreakdown)`

SetModulesBreakdown sets ModulesBreakdown field to given value.

### HasModulesBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasModulesBreakdown() bool`

HasModulesBreakdown returns a boolean if a field has been set.

### GetMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetMonitor() int32`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetMonitorOk() (*int32, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetMonitor(v int32)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetRemote() int32`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetRemoteOk() (*int32, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetRemote(v int32)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetSyslog() int32`

GetSyslog returns the Syslog field if non-nil, zero value otherwise.

### GetSyslogOk

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) GetSyslogOk() (*int32, bool)`

GetSyslogOk returns a tuple with the Syslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) SetSyslog(v int32)`

SetSyslog sets Syslog field to given value.

### HasSyslog

`func (o *WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdownDroppedBreakdown) HasSyslog() bool`

HasSyslog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


