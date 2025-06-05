# WazuhStatsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]WazuhStatsInnerAlertsInner**](WazuhStatsInnerAlertsInner.md) |  | [optional] 
**Events** | Pointer to **int32** | Number of events processed during the specified hour | [optional] 
**Firewall** | Pointer to **int32** | Number of firewall alerts raised during the specified hour | [optional] 
**Hour** | Pointer to **int32** | Hour of the day in 24h format | [optional] 
**Syscheck** | Pointer to **int32** | Number of syscheck alerts raised during the specified hour | [optional] 
**TotalAlerts** | Pointer to **int32** | Number of alerts raised during the specified hour | [optional] 

## Methods

### NewWazuhStatsInner

`func NewWazuhStatsInner() *WazuhStatsInner`

NewWazuhStatsInner instantiates a new WazuhStatsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhStatsInnerWithDefaults

`func NewWazuhStatsInnerWithDefaults() *WazuhStatsInner`

NewWazuhStatsInnerWithDefaults instantiates a new WazuhStatsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *WazuhStatsInner) GetAlerts() []WazuhStatsInnerAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *WazuhStatsInner) GetAlertsOk() (*[]WazuhStatsInnerAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *WazuhStatsInner) SetAlerts(v []WazuhStatsInnerAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *WazuhStatsInner) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetEvents

`func (o *WazuhStatsInner) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WazuhStatsInner) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WazuhStatsInner) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WazuhStatsInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFirewall

`func (o *WazuhStatsInner) GetFirewall() int32`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *WazuhStatsInner) GetFirewallOk() (*int32, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *WazuhStatsInner) SetFirewall(v int32)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *WazuhStatsInner) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetHour

`func (o *WazuhStatsInner) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *WazuhStatsInner) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *WazuhStatsInner) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *WazuhStatsInner) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetSyscheck

`func (o *WazuhStatsInner) GetSyscheck() int32`

GetSyscheck returns the Syscheck field if non-nil, zero value otherwise.

### GetSyscheckOk

`func (o *WazuhStatsInner) GetSyscheckOk() (*int32, bool)`

GetSyscheckOk returns a tuple with the Syscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscheck

`func (o *WazuhStatsInner) SetSyscheck(v int32)`

SetSyscheck sets Syscheck field to given value.

### HasSyscheck

`func (o *WazuhStatsInner) HasSyscheck() bool`

HasSyscheck returns a boolean if a field has been set.

### GetTotalAlerts

`func (o *WazuhStatsInner) GetTotalAlerts() int32`

GetTotalAlerts returns the TotalAlerts field if non-nil, zero value otherwise.

### GetTotalAlertsOk

`func (o *WazuhStatsInner) GetTotalAlertsOk() (*int32, bool)`

GetTotalAlertsOk returns a tuple with the TotalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAlerts

`func (o *WazuhStatsInner) SetTotalAlerts(v int32)`

SetTotalAlerts sets TotalAlerts field to given value.

### HasTotalAlerts

`func (o *WazuhStatsInner) HasTotalAlerts() bool`

HasTotalAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


