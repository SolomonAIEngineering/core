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
import type { CommunityProfile } from './CommunityProfile';
import {
    CommunityProfileFromJSON,
    CommunityProfileFromJSONTyped,
    CommunityProfileToJSON,
} from './CommunityProfile';

/**
 * 
 * @export
 * @interface EditCommunityProfileResponse
 */
export interface EditCommunityProfileResponse {
    /**
     * 
     * @type {CommunityProfile}
     * @memberof EditCommunityProfileResponse
     */
    profile?: CommunityProfile;
}

/**
 * Check if a given object implements the EditCommunityProfileResponse interface.
 */
export function instanceOfEditCommunityProfileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EditCommunityProfileResponseFromJSON(json: any): EditCommunityProfileResponse {
    return EditCommunityProfileResponseFromJSONTyped(json, false);
}

export function EditCommunityProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EditCommunityProfileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profile': !exists(json, 'profile') ? undefined : CommunityProfileFromJSON(json['profile']),
    };
}

export function EditCommunityProfileResponseToJSON(value?: EditCommunityProfileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profile': CommunityProfileToJSON(value.profile),
    };
}

