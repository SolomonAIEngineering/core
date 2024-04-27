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
import type { Pocket } from './Pocket';
import {
    PocketFromJSON,
    PocketFromJSONTyped,
    PocketToJSON,
} from './Pocket';

/**
 * 
 * @export
 * @interface UpdatePocketResponse
 */
export interface UpdatePocketResponse {
    /**
     * 
     * @type {Pocket}
     * @memberof UpdatePocketResponse
     */
    pocket?: Pocket;
}

/**
 * Check if a given object implements the UpdatePocketResponse interface.
 */
export function instanceOfUpdatePocketResponse(value: object): boolean {
    return true;
}

export function UpdatePocketResponseFromJSON(json: any): UpdatePocketResponse {
    return UpdatePocketResponseFromJSONTyped(json, false);
}

export function UpdatePocketResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdatePocketResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'pocket': json['pocket'] == null ? undefined : PocketFromJSON(json['pocket']),
    };
}

export function UpdatePocketResponseToJSON(value?: UpdatePocketResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'pocket': PocketToJSON(value['pocket']),
    };
}

