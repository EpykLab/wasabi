# WazuhAnalysisdStatsAgentsItemAgentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **time.Time** | When the count of the metrics started | [optional] 
**Id** | Pointer to **int32** | Agent ID | [optional] 
**Metrics** | Pointer to [**WazuhAnalysisdStatsAgentsItemAgentsInnerMetrics**](WazuhAnalysisdStatsAgentsItemAgentsInnerMetrics.md) |  | [optional] 

## Methods

### NewWazuhAnalysisdStatsAgentsItemAgentsInner

`func NewWazuhAnalysisdStatsAgentsItemAgentsInner() *WazuhAnalysisdStatsAgentsItemAgentsInner`

NewWazuhAnalysisdStatsAgentsItemAgentsInner instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsAgentsItemAgentsInnerWithDefaults

`func NewWazuhAnalysisdStatsAgentsItemAgentsInnerWithDefaults() *WazuhAnalysisdStatsAgentsItemAgentsInner`

NewWazuhAnalysisdStatsAgentsItemAgentsInnerWithDefaults instantiates a new WazuhAnalysisdStatsAgentsItemAgentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetId

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetrics

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetMetrics() WazuhAnalysisdStatsAgentsItemAgentsInnerMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) GetMetricsOk() (*WazuhAnalysisdStatsAgentsItemAgentsInnerMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) SetMetrics(v WazuhAnalysisdStatsAgentsItemAgentsInnerMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *WazuhAnalysisdStatsAgentsItemAgentsInner) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


