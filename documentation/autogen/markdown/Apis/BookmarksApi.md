# BookmarksApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getBookmarkedPosts**](BookmarksApi.md#getBookmarkedPosts) | **GET** /social-microservice/api/v1/users/bookmarks/{userId} | Get Bookmarked Posts |


<a name="getBookmarkedPosts"></a>
# **getBookmarkedPosts**
> GetBookmarkedPostsResponse getBookmarkedPosts(userId)

Get Bookmarked Posts

    This endpoint enables a client to get all bookmarked posts of a given user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |

### Return type

[**GetBookmarkedPostsResponse**](../Models/GetBookmarkedPostsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

