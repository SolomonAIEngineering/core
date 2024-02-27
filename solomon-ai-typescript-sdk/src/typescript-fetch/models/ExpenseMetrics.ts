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
 * ExpenseMetrics
 * This message is used to represent the expense metrics of a user.
 * @export
 * @interface ExpenseMetrics
 */
export interface ExpenseMetrics {
    /**
     * 
     * @type {number}
     * @memberof ExpenseMetrics
     */
    month?: number;
    /**
     * 
     * @type {string}
     * @memberof ExpenseMetrics
     */
    personalFinanceCategoryPrimary?: string;
    /**
     * 
     * @type {string}
     * @memberof ExpenseMetrics
     */
    transactionCount?: string;
    /**
     * 
     * @type {number}
     * @memberof ExpenseMetrics
     */
    totalExpenses?: number;
    /**
     * 
     * @type {string}
     * @memberof ExpenseMetrics
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof ExpenseMetrics
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the ExpenseMetrics interface.
 */
export function instanceOfExpenseMetrics(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ExpenseMetricsFromJSON(json: any): ExpenseMetrics {
    return ExpenseMetricsFromJSONTyped(json, false);
}

export function ExpenseMetricsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExpenseMetrics {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'month': !exists(json, 'month') ? undefined : json['month'],
        'personalFinanceCategoryPrimary': !exists(json, 'personalFinanceCategoryPrimary') ? undefined : json['personalFinanceCategoryPrimary'],
        'transactionCount': !exists(json, 'transactionCount') ? undefined : json['transactionCount'],
        'totalExpenses': !exists(json, 'totalExpenses') ? undefined : json['totalExpenses'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function ExpenseMetricsToJSON(value?: ExpenseMetrics | null): any {
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
        'totalExpenses': value.totalExpenses,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}

