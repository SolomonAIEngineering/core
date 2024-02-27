# FollowApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getAccountsFollowing**](FollowApi.md#getAccountsFollowing) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/account-type/{accountType}/following | Get Communities and users you are following |
| [**getFollowers**](FollowApi.md#getFollowers) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/followers | Get Users Following you |
| [**getPendingFollows**](FollowApi.md#getPendingFollows) | **GET** /social-microservice/api/v1/users/{userId}/follow/pending-requests | Get Pending Follow Requests |


<a name="getAccountsFollowing"></a>
# **getAccountsFollowing**
> GetAccountsFollowingResponse getAccountsFollowing(userId, profileId, accountType, limit)

Get Communities and users you are following

    This endpoint enables a client to get all the accounts a given user follows

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **profileId** | **String**|  | [default to null] |
| **accountType** | **String**| the account type of the user whoses followers are being requested | [default to null] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **limit** | **String**|  | [default to null] |

### Return type

[**GetAccountsFollowingResponse**](../Models/GetAccountsFollowingResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getFollowers"></a>
# **getFollowers**
> GetFollowersResponse getFollowers(userId, profileId, limit)

Get Users Following you

    This endpoint enables a client to get all the followers following a given user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **profileId** | **String**|  | [default to null] |
| **limit** | **String**|  | [default to null] |

### Return type

[**GetFollowersResponse**](../Models/GetFollowersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPendingFollows"></a>
# **getPendingFollows**
> GetPendingFollowsResponse getPendingFollows(userId)

Get Pending Follow Requests

    This endpoint enables a client to get all the pending follow requests

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID who&#39;s pending request we want to obtain (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**GetPendingFollowsResponse**](../Models/GetPendingFollowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

