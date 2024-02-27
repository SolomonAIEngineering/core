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
 * MonthlyBalance
 * This message is used to represent the monthly balance of a user.
 * @export
 * @interface MonthlyBalance
 */
export interface MonthlyBalance {
    /**
     * 
     * @type {number}
     * @memberof MonthlyBalance
     */
    month?: number;
    /**
     * 
     * @type {number}
     * @memberof MonthlyBalance
     */
    netBalance?: number;
    /**
     * 
     * @type {string}
     * @memberof MonthlyBalance
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof MonthlyBalance
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the MonthlyBalance interface.
 */
export function instanceOfMonthlyBalance(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MonthlyBalanceFromJSON(json: any): MonthlyBalance {
    return MonthlyBalanceFromJSONTyped(json, false);
}

export function MonthlyBalanceFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonthlyBalance {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'month': !exists(json, 'month') ? undefined : json['month'],
        'netBalance': !exists(json, 'netBalance') ? undefined : json['netBalance'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function MonthlyBalanceToJSON(value?: MonthlyBalance | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'month': value.month,
        'netBalance': value.netBalance,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}

