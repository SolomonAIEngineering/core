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
 * 
 * @export
 * @interface PlaidExchangeTokenRequest
 */
export interface PlaidExchangeTokenRequest {
    /**
     * 
     * @type {string}
     * @memberof PlaidExchangeTokenRequest
     */
    userId: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidExchangeTokenRequest
     */
    publicToken: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidExchangeTokenRequest
     */
    institutionId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidExchangeTokenRequest
     */
    institutionName?: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof PlaidExchangeTokenRequest
     */
    profileType: FinancialUserProfileType;
}

/**
 * Check if a given object implements the PlaidExchangeTokenRequest interface.
 */
export function instanceOfPlaidExchangeTokenRequest(value: object): boolean {
    if (!('userId' in value)) return false;
    if (!('publicToken' in value)) return false;
    if (!('profileType' in value)) return false;
    return true;
}

export function PlaidExchangeTokenRequestFromJSON(json: any): PlaidExchangeTokenRequest {
    return PlaidExchangeTokenRequestFromJSONTyped(json, false);
}

export function PlaidExchangeTokenRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlaidExchangeTokenRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'userId': json['userId'],
        'publicToken': json['publicToken'],
        'institutionId': json['institutionId'] == null ? undefined : json['institutionId'],
        'institutionName': json['institutionName'] == null ? undefined : json['institutionName'],
        'profileType': FinancialUserProfileTypeFromJSON(json['profileType']),
    };
}

export function PlaidExchangeTokenRequestToJSON(value?: PlaidExchangeTokenRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'userId': value['userId'],
        'publicToken': value['publicToken'],
        'institutionId': value['institutionId'],
        'institutionName': value['institutionName'],
        'profileType': FinancialUserProfileTypeToJSON(value['profileType']),
    };
}

