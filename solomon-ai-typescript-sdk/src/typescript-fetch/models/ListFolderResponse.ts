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
import type { FolderMetadata } from './FolderMetadata';
import {
    FolderMetadataFromJSON,
    FolderMetadataFromJSONTyped,
    FolderMetadataToJSON,
} from './FolderMetadata';

/**
 * 
 * @export
 * @interface ListFolderResponse
 */
export interface ListFolderResponse {
    /**
     * 
     * @type {Array<FolderMetadata>}
     * @memberof ListFolderResponse
     */
    folder?: Array<FolderMetadata>;
}

/**
 * Check if a given object implements the ListFolderResponse interface.
 */
export function instanceOfListFolderResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ListFolderResponseFromJSON(json: any): ListFolderResponse {
    return ListFolderResponseFromJSONTyped(json, false);
}

export function ListFolderResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListFolderResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'folder': !exists(json, 'folder') ? undefined : ((json['folder'] as Array<any>).map(FolderMetadataFromJSON)),
    };
}

export function ListFolderResponseToJSON(value?: ListFolderResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'folder': value.folder === undefined ? undefined : ((value.folder as Array<any>).map(FolderMetadataToJSON)),
    };
}

