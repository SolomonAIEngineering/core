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
 * @interface PlaidInitiateTokenUpdateResponse
 */
export interface PlaidInitiateTokenUpdateResponse {
    /**
     * 
     * @type {string}
     * @memberof PlaidInitiateTokenUpdateResponse
     */
    linkToken?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidInitiateTokenUpdateResponse
     */
    expiration?: string;
}

/**
 * Check if a given object implements the PlaidInitiateTokenUpdateResponse interface.
 */
export function instanceOfPlaidInitiateTokenUpdateResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PlaidInitiateTokenUpdateResponseFromJSON(json: any): PlaidInitiateTokenUpdateResponse {
    return PlaidInitiateTokenUpdateResponseFromJSONTyped(json, false);
}

export function PlaidInitiateTokenUpdateResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlaidInitiateTokenUpdateResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'linkToken': !exists(json, 'linkToken') ? undefined : json['linkToken'],
        'expiration': !exists(json, 'expiration') ? undefined : json['expiration'],
    };
}

export function PlaidInitiateTokenUpdateResponseToJSON(value?: PlaidInitiateTokenUpdateResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'linkToken': value.linkToken,
        'expiration': value.expiration,
    };
}
