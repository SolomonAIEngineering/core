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
import type { Transaction } from './Transaction';
import {
    TransactionFromJSON,
    TransactionFromJSONTyped,
    TransactionToJSON,
} from './Transaction';

/**
 * 
 * @export
 * @interface GetTransactionsForBankAccountResponse
 */
export interface GetTransactionsForBankAccountResponse {
    /**
     * 
     * @type {Array<Transaction>}
     * @memberof GetTransactionsForBankAccountResponse
     */
    transactions?: Array<Transaction>;
    /**
     * 
     * @type {string}
     * @memberof GetTransactionsForBankAccountResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetTransactionsForBankAccountResponse interface.
 */
export function instanceOfGetTransactionsForBankAccountResponse(value: object): boolean {
    return true;
}

export function GetTransactionsForBankAccountResponseFromJSON(json: any): GetTransactionsForBankAccountResponse {
    return GetTransactionsForBankAccountResponseFromJSONTyped(json, false);
}

export function GetTransactionsForBankAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetTransactionsForBankAccountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'transactions': json['transactions'] == null ? undefined : ((json['transactions'] as Array<any>).map(TransactionFromJSON)),
        'nextPageNumber': json['nextPageNumber'] == null ? undefined : json['nextPageNumber'],
    };
}

export function GetTransactionsForBankAccountResponseToJSON(value?: GetTransactionsForBankAccountResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'transactions': value['transactions'] == null ? undefined : ((value['transactions'] as Array<any>).map(TransactionToJSON)),
        'nextPageNumber': value['nextPageNumber'],
    };
}

