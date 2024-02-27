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
import type { SmartNote } from './SmartNote';
import {
    SmartNoteFromJSON,
    SmartNoteFromJSONTyped,
    SmartNoteToJSON,
} from './SmartNote';

/**
 * 
 * @export
 * @interface UpdateNoteToRecurringTransactionRequest
 */
export interface UpdateNoteToRecurringTransactionRequest {
    /**
     * 
     * @type {SmartNote}
     * @memberof UpdateNoteToRecurringTransactionRequest
     */
    note: SmartNote;
}

/**
 * Check if a given object implements the UpdateNoteToRecurringTransactionRequest interface.
 */
export function instanceOfUpdateNoteToRecurringTransactionRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "note" in value;

    return isInstance;
}

export function UpdateNoteToRecurringTransactionRequestFromJSON(json: any): UpdateNoteToRecurringTransactionRequest {
    return UpdateNoteToRecurringTransactionRequestFromJSONTyped(json, false);
}

export function UpdateNoteToRecurringTransactionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateNoteToRecurringTransactionRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'note': SmartNoteFromJSON(json['note']),
    };
}

export function UpdateNoteToRecurringTransactionRequestToJSON(value?: UpdateNoteToRecurringTransactionRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'note': SmartNoteToJSON(value.note),
    };
}

