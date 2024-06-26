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
import type { BankAccount } from './BankAccount';
import {
    BankAccountFromJSON,
    BankAccountFromJSONTyped,
    BankAccountToJSON,
} from './BankAccount';

/**
 * 
 * @export
 * @interface GetBankAccountResponse
 */
export interface GetBankAccountResponse {
    /**
     * 
     * @type {BankAccount}
     * @memberof GetBankAccountResponse
     */
    bankAccount?: BankAccount;
}

/**
 * Check if a given object implements the GetBankAccountResponse interface.
 */
export function instanceOfGetBankAccountResponse(value: object): boolean {
    return true;
}

export function GetBankAccountResponseFromJSON(json: any): GetBankAccountResponse {
    return GetBankAccountResponseFromJSONTyped(json, false);
}

export function GetBankAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetBankAccountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'bankAccount': json['bankAccount'] == null ? undefined : BankAccountFromJSON(json['bankAccount']),
    };
}

export function GetBankAccountResponseToJSON(value?: GetBankAccountResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'bankAccount': BankAccountToJSON(value['bankAccount']),
    };
}

