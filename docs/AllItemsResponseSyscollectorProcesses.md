# AllItemsResponseSyscollectorProcesses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAffectedItems** | **int32** | Number of items that have successfully applied the requested operation | 
**FailedItems** | [**[]SimpleApiError**](SimpleApiError.md) | List of items that have failed applying the requested operation | 
**TotalFailedItems** | **int32** | Number of items that have failed applying the requested operation | 
**AffectedItems** | [**[]SyscollectorProcess**](SyscollectorProcess.md) | Items that successfully applied the API call action | 

## Methods

### NewAllItemsResponseSyscollectorProcesses

`func NewAllItemsResponseSyscollectorProcesses(totalAffectedItems int32, failedItems []SimpleApiError, totalFailedItems int32, affectedItems []SyscollectorProcess, ) *AllItemsResponseSyscollectorProcesses`

NewAllItemsResponseSyscollectorProcesses instantiates a new AllItemsResponseSyscollectorProcesses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseSyscollectorProcessesWithDefaults

`func NewAllItemsResponseSyscollectorProcessesWithDefaults() *AllItemsResponseSyscollectorProcesses`

NewAllItemsResponseSyscollectorProcessesWithDefaults instantiates a new AllItemsResponseSyscollectorProcesses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAffectedItems

`func (o *AllItemsResponseSyscollectorProcesses) GetTotalAffectedItems() int32`

GetTotalAffectedItems returns the TotalAffectedItems field if non-nil, zero value otherwise.

### GetTotalAffectedItemsOk

`func (o *AllItemsResponseSyscollectorProcesses) GetTotalAffectedItemsOk() (*int32, bool)`

GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedItems

`func (o *AllItemsResponseSyscollectorProcesses) SetTotalAffectedItems(v int32)`

SetTotalAffectedItems sets TotalAffectedItems field to given value.


### GetFailedItems

`func (o *AllItemsResponseSyscollectorProcesses) GetFailedItems() []SimpleApiError`

GetFailedItems returns the FailedItems field if non-nil, zero value otherwise.

### GetFailedItemsOk

`func (o *AllItemsResponseSyscollectorProcesses) GetFailedItemsOk() (*[]SimpleApiError, bool)`

GetFailedItemsOk returns a tuple with the FailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItems

`func (o *AllItemsResponseSyscollectorProcesses) SetFailedItems(v []SimpleApiError)`

SetFailedItems sets FailedItems field to given value.


### GetTotalFailedItems

`func (o *AllItemsResponseSyscollectorProcesses) GetTotalFailedItems() int32`

GetTotalFailedItems returns the TotalFailedItems field if non-nil, zero value otherwise.

### GetTotalFailedItemsOk

`func (o *AllItemsResponseSyscollectorProcesses) GetTotalFailedItemsOk() (*int32, bool)`

GetTotalFailedItemsOk returns a tuple with the TotalFailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailedItems

`func (o *AllItemsResponseSyscollectorProcesses) SetTotalFailedItems(v int32)`

SetTotalFailedItems sets TotalFailedItems field to given value.


### GetAffectedItems

`func (o *AllItemsResponseSyscollectorProcesses) GetAffectedItems() []SyscollectorProcess`

GetAffectedItems returns the AffectedItems field if non-nil, zero value otherwise.

### GetAffectedItemsOk

`func (o *AllItemsResponseSyscollectorProcesses) GetAffectedItemsOk() (*[]SyscollectorProcess, bool)`

GetAffectedItemsOk returns a tuple with the AffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedItems

`func (o *AllItemsResponseSyscollectorProcesses) SetAffectedItems(v []SyscollectorProcess)`

SetAffectedItems sets AffectedItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


