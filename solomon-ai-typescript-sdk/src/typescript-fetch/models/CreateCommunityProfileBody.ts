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
 * @interface CreateCommunityProfileBody
 */
export interface CreateCommunityProfileBody {
    /**
     * 
     * @type {CommunityProfile}
     * @memberof CreateCommunityProfileBody
     */
    profile: CommunityProfile;
}

/**
 * Check if a given object implements the CreateCommunityProfileBody interface.
 */
export function instanceOfCreateCommunityProfileBody(value: object): boolean {
    if (!('profile' in value)) return false;
    return true;
}

export function CreateCommunityProfileBodyFromJSON(json: any): CreateCommunityProfileBody {
    return CreateCommunityProfileBodyFromJSONTyped(json, false);
}

export function CreateCommunityProfileBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateCommunityProfileBody {
    if (json == null) {
        return json;
    }
    return {
        
        'profile': CommunityProfileFromJSON(json['profile']),
    };
}

export function CreateCommunityProfileBodyToJSON(value?: CreateCommunityProfileBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'profile': CommunityProfileToJSON(value['profile']),
    };
}

