# SecurityRulesRequestNoRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Rule name | [optional] 
**Rule** | Pointer to **map[string]interface{}** | Rule body | [optional] 

## Methods

### NewSecurityRulesRequestNoRequired

`func NewSecurityRulesRequestNoRequired() *SecurityRulesRequestNoRequired`

NewSecurityRulesRequestNoRequired instantiates a new SecurityRulesRequestNoRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesRequestNoRequiredWithDefaults

`func NewSecurityRulesRequestNoRequiredWithDefaults() *SecurityRulesRequestNoRequired`

NewSecurityRulesRequestNoRequiredWithDefaults instantiates a new SecurityRulesRequestNoRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityRulesRequestNoRequired) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityRulesRequestNoRequired) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityRulesRequestNoRequired) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityRulesRequestNoRequired) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRule

`func (o *SecurityRulesRequestNoRequired) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *SecurityRulesRequestNoRequired) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *SecurityRulesRequestNoRequired) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.

### HasRule

`func (o *SecurityRulesRequestNoRequired) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


