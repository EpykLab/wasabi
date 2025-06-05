# WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | Events coming from agentd (this agent) | [optional] 
**Dbsync** | Pointer to **int32** | Synchronization events (this agent) | [optional] 
**IntegrationsBreakdown** | Pointer to [**WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown**](WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown.md) |  | [optional] 
**ModulesBreakdown** | Pointer to [**WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown**](WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown.md) |  | [optional] 
**Monitor** | Pointer to **int32** | Events coming from monitord (this agent) | [optional] 
**Remote** | Pointer to **int32** | Events coming from remoted (this agent) | [optional] 

## Methods

### NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown

`func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown`

NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults

`func NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults() *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown`

NewWazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownWithDefaults instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetDbsync

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetDbsync() int32`

GetDbsync returns the Dbsync field if non-nil, zero value otherwise.

### GetDbsyncOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetDbsyncOk() (*int32, bool)`

GetDbsyncOk returns a tuple with the Dbsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbsync

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetDbsync(v int32)`

SetDbsync sets Dbsync field to given value.

### HasDbsync

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasDbsync() bool`

HasDbsync returns a boolean if a field has been set.

### GetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetIntegrationsBreakdown() WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown`

GetIntegrationsBreakdown returns the IntegrationsBreakdown field if non-nil, zero value otherwise.

### GetIntegrationsBreakdownOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetIntegrationsBreakdownOk() (*WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown, bool)`

GetIntegrationsBreakdownOk returns a tuple with the IntegrationsBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetIntegrationsBreakdown(v WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownIntegrationsBreakdown)`

SetIntegrationsBreakdown sets IntegrationsBreakdown field to given value.

### HasIntegrationsBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasIntegrationsBreakdown() bool`

HasIntegrationsBreakdown returns a boolean if a field has been set.

### GetModulesBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetModulesBreakdown() WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown`

GetModulesBreakdown returns the ModulesBreakdown field if non-nil, zero value otherwise.

### GetModulesBreakdownOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetModulesBreakdownOk() (*WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown, bool)`

GetModulesBreakdownOk returns a tuple with the ModulesBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetModulesBreakdown(v WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdownModulesBreakdown)`

SetModulesBreakdown sets ModulesBreakdown field to given value.

### HasModulesBreakdown

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasModulesBreakdown() bool`

HasModulesBreakdown returns a boolean if a field has been set.

### GetMonitor

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetMonitor() int32`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetMonitorOk() (*int32, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetMonitor(v int32)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetRemote

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetRemote() int32`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) GetRemoteOk() (*int32, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) SetRemote(v int32)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInnerMetricsEventsReceivedBreakdownDecodedBreakdown) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


