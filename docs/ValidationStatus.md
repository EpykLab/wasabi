# ValidationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node name | [optional] 
**Status** | Pointer to **string** | Status value | [optional] 

## Methods

### NewValidationStatus

`func NewValidationStatus() *ValidationStatus`

NewValidationStatus instantiates a new ValidationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationStatusWithDefaults

`func NewValidationStatusWithDefaults() *ValidationStatus`

NewValidationStatusWithDefaults instantiates a new ValidationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidationStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidationStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidationStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValidationStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ValidationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ValidationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


