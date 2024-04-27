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
 * @interface DeletePostResponse
 */
export interface DeletePostResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeletePostResponse
     */
    success?: boolean;
}

/**
 * Check if a given object implements the DeletePostResponse interface.
 */
export function instanceOfDeletePostResponse(value: object): boolean {
    return true;
}

export function DeletePostResponseFromJSON(json: any): DeletePostResponse {
    return DeletePostResponseFromJSONTyped(json, false);
}

export function DeletePostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeletePostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'success': json['success'] == null ? undefined : json['success'],
    };
}

export function DeletePostResponseToJSON(value?: DeletePostResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'success': value['success'],
    };
}

