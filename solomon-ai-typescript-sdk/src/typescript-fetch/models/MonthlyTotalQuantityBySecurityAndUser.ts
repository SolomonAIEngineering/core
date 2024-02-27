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
 * MonthlyTotalQuantityBySecurityAndUser
 * This message is used to represent the monthly total quantity of a security.
 * @export
 * @interface MonthlyTotalQuantityBySecurityAndUser
 */
export interface MonthlyTotalQuantityBySecurityAndUser {
    /**
     * 
     * @type {number}
     * @memberof MonthlyTotalQuantityBySecurityAndUser
     */
    month?: number;
    /**
     * 
     * @type {string}
     * @memberof MonthlyTotalQuantityBySecurityAndUser
     */
    securityId?: string;
    /**
     * 
     * @type {number}
     * @memberof MonthlyTotalQuantityBySecurityAndUser
     */
    totalQuantity?: number;
    /**
     * 
     * @type {string}
     * @memberof MonthlyTotalQuantityBySecurityAndUser
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof MonthlyTotalQuantityBySecurityAndUser
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the MonthlyTotalQuantityBySecurityAndUser interface.
 */
export function instanceOfMonthlyTotalQuantityBySecurityAndUser(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MonthlyTotalQuantityBySecurityAndUserFromJSON(json: any): MonthlyTotalQuantityBySecurityAndUser {
    return MonthlyTotalQuantityBySecurityAndUserFromJSONTyped(json, false);
}

export function MonthlyTotalQuantityBySecurityAndUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonthlyTotalQuantityBySecurityAndUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'month': !exists(json, 'month') ? undefined : json['month'],
        'securityId': !exists(json, 'securityId') ? undefined : json['securityId'],
        'totalQuantity': !exists(json, 'totalQuantity') ? undefined : json['totalQuantity'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function MonthlyTotalQuantityBySecurityAndUserToJSON(value?: MonthlyTotalQuantityBySecurityAndUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'month': value.month,
        'securityId': value.securityId,
        'totalQuantity': value.totalQuantity,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}

