# ApiControllersDecoderControllerGetDecoders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseDecoders**](AllItemsResponseDecoders.md) |  | [optional] 

## Methods

### NewApiControllersDecoderControllerGetDecoders200Response

`func NewApiControllersDecoderControllerGetDecoders200Response() *ApiControllersDecoderControllerGetDecoders200Response`

NewApiControllersDecoderControllerGetDecoders200Response instantiates a new ApiControllersDecoderControllerGetDecoders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersDecoderControllerGetDecoders200ResponseWithDefaults

`func NewApiControllersDecoderControllerGetDecoders200ResponseWithDefaults() *ApiControllersDecoderControllerGetDecoders200Response`

NewApiControllersDecoderControllerGetDecoders200ResponseWithDefaults instantiates a new ApiControllersDecoderControllerGetDecoders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersDecoderControllerGetDecoders200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersDecoderControllerGetDecoders200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersDecoderControllerGetDecoders200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersDecoderControllerGetDecoders200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersDecoderControllerGetDecoders200Response) GetData() AllItemsResponseDecoders`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersDecoderControllerGetDecoders200Response) GetDataOk() (*AllItemsResponseDecoders, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersDecoderControllerGetDecoders200Response) SetData(v AllItemsResponseDecoders)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersDecoderControllerGetDecoders200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


