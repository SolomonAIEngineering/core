# PollApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createPoll**](PollApi.md#createPoll) | **POST** /social-microservice/api/v1/users/{userId}/poll | Create a poll |
| [**deletePoll**](PollApi.md#deletePoll) | **DELETE** /social-microservice/api/v1/users/{userId}/poll/{postId} | Delete a poll |
| [**getPoll**](PollApi.md#getPoll) | **GET** /social-microservice/api/v1/users/{userId}/poll/{postId} | Get a poll |
| [**getPolls**](PollApi.md#getPolls) | **GET** /social-microservice/api/v1/users/{userId}/polls | Get all the polls of a given user |
| [**respondToPoll**](PollApi.md#respondToPoll) | **POST** /social-microservice/api/v1/users/{userId}/poll/{pollId} | Adds a user response to a given poll by a user |


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

