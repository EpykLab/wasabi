# GroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**GroupConfigurationFilters**](GroupConfigurationFilters.md) |  | 
**Config** | **map[string]interface{}** | Group configuration. The fields on this object depend on the actual group configuration | 

## Methods

### NewGroupConfiguration

`func NewGroupConfiguration(filters GroupConfigurationFilters, config map[string]interface{}, ) *GroupConfiguration`

NewGroupConfiguration instantiates a new GroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupConfigurationWithDefaults

`func NewGroupConfigurationWithDefaults() *GroupConfiguration`

NewGroupConfigurationWithDefaults instantiates a new GroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GroupConfiguration) GetFilters() GroupConfigurationFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GroupConfiguration) GetFiltersOk() (*GroupConfigurationFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GroupConfiguration) SetFilters(v GroupConfigurationFilters)`

SetFilters sets Filters field to given value.


### GetConfig

`func (o *GroupConfiguration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GroupConfiguration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GroupConfiguration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


