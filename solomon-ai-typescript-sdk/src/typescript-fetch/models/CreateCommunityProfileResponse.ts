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
import type { CommunityProfile } from './CommunityProfile';
import {
    CommunityProfileFromJSON,
    CommunityProfileFromJSONTyped,
    CommunityProfileToJSON,
} from './CommunityProfile';

/**
 * 
 * @export
 * @interface CreateCommunityProfileResponse
 */
export interface CreateCommunityProfileResponse {
    /**
     * 
     * @type {CommunityProfile}
     * @memberof CreateCommunityProfileResponse
     */
    profile?: CommunityProfile;
}

/**
 * Check if a given object implements the CreateCommunityProfileResponse interface.
 */
export function instanceOfCreateCommunityProfileResponse(value: object): boolean {
    return true;
}

export function CreateCommunityProfileResponseFromJSON(json: any): CreateCommunityProfileResponse {
    return CreateCommunityProfileResponseFromJSONTyped(json, false);
}

export function CreateCommunityProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateCommunityProfileResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'profile': json['profile'] == null ? undefined : CommunityProfileFromJSON(json['profile']),
    };
}

export function CreateCommunityProfileResponseToJSON(value?: CreateCommunityProfileResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'profile': CommunityProfileToJSON(value['profile']),
    };
}

