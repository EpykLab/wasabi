# WazuhDBStatsItemMetricsQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to **int32** | Total number of queries through WazuhDB socket | [optional] 
**ReceivedBreakdown** | Pointer to [**WazuhDBStatsItemMetricsQueriesReceivedBreakdown**](WazuhDBStatsItemMetricsQueriesReceivedBreakdown.md) |  | [optional] 

## Methods

### NewWazuhDBStatsItemMetricsQueries

`func NewWazuhDBStatsItemMetricsQueries() *WazuhDBStatsItemMetricsQueries`

NewWazuhDBStatsItemMetricsQueries instantiates a new WazuhDBStatsItemMetricsQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhDBStatsItemMetricsQueriesWithDefaults

`func NewWazuhDBStatsItemMetricsQueriesWithDefaults() *WazuhDBStatsItemMetricsQueries`

NewWazuhDBStatsItemMetricsQueriesWithDefaults instantiates a new WazuhDBStatsItemMetricsQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *WazuhDBStatsItemMetricsQueries) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *WazuhDBStatsItemMetricsQueries) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *WazuhDBStatsItemMetricsQueries) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *WazuhDBStatsItemMetricsQueries) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetReceivedBreakdown

`func (o *WazuhDBStatsItemMetricsQueries) GetReceivedBreakdown() WazuhDBStatsItemMetricsQueriesReceivedBreakdown`

GetReceivedBreakdown returns the ReceivedBreakdown field if non-nil, zero value otherwise.

### GetReceivedBreakdownOk

`func (o *WazuhDBStatsItemMetricsQueries) GetReceivedBreakdownOk() (*WazuhDBStatsItemMetricsQueriesReceivedBreakdown, bool)`

GetReceivedBreakdownOk returns a tuple with the ReceivedBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBreakdown

`func (o *WazuhDBStatsItemMetricsQueries) SetReceivedBreakdown(v WazuhDBStatsItemMetricsQueriesReceivedBreakdown)`

SetReceivedBreakdown sets ReceivedBreakdown field to given value.

### HasReceivedBreakdown

`func (o *WazuhDBStatsItemMetricsQueries) HasReceivedBreakdown() bool`

HasReceivedBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


