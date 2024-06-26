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
import type { PlaidAccountRecurringTransaction } from './PlaidAccountRecurringTransaction';
import {
    PlaidAccountRecurringTransactionFromJSON,
    PlaidAccountRecurringTransactionFromJSONTyped,
    PlaidAccountRecurringTransactionToJSON,
} from './PlaidAccountRecurringTransaction';

/**
 * 
 * @export
 * @interface GetRecurringTransactionsForUserResponse
 */
export interface GetRecurringTransactionsForUserResponse {
    /**
     * 
     * @type {Array<PlaidAccountRecurringTransaction>}
     * @memberof GetRecurringTransactionsForUserResponse
     */
    transactions?: Array<PlaidAccountRecurringTransaction>;
    /**
     * 
     * @type {string}
     * @memberof GetRecurringTransactionsForUserResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetRecurringTransactionsForUserResponse interface.
 */
export function instanceOfGetRecurringTransactionsForUserResponse(value: object): boolean {
    return true;
}

export function GetRecurringTransactionsForUserResponseFromJSON(json: any): GetRecurringTransactionsForUserResponse {
    return GetRecurringTransactionsForUserResponseFromJSONTyped(json, false);
}

export function GetRecurringTransactionsForUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetRecurringTransactionsForUserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'transactions': json['transactions'] == null ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountRecurringTransactionFromJSON)),
        'nextPageNumber': json['nextPageNumber'] == null ? undefined : json['nextPageNumber'],
    };
}

export function GetRecurringTransactionsForUserResponseToJSON(value?: GetRecurringTransactionsForUserResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'transactions': value['transactions'] == null ? undefined : ((value['transactions'] as Array<any>).map(PlaidAccountRecurringTransactionToJSON)),
        'nextPageNumber': value['nextPageNumber'],
    };
}

