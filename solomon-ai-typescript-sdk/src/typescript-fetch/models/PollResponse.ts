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
 * @interface PollResponse
 */
export interface PollResponse {
    /**
     * 
     * @type {string}
     * @memberof PollResponse
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof PollResponse
     */
    userId?: string;
    /**
     * 
     * @type {string}
     * @memberof PollResponse
     */
    responseValue?: string;
    /**
     * 
     * @type {string}
     * @memberof PollResponse
     */
    responseIdx?: string;
}

/**
 * Check if a given object implements the PollResponse interface.
 */
export function instanceOfPollResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PollResponseFromJSON(json: any): PollResponse {
    return PollResponseFromJSONTyped(json, false);
}

export function PollResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PollResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'responseValue': !exists(json, 'responseValue') ? undefined : json['responseValue'],
        'responseIdx': !exists(json, 'responseIdx') ? undefined : json['responseIdx'],
    };
}

export function PollResponseToJSON(value?: PollResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'userId': value.userId,
        'responseValue': value.responseValue,
        'responseIdx': value.responseIdx,
    };
}
