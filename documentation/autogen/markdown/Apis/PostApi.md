# PostApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addPostQualityScore**](PostApi.md#addPostQualityScore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/quality | Adds a quality score to a post |
| [**createPost**](PostApi.md#createPost) | **POST** /social-microservice/api/v1/users/{userId}/post | Create a post |
| [**deletePost**](PostApi.md#deletePost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Delete a post |
| [**editPost**](PostApi.md#editPost) | **PUT** /social-microservice/api/v1/post/{postId}/type/{postType} | Edits a post by id |
| [**getBlogPostsByTag**](PostApi.md#getBlogPostsByTag) | **GET** /social-microservice/api/v1/posts/blog/tag/{tag} | Get blog posts by tag |
| [**getCommunityBlogPosts**](PostApi.md#getCommunityBlogPosts) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/posts/blog | Get community blog posts |
| [**getPost**](PostApi.md#getPost) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Get a post |
| [**getPostsByCategory**](PostApi.md#getPostsByCategory) | **GET** /social-microservice/api/v1/user/{userId}/category/{category}/posts/{postType}/limit/{limit}/offset/{offset} | Get all posts associated with a category |
| [**getPostsByTopic**](PostApi.md#getPostsByTopic) | **GET** /social-microservice/api/v1/community/{communityProfileId}/topic/{topicName}/posts | Get all posts associated with a topic |
| [**reportPost**](PostApi.md#reportPost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType}/report | Report a post |


<a name="addPostQualityScore"></a>
# **addPostQualityScore**
> AddPostQualityScoreResponse addPostQualityScore(userId, postId, postType, body)

Adds a quality score to a post

    This endpoint enables a client add a quality score to a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **postType** | **String**| The type of post being reacted to | type: string | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **body** | **String**|  | |

### Return type

[**AddPostQualityScoreResponse**](../Models/AddPostQualityScoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createPost"></a>
# **createPost**
> CreatePostResponse createPost(userId, accountType, Post, communityProfileId)

Create a post

    This endpoint enables a client to create a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to create a post (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **accountType** | **String**| The type of profile making the request | type: string | [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **Post** | [**Post**](../Models/Post.md)| The post payload | type: json_object | |
| **communityProfileId** | **String**| The ID of the community profile attempting to create the post  | type: uint64 | [optional] [default to null] |

### Return type

[**CreatePostResponse**](../Models/CreatePostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deletePost"></a>
# **deletePost**
> DeletePostResponse deletePost(userId, postId, postType)

Delete a post

    This endpoint enables a client to delete a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete a post (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **postId** | **String**| The ID of the post attempted to be delete | type: string | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**DeletePostResponse**](../Models/DeletePostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="editPost"></a>
# **editPost**
> EditPostResponse editPost(postId, postType, Post)

Edits a post by id

    This endpoint enables a client to edit a post by id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **postId** | **String**| The ID of the post to be updated | type: string | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **Post** | [**Post**](../Models/Post.md)| The post payload | type: json_object | |

### Return type

[**EditPostResponse**](../Models/EditPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getBlogPostsByTag"></a>
# **getBlogPostsByTag**
> GetBlogPostsByTagResponse getBlogPostsByTag(tag, postType)

Get blog posts by tag

    This endpoint enables a client to query a set of blog posts tied to a tag

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tag** | **String**|  | [default to null] |
| **postType** | **String**| The type of post being reacted to | type: string | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**GetBlogPostsByTagResponse**](../Models/GetBlogPostsByTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCommunityBlogPosts"></a>
# **getCommunityBlogPosts**
> GetCommunityBlogPostsResponse getCommunityBlogPosts(communityProfileId)

Get community blog posts

    This endpoint enables a client to get community blog posts

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **communityProfileId** | **String**|  | [default to null] |

### Return type

[**GetCommunityBlogPostsResponse**](../Models/GetCommunityBlogPostsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPost"></a>
# **getPost**
> GetPostResponse getPost(userId, postId, postType)

Get a post

    This endpoint enables a client to get a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to obtain a post (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **postId** | **String**| The ID of the post to obtain | type: string | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**GetPostResponse**](../Models/GetPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPostsByCategory"></a>
# **getPostsByCategory**
> GetPostsByCategoryResponse getPostsByCategory(userId, category, postType, limit, offset)

Get all posts associated with a category

    This endpoint enables a client to get all posts tied to a category

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **category** | **String**|  | [default to null] [enum: CATEGORY_UNSPECIFIED, CATEGORY_WORLD, CATEGORY_BUSINESS, CATEGORY_ECONOMICS, CATEGORY_FOREIGN_POLICY, CATEGORY_POLITICS, CATEGORY_TECHNOLOGY, CATEGORY_OTHER] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **limit** | **String**|  | [default to null] |
| **offset** | **String**|  | [default to null] |

### Return type

[**GetPostsByCategoryResponse**](../Models/GetPostsByCategoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPostsByTopic"></a>
# **getPostsByTopic**
> GetPostsByTopicResponse getPostsByTopic(communityProfileId, topicName)

Get all posts associated with a topic

    This endpoint enables a client to get all posts tied to a topic

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **communityProfileId** | **String**| the community profile to associate the topic to | type: uint64 | [default to null] |
| **topicName** | **String**| the name of a given topic | type: string | [default to null] |

### Return type

[**GetPostsByTopicResponse**](../Models/GetPostsByTopicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="reportPost"></a>
# **reportPost**
> ReportPostResponse reportPost(userId, postId, postType)

Report a post

    This endpoint enables a client to report a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**ReportPostResponse**](../Models/ReportPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

