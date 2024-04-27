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
import type { Account } from './Account';
import {
    AccountFromJSON,
    AccountFromJSONTyped,
    AccountToJSON,
} from './Account';

/**
 * 
 * @export
 * @interface GetAccountResponse
 */
export interface GetAccountResponse {
    /**
     * 
     * @type {Account}
     * @memberof GetAccountResponse
     */
    account?: Account;
}

/**
 * Check if a given object implements the GetAccountResponse interface.
 */
export function instanceOfGetAccountResponse(value: object): boolean {
    return true;
}

export function GetAccountResponseFromJSON(json: any): GetAccountResponse {
    return GetAccountResponseFromJSONTyped(json, false);
}

export function GetAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetAccountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'account': json['account'] == null ? undefined : AccountFromJSON(json['account']),
    };
}

export function GetAccountResponseToJSON(value?: GetAccountResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'account': AccountToJSON(value['account']),
    };
}

