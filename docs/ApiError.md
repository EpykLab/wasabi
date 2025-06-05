# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Remediation** | Pointer to **string** |  | [optional] 
**DapiErrors** | Pointer to [**map[string]ApiErrorDapiErrorsValue**](ApiErrorDapiErrorsValue.md) |  | [optional] 

## Methods

### NewApiError

`func NewApiError(title string, detail string, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ApiError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ApiError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ApiError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ApiError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ApiError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ApiError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ApiError) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ApiError) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCode

`func (o *ApiError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRemediation

`func (o *ApiError) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *ApiError) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *ApiError) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *ApiError) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetDapiErrors

`func (o *ApiError) GetDapiErrors() map[string]ApiErrorDapiErrorsValue`

GetDapiErrors returns the DapiErrors field if non-nil, zero value otherwise.

### GetDapiErrorsOk

`func (o *ApiError) GetDapiErrorsOk() (*map[string]ApiErrorDapiErrorsValue, bool)`

GetDapiErrorsOk returns a tuple with the DapiErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDapiErrors

`func (o *ApiError) SetDapiErrors(v map[string]ApiErrorDapiErrorsValue)`

SetDapiErrors sets DapiErrors field to given value.

### HasDapiErrors

`func (o *ApiError) HasDapiErrors() bool`

HasDapiErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


