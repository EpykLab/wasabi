# NodeHealthcheckNameInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node ID | [optional] 
**Info** | Pointer to [**NodeHealthcheckNameInfoInfo**](NodeHealthcheckNameInfoInfo.md) |  | [optional] 

## Methods

### NewNodeHealthcheckNameInfo

`func NewNodeHealthcheckNameInfo() *NodeHealthcheckNameInfo`

NewNodeHealthcheckNameInfo instantiates a new NodeHealthcheckNameInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHealthcheckNameInfoWithDefaults

`func NewNodeHealthcheckNameInfoWithDefaults() *NodeHealthcheckNameInfo`

NewNodeHealthcheckNameInfoWithDefaults instantiates a new NodeHealthcheckNameInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeHealthcheckNameInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeHealthcheckNameInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeHealthcheckNameInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeHealthcheckNameInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *NodeHealthcheckNameInfo) GetInfo() NodeHealthcheckNameInfoInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *NodeHealthcheckNameInfo) GetInfoOk() (*NodeHealthcheckNameInfoInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *NodeHealthcheckNameInfo) SetInfo(v NodeHealthcheckNameInfoInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *NodeHealthcheckNameInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


