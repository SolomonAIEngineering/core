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
 * PaymentChannelMetricsFinancialSubProfile
 * This message is used to represent the financial sub profile of a payment channel.
 * @export
 * @interface PaymentChannelMetricsFinancialSubProfile
 */
export interface PaymentChannelMetricsFinancialSubProfile {
    /**
     * 
     * @type {string}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    paymentChannel?: string;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastWeek?: number;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastTwoWeeks?: number;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastMonth?: number;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastSixMonths?: number;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastYear?: number;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    spentLastTwoYears?: number;
    /**
     * 
     * @type {string}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    userId?: string;
    /**
     * 
     * @type {number}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    month?: number;
    /**
     * 
     * @type {string}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    transactionCount?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof PaymentChannelMetricsFinancialSubProfile
     */
    profileType?: FinancialUserProfileType;
}

/**
 * Check if a given object implements the PaymentChannelMetricsFinancialSubProfile interface.
 */
export function instanceOfPaymentChannelMetricsFinancialSubProfile(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PaymentChannelMetricsFinancialSubProfileFromJSON(json: any): PaymentChannelMetricsFinancialSubProfile {
    return PaymentChannelMetricsFinancialSubProfileFromJSONTyped(json, false);
}

export function PaymentChannelMetricsFinancialSubProfileFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaymentChannelMetricsFinancialSubProfile {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'paymentChannel': !exists(json, 'paymentChannel') ? undefined : json['paymentChannel'],
        'spentLastWeek': !exists(json, 'spentLastWeek') ? undefined : json['spentLastWeek'],
        'spentLastTwoWeeks': !exists(json, 'spentLastTwoWeeks') ? undefined : json['spentLastTwoWeeks'],
        'spentLastMonth': !exists(json, 'spentLastMonth') ? undefined : json['spentLastMonth'],
        'spentLastSixMonths': !exists(json, 'spentLastSixMonths') ? undefined : json['spentLastSixMonths'],
        'spentLastYear': !exists(json, 'spentLastYear') ? undefined : json['spentLastYear'],
        'spentLastTwoYears': !exists(json, 'spentLastTwoYears') ? undefined : json['spentLastTwoYears'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'month': !exists(json, 'month') ? undefined : json['month'],
        'transactionCount': !exists(json, 'transactionCount') ? undefined : json['transactionCount'],
        'profileType': !exists(json, 'profileType') ? undefined : FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function PaymentChannelMetricsFinancialSubProfileToJSON(value?: PaymentChannelMetricsFinancialSubProfile | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'paymentChannel': value.paymentChannel,
        'spentLastWeek': value.spentLastWeek,
        'spentLastTwoWeeks': value.spentLastTwoWeeks,
        'spentLastMonth': value.spentLastMonth,
        'spentLastSixMonths': value.spentLastSixMonths,
        'spentLastYear': value.spentLastYear,
        'spentLastTwoYears': value.spentLastTwoYears,
        'userId': value.userId,
        'month': value.month,
        'transactionCount': value.transactionCount,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
    };
}
