# AllItemsResponseLists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAffectedItems** | **int32** | Number of items that have successfully applied the requested operation | 
**FailedItems** | [**[]SimpleApiError**](SimpleApiError.md) | List of items that have failed applying the requested operation | 
**TotalFailedItems** | **int32** | Number of items that have failed applying the requested operation | 
**AffectedItems** | [**[]CDBList**](CDBList.md) | Items that successfully applied the API call action | 

## Methods

### NewAllItemsResponseLists

`func NewAllItemsResponseLists(totalAffectedItems int32, failedItems []SimpleApiError, totalFailedItems int32, affectedItems []CDBList, ) *AllItemsResponseLists`

NewAllItemsResponseLists instantiates a new AllItemsResponseLists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseListsWithDefaults

`func NewAllItemsResponseListsWithDefaults() *AllItemsResponseLists`

NewAllItemsResponseListsWithDefaults instantiates a new AllItemsResponseLists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAffectedItems

`func (o *AllItemsResponseLists) GetTotalAffectedItems() int32`

GetTotalAffectedItems returns the TotalAffectedItems field if non-nil, zero value otherwise.

### GetTotalAffectedItemsOk

`func (o *AllItemsResponseLists) GetTotalAffectedItemsOk() (*int32, bool)`

GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedItems

`func (o *AllItemsResponseLists) SetTotalAffectedItems(v int32)`

SetTotalAffectedItems sets TotalAffectedItems field to given value.


### GetFailedItems

`func (o *AllItemsResponseLists) GetFailedItems() []SimpleApiError`

GetFailedItems returns the FailedItems field if non-nil, zero value otherwise.

### GetFailedItemsOk

`func (o *AllItemsResponseLists) GetFailedItemsOk() (*[]SimpleApiError, bool)`

GetFailedItemsOk returns a tuple with the FailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItems

`func (o *AllItemsResponseLists) SetFailedItems(v []SimpleApiError)`

SetFailedItems sets FailedItems field to given value.


### GetTotalFailedItems

`func (o *AllItemsResponseLists) GetTotalFailedItems() int32`

GetTotalFailedItems returns the TotalFailedItems field if non-nil, zero value otherwise.

### GetTotalFailedItemsOk

`func (o *AllItemsResponseLists) GetTotalFailedItemsOk() (*int32, bool)`

GetTotalFailedItemsOk returns a tuple with the TotalFailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailedItems

`func (o *AllItemsResponseLists) SetTotalFailedItems(v int32)`

SetTotalFailedItems sets TotalFailedItems field to given value.


### GetAffectedItems

`func (o *AllItemsResponseLists) GetAffectedItems() []CDBList`

GetAffectedItems returns the AffectedItems field if non-nil, zero value otherwise.

### GetAffectedItemsOk

`func (o *AllItemsResponseLists) GetAffectedItemsOk() (*[]CDBList, bool)`

GetAffectedItemsOk returns a tuple with the AffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedItems

`func (o *AllItemsResponseLists) SetAffectedItems(v []CDBList)`

SetAffectedItems sets AffectedItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


