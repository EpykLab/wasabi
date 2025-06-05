# WazuhRemotedStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CtrlMsgCount** | Pointer to **float32** | Number of control messages received from all agents during the last five seconds | [optional] 
**DiscardedCount** | Pointer to **float32** | Number of discarded events received from agents during the last five seconds | [optional] 
**EvtCount** | Pointer to **float32** | Number of events sent to analysisd during the last five seconds | [optional] 
**SentBytes** | Pointer to **float32** | Number of sent bytes to the agents during the last five seconds | [optional] 
**QueueSize** | Pointer to **float32** | Usage of the queue to storage events from agents | [optional] 
**RecvBytes** | Pointer to **float32** | Number of received bytes from all agents during the last five seconds | [optional] 
**TcpSessions** | Pointer to **float32** | Number of TCP active sessions during the last five seconds | [optional] 
**TotalQueueSize** | Pointer to **float32** | Total queue size to store events from agents | [optional] 

## Methods

### NewWazuhRemotedStats

`func NewWazuhRemotedStats() *WazuhRemotedStats`

NewWazuhRemotedStats instantiates a new WazuhRemotedStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsWithDefaults

`func NewWazuhRemotedStatsWithDefaults() *WazuhRemotedStats`

NewWazuhRemotedStatsWithDefaults instantiates a new WazuhRemotedStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtrlMsgCount

`func (o *WazuhRemotedStats) GetCtrlMsgCount() float32`

GetCtrlMsgCount returns the CtrlMsgCount field if non-nil, zero value otherwise.

### GetCtrlMsgCountOk

`func (o *WazuhRemotedStats) GetCtrlMsgCountOk() (*float32, bool)`

GetCtrlMsgCountOk returns a tuple with the CtrlMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtrlMsgCount

`func (o *WazuhRemotedStats) SetCtrlMsgCount(v float32)`

SetCtrlMsgCount sets CtrlMsgCount field to given value.

### HasCtrlMsgCount

`func (o *WazuhRemotedStats) HasCtrlMsgCount() bool`

HasCtrlMsgCount returns a boolean if a field has been set.

### GetDiscardedCount

`func (o *WazuhRemotedStats) GetDiscardedCount() float32`

GetDiscardedCount returns the DiscardedCount field if non-nil, zero value otherwise.

### GetDiscardedCountOk

`func (o *WazuhRemotedStats) GetDiscardedCountOk() (*float32, bool)`

GetDiscardedCountOk returns a tuple with the DiscardedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedCount

`func (o *WazuhRemotedStats) SetDiscardedCount(v float32)`

SetDiscardedCount sets DiscardedCount field to given value.

### HasDiscardedCount

`func (o *WazuhRemotedStats) HasDiscardedCount() bool`

HasDiscardedCount returns a boolean if a field has been set.

### GetEvtCount

`func (o *WazuhRemotedStats) GetEvtCount() float32`

GetEvtCount returns the EvtCount field if non-nil, zero value otherwise.

### GetEvtCountOk

`func (o *WazuhRemotedStats) GetEvtCountOk() (*float32, bool)`

GetEvtCountOk returns a tuple with the EvtCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtCount

`func (o *WazuhRemotedStats) SetEvtCount(v float32)`

SetEvtCount sets EvtCount field to given value.

### HasEvtCount

`func (o *WazuhRemotedStats) HasEvtCount() bool`

HasEvtCount returns a boolean if a field has been set.

### GetSentBytes

`func (o *WazuhRemotedStats) GetSentBytes() float32`

GetSentBytes returns the SentBytes field if non-nil, zero value otherwise.

### GetSentBytesOk

`func (o *WazuhRemotedStats) GetSentBytesOk() (*float32, bool)`

GetSentBytesOk returns a tuple with the SentBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentBytes

`func (o *WazuhRemotedStats) SetSentBytes(v float32)`

SetSentBytes sets SentBytes field to given value.

### HasSentBytes

`func (o *WazuhRemotedStats) HasSentBytes() bool`

HasSentBytes returns a boolean if a field has been set.

### GetQueueSize

`func (o *WazuhRemotedStats) GetQueueSize() float32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *WazuhRemotedStats) GetQueueSizeOk() (*float32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *WazuhRemotedStats) SetQueueSize(v float32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *WazuhRemotedStats) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.

### GetRecvBytes

`func (o *WazuhRemotedStats) GetRecvBytes() float32`

GetRecvBytes returns the RecvBytes field if non-nil, zero value otherwise.

### GetRecvBytesOk

`func (o *WazuhRemotedStats) GetRecvBytesOk() (*float32, bool)`

GetRecvBytesOk returns a tuple with the RecvBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecvBytes

`func (o *WazuhRemotedStats) SetRecvBytes(v float32)`

SetRecvBytes sets RecvBytes field to given value.

### HasRecvBytes

`func (o *WazuhRemotedStats) HasRecvBytes() bool`

HasRecvBytes returns a boolean if a field has been set.

### GetTcpSessions

`func (o *WazuhRemotedStats) GetTcpSessions() float32`

GetTcpSessions returns the TcpSessions field if non-nil, zero value otherwise.

### GetTcpSessionsOk

`func (o *WazuhRemotedStats) GetTcpSessionsOk() (*float32, bool)`

GetTcpSessionsOk returns a tuple with the TcpSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSessions

`func (o *WazuhRemotedStats) SetTcpSessions(v float32)`

SetTcpSessions sets TcpSessions field to given value.

### HasTcpSessions

`func (o *WazuhRemotedStats) HasTcpSessions() bool`

HasTcpSessions returns a boolean if a field has been set.

### GetTotalQueueSize

`func (o *WazuhRemotedStats) GetTotalQueueSize() float32`

GetTotalQueueSize returns the TotalQueueSize field if non-nil, zero value otherwise.

### GetTotalQueueSizeOk

`func (o *WazuhRemotedStats) GetTotalQueueSizeOk() (*float32, bool)`

GetTotalQueueSizeOk returns a tuple with the TotalQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQueueSize

`func (o *WazuhRemotedStats) SetTotalQueueSize(v float32)`

SetTotalQueueSize sets TotalQueueSize field to given value.

### HasTotalQueueSize

`func (o *WazuhRemotedStats) HasTotalQueueSize() bool`

HasTotalQueueSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


