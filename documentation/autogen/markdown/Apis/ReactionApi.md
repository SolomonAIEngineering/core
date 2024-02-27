# ReactionApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**reactToComment**](ReactionApi.md#reactToComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment |
| [**reactToCommentReply**](ReactionApi.md#reactToCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment reply |
| [**reactToPost**](ReactionApi.md#reactToPost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/account-type/{accountType}/reaction/{reaction} | Reacts to a post |


<a name="reactToComment"></a>
# **reactToComment**
> ReactToCommentResponse reactToComment(userId, postId, commentId, accountType, reaction, postType)

Reacts to a comment

    This endpoint enables a client to react to a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **accountType** | **String**|  | [default to null] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **reaction** | **String**|  | [default to null] [enum: REACTION_UNSPECIFIED, REACTION_LIKE, REACTION_LOVE, REACTION_HAHA, REACTION_WOW, REACTION_SAD, REACTION_ANGRY, REACTION_DISLIKE] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**ReactToCommentResponse**](../Models/ReactToCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reactToCommentReply"></a>
# **reactToCommentReply**
> ReactToCommentReplyResponse reactToCommentReply(userId, postId, commentId, replyId, accountType, reaction, postType)

Reacts to a comment reply

    This endpoint enables a client to react to a comment reply

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **replyId** | **String**|  | [default to null] |
| **accountType** | **String**|  | [default to null] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **reaction** | **String**|  | [default to null] [enum: REACTION_UNSPECIFIED, REACTION_LIKE, REACTION_LOVE, REACTION_HAHA, REACTION_WOW, REACTION_SAD, REACTION_ANGRY, REACTION_DISLIKE] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**ReactToCommentReplyResponse**](../Models/ReactToCommentReplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reactToPost"></a>
# **reactToPost**
> ReactToPostResponse reactToPost(userId, postId, accountType, reaction, postType)

Reacts to a post

    This endpoint enables a client to react to a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **accountType** | **String**|  | [default to null] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **reaction** | **String**|  | [default to null] [enum: REACTION_UNSPECIFIED, REACTION_LIKE, REACTION_LOVE, REACTION_HAHA, REACTION_WOW, REACTION_SAD, REACTION_ANGRY, REACTION_DISLIKE] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**ReactToPostResponse**](../Models/ReactToPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

