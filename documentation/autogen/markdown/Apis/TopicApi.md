# TopicApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTopic**](TopicApi.md#createTopic) | **POST** /social-microservice/api/v1/users/{userId}/community/{communityProfileId}/topic | Create A Topic |
| [**getTopicsOfCommunitiesUserFollows**](TopicApi.md#getTopicsOfCommunitiesUserFollows) | **GET** /social-microservice/api/v1/users/{userId}/topics | Get Topics Of Communities User Follows |


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

