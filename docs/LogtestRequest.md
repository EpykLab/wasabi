# LogtestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | Token for the logtest session | [optional] 
**LogFormat** | **string** | Allowed values: syslog, json, snort-full, squid, eventlog, eventchannel, audit, mysql_log, postgresql_log, nmapg, iis, command, full_command, djb-multilog, multi-line | 
**Location** | **string** | Path string | 
**Event** | **string** | Event to look for | 

## Methods

### NewLogtestRequest

`func NewLogtestRequest(logFormat string, location string, event string, ) *LogtestRequest`

NewLogtestRequest instantiates a new LogtestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogtestRequestWithDefaults

`func NewLogtestRequestWithDefaults() *LogtestRequest`

NewLogtestRequestWithDefaults instantiates a new LogtestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LogtestRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogtestRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogtestRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogtestRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLogFormat

`func (o *LogtestRequest) GetLogFormat() string`

GetLogFormat returns the LogFormat field if non-nil, zero value otherwise.

### GetLogFormatOk

`func (o *LogtestRequest) GetLogFormatOk() (*string, bool)`

GetLogFormatOk returns a tuple with the LogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFormat

`func (o *LogtestRequest) SetLogFormat(v string)`

SetLogFormat sets LogFormat field to given value.


### GetLocation

`func (o *LogtestRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LogtestRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LogtestRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEvent

`func (o *LogtestRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *LogtestRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *LogtestRequest) SetEvent(v string)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


