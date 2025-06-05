# WazuhAnalysisdStatsItemMetricsEps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableCredits** | Pointer to **int32** | Available credits to process events in the current timeframe | [optional] 
**AvailableCreditsPrev** | Pointer to **int32** | Available credits to process events in the previous timeframe | [optional] 
**EventsDropped** | Pointer to **int32** | Events discarded because the EPS limit was reached and queues were full | [optional] 
**EventsDroppedNotEps** | Pointer to **int32** | Events discarded due to causes unrelated to EPS limit | [optional] 
**SecondsOverLimit** | Pointer to **int32** | Time in seconds the EPS limit was exceeded | [optional] 

## Methods

### NewWazuhAnalysisdStatsItemMetricsEps

`func NewWazuhAnalysisdStatsItemMetricsEps() *WazuhAnalysisdStatsItemMetricsEps`

NewWazuhAnalysisdStatsItemMetricsEps instantiates a new WazuhAnalysisdStatsItemMetricsEps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsItemMetricsEpsWithDefaults

`func NewWazuhAnalysisdStatsItemMetricsEpsWithDefaults() *WazuhAnalysisdStatsItemMetricsEps`

NewWazuhAnalysisdStatsItemMetricsEpsWithDefaults instantiates a new WazuhAnalysisdStatsItemMetricsEps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableCredits

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetAvailableCredits() int32`

GetAvailableCredits returns the AvailableCredits field if non-nil, zero value otherwise.

### GetAvailableCreditsOk

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetAvailableCreditsOk() (*int32, bool)`

GetAvailableCreditsOk returns a tuple with the AvailableCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredits

`func (o *WazuhAnalysisdStatsItemMetricsEps) SetAvailableCredits(v int32)`

SetAvailableCredits sets AvailableCredits field to given value.

### HasAvailableCredits

`func (o *WazuhAnalysisdStatsItemMetricsEps) HasAvailableCredits() bool`

HasAvailableCredits returns a boolean if a field has been set.

### GetAvailableCreditsPrev

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetAvailableCreditsPrev() int32`

GetAvailableCreditsPrev returns the AvailableCreditsPrev field if non-nil, zero value otherwise.

### GetAvailableCreditsPrevOk

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetAvailableCreditsPrevOk() (*int32, bool)`

GetAvailableCreditsPrevOk returns a tuple with the AvailableCreditsPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCreditsPrev

`func (o *WazuhAnalysisdStatsItemMetricsEps) SetAvailableCreditsPrev(v int32)`

SetAvailableCreditsPrev sets AvailableCreditsPrev field to given value.

### HasAvailableCreditsPrev

`func (o *WazuhAnalysisdStatsItemMetricsEps) HasAvailableCreditsPrev() bool`

HasAvailableCreditsPrev returns a boolean if a field has been set.

### GetEventsDropped

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetEventsDropped() int32`

GetEventsDropped returns the EventsDropped field if non-nil, zero value otherwise.

### GetEventsDroppedOk

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetEventsDroppedOk() (*int32, bool)`

GetEventsDroppedOk returns a tuple with the EventsDropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDropped

`func (o *WazuhAnalysisdStatsItemMetricsEps) SetEventsDropped(v int32)`

SetEventsDropped sets EventsDropped field to given value.

### HasEventsDropped

`func (o *WazuhAnalysisdStatsItemMetricsEps) HasEventsDropped() bool`

HasEventsDropped returns a boolean if a field has been set.

### GetEventsDroppedNotEps

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetEventsDroppedNotEps() int32`

GetEventsDroppedNotEps returns the EventsDroppedNotEps field if non-nil, zero value otherwise.

### GetEventsDroppedNotEpsOk

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetEventsDroppedNotEpsOk() (*int32, bool)`

GetEventsDroppedNotEpsOk returns a tuple with the EventsDroppedNotEps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDroppedNotEps

`func (o *WazuhAnalysisdStatsItemMetricsEps) SetEventsDroppedNotEps(v int32)`

SetEventsDroppedNotEps sets EventsDroppedNotEps field to given value.

### HasEventsDroppedNotEps

`func (o *WazuhAnalysisdStatsItemMetricsEps) HasEventsDroppedNotEps() bool`

HasEventsDroppedNotEps returns a boolean if a field has been set.

### GetSecondsOverLimit

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetSecondsOverLimit() int32`

GetSecondsOverLimit returns the SecondsOverLimit field if non-nil, zero value otherwise.

### GetSecondsOverLimitOk

`func (o *WazuhAnalysisdStatsItemMetricsEps) GetSecondsOverLimitOk() (*int32, bool)`

GetSecondsOverLimitOk returns a tuple with the SecondsOverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsOverLimit

`func (o *WazuhAnalysisdStatsItemMetricsEps) SetSecondsOverLimit(v int32)`

SetSecondsOverLimit sets SecondsOverLimit field to given value.

### HasSecondsOverLimit

`func (o *WazuhAnalysisdStatsItemMetricsEps) HasSecondsOverLimit() bool`

HasSecondsOverLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


