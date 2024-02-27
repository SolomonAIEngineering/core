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
 * @interface DeleteNoteFromTransactionResponse
 */
export interface DeleteNoteFromTransactionResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteNoteFromTransactionResponse
     */
    deleted?: boolean;
}

/**
 * Check if a given object implements the DeleteNoteFromTransactionResponse interface.
 */
export function instanceOfDeleteNoteFromTransactionResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DeleteNoteFromTransactionResponseFromJSON(json: any): DeleteNoteFromTransactionResponse {
    return DeleteNoteFromTransactionResponseFromJSONTyped(json, false);
}

export function DeleteNoteFromTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteNoteFromTransactionResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deleted': !exists(json, 'deleted') ? undefined : json['deleted'],
    };
}

export function DeleteNoteFromTransactionResponseToJSON(value?: DeleteNoteFromTransactionResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deleted': value.deleted,
    };
}

