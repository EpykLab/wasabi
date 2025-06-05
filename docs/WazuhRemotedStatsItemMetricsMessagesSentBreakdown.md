# WazuhRemotedStatsItemMetricsMessagesSentBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ack** | Pointer to **int32** | ACK messages (response to keepalive, startup and shutdown) sent to agents | [optional] 
**Ar** | Pointer to **int32** | Active response messages sent to agents | [optional] 
**Discarded** | Pointer to **int32** | Messages discarded because the send queue was full | [optional] 
**Request** | Pointer to **int32** | Request messages (for example, WPK chunks) sent to agents | [optional] 
**Sca** | Pointer to **int32** | SCA messages sent to agents | [optional] 
**Shared** | Pointer to **int32** | Shared configuration messages (merged.mg) sent to agents | [optional] 

## Methods

### NewWazuhRemotedStatsItemMetricsMessagesSentBreakdown

`func NewWazuhRemotedStatsItemMetricsMessagesSentBreakdown() *WazuhRemotedStatsItemMetricsMessagesSentBreakdown`

NewWazuhRemotedStatsItemMetricsMessagesSentBreakdown instantiates a new WazuhRemotedStatsItemMetricsMessagesSentBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsItemMetricsMessagesSentBreakdownWithDefaults

`func NewWazuhRemotedStatsItemMetricsMessagesSentBreakdownWithDefaults() *WazuhRemotedStatsItemMetricsMessagesSentBreakdown`

NewWazuhRemotedStatsItemMetricsMessagesSentBreakdownWithDefaults instantiates a new WazuhRemotedStatsItemMetricsMessagesSentBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAck

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetAck() int32`

GetAck returns the Ack field if non-nil, zero value otherwise.

### GetAckOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetAckOk() (*int32, bool)`

GetAckOk returns a tuple with the Ack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAck

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetAck(v int32)`

SetAck sets Ack field to given value.

### HasAck

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasAck() bool`

HasAck returns a boolean if a field has been set.

### GetAr

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetAr() int32`

GetAr returns the Ar field if non-nil, zero value otherwise.

### GetArOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetArOk() (*int32, bool)`

GetArOk returns a tuple with the Ar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAr

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetAr(v int32)`

SetAr sets Ar field to given value.

### HasAr

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasAr() bool`

HasAr returns a boolean if a field has been set.

### GetDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetDiscarded() int32`

GetDiscarded returns the Discarded field if non-nil, zero value otherwise.

### GetDiscardedOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetDiscardedOk() (*int32, bool)`

GetDiscardedOk returns a tuple with the Discarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetDiscarded(v int32)`

SetDiscarded sets Discarded field to given value.

### HasDiscarded

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasDiscarded() bool`

HasDiscarded returns a boolean if a field has been set.

### GetRequest

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetRequest() int32`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetRequestOk() (*int32, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetRequest(v int32)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSca

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetSca() int32`

GetSca returns the Sca field if non-nil, zero value otherwise.

### GetScaOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetScaOk() (*int32, bool)`

GetScaOk returns a tuple with the Sca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSca

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetSca(v int32)`

SetSca sets Sca field to given value.

### HasSca

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasSca() bool`

HasSca returns a boolean if a field has been set.

### GetShared

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetShared() int32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) GetSharedOk() (*int32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) SetShared(v int32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *WazuhRemotedStatsItemMetricsMessagesSentBreakdown) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


