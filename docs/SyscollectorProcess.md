# SyscollectorProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Argvs** | Pointer to **string** | Process received arguments | [optional] 
**Cmd** | Pointer to **string** | Executed command | [optional] 
**Egroup** | Pointer to **string** | Effective group | [optional] 
**Euser** | Pointer to **string** | Effective user | [optional] 
**Fgroup** | Pointer to **string** | Filesystem group name | [optional] 
**Name** | Pointer to **string** | Process name | [optional] 
**Nice** | Pointer to **int32** | Nice value of the process | [optional] 
**Nlwp** | Pointer to **int32** | Number of light weight processes | [optional] 
**Pgrp** | Pointer to **int32** | Process group | [optional] 
**Pid** | Pointer to **string** | Process PID | [optional] 
**Ppid** | Pointer to **int32** | Process parent PID | [optional] 
**Priority** | Pointer to **int32** | Kernel scheduling priority | [optional] 
**Processor** | Pointer to **int32** | Processor number which is running the process | [optional] 
**Resident** | Pointer to **int32** | Process resident size in bytes | [optional] 
**Rgroup** | Pointer to **string** | Real group | [optional] 
**Ruser** | Pointer to **string** | Real user | [optional] 
**Scan** | Pointer to [**ScanIdTime**](ScanIdTime.md) |  | [optional] 
**Session** | Pointer to **int32** | Process session | [optional] 
**Sgroup** | Pointer to **string** | Saved-set group | [optional] 
**Share** | Pointer to **int32** | Shared memory | [optional] 
**Size** | Pointer to **int32** | Process size in bytes | [optional] 
**StartTime** | Pointer to **int64** | Time when the process started | [optional] 
**State** | Pointer to **string** | Process state | [optional] 
**Stime** | Pointer to **int32** | Time spent executing system code | [optional] 
**Suser** | Pointer to **string** | Saved-set user | [optional] 
**Tgid** | Pointer to **int32** | Thread Group ID | [optional] 
**Tty** | Pointer to **int32** | Process TTY number | [optional] 
**Utime** | Pointer to **int32** | Time spent executing user code | [optional] 
**VmSize** | Pointer to **int32** | Total VM size (KB) | [optional] 
**AgentId** | Pointer to **string** | Agent ID | [optional] 

## Methods

### NewSyscollectorProcess

`func NewSyscollectorProcess() *SyscollectorProcess`

NewSyscollectorProcess instantiates a new SyscollectorProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyscollectorProcessWithDefaults

`func NewSyscollectorProcessWithDefaults() *SyscollectorProcess`

NewSyscollectorProcessWithDefaults instantiates a new SyscollectorProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgvs

`func (o *SyscollectorProcess) GetArgvs() string`

GetArgvs returns the Argvs field if non-nil, zero value otherwise.

### GetArgvsOk

`func (o *SyscollectorProcess) GetArgvsOk() (*string, bool)`

GetArgvsOk returns a tuple with the Argvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgvs

`func (o *SyscollectorProcess) SetArgvs(v string)`

SetArgvs sets Argvs field to given value.

### HasArgvs

`func (o *SyscollectorProcess) HasArgvs() bool`

HasArgvs returns a boolean if a field has been set.

### GetCmd

`func (o *SyscollectorProcess) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *SyscollectorProcess) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *SyscollectorProcess) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *SyscollectorProcess) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetEgroup

`func (o *SyscollectorProcess) GetEgroup() string`

GetEgroup returns the Egroup field if non-nil, zero value otherwise.

### GetEgroupOk

`func (o *SyscollectorProcess) GetEgroupOk() (*string, bool)`

GetEgroupOk returns a tuple with the Egroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgroup

`func (o *SyscollectorProcess) SetEgroup(v string)`

SetEgroup sets Egroup field to given value.

### HasEgroup

`func (o *SyscollectorProcess) HasEgroup() bool`

HasEgroup returns a boolean if a field has been set.

### GetEuser

`func (o *SyscollectorProcess) GetEuser() string`

GetEuser returns the Euser field if non-nil, zero value otherwise.

### GetEuserOk

`func (o *SyscollectorProcess) GetEuserOk() (*string, bool)`

GetEuserOk returns a tuple with the Euser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuser

`func (o *SyscollectorProcess) SetEuser(v string)`

SetEuser sets Euser field to given value.

### HasEuser

`func (o *SyscollectorProcess) HasEuser() bool`

HasEuser returns a boolean if a field has been set.

### GetFgroup

`func (o *SyscollectorProcess) GetFgroup() string`

GetFgroup returns the Fgroup field if non-nil, zero value otherwise.

### GetFgroupOk

`func (o *SyscollectorProcess) GetFgroupOk() (*string, bool)`

GetFgroupOk returns a tuple with the Fgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgroup

`func (o *SyscollectorProcess) SetFgroup(v string)`

SetFgroup sets Fgroup field to given value.

### HasFgroup

`func (o *SyscollectorProcess) HasFgroup() bool`

HasFgroup returns a boolean if a field has been set.

### GetName

`func (o *SyscollectorProcess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyscollectorProcess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyscollectorProcess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyscollectorProcess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNice

`func (o *SyscollectorProcess) GetNice() int32`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *SyscollectorProcess) GetNiceOk() (*int32, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *SyscollectorProcess) SetNice(v int32)`

SetNice sets Nice field to given value.

### HasNice

`func (o *SyscollectorProcess) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetNlwp

`func (o *SyscollectorProcess) GetNlwp() int32`

GetNlwp returns the Nlwp field if non-nil, zero value otherwise.

### GetNlwpOk

`func (o *SyscollectorProcess) GetNlwpOk() (*int32, bool)`

GetNlwpOk returns a tuple with the Nlwp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNlwp

`func (o *SyscollectorProcess) SetNlwp(v int32)`

SetNlwp sets Nlwp field to given value.

### HasNlwp

`func (o *SyscollectorProcess) HasNlwp() bool`

HasNlwp returns a boolean if a field has been set.

### GetPgrp

`func (o *SyscollectorProcess) GetPgrp() int32`

GetPgrp returns the Pgrp field if non-nil, zero value otherwise.

### GetPgrpOk

`func (o *SyscollectorProcess) GetPgrpOk() (*int32, bool)`

GetPgrpOk returns a tuple with the Pgrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgrp

`func (o *SyscollectorProcess) SetPgrp(v int32)`

SetPgrp sets Pgrp field to given value.

### HasPgrp

`func (o *SyscollectorProcess) HasPgrp() bool`

HasPgrp returns a boolean if a field has been set.

### GetPid

`func (o *SyscollectorProcess) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *SyscollectorProcess) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *SyscollectorProcess) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *SyscollectorProcess) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPpid

`func (o *SyscollectorProcess) GetPpid() int32`

GetPpid returns the Ppid field if non-nil, zero value otherwise.

### GetPpidOk

`func (o *SyscollectorProcess) GetPpidOk() (*int32, bool)`

GetPpidOk returns a tuple with the Ppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpid

`func (o *SyscollectorProcess) SetPpid(v int32)`

SetPpid sets Ppid field to given value.

### HasPpid

`func (o *SyscollectorProcess) HasPpid() bool`

HasPpid returns a boolean if a field has been set.

### GetPriority

`func (o *SyscollectorProcess) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SyscollectorProcess) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SyscollectorProcess) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SyscollectorProcess) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProcessor

`func (o *SyscollectorProcess) GetProcessor() int32`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *SyscollectorProcess) GetProcessorOk() (*int32, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *SyscollectorProcess) SetProcessor(v int32)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *SyscollectorProcess) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetResident

`func (o *SyscollectorProcess) GetResident() int32`

GetResident returns the Resident field if non-nil, zero value otherwise.

### GetResidentOk

`func (o *SyscollectorProcess) GetResidentOk() (*int32, bool)`

GetResidentOk returns a tuple with the Resident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResident

`func (o *SyscollectorProcess) SetResident(v int32)`

SetResident sets Resident field to given value.

### HasResident

`func (o *SyscollectorProcess) HasResident() bool`

HasResident returns a boolean if a field has been set.

### GetRgroup

`func (o *SyscollectorProcess) GetRgroup() string`

GetRgroup returns the Rgroup field if non-nil, zero value otherwise.

### GetRgroupOk

`func (o *SyscollectorProcess) GetRgroupOk() (*string, bool)`

GetRgroupOk returns a tuple with the Rgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRgroup

`func (o *SyscollectorProcess) SetRgroup(v string)`

SetRgroup sets Rgroup field to given value.

### HasRgroup

`func (o *SyscollectorProcess) HasRgroup() bool`

HasRgroup returns a boolean if a field has been set.

### GetRuser

`func (o *SyscollectorProcess) GetRuser() string`

GetRuser returns the Ruser field if non-nil, zero value otherwise.

### GetRuserOk

`func (o *SyscollectorProcess) GetRuserOk() (*string, bool)`

GetRuserOk returns a tuple with the Ruser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuser

`func (o *SyscollectorProcess) SetRuser(v string)`

SetRuser sets Ruser field to given value.

### HasRuser

`func (o *SyscollectorProcess) HasRuser() bool`

HasRuser returns a boolean if a field has been set.

### GetScan

`func (o *SyscollectorProcess) GetScan() ScanIdTime`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *SyscollectorProcess) GetScanOk() (*ScanIdTime, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *SyscollectorProcess) SetScan(v ScanIdTime)`

SetScan sets Scan field to given value.

### HasScan

`func (o *SyscollectorProcess) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetSession

`func (o *SyscollectorProcess) GetSession() int32`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SyscollectorProcess) GetSessionOk() (*int32, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SyscollectorProcess) SetSession(v int32)`

SetSession sets Session field to given value.

### HasSession

`func (o *SyscollectorProcess) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSgroup

`func (o *SyscollectorProcess) GetSgroup() string`

GetSgroup returns the Sgroup field if non-nil, zero value otherwise.

### GetSgroupOk

`func (o *SyscollectorProcess) GetSgroupOk() (*string, bool)`

GetSgroupOk returns a tuple with the Sgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgroup

`func (o *SyscollectorProcess) SetSgroup(v string)`

SetSgroup sets Sgroup field to given value.

### HasSgroup

`func (o *SyscollectorProcess) HasSgroup() bool`

HasSgroup returns a boolean if a field has been set.

### GetShare

`func (o *SyscollectorProcess) GetShare() int32`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *SyscollectorProcess) GetShareOk() (*int32, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *SyscollectorProcess) SetShare(v int32)`

SetShare sets Share field to given value.

### HasShare

`func (o *SyscollectorProcess) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSize

`func (o *SyscollectorProcess) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SyscollectorProcess) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SyscollectorProcess) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SyscollectorProcess) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartTime

`func (o *SyscollectorProcess) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SyscollectorProcess) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SyscollectorProcess) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SyscollectorProcess) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *SyscollectorProcess) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SyscollectorProcess) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SyscollectorProcess) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SyscollectorProcess) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStime

`func (o *SyscollectorProcess) GetStime() int32`

GetStime returns the Stime field if non-nil, zero value otherwise.

### GetStimeOk

`func (o *SyscollectorProcess) GetStimeOk() (*int32, bool)`

GetStimeOk returns a tuple with the Stime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStime

`func (o *SyscollectorProcess) SetStime(v int32)`

SetStime sets Stime field to given value.

### HasStime

`func (o *SyscollectorProcess) HasStime() bool`

HasStime returns a boolean if a field has been set.

### GetSuser

`func (o *SyscollectorProcess) GetSuser() string`

GetSuser returns the Suser field if non-nil, zero value otherwise.

### GetSuserOk

`func (o *SyscollectorProcess) GetSuserOk() (*string, bool)`

GetSuserOk returns a tuple with the Suser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuser

`func (o *SyscollectorProcess) SetSuser(v string)`

SetSuser sets Suser field to given value.

### HasSuser

`func (o *SyscollectorProcess) HasSuser() bool`

HasSuser returns a boolean if a field has been set.

### GetTgid

`func (o *SyscollectorProcess) GetTgid() int32`

GetTgid returns the Tgid field if non-nil, zero value otherwise.

### GetTgidOk

`func (o *SyscollectorProcess) GetTgidOk() (*int32, bool)`

GetTgidOk returns a tuple with the Tgid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgid

`func (o *SyscollectorProcess) SetTgid(v int32)`

SetTgid sets Tgid field to given value.

### HasTgid

`func (o *SyscollectorProcess) HasTgid() bool`

HasTgid returns a boolean if a field has been set.

### GetTty

`func (o *SyscollectorProcess) GetTty() int32`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *SyscollectorProcess) GetTtyOk() (*int32, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *SyscollectorProcess) SetTty(v int32)`

SetTty sets Tty field to given value.

### HasTty

`func (o *SyscollectorProcess) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUtime

`func (o *SyscollectorProcess) GetUtime() int32`

GetUtime returns the Utime field if non-nil, zero value otherwise.

### GetUtimeOk

`func (o *SyscollectorProcess) GetUtimeOk() (*int32, bool)`

GetUtimeOk returns a tuple with the Utime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtime

`func (o *SyscollectorProcess) SetUtime(v int32)`

SetUtime sets Utime field to given value.

### HasUtime

`func (o *SyscollectorProcess) HasUtime() bool`

HasUtime returns a boolean if a field has been set.

### GetVmSize

`func (o *SyscollectorProcess) GetVmSize() int32`

GetVmSize returns the VmSize field if non-nil, zero value otherwise.

### GetVmSizeOk

`func (o *SyscollectorProcess) GetVmSizeOk() (*int32, bool)`

GetVmSizeOk returns a tuple with the VmSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSize

`func (o *SyscollectorProcess) SetVmSize(v int32)`

SetVmSize sets VmSize field to given value.

### HasVmSize

`func (o *SyscollectorProcess) HasVmSize() bool`

HasVmSize returns a boolean if a field has been set.

### GetAgentId

`func (o *SyscollectorProcess) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *SyscollectorProcess) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *SyscollectorProcess) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *SyscollectorProcess) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


