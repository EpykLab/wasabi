# ApiControllersAgentControllerGetAgents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseAgents**](AllItemsResponseAgents.md) |  | [optional] 

## Methods

### NewApiControllersAgentControllerGetAgents200Response

`func NewApiControllersAgentControllerGetAgents200Response() *ApiControllersAgentControllerGetAgents200Response`

NewApiControllersAgentControllerGetAgents200Response instantiates a new ApiControllersAgentControllerGetAgents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerGetAgents200ResponseWithDefaults

`func NewApiControllersAgentControllerGetAgents200ResponseWithDefaults() *ApiControllersAgentControllerGetAgents200Response`

NewApiControllersAgentControllerGetAgents200ResponseWithDefaults instantiates a new ApiControllersAgentControllerGetAgents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerGetAgents200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerGetAgents200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerGetAgents200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerGetAgents200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerGetAgents200Response) GetData() AllItemsResponseAgents`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerGetAgents200Response) GetDataOk() (*AllItemsResponseAgents, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerGetAgents200Response) SetData(v AllItemsResponseAgents)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerGetAgents200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


