# ApiControllersAgentControllerGetAgentFields200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseAgentsDistinct**](AllItemsResponseAgentsDistinct.md) |  | [optional] 

## Methods

### NewApiControllersAgentControllerGetAgentFields200Response

`func NewApiControllersAgentControllerGetAgentFields200Response() *ApiControllersAgentControllerGetAgentFields200Response`

NewApiControllersAgentControllerGetAgentFields200Response instantiates a new ApiControllersAgentControllerGetAgentFields200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersAgentControllerGetAgentFields200ResponseWithDefaults

`func NewApiControllersAgentControllerGetAgentFields200ResponseWithDefaults() *ApiControllersAgentControllerGetAgentFields200Response`

NewApiControllersAgentControllerGetAgentFields200ResponseWithDefaults instantiates a new ApiControllersAgentControllerGetAgentFields200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersAgentControllerGetAgentFields200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersAgentControllerGetAgentFields200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersAgentControllerGetAgentFields200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersAgentControllerGetAgentFields200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersAgentControllerGetAgentFields200Response) GetData() AllItemsResponseAgentsDistinct`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersAgentControllerGetAgentFields200Response) GetDataOk() (*AllItemsResponseAgentsDistinct, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersAgentControllerGetAgentFields200Response) SetData(v AllItemsResponseAgentsDistinct)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersAgentControllerGetAgentFields200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


