# WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Control** | Pointer to **int32** | Control messages received from agents | [optional] 
**ControlBreakdown** | Pointer to [**WazuhRemotedStatsItemMetricsMessagesReceivedBreakdownControlBreakdown**](WazuhRemotedStatsItemMetricsMessagesReceivedBreakdownControlBreakdown.md) |  | [optional] 
**DequeuedAfter** | Pointer to **int32** | Messages dequeued after newer messages (counter &lt; current counter) | [optional] 
**Discarded** | Pointer to **int32** | Messages discarded because the received queue was full | [optional] 
**Event** | Pointer to **int32** | Event messages (syscheck, syscollector, logcollector, etc.) received from agents | [optional] 
**Ping** | Pointer to **int32** | Ping messages received | [optional] 
**Unknown** | Pointer to **int32** | Not recognized messages | [optional] 

## Methods

### NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdown

`func NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdown() *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown`

NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdown instantiates a new WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdownWithDefaults

`func NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdownWithDefaults() *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown`

NewWazuhRemotedStatsItemMetricsMessagesReceivedBreakdownWithDefaults instantiates a new WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControl

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetControl() int32`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetControlOk() (*int32, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetControl(v int32)`

SetControl sets Control field to given value.

### HasControl

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasControl() bool`

HasControl returns a boolean if a field has been set.

### GetControlBreakdown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetControlBreakdown() WazuhRemotedStatsItemMetricsMessagesReceivedBreakdownControlBreakdown`

GetControlBreakdown returns the ControlBreakdown field if non-nil, zero value otherwise.

### GetControlBreakdownOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetControlBreakdownOk() (*WazuhRemotedStatsItemMetricsMessagesReceivedBreakdownControlBreakdown, bool)`

GetControlBreakdownOk returns a tuple with the ControlBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlBreakdown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetControlBreakdown(v WazuhRemotedStatsItemMetricsMessagesReceivedBreakdownControlBreakdown)`

SetControlBreakdown sets ControlBreakdown field to given value.

### HasControlBreakdown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasControlBreakdown() bool`

HasControlBreakdown returns a boolean if a field has been set.

### GetDequeuedAfter

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetDequeuedAfter() int32`

GetDequeuedAfter returns the DequeuedAfter field if non-nil, zero value otherwise.

### GetDequeuedAfterOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetDequeuedAfterOk() (*int32, bool)`

GetDequeuedAfterOk returns a tuple with the DequeuedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDequeuedAfter

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetDequeuedAfter(v int32)`

SetDequeuedAfter sets DequeuedAfter field to given value.

### HasDequeuedAfter

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasDequeuedAfter() bool`

HasDequeuedAfter returns a boolean if a field has been set.

### GetDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetDiscarded() int32`

GetDiscarded returns the Discarded field if non-nil, zero value otherwise.

### GetDiscardedOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetDiscardedOk() (*int32, bool)`

GetDiscardedOk returns a tuple with the Discarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetDiscarded(v int32)`

SetDiscarded sets Discarded field to given value.

### HasDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasDiscarded() bool`

HasDiscarded returns a boolean if a field has been set.

### GetEvent

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetEvent() int32`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetEventOk() (*int32, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetEvent(v int32)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetPing

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetPing() int32`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetPingOk() (*int32, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetPing(v int32)`

SetPing sets Ping field to given value.

### HasPing

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasPing() bool`

HasPing returns a boolean if a field has been set.

### GetUnknown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *WazuhRemotedStatsItemMetricsMessagesReceivedBreakdown) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


