# NoteApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNote**](NoteApi.md#createNote) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/note | Creates and associates a note to a given post |
| [**deleteNote**](NoteApi.md#deleteNote) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Delete a note |
| [**editNote**](NoteApi.md#editNote) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Update a note |


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

