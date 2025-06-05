# WazuhAnalysisdStatsItemMetricsEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processed** | Pointer to **int32** | Total processed events (analyzed by rules) | [optional] 
**Received** | Pointer to **int32** | Total received events from agents and local modules | [optional] 
**ReceivedBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown**](WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown.md) |  | [optional] 
**WrittenBreakdown** | Pointer to [**WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown**](WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown.md) |  | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEvents

`func NewWazuhAnalysisdStatsItemMetricsEvents() *WazuhAnalysisdStatsItemMetricsEvents`

NewWazuhAnalysisdStatsItemMetricsEvents instantiates a new WazuhAnalysisdStatsItemMetricsEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEventsWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEventsWithDefaults() *WazuhAnalysisdStatsItemMetricsEvents`

NewWazuhAnalysisdStatsItemMetricsEventsWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessed

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetProcessed() int32`

GetProcessed returns the Processed field if non-nil, zero value otherwise.

### GetProcessedOk

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetProcessedOk() (*int32, bool)`

GetProcessedOk returns a tuple with the Processed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessed

`func (o *WazuhAnalysisdStatsItemMetricsEvents) SetProcessed(v int32)`

SetProcessed sets Processed field to given value.

### HasProcessed

`func (o *WazuhAnalysisdStatsItemMetricsEvents) HasProcessed() bool`

HasProcessed returns a boolean if a field has been set.

### GetReceived

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *WazuhAnalysisdStatsItemMetricsEvents) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *WazuhAnalysisdStatsItemMetricsEvents) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetReceivedBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedBreakdown() WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown`

GetReceivedBreakdown returns the ReceivedBreakdown field if non-nil, zero value otherwise.

### GetReceivedBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetReceivedBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown, bool)`

GetReceivedBreakdownOk returns a tuple with the ReceivedBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) SetReceivedBreakdown(v WazuhAnalysisdStatsItemMetricsEventsReceivedBreakdown)`

SetReceivedBreakdown sets ReceivedBreakdown field to given value.

### HasReceivedBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) HasReceivedBreakdown() bool`

HasReceivedBreakdown returns a boolean if a field has been set.

### GetWrittenBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetWrittenBreakdown() WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown`

GetWrittenBreakdown returns the WrittenBreakdown field if non-nil, zero value otherwise.

### GetWrittenBreakdownOk

`func (o *WazuhAnalysisdStatsItemMetricsEvents) GetWrittenBreakdownOk() (*WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown, bool)`

GetWrittenBreakdownOk returns a tuple with the WrittenBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) SetWrittenBreakdown(v WazuhAnalysisdStatsItemMetricsEventsWrittenBreakdown)`

SetWrittenBreakdown sets WrittenBreakdown field to given value.

### HasWrittenBreakdown

`func (o *WazuhAnalysisdStatsItemMetricsEvents) HasWrittenBreakdown() bool`

HasWrittenBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


