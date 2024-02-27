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
import type { AccountType } from './AccountType';
import {
    AccountTypeFromJSON,
    AccountTypeFromJSONTyped,
    AccountTypeToJSON,
} from './AccountType';
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
 * @interface Actor
 */
export interface Actor {
    /**
     * 
     * @type {UserProfile}
     * @memberof Actor
     */
    userProfile: UserProfile;
    /**
     * 
     * @type {CommunityProfile}
     * @memberof Actor
     */
    community: CommunityProfile;
    /**
     * 
     * @type {AccountType}
     * @memberof Actor
     */
    actorType: AccountType;
}

/**
 * Check if a given object implements the Actor interface.
 */
export function instanceOfActor(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "userProfile" in value;
    isInstance = isInstance && "community" in value;
    isInstance = isInstance && "actorType" in value;

    return isInstance;
}

export function ActorFromJSON(json: any): Actor {
    return ActorFromJSONTyped(json, false);
}

export function ActorFromJSONTyped(json: any, ignoreDiscriminator: boolean): Actor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userProfile': UserProfileFromJSON(json['userProfile']),
        'community': CommunityProfileFromJSON(json['community']),
        'actorType': AccountTypeFromJSON(json['actorType']),
    };
}

export function ActorToJSON(value?: Actor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userProfile': UserProfileToJSON(value.userProfile),
        'community': CommunityProfileToJSON(value.community),
        'actorType': AccountTypeToJSON(value.actorType),
    };
}

