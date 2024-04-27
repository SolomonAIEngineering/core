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
 * LocationFinancialSubProfile
 * This message is used to represent the financial sub profile of a location.
 * @export
 * @interface LocationFinancialSubProfile
 */
export interface LocationFinancialSubProfile {
    /**
     * 
     * @type {string}
     * @memberof LocationFinancialSubProfile
     */
    locationCity?: string;
    /**
     * 
     * @type {string}
     * @memberof LocationFinancialSubProfile
     */
    transactionCount?: string;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastWeek?: number;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastTwoWeeks?: number;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastMonth?: number;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastSixMonths?: number;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastYear?: number;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    spentLastTwoYears?: number;
    /**
     * 
     * @type {string}
     * @memberof LocationFinancialSubProfile
     */
    userId?: string;
    /**
     * 
     * @type {number}
     * @memberof LocationFinancialSubProfile
     */
    month?: number;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof LocationFinancialSubProfile
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the LocationFinancialSubProfile interface.
 */
export function instanceOfLocationFinancialSubProfile(value: object): boolean {
    return true;
}

export function LocationFinancialSubProfileFromJSON(json: any): LocationFinancialSubProfile {
    return LocationFinancialSubProfileFromJSONTyped(json, false);
}

export function LocationFinancialSubProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): LocationFinancialSubProfile {
    if (json == null) {
        return json;
    }
    return {
        
        'locationCity': json['locationCity'] == null ? undefined : json['locationCity'],
        'transactionCount': json['transactionCount'] == null ? undefined : json['transactionCount'],
        'spentLastWeek': json['spentLastWeek'] == null ? undefined : json['spentLastWeek'],
        'spentLastTwoWeeks': json['spentLastTwoWeeks'] == null ? undefined : json['spentLastTwoWeeks'],
        'spentLastMonth': json['spentLastMonth'] == null ? undefined : json['spentLastMonth'],
        'spentLastSixMonths': json['spentLastSixMonths'] == null ? undefined : json['spentLastSixMonths'],
        'spentLastYear': json['spentLastYear'] == null ? undefined : json['spentLastYear'],
        'spentLastTwoYears': json['spentLastTwoYears'] == null ? undefined : json['spentLastTwoYears'],
        'userId': json['userId'] == null ? undefined : json['userId'],
        'month': json['month'] == null ? undefined : json['month'],
        'profileType': json['profileType'] == null ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function LocationFinancialSubProfileToJSON(value?: LocationFinancialSubProfile | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'locationCity': value['locationCity'],
        'transactionCount': value['transactionCount'],
        'spentLastWeek': value['spentLastWeek'],
        'spentLastTwoWeeks': value['spentLastTwoWeeks'],
        'spentLastMonth': value['spentLastMonth'],
        'spentLastSixMonths': value['spentLastSixMonths'],
        'spentLastYear': value['spentLastYear'],
        'spentLastTwoYears': value['spentLastTwoYears'],
        'userId': value['userId'],
        'month': value['month'],
        'profileType': FinancialUserProfileTypeToJSON(value['profileType']),
    };
}

