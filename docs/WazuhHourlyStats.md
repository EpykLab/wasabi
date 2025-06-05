# WazuhHourlyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Averages** | Pointer to **[]int32** | Array containing the number of alerts for every hour | [optional] 
**Interactions** | Pointer to **int32** |  | [optional] 

## Methods

### NewWazuhHourlyStats

`func NewWazuhHourlyStats() *WazuhHourlyStats`

NewWazuhHourlyStats instantiates a new WazuhHourlyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhHourlyStatsWithDefaults

`func NewWazuhHourlyStatsWithDefaults() *WazuhHourlyStats`

NewWazuhHourlyStatsWithDefaults instantiates a new WazuhHourlyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverages

`func (o *WazuhHourlyStats) GetAverages() []int32`

GetAverages returns the Averages field if non-nil, zero value otherwise.

### GetAveragesOk

`func (o *WazuhHourlyStats) GetAveragesOk() (*[]int32, bool)`

GetAveragesOk returns a tuple with the Averages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverages

`func (o *WazuhHourlyStats) SetAverages(v []int32)`

SetAverages sets Averages field to given value.

### HasAverages

`func (o *WazuhHourlyStats) HasAverages() bool`

HasAverages returns a boolean if a field has been set.

### GetInteractions

`func (o *WazuhHourlyStats) GetInteractions() int32`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *WazuhHourlyStats) GetInteractionsOk() (*int32, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *WazuhHourlyStats) SetInteractions(v int32)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *WazuhHourlyStats) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


