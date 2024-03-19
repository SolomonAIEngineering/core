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
import type { MergeBusinessProfile } from './MergeBusinessProfile';
import {
    MergeBusinessProfileFromJSON,
    MergeBusinessProfileFromJSONTyped,
    MergeBusinessProfileToJSON,
} from './MergeBusinessProfile';

/**
 * Defines a message named UpdateAccountingProfileResponse.
 * @export
 * @interface UpdateAccountingProfileResponse
 */
export interface UpdateAccountingProfileResponse {
    /**
     * 
     * @type {MergeBusinessProfile}
     * @memberof UpdateAccountingProfileResponse
     */
    profile?: MergeBusinessProfile;
}

/**
 * Check if a given object implements the UpdateAccountingProfileResponse interface.
 */
export function instanceOfUpdateAccountingProfileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UpdateAccountingProfileResponseFromJSON(json: any): UpdateAccountingProfileResponse {
    return UpdateAccountingProfileResponseFromJSONTyped(json, false);
}

export function UpdateAccountingProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateAccountingProfileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profile': !exists(json, 'profile') ? undefined : MergeBusinessProfileFromJSON(json['profile']),
    };
}

export function UpdateAccountingProfileResponseToJSON(value?: UpdateAccountingProfileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profile': MergeBusinessProfileToJSON(value.profile),
    };
}
