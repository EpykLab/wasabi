# SyscollectorPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inode** | Pointer to **int64** | Port inode | [optional] 
**Local** | Pointer to [**LocalPortInfo**](LocalPortInfo.md) |  | [optional] 
**Protocol** | Pointer to **string** | Protocol used in the communication | [optional] 
**Remote** | Pointer to [**RemotePortInfo**](RemotePortInfo.md) |  | [optional] 
**RxQueue** | Pointer to **int32** | Packets at the receiver queue | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**State** | Pointer to **string** | Communication status | [optional] 
**TxQueue** | Pointer to **int32** | Packets pending to be transmitted | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 
**Pid** | Pointer to **int32** | PID owner of the opened port | [optional] 
**Process** | Pointer to **string** | Name of the PID | [optional] 

## Methods

### NewSyscollectorPorts

`func NewSyscollectorPorts() *SyscollectorPorts`

NewSyscollectorPorts instantiates a new SyscollectorPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorPortsWithDefaults

`func NewSyscollectorPortsWithDefaults() *SyscollectorPorts`

NewSyscollectorPortsWithDefaults instantiates a new SyscollectorPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInode

`func (o *SyscollectorPorts) GetInode() int64`

GetInode returns the Inode field if non-nil, zero value otherwise.

### GetInodeOk

`func (o *SyscollectorPorts) GetInodeOk() (*int64, bool)`

GetInodeOk returns a tuple with the Inode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInode

`func (o *SyscollectorPorts) SetInode(v int64)`

SetInode sets Inode field to given value.

### HasInode

`func (o *SyscollectorPorts) HasInode() bool`

HasInode returns a boolean if a field has been set.

### GetLocal

`func (o *SyscollectorPorts) GetLocal() LocalPortInfo`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *SyscollectorPorts) GetLocalOk() (*LocalPortInfo, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *SyscollectorPorts) SetLocal(v LocalPortInfo)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *SyscollectorPorts) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetProtocol

`func (o *SyscollectorPorts) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyscollectorPorts) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyscollectorPorts) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyscollectorPorts) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemote

`func (o *SyscollectorPorts) GetRemote() RemotePortInfo`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *SyscollectorPorts) GetRemoteOk() (*RemotePortInfo, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *SyscollectorPorts) SetRemote(v RemotePortInfo)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *SyscollectorPorts) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetRxQueue

`func (o *SyscollectorPorts) GetRxQueue() int32`

GetRxQueue returns the RxQueue field if non-nil, zero value otherwise.

### GetRxQueueOk

`func (o *SyscollectorPorts) GetRxQueueOk() (*int32, bool)`

GetRxQueueOk returns a tuple with the RxQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueue

`func (o *SyscollectorPorts) SetRxQueue(v int32)`

SetRxQueue sets RxQueue field to given value.

### HasRxQueue

`func (o *SyscollectorPorts) HasRxQueue() bool`

HasRxQueue returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorPorts) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorPorts) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorPorts) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorPorts) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetState

`func (o *SyscollectorPorts) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SyscollectorPorts) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SyscollectorPorts) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SyscollectorPorts) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTxQueue

`func (o *SyscollectorPorts) GetTxQueue() int32`

GetTxQueue returns the TxQueue field if non-nil, zero value otherwise.

### GetTxQueueOk

`func (o *SyscollectorPorts) GetTxQueueOk() (*int32, bool)`

GetTxQueueOk returns a tuple with the TxQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueue

`func (o *SyscollectorPorts) SetTxQueue(v int32)`

SetTxQueue sets TxQueue field to given value.

### HasTxQueue

`func (o *SyscollectorPorts) HasTxQueue() bool`

HasTxQueue returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorPorts) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorPorts) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorPorts) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorPorts) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetPid

`func (o *SyscollectorPorts) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *SyscollectorPorts) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *SyscollectorPorts) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *SyscollectorPorts) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetProcess

`func (o *SyscollectorPorts) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *SyscollectorPorts) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *SyscollectorPorts) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *SyscollectorPorts) HasProcess() bool`

HasProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


