# RolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Role id | [optional] 
**Name** | Pointer to **string** | Role name | [optional] 
**Rule** | Pointer to **map[string]interface{}** | Role rule | [optional] 

## Methods

### NewRolesResponse

`func NewRolesResponse() *RolesResponse`

NewRolesResponse instantiates a new RolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesResponseWithDefaults

`func NewRolesResponseWithDefaults() *RolesResponse`

NewRolesResponseWithDefaults instantiates a new RolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolesResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolesResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolesResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RolesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RolesResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolesResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolesResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolesResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRule

`func (o *RolesResponse) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RolesResponse) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RolesResponse) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.

### HasRule

`func (o *RolesResponse) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


