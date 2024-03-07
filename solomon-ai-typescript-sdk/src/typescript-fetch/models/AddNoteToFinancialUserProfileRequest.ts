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
 * @interface AddNoteToFinancialUserProfileRequest
 */
export interface AddNoteToFinancialUserProfileRequest {
    /**
     * 
     * @type {string}
     * @memberof AddNoteToFinancialUserProfileRequest
     */
    businessAccountId: string;
    /**
     * 
     * @type {SmartNote}
     * @memberof AddNoteToFinancialUserProfileRequest
     */
    note: SmartNote;
    /**
     * 
     * @type {string}
     * @memberof AddNoteToFinancialUserProfileRequest
     */
    userId: string;
}

/**
 * Check if a given object implements the AddNoteToFinancialUserProfileRequest interface.
 */
export function instanceOfAddNoteToFinancialUserProfileRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "businessAccountId" in value;
    isInstance = isInstance && "note" in value;
    isInstance = isInstance && "userId" in value;

    return isInstance;
}

export function AddNoteToFinancialUserProfileRequestFromJSON(json: any): AddNoteToFinancialUserProfileRequest {
    return AddNoteToFinancialUserProfileRequestFromJSONTyped(json, false);
}

export function AddNoteToFinancialUserProfileRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddNoteToFinancialUserProfileRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'businessAccountId': json['businessAccountId'],
        'note': SmartNoteFromJSON(json['note']),
        'userId': json['userId'],
    };
}

export function AddNoteToFinancialUserProfileRequestToJSON(value?: AddNoteToFinancialUserProfileRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'businessAccountId': value.businessAccountId,
        'note': SmartNoteToJSON(value.note),
        'userId': value.userId,
    };
}

