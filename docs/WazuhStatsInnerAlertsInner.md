# WazuhStatsInnerAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sigid** | Pointer to **int32** | Rule ID that matched the event | [optional] 
**Level** | Pointer to **int32** | Alert level | [optional] 
**Times** | Pointer to **int32** | Number of times the alert was raised during the specified hour | [optional] 

## Methods

### NewWazuhStatsInnerAlertsInner

`func NewWazuhStatsInnerAlertsInner() *WazuhStatsInnerAlertsInner`

NewWazuhStatsInnerAlertsInner instantiates a new WazuhStatsInnerAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhStatsInnerAlertsInnerWithDefaults

`func NewWazuhStatsInnerAlertsInnerWithDefaults() *WazuhStatsInnerAlertsInner`

NewWazuhStatsInnerAlertsInnerWithDefaults instantiates a new WazuhStatsInnerAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigid

`func (o *WazuhStatsInnerAlertsInner) GetSigid() int32`

GetSigid returns the Sigid field if non-nil, zero value otherwise.

### GetSigidOk

`func (o *WazuhStatsInnerAlertsInner) GetSigidOk() (*int32, bool)`

GetSigidOk returns a tuple with the Sigid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigid

`func (o *WazuhStatsInnerAlertsInner) SetSigid(v int32)`

SetSigid sets Sigid field to given value.

### HasSigid

`func (o *WazuhStatsInnerAlertsInner) HasSigid() bool`

HasSigid returns a boolean if a field has been set.

### GetLevel

`func (o *WazuhStatsInnerAlertsInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WazuhStatsInnerAlertsInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WazuhStatsInnerAlertsInner) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *WazuhStatsInnerAlertsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTimes

`func (o *WazuhStatsInnerAlertsInner) GetTimes() int32`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *WazuhStatsInnerAlertsInner) GetTimesOk() (*int32, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *WazuhStatsInnerAlertsInner) SetTimes(v int32)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *WazuhStatsInnerAlertsInner) HasTimes() bool`

HasTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


