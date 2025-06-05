# PoliciesResponsePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | Actions to perform | [optional] 
**Resources** | Pointer to **[]string** | Resources to apply the actions on | [optional] 
**Effect** | Pointer to **string** | Effect of the policy | [optional] 

## Methods

### NewPoliciesResponsePolicy

`func NewPoliciesResponsePolicy() *PoliciesResponsePolicy`

NewPoliciesResponsePolicy instantiates a new PoliciesResponsePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesResponsePolicyWithDefaults

`func NewPoliciesResponsePolicyWithDefaults() *PoliciesResponsePolicy`

NewPoliciesResponsePolicyWithDefaults instantiates a new PoliciesResponsePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *PoliciesResponsePolicy) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PoliciesResponsePolicy) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PoliciesResponsePolicy) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PoliciesResponsePolicy) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetResources

`func (o *PoliciesResponsePolicy) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PoliciesResponsePolicy) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PoliciesResponsePolicy) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PoliciesResponsePolicy) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetEffect

`func (o *PoliciesResponsePolicy) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PoliciesResponsePolicy) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PoliciesResponsePolicy) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *PoliciesResponsePolicy) HasEffect() bool`

HasEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


