# PostsApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getPosts**](PostsApi.md#getPosts) | **GET** /social-microservice/api/v1/users/{userId}/posts | Get all the posts of a given user |


<a name="getPosts"></a>
# **getPosts**
> GetPostsResponse getPosts(userId)

Get all the posts of a given user

    This endpoint enables a client to query all posts tied to a given user id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**GetPostsResponse**](../Models/GetPostsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

