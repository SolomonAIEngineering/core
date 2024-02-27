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
 * @interface ListRecurringTransactionNotesResponse
 */
export interface ListRecurringTransactionNotesResponse {
    /**
     * 
     * @type {Array<SmartNote>}
     * @memberof ListRecurringTransactionNotesResponse
     */
    notes?: Array<SmartNote>;
}

/**
 * Check if a given object implements the ListRecurringTransactionNotesResponse interface.
 */
export function instanceOfListRecurringTransactionNotesResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ListRecurringTransactionNotesResponseFromJSON(json: any): ListRecurringTransactionNotesResponse {
    return ListRecurringTransactionNotesResponseFromJSONTyped(json, false);
}

export function ListRecurringTransactionNotesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListRecurringTransactionNotesResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'notes': !exists(json, 'notes') ? undefined : ((json['notes'] as Array<any>).map(SmartNoteFromJSON)),
    };
}

export function ListRecurringTransactionNotesResponseToJSON(value?: ListRecurringTransactionNotesResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'notes': value.notes === undefined ? undefined : ((value.notes as Array<any>).map(SmartNoteToJSON)),
    };
}

