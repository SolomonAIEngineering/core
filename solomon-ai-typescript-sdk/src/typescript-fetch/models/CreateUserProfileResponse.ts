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
 * @interface CreateUserProfileResponse
 */
export interface CreateUserProfileResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateUserProfileResponse
     */
    virtualProfileId?: string;
}

/**
 * Check if a given object implements the CreateUserProfileResponse interface.
 */
export function instanceOfCreateUserProfileResponse(value: object): boolean {
    return true;
}

export function CreateUserProfileResponseFromJSON(json: any): CreateUserProfileResponse {
    return CreateUserProfileResponseFromJSONTyped(json, false);
}

export function CreateUserProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateUserProfileResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'virtualProfileId': json['virtualProfileId'] == null ? undefined : json['virtualProfileId'],
    };
}

export function CreateUserProfileResponseToJSON(value?: CreateUserProfileResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'virtualProfileId': value['virtualProfileId'],
    };
}

