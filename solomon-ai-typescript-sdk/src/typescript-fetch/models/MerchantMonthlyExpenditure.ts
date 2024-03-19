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
 * MerchantMonthlyExpenditure
 * This message is used to represent the monthly expenditure of a merchant.
 * @export
 * @interface MerchantMonthlyExpenditure
 */
export interface MerchantMonthlyExpenditure {
    /**
     * 
     * @type {number}
     * @memberof MerchantMonthlyExpenditure
     */
    month?: number;
    /**
     * 
     * @type {string}
     * @memberof MerchantMonthlyExpenditure
     */
    merchantName?: string;
    /**
     * 
     * @type {number}
     * @memberof MerchantMonthlyExpenditure
     */
    totalSpending?: number;
    /**
     * 
     * @type {string}
     * @memberof MerchantMonthlyExpenditure
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof MerchantMonthlyExpenditure
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the MerchantMonthlyExpenditure interface.
 */
export function instanceOfMerchantMonthlyExpenditure(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MerchantMonthlyExpenditureFromJSON(json: any): MerchantMonthlyExpenditure {
    return MerchantMonthlyExpenditureFromJSONTyped(json, false);
}

export function MerchantMonthlyExpenditureFromJSONTyped(json: any, ignoreDiscriminator: boolean): MerchantMonthlyExpenditure {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'month': !exists(json, 'month') ? undefined : json['month'],
        'merchantName': !exists(json, 'merchantName') ? undefined : json['merchantName'],
        'totalSpending': !exists(json, 'totalSpending') ? undefined : json['totalSpending'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function MerchantMonthlyExpenditureToJSON(value?: MerchantMonthlyExpenditure | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'month': value.month,
        'merchantName': value.merchantName,
        'totalSpending': value.totalSpending,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}
