# FollowCommunityProfileApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**followCommunityProfile**](FollowCommunityProfileApi.md#followCommunityProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/community-profiles/{targetCommunityProfileId} | Follows A Community Profile |


<a name="followCommunityProfile"></a>
# **followCommunityProfile**
> FollowCommunityProfileResponse followCommunityProfile(sourceUserId, targetCommunityProfileId)

Follows A Community Profile

    This endpoint enables a client to follow a community profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **sourceUserId** | **String**| the user ID trying to follow another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **targetCommunityProfileId** | **String**| the targetCommunityProfileID ID being followed by another user | [default to null] |

### Return type

[**FollowCommunityProfileResponse**](../Models/FollowCommunityProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

