# CommunityProfileApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createCommunityProfile**](CommunityProfileApi.md#createCommunityProfile) | **POST** /social-microservice/api/v1/community-profiles/{userId} | Create a community Profile |
| [**deleteCommunityProfile**](CommunityProfileApi.md#deleteCommunityProfile) | **DELETE** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Delete Community Profile |
| [**editCommunityProfile**](CommunityProfileApi.md#editCommunityProfile) | **PUT** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Edit a community Profile |
| [**getCommunitiesUserFollows**](CommunityProfileApi.md#getCommunitiesUserFollows) | **GET** /social-microservice/api/v1/users/{userId}/communities-followed | Gets all the communities a user follows |
| [**getCommunityProfile**](CommunityProfileApi.md#getCommunityProfile) | **GET** /social-microservice/api/v1/social/community-profiles/{communityId}/requestor/{requestorProfileId}/profile-type/{requestorProfileType} | Get a community Profile |
| [**getCommunityProfiles**](CommunityProfileApi.md#getCommunityProfiles) | **GET** /social-microservice/api/v1/community-profiles/page-size/{pageSize}/page-number/{pageNumber} | Get Community Profiles |


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

