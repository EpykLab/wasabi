# AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Agents** | Pointer to [**[]WazuhAnalysisdStatsAgentsItemAgentsInner**](WazuhAnalysisdStatsAgentsItemAgentsInner.md) |  | [optional] 

## Methods

### NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems

`func NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems() *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems`

NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems instantiates a new AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItemsWithDefaults

`func NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItemsWithDefaults() *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems`

NewAllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItemsWithDefaults instantiates a new AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAgents

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetAgents() []WazuhAnalysisdStatsAgentsItemAgentsInner`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) GetAgentsOk() (*[]WazuhAnalysisdStatsAgentsItemAgentsInner, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) SetAgents(v []WazuhAnalysisdStatsAgentsItemAgentsInner)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AllItemsResponseWazuhDaemonStatsAgentsAllOfAffectedItems) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


