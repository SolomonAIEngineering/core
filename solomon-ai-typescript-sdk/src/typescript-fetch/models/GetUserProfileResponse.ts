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
import type { SocialRelationshipMetadata } from './SocialRelationshipMetadata';
import {
    SocialRelationshipMetadataFromJSON,
    SocialRelationshipMetadataFromJSONTyped,
    SocialRelationshipMetadataToJSON,
} from './SocialRelationshipMetadata';
import type { UserProfile } from './UserProfile';
import {
    UserProfileFromJSON,
    UserProfileFromJSONTyped,
    UserProfileToJSON,
} from './UserProfile';

/**
 * 
 * @export
 * @interface GetUserProfileResponse
 */
export interface GetUserProfileResponse {
    /**
     * 
     * @type {UserProfile}
     * @memberof GetUserProfileResponse
     */
    profile?: UserProfile;
    /**
     * 
     * @type {SocialRelationshipMetadata}
     * @memberof GetUserProfileResponse
     */
    metadata?: SocialRelationshipMetadata;
}

/**
 * Check if a given object implements the GetUserProfileResponse interface.
 */
export function instanceOfGetUserProfileResponse(value: object): boolean {
    return true;
}

export function GetUserProfileResponseFromJSON(json: any): GetUserProfileResponse {
    return GetUserProfileResponseFromJSONTyped(json, false);
}

export function GetUserProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetUserProfileResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'profile': json['profile'] == null ? undefined : UserProfileFromJSON(json['profile']),
        'metadata': json['metadata'] == null ? undefined : SocialRelationshipMetadataFromJSON(json['metadata']),
    };
}

export function GetUserProfileResponseToJSON(value?: GetUserProfileResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'profile': UserProfileToJSON(value['profile']),
        'metadata': SocialRelationshipMetadataToJSON(value['metadata']),
    };
}

