# ApiControllersAgentControllerRestartAgent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**ItemAffected**](ItemAffected.md) |  | [optional] 

## Methods

### NewApiControllersAgentControllerRestartAgent200Response

`func NewApiControllersAgentControllerRestartAgent200Response() *ApiControllersAgentControllerRestartAgent200Response`

NewApiControllersAgentControllerRestartAgent200Response instantiates a new ApiControllersAgentControllerRestartAgent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerRestartAgent200ResponseWithDefaults

`func NewApiControllersAgentControllerRestartAgent200ResponseWithDefaults() *ApiControllersAgentControllerRestartAgent200Response`

NewApiControllersAgentControllerRestartAgent200ResponseWithDefaults instantiates a new ApiControllersAgentControllerRestartAgent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerRestartAgent200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerRestartAgent200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerRestartAgent200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerRestartAgent200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerRestartAgent200Response) GetData() ItemAffected`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerRestartAgent200Response) GetDataOk() (*ItemAffected, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerRestartAgent200Response) SetData(v ItemAffected)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerRestartAgent200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


