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
import type { UserProfile } from './UserProfile';
import {
    UserProfileFromJSON,
    UserProfileFromJSONTyped,
    UserProfileToJSON,
} from './UserProfile';

/**
 * 
 * @export
 * @interface GetAccountsFollowingResponse
 */
export interface GetAccountsFollowingResponse {
    /**
     * 
     * @type {Array<UserProfile>}
     * @memberof GetAccountsFollowingResponse
     */
    users?: Array<UserProfile>;
    /**
     * 
     * @type {Array<CommunityProfile>}
     * @memberof GetAccountsFollowingResponse
     */
    communities?: Array<CommunityProfile>;
}

/**
 * Check if a given object implements the GetAccountsFollowingResponse interface.
 */
export function instanceOfGetAccountsFollowingResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetAccountsFollowingResponseFromJSON(json: any): GetAccountsFollowingResponse {
    return GetAccountsFollowingResponseFromJSONTyped(json, false);
}

export function GetAccountsFollowingResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetAccountsFollowingResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'users': !exists(json, 'users') ? undefined : ((json['users'] as Array<any>).map(UserProfileFromJSON)),
        'communities': !exists(json, 'communities') ? undefined : ((json['communities'] as Array<any>).map(CommunityProfileFromJSON)),
    };
}

export function GetAccountsFollowingResponseToJSON(value?: GetAccountsFollowingResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'users': value.users === undefined ? undefined : ((value.users as Array<any>).map(UserProfileToJSON)),
        'communities': value.communities === undefined ? undefined : ((value.communities as Array<any>).map(CommunityProfileToJSON)),
    };
}
