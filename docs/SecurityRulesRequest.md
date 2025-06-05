# SecurityRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Rule name | 
**Rule** | **map[string]interface{}** | Rule body | 

## Methods

### NewSecurityRulesRequest

`func NewSecurityRulesRequest(name string, rule map[string]interface{}, ) *SecurityRulesRequest`

NewSecurityRulesRequest instantiates a new SecurityRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRulesRequestWithDefaults

`func NewSecurityRulesRequestWithDefaults() *SecurityRulesRequest`

NewSecurityRulesRequestWithDefaults instantiates a new SecurityRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityRulesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityRulesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityRulesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRule

`func (o *SecurityRulesRequest) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *SecurityRulesRequest) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *SecurityRulesRequest) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


