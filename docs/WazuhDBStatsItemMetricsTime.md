# WazuhDBStatsItemMetricsTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Execution** | Pointer to **int32** | Total time taken by all the queries (milliseconds) | [optional] 
**ExecutionBreakdown** | Pointer to [**WazuhDBStatsItemMetricsTimeExecutionBreakdown**](WazuhDBStatsItemMetricsTimeExecutionBreakdown.md) |  | [optional] 

## Methods

### NewWazuhDBStatsItemMetricsTime

`func NewWazuhDBStatsItemMetricsTime() *WazuhDBStatsItemMetricsTime`

NewWazuhDBStatsItemMetricsTime instantiates a new WazuhDBStatsItemMetricsTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhDBStatsItemMetricsTimeWithDefaults

`func NewWazuhDBStatsItemMetricsTimeWithDefaults() *WazuhDBStatsItemMetricsTime`

NewWazuhDBStatsItemMetricsTimeWithDefaults instantiates a new WazuhDBStatsItemMetricsTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecution

`func (o *WazuhDBStatsItemMetricsTime) GetExecution() int32`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *WazuhDBStatsItemMetricsTime) GetExecutionOk() (*int32, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *WazuhDBStatsItemMetricsTime) SetExecution(v int32)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *WazuhDBStatsItemMetricsTime) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetExecutionBreakdown

`func (o *WazuhDBStatsItemMetricsTime) GetExecutionBreakdown() WazuhDBStatsItemMetricsTimeExecutionBreakdown`

GetExecutionBreakdown returns the ExecutionBreakdown field if non-nil, zero value otherwise.

### GetExecutionBreakdownOk

`func (o *WazuhDBStatsItemMetricsTime) GetExecutionBreakdownOk() (*WazuhDBStatsItemMetricsTimeExecutionBreakdown, bool)`

GetExecutionBreakdownOk returns a tuple with the ExecutionBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionBreakdown

`func (o *WazuhDBStatsItemMetricsTime) SetExecutionBreakdown(v WazuhDBStatsItemMetricsTimeExecutionBreakdown)`

SetExecutionBreakdown sets ExecutionBreakdown field to given value.

### HasExecutionBreakdown

`func (o *WazuhDBStatsItemMetricsTime) HasExecutionBreakdown() bool`

HasExecutionBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


