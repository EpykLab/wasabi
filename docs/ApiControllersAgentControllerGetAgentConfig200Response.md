# ApiControllersAgentControllerGetAgentConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to **map[string]interface{}** | Current agent&#39;s configuration. The output varies with requested component and the agent configuration | [optional] 

## Methods

### NewApiControllersAgentControllerGetAgentConfig200Response

`func NewApiControllersAgentControllerGetAgentConfig200Response() *ApiControllersAgentControllerGetAgentConfig200Response`

NewApiControllersAgentControllerGetAgentConfig200Response instantiates a new ApiControllersAgentControllerGetAgentConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerGetAgentConfig200ResponseWithDefaults

`func NewApiControllersAgentControllerGetAgentConfig200ResponseWithDefaults() *ApiControllersAgentControllerGetAgentConfig200Response`

NewApiControllersAgentControllerGetAgentConfig200ResponseWithDefaults instantiates a new ApiControllersAgentControllerGetAgentConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerGetAgentConfig200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


