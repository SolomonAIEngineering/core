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
 * @interface DeleteNoteResponse
 */
export interface DeleteNoteResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteNoteResponse
     */
    success?: boolean;
}

/**
 * Check if a given object implements the DeleteNoteResponse interface.
 */
export function instanceOfDeleteNoteResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DeleteNoteResponseFromJSON(json: any): DeleteNoteResponse {
    return DeleteNoteResponseFromJSONTyped(json, false);
}

export function DeleteNoteResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteNoteResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'success': !exists(json, 'success') ? undefined : json['success'],
    };
}

export function DeleteNoteResponseToJSON(value?: DeleteNoteResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'success': value.success,
    };
}

