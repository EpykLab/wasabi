# SimpleApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**map[string]SimpleApiErrorErrorValue**](SimpleApiErrorErrorValue.md) |  | 
**Id** | Pointer to [**[]SimpleApiErrorIdInner**](SimpleApiErrorIdInner.md) |  | [optional] 

## Methods

### NewSimpleApiError

`func NewSimpleApiError(error_ map[string]SimpleApiErrorErrorValue, ) *SimpleApiError`

NewSimpleApiError instantiates a new SimpleApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiErrorWithDefaults

`func NewSimpleApiErrorWithDefaults() *SimpleApiError`

NewSimpleApiErrorWithDefaults instantiates a new SimpleApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *SimpleApiError) GetError() map[string]SimpleApiErrorErrorValue`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SimpleApiError) GetErrorOk() (*map[string]SimpleApiErrorErrorValue, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SimpleApiError) SetError(v map[string]SimpleApiErrorErrorValue)`

SetError sets Error field to given value.


### GetId

`func (o *SimpleApiError) GetId() []SimpleApiErrorIdInner`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleApiError) GetIdOk() (*[]SimpleApiErrorIdInner, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleApiError) SetId(v []SimpleApiErrorIdInner)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleApiError) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


