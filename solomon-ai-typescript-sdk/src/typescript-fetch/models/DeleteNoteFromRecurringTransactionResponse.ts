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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface DeleteNoteFromRecurringTransactionResponse
 */
export interface DeleteNoteFromRecurringTransactionResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteNoteFromRecurringTransactionResponse
     */
    deleted?: boolean;
}

/**
 * Check if a given object implements the DeleteNoteFromRecurringTransactionResponse interface.
 */
export function instanceOfDeleteNoteFromRecurringTransactionResponse(value: object): boolean {
    return true;
}

export function DeleteNoteFromRecurringTransactionResponseFromJSON(json: any): DeleteNoteFromRecurringTransactionResponse {
    return DeleteNoteFromRecurringTransactionResponseFromJSONTyped(json, false);
}

export function DeleteNoteFromRecurringTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteNoteFromRecurringTransactionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'deleted': json['deleted'] == null ? undefined : json['deleted'],
    };
}

export function DeleteNoteFromRecurringTransactionResponseToJSON(value?: DeleteNoteFromRecurringTransactionResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'deleted': value['deleted'],
    };
}

