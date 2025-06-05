# \SecurityAPI

All URIs are relative to *https://localhost:55000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiControllersSecurityControllerAddPolicy**](SecurityAPI.md#ApiControllersSecurityControllerAddPolicy) | **Post** /security/policies | Add policy
[**ApiControllersSecurityControllerAddRole**](SecurityAPI.md#ApiControllersSecurityControllerAddRole) | **Post** /security/roles | Add role
[**ApiControllersSecurityControllerAddRule**](SecurityAPI.md#ApiControllersSecurityControllerAddRule) | **Post** /security/rules | Add security rule
[**ApiControllersSecurityControllerCreateUser**](SecurityAPI.md#ApiControllersSecurityControllerCreateUser) | **Post** /security/users | Add user
[**ApiControllersSecurityControllerDeleteSecurityConfig**](SecurityAPI.md#ApiControllersSecurityControllerDeleteSecurityConfig) | **Delete** /security/config | Restore default security config
[**ApiControllersSecurityControllerDeleteUsers**](SecurityAPI.md#ApiControllersSecurityControllerDeleteUsers) | **Delete** /security/users | Delete users
[**ApiControllersSecurityControllerDeprecatedLoginUser**](SecurityAPI.md#ApiControllersSecurityControllerDeprecatedLoginUser) | **Get** /security/user/authenticate | Login
[**ApiControllersSecurityControllerEditRunAs**](SecurityAPI.md#ApiControllersSecurityControllerEditRunAs) | **Put** /security/users/{user_id}/run_as | Enable/Disable run_as
[**ApiControllersSecurityControllerGetPolicies**](SecurityAPI.md#ApiControllersSecurityControllerGetPolicies) | **Get** /security/policies | List policies
[**ApiControllersSecurityControllerGetRbacActions**](SecurityAPI.md#ApiControllersSecurityControllerGetRbacActions) | **Get** /security/actions | List RBAC actions
[**ApiControllersSecurityControllerGetRbacResources**](SecurityAPI.md#ApiControllersSecurityControllerGetRbacResources) | **Get** /security/resources | List RBAC resources
[**ApiControllersSecurityControllerGetRoles**](SecurityAPI.md#ApiControllersSecurityControllerGetRoles) | **Get** /security/roles | List roles
[**ApiControllersSecurityControllerGetRules**](SecurityAPI.md#ApiControllersSecurityControllerGetRules) | **Get** /security/rules | List security rules
[**ApiControllersSecurityControllerGetSecurityConfig**](SecurityAPI.md#ApiControllersSecurityControllerGetSecurityConfig) | **Get** /security/config | Get security config
[**ApiControllersSecurityControllerGetUserMe**](SecurityAPI.md#ApiControllersSecurityControllerGetUserMe) | **Get** /security/users/me | Get current user info
[**ApiControllersSecurityControllerGetUserMePolicies**](SecurityAPI.md#ApiControllersSecurityControllerGetUserMePolicies) | **Get** /security/users/me/policies | Get current user processed policies
[**ApiControllersSecurityControllerGetUsers**](SecurityAPI.md#ApiControllersSecurityControllerGetUsers) | **Get** /security/users | List users
[**ApiControllersSecurityControllerLoginUser**](SecurityAPI.md#ApiControllersSecurityControllerLoginUser) | **Post** /security/user/authenticate | Login
[**ApiControllersSecurityControllerLogoutUser**](SecurityAPI.md#ApiControllersSecurityControllerLogoutUser) | **Delete** /security/user/authenticate | Logout current user
[**ApiControllersSecurityControllerPutSecurityConfig**](SecurityAPI.md#ApiControllersSecurityControllerPutSecurityConfig) | **Put** /security/config | Update security config
[**ApiControllersSecurityControllerRemovePolicies**](SecurityAPI.md#ApiControllersSecurityControllerRemovePolicies) | **Delete** /security/policies | Delete policies
[**ApiControllersSecurityControllerRemoveRolePolicy**](SecurityAPI.md#ApiControllersSecurityControllerRemoveRolePolicy) | **Delete** /security/roles/{role_id}/policies | Remove policies from role
[**ApiControllersSecurityControllerRemoveRoleRule**](SecurityAPI.md#ApiControllersSecurityControllerRemoveRoleRule) | **Delete** /security/roles/{role_id}/rules | Remove security rules from role
[**ApiControllersSecurityControllerRemoveRoles**](SecurityAPI.md#ApiControllersSecurityControllerRemoveRoles) | **Delete** /security/roles | Delete roles
[**ApiControllersSecurityControllerRemoveRules**](SecurityAPI.md#ApiControllersSecurityControllerRemoveRules) | **Delete** /security/rules | Delete security rules
[**ApiControllersSecurityControllerRemoveUserRole**](SecurityAPI.md#ApiControllersSecurityControllerRemoveUserRole) | **Delete** /security/users/{user_id}/roles | Remove roles from user
[**ApiControllersSecurityControllerRevokeAllTokens**](SecurityAPI.md#ApiControllersSecurityControllerRevokeAllTokens) | **Put** /security/user/revoke | Revoke JWT tokens
[**ApiControllersSecurityControllerRunAsLogin**](SecurityAPI.md#ApiControllersSecurityControllerRunAsLogin) | **Post** /security/user/authenticate/run_as | Login auth_context
[**ApiControllersSecurityControllerSetRolePolicy**](SecurityAPI.md#ApiControllersSecurityControllerSetRolePolicy) | **Post** /security/roles/{role_id}/policies | Add policies to role
[**ApiControllersSecurityControllerSetRoleRule**](SecurityAPI.md#ApiControllersSecurityControllerSetRoleRule) | **Post** /security/roles/{role_id}/rules | Add security rules to role
[**ApiControllersSecurityControllerSetUserRole**](SecurityAPI.md#ApiControllersSecurityControllerSetUserRole) | **Post** /security/users/{user_id}/roles | Add roles to user
[**ApiControllersSecurityControllerUpdatePolicy**](SecurityAPI.md#ApiControllersSecurityControllerUpdatePolicy) | **Put** /security/policies/{policy_id} | Update policy
[**ApiControllersSecurityControllerUpdateRole**](SecurityAPI.md#ApiControllersSecurityControllerUpdateRole) | **Put** /security/roles/{role_id} | Update role
[**ApiControllersSecurityControllerUpdateRule**](SecurityAPI.md#ApiControllersSecurityControllerUpdateRule) | **Put** /security/rules/{rule_id} | Update security rule
[**ApiControllersSecurityControllerUpdateUser**](SecurityAPI.md#ApiControllersSecurityControllerUpdateUser) | **Put** /security/users/{user_id} | Update users



## ApiControllersSecurityControllerAddPolicy

> ApiControllersSecurityControllerGetPolicies200Response ApiControllersSecurityControllerAddPolicy(ctx).Pretty(pretty).WaitForComplete(waitForComplete).PoliciesRequest(policiesRequest).Execute()

Add policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	policiesRequest := *openapiclient.NewPoliciesRequest("Name_example", *openapiclient.NewPoliciesRequestPolicy([]string{"Actions_example"}, []string{"Resources_example"}, "Effect_example")) // PoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerAddPolicy(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).PoliciesRequest(policiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerAddPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerAddPolicy`: ApiControllersSecurityControllerGetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerAddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **policiesRequest** | [**PoliciesRequest**](PoliciesRequest.md) |  |

### Return type

[**ApiControllersSecurityControllerGetPolicies200Response**](ApiControllersSecurityControllerGetPolicies200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerAddRole

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerAddRole(ctx).Pretty(pretty).WaitForComplete(waitForComplete).RolesRequest(rolesRequest).Execute()

Add role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	rolesRequest := *openapiclient.NewRolesRequest("Name_example") // RolesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerAddRole(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).RolesRequest(rolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerAddRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerAddRole`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerAddRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerAddRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **rolesRequest** | [**RolesRequest**](RolesRequest.md) |  |

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerAddRule

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerAddRule(ctx).Pretty(pretty).WaitForComplete(waitForComplete).SecurityRulesRequest(securityRulesRequest).Execute()

Add security rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	securityRulesRequest := *openapiclient.NewSecurityRulesRequest("Name_example", map[string]interface{}(123)) // SecurityRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerAddRule(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).SecurityRulesRequest(securityRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerAddRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerAddRule`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerAddRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerAddRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **securityRulesRequest** | [**SecurityRulesRequest**](SecurityRulesRequest.md) |  |

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerCreateUser

> ApiControllersSecurityControllerCreateUser200Response ApiControllersSecurityControllerCreateUser(ctx).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersSecurityControllerCreateUserRequest(apiControllersSecurityControllerCreateUserRequest).Execute()

Add user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	apiControllersSecurityControllerCreateUserRequest := *openapiclient.NewApiControllersSecurityControllerCreateUserRequest("Username_example", "Password_example") // ApiControllersSecurityControllerCreateUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerCreateUser(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersSecurityControllerCreateUserRequest(apiControllersSecurityControllerCreateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerCreateUser`: ApiControllersSecurityControllerCreateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerCreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **apiControllersSecurityControllerCreateUserRequest** | [**ApiControllersSecurityControllerCreateUserRequest**](ApiControllersSecurityControllerCreateUserRequest.md) |  |

### Return type

[**ApiControllersSecurityControllerCreateUser200Response**](ApiControllersSecurityControllerCreateUser200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerDeleteSecurityConfig

> map[string]interface{} ApiControllersSecurityControllerDeleteSecurityConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Restore default security config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerDeleteSecurityConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerDeleteSecurityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerDeleteSecurityConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerDeleteSecurityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerDeleteSecurityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerDeleteUsers

> ApiControllersSecurityControllerDeleteUsers200Response ApiControllersSecurityControllerDeleteUsers(ctx).UserIds(userIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userIds := []string{"Inner_example"} // []string | List of user IDs (separated by comma), use the keyword 'all' to select all users
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerDeleteUsers(context.Background()).UserIds(userIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerDeleteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerDeleteUsers`: ApiControllersSecurityControllerDeleteUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerDeleteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | **[]string** | List of user IDs (separated by comma), use the keyword &#39;all&#39; to select all users |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerDeleteUsers200Response**](ApiControllersSecurityControllerDeleteUsers200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerDeprecatedLoginUser

> ApiControllersSecurityControllerDeprecatedLoginUser200Response ApiControllersSecurityControllerDeprecatedLoginUser(ctx).Raw(raw).Execute()

Login



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	raw := true // bool | Format response in plain text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerDeprecatedLoginUser(context.Background()).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerDeprecatedLoginUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerDeprecatedLoginUser`: ApiControllersSecurityControllerDeprecatedLoginUser200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerDeprecatedLoginUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerDeprecatedLoginUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **raw** | **bool** | Format response in plain text |

### Return type

[**ApiControllersSecurityControllerDeprecatedLoginUser200Response**](ApiControllersSecurityControllerDeprecatedLoginUser200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerEditRunAs

> ApiControllersSecurityControllerEditRunAs200Response ApiControllersSecurityControllerEditRunAs(ctx, userId).AllowRunAs(allowRunAs).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Enable/Disable run_as



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userId := "userId_example" // string | User ID
	allowRunAs := true // bool | Value for the allow_run_as flag (optional) (default to false)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerEditRunAs(context.Background(), userId).AllowRunAs(allowRunAs).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerEditRunAs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerEditRunAs`: ApiControllersSecurityControllerEditRunAs200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerEditRunAs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerEditRunAsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowRunAs** | **bool** | Value for the allow_run_as flag | [default to false]
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerEditRunAs200Response**](ApiControllersSecurityControllerEditRunAs200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetPolicies

> ApiControllersSecurityControllerGetPolicies200Response ApiControllersSecurityControllerGetPolicies(ctx).PolicyIds(policyIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()

List policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	policyIds := []string{"Inner_example"} // []string | List of policy IDs (optional)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetPolicies(context.Background()).PolicyIds(policyIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetPolicies`: ApiControllersSecurityControllerGetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyIds** | **[]string** | List of policy IDs |
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersSecurityControllerGetPolicies200Response**](ApiControllersSecurityControllerGetPolicies200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetRbacActions

> ApiControllersSecurityControllerGetRbacActions200Response ApiControllersSecurityControllerGetRbacActions(ctx).Endpoint(endpoint).Pretty(pretty).Execute()

List RBAC actions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	endpoint := "endpoint_example" // string | Look for the RBAC actions which are related to the specified endpoint (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetRbacActions(context.Background()).Endpoint(endpoint).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetRbacActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetRbacActions`: ApiControllersSecurityControllerGetRbacActions200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetRbacActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetRbacActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **string** | Look for the RBAC actions which are related to the specified endpoint |
 **pretty** | **bool** | Show results in human-readable format | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRbacActions200Response**](ApiControllersSecurityControllerGetRbacActions200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetRbacResources

> ApiControllersSecurityControllerGetRbacActions200Response ApiControllersSecurityControllerGetRbacResources(ctx).Resource(resource).Pretty(pretty).Execute()

List RBAC resources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	resource := "resource_example" // string | List of current RBAC's resources. (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetRbacResources(context.Background()).Resource(resource).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetRbacResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetRbacResources`: ApiControllersSecurityControllerGetRbacActions200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetRbacResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetRbacResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource** | **string** | List of current RBAC&#39;s resources. |
 **pretty** | **bool** | Show results in human-readable format | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRbacActions200Response**](ApiControllersSecurityControllerGetRbacActions200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetRoles

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerGetRoles(ctx).RoleIds(roleIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleIds := []string{"Inner_example"} // []string | List of role IDs (separated by comma) (optional)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetRoles(context.Background()).RoleIds(roleIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetRoles`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleIds** | **[]string** | List of role IDs (separated by comma) |
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetRules

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerGetRules(ctx).RuleIds(ruleIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()

List security rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	ruleIds := []string{"Inner_example"} // []string | List of rule IDs (separated by comma) (optional)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetRules(context.Background()).RuleIds(ruleIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetRules`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleIds** | **[]string** | List of rule IDs (separated by comma) |
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetSecurityConfig

> ApiResponse ApiControllersSecurityControllerGetSecurityConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get security config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetSecurityConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetSecurityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetSecurityConfig`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetSecurityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetSecurityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetUserMe

> ApiControllersSecurityControllerGetUserMe200Response ApiControllersSecurityControllerGetUserMe(ctx).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Get current user info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetUserMe(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetUserMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetUserMe`: ApiControllersSecurityControllerGetUserMe200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetUserMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetUserMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerGetUserMe200Response**](ApiControllersSecurityControllerGetUserMe200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetUserMePolicies

> ApiResponse ApiControllersSecurityControllerGetUserMePolicies(ctx).Pretty(pretty).Execute()

Get current user processed policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetUserMePolicies(context.Background()).Pretty(pretty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetUserMePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetUserMePolicies`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetUserMePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetUserMePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerGetUsers

> ApiControllersSecurityControllerGetUsers200Response ApiControllersSecurityControllerGetUsers(ctx).UserIds(userIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()

List users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userIds := []string{"Inner_example"} // []string | List of user IDs (separated by comma) (optional)
	limit := int32(56) // int32 | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  (optional) (default to 500)
	offset := int32(56) // int32 | First element to return in the collection (optional) (default to 0)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	search := "search_example" // string | Look for elements containing the specified string. To obtain a complementary search, use '-' at the beginning (optional)
	select_ := []string{"Inner_example"} // []string | Select which fields to return (separated by comma). Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	sort := "sort_example" // string | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use '.' for nested fields. For example, '{field1: field2}' may be selected with 'field1.field2' (optional)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	q := "q_example" // string | Query to filter results by. For example q=&quot;status=active&quot; (optional)
	distinct := true // bool | Look for distinct values. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerGetUsers(context.Background()).UserIds(userIds).Limit(limit).Offset(offset).Pretty(pretty).Search(search).Select_(select_).Sort(sort).WaitForComplete(waitForComplete).Q(q).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerGetUsers`: ApiControllersSecurityControllerGetUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | **[]string** | List of user IDs (separated by comma) |
 **limit** | **int32** | Maximum number of elements to return. Although up to 100.000 can be specified, it is recommended not to exceed 500 elements. Responses may be slower the more this number is exceeded.  | [default to 500]
 **offset** | **int32** | First element to return in the collection | [default to 0]
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **search** | **string** | Look for elements containing the specified string. To obtain a complementary search, use &#39;-&#39; at the beginning |
 **select_** | **[]string** | Select which fields to return (separated by comma). Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **sort** | **string** | Sort the collection by a field or fields (separated by comma). Use +/- at the beggining to list in ascending or descending order. Use &#39;.&#39; for nested fields. For example, &#39;{field1: field2}&#39; may be selected with &#39;field1.field2&#39; |
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **q** | **string** | Query to filter results by. For example q&#x3D;&amp;quot;status&#x3D;active&amp;quot; |
 **distinct** | **bool** | Look for distinct values. | [default to false]

### Return type

[**ApiControllersSecurityControllerGetUsers200Response**](ApiControllersSecurityControllerGetUsers200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerLoginUser

> ApiControllersSecurityControllerDeprecatedLoginUser200Response ApiControllersSecurityControllerLoginUser(ctx).Raw(raw).Execute()

Login



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	raw := true // bool | Format response in plain text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerLoginUser(context.Background()).Raw(raw).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerLoginUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerLoginUser`: ApiControllersSecurityControllerDeprecatedLoginUser200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerLoginUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerLoginUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **raw** | **bool** | Format response in plain text |

### Return type

[**ApiControllersSecurityControllerDeprecatedLoginUser200Response**](ApiControllersSecurityControllerDeprecatedLoginUser200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerLogoutUser

> ApiResponse ApiControllersSecurityControllerLogoutUser(ctx).Execute()

Logout current user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerLogoutUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerLogoutUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerLogoutUser`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerLogoutUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerLogoutUserRequest struct via the builder pattern


### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerPutSecurityConfig

> map[string]interface{} ApiControllersSecurityControllerPutSecurityConfig(ctx).Pretty(pretty).WaitForComplete(waitForComplete).SecurityConfiguration(securityConfiguration).Execute()

Update security config



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	securityConfiguration := *openapiclient.NewSecurityConfiguration() // SecurityConfiguration |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerPutSecurityConfig(context.Background()).Pretty(pretty).WaitForComplete(waitForComplete).SecurityConfiguration(securityConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerPutSecurityConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerPutSecurityConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerPutSecurityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerPutSecurityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **securityConfiguration** | [**SecurityConfiguration**](SecurityConfiguration.md) |  |

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemovePolicies

> ApiControllersSecurityControllerGetPolicies200Response ApiControllersSecurityControllerRemovePolicies(ctx).PolicyIds(policyIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	policyIds := []string{"Inner_example"} // []string | List of policy IDs (separated by comma), use the keyword 'all' to select all policies
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemovePolicies(context.Background()).PolicyIds(policyIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemovePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemovePolicies`: ApiControllersSecurityControllerGetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemovePolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemovePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyIds** | **[]string** | List of policy IDs (separated by comma), use the keyword &#39;all&#39; to select all policies |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerGetPolicies200Response**](ApiControllersSecurityControllerGetPolicies200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemoveRolePolicy

> ApiResponse ApiControllersSecurityControllerRemoveRolePolicy(ctx, roleId).PolicyIds(policyIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Remove policies from role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleId := "roleId_example" // string | Specify a role ID
	policyIds := []string{"Inner_example"} // []string | List of policy IDs (separated by comma), use the keyword 'all' to select all policies
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemoveRolePolicy(context.Background(), roleId).PolicyIds(policyIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemoveRolePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemoveRolePolicy`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemoveRolePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Specify a role ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemoveRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyIds** | **[]string** | List of policy IDs (separated by comma), use the keyword &#39;all&#39; to select all policies |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemoveRoleRule

> ApiResponse ApiControllersSecurityControllerRemoveRoleRule(ctx, roleId).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Remove security rules from role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleId := "roleId_example" // string | Specify a role ID
	ruleIds := []string{"Inner_example"} // []string | List of rule IDs (separated by comma), use the keyword 'all' to select all rules
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemoveRoleRule(context.Background(), roleId).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemoveRoleRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemoveRoleRule`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemoveRoleRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Specify a role ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemoveRoleRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleIds** | **[]string** | List of rule IDs (separated by comma), use the keyword &#39;all&#39; to select all rules |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemoveRoles

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerRemoveRoles(ctx).RoleIds(roleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleIds := []string{"Inner_example"} // []string | List of role IDs (separated by comma), use the keyword 'all' to select all roles
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemoveRoles(context.Background()).RoleIds(roleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemoveRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemoveRoles`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemoveRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemoveRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleIds** | **[]string** | List of role IDs (separated by comma), use the keyword &#39;all&#39; to select all roles |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemoveRules

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerRemoveRules(ctx).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Delete security rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	ruleIds := []string{"Inner_example"} // []string | List of rule IDs (separated by comma), use the keyword 'all' to select all rules
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemoveRules(context.Background()).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemoveRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemoveRules`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemoveRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemoveRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ruleIds** | **[]string** | List of rule IDs (separated by comma), use the keyword &#39;all&#39; to select all rules |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRemoveUserRole

> ApiControllersSecurityControllerSetUserRole200Response ApiControllersSecurityControllerRemoveUserRole(ctx, userId).RoleIds(roleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Remove roles from user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userId := "userId_example" // string | User ID
	roleIds := []string{"Inner_example"} // []string | List of role IDs (separated by comma), use the keyword 'all' to select all roles
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRemoveUserRole(context.Background(), userId).RoleIds(roleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRemoveUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRemoveUserRole`: ApiControllersSecurityControllerSetUserRole200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRemoveUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRemoveUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleIds** | **[]string** | List of role IDs (separated by comma), use the keyword &#39;all&#39; to select all roles |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerSetUserRole200Response**](ApiControllersSecurityControllerSetUserRole200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRevokeAllTokens

> map[string]interface{} ApiControllersSecurityControllerRevokeAllTokens(ctx).Execute()

Revoke JWT tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRevokeAllTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRevokeAllTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRevokeAllTokens`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRevokeAllTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRevokeAllTokensRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerRunAsLogin

> ApiControllersSecurityControllerRunAsLogin200Response ApiControllersSecurityControllerRunAsLogin(ctx).Raw(raw).Body(body).Execute()

Login auth_context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	raw := true // bool | Format response in plain text (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerRunAsLogin(context.Background()).Raw(raw).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerRunAsLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerRunAsLogin`: ApiControllersSecurityControllerRunAsLogin200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerRunAsLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerRunAsLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **raw** | **bool** | Format response in plain text |
 **body** | **map[string]interface{}** |  |

### Return type

[**ApiControllersSecurityControllerRunAsLogin200Response**](ApiControllersSecurityControllerRunAsLogin200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerSetRolePolicy

> ApiResponse ApiControllersSecurityControllerSetRolePolicy(ctx, roleId).PolicyIds(policyIds).Position(position).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Add policies to role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleId := "roleId_example" // string | Specify a role ID
	policyIds := []string{"Inner_example"} // []string | List of policy IDs
	position := int32(56) // int32 | Security position for roles/policies (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerSetRolePolicy(context.Background(), roleId).PolicyIds(policyIds).Position(position).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerSetRolePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerSetRolePolicy`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerSetRolePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Specify a role ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerSetRolePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyIds** | **[]string** | List of policy IDs |
 **position** | **int32** | Security position for roles/policies |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerSetRoleRule

> ApiResponse ApiControllersSecurityControllerSetRoleRule(ctx, roleId).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Add security rules to role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleId := "roleId_example" // string | Specify a role ID
	ruleIds := []string{"Inner_example"} // []string | List of rule IDs (separated by comma)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerSetRoleRule(context.Background(), roleId).RuleIds(ruleIds).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerSetRoleRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerSetRoleRule`: ApiResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerSetRoleRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Specify a role ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerSetRoleRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ruleIds** | **[]string** | List of rule IDs (separated by comma) |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerSetUserRole

> ApiControllersSecurityControllerSetUserRole200Response ApiControllersSecurityControllerSetUserRole(ctx, userId).RoleIds(roleIds).Position(position).Pretty(pretty).WaitForComplete(waitForComplete).Execute()

Add roles to user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userId := "userId_example" // string | User ID
	roleIds := []string{"Inner_example"} // []string | List of role IDs (separated by comma)
	position := int32(56) // int32 | Security position for roles/policies (optional)
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerSetUserRole(context.Background(), userId).RoleIds(roleIds).Position(position).Pretty(pretty).WaitForComplete(waitForComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerSetUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerSetUserRole`: ApiControllersSecurityControllerSetUserRole200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerSetUserRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerSetUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleIds** | **[]string** | List of role IDs (separated by comma) |
 **position** | **int32** | Security position for roles/policies |
 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]

### Return type

[**ApiControllersSecurityControllerSetUserRole200Response**](ApiControllersSecurityControllerSetUserRole200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerUpdatePolicy

> ApiControllersSecurityControllerGetPolicies200Response ApiControllersSecurityControllerUpdatePolicy(ctx, policyId).Pretty(pretty).WaitForComplete(waitForComplete).PoliciesRequestNoRequired(policiesRequestNoRequired).Execute()

Update policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	policyId := "policyId_example" // string | Specify a policy id
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	policiesRequestNoRequired := *openapiclient.NewPoliciesRequestNoRequired() // PoliciesRequestNoRequired |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerUpdatePolicy(context.Background(), policyId).Pretty(pretty).WaitForComplete(waitForComplete).PoliciesRequestNoRequired(policiesRequestNoRequired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerUpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerUpdatePolicy`: ApiControllersSecurityControllerGetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerUpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | Specify a policy id |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **policiesRequestNoRequired** | [**PoliciesRequestNoRequired**](PoliciesRequestNoRequired.md) |  |

### Return type

[**ApiControllersSecurityControllerGetPolicies200Response**](ApiControllersSecurityControllerGetPolicies200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerUpdateRole

> ApiControllersSecurityControllerGetRoles200Response ApiControllersSecurityControllerUpdateRole(ctx, roleId).Pretty(pretty).WaitForComplete(waitForComplete).RolesRequestNoRequired(rolesRequestNoRequired).Execute()

Update role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	roleId := "roleId_example" // string | Specify a role ID
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	rolesRequestNoRequired := *openapiclient.NewRolesRequestNoRequired() // RolesRequestNoRequired |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerUpdateRole(context.Background(), roleId).Pretty(pretty).WaitForComplete(waitForComplete).RolesRequestNoRequired(rolesRequestNoRequired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerUpdateRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerUpdateRole`: ApiControllersSecurityControllerGetRoles200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerUpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | Specify a role ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **rolesRequestNoRequired** | [**RolesRequestNoRequired**](RolesRequestNoRequired.md) |  |

### Return type

[**ApiControllersSecurityControllerGetRoles200Response**](ApiControllersSecurityControllerGetRoles200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerUpdateRule

> ApiControllersSecurityControllerUpdateRule200Response ApiControllersSecurityControllerUpdateRule(ctx, ruleId).Pretty(pretty).WaitForComplete(waitForComplete).SecurityRulesRequestNoRequired(securityRulesRequestNoRequired).Execute()

Update security rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	ruleId := "ruleId_example" // string | Specify a rule ID
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	securityRulesRequestNoRequired := *openapiclient.NewSecurityRulesRequestNoRequired() // SecurityRulesRequestNoRequired |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerUpdateRule(context.Background(), ruleId).Pretty(pretty).WaitForComplete(waitForComplete).SecurityRulesRequestNoRequired(securityRulesRequestNoRequired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerUpdateRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerUpdateRule`: ApiControllersSecurityControllerUpdateRule200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerUpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | Specify a rule ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **securityRulesRequestNoRequired** | [**SecurityRulesRequestNoRequired**](SecurityRulesRequestNoRequired.md) |  |

### Return type

[**ApiControllersSecurityControllerUpdateRule200Response**](ApiControllersSecurityControllerUpdateRule200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiControllersSecurityControllerUpdateUser

> ApiControllersSecurityControllerUpdateUser200Response ApiControllersSecurityControllerUpdateUser(ctx, userId).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersSecurityControllerUpdateUserRequest(apiControllersSecurityControllerUpdateUserRequest).Execute()

Update users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/EpykLab/wasabi"
)

func main() {
	userId := "userId_example" // string | User ID
	pretty := true // bool | Show results in human-readable format (optional) (default to false)
	waitForComplete := true // bool | Disable timeout response (optional) (default to false)
	apiControllersSecurityControllerUpdateUserRequest := *openapiclient.NewApiControllersSecurityControllerUpdateUserRequest() // ApiControllersSecurityControllerUpdateUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiControllersSecurityControllerUpdateUser(context.Background(), userId).Pretty(pretty).WaitForComplete(waitForComplete).ApiControllersSecurityControllerUpdateUserRequest(apiControllersSecurityControllerUpdateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiControllersSecurityControllerUpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiControllersSecurityControllerUpdateUser`: ApiControllersSecurityControllerUpdateUser200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiControllersSecurityControllerUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID |

### Other Parameters

Other parameters are passed through a pointer to a apiApiControllersSecurityControllerUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **bool** | Show results in human-readable format | [default to false]
 **waitForComplete** | **bool** | Disable timeout response | [default to false]
 **apiControllersSecurityControllerUpdateUserRequest** | [**ApiControllersSecurityControllerUpdateUserRequest**](ApiControllersSecurityControllerUpdateUserRequest.md) |  |

### Return type

[**ApiControllersSecurityControllerUpdateUser200Response**](ApiControllersSecurityControllerUpdateUser200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
