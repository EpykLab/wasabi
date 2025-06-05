# WazuhRemotedStatsAgentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Agents** | Pointer to [**[]WazuhRemotedStatsAgentsItemAgentsInner**](WazuhRemotedStatsAgentsItemAgentsInner.md) |  | [optional] 

## Methods

### NewWazuhRemotedStatsAgentsItem

`func NewWazuhRemotedStatsAgentsItem() *WazuhRemotedStatsAgentsItem`

NewWazuhRemotedStatsAgentsItem instantiates a new WazuhRemotedStatsAgentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhRemotedStatsAgentsItemWithDefaults

`func NewWazuhRemotedStatsAgentsItemWithDefaults() *WazuhRemotedStatsAgentsItem`

NewWazuhRemotedStatsAgentsItemWithDefaults instantiates a new WazuhRemotedStatsAgentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WazuhRemotedStatsAgentsItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WazuhRemotedStatsAgentsItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WazuhRemotedStatsAgentsItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WazuhRemotedStatsAgentsItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *WazuhRemotedStatsAgentsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WazuhRemotedStatsAgentsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WazuhRemotedStatsAgentsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WazuhRemotedStatsAgentsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAgents

`func (o *WazuhRemotedStatsAgentsItem) GetAgents() []WazuhRemotedStatsAgentsItemAgentsInner`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *WazuhRemotedStatsAgentsItem) GetAgentsOk() (*[]WazuhRemotedStatsAgentsItemAgentsInner, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *WazuhRemotedStatsAgentsItem) SetAgents(v []WazuhRemotedStatsAgentsItemAgentsInner)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *WazuhRemotedStatsAgentsItem) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


