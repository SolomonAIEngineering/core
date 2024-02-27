# BookmarkApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**bookmarkPost**](BookmarkApi.md#bookmarkPost) | **POST** /social-microservice/api/v1/users/{userId}/post/bookmark/{postId} | Bookmarks a post |
| [**bookmarkPublication**](BookmarkApi.md#bookmarkPublication) | **POST** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Bookmarks a publication |
| [**removeBookmarkedPost**](BookmarkApi.md#removeBookmarkedPost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/bookmark | Deletes A Bookmarked Post |
| [**removeBookmarkedPublication**](BookmarkApi.md#removeBookmarkedPublication) | **DELETE** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Deletes A Bookmarked Publication |


<a name="bookmarkPost"></a>
# **bookmarkPost**
> BookmarkPostResponse bookmarkPost(userId, postId)

Bookmarks a post

    This endpoint enables a client to bookmark a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |

### Return type

[**BookmarkPostResponse**](../Models/BookmarkPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="bookmarkPublication"></a>
# **bookmarkPublication**
> BookmarkPublicationResponse bookmarkPublication(userId, publicationId)

Bookmarks a publication

    This endpoint enables a client to bookmark a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |

### Return type

[**BookmarkPublicationResponse**](../Models/BookmarkPublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removeBookmarkedPost"></a>
# **removeBookmarkedPost**
> RemoveBookmarkedPostResponse removeBookmarkedPost(userId, postId, postType)

Deletes A Bookmarked Post

    This endpoint enables a client to delete a bookmarked post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**RemoveBookmarkedPostResponse**](../Models/RemoveBookmarkedPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removeBookmarkedPublication"></a>
# **removeBookmarkedPublication**
> RemoveBookmarkedPostResponse removeBookmarkedPublication(userId, publicationId)

Deletes A Bookmarked Publication

    This endpoint enables a client to delete a bookmarked publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |

### Return type

[**RemoveBookmarkedPostResponse**](../Models/RemoveBookmarkedPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

