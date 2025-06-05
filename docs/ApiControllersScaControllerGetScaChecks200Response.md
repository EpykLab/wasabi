# ApiControllersScaControllerGetScaChecks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseSCAChecks**](AllItemsResponseSCAChecks.md) |  | [optional] 

## Methods

### NewApiControllersScaControllerGetScaChecks200Response

`func NewApiControllersScaControllerGetScaChecks200Response() *ApiControllersScaControllerGetScaChecks200Response`

NewApiControllersScaControllerGetScaChecks200Response instantiates a new ApiControllersScaControllerGetScaChecks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersScaControllerGetScaChecks200ResponseWithDefaults

`func NewApiControllersScaControllerGetScaChecks200ResponseWithDefaults() *ApiControllersScaControllerGetScaChecks200Response`

NewApiControllersScaControllerGetScaChecks200ResponseWithDefaults instantiates a new ApiControllersScaControllerGetScaChecks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersScaControllerGetScaChecks200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersScaControllerGetScaChecks200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersScaControllerGetScaChecks200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersScaControllerGetScaChecks200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersScaControllerGetScaChecks200Response) GetData() AllItemsResponseSCAChecks`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersScaControllerGetScaChecks200Response) GetDataOk() (*AllItemsResponseSCAChecks, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersScaControllerGetScaChecks200Response) SetData(v AllItemsResponseSCAChecks)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersScaControllerGetScaChecks200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


