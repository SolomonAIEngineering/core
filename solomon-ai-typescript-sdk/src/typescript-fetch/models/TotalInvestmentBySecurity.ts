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
 * TotalInvestmentBySecurity
 * This message is used to represent the total investment of a security.
 * @export
 * @interface TotalInvestmentBySecurity
 */
export interface TotalInvestmentBySecurity {
    /**
     * 
     * @type {string}
     * @memberof TotalInvestmentBySecurity
     */
    securityId?: string;
    /**
     * 
     * @type {number}
     * @memberof TotalInvestmentBySecurity
     */
    totalInvestment?: number;
    /**
     * 
     * @type {string}
     * @memberof TotalInvestmentBySecurity
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof TotalInvestmentBySecurity
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the TotalInvestmentBySecurity interface.
 */
export function instanceOfTotalInvestmentBySecurity(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TotalInvestmentBySecurityFromJSON(json: any): TotalInvestmentBySecurity {
    return TotalInvestmentBySecurityFromJSONTyped(json, false);
}

export function TotalInvestmentBySecurityFromJSONTyped(json: any, ignoreDiscriminator: boolean): TotalInvestmentBySecurity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'securityId': !exists(json, 'securityId') ? undefined : json['securityId'],
        'totalInvestment': !exists(json, 'totalInvestment') ? undefined : json['totalInvestment'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function TotalInvestmentBySecurityToJSON(value?: TotalInvestmentBySecurity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'securityId': value.securityId,
        'totalInvestment': value.totalInvestment,
        'userId': value.userId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}
