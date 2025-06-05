# WazuhRemotedStatsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **time.Time** | When the count of the metrics started | [optional] 
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Metrics** | Pointer to [**WazuhRemotedStatsItemMetrics**](WazuhRemotedStatsItemMetrics.md) |  | [optional] 

## Methods

### NewWazuhRemotedStatsItem

`func NewWazuhRemotedStatsItem() *WazuhRemotedStatsItem`

NewWazuhRemotedStatsItem instantiates a new WazuhRemotedStatsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsItemWithDefaults

`func NewWazuhRemotedStatsItemWithDefaults() *WazuhRemotedStatsItem`

NewWazuhRemotedStatsItemWithDefaults instantiates a new WazuhRemotedStatsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *WazuhRemotedStatsItem) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *WazuhRemotedStatsItem) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *WazuhRemotedStatsItem) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *WazuhRemotedStatsItem) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTimestamp

`func (o *WazuhRemotedStatsItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WazuhRemotedStatsItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WazuhRemotedStatsItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WazuhRemotedStatsItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *WazuhRemotedStatsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WazuhRemotedStatsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WazuhRemotedStatsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WazuhRemotedStatsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetrics

`func (o *WazuhRemotedStatsItem) GetMetrics() WazuhRemotedStatsItemMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *WazuhRemotedStatsItem) GetMetricsOk() (*WazuhRemotedStatsItemMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *WazuhRemotedStatsItem) SetMetrics(v WazuhRemotedStatsItemMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *WazuhRemotedStatsItem) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


