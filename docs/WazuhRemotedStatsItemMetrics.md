# WazuhRemotedStatsItemMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to [**WazuhRemotedStatsItemMetricsBytes**](WazuhRemotedStatsItemMetricsBytes.md) |  | [optional] 
**KeysReloadCount** | Pointer to **int32** | Number of times keys were reloaded into memory | [optional] 
**Messages** | Pointer to [**WazuhRemotedStatsItemMetricsMessages**](WazuhRemotedStatsItemMetricsMessages.md) |  | [optional] 
**Queues** | Pointer to [**WazuhRemotedStatsItemMetricsQueues**](WazuhRemotedStatsItemMetricsQueues.md) |  | [optional] 
**TcpSessions** | Pointer to **int32** | Current active TCP sessions (agents) | [optional] 

## Methods

### NewWazuhRemotedStatsItemMetrics

`func NewWazuhRemotedStatsItemMetrics() *WazuhRemotedStatsItemMetrics`

NewWazuhRemotedStatsItemMetrics instantiates a new WazuhRemotedStatsItemMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsItemMetricsWithDefaults

`func NewWazuhRemotedStatsItemMetricsWithDefaults() *WazuhRemotedStatsItemMetrics`

NewWazuhRemotedStatsItemMetricsWithDefaults instantiates a new WazuhRemotedStatsItemMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *WazuhRemotedStatsItemMetrics) GetBytes() WazuhRemotedStatsItemMetricsBytes`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *WazuhRemotedStatsItemMetrics) GetBytesOk() (*WazuhRemotedStatsItemMetricsBytes, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *WazuhRemotedStatsItemMetrics) SetBytes(v WazuhRemotedStatsItemMetricsBytes)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *WazuhRemotedStatsItemMetrics) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetKeysReloadCount

`func (o *WazuhRemotedStatsItemMetrics) GetKeysReloadCount() int32`

GetKeysReloadCount returns the KeysReloadCount field if non-nil, zero value otherwise.

### GetKeysReloadCountOk

`func (o *WazuhRemotedStatsItemMetrics) GetKeysReloadCountOk() (*int32, bool)`

GetKeysReloadCountOk returns a tuple with the KeysReloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysReloadCount

`func (o *WazuhRemotedStatsItemMetrics) SetKeysReloadCount(v int32)`

SetKeysReloadCount sets KeysReloadCount field to given value.

### HasKeysReloadCount

`func (o *WazuhRemotedStatsItemMetrics) HasKeysReloadCount() bool`

HasKeysReloadCount returns a boolean if a field has been set.

### GetMessages

`func (o *WazuhRemotedStatsItemMetrics) GetMessages() WazuhRemotedStatsItemMetricsMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *WazuhRemotedStatsItemMetrics) GetMessagesOk() (*WazuhRemotedStatsItemMetricsMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *WazuhRemotedStatsItemMetrics) SetMessages(v WazuhRemotedStatsItemMetricsMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *WazuhRemotedStatsItemMetrics) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetQueues

`func (o *WazuhRemotedStatsItemMetrics) GetQueues() WazuhRemotedStatsItemMetricsQueues`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *WazuhRemotedStatsItemMetrics) GetQueuesOk() (*WazuhRemotedStatsItemMetricsQueues, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *WazuhRemotedStatsItemMetrics) SetQueues(v WazuhRemotedStatsItemMetricsQueues)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *WazuhRemotedStatsItemMetrics) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetTcpSessions

`func (o *WazuhRemotedStatsItemMetrics) GetTcpSessions() int32`

GetTcpSessions returns the TcpSessions field if non-nil, zero value otherwise.

### GetTcpSessionsOk

`func (o *WazuhRemotedStatsItemMetrics) GetTcpSessionsOk() (*int32, bool)`

GetTcpSessionsOk returns a tuple with the TcpSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSessions

`func (o *WazuhRemotedStatsItemMetrics) SetTcpSessions(v int32)`

SetTcpSessions sets TcpSessions field to given value.

### HasTcpSessions

`func (o *WazuhRemotedStatsItemMetrics) HasTcpSessions() bool`

HasTcpSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


