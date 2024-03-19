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
import type { FinancialUserProfileType } from './FinancialUserProfileType';
import {
    FinancialUserProfileTypeFromJSON,
    FinancialUserProfileTypeFromJSONTyped,
    FinancialUserProfileTypeToJSON,
} from './FinancialUserProfileType';

/**
 * MerchantFinancialSubProfile
 * This message is used to represent the financial sub profile of a merchant.
 * @export
 * @interface MerchantMetricsFinancialSubProfile
 */
export interface MerchantMetricsFinancialSubProfile {
    /**
     * 
     * @type {string}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    merchantName?: string;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastWeek?: number;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastTwoWeeks?: number;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastMonth?: number;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastSixMonths?: number;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastYear?: number;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    spentLastTwoYears?: number;
    /**
     * 
     * @type {string}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    userId?: string;
    /**
     * 
     * @type {number}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    month?: number;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof MerchantMetricsFinancialSubProfile
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the MerchantMetricsFinancialSubProfile interface.
 */
export function instanceOfMerchantMetricsFinancialSubProfile(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MerchantMetricsFinancialSubProfileFromJSON(json: any): MerchantMetricsFinancialSubProfile {
    return MerchantMetricsFinancialSubProfileFromJSONTyped(json, false);
}

export function MerchantMetricsFinancialSubProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): MerchantMetricsFinancialSubProfile {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'merchantName': !exists(json, 'merchantName') ? undefined : json['merchantName'],
        'spentLastWeek': !exists(json, 'spentLastWeek') ? undefined : json['spentLastWeek'],
        'spentLastTwoWeeks': !exists(json, 'spentLastTwoWeeks') ? undefined : json['spentLastTwoWeeks'],
        'spentLastMonth': !exists(json, 'spentLastMonth') ? undefined : json['spentLastMonth'],
        'spentLastSixMonths': !exists(json, 'spentLastSixMonths') ? undefined : json['spentLastSixMonths'],
        'spentLastYear': !exists(json, 'spentLastYear') ? undefined : json['spentLastYear'],
        'spentLastTwoYears': !exists(json, 'spentLastTwoYears') ? undefined : json['spentLastTwoYears'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'month': !exists(json, 'month') ? undefined : json['month'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function MerchantMetricsFinancialSubProfileToJSON(value?: MerchantMetricsFinancialSubProfile | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'merchantName': value.merchantName,
        'spentLastWeek': value.spentLastWeek,
        'spentLastTwoWeeks': value.spentLastTwoWeeks,
        'spentLastMonth': value.spentLastMonth,
        'spentLastSixMonths': value.spentLastSixMonths,
        'spentLastYear': value.spentLastYear,
        'spentLastTwoYears': value.spentLastTwoYears,
        'userId': value.userId,
        'month': value.month,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}
