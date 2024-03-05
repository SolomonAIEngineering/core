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
 * @interface UpdateNoteToSmartGoalResponse
 */
export interface UpdateNoteToSmartGoalResponse {
    /**
     * 
     * @type {SmartNote}
     * @memberof UpdateNoteToSmartGoalResponse
     */
    note?: SmartNote;
}

/**
 * Check if a given object implements the UpdateNoteToSmartGoalResponse interface.
 */
export function instanceOfUpdateNoteToSmartGoalResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UpdateNoteToSmartGoalResponseFromJSON(json: any): UpdateNoteToSmartGoalResponse {
    return UpdateNoteToSmartGoalResponseFromJSONTyped(json, false);
}

export function UpdateNoteToSmartGoalResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateNoteToSmartGoalResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'note': !exists(json, 'note') ? undefined : SmartNoteFromJSON(json['note']),
    };
}

export function UpdateNoteToSmartGoalResponseToJSON(value?: UpdateNoteToSmartGoalResponse | null): any {
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
