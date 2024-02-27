# ThreadApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addPostToThread**](ThreadApi.md#addPostToThread) | **POST** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType} | Adds A Post To A Thread |
| [**getPostThread**](ThreadApi.md#getPostThread) | **GET** /social-microservice/api/v1/users/{userId}/post/thread/{postId} | Gets A Post&#39;s Thread |
| [**removePostFromThread**](ThreadApi.md#removePostFromThread) | **DELETE** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType}/target/{participantPostId} | Deletes A Post From A Thread |


<a name="addPostToThread"></a>
# **addPostToThread**
> AddPostToThreadResponse addPostToThread(userId, parentPostId, postType, Post)

Adds A Post To A Thread

    This endpoint enables a client to add a post to a thread

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **parentPostId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **Post** | [**Post**](../Models/Post.md)|  | |

### Return type

[**AddPostToThreadResponse**](../Models/AddPostToThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getPostThread"></a>
# **getPostThread**
> GetPostThreadResponse getPostThread(userId, postId, postType)

Gets A Post&#39;s Thread

    This endpoint enables a client to query a post&#39;s thread

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**GetPostThreadResponse**](../Models/GetPostThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="removePostFromThread"></a>
# **removePostFromThread**
> RemovePostFromThreadResponse removePostFromThread(userId, parentPostId, postType, participantPostId)

Deletes A Post From A Thread

    This endpoint enables a client to delete a post from a thread

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **parentPostId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **participantPostId** | **String**|  | [default to null] |

### Return type

[**RemovePostFromThreadResponse**](../Models/RemovePostFromThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

