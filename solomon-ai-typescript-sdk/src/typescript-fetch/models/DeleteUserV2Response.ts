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
 * @interface DeleteUserV2Response
 */
export interface DeleteUserV2Response {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteUserV2Response
     */
    accountDeleted?: boolean;
}

/**
 * Check if a given object implements the DeleteUserV2Response interface.
 */
export function instanceOfDeleteUserV2Response(value: object): boolean {
    return true;
}

export function DeleteUserV2ResponseFromJSON(json: any): DeleteUserV2Response {
    return DeleteUserV2ResponseFromJSONTyped(json, false);
}

export function DeleteUserV2ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteUserV2Response {
    if (json == null) {
        return json;
    }
    return {
        
        'accountDeleted': json['accountDeleted'] == null ? undefined : json['accountDeleted'],
    };
}

export function DeleteUserV2ResponseToJSON(value?: DeleteUserV2Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'accountDeleted': value['accountDeleted'],
    };
}

