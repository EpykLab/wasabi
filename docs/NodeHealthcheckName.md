# NodeHealthcheckName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**NodeHealthcheckNameInfo**](NodeHealthcheckNameInfo.md) |  | [optional] 
**Status** | Pointer to [**NodeHealthcheckNameStatus**](NodeHealthcheckNameStatus.md) |  | [optional] 

## Methods

### NewNodeHealthcheckName

`func NewNodeHealthcheckName() *NodeHealthcheckName`

NewNodeHealthcheckName instantiates a new NodeHealthcheckName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeHealthcheckNameWithDefaults

`func NewNodeHealthcheckNameWithDefaults() *NodeHealthcheckName`

NewNodeHealthcheckNameWithDefaults instantiates a new NodeHealthcheckName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *NodeHealthcheckName) GetInfo() NodeHealthcheckNameInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *NodeHealthcheckName) GetInfoOk() (*NodeHealthcheckNameInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *NodeHealthcheckName) SetInfo(v NodeHealthcheckNameInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *NodeHealthcheckName) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetStatus

`func (o *NodeHealthcheckName) GetStatus() NodeHealthcheckNameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeHealthcheckName) GetStatusOk() (*NodeHealthcheckNameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeHealthcheckName) SetStatus(v NodeHealthcheckNameStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeHealthcheckName) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


