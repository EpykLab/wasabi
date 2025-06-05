# ApiControllersAgentControllerRestartAgents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseAgentIDs**](AllItemsResponseAgentIDs.md) |  | [optional] 

## Methods

### NewApiControllersAgentControllerRestartAgents200Response

`func NewApiControllersAgentControllerRestartAgents200Response() *ApiControllersAgentControllerRestartAgents200Response`

NewApiControllersAgentControllerRestartAgents200Response instantiates a new ApiControllersAgentControllerRestartAgents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerRestartAgents200ResponseWithDefaults

`func NewApiControllersAgentControllerRestartAgents200ResponseWithDefaults() *ApiControllersAgentControllerRestartAgents200Response`

NewApiControllersAgentControllerRestartAgents200ResponseWithDefaults instantiates a new ApiControllersAgentControllerRestartAgents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerRestartAgents200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerRestartAgents200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerRestartAgents200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerRestartAgents200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerRestartAgents200Response) GetData() AllItemsResponseAgentIDs`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerRestartAgents200Response) GetDataOk() (*AllItemsResponseAgentIDs, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerRestartAgents200Response) SetData(v AllItemsResponseAgentIDs)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerRestartAgents200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


