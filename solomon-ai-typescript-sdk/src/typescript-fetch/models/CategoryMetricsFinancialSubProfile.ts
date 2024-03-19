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
 * CategoryMetricsFinancialSubProfile
 * This message is used to represent the financial sub profile of a category.
 * @export
 * @interface CategoryMetricsFinancialSubProfile
 */
export interface CategoryMetricsFinancialSubProfile {
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    month?: number;
    /**
     * 
     * @type {string}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    personalFinanceCategoryPrimary?: string;
    /**
     * 
     * @type {string}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    transactionCount?: string;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastWeek?: number;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastTwoWeeks?: number;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastMonth?: number;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastSixMonths?: number;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastYear?: number;
    /**
     * 
     * @type {number}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    spentLastTwoYears?: number;
    /**
     * 
     * @type {string}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof CategoryMetricsFinancialSubProfile
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the CategoryMetricsFinancialSubProfile interface.
 */
export function instanceOfCategoryMetricsFinancialSubProfile(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CategoryMetricsFinancialSubProfileFromJSON(json: any): CategoryMetricsFinancialSubProfile {
    return CategoryMetricsFinancialSubProfileFromJSONTyped(json, false);
}

export function CategoryMetricsFinancialSubProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): CategoryMetricsFinancialSubProfile {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'month': !exists(json, 'month') ? undefined : json['month'],
        'personalFinanceCategoryPrimary': !exists(json, 'personalFinanceCategoryPrimary') ? undefined : json['personalFinanceCategoryPrimary'],
        'transactionCount': !exists(json, 'transactionCount') ? undefined : json['transactionCount'],
        'spentLastWeek': !exists(json, 'spentLastWeek') ? undefined : json['spentLastWeek'],
        'spentLastTwoWeeks': !exists(json, 'spentLastTwoWeeks') ? undefined : json['spentLastTwoWeeks'],
        'spentLastMonth': !exists(json, 'spentLastMonth') ? undefined : json['spentLastMonth'],
        'spentLastSixMonths': !exists(json, 'spentLastSixMonths') ? undefined : json['spentLastSixMonths'],
        'spentLastYear': !exists(json, 'spentLastYear') ? undefined : json['spentLastYear'],
        'spentLastTwoYears': !exists(json, 'spentLastTwoYears') ? undefined : json['spentLastTwoYears'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function CategoryMetricsFinancialSubProfileToJSON(value?: CategoryMetricsFinancialSubProfile | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'month': value.month,
        'personalFinanceCategoryPrimary': value.personalFinanceCategoryPrimary,
        'transactionCount': value.transactionCount,
        'spentLastWeek': value.spentLastWeek,
        'spentLastTwoWeeks': value.spentLastTwoWeeks,
        'spentLastMonth': value.spentLastMonth,
        'spentLastSixMonths': value.spentLastSixMonths,
        'spentLastYear': value.spentLastYear,
        'spentLastTwoYears': value.spentLastTwoYears,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}
