# WorkspaceServiceRestApi

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**workspaceMicroserviceRestApiV1FileUploadPost**](WorkspaceServiceRestApi.md#workspaceMicroserviceRestApiV1FileUploadPost) | **POST** /workspace-microservice/rest-api/v1/file/upload | Uploads a file to the server |


<a name="workspaceMicroserviceRestApiV1FileUploadPost"></a>
# **workspaceMicroserviceRestApiV1FileUploadPost**
> workspaceservicehttp.FileUploadResponse workspaceMicroserviceRestApiV1FileUploadPost(workspaceId, folderId, userId, attachment, filename)

Uploads a file to the server

    This endpoint allows for the uploading of a file to the server. Upon successful upload,

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceId** | **String**| Identifier of the workspace where the file will be uploaded | [default to null] |
| **folderId** | **String**| Identifier of the folder within the workspace where the file will be stored | [default to null] |
| **userId** | **String**| Identifier of the user uploading the file | [default to null] |
| **attachment** | **File**| The file to be uploaded | [default to null] |
| **filename** | **String**| The name of the file to be saved (optional) | [optional] [default to null] |

### Return type

[**workspaceservicehttp.FileUploadResponse**](../Models/workspaceservicehttp.FileUploadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

