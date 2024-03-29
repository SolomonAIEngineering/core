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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface WorkspaceservicehttpFileUploadResponse
 */
export interface WorkspaceservicehttpFileUploadResponse {
    /**
     * 
     * @type {string}
     * @memberof WorkspaceservicehttpFileUploadResponse
     */
    fileUrl?: string;
    /**
     * 
     * @type {number}
     * @memberof WorkspaceservicehttpFileUploadResponse
     */
    fileId?: number;
}

/**
 * Check if a given object implements the WorkspaceservicehttpFileUploadResponse interface.
 */
export function instanceOfWorkspaceservicehttpFileUploadResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function WorkspaceservicehttpFileUploadResponseFromJSON(json: any): WorkspaceservicehttpFileUploadResponse {
    return WorkspaceservicehttpFileUploadResponseFromJSONTyped(json, false);
}

export function WorkspaceservicehttpFileUploadResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): WorkspaceservicehttpFileUploadResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fileUrl': !exists(json, 'fileUrl') ? undefined : json['fileUrl'],
        'fileId': !exists(json, 'file_id') ? undefined : json['file_id'],
    };
}

export function WorkspaceservicehttpFileUploadResponseToJSON(value?: WorkspaceservicehttpFileUploadResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fileUrl': value.fileUrl,
        'file_id': value.fileId,
    };
}

