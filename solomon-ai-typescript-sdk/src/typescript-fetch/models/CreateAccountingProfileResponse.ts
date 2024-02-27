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
 * Defines a message named CreatesAccountingProfileResponse.
 * @export
 * @interface CreateAccountingProfileResponse
 */
export interface CreateAccountingProfileResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateAccountingProfileResponse
     */
    profileId?: string;
}

/**
 * Check if a given object implements the CreateAccountingProfileResponse interface.
 */
export function instanceOfCreateAccountingProfileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateAccountingProfileResponseFromJSON(json: any): CreateAccountingProfileResponse {
    return CreateAccountingProfileResponseFromJSONTyped(json, false);
}

export function CreateAccountingProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateAccountingProfileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileId': !exists(json, 'profileId') ? undefined : json['profileId'],
    };
}

export function CreateAccountingProfileResponseToJSON(value?: CreateAccountingProfileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileId': value.profileId,
    };
}

