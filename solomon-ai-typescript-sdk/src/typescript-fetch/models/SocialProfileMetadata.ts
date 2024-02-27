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

/**
 * 
 * @export
 * @interface SocialProfileMetadata
 */
export interface SocialProfileMetadata {
    /**
     * 
     * @type {AccountType}
     * @memberof SocialProfileMetadata
     */
    profileType?: AccountType;
    /**
     * 
     * @type {string}
     * @memberof SocialProfileMetadata
     */
    profileId?: string;
}

/**
 * Check if a given object implements the SocialProfileMetadata interface.
 */
export function instanceOfSocialProfileMetadata(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SocialProfileMetadataFromJSON(json: any): SocialProfileMetadata {
    return SocialProfileMetadataFromJSONTyped(json, false);
}

export function SocialProfileMetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): SocialProfileMetadata {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileType': !exists(json, 'profileType') ? undefined : AccountTypeFromJSON(json['profileType']),
        'profileId': !exists(json, 'profileId') ? undefined : json['profileId'],
    };
}

export function SocialProfileMetadataToJSON(value?: SocialProfileMetadata | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileType': AccountTypeToJSON(value.profileType),
        'profileId': value.profileId,
    };
}

