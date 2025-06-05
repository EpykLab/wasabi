# WazuhAnalysisdStatsAgentsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | Daemon stats request time | [optional] 
**Name** | Pointer to **string** | Daemon name | [optional] 
**Agents** | Pointer to [**[]WazuhAnalysisdStatsAgentsItemAgentsInner**](WazuhAnalysisdStatsAgentsItemAgentsInner.md) |  | [optional] 

## Methods

### NewWazuhAnalysisdStatsAgentsItem

`func NewWazuhAnalysisdStatsAgentsItem() *WazuhAnalysisdStatsAgentsItem`

NewWazuhAnalysisdStatsAgentsItem instantiates a new WazuhAnalysisdStatsAgentsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhAnalysisdStatsAgentsItemWithDefaults

`func NewWazuhAnalysisdStatsAgentsItemWithDefaults() *WazuhAnalysisdStatsAgentsItem`

NewWazuhAnalysisdStatsAgentsItemWithDefaults instantiates a new WazuhAnalysisdStatsAgentsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WazuhAnalysisdStatsAgentsItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WazuhAnalysisdStatsAgentsItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WazuhAnalysisdStatsAgentsItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WazuhAnalysisdStatsAgentsItem) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *WazuhAnalysisdStatsAgentsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WazuhAnalysisdStatsAgentsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WazuhAnalysisdStatsAgentsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WazuhAnalysisdStatsAgentsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAgents

`func (o *WazuhAnalysisdStatsAgentsItem) GetAgents() []WazuhAnalysisdStatsAgentsItemAgentsInner`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *WazuhAnalysisdStatsAgentsItem) GetAgentsOk() (*[]WazuhAnalysisdStatsAgentsItemAgentsInner, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *WazuhAnalysisdStatsAgentsItem) SetAgents(v []WazuhAnalysisdStatsAgentsItemAgentsInner)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *WazuhAnalysisdStatsAgentsItem) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


