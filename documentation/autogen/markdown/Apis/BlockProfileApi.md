# BlockProfileApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**blockUserProfile**](BlockProfileApi.md#blockUserProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/block/target/{targetUserId} | blocks a user profile |


<a name="blockUserProfile"></a>
# **blockUserProfile**
> BlockUserProfileResponse blockUserProfile(sourceUserId, targetUserId)

blocks a user profile

    This endpoint enables a client to block a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **sourceUserId** | **String**| the user ID trying to block another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **targetUserId** | **String**| the user ID being blocked by another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**BlockUserProfileResponse**](../Models/BlockUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

