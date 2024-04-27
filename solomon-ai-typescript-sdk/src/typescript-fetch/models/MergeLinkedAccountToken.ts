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
 * @interface MergeLinkedAccountToken
 */
export interface MergeLinkedAccountToken {
    /**
     * 
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    itemId?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    keyId?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    accessToken?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    version?: string;
    /**
     * This is what you'll pass to Merge as the end_user_origin_id.
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    mergeEndUserOriginId?: string;
    /**
     * The integration slug/identifier. This is returned at the end of the linking flow.
     * @type {string}
     * @memberof MergeLinkedAccountToken
     */
    mergeIntegrationSlug?: string;
}

/**
 * Check if a given object implements the MergeLinkedAccountToken interface.
 */
export function instanceOfMergeLinkedAccountToken(value: object): boolean {
    return true;
}

export function MergeLinkedAccountTokenFromJSON(json: any): MergeLinkedAccountToken {
    return MergeLinkedAccountTokenFromJSONTyped(json, false);
}

export function MergeLinkedAccountTokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): MergeLinkedAccountToken {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'itemId': json['itemId'] == null ? undefined : json['itemId'],
        'keyId': json['keyId'] == null ? undefined : json['keyId'],
        'accessToken': json['accessToken'] == null ? undefined : json['accessToken'],
        'version': json['version'] == null ? undefined : json['version'],
        'mergeEndUserOriginId': json['mergeEndUserOriginId'] == null ? undefined : json['mergeEndUserOriginId'],
        'mergeIntegrationSlug': json['mergeIntegrationSlug'] == null ? undefined : json['mergeIntegrationSlug'],
    };
}

export function MergeLinkedAccountTokenToJSON(value?: MergeLinkedAccountToken | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'itemId': value['itemId'],
        'keyId': value['keyId'],
        'accessToken': value['accessToken'],
        'version': value['version'],
        'mergeEndUserOriginId': value['mergeEndUserOriginId'],
        'mergeIntegrationSlug': value['mergeIntegrationSlug'],
    };
}

