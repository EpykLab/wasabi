# ApiControllersSecurityControllerGetUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human readable description to explain the result of the request | [optional] 
**Data** | Pointer to [**AllItemsResponseUsers**](AllItemsResponseUsers.md) |  | [optional] 

## Methods

### NewApiControllersSecurityControllerGetUsers200Response

`func NewApiControllersSecurityControllerGetUsers200Response() *ApiControllersSecurityControllerGetUsers200Response`

NewApiControllersSecurityControllerGetUsers200Response instantiates a new ApiControllersSecurityControllerGetUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiControllersSecurityControllerGetUsers200ResponseWithDefaults

`func NewApiControllersSecurityControllerGetUsers200ResponseWithDefaults() *ApiControllersSecurityControllerGetUsers200Response`

NewApiControllersSecurityControllerGetUsers200ResponseWithDefaults instantiates a new ApiControllersSecurityControllerGetUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiControllersSecurityControllerGetUsers200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiControllersSecurityControllerGetUsers200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiControllersSecurityControllerGetUsers200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiControllersSecurityControllerGetUsers200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *ApiControllersSecurityControllerGetUsers200Response) GetData() AllItemsResponseUsers`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiControllersSecurityControllerGetUsers200Response) GetDataOk() (*AllItemsResponseUsers, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiControllersSecurityControllerGetUsers200Response) SetData(v AllItemsResponseUsers)`

SetData sets Data field to given value.

### HasData

`func (o *ApiControllersSecurityControllerGetUsers200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


