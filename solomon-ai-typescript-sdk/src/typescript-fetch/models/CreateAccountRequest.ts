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
 * @interface CreateAccountRequest
 */
export interface CreateAccountRequest {
    /**
     * 
     * @type {string}
     * @memberof CreateAccountRequest
     */
    authZeroUserId: string;
}

/**
 * Check if a given object implements the CreateAccountRequest interface.
 */
export function instanceOfCreateAccountRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "authZeroUserId" in value;

    return isInstance;
}

export function CreateAccountRequestFromJSON(json: any): CreateAccountRequest {
    return CreateAccountRequestFromJSONTyped(json, false);
}

export function CreateAccountRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateAccountRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authZeroUserId': json['authZeroUserId'],
    };
}

export function CreateAccountRequestToJSON(value?: CreateAccountRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authZeroUserId': value.authZeroUserId,
    };
}
