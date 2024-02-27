# UserApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**editUserProfile**](UserApi.md#editUserProfile) | **PUT** /api/v1/users/{userId} | update a user profile |


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

