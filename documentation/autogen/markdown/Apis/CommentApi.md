# CommentApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addCommentQualityScore**](CommentApi.md#addCommentQualityScore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/quality | Add Comment Quality Score |
| [**createComment**](CommentApi.md#createComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment | Create A Commnet |
| [**deleteComment**](CommentApi.md#deleteComment) | **DELETE** /social-microservice/api/v1/post/{postId}/comment/{commentId} | Delete A Comment |
| [**reportComment**](CommentApi.md#reportComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/report | Report A Comment |


<a name="addCommentQualityScore"></a>
# **addCommentQualityScore**
> AddCommentQualityScoreResponse addCommentQualityScore(userId, postId, commentId, postType, body)

Add Comment Quality Score

    This endpoint enables a client to add a quality score to a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **postType** | **String**| The type of post being reacted to | type: string | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **body** | **String**|  | |

### Return type

[**AddCommentQualityScoreResponse**](../Models/AddCommentQualityScoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createComment"></a>
# **createComment**
> CreateCommentResponse createComment(userId, postId, accountType, postType, Comment, communityProfileId)

Create A Commnet

    This endpoint enables a client to create a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to create a comment (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **postId** | **String**| The ID of the post to whom to add the comment | type: string | [default to null] |
| **accountType** | **String**| The type of account making the request to create a comment | type: string | [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **postType** | **String**| The type of post being reacted to | type: string | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **Comment** | [**Comment**](../Models/Comment.md)| The actual comment payload | type: json_object | |
| **communityProfileId** | **String**| The ID of the community trying to create a comment | type: uint64 | [optional] [default to null] |

### Return type

[**CreateCommentResponse**](../Models/CreateCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteComment"></a>
# **deleteComment**
> CreateCommentResponse deleteComment(postId, commentId, postType)

Delete A Comment

    This endpoint enables a client to delete a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **postId** | **String**| The ID of the post trying to be deleted | type: string | [default to null] |
| **commentId** | **String**| The ID of the comment trying to be delete | type: string | [default to null] |
| **postType** | **String**| The type of post being reacted to | type: string | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**CreateCommentResponse**](../Models/CreateCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reportComment"></a>
# **reportComment**
> ReportCommentResponse reportComment(userId, postId, commentId, ReportCommentBody)

Report A Comment

    This endpoint enables a client to report a comment

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **commentId** | **String**|  | [default to null] |
| **ReportCommentBody** | [**ReportCommentBody**](../Models/ReportCommentBody.md)|  | |

### Return type

[**ReportCommentResponse**](../Models/ReportCommentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

