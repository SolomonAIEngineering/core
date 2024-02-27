# UserServiceV1Api

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**checkEmailExists**](UserServiceV1Api.md#checkEmailExists) | **GET** /user-microservice/api/v1/user-service/user/email/{email}/exists | Checks that an email exists or not |
| [**checkUsernameExists**](UserServiceV1Api.md#checkUsernameExists) | **GET** /user-microservice/api/v1/user-service/user/username/{username}/exists | Checks that a username exists or not |
| [**deleteUser**](UserServiceV1Api.md#deleteUser) | **DELETE** /user-microservice/api/v1/user-service/user/{userId} | deletes a user account |
| [**getUser**](UserServiceV1Api.md#getUser) | **GET** /user-microservice/api/v1/user-service/user/{userId} | Gets a user account |
| [**getUserByEmail**](UserServiceV1Api.md#getUserByEmail) | **GET** /user-microservice/api/v1/user-service/user/email/{email} | Gets a user account by email |
| [**getUserByEmailOrUsername**](UserServiceV1Api.md#getUserByEmailOrUsername) | **GET** /user-microservice/api/v1/user-service/user/query-account-by-email-or-username | gets a user account by either email or username |
| [**getUserByUsername**](UserServiceV1Api.md#getUserByUsername) | **GET** /user-microservice/api/v1/user-service/user/username/{username} | Gets a user account by user name |
| [**getUserId**](UserServiceV1Api.md#getUserId) | **GET** /user-microservice/api/v1/user-service/user/query-id | get a user account id |
| [**healthCheck**](UserServiceV1Api.md#healthCheck) | **GET** /user-microservice/api/v1/user-service/user/health | health check |
| [**passwordReset**](UserServiceV1Api.md#passwordReset) | **POST** /user-microservice/api/v1/user-service/user/webhook/password-reset | password reset |
| [**readynessCheck**](UserServiceV1Api.md#readynessCheck) | **GET** /user-microservice/api/v1/user-service/user/ready | readyness check |
| [**updateUser**](UserServiceV1Api.md#updateUser) | **PUT** /user-microservice/api/v1/user-service/user | update a user account |
| [**verification**](UserServiceV1Api.md#verification) | **POST** /user-microservice/api/v1/user-service/user/verification/{userId} | user verification |


<a name="checkEmailExists"></a>
# **checkEmailExists**
> CheckEmailExistsResponse checkEmailExists(email)

Checks that an email exists or not

    Checks if an email exists or not

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**| The email of the user of interest | [default to null] |

### Return type

[**CheckEmailExistsResponse**](../Models/CheckEmailExistsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="checkUsernameExists"></a>
# **checkUsernameExists**
> CheckUsernameExistsResponse checkUsernameExists(username)

Checks that a username exists or not

    Checks if a username exists or not

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The username of the user of interest | [default to null] |

### Return type

[**CheckUsernameExistsResponse**](../Models/CheckUsernameExistsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteUser"></a>
# **deleteUser**
> DeleteUserResponse deleteUser(userId)

deletes a user account

    This endpoint performs a delete operation on a user account based on the provided parametersThis deletion operation spans multiple services as user details are stored across a suite of our backend servicesThe operation itself is an atomic one of nature. Either all services successfully delete the recod or we fail the requestDivergent state is not expected to be encountered with this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user to delete Validations: - user_id must be greater than 0 | [default to null] |

### Return type

[**DeleteUserResponse**](../Models/DeleteUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUser"></a>
# **getUser**
> GetUserResponse getUser(userId)

Gets a user account

    Queries and obtains a user account based on the provided parameters

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user Validations: - user_id must be greater than 0 | [default to null] |

### Return type

[**GetUserResponse**](../Models/GetUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByEmail"></a>
# **getUserByEmail**
> GetUserByEmailResponse getUserByEmail(email)

Gets a user account by email

    Queries and obtains a user account based on the email

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**| The email of the user of interest | [default to null] |

### Return type

[**GetUserByEmailResponse**](../Models/GetUserByEmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByEmailOrUsername"></a>
# **getUserByEmailOrUsername**
> GetUserByEmailOrUsernameResponse getUserByEmailOrUsername(email, username)

gets a user account by either email or username

    This endpoint returns a user account by either provided email or username

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**| The email of the user of interest | [optional] [default to null] |
| **username** | **String**| The username of the user of interest | [optional] [default to null] |

### Return type

[**GetUserByEmailOrUsernameResponse**](../Models/GetUserByEmailOrUsernameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserByUsername"></a>
# **getUserByUsername**
> GetUserByUsernameResponse getUserByUsername(username)

Gets a user account by user name

    Queries and obtains a user account based on the username

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **username** | **String**| The username of the user of interest | [default to null] |

### Return type

[**GetUserByUsernameResponse**](../Models/GetUserByUsernameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserId"></a>
# **getUserId**
> GetUserIdResponse getUserId(email, username)

get a user account id

    This endpoint returns the user record id if the user record exists example: /api/v1/user-service?email&#x3D;testuser@gmail.com&amp;&amp;username&#x3D;testuser

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**| The email of the user of interest | [optional] [default to null] |
| **username** | **String**| The username of the user of interest | [optional] [default to null] |

### Return type

[**GetUserIdResponse**](../Models/GetUserIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="healthCheck"></a>
# **healthCheck**
> HealthCheckResponse healthCheck()

health check

    This endpoint performs a healc check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheckResponse**](../Models/HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="passwordReset"></a>
# **passwordReset**
> PasswordResetWebhookResponse passwordReset(accountId, token)

password reset

    This endpoint performs password reset for a given user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**|  | [optional] [default to null] |
| **token** | **String**|  | [optional] [default to null] |

### Return type

[**PasswordResetWebhookResponse**](../Models/PasswordResetWebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readynessCheck"></a>
# **readynessCheck**
> ReadynessCheckResponse readynessCheck()

readyness check

    This endpoint performs a readiness check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**ReadynessCheckResponse**](../Models/ReadynessCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateUser"></a>
# **updateUser**
> UpdateUserResponse updateUser(UpdateUserRequest)

update a user account

    This endpoint performs an updates operation on a user account based on the provided parametersThis update operation can span multiple services on specific cases (such as when the client is explicitly attempting to update the email of the user)All update operations are atomic by nature hence we should not expect any form of divergent state

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateUserRequest** | [**UpdateUserRequest**](../Models/UpdateUserRequest.md)|  | |

### Return type

[**UpdateUserResponse**](../Models/UpdateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verification"></a>
# **verification**
> VerifyUserResponse verification(userId)

user verification

    This endpoint performs verification of a user account email

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user | [default to null] |

### Return type

[**VerifyUserResponse**](../Models/VerifyUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

