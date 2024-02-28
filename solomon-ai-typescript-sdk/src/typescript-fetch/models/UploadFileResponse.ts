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
import type { FileMetadata } from './FileMetadata';
import {
    FileMetadataFromJSON,
    FileMetadataFromJSONTyped,
    FileMetadataToJSON,
} from './FileMetadata';

/**
 * 
 * @export
 * @interface UploadFileResponse
 */
export interface UploadFileResponse {
    /**
     * 
     * @type {FileMetadata}
     * @memberof UploadFileResponse
     */
    metadata?: FileMetadata;
}

/**
 * Check if a given object implements the UploadFileResponse interface.
 */
export function instanceOfUploadFileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UploadFileResponseFromJSON(json: any): UploadFileResponse {
    return UploadFileResponseFromJSONTyped(json, false);
}

export function UploadFileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UploadFileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'metadata': !exists(json, 'metadata') ? undefined : FileMetadataFromJSON(json['metadata']),
    };
}

export function UploadFileResponseToJSON(value?: UploadFileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'metadata': FileMetadataToJSON(value.metadata),
    };
}

