# CommentReplyApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createCommentReply**](CommentReplyApi.md#createCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply | Reply To A Comment |
| [**deleteCommentReply**](CommentReplyApi.md#deleteCommentReply) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Delete A Comment Reply |
| [**editCommentReply**](CommentReplyApi.md#editCommentReply) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Edit A Comment Reply |
| [**getCommentReplies**](CommentReplyApi.md#getCommentReplies) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/replies | Get Comment Replies |
| [**reportCommentReply**](CommentReplyApi.md#reportCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/report | Report A Comment Reply |


<a name="createCommentReply"></a>
# **createCommentReply**
> CreateCommentReplyResponse createCommentReply(userId, postId, commentId, CreateCommentReplyBody)

Reply To A Comment

    This endpoint enables a client to reply to a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **CreateCommentReplyBody** | [**CreateCommentReplyBody**](../Models/CreateCommentReplyBody.md)|  | |

### Return type

[**CreateCommentReplyResponse**](../Models/CreateCommentReplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteCommentReply"></a>
# **deleteCommentReply**
> DeleteCommentReplyResponse deleteCommentReply(userId, postId, commentId, replyId, postType)

Delete A Comment Reply

    This endpoint enables a client to delete a comment reply

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **replyId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**DeleteCommentReplyResponse**](../Models/DeleteCommentReplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="editCommentReply"></a>
# **editCommentReply**
> EditCommentReplyResponse editCommentReply(userId, postId, commentId, replyId, EditCommentReplyBody)

Edit A Comment Reply

    This endpoint enables a client to edit a comment reply

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **replyId** | **String**|  | [default to null] |
| **EditCommentReplyBody** | [**EditCommentReplyBody**](../Models/EditCommentReplyBody.md)|  | |

### Return type

[**EditCommentReplyResponse**](../Models/EditCommentReplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getCommentReplies"></a>
# **getCommentReplies**
> GetCommentRepliesResponse getCommentReplies(userId, postId, commentId, postType)

Get Comment Replies

    This endpoint enables a client to get comment replies

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**GetCommentRepliesResponse**](../Models/GetCommentRepliesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reportCommentReply"></a>
# **reportCommentReply**
> ReportCommentReplyResponse reportCommentReply(userId, postId, commentId, replyId, ReportCommentReplyBody)

Report A Comment Reply

    This endpoint enables a client to report a comment reply

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **replyId** | **String**|  | [default to null] |
| **ReportCommentReplyBody** | [**ReportCommentReplyBody**](../Models/ReportCommentReplyBody.md)|  | |

### Return type

[**ReportCommentReplyResponse**](../Models/ReportCommentReplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

