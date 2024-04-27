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
 * @interface ReadynessCheckResponse
 */
export interface ReadynessCheckResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ReadynessCheckResponse
     */
    healthy?: boolean;
}

/**
 * Check if a given object implements the ReadynessCheckResponse interface.
 */
export function instanceOfReadynessCheckResponse(value: object): boolean {
    return true;
}

export function ReadynessCheckResponseFromJSON(json: any): ReadynessCheckResponse {
    return ReadynessCheckResponseFromJSONTyped(json, false);
}

export function ReadynessCheckResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReadynessCheckResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'healthy': json['healthy'] == null ? undefined : json['healthy'],
    };
}

export function ReadynessCheckResponseToJSON(value?: ReadynessCheckResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'healthy': value['healthy'],
    };
}

