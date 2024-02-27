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
 * @interface CheckEmailExistsV2Response
 */
export interface CheckEmailExistsV2Response {
    /**
     * 
     * @type {boolean}
     * @memberof CheckEmailExistsV2Response
     */
    _exists?: boolean;
}

/**
 * Check if a given object implements the CheckEmailExistsV2Response interface.
 */
export function instanceOfCheckEmailExistsV2Response(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CheckEmailExistsV2ResponseFromJSON(json: any): CheckEmailExistsV2Response {
    return CheckEmailExistsV2ResponseFromJSONTyped(json, false);
}

export function CheckEmailExistsV2ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CheckEmailExistsV2Response {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        '_exists': !exists(json, 'exists') ? undefined : json['exists'],
    };
}

export function CheckEmailExistsV2ResponseToJSON(value?: CheckEmailExistsV2Response | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'exists': value._exists,
    };
}

