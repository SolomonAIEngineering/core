# UserProfileApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createUserProfile**](UserProfileApi.md#createUserProfile) | **POST** /social-microservice/api/v1/users | creates a user profile |
| [**deleteUserProfile**](UserProfileApi.md#deleteUserProfile) | **DELETE** /social-microservice/api/v1/users/{userId} | deletes a user profile |
| [**discoverProfiles**](UserProfileApi.md#discoverProfiles) | **GET** /social-microservice/api/v1/users/{userId}/discover/limit/{limit} | Discover Profiles |
| [**editUserProfile**](UserProfileApi.md#editUserProfile) | **PUT** /social-microservice/api/v1/users/{userId} | update a user profile |
| [**getCannyUserSSOToken**](UserProfileApi.md#getCannyUserSSOToken) | **GET** /social-microservice/api/v1/user/{userId}/canny/email/{email} | Retrieves user sso token for canny |
| [**getUserProfile**](UserProfileApi.md#getUserProfile) | **GET** /social-microservice/api/v1/users/{userId} | gets a user profile |
| [**getUserProfiles**](UserProfileApi.md#getUserProfiles) | **GET** /social-microservice/api/v1/users/page-size/{pageSize}/page-number/{pageNumber} | Gets a set of user profiles |


<a name="createUserProfile"></a>
# **createUserProfile**
> CreateUserProfileResponse createUserProfile(CreateUserProfileRequest)

creates a user profile

    This endpoint enables a client to create a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateUserProfileRequest** | [**CreateUserProfileRequest**](../Models/CreateUserProfileRequest.md)|  | |

### Return type

[**CreateUserProfileResponse**](../Models/CreateUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteUserProfile"></a>
# **deleteUserProfile**
> DeleteUserProfileResponse deleteUserProfile(userId)

deletes a user profile

    This endpoint enables a client to delete a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**DeleteUserProfileResponse**](../Models/DeleteUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="discoverProfiles"></a>
# **discoverProfiles**
> DiscoverProfilesResponse discoverProfiles(userId, limit)

Discover Profiles

    This endpoint enables a client to discover a set of profiles he/she does not follow

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **limit** | **String**|  | [default to null] |

### Return type

[**DiscoverProfilesResponse**](../Models/DiscoverProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="editUserProfile"></a>
# **editUserProfile**
> EditUserProfileResponse editUserProfile(userId, UserProfile)

update a user profile

    This endpoint performs an updates operation on a user profile based on the provided parametersThis update operation can span multiple services on specific cases (such as when the client is explicitly attempting to update the email of the user)All update operations are atomic by nature hence we should not expect any form of divergent state

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to update this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **UserProfile** | [**UserProfile**](../Models/UserProfile.md)| the profile payload | type: json_object | |

### Return type

[**EditUserProfileResponse**](../Models/EditUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getCannyUserSSOToken"></a>
# **getCannyUserSSOToken**
> GetCannyUserSSOTokenResponse getCannyUserSSOToken(userId, email)

Retrieves user sso token for canny

    Fetches a user sso token for canny

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **email** | **String**|  | [default to null] |

### Return type

[**GetCannyUserSSOTokenResponse**](../Models/GetCannyUserSSOTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserProfile"></a>
# **getUserProfile**
> GetUserProfileResponse getUserProfile(userId, requestorProfileId, requestorProfileType)

gets a user profile

    This endpoint performs a query against the social service to obtain a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user ID associated with the profile we want to get | type: uint64 | [default to null] |
| **requestorProfileId** | **String**| The RequestorProfileID is an optional parameter used to check if the profileID (requestor) making a request for the record actually follows the record | [optional] [default to null] |
| **requestorProfileType** | **String**| The RequestorProfileType is an optional parameter which tells us what type of profile is the requestor | [optional] [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |

### Return type

[**GetUserProfileResponse**](../Models/GetUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserProfiles"></a>
# **getUserProfiles**
> GetUserProfilesResponse getUserProfiles(pageSize, pageNumber)

Gets a set of user profiles

    This endpoint enables a client to get a number of user profiles in a paginated manner

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pageSize** | **Integer**|  | [default to null] |
| **pageNumber** | **Integer**|  | [default to null] |

### Return type

[**GetUserProfilesResponse**](../Models/GetUserProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

