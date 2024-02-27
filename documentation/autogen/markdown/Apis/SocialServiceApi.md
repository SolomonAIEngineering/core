# SocialServiceApi

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptFollowProfile**](SocialServiceApi.md#acceptFollowProfile) | **POST** /social-microservice/api/v1/follow-requests/{followRecordId}/accept | Accepts a user&#39;s follow request |
| [**addCommentQualityScore**](SocialServiceApi.md#addCommentQualityScore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/quality | Add Comment Quality Score |
| [**addPostQualityScore**](SocialServiceApi.md#addPostQualityScore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/quality | Adds a quality score to a post |
| [**addPostToPublication**](SocialServiceApi.md#addPostToPublication) | **POST** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId} | Add a post to a publication |
| [**addPostToThread**](SocialServiceApi.md#addPostToThread) | **POST** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType} | Adds A Post To A Thread |
| [**addPublicationEditor**](SocialServiceApi.md#addPublicationEditor) | **PUT** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Adds an editor to a publication |
| [**blockUserProfile**](SocialServiceApi.md#blockUserProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/block/target/{targetUserId} | blocks a user profile |
| [**bookmarkPost**](SocialServiceApi.md#bookmarkPost) | **POST** /social-microservice/api/v1/users/{userId}/post/bookmark/{postId} | Bookmarks a post |
| [**bookmarkPublication**](SocialServiceApi.md#bookmarkPublication) | **POST** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Bookmarks a publication |
| [**createComment**](SocialServiceApi.md#createComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment | Create A Commnet |
| [**createCommentReply**](SocialServiceApi.md#createCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply | Reply To A Comment |
| [**createCommunityProfile**](SocialServiceApi.md#createCommunityProfile) | **POST** /social-microservice/api/v1/community-profiles/{userId} | Create a community Profile |
| [**createNote**](SocialServiceApi.md#createNote) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/note | Creates and associates a note to a given post |
| [**createPoll**](SocialServiceApi.md#createPoll) | **POST** /social-microservice/api/v1/users/{userId}/poll | Create a poll |
| [**createPost**](SocialServiceApi.md#createPost) | **POST** /social-microservice/api/v1/users/{userId}/post | Create a post |
| [**createPublication**](SocialServiceApi.md#createPublication) | **POST** /social-microservice/api/v1/users/{userId}/publication | Creates a publication |
| [**createTopic**](SocialServiceApi.md#createTopic) | **POST** /social-microservice/api/v1/users/{userId}/community/{communityProfileId}/topic | Create A Topic |
| [**createUserProfile**](SocialServiceApi.md#createUserProfile) | **POST** /social-microservice/api/v1/users | creates a user profile |
| [**deleteComment**](SocialServiceApi.md#deleteComment) | **DELETE** /social-microservice/api/v1/post/{postId}/comment/{commentId} | Delete A Comment |
| [**deleteCommentReply**](SocialServiceApi.md#deleteCommentReply) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Delete A Comment Reply |
| [**deleteCommunityProfile**](SocialServiceApi.md#deleteCommunityProfile) | **DELETE** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Delete Community Profile |
| [**deleteNote**](SocialServiceApi.md#deleteNote) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Delete a note |
| [**deletePoll**](SocialServiceApi.md#deletePoll) | **DELETE** /social-microservice/api/v1/users/{userId}/poll/{postId} | Delete a poll |
| [**deletePost**](SocialServiceApi.md#deletePost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Delete a post |
| [**deletePostFromPublication**](SocialServiceApi.md#deletePostFromPublication) | **DELETE** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId}/post/{postId} | Deletes a post from a publication |
| [**deletePublication**](SocialServiceApi.md#deletePublication) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId} | Deletes a publication |
| [**deletePublicationEditor**](SocialServiceApi.md#deletePublicationEditor) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Deletes an editor to a publication |
| [**deleteUserProfile**](SocialServiceApi.md#deleteUserProfile) | **DELETE** /social-microservice/api/v1/users/{userId} | deletes a user profile |
| [**discoverProfiles**](SocialServiceApi.md#discoverProfiles) | **GET** /social-microservice/api/v1/users/{userId}/discover/limit/{limit} | Discover Profiles |
| [**editCommentReply**](SocialServiceApi.md#editCommentReply) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Edit A Comment Reply |
| [**editCommunityProfile**](SocialServiceApi.md#editCommunityProfile) | **PUT** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Edit a community Profile |
| [**editNote**](SocialServiceApi.md#editNote) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Update a note |
| [**editPost**](SocialServiceApi.md#editPost) | **PUT** /social-microservice/api/v1/post/{postId}/type/{postType} | Edits a post by id |
| [**editUserProfile**](SocialServiceApi.md#editUserProfile) | **PUT** /social-microservice/api/v1/users/{userId} | update a user profile |
| [**followCommunityProfile**](SocialServiceApi.md#followCommunityProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/community-profiles/{targetCommunityProfileId} | Follows A Community Profile |
| [**followProfile**](SocialServiceApi.md#followProfile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/target/{targetUserId} | follow a user profile |
| [**getAccountsFollowing**](SocialServiceApi.md#getAccountsFollowing) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/account-type/{accountType}/following | Get Communities and users you are following |
| [**getBlogPostsByTag**](SocialServiceApi.md#getBlogPostsByTag) | **GET** /social-microservice/api/v1/posts/blog/tag/{tag} | Get blog posts by tag |
| [**getBookmarkedPosts**](SocialServiceApi.md#getBookmarkedPosts) | **GET** /social-microservice/api/v1/users/bookmarks/{userId} | Get Bookmarked Posts |
| [**getCannyUserSSOToken1**](SocialServiceApi.md#getCannyUserSSOToken1) | **GET** /social-microservice/api/v1/user/{userId}/canny/email/{email} | Retrieves user sso token for canny |
| [**getCommentReplies**](SocialServiceApi.md#getCommentReplies) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/replies | Get Comment Replies |
| [**getCommunitiesUserFollows**](SocialServiceApi.md#getCommunitiesUserFollows) | **GET** /social-microservice/api/v1/users/{userId}/communities-followed | Gets all the communities a user follows |
| [**getCommunityBlogPosts**](SocialServiceApi.md#getCommunityBlogPosts) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/posts/blog | Get community blog posts |
| [**getCommunityFeed**](SocialServiceApi.md#getCommunityFeed) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/timeline | Gets A Community Feed |
| [**getCommunityProfile**](SocialServiceApi.md#getCommunityProfile) | **GET** /social-microservice/api/v1/social/community-profiles/{communityId}/requestor/{requestorProfileId}/profile-type/{requestorProfileType} | Get a community Profile |
| [**getCommunityProfiles**](SocialServiceApi.md#getCommunityProfiles) | **GET** /social-microservice/api/v1/community-profiles/page-size/{pageSize}/page-number/{pageNumber} | Get Community Profiles |
| [**getFollowers**](SocialServiceApi.md#getFollowers) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/followers | Get Users Following you |
| [**getPendingFollows**](SocialServiceApi.md#getPendingFollows) | **GET** /social-microservice/api/v1/users/{userId}/follow/pending-requests | Get Pending Follow Requests |
| [**getPoll**](SocialServiceApi.md#getPoll) | **GET** /social-microservice/api/v1/users/{userId}/poll/{postId} | Get a poll |
| [**getPolls**](SocialServiceApi.md#getPolls) | **GET** /social-microservice/api/v1/users/{userId}/polls | Get all the polls of a given user |
| [**getPost**](SocialServiceApi.md#getPost) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Get a post |
| [**getPostThread**](SocialServiceApi.md#getPostThread) | **GET** /social-microservice/api/v1/users/{userId}/post/thread/{postId} | Gets A Post&#39;s Thread |
| [**getPosts**](SocialServiceApi.md#getPosts) | **GET** /social-microservice/api/v1/users/{userId}/posts | Get all the posts of a given user |
| [**getPostsByCategory**](SocialServiceApi.md#getPostsByCategory) | **GET** /social-microservice/api/v1/user/{userId}/category/{category}/posts/{postType}/limit/{limit}/offset/{offset} | Get all posts associated with a category |
| [**getPostsByTopic**](SocialServiceApi.md#getPostsByTopic) | **GET** /social-microservice/api/v1/community/{communityProfileId}/topic/{topicName}/posts | Get all posts associated with a topic |
| [**getPublication**](SocialServiceApi.md#getPublication) | **GET** /social-microservice/api/v1/users/{userId}/publication/{publicationId} | Gets a publication |
| [**getTopicsOfCommunitiesUserFollows**](SocialServiceApi.md#getTopicsOfCommunitiesUserFollows) | **GET** /social-microservice/api/v1/users/{userId}/topics | Get Topics Of Communities User Follows |
| [**getUserFeed**](SocialServiceApi.md#getUserFeed) | **GET** /social-microservice/api/v1/users/{userId}/timeline | Gets A Userfeed |
| [**getUserProfile**](SocialServiceApi.md#getUserProfile) | **GET** /social-microservice/api/v1/users/{userId} | gets a user profile |
| [**getUserProfiles**](SocialServiceApi.md#getUserProfiles) | **GET** /social-microservice/api/v1/users/page-size/{pageSize}/page-number/{pageNumber} | Gets a set of user profiles |
| [**healthCheck1**](SocialServiceApi.md#healthCheck1) | **GET** /social-microservice/api/v1/health | health check |
| [**reactToComment**](SocialServiceApi.md#reactToComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment |
| [**reactToCommentReply**](SocialServiceApi.md#reactToCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment reply |
| [**reactToPost**](SocialServiceApi.md#reactToPost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/account-type/{accountType}/reaction/{reaction} | Reacts to a post |
| [**readynessCheck1**](SocialServiceApi.md#readynessCheck1) | **GET** /social-microservice/api/v1/ready | readyness check |
| [**removeBookmarkedPost**](SocialServiceApi.md#removeBookmarkedPost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/bookmark | Deletes A Bookmarked Post |
| [**removeBookmarkedPublication**](SocialServiceApi.md#removeBookmarkedPublication) | **DELETE** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Deletes A Bookmarked Publication |
| [**removePostFromThread**](SocialServiceApi.md#removePostFromThread) | **DELETE** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType}/target/{participantPostId} | Deletes A Post From A Thread |
| [**reportComment**](SocialServiceApi.md#reportComment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/report | Report A Comment |
| [**reportCommentReply**](SocialServiceApi.md#reportCommentReply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/report | Report A Comment Reply |
| [**reportPost**](SocialServiceApi.md#reportPost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType}/report | Report a post |
| [**respondToPoll**](SocialServiceApi.md#respondToPoll) | **POST** /social-microservice/api/v1/users/{userId}/poll/{pollId} | Adds a user response to a given poll by a user |
| [**sharePost**](SocialServiceApi.md#sharePost) | **POST** /social-microservice/api/v1/users/{userId}/post/share/{parentPostId}/type/{parentPostType} | Share a post |


<a name="acceptFollowProfile"></a>
# **acceptFollowProfile**
> AcceptFollowProfileResponse acceptFollowProfile(followRecordId)

Accepts a user&#39;s follow request

    This endpoint enables a client to accept a follow request from a source a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **followRecordId** | **String**| The id of the follow record | type: uint64 | [default to null] |

### Return type

[**AcceptFollowProfileResponse**](../Models/AcceptFollowProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="addPostToPublication"></a>
# **addPostToPublication**
> AddPostToPublicationResponse addPostToPublication(editorUserId, publicationId, Post)

Add a post to a publication

    This endpoint enables a client to add a post to a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **editorUserId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |
| **Post** | [**Post**](../Models/Post.md)|  | |

### Return type

[**AddPostToPublicationResponse**](../Models/AddPostToPublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

<a name="addPublicationEditor"></a>
# **addPublicationEditor**
> AddPublicationEditorResponse addPublicationEditor(adminUserId, publicationId, editorUserId)

Adds an editor to a publication

    This endpoint enables a client to add an editor to a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adminUserId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |
| **editorUserId** | **String**|  | [default to null] |

### Return type

[**AddPublicationEditorResponse**](../Models/AddPublicationEditorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="blockUserProfile"></a>
# **blockUserProfile**
> BlockUserProfileResponse blockUserProfile(sourceUserId, targetUserId)

blocks a user profile

    This endpoint enables a client to block a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **sourceUserId** | **String**| the user ID trying to block another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **targetUserId** | **String**| the user ID being blocked by another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**BlockUserProfileResponse**](../Models/BlockUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="createCommunityProfile"></a>
# **createCommunityProfile**
> CreateCommunityProfileResponse createCommunityProfile(userId, CreateCommunityProfileBody)

Create a community Profile

    This endpoint enables a client to create a community profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to create this community profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **CreateCommunityProfileBody** | [**CreateCommunityProfileBody**](../Models/CreateCommunityProfileBody.md)|  | |

### Return type

[**CreateCommunityProfileResponse**](../Models/CreateCommunityProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createNote"></a>
# **createNote**
> CreateNoteResponse createNote(userId, postId, CreateNoteBody)

Creates and associates a note to a given post

    This endpoint enables a client to create and associate a not to a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **CreateNoteBody** | [**CreateNoteBody**](../Models/CreateNoteBody.md)|  | |

### Return type

[**CreateNoteResponse**](../Models/CreateNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createPoll"></a>
# **createPoll**
> CreatePollResponse createPoll(userId, PollPost)

Create a poll

    This endpoint enables a client to create a poll

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to create a post (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **PollPost** | [**PollPost**](../Models/PollPost.md)| The post payload | type: json_object | |

### Return type

[**CreatePollResponse**](../Models/CreatePollResponse.md)

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

<a name="createPublication"></a>
# **createPublication**
> CreatePublicationResponse createPublication(userId, Publication)

Creates a publication

    This endpoint enables a client to creare a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **Publication** | [**Publication**](../Models/Publication.md)|  | |

### Return type

[**CreatePublicationResponse**](../Models/CreatePublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createTopic"></a>
# **createTopic**
> CreateTopicResponse createTopic(userId, communityProfileId, Topic)

Create A Topic

    This endpoint enables a client to create a topic

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to whom the community is tied to (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **communityProfileId** | **String**| the community profile to associate the topic to | type: uint64 | [default to null] |
| **Topic** | [**Topic**](../Models/Topic.md)| topic payload | type: json_object | |

### Return type

[**CreateTopicResponse**](../Models/CreateTopicResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createUserProfile"></a>
# **createUserProfile**
> CreateUserProfileResponse createUserProfile(CreateUserProfileRequest)

creates a user profile

    This endpoint enables a client to create a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateUserProfileRequest** | [**CreateUserProfileRequest**](../Models/CreateUserProfileRequest.md)|  | |

### Return type

[**CreateUserProfileResponse**](../Models/CreateUserProfileResponse.md)

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

<a name="deleteCommunityProfile"></a>
# **deleteCommunityProfile**
> DeleteCommunityProfileResponse deleteCommunityProfile(userId, communityProfileId)

Delete Community Profile

    This endpoint enables a client to delete a community profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this community profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **communityProfileId** | **String**|  | [default to null] |

### Return type

[**DeleteCommunityProfileResponse**](../Models/DeleteCommunityProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteNote"></a>
# **deleteNote**
> DeleteNoteResponse deleteNote(userId, postId, noteId, postType)

Delete a note

    This endpoint enables a client to delete a note

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **noteId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**DeleteNoteResponse**](../Models/DeleteNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deletePoll"></a>
# **deletePoll**
> DeletePollResponse deletePoll(userId, postId)

Delete a poll

    This endpoint enables a client to delete a poll

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **postId** | **String**| The ID of the post attempted to be delete | type: string | [default to null] |

### Return type

[**DeletePollResponse**](../Models/DeletePollResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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

<a name="deletePostFromPublication"></a>
# **deletePostFromPublication**
> DeletePostFromPublicationResponse deletePostFromPublication(editorUserId, publicationId, postId, postType)

Deletes a post from a publication

    This endpoint enables a client to delete a post from a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **editorUserId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |

### Return type

[**DeletePostFromPublicationResponse**](../Models/DeletePostFromPublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deletePublication"></a>
# **deletePublication**
> DeletePublicationResponse deletePublication(adminUserId, publicationId)

Deletes a publication

    This endpoint enables a client to delete a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adminUserId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |

### Return type

[**DeletePublicationResponse**](../Models/DeletePublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deletePublicationEditor"></a>
# **deletePublicationEditor**
> DeletePublicationEditorResponse deletePublicationEditor(adminUserId, publicationId, editorUserId)

Deletes an editor to a publication

    This endpoint enables a client to add an editor to a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adminUserId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |
| **editorUserId** | **String**|  | [default to null] |

### Return type

[**DeletePublicationEditorResponse**](../Models/DeletePublicationEditorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteUserProfile"></a>
# **deleteUserProfile**
> DeleteUserProfileResponse deleteUserProfile(userId)

deletes a user profile

    This endpoint enables a client to delete a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**DeleteUserProfileResponse**](../Models/DeleteUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="discoverProfiles"></a>
# **discoverProfiles**
> DiscoverProfilesResponse discoverProfiles(userId, limit)

Discover Profiles

    This endpoint enables a client to discover a set of profiles he/she does not follow

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **limit** | **String**|  | [default to null] |

### Return type

[**DiscoverProfilesResponse**](../Models/DiscoverProfilesResponse.md)

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

<a name="editCommunityProfile"></a>
# **editCommunityProfile**
> EditCommunityProfileRequest editCommunityProfile(userId, communityProfileId, CommunityProfile)

Edit a community Profile

    This endpoint enables a client to update a community profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to update this community profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **communityProfileId** | **String**| The community profile ID of the community being updated | type: uint64 | [default to null] |
| **CommunityProfile** | [**CommunityProfile**](../Models/CommunityProfile.md)| The community profile being updated | type: json_object | |

### Return type

[**EditCommunityProfileRequest**](../Models/EditCommunityProfileRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="editNote"></a>
# **editNote**
> EditNoteResponse editNote(userId, postId, noteId, postType, Note)

Update a note

    This endpoint enables a client to update a note

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **postId** | **String**|  | [default to null] |
| **noteId** | **String**|  | [default to null] |
| **postType** | **String**|  | [default to POST_TYPE_UNSPECIFIED] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **Note** | [**Note**](../Models/Note.md)|  | |

### Return type

[**EditNoteResponse**](../Models/EditNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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

<a name="editUserProfile"></a>
# **editUserProfile**
> EditUserProfileResponse editUserProfile(userId, UserProfile)

update a user profile

    This endpoint performs an updates operation on a user profile based on the provided parametersThis update operation can span multiple services on specific cases (such as when the client is explicitly attempting to update the email of the user)All update operations are atomic by nature hence we should not expect any form of divergent state

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to update this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **UserProfile** | [**UserProfile**](../Models/UserProfile.md)| the profile payload | type: json_object | |

### Return type

[**EditUserProfileResponse**](../Models/EditUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

<a name="followProfile"></a>
# **followProfile**
> FollowProfileResponse followProfile(sourceUserId, targetUserId)

follow a user profile

    This endpoint enables a client to follow a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **sourceUserId** | **String**| the user ID trying to follow another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **targetUserId** | **String**| the user ID being followed by another user (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**FollowProfileResponse**](../Models/FollowProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="getCannyUserSSOToken1"></a>
# **getCannyUserSSOToken1**
> GetCannyUserSSOTokenResponse getCannyUserSSOToken1(userId, email)

Retrieves user sso token for canny

    Fetches a user sso token for canny

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **email** | **String**|  | [default to null] |

### Return type

[**GetCannyUserSSOTokenResponse**](../Models/GetCannyUserSSOTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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

<a name="getCommunitiesUserFollows"></a>
# **getCommunitiesUserFollows**
> GetCommunitiesUserFollowsResponse getCommunitiesUserFollows(userId, limit)

Gets all the communities a user follows

    This endpoint enables a client to get all the communities a user follows

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID who&#39;s communities follow set we want to obtain (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **limit** | **String**| the max number of communities to return | type: uint64 | [default to null] |

### Return type

[**GetCommunitiesUserFollowsResponse**](../Models/GetCommunitiesUserFollowsResponse.md)

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

<a name="getCommunityFeed"></a>
# **getCommunityFeed**
> GetCommunityFeedResponse getCommunityFeed(communityProfileId, feedType, accountType, nextPageToken)

Gets A Community Feed

    This endpoint enables a client query a community feed

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **communityProfileId** | **String**| The ID of the community we are trying to obtain the feed for | type: uint64 | [default to null] |
| **feedType** | **String**| The type of feed aiming to be obtained | type: string   - FEED_TYPE_PERSONAL: UserFeed is a profile&#39;s personal feed  - FEED_TYPE_NEWS: NewsFeed is a profile&#39;s timeline  - FEED_TYPE_NOTIFICATION: NotificationFeed encompasses a profile&#39;s notification feed | [default to FEED_TYPE_UNSPECIFIED] [enum: FEED_TYPE_UNSPECIFIED, FEED_TYPE_PERSONAL, FEED_TYPE_NEWS, FEED_TYPE_NOTIFICATION] |
| **accountType** | **String**| The type of account making the request to obtain the feed | type: string | [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **nextPageToken** | **String**|  | [optional] [default to null] |

### Return type

[**GetCommunityFeedResponse**](../Models/GetCommunityFeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCommunityProfile"></a>
# **getCommunityProfile**
> GetCommunityProfileResponse getCommunityProfile(communityId, requestorProfileId, requestorProfileType)

Get a community Profile

    This endpoint enables a client to get a community profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **communityId** | **String**| The communityID associated with the community being requested  | type: uint64 | [default to null] |
| **requestorProfileId** | **String**| The RequestorProfileID is an optional parameter used to check if the profileID (requestor) making a request for the record actually follows the record | [default to null] |
| **requestorProfileType** | **String**| The RequestorProfileType is an optional parameter which tells us what type of profile is the requestor | [default to null] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |

### Return type

[**GetCommunityProfileResponse**](../Models/GetCommunityProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCommunityProfiles"></a>
# **getCommunityProfiles**
> GetCommunityProfilesResponse getCommunityProfiles(pageSize, pageNumber)

Get Community Profiles

    This endpoint enables a client to get community profiles

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pageSize** | **Integer**|  | [default to null] |
| **pageNumber** | **Integer**|  | [default to null] |

### Return type

[**GetCommunityProfilesResponse**](../Models/GetCommunityProfilesResponse.md)

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

<a name="getPoll"></a>
# **getPoll**
> GetPollResponse getPoll(userId, postId)

Get a poll

    This endpoint enables a client to query a poll by id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **postId** | **String**| The ID of the post attempted to be delete | type: string | [default to null] |

### Return type

[**GetPollResponse**](../Models/GetPollResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPolls"></a>
# **getPolls**
> GetPollsResponse getPolls(userId)

Get all the polls of a given user

    This endpoint enables a client to query all polls tied to a given user id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to delete this user profile (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |

### Return type

[**GetPollsResponse**](../Models/GetPollsResponse.md)

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

<a name="getPublication"></a>
# **getPublication**
> GetPublicationResponse getPublication(userId, publicationId)

Gets a publication

    This endpoint enables a client to get a publication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **publicationId** | **String**|  | [default to null] |

### Return type

[**GetPublicationResponse**](../Models/GetPublicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTopicsOfCommunitiesUserFollows"></a>
# **getTopicsOfCommunitiesUserFollows**
> GetTopicsOfCommunitiesUserFollowsResponse getTopicsOfCommunitiesUserFollows(userId, limit)

Get Topics Of Communities User Follows

    This endpoint enables a client to get topics of communities a user follows

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID whose communities topics follow set we want to obtain (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **limit** | **String**|  | [default to null] |

### Return type

[**GetTopicsOfCommunitiesUserFollowsResponse**](../Models/GetTopicsOfCommunitiesUserFollowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserFeed"></a>
# **getUserFeed**
> GetUserFeedResponse getUserFeed(userId, feedType, accountType, nextPageToken)

Gets A Userfeed

    This endpoint enables a client query a user feed

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the user ID trying to obtain a given feed (NOTE: userID refers to the ID from the vantage point of the user service. This ID is the single source of truth for a given user across our suite of services) | type: uint64 | [default to null] |
| **feedType** | **String**| The type of feed aiming to be obtained | type: string   - FEED_TYPE_PERSONAL: UserFeed is a profile&#39;s personal feed  - FEED_TYPE_NEWS: NewsFeed is a profile&#39;s timeline  - FEED_TYPE_NOTIFICATION: NotificationFeed encompasses a profile&#39;s notification feed | [default to FEED_TYPE_UNSPECIFIED] [enum: FEED_TYPE_UNSPECIFIED, FEED_TYPE_PERSONAL, FEED_TYPE_NEWS, FEED_TYPE_NOTIFICATION] |
| **accountType** | **String**| The type of account making the request to obtain the feed | type: string | [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |
| **nextPageToken** | **String**|  | [optional] [default to null] |

### Return type

[**GetUserFeedResponse**](../Models/GetUserFeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserProfile"></a>
# **getUserProfile**
> GetUserProfileResponse getUserProfile(userId, requestorProfileId, requestorProfileType)

gets a user profile

    This endpoint performs a query against the social service to obtain a user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user ID associated with the profile we want to get | type: uint64 | [default to null] |
| **requestorProfileId** | **String**| The RequestorProfileID is an optional parameter used to check if the profileID (requestor) making a request for the record actually follows the record | [optional] [default to null] |
| **requestorProfileType** | **String**| The RequestorProfileType is an optional parameter which tells us what type of profile is the requestor | [optional] [default to ACCOUNT_TYPE_UNSPECIFIED] [enum: ACCOUNT_TYPE_UNSPECIFIED, ACCOUNT_TYPE_USER, ACCOUNT_TYPE_COMMUNITY] |

### Return type

[**GetUserProfileResponse**](../Models/GetUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserProfiles"></a>
# **getUserProfiles**
> GetUserProfilesResponse getUserProfiles(pageSize, pageNumber)

Gets a set of user profiles

    This endpoint enables a client to get a number of user profiles in a paginated manner

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pageSize** | **Integer**|  | [default to null] |
| **pageNumber** | **Integer**|  | [default to null] |

### Return type

[**GetUserProfilesResponse**](../Models/GetUserProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="healthCheck1"></a>
# **healthCheck1**
> HealthCheckResponse1 healthCheck1()

health check

    This endpoint performs a healc check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheckResponse1**](../Models/HealthCheckResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="readynessCheck1"></a>
# **readynessCheck1**
> ReadynessCheckResponse1 readynessCheck1()

readyness check

    This endpoint performs a readiness check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**ReadynessCheckResponse1**](../Models/ReadynessCheckResponse1.md)

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

<a name="respondToPoll"></a>
# **respondToPoll**
> RespondToPollResponse respondToPoll(userId, pollId, RespondToPollBody)

Adds a user response to a given poll by a user

    This endpoint enables a client to response to a poll

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **pollId** | **String**|  | [default to null] |
| **RespondToPollBody** | [**RespondToPollBody**](../Models/RespondToPollBody.md)|  | |

### Return type

[**RespondToPollResponse**](../Models/RespondToPollResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="sharePost"></a>
# **sharePost**
> SharePostResponse sharePost(userId, parentPostId, parentPostType, body)

Share a post

    This endpoint enables a client to share a post

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **parentPostId** | **String**|  | [default to null] |
| **parentPostType** | **String**|  | [default to null] [enum: POST_TYPE_UNSPECIFIED, POST_TYPE_POST, POST_TYPE_REPOST, POST_TYPE_QUESTION, POST_TYPE_ACHIEVEMENT, POST_TYPE_ANNOUNCEMENT, POST_TYPE_POLL, POST_TYPE_ARTICLE, POST_TYPE_SHORT_STORY] |
| **body** | **String**|  | |

### Return type

[**SharePostResponse**](../Models/SharePostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

