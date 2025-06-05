# ActiveResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **[]string** | Command arguments | [optional] 
**Command** | **string** | Command running in the agent. If this value starts with &#x60;!&#x60;, then it refers to a script name instead of a command name | 
**Alert** | Pointer to [**ActiveResponseBodyAlert**](ActiveResponseBodyAlert.md) |  | [optional] 

## Methods

### NewActiveResponseBody

`func NewActiveResponseBody(command string, ) *ActiveResponseBody`

NewActiveResponseBody instantiates a new ActiveResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveResponseBodyWithDefaults

`func NewActiveResponseBodyWithDefaults() *ActiveResponseBody`

NewActiveResponseBodyWithDefaults instantiates a new ActiveResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *ActiveResponseBody) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ActiveResponseBody) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ActiveResponseBody) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ActiveResponseBody) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCommand

`func (o *ActiveResponseBody) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ActiveResponseBody) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ActiveResponseBody) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetAlert

`func (o *ActiveResponseBody) GetAlert() ActiveResponseBodyAlert`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *ActiveResponseBody) GetAlertOk() (*ActiveResponseBodyAlert, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *ActiveResponseBody) SetAlert(v ActiveResponseBodyAlert)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *ActiveResponseBody) HasAlert() bool`

HasAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


