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
    let isInstance = true;

    return isInstance;
}

export function UpdatePocketResponseFromJSON(json: any): UpdatePocketResponse {
    return UpdatePocketResponseFromJSONTyped(json, false);
}

export function UpdatePocketResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdatePocketResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pocket': !exists(json, 'pocket') ? undefined : PocketFromJSON(json['pocket']),
    };
}

export function UpdatePocketResponseToJSON(value?: UpdatePocketResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pocket': PocketToJSON(value.pocket),
    };
}

