# PoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Policy id | [optional] 
**Name** | Pointer to **string** | Policy name | [optional] 
**Policy** | Pointer to [**PoliciesResponsePolicy**](PoliciesResponsePolicy.md) |  | [optional] 

## Methods

### NewPoliciesResponse

`func NewPoliciesResponse() *PoliciesResponse`

NewPoliciesResponse instantiates a new PoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesResponseWithDefaults

`func NewPoliciesResponseWithDefaults() *PoliciesResponse`

NewPoliciesResponseWithDefaults instantiates a new PoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoliciesResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoliciesResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoliciesResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PoliciesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PoliciesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoliciesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoliciesResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoliciesResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *PoliciesResponse) GetPolicy() PoliciesResponsePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PoliciesResponse) GetPolicyOk() (*PoliciesResponsePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PoliciesResponse) SetPolicy(v PoliciesResponsePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PoliciesResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


