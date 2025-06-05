# ApiControllersAgentControllerGetAgentKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseAgentsKeys**](AllItemsResponseAgentsKeys.md) |  | [optional] 

## Methods

### NewApiControllersAgentControllerGetAgentKey200Response

`func NewApiControllersAgentControllerGetAgentKey200Response() *ApiControllersAgentControllerGetAgentKey200Response`

NewApiControllersAgentControllerGetAgentKey200Response instantiates a new ApiControllersAgentControllerGetAgentKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerGetAgentKey200ResponseWithDefaults

`func NewApiControllersAgentControllerGetAgentKey200ResponseWithDefaults() *ApiControllersAgentControllerGetAgentKey200Response`

NewApiControllersAgentControllerGetAgentKey200ResponseWithDefaults instantiates a new ApiControllersAgentControllerGetAgentKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerGetAgentKey200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerGetAgentKey200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerGetAgentKey200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerGetAgentKey200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerGetAgentKey200Response) GetData() AllItemsResponseAgentsKeys`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerGetAgentKey200Response) GetDataOk() (*AllItemsResponseAgentsKeys, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerGetAgentKey200Response) SetData(v AllItemsResponseAgentsKeys)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerGetAgentKey200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


