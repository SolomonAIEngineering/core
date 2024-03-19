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
import type { AccountingIntegrationMergeLink } from './AccountingIntegrationMergeLink';
import {
    AccountingIntegrationMergeLinkFromJSON,
    AccountingIntegrationMergeLinkFromJSONTyped,
    AccountingIntegrationMergeLinkToJSON,
} from './AccountingIntegrationMergeLink';
import type { BusinessActionableInsight } from './BusinessActionableInsight';
import {
    BusinessActionableInsightFromJSON,
    BusinessActionableInsightFromJSONTyped,
    BusinessActionableInsightToJSON,
} from './BusinessActionableInsight';
import type { HrisIntegrationMergeLink } from './HrisIntegrationMergeLink';
import {
    HrisIntegrationMergeLinkFromJSON,
    HrisIntegrationMergeLinkFromJSONTyped,
    HrisIntegrationMergeLinkToJSON,
} from './HrisIntegrationMergeLink';

/**
 * 
 * @export
 * @interface MergeBusinessProfile
 */
export interface MergeBusinessProfile {
    /**
     * 
     * @type {string}
     * @memberof MergeBusinessProfile
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeBusinessProfile
     */
    authZeroUserId: string;
    /**
     * 
     * @type {string}
     * @memberof MergeBusinessProfile
     */
    companyName?: string;
    /**
     * 
     * @type {string}
     * @memberof MergeBusinessProfile
     */
    email: string;
    /**
     * 
     * @type {Array<AccountingIntegrationMergeLink>}
     * @memberof MergeBusinessProfile
     */
    accountingIntegrationMergeLink?: Array<AccountingIntegrationMergeLink>;
    /**
     * 
     * @type {Array<HrisIntegrationMergeLink>}
     * @memberof MergeBusinessProfile
     */
    payrollIntegrationMergeLink?: Array<HrisIntegrationMergeLink>;
    /**
     * 
     * @type {Array<BusinessActionableInsight>}
     * @memberof MergeBusinessProfile
     */
    actionablePersonalInsights?: Array<BusinessActionableInsight>;
}

/**
 * Check if a given object implements the MergeBusinessProfile interface.
 */
export function instanceOfMergeBusinessProfile(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "authZeroUserId" in value;
    isInstance = isInstance && "email" in value;

    return isInstance;
}

export function MergeBusinessProfileFromJSON(json: any): MergeBusinessProfile {
    return MergeBusinessProfileFromJSONTyped(json, false);
}

export function MergeBusinessProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): MergeBusinessProfile {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'authZeroUserId': json['authZeroUserId'],
        'companyName': !exists(json, 'companyName') ? undefined : json['companyName'],
        'email': json['email'],
        'accountingIntegrationMergeLink': !exists(json, 'accountingIntegrationMergeLink') ? undefined : ((json['accountingIntegrationMergeLink'] as Array<any>).map(AccountingIntegrationMergeLinkFromJSON)),
        'payrollIntegrationMergeLink': !exists(json, 'payrollIntegrationMergeLink') ? undefined : ((json['payrollIntegrationMergeLink'] as Array<any>).map(HrisIntegrationMergeLinkFromJSON)),
        'actionablePersonalInsights': !exists(json, 'actionablePersonalInsights') ? undefined : ((json['actionablePersonalInsights'] as Array<any>).map(BusinessActionableInsightFromJSON)),
    };
}

export function MergeBusinessProfileToJSON(value?: MergeBusinessProfile | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'authZeroUserId': value.authZeroUserId,
        'companyName': value.companyName,
        'email': value.email,
        'accountingIntegrationMergeLink': value.accountingIntegrationMergeLink === undefined ? undefined : ((value.accountingIntegrationMergeLink as Array<any>).map(AccountingIntegrationMergeLinkToJSON)),
        'payrollIntegrationMergeLink': value.payrollIntegrationMergeLink === undefined ? undefined : ((value.payrollIntegrationMergeLink as Array<any>).map(HrisIntegrationMergeLinkToJSON)),
        'actionablePersonalInsights': value.actionablePersonalInsights === undefined ? undefined : ((value.actionablePersonalInsights as Array<any>).map(BusinessActionableInsightToJSON)),
    };
}
