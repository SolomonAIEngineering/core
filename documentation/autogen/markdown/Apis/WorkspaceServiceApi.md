# WorkspaceServiceApi

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createAccount**](WorkspaceServiceApi.md#createAccount) | **POST** /workspace-microservice/api/v1/accounts | Create a new account |
| [**createFolder**](WorkspaceServiceApi.md#createFolder) | **POST** /workspace-microservice/api/v1/folders | Create a folder |
| [**createWorkspace**](WorkspaceServiceApi.md#createWorkspace) | **POST** /workspace-microservice/api/v1/workspaces | Create a workspace |
| [**deleteAccount**](WorkspaceServiceApi.md#deleteAccount) | **DELETE** /workspace-microservice/api/v1/accounts/{authZeroUserId} | Delete an account |
| [**deleteFile**](WorkspaceServiceApi.md#deleteFile) | **DELETE** /workspace-microservice/api/v1/files/{fileId} | Delete a file |
| [**deleteFolder**](WorkspaceServiceApi.md#deleteFolder) | **DELETE** /workspace-microservice/api/v1/folders/{folderId} | Delete a folder |
| [**deleteWorkspace**](WorkspaceServiceApi.md#deleteWorkspace) | **DELETE** /workspace-microservice/api/v1/workspaces/{workspaceId} | Delete a workspace |
| [**downloadFile**](WorkspaceServiceApi.md#downloadFile) | **GET** /workspace-microservice/api/v1/files/{fileId} | Download a file |
| [**getAccount**](WorkspaceServiceApi.md#getAccount) | **GET** /workspace-microservice/api/v1/accounts/{authZeroUserId} | Get account by ID |
| [**listFolder**](WorkspaceServiceApi.md#listFolder) | **GET** /workspace-microservice/api/v1/folders | List folders |
| [**listWorkspace**](WorkspaceServiceApi.md#listWorkspace) | **GET** /workspace-microservice/api/v1/workspaces | List workspaces |
| [**updateFile**](WorkspaceServiceApi.md#updateFile) | **PUT** /workspace-microservice/api/v1/files | Update a file |
| [**updateFolder**](WorkspaceServiceApi.md#updateFolder) | **PUT** /workspace-microservice/api/v1/folders | Update a folder |
| [**updateWorkspace**](WorkspaceServiceApi.md#updateWorkspace) | **PUT** /workspace-microservice/api/v1/workspaces | Update a workspace |
| [**uploadFile**](WorkspaceServiceApi.md#uploadFile) | **POST** /workspace-microservice/api/v1/files/upload | Upload a file |


<a name="createAccount"></a>
# **createAccount**
> CreateAccountResponse createAccount(CreateAccountRequest)

Create a new account

    This endpoint creates a new user account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateAccountRequest** | [**CreateAccountRequest**](../Models/CreateAccountRequest.md)|  | |

### Return type

[**CreateAccountResponse**](../Models/CreateAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createFolder"></a>
# **createFolder**
> CreateFolderResponse createFolder(CreateFolderRequest)

Create a folder

    This endpoint creates a new folder

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateFolderRequest** | [**CreateFolderRequest**](../Models/CreateFolderRequest.md)|  | |

### Return type

[**CreateFolderResponse**](../Models/CreateFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createWorkspace"></a>
# **createWorkspace**
> CreateWorkspaceResponse createWorkspace(CreateWorkspaceRequest)

Create a workspace

    This endpoint creates a new workspace

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateWorkspaceRequest** | [**CreateWorkspaceRequest**](../Models/CreateWorkspaceRequest.md)|  | |

### Return type

[**CreateWorkspaceResponse**](../Models/CreateWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteAccount"></a>
# **deleteAccount**
> DeleteAccountResponse deleteAccount(authZeroUserId)

Delete an account

    This endpoint deletes an account by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |

### Return type

[**DeleteAccountResponse**](../Models/DeleteAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteFile"></a>
# **deleteFile**
> DeleteFileResponse deleteFile(fileId, authZeroUserId, folderId, workspaceId)

Delete a file

    This endpoint deletes a file by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **fileId** | **String**|  | [default to null] |
| **authZeroUserId** | **String**|  | [default to null] |
| **folderId** | **String**|  | [optional] [default to null] |
| **workspaceId** | **String**|  | [optional] [default to null] |

### Return type

[**DeleteFileResponse**](../Models/DeleteFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteFolder"></a>
# **deleteFolder**
> DeleteFolderResponse deleteFolder(folderId, authZeroUserId, workspaceId)

Delete a folder

    This endpoint deletes a folder by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **folderId** | **String**|  | [default to null] |
| **authZeroUserId** | **String**|  | [default to null] |
| **workspaceId** | **String**|  | [optional] [default to null] |

### Return type

[**DeleteFolderResponse**](../Models/DeleteFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteWorkspace"></a>
# **deleteWorkspace**
> DeleteWorkspaceResponse deleteWorkspace(workspaceId, authZeroUserId)

Delete a workspace

    This endpoint deletes a workspace by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workspaceId** | **String**|  | [default to null] |
| **authZeroUserId** | **String**|  | [default to null] |

### Return type

[**DeleteWorkspaceResponse**](../Models/DeleteWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="downloadFile"></a>
# **downloadFile**
> DownloadFileResponse downloadFile(fileId, authZeroUserId, folderId, workspaceId)

Download a file

    This endpoint downloads a file by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **fileId** | **String**|  | [default to null] |
| **authZeroUserId** | **String**|  | [default to null] |
| **folderId** | **String**|  | [optional] [default to null] |
| **workspaceId** | **String**|  | [optional] [default to null] |

### Return type

[**DownloadFileResponse**](../Models/DownloadFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAccount"></a>
# **getAccount**
> GetAccountResponse getAccount(authZeroUserId)

Get account by ID

    This endpoint retrieves account details by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |

### Return type

[**GetAccountResponse**](../Models/GetAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listFolder"></a>
# **listFolder**
> ListFolderResponse listFolder(authZeroUserId, folderId, workspaceId)

List folders

    This endpoint lists all folders

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |
| **folderId** | **String**|  | [optional] [default to null] |
| **workspaceId** | **String**|  | [optional] [default to null] |

### Return type

[**ListFolderResponse**](../Models/ListFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listWorkspace"></a>
# **listWorkspace**
> ListWorkspaceResponse listWorkspace(authZeroUserId)

List workspaces

    This endpoint lists all workspaces

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |

### Return type

[**ListWorkspaceResponse**](../Models/ListWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateFile"></a>
# **updateFile**
> UpdateFileResponse updateFile(UpdateFileRequest)

Update a file

    This endpoint updates a file by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateFileRequest** | [**UpdateFileRequest**](../Models/UpdateFileRequest.md)|  | |

### Return type

[**UpdateFileResponse**](../Models/UpdateFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateFolder"></a>
# **updateFolder**
> UpdateFolderResponse updateFolder(UpdateFolderRequest)

Update a folder

    This endpoint updates a folder by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateFolderRequest** | [**UpdateFolderRequest**](../Models/UpdateFolderRequest.md)|  | |

### Return type

[**UpdateFolderResponse**](../Models/UpdateFolderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateWorkspace"></a>
# **updateWorkspace**
> UpdateWorkspaceResponse updateWorkspace(UpdateWorkspaceRequest)

Update a workspace

    This endpoint updates a workspace by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateWorkspaceRequest** | [**UpdateWorkspaceRequest**](../Models/UpdateWorkspaceRequest.md)|  | |

### Return type

[**UpdateWorkspaceResponse**](../Models/UpdateWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="uploadFile"></a>
# **uploadFile**
> UploadFileResponse uploadFile(UploadFileRequest)

Upload a file

    This endpoint uploads a file

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UploadFileRequest** | [**UploadFileRequest**](../Models/UploadFileRequest.md)|  (streaming inputs) | |

### Return type

[**UploadFileResponse**](../Models/UploadFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

