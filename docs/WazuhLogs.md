# WazuhLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Log message | [optional] 
**Level** | Pointer to **string** | Log level | [optional] 
**Tag** | Pointer to **string** | Wazuh component that logged the event | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWazuhLogs

`func NewWazuhLogs() *WazuhLogs`

NewWazuhLogs instantiates a new WazuhLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWazuhLogsWithDefaults

`func NewWazuhLogsWithDefaults() *WazuhLogs`

NewWazuhLogsWithDefaults instantiates a new WazuhLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WazuhLogs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WazuhLogs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WazuhLogs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WazuhLogs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLevel

`func (o *WazuhLogs) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WazuhLogs) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WazuhLogs) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *WazuhLogs) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTag

`func (o *WazuhLogs) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *WazuhLogs) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *WazuhLogs) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *WazuhLogs) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTimestamp

`func (o *WazuhLogs) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WazuhLogs) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WazuhLogs) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WazuhLogs) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


