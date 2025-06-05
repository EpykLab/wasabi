# AgentGroupDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedItems** | **[]string** | List of removed groups, agents belonging exclusively to the removed groups will be reassigned to group default | 

## Methods

### NewAgentGroupDeleted

`func NewAgentGroupDeleted(affectedItems []string, ) *AgentGroupDeleted`

NewAgentGroupDeleted instantiates a new AgentGroupDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentGroupDeletedWithDefaults

`func NewAgentGroupDeletedWithDefaults() *AgentGroupDeleted`

NewAgentGroupDeletedWithDefaults instantiates a new AgentGroupDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedItems

`func (o *AgentGroupDeleted) GetAffectedItems() []string`

GetAffectedItems returns the AffectedItems field if non-nil, zero value otherwise.

### GetAffectedItemsOk

`func (o *AgentGroupDeleted) GetAffectedItemsOk() (*[]string, bool)`

GetAffectedItemsOk returns a tuple with the AffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedItems

`func (o *AgentGroupDeleted) SetAffectedItems(v []string)`

SetAffectedItems sets AffectedItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


