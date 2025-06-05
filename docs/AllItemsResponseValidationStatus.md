# AllItemsResponseValidationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAffectedItems** | **int32** | Number of items that have successfully applied the requested operation | 
**FailedItems** | [**[]SimpleApiError**](SimpleApiError.md) | List of items that have failed applying the requested operation | 
**TotalFailedItems** | **int32** | Number of items that have failed applying the requested operation | 
**AffectedItems** | [**[]ValidationStatus**](ValidationStatus.md) | Items that successfully applied the API call action | 

## Methods

### NewAllItemsResponseValidationStatus

`func NewAllItemsResponseValidationStatus(totalAffectedItems int32, failedItems []SimpleApiError, totalFailedItems int32, affectedItems []ValidationStatus, ) *AllItemsResponseValidationStatus`

NewAllItemsResponseValidationStatus instantiates a new AllItemsResponseValidationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseValidationStatusWithDefaults

`func NewAllItemsResponseValidationStatusWithDefaults() *AllItemsResponseValidationStatus`

NewAllItemsResponseValidationStatusWithDefaults instantiates a new AllItemsResponseValidationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAffectedItems

`func (o *AllItemsResponseValidationStatus) GetTotalAffectedItems() int32`

GetTotalAffectedItems returns the TotalAffectedItems field if non-nil, zero value otherwise.

### GetTotalAffectedItemsOk

`func (o *AllItemsResponseValidationStatus) GetTotalAffectedItemsOk() (*int32, bool)`

GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedItems

`func (o *AllItemsResponseValidationStatus) SetTotalAffectedItems(v int32)`

SetTotalAffectedItems sets TotalAffectedItems field to given value.


### GetFailedItems

`func (o *AllItemsResponseValidationStatus) GetFailedItems() []SimpleApiError`

GetFailedItems returns the FailedItems field if non-nil, zero value otherwise.

### GetFailedItemsOk

`func (o *AllItemsResponseValidationStatus) GetFailedItemsOk() (*[]SimpleApiError, bool)`

GetFailedItemsOk returns a tuple with the FailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedItems

`func (o *AllItemsResponseValidationStatus) SetFailedItems(v []SimpleApiError)`

SetFailedItems sets FailedItems field to given value.


### GetTotalFailedItems

`func (o *AllItemsResponseValidationStatus) GetTotalFailedItems() int32`

GetTotalFailedItems returns the TotalFailedItems field if non-nil, zero value otherwise.

### GetTotalFailedItemsOk

`func (o *AllItemsResponseValidationStatus) GetTotalFailedItemsOk() (*int32, bool)`

GetTotalFailedItemsOk returns a tuple with the TotalFailedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailedItems

`func (o *AllItemsResponseValidationStatus) SetTotalFailedItems(v int32)`

SetTotalFailedItems sets TotalFailedItems field to given value.


### GetAffectedItems

`func (o *AllItemsResponseValidationStatus) GetAffectedItems() []ValidationStatus`

GetAffectedItems returns the AffectedItems field if non-nil, zero value otherwise.

### GetAffectedItemsOk

`func (o *AllItemsResponseValidationStatus) GetAffectedItemsOk() (*[]ValidationStatus, bool)`

GetAffectedItemsOk returns a tuple with the AffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedItems

`func (o *AllItemsResponseValidationStatus) SetAffectedItems(v []ValidationStatus)`

SetAffectedItems sets AffectedItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


