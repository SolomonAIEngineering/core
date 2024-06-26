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
import type { HrisLinkedAccount } from './HrisLinkedAccount';
import {
    HrisLinkedAccountFromJSON,
    HrisLinkedAccountFromJSONTyped,
    HrisLinkedAccountToJSON,
} from './HrisLinkedAccount';
import type { MergeLinkedAccountToken } from './MergeLinkedAccountToken';
import {
    MergeLinkedAccountTokenFromJSON,
    MergeLinkedAccountTokenFromJSONTyped,
    MergeLinkedAccountTokenToJSON,
} from './MergeLinkedAccountToken';

/**
 * 
 * @export
 * @interface HrisIntegrationMergeLink
 */
export interface HrisIntegrationMergeLink {
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    integration?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    integrationSlug?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    category?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    endUserOriginId?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    endUserOrganizationName?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    endUserEmailAddress?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    status?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    webhookListenerUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof HrisIntegrationMergeLink
     */
    isDuplicate?: boolean;
    /**
     * 
     * @type {MergeLinkedAccountToken}
     * @memberof HrisIntegrationMergeLink
     */
    token?: MergeLinkedAccountToken;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    integrationName?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    integrationImage?: string;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    integrationSquareImage?: string;
    /**
     * 
     * @type {HrisLinkedAccount}
     * @memberof HrisIntegrationMergeLink
     */
    account?: HrisLinkedAccount;
    /**
     * 
     * @type {string}
     * @memberof HrisIntegrationMergeLink
     */
    mergeLinkedAccountId?: string;
    /**
     * 
     * @type {Date}
     * @memberof HrisIntegrationMergeLink
     */
    lastModifiedAt?: Date;
}

/**
 * Check if a given object implements the HrisIntegrationMergeLink interface.
 */
export function instanceOfHrisIntegrationMergeLink(value: object): boolean {
    return true;
}

export function HrisIntegrationMergeLinkFromJSON(json: any): HrisIntegrationMergeLink {
    return HrisIntegrationMergeLinkFromJSONTyped(json, false);
}

export function HrisIntegrationMergeLinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): HrisIntegrationMergeLink {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'integration': json['integration'] == null ? undefined : json['integration'],
        'integrationSlug': json['integrationSlug'] == null ? undefined : json['integrationSlug'],
        'category': json['category'] == null ? undefined : json['category'],
        'endUserOriginId': json['endUserOriginId'] == null ? undefined : json['endUserOriginId'],
        'endUserOrganizationName': json['endUserOrganizationName'] == null ? undefined : json['endUserOrganizationName'],
        'endUserEmailAddress': json['endUserEmailAddress'] == null ? undefined : json['endUserEmailAddress'],
        'status': json['status'] == null ? undefined : json['status'],
        'webhookListenerUrl': json['webhookListenerUrl'] == null ? undefined : json['webhookListenerUrl'],
        'isDuplicate': json['isDuplicate'] == null ? undefined : json['isDuplicate'],
        'token': json['token'] == null ? undefined : MergeLinkedAccountTokenFromJSON(json['token']),
        'integrationName': json['integrationName'] == null ? undefined : json['integrationName'],
        'integrationImage': json['integrationImage'] == null ? undefined : json['integrationImage'],
        'integrationSquareImage': json['integrationSquareImage'] == null ? undefined : json['integrationSquareImage'],
        'account': json['account'] == null ? undefined : HrisLinkedAccountFromJSON(json['account']),
        'mergeLinkedAccountId': json['mergeLinkedAccountId'] == null ? undefined : json['mergeLinkedAccountId'],
        'lastModifiedAt': json['lastModifiedAt'] == null ? undefined : (new Date(json['lastModifiedAt'])),
    };
}

export function HrisIntegrationMergeLinkToJSON(value?: HrisIntegrationMergeLink | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'integration': value['integration'],
        'integrationSlug': value['integrationSlug'],
        'category': value['category'],
        'endUserOriginId': value['endUserOriginId'],
        'endUserOrganizationName': value['endUserOrganizationName'],
        'endUserEmailAddress': value['endUserEmailAddress'],
        'status': value['status'],
        'webhookListenerUrl': value['webhookListenerUrl'],
        'isDuplicate': value['isDuplicate'],
        'token': MergeLinkedAccountTokenToJSON(value['token']),
        'integrationName': value['integrationName'],
        'integrationImage': value['integrationImage'],
        'integrationSquareImage': value['integrationSquareImage'],
        'account': HrisLinkedAccountToJSON(value['account']),
        'mergeLinkedAccountId': value['mergeLinkedAccountId'],
        'lastModifiedAt': value['lastModifiedAt'] == null ? undefined : ((value['lastModifiedAt']).toISOString()),
    };
}

