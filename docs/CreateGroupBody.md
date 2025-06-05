# CreateGroupBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Group name. It can contain any of the characters between a-z, A-Z, 0-9, &#39;_&#39;, &#39;-&#39; and &#39;.&#39;. Names &#39;.&#39; and &#39;..&#39; are restricted. | 

## Methods

### NewCreateGroupBody

`func NewCreateGroupBody(groupId string, ) *CreateGroupBody`

NewCreateGroupBody instantiates a new CreateGroupBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupBodyWithDefaults

`func NewCreateGroupBodyWithDefaults() *CreateGroupBody`

NewCreateGroupBodyWithDefaults instantiates a new CreateGroupBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CreateGroupBody) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CreateGroupBody) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CreateGroupBody) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


