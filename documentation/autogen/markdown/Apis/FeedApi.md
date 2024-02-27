# FeedApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCommunityFeed**](FeedApi.md#getCommunityFeed) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/timeline | Gets A Community Feed |
| [**getUserFeed**](FeedApi.md#getUserFeed) | **GET** /social-microservice/api/v1/users/{userId}/timeline | Gets A Userfeed |
| [**sharePost**](FeedApi.md#sharePost) | **POST** /social-microservice/api/v1/users/{userId}/post/share/{parentPostId}/type/{parentPostType} | Share a post |


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

