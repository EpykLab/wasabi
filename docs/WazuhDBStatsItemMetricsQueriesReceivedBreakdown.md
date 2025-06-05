# WazuhDBStatsItemMetricsQueriesReceivedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | Number of agent queries through WazuhDB socket | [optional] 
**AgentBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdownAgentBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdownAgentBreakdown.md) |  | [optional] 
**Global** | Pointer to **int32** | Number of global queries through WazuhDB socket | [optional] 
**GlobalBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdown.md) |  | [optional] 
**Mitre** | Pointer to **int32** | Number of mitre queries through WazuhDB socket | [optional] 
**MitreBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdownMitreBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdownMitreBreakdown.md) |  | [optional] 
**Task** | Pointer to **int32** | Number of task queries through WazuhDB socket | [optional] 
**TaskBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdownTaskBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdownTaskBreakdown.md) |  | [optional] 
**Wazuhdb** | Pointer to **int32** | Number of wazuhdb queries through WazuhDB socket | [optional] 
**WazuhdbBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdownWazuhdbBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdownWazuhdbBreakdown.md) |  | [optional] 

## Methods

### NewWazuhDBStatsItemMetricsQueriesReceivedBreakdown

`func NewWazuhDBStatsItemMetricsQueriesReceivedBreakdown() *WazuhDBStatsItemMetricsQueriesReceivedBreakdown`

NewWazuhDBStatsItemMetricsQueriesReceivedBreakdown instantiates a new WazuhDBStatsItemMetricsQueriesReceivedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhDBStatsItemMetricsQueriesReceivedBreakdownWithDefaults

`func NewWazuhDBStatsItemMetricsQueriesReceivedBreakdownWithDefaults() *WazuhDBStatsItemMetricsQueriesReceivedBreakdown`

NewWazuhDBStatsItemMetricsQueriesReceivedBreakdownWithDefaults instantiates a new WazuhDBStatsItemMetricsQueriesReceivedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAgentBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetAgentBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdownAgentBreakdown`

GetAgentBreakdown returns the AgentBreakdown field if non-nil, zero value otherwise.

### GetAgentBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetAgentBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownAgentBreakdown, bool)`

GetAgentBreakdownOk returns a tuple with the AgentBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetAgentBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownAgentBreakdown)`

SetAgentBreakdown sets AgentBreakdown field to given value.

### HasAgentBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasAgentBreakdown() bool`

HasAgentBreakdown returns a boolean if a field has been set.

### GetGlobal

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetGlobal() int32`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetGlobalOk() (*int32, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetGlobal(v int32)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetGlobalBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdown`

GetGlobalBreakdown returns the GlobalBreakdown field if non-nil, zero value otherwise.

### GetGlobalBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetGlobalBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdown, bool)`

GetGlobalBreakdownOk returns a tuple with the GlobalBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetGlobalBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownGlobalBreakdown)`

SetGlobalBreakdown sets GlobalBreakdown field to given value.

### HasGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasGlobalBreakdown() bool`

HasGlobalBreakdown returns a boolean if a field has been set.

### GetMitre

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetMitre() int32`

GetMitre returns the Mitre field if non-nil, zero value otherwise.

### GetMitreOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetMitreOk() (*int32, bool)`

GetMitreOk returns a tuple with the Mitre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitre

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetMitre(v int32)`

SetMitre sets Mitre field to given value.

### HasMitre

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasMitre() bool`

HasMitre returns a boolean if a field has been set.

### GetMitreBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetMitreBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdownMitreBreakdown`

GetMitreBreakdown returns the MitreBreakdown field if non-nil, zero value otherwise.

### GetMitreBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetMitreBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownMitreBreakdown, bool)`

GetMitreBreakdownOk returns a tuple with the MitreBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetMitreBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownMitreBreakdown)`

SetMitreBreakdown sets MitreBreakdown field to given value.

### HasMitreBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasMitreBreakdown() bool`

HasMitreBreakdown returns a boolean if a field has been set.

### GetTask

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetTaskBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdownTaskBreakdown`

GetTaskBreakdown returns the TaskBreakdown field if non-nil, zero value otherwise.

### GetTaskBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetTaskBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownTaskBreakdown, bool)`

GetTaskBreakdownOk returns a tuple with the TaskBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetTaskBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownTaskBreakdown)`

SetTaskBreakdown sets TaskBreakdown field to given value.

### HasTaskBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasTaskBreakdown() bool`

HasTaskBreakdown returns a boolean if a field has been set.

### GetWazuhdb

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetWazuhdb() int32`

GetWazuhdb returns the Wazuhdb field if non-nil, zero value otherwise.

### GetWazuhdbOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetWazuhdbOk() (*int32, bool)`

GetWazuhdbOk returns a tuple with the Wazuhdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhdb

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetWazuhdb(v int32)`

SetWazuhdb sets Wazuhdb field to given value.

### HasWazuhdb

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasWazuhdb() bool`

HasWazuhdb returns a boolean if a field has been set.

### GetWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetWazuhdbBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdownWazuhdbBreakdown`

GetWazuhdbBreakdown returns the WazuhdbBreakdown field if non-nil, zero value otherwise.

### GetWazuhdbBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) GetWazuhdbBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdownWazuhdbBreakdown, bool)`

GetWazuhdbBreakdownOk returns a tuple with the WazuhdbBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) SetWazuhdbBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdownWazuhdbBreakdown)`

SetWazuhdbBreakdown sets WazuhdbBreakdown field to given value.

### HasWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsQueriesReceivedBreakdown) HasWazuhdbBreakdown() bool`

HasWazuhdbBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


