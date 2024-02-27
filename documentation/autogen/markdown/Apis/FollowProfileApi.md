# FollowProfileApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptFollowProfile**](FollowProfileApi.md#acceptFollowProfile) | **POST** /social-microservice/api/v1/follow-requests/{followRecordId}/accept | Accepts a user&#39;s follow request |
| [**followProfile**](FollowProfileApi.md#followProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/target/{targetUserId} | follow a user profile |


<a name="acceptFollowProfile"></a>
# **acceptFollowProfile**
> AcceptFollowProfileResponse acceptFollowProfile(followRecordId)

Accepts a user&#39;s follow request

    This endpoint enables a client to accept a follow request from a source a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **followRecordId** | **String**| The id of the follow record | type: uint64 | [default to null] |

### Return type

[**AcceptFollowProfileResponse**](../Models/AcceptFollowProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="followProfile"></a>
# **followProfile**
> FollowProfileResponse followProfile(sourceUserId, targetUserId)

follow a user profile

    This endpoint enables a client to follow a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **sourceUserId** | **String**| the user ID trying to follow another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **targetUserId** | **String**| the user ID being followed by another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**FollowProfileResponse**](../Models/FollowProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

