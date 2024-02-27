# UserServiceV2Api

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**checkEmailExistsV2**](UserServiceV2Api.md#checkEmailExistsV2) | **GET** /user-microservice/api/v2/user-service/user/email/{email}/exists | Checks that an email exists or not |
| [**checkUsernameExistsV2**](UserServiceV2Api.md#checkUsernameExistsV2) | **GET** /user-microservice/api/v2/user-service/user/username/{username}/exists | Checks that a username exists or not |
| [**createRole**](UserServiceV2Api.md#createRole) | **POST** /user-microservice/api/v2/user-service/user/role | Creates a new role |
| [**createUserV2**](UserServiceV2Api.md#createUserV2) | **POST** /user-microservice/api/v2/user-service/user | create a user account |
| [**deleteRole**](UserServiceV2Api.md#deleteRole) | **DELETE** /user-microservice/api/v2/user-service/user/role/{id} | Deletes a role |
| [**deleteUserV2**](UserServiceV2Api.md#deleteUserV2) | **DELETE** /user-microservice/api/v2/user-service/user/{userId} | deletes a user account |
| [**getCannyUserSSOToken**](UserServiceV2Api.md#getCannyUserSSOToken) | **GET** /user-microservice/api/v2/user-service/user/canny/{userId} | Retrieves user sso token for canny |
| [**getRole**](UserServiceV2Api.md#getRole) | **GET** /user-microservice/api/v2/user-service/user/role/{id} | Retrieves a role |
| [**getUserByAuth0ID**](UserServiceV2Api.md#getUserByAuth0ID) | **GET** /user-microservice/api/v2/user-service/user/auth-zero/{auth0UserId} | Retrieve user account details by auth0 id and profile type |
| [**getUserByAuthnIDV2**](UserServiceV2Api.md#getUserByAuthnIDV2) | **GET** /user-microservice/api/v2/user-service/user/authn/{authnId} | Retrieve user account details by authn id |
| [**getUserByEmailOrUsernameV2**](UserServiceV2Api.md#getUserByEmailOrUsernameV2) | **GET** /user-microservice/api/v2/user-service/user/account/query-by-email-or-username | Retrieve user account by email or username |
| [**getUserByEmailV2**](UserServiceV2Api.md#getUserByEmailV2) | **GET** /user-microservice/api/v2/user-service/user/email/{email} | Retrieve user details by email |
| [**getUserByUsernameV2**](UserServiceV2Api.md#getUserByUsernameV2) | **GET** /user-microservice/api/v2/user-service/user/username/{username} | Retrieve user details by username |
| [**getUserIdV2**](UserServiceV2Api.md#getUserIdV2) | **GET** /user-microservice/api/v2/user-service/user/query-id | get a user account id |
| [**getUserV2**](UserServiceV2Api.md#getUserV2) | **GET** /user-microservice/api/v2/user-service/user/{userId} | Retrieve user account details |
| [**listRoles**](UserServiceV2Api.md#listRoles) | **GET** /user-microservice/api/v2/user-service/user/roles | Lists all roles |
| [**passwordResetWebhookV2**](UserServiceV2Api.md#passwordResetWebhookV2) | **POST** /user-microservice/api/v2/user-service/user/webhook/password-reset | Webhook for Password Reset |
| [**retrieveBusinessSettings**](UserServiceV2Api.md#retrieveBusinessSettings) | **GET** /user-microservice/api/v2/user-service/user/business/settings/{userId} | Retrieve Business Account Settings |
| [**updateRole**](UserServiceV2Api.md#updateRole) | **PATCH** /user-microservice/api/v2/user/role | Updates an existing role |
| [**updateUserV2**](UserServiceV2Api.md#updateUserV2) | **PUT** /user-microservice/api/v2/user-service/user | update a user account |
| [**verifyUserV2**](UserServiceV2Api.md#verifyUserV2) | **POST** /user-microservice/api/v2/user-service/user/verification/{userId}/profile-type/{profileType} | User Email Verification |


<a name="checkEmailExistsV2"></a>
# **checkEmailExistsV2**
> CheckEmailExistsV2Response checkEmailExistsV2(email, profileType)

Checks that an email exists or not

    Checks if an email exists or not

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**|  | [default to null] |
| **profileType** | **String**| the profile type of the given user we hope to query | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**CheckEmailExistsV2Response**](../Models/CheckEmailExistsV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="checkUsernameExistsV2"></a>
# **checkUsernameExistsV2**
> CheckUsernameExistsV2Response checkUsernameExistsV2(username, profileType)

Checks that a username exists or not

    Checks if a username exists or not

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The username of the user of interest | [default to null] |
| **profileType** | **String**| the profile type of the given user we hope to query | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**CheckUsernameExistsV2Response**](../Models/CheckUsernameExistsV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createRole"></a>
# **createRole**
> CreateRoleResponse createRole(Role)

Creates a new role

    This endpoint adds a new role to the system. It requires role details such as name, type, and permissions.The creation process involves adding the role to the database and initializing its permissions.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Role** | [**Role**](../Models/Role.md)|  | |

### Return type

[**CreateRoleResponse**](../Models/CreateRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createUserV2"></a>
# **createUserV2**
> CreateUserV2Response createUserV2(CreateUserV2Request)

create a user account

    This endpoint performs an a creation operation of a user account based on the provided parametersThis operation is implemented as a distributed transactions as this operation spans multiple services

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateUserV2Request** | [**CreateUserV2Request**](../Models/CreateUserV2Request.md)|  | |

### Return type

[**CreateUserV2Response**](../Models/CreateUserV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteRole"></a>
# **deleteRole**
> DeleteRoleResponse deleteRole(id)

Deletes a role

    This endpoint deletes a role from the system based on the provided role ID.The deletion process ensures that all related data and permissions are properly removed.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**|  | [default to null] |

### Return type

[**DeleteRoleResponse**](../Models/DeleteRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteUserV2"></a>
# **deleteUserV2**
> DeleteUserV2Response deleteUserV2(userId, profileType)

deletes a user account

    This endpoint performs a delete operation on a user account based on the provided parametersThis deletion operation spans multiple services as user details are stored across a suite of our backend servicesThe operation itself is an atomic one of nature. Either all services successfully delete the recod or we fail the requestDivergent state is not expected to be encountered with this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user to delete Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**| the profile type of the given user we hope to query | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**DeleteUserV2Response**](../Models/DeleteUserV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCannyUserSSOToken"></a>
# **getCannyUserSSOToken**
> GetCannyUserSSOTokenResponse getCannyUserSSOToken(userId, profileType)

Retrieves user sso token for canny

    Fetches a user sso token for canny

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user | [default to null] |
| **profileType** | **String**| Indicates the profile type to be queried. For example: \&quot;username:testuser\&quot; | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetCannyUserSSOTokenResponse**](../Models/GetCannyUserSSOTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRole"></a>
# **getRole**
> GetRoleResponse getRole(id)

Retrieves a role

    This endpoint fetches details of a specific role using the role ID.It retrieves the role&#39;s name, type, permissions, and audit history.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**|  | [default to null] |

### Return type

[**GetRoleResponse**](../Models/GetRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByAuth0ID"></a>
# **getUserByAuth0ID**
> GetUserByAuth0IDResponse getUserByAuth0ID(auth0UserId, profileType)

Retrieve user account details by auth0 id and profile type

    Fetches detailed information about a user account based on the specified authn ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **auth0UserId** | **String**| The account ID associated with the user | [default to null] |
| **profileType** | **String**| Indicates the profile type to be queried. For example: \&quot;username:testuser\&quot; | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserByAuth0IDResponse**](../Models/GetUserByAuth0IDResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByAuthnIDV2"></a>
# **getUserByAuthnIDV2**
> GetUserByAuthnIDV2Response getUserByAuthnIDV2(authnId, profileType)

Retrieve user account details by authn id

    Fetches detailed information about a user account based on the specified authn ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authnId** | **String**| Specifies the user&#39;s authn account ID.  Validations: - Must be greater than 0. | [default to null] |
| **profileType** | **String**| Indicates the profile type to be queried. For example: \&quot;username:testuser\&quot; | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserByAuthnIDV2Response**](../Models/GetUserByAuthnIDV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByEmailOrUsernameV2"></a>
# **getUserByEmailOrUsernameV2**
> GetUserByEmailOrUsernameV2Response getUserByEmailOrUsernameV2(profileType, email, username)

Retrieve user account by email or username

    Fetches a user account using the provided email or username.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **profileType** | **String**|  | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |
| **email** | **String**| The email of the user of interest | [optional] [default to null] |
| **username** | **String**| The username of the user of interest | [optional] [default to null] |

### Return type

[**GetUserByEmailOrUsernameV2Response**](../Models/GetUserByEmailOrUsernameV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByEmailV2"></a>
# **getUserByEmailV2**
> GetUserByEmailV2Response getUserByEmailV2(email, profileType)

Retrieve user details by email

    Provides detailed information of a user based on the given email.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**|  | [default to null] |
| **profileType** | **String**| the profile type of the given user we hope to query | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserByEmailV2Response**](../Models/GetUserByEmailV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByUsernameV2"></a>
# **getUserByUsernameV2**
> GetUserByUsernameV2Response getUserByUsernameV2(username, profileType)

Retrieve user details by username

    Provides detailed information of a user based on the given username.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The username of the user of interest | [default to null] |
| **profileType** | **String**| the profile type of the given user we hope to query | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserByUsernameV2Response**](../Models/GetUserByUsernameV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserIdV2"></a>
# **getUserIdV2**
> GetUserIdV2Response getUserIdV2(profileType, email, username)

get a user account id

    This endpoint returns the user record id if the user record exists example: /api/v1/user-service?email&#x3D;testuser@gmail.com&amp;&amp;username&#x3D;testuser

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **profileType** | **String**|  | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |
| **email** | **String**| The email of the user of interest | [optional] [default to null] |
| **username** | **String**| The username of the user of interest | [optional] [default to null] |

### Return type

[**GetUserIdV2Response**](../Models/GetUserIdV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserV2"></a>
# **getUserV2**
> GetUserV2Response getUserV2(userId, profileType)

Retrieve user account details

    Fetches detailed information about a user account based on the specified user ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| Specifies the user&#39;s account ID.  Validations: - Must be greater than 0. | [default to null] |
| **profileType** | **String**| Indicates the profile type to be queried. For example: \&quot;username:testuser\&quot; | [default to PROFILE_TYPE_UNSPECIFIED] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserV2Response**](../Models/GetUserV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRoles"></a>
# **listRoles**
> ListRolesResponse listRoles(page, pageSize)

Lists all roles

    This endpoint retrieves a list of all roles in the system. It supports pagination to handle large sets of data.Each role in the list includes details like name, type, and permissions.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **page** | **Integer**|  | [optional] [default to null] |
| **pageSize** | **Integer**|  | [optional] [default to null] |

### Return type

[**ListRolesResponse**](../Models/ListRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="passwordResetWebhookV2"></a>
# **passwordResetWebhookV2**
> PasswordResetWebhookV2Response passwordResetWebhookV2(accountId, token)

Webhook for Password Reset

    Handles password reset operations for a specified user account through a webhook.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**|  | [optional] [default to null] |
| **token** | **String**|  | [optional] [default to null] |

### Return type

[**PasswordResetWebhookV2Response**](../Models/PasswordResetWebhookV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="retrieveBusinessSettings"></a>
# **retrieveBusinessSettings**
> GetBusinessSettingsResponse retrieveBusinessSettings(userId)

Retrieve Business Account Settings

    Fetches settings associated with a specified business account using the user ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user (business) | [default to null] |

### Return type

[**GetBusinessSettingsResponse**](../Models/GetBusinessSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateRole"></a>
# **updateRole**
> UpdateRoleResponse updateRole(Role)

Updates an existing role

    This endpoint updates the details of an existing role. The role ID is used to identify the role to be updated.The update operation can modify the role&#39;s name, type, and permissions.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Role** | [**Role**](../Models/Role.md)|  | |

### Return type

[**UpdateRoleResponse**](../Models/UpdateRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserV2"></a>
# **updateUserV2**
> UpdateUserV2Response updateUserV2(UpdateUserV2Request)

update a user account

    This endpoint performs an updates operation on a user account based on the provided parametersThis update operation can span multiple services on specific cases (such as when the client is explicitly attempting to update the email of the user)All update operations are atomic by nature hence we should not expect any form of divergent state

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateUserV2Request** | [**UpdateUserV2Request**](../Models/UpdateUserV2Request.md)|  | |

### Return type

[**UpdateUserV2Response**](../Models/UpdateUserV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verifyUserV2"></a>
# **verifyUserV2**
> VerifyUserV2Response verifyUserV2(userId, profileType)

User Email Verification

    Performs verification of a user account&#39;s email based on the provided user ID and profile type.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user | [default to null] |
| **profileType** | **String**|  | [default to null] [enum: PROFILE_TYPE_UNSPECIFIED, PROFILE_TYPE_USER, PROFILE_TYPE_BUSINESS] |

### Return type

[**VerifyUserV2Response**](../Models/VerifyUserV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

