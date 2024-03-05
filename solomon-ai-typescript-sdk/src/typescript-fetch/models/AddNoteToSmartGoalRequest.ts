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
 * @interface AddNoteToSmartGoalRequest
 */
export interface AddNoteToSmartGoalRequest {
    /**
     * 
     * @type {string}
     * @memberof AddNoteToSmartGoalRequest
     */
    smartGoalId: string;
    /**
     * 
     * @type {SmartNote}
     * @memberof AddNoteToSmartGoalRequest
     */
    note: SmartNote;
}

/**
 * Check if a given object implements the AddNoteToSmartGoalRequest interface.
 */
export function instanceOfAddNoteToSmartGoalRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "smartGoalId" in value;
    isInstance = isInstance && "note" in value;

    return isInstance;
}

export function AddNoteToSmartGoalRequestFromJSON(json: any): AddNoteToSmartGoalRequest {
    return AddNoteToSmartGoalRequestFromJSONTyped(json, false);
}

export function AddNoteToSmartGoalRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddNoteToSmartGoalRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'smartGoalId': json['smartGoalId'],
        'note': SmartNoteFromJSON(json['note']),
    };
}

export function AddNoteToSmartGoalRequestToJSON(value?: AddNoteToSmartGoalRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'smartGoalId': value.smartGoalId,
        'note': SmartNoteToJSON(value.note),
    };
}
