# WazuhDBStatsItemMetricsTimeExecutionBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **int32** | Time taken by all agent queries (milliseconds) | [optional] 
**AgentBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdownAgentBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdownAgentBreakdown.md) |  | [optional] 
**Global** | Pointer to **int32** | Time taken by all global queries (milliseconds) | [optional] 
**GlobalBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdown.md) |  | [optional] 
**Mitre** | Pointer to **int32** | Time taken by all mitre queries (milliseconds) | [optional] 
**MitreBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdownMitreBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdownMitreBreakdown.md) |  | [optional] 
**Task** | Pointer to **int32** | Time taken by all task queries (milliseconds) | [optional] 
**TaskBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdownTaskBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdownTaskBreakdown.md) |  | [optional] 
**Wazuhdb** | Pointer to **int32** | Time taken by all wazuhdb queries (milliseconds) | [optional] 
**WazuhdbBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdownWazuhdbBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdownWazuhdbBreakdown.md) |  | [optional] 

## Methods

### NewWazuhDBStatsItemMetricsTimeExecutionBreakdown

`func NewWazuhDBStatsItemMetricsTimeExecutionBreakdown() *WazuhDBStatsItemMetricsTimeExecutionBreakdown`

NewWazuhDBStatsItemMetricsTimeExecutionBreakdown instantiates a new WazuhDBStatsItemMetricsTimeExecutionBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhDBStatsItemMetricsTimeExecutionBreakdownWithDefaults

`func NewWazuhDBStatsItemMetricsTimeExecutionBreakdownWithDefaults() *WazuhDBStatsItemMetricsTimeExecutionBreakdown`

NewWazuhDBStatsItemMetricsTimeExecutionBreakdownWithDefaults instantiates a new WazuhDBStatsItemMetricsTimeExecutionBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetAgent() int32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetAgentOk() (*int32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetAgent(v int32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAgentBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetAgentBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdownAgentBreakdown`

GetAgentBreakdown returns the AgentBreakdown field if non-nil, zero value otherwise.

### GetAgentBreakdownOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetAgentBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdownAgentBreakdown, bool)`

GetAgentBreakdownOk returns a tuple with the AgentBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetAgentBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdownAgentBreakdown)`

SetAgentBreakdown sets AgentBreakdown field to given value.

### HasAgentBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasAgentBreakdown() bool`

HasAgentBreakdown returns a boolean if a field has been set.

### GetGlobal

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetGlobal() int32`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetGlobalOk() (*int32, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetGlobal(v int32)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetGlobalBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdown`

GetGlobalBreakdown returns the GlobalBreakdown field if non-nil, zero value otherwise.

### GetGlobalBreakdownOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetGlobalBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdown, bool)`

GetGlobalBreakdownOk returns a tuple with the GlobalBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetGlobalBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdownGlobalBreakdown)`

SetGlobalBreakdown sets GlobalBreakdown field to given value.

### HasGlobalBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasGlobalBreakdown() bool`

HasGlobalBreakdown returns a boolean if a field has been set.

### GetMitre

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetMitre() int32`

GetMitre returns the Mitre field if non-nil, zero value otherwise.

### GetMitreOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetMitreOk() (*int32, bool)`

GetMitreOk returns a tuple with the Mitre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitre

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetMitre(v int32)`

SetMitre sets Mitre field to given value.

### HasMitre

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasMitre() bool`

HasMitre returns a boolean if a field has been set.

### GetMitreBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetMitreBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdownMitreBreakdown`

GetMitreBreakdown returns the MitreBreakdown field if non-nil, zero value otherwise.

### GetMitreBreakdownOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetMitreBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdownMitreBreakdown, bool)`

GetMitreBreakdownOk returns a tuple with the MitreBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetMitreBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdownMitreBreakdown)`

SetMitreBreakdown sets MitreBreakdown field to given value.

### HasMitreBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasMitreBreakdown() bool`

HasMitreBreakdown returns a boolean if a field has been set.

### GetTask

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetTask() int32`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetTaskOk() (*int32, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetTask(v int32)`

SetTask sets Task field to given value.

### HasTask

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTaskBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetTaskBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdownTaskBreakdown`

GetTaskBreakdown returns the TaskBreakdown field if non-nil, zero value otherwise.

### GetTaskBreakdownOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetTaskBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdownTaskBreakdown, bool)`

GetTaskBreakdownOk returns a tuple with the TaskBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetTaskBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdownTaskBreakdown)`

SetTaskBreakdown sets TaskBreakdown field to given value.

### HasTaskBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasTaskBreakdown() bool`

HasTaskBreakdown returns a boolean if a field has been set.

### GetWazuhdb

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetWazuhdb() int32`

GetWazuhdb returns the Wazuhdb field if non-nil, zero value otherwise.

### GetWazuhdbOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetWazuhdbOk() (*int32, bool)`

GetWazuhdbOk returns a tuple with the Wazuhdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhdb

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetWazuhdb(v int32)`

SetWazuhdb sets Wazuhdb field to given value.

### HasWazuhdb

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasWazuhdb() bool`

HasWazuhdb returns a boolean if a field has been set.

### GetWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetWazuhdbBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdownWazuhdbBreakdown`

GetWazuhdbBreakdown returns the WazuhdbBreakdown field if non-nil, zero value otherwise.

### GetWazuhdbBreakdownOk

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) GetWazuhdbBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdownWazuhdbBreakdown, bool)`

GetWazuhdbBreakdownOk returns a tuple with the WazuhdbBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) SetWazuhdbBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdownWazuhdbBreakdown)`

SetWazuhdbBreakdown sets WazuhdbBreakdown field to given value.

### HasWazuhdbBreakdown

`func (o *WazuhDBStatsItemMetricsTimeExecutionBreakdown) HasWazuhdbBreakdown() bool`

HasWazuhdbBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


