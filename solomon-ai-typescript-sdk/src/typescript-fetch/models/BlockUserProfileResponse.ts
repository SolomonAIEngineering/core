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
/**
 * 
 * @export
 * @interface BlockUserProfileResponse
 */
export interface BlockUserProfileResponse {
    /**
     * 
     * @type {boolean}
     * @memberof BlockUserProfileResponse
     */
    success?: boolean;
}

/**
 * Check if a given object implements the BlockUserProfileResponse interface.
 */
export function instanceOfBlockUserProfileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockUserProfileResponseFromJSON(json: any): BlockUserProfileResponse {
    return BlockUserProfileResponseFromJSONTyped(json, false);
}

export function BlockUserProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockUserProfileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'success': !exists(json, 'success') ? undefined : json['success'],
    };
}

export function BlockUserProfileResponseToJSON(value?: BlockUserProfileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'success': value.success,
    };
}

