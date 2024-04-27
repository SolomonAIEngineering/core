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
import type { FinancialUserProfileType } from './FinancialUserProfileType';
import {
    FinancialUserProfileTypeFromJSON,
    FinancialUserProfileTypeFromJSONTyped,
    FinancialUserProfileTypeToJSON,
} from './FinancialUserProfileType';

/**
 * MonthlySavings
 * This message is used to represent the monthly savings of a user.
 * @export
 * @interface MonthlySavings
 */
export interface MonthlySavings {
    /**
     * 
     * @type {number}
     * @memberof MonthlySavings
     */
    month?: number;
    /**
     * 
     * @type {number}
     * @memberof MonthlySavings
     */
    netSavings?: number;
    /**
     * 
     * @type {string}
     * @memberof MonthlySavings
     */
    userId?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof MonthlySavings
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the MonthlySavings interface.
 */
export function instanceOfMonthlySavings(value: object): boolean {
    return true;
}

export function MonthlySavingsFromJSON(json: any): MonthlySavings {
    return MonthlySavingsFromJSONTyped(json, false);
}

export function MonthlySavingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): MonthlySavings {
    if (json == null) {
        return json;
    }
    return {
        
        'month': json['month'] == null ? undefined : json['month'],
        'netSavings': json['netSavings'] == null ? undefined : json['netSavings'],
        'userId': json['userId'] == null ? undefined : json['userId'],
        'profileType': json['profileType'] == null ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function MonthlySavingsToJSON(value?: MonthlySavings | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'month': value['month'],
        'netSavings': value['netSavings'],
        'userId': value['userId'],
        'profileType': FinancialUserProfileTypeToJSON(value['profileType']),
    };
}

