# AllItemsResponseWazuhDaemonStatsAllOfAffectedItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | Pointer to **time.Time** | When the count of the metrics started | [optional] 
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Metrics** | Pointer to [**WazuhDBStatsItemMetrics**](WazuhDBStatsItemMetrics.md) |  | [optional] 

## Methods

### NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItems

`func NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItems() *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems`

NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItems instantiates a new AllItemsResponseWazuhDaemonStatsAllOfAffectedItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItemsWithDefaults

`func NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItemsWithDefaults() *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems`

NewAllItemsResponseWazuhDaemonStatsAllOfAffectedItemsWithDefaults instantiates a new AllItemsResponseWazuhDaemonStatsAllOfAffectedItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUptime

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetrics

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetMetrics() WazuhDBStatsItemMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) GetMetricsOk() (*WazuhDBStatsItemMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) SetMetrics(v WazuhDBStatsItemMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AllItemsResponseWazuhDaemonStatsAllOfAffectedItems) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


