# ApiControllersSecurityControllerGetUserMe200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseUsers**](AllItemsResponseUsers.md) |  | [optional] 

## Methods

### NewApiControllersSecurityControllerGetUserMe200Response

`func NewApiControllersSecurityControllerGetUserMe200Response() *ApiControllersSecurityControllerGetUserMe200Response`

NewApiControllersSecurityControllerGetUserMe200Response instantiates a new ApiControllersSecurityControllerGetUserMe200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersSecurityControllerGetUserMe200ResponseWithDefaults

`func NewApiControllersSecurityControllerGetUserMe200ResponseWithDefaults() *ApiControllersSecurityControllerGetUserMe200Response`

NewApiControllersSecurityControllerGetUserMe200ResponseWithDefaults instantiates a new ApiControllersSecurityControllerGetUserMe200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersSecurityControllerGetUserMe200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersSecurityControllerGetUserMe200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersSecurityControllerGetUserMe200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersSecurityControllerGetUserMe200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersSecurityControllerGetUserMe200Response) GetData() AllItemsResponseUsers`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersSecurityControllerGetUserMe200Response) GetDataOk() (*AllItemsResponseUsers, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersSecurityControllerGetUserMe200Response) SetData(v AllItemsResponseUsers)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersSecurityControllerGetUserMe200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


