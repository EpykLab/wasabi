# WazuhAnalysisdStatsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **time.Time** | When the count of the metrics started | [optional] 
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Metrics** | Pointer to [**WazuhAnalysisdStatsItemMetrics**](WazuhAnalysisdStatsItemMetrics.md) |  | [optional] 

## Methods

### NewWazuhAnalysisdStatsItem

`func NewWazuhAnalysisdStatsItem() *WazuhAnalysisdStatsItem`

NewWazuhAnalysisdStatsItem instantiates a new WazuhAnalysisdStatsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemWithDefaults

`func NewWazuhAnalysisdStatsItemWithDefaults() *WazuhAnalysisdStatsItem`

NewWazuhAnalysisdStatsItemWithDefaults instantiates a new WazuhAnalysisdStatsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *WazuhAnalysisdStatsItem) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *WazuhAnalysisdStatsItem) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *WazuhAnalysisdStatsItem) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *WazuhAnalysisdStatsItem) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTimestamp

`func (o *WazuhAnalysisdStatsItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WazuhAnalysisdStatsItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WazuhAnalysisdStatsItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WazuhAnalysisdStatsItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *WazuhAnalysisdStatsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WazuhAnalysisdStatsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WazuhAnalysisdStatsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WazuhAnalysisdStatsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetrics

`func (o *WazuhAnalysisdStatsItem) GetMetrics() WazuhAnalysisdStatsItemMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WazuhAnalysisdStatsItem) GetMetricsOk() (*WazuhAnalysisdStatsItemMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WazuhAnalysisdStatsItem) SetMetrics(v WazuhAnalysisdStatsItemMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *WazuhAnalysisdStatsItem) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


