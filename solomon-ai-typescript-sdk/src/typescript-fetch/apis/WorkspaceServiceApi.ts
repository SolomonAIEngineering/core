/* tslint:disable */
/* eslint-disable */
/**
 * User Service
 * Solomon AI User Service API
 *
 * The version of the OpenAPI document: 0.1
 * Contact: yoanyomba@solomon-ai.co
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  CreateAccountRequest,
  CreateAccountResponse,
  CreateFolderRequest,
  CreateFolderResponse,
  CreateWorkspaceRequest,
  CreateWorkspaceResponse,
  DeleteAccountResponse,
  DeleteFileResponse,
  DeleteFolderResponse,
  DeleteWorkspaceResponse,
  DownloadFileResponse,
  GetAccountResponse,
  InternalErrorMessageResponse,
  ListFolderResponse,
  ListWorkspaceResponse,
  PathUnknownErrorMessageResponse,
  Status,
  UpdateFileRequest,
  UpdateFileResponse,
  UpdateFolderRequest,
  UpdateFolderResponse,
  UpdateWorkspaceRequest,
  UpdateWorkspaceResponse,
  UploadFileRequest,
  UploadFileResponse,
  ValidationErrorMessageResponse,
} from '../models/index';
import {
    CreateAccountRequestFromJSON,
    CreateAccountRequestToJSON,
    CreateAccountResponseFromJSON,
    CreateAccountResponseToJSON,
    CreateFolderRequestFromJSON,
    CreateFolderRequestToJSON,
    CreateFolderResponseFromJSON,
    CreateFolderResponseToJSON,
    CreateWorkspaceRequestFromJSON,
    CreateWorkspaceRequestToJSON,
    CreateWorkspaceResponseFromJSON,
    CreateWorkspaceResponseToJSON,
    DeleteAccountResponseFromJSON,
    DeleteAccountResponseToJSON,
    DeleteFileResponseFromJSON,
    DeleteFileResponseToJSON,
    DeleteFolderResponseFromJSON,
    DeleteFolderResponseToJSON,
    DeleteWorkspaceResponseFromJSON,
    DeleteWorkspaceResponseToJSON,
    DownloadFileResponseFromJSON,
    DownloadFileResponseToJSON,
    GetAccountResponseFromJSON,
    GetAccountResponseToJSON,
    InternalErrorMessageResponseFromJSON,
    InternalErrorMessageResponseToJSON,
    ListFolderResponseFromJSON,
    ListFolderResponseToJSON,
    ListWorkspaceResponseFromJSON,
    ListWorkspaceResponseToJSON,
    PathUnknownErrorMessageResponseFromJSON,
    PathUnknownErrorMessageResponseToJSON,
    StatusFromJSON,
    StatusToJSON,
    UpdateFileRequestFromJSON,
    UpdateFileRequestToJSON,
    UpdateFileResponseFromJSON,
    UpdateFileResponseToJSON,
    UpdateFolderRequestFromJSON,
    UpdateFolderRequestToJSON,
    UpdateFolderResponseFromJSON,
    UpdateFolderResponseToJSON,
    UpdateWorkspaceRequestFromJSON,
    UpdateWorkspaceRequestToJSON,
    UpdateWorkspaceResponseFromJSON,
    UpdateWorkspaceResponseToJSON,
    UploadFileRequestFromJSON,
    UploadFileRequestToJSON,
    UploadFileResponseFromJSON,
    UploadFileResponseToJSON,
    ValidationErrorMessageResponseFromJSON,
    ValidationErrorMessageResponseToJSON,
} from '../models/index';

export interface CreateAccountOperationRequest {
    createAccountRequest: CreateAccountRequest;
}

export interface CreateFolderOperationRequest {
    createFolderRequest: CreateFolderRequest;
}

export interface CreateWorkspaceOperationRequest {
    createWorkspaceRequest: CreateWorkspaceRequest;
}

export interface DeleteAccountRequest {
    authZeroUserId: string;
}

export interface DeleteFileRequest {
    fileId: string;
    authZeroUserId: string;
    folderId?: string;
    workspaceId?: string;
}

export interface DeleteFolderRequest {
    folderId: string;
    authZeroUserId: string;
    workspaceId?: string;
}

export interface DeleteWorkspaceRequest {
    workspaceId: string;
    authZeroUserId: string;
}

export interface DownloadFileRequest {
    fileId: string;
    authZeroUserId: string;
    folderId?: string;
    workspaceId?: string;
}

export interface GetAccountRequest {
    authZeroUserId: string;
}

export interface ListFolderRequest {
    authZeroUserId: string;
    folderId?: string;
    workspaceId?: string;
}

export interface ListWorkspaceRequest {
    authZeroUserId: string;
}

export interface UpdateFileOperationRequest {
    updateFileRequest: UpdateFileRequest;
}

export interface UpdateFolderOperationRequest {
    updateFolderRequest: UpdateFolderRequest;
}

export interface UpdateWorkspaceOperationRequest {
    updateWorkspaceRequest: UpdateWorkspaceRequest;
}

export interface UploadFileOperationRequest {
    uploadFileRequest: UploadFileRequest;
}

/**
 * 
 */
export class WorkspaceServiceApi extends runtime.BaseAPI {

    /**
     * This endpoint creates a new user account
     * Create a new account
     */
    async createAccountRaw(requestParameters: CreateAccountOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateAccountResponse>> {
        if (requestParameters.createAccountRequest === null || requestParameters.createAccountRequest === undefined) {
            throw new runtime.RequiredError('createAccountRequest','Required parameter requestParameters.createAccountRequest was null or undefined when calling createAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/accounts`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateAccountRequestToJSON(requestParameters.createAccountRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateAccountResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint creates a new user account
     * Create a new account
     */
    async createAccount(requestParameters: CreateAccountOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateAccountResponse> {
        const response = await this.createAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint creates a new folder
     * Create a folder
     */
    async createFolderRaw(requestParameters: CreateFolderOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateFolderResponse>> {
        if (requestParameters.createFolderRequest === null || requestParameters.createFolderRequest === undefined) {
            throw new runtime.RequiredError('createFolderRequest','Required parameter requestParameters.createFolderRequest was null or undefined when calling createFolder.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/folders`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateFolderRequestToJSON(requestParameters.createFolderRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateFolderResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint creates a new folder
     * Create a folder
     */
    async createFolder(requestParameters: CreateFolderOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateFolderResponse> {
        const response = await this.createFolderRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint creates a new workspace
     * Create a workspace
     */
    async createWorkspaceRaw(requestParameters: CreateWorkspaceOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateWorkspaceResponse>> {
        if (requestParameters.createWorkspaceRequest === null || requestParameters.createWorkspaceRequest === undefined) {
            throw new runtime.RequiredError('createWorkspaceRequest','Required parameter requestParameters.createWorkspaceRequest was null or undefined when calling createWorkspace.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/workspaces`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateWorkspaceRequestToJSON(requestParameters.createWorkspaceRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateWorkspaceResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint creates a new workspace
     * Create a workspace
     */
    async createWorkspace(requestParameters: CreateWorkspaceOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateWorkspaceResponse> {
        const response = await this.createWorkspaceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint deletes an account by ID
     * Delete an account
     */
    async deleteAccountRaw(requestParameters: DeleteAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeleteAccountResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling deleteAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/accounts/{authZeroUserId}`.replace(`{${"authZeroUserId"}}`, encodeURIComponent(String(requestParameters.authZeroUserId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DeleteAccountResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint deletes an account by ID
     * Delete an account
     */
    async deleteAccount(requestParameters: DeleteAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeleteAccountResponse> {
        const response = await this.deleteAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint deletes a file by ID
     * Delete a file
     */
    async deleteFileRaw(requestParameters: DeleteFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeleteFileResponse>> {
        if (requestParameters.fileId === null || requestParameters.fileId === undefined) {
            throw new runtime.RequiredError('fileId','Required parameter requestParameters.fileId was null or undefined when calling deleteFile.');
        }

        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling deleteFile.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        if (requestParameters.folderId !== undefined) {
            queryParameters['folderId'] = requestParameters.folderId;
        }

        if (requestParameters.workspaceId !== undefined) {
            queryParameters['workspaceId'] = requestParameters.workspaceId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/files/{fileId}`.replace(`{${"fileId"}}`, encodeURIComponent(String(requestParameters.fileId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DeleteFileResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint deletes a file by ID
     * Delete a file
     */
    async deleteFile(requestParameters: DeleteFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeleteFileResponse> {
        const response = await this.deleteFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint deletes a folder by ID
     * Delete a folder
     */
    async deleteFolderRaw(requestParameters: DeleteFolderRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeleteFolderResponse>> {
        if (requestParameters.folderId === null || requestParameters.folderId === undefined) {
            throw new runtime.RequiredError('folderId','Required parameter requestParameters.folderId was null or undefined when calling deleteFolder.');
        }

        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling deleteFolder.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        if (requestParameters.workspaceId !== undefined) {
            queryParameters['workspaceId'] = requestParameters.workspaceId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/folders/{folderId}`.replace(`{${"folderId"}}`, encodeURIComponent(String(requestParameters.folderId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DeleteFolderResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint deletes a folder by ID
     * Delete a folder
     */
    async deleteFolder(requestParameters: DeleteFolderRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeleteFolderResponse> {
        const response = await this.deleteFolderRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint deletes a workspace by ID
     * Delete a workspace
     */
    async deleteWorkspaceRaw(requestParameters: DeleteWorkspaceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeleteWorkspaceResponse>> {
        if (requestParameters.workspaceId === null || requestParameters.workspaceId === undefined) {
            throw new runtime.RequiredError('workspaceId','Required parameter requestParameters.workspaceId was null or undefined when calling deleteWorkspace.');
        }

        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling deleteWorkspace.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/workspaces/{workspaceId}`.replace(`{${"workspaceId"}}`, encodeURIComponent(String(requestParameters.workspaceId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DeleteWorkspaceResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint deletes a workspace by ID
     * Delete a workspace
     */
    async deleteWorkspace(requestParameters: DeleteWorkspaceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeleteWorkspaceResponse> {
        const response = await this.deleteWorkspaceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint downloads a file by ID
     * Download a file
     */
    async downloadFileRaw(requestParameters: DownloadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DownloadFileResponse>> {
        if (requestParameters.fileId === null || requestParameters.fileId === undefined) {
            throw new runtime.RequiredError('fileId','Required parameter requestParameters.fileId was null or undefined when calling downloadFile.');
        }

        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling downloadFile.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        if (requestParameters.folderId !== undefined) {
            queryParameters['folderId'] = requestParameters.folderId;
        }

        if (requestParameters.workspaceId !== undefined) {
            queryParameters['workspaceId'] = requestParameters.workspaceId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/files/{fileId}`.replace(`{${"fileId"}}`, encodeURIComponent(String(requestParameters.fileId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DownloadFileResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint downloads a file by ID
     * Download a file
     */
    async downloadFile(requestParameters: DownloadFileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DownloadFileResponse> {
        const response = await this.downloadFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint retrieves account details by ID
     * Get account by ID
     */
    async getAccountRaw(requestParameters: GetAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetAccountResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling getAccount.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/accounts/{authZeroUserId}`.replace(`{${"authZeroUserId"}}`, encodeURIComponent(String(requestParameters.authZeroUserId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetAccountResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint retrieves account details by ID
     * Get account by ID
     */
    async getAccount(requestParameters: GetAccountRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetAccountResponse> {
        const response = await this.getAccountRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint lists all folders
     * List folders
     */
    async listFolderRaw(requestParameters: ListFolderRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ListFolderResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling listFolder.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        if (requestParameters.folderId !== undefined) {
            queryParameters['folderId'] = requestParameters.folderId;
        }

        if (requestParameters.workspaceId !== undefined) {
            queryParameters['workspaceId'] = requestParameters.workspaceId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/folders`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ListFolderResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint lists all folders
     * List folders
     */
    async listFolder(requestParameters: ListFolderRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ListFolderResponse> {
        const response = await this.listFolderRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint lists all workspaces
     * List workspaces
     */
    async listWorkspaceRaw(requestParameters: ListWorkspaceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ListWorkspaceResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling listWorkspace.');
        }

        const queryParameters: any = {};

        if (requestParameters.authZeroUserId !== undefined) {
            queryParameters['authZeroUserId'] = requestParameters.authZeroUserId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/workspace-microservice/api/v1/workspaces`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ListWorkspaceResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint lists all workspaces
     * List workspaces
     */
    async listWorkspace(requestParameters: ListWorkspaceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ListWorkspaceResponse> {
        const response = await this.listWorkspaceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint updates a file by ID
     * Update a file
     */
    async updateFileRaw(requestParameters: UpdateFileOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UpdateFileResponse>> {
        if (requestParameters.updateFileRequest === null || requestParameters.updateFileRequest === undefined) {
            throw new runtime.RequiredError('updateFileRequest','Required parameter requestParameters.updateFileRequest was null or undefined when calling updateFile.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/files`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateFileRequestToJSON(requestParameters.updateFileRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UpdateFileResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint updates a file by ID
     * Update a file
     */
    async updateFile(requestParameters: UpdateFileOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UpdateFileResponse> {
        const response = await this.updateFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint updates a folder by ID
     * Update a folder
     */
    async updateFolderRaw(requestParameters: UpdateFolderOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UpdateFolderResponse>> {
        if (requestParameters.updateFolderRequest === null || requestParameters.updateFolderRequest === undefined) {
            throw new runtime.RequiredError('updateFolderRequest','Required parameter requestParameters.updateFolderRequest was null or undefined when calling updateFolder.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/folders`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateFolderRequestToJSON(requestParameters.updateFolderRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UpdateFolderResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint updates a folder by ID
     * Update a folder
     */
    async updateFolder(requestParameters: UpdateFolderOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UpdateFolderResponse> {
        const response = await this.updateFolderRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint updates a workspace by ID
     * Update a workspace
     */
    async updateWorkspaceRaw(requestParameters: UpdateWorkspaceOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UpdateWorkspaceResponse>> {
        if (requestParameters.updateWorkspaceRequest === null || requestParameters.updateWorkspaceRequest === undefined) {
            throw new runtime.RequiredError('updateWorkspaceRequest','Required parameter requestParameters.updateWorkspaceRequest was null or undefined when calling updateWorkspace.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/workspaces`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateWorkspaceRequestToJSON(requestParameters.updateWorkspaceRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UpdateWorkspaceResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint updates a workspace by ID
     * Update a workspace
     */
    async updateWorkspace(requestParameters: UpdateWorkspaceOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UpdateWorkspaceResponse> {
        const response = await this.updateWorkspaceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * This endpoint uploads a file
     * Upload a file
     */
    async uploadFileRaw(requestParameters: UploadFileOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UploadFileResponse>> {
        if (requestParameters.uploadFileRequest === null || requestParameters.uploadFileRequest === undefined) {
            throw new runtime.RequiredError('uploadFileRequest','Required parameter requestParameters.uploadFileRequest was null or undefined when calling uploadFile.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/workspace-microservice/api/v1/files/upload`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: UploadFileRequestToJSON(requestParameters.uploadFileRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UploadFileResponseFromJSON(jsonValue));
    }

    /**
     * This endpoint uploads a file
     * Upload a file
     */
    async uploadFile(requestParameters: UploadFileOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UploadFileResponse> {
        const response = await this.uploadFileRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
