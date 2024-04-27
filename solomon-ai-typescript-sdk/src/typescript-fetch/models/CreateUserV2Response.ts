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
 * @interface CreateUserV2Response
 */
export interface CreateUserV2Response {
    /**
     * 
     * @type {string}
     * @memberof CreateUserV2Response
     */
    userId?: string;
}

/**
 * Check if a given object implements the CreateUserV2Response interface.
 */
export function instanceOfCreateUserV2Response(value: object): boolean {
    return true;
}

export function CreateUserV2ResponseFromJSON(json: any): CreateUserV2Response {
    return CreateUserV2ResponseFromJSONTyped(json, false);
}

export function CreateUserV2ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateUserV2Response {
    if (json == null) {
        return json;
    }
    return {
        
        'userId': json['userId'] == null ? undefined : json['userId'],
    };
}

export function CreateUserV2ResponseToJSON(value?: CreateUserV2Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'userId': value['userId'],
    };
}

