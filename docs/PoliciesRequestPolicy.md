# PoliciesRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | Actions to perform | 
**Resources** | **[]string** | Resources to apply the actions on | 
**Effect** | **string** | Effect of the policy | 

## Methods

### NewPoliciesRequestPolicy

`func NewPoliciesRequestPolicy(actions []string, resources []string, effect string, ) *PoliciesRequestPolicy`

NewPoliciesRequestPolicy instantiates a new PoliciesRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoliciesRequestPolicyWithDefaults

`func NewPoliciesRequestPolicyWithDefaults() *PoliciesRequestPolicy`

NewPoliciesRequestPolicyWithDefaults instantiates a new PoliciesRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *PoliciesRequestPolicy) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PoliciesRequestPolicy) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PoliciesRequestPolicy) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetResources

`func (o *PoliciesRequestPolicy) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PoliciesRequestPolicy) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PoliciesRequestPolicy) SetResources(v []string)`

SetResources sets Resources field to given value.


### GetEffect

`func (o *PoliciesRequestPolicy) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PoliciesRequestPolicy) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PoliciesRequestPolicy) SetEffect(v string)`

SetEffect sets Effect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


