# NodeRulesetSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node name | [optional] 
**Synced** | Pointer to **bool** | Whether the ruleset is synchronized or not | [optional] 

## Methods

### NewNodeRulesetSyncStatus

`func NewNodeRulesetSyncStatus() *NodeRulesetSyncStatus`

NewNodeRulesetSyncStatus instantiates a new NodeRulesetSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRulesetSyncStatusWithDefaults

`func NewNodeRulesetSyncStatusWithDefaults() *NodeRulesetSyncStatus`

NewNodeRulesetSyncStatusWithDefaults instantiates a new NodeRulesetSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeRulesetSyncStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeRulesetSyncStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeRulesetSyncStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeRulesetSyncStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSynced

`func (o *NodeRulesetSyncStatus) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *NodeRulesetSyncStatus) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *NodeRulesetSyncStatus) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *NodeRulesetSyncStatus) HasSynced() bool`

HasSynced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


