# PoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Policy name | 
**Policy** | [**PoliciesRequestPolicy**](PoliciesRequestPolicy.md) |  | 

## Methods

### NewPoliciesRequest

`func NewPoliciesRequest(name string, policy PoliciesRequestPolicy, ) *PoliciesRequest`

NewPoliciesRequest instantiates a new PoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesRequestWithDefaults

`func NewPoliciesRequestWithDefaults() *PoliciesRequest`

NewPoliciesRequestWithDefaults instantiates a new PoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoliciesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoliciesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoliciesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *PoliciesRequest) GetPolicy() PoliciesRequestPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PoliciesRequest) GetPolicyOk() (*PoliciesRequestPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PoliciesRequest) SetPolicy(v PoliciesRequestPolicy)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


