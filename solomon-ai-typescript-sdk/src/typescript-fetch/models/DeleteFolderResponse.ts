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
 * @interface DeleteFolderResponse
 */
export interface DeleteFolderResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteFolderResponse
     */
    isDeleted?: boolean;
}

/**
 * Check if a given object implements the DeleteFolderResponse interface.
 */
export function instanceOfDeleteFolderResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DeleteFolderResponseFromJSON(json: any): DeleteFolderResponse {
    return DeleteFolderResponseFromJSONTyped(json, false);
}

export function DeleteFolderResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteFolderResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'isDeleted': !exists(json, 'isDeleted') ? undefined : json['isDeleted'],
    };
}

export function DeleteFolderResponseToJSON(value?: DeleteFolderResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'isDeleted': value.isDeleted,
    };
}
