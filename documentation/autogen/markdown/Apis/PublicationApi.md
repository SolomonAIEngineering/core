# PublicationApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addPostToPublication**](PublicationApi.md#addPostToPublication) | **POST** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId} | Add a post to a publication |
| [**addPublicationEditor**](PublicationApi.md#addPublicationEditor) | **PUT** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Adds an editor to a publication |
| [**createPublication**](PublicationApi.md#createPublication) | **POST** /social-microservice/api/v1/users/{userId}/publication | Creates a publication |
| [**deletePostFromPublication**](PublicationApi.md#deletePostFromPublication) | **DELETE** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId}/post/{postId} | Deletes a post from a publication |
| [**deletePublication**](PublicationApi.md#deletePublication) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId} | Deletes a publication |
| [**deletePublicationEditor**](PublicationApi.md#deletePublicationEditor) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Deletes an editor to a publication |
| [**getPublication**](PublicationApi.md#getPublication) | **GET** /social-microservice/api/v1/users/{userId}/publication/{publicationId} | Gets a publication |


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

