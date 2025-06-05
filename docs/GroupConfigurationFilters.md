# GroupConfigurationFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to **string** | OS family where the configuration is being applied | [optional] 
**Name** | Pointer to **string** | The name of the agent where that configuration is being applied | [optional] 
**Profile** | Pointer to **string** | Profile name. Any agent configured to use the defined profile may use the block | [optional] 

## Methods

### NewGroupConfigurationFilters

`func NewGroupConfigurationFilters() *GroupConfigurationFilters`

NewGroupConfigurationFilters instantiates a new GroupConfigurationFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupConfigurationFiltersWithDefaults

`func NewGroupConfigurationFiltersWithDefaults() *GroupConfigurationFilters`

NewGroupConfigurationFiltersWithDefaults instantiates a new GroupConfigurationFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *GroupConfigurationFilters) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *GroupConfigurationFilters) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *GroupConfigurationFilters) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *GroupConfigurationFilters) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetName

`func (o *GroupConfigurationFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupConfigurationFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupConfigurationFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupConfigurationFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *GroupConfigurationFilters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GroupConfigurationFilters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GroupConfigurationFilters) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GroupConfigurationFilters) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


