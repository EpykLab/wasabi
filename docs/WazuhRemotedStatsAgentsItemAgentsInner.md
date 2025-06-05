# WazuhRemotedStatsAgentsItemAgentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **time.Time** | When the count of the metrics started | [optional] 
**Id** | Pointer to **int32** | Agent ID | [optional] 
**Metrics** | Pointer to [**WazuhRemotedStatsAgentsItemAgentsInnerMetrics**](WazuhRemotedStatsAgentsItemAgentsInnerMetrics.md) |  | [optional] 

## Methods

### NewWazuhRemotedStatsAgentsItemAgentsInner

`func NewWazuhRemotedStatsAgentsItemAgentsInner() *WazuhRemotedStatsAgentsItemAgentsInner`

NewWazuhRemotedStatsAgentsItemAgentsInner instantiates a new WazuhRemotedStatsAgentsItemAgentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsAgentsItemAgentsInnerWithDefaults

`func NewWazuhRemotedStatsAgentsItemAgentsInnerWithDefaults() *WazuhRemotedStatsAgentsItemAgentsInner`

NewWazuhRemotedStatsAgentsItemAgentsInnerWithDefaults instantiates a new WazuhRemotedStatsAgentsItemAgentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetId

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetrics

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetMetrics() WazuhRemotedStatsAgentsItemAgentsInnerMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) GetMetricsOk() (*WazuhRemotedStatsAgentsItemAgentsInnerMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) SetMetrics(v WazuhRemotedStatsAgentsItemAgentsInnerMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *WazuhRemotedStatsAgentsItemAgentsInner) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


