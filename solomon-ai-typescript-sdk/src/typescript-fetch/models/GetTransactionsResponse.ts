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
import type { Transaction } from './Transaction';
import {
    TransactionFromJSON,
    TransactionFromJSONTyped,
    TransactionToJSON,
} from './Transaction';

/**
 * 
 * @export
 * @interface GetTransactionsResponse
 */
export interface GetTransactionsResponse {
    /**
     * 
     * @type {Array<Transaction>}
     * @memberof GetTransactionsResponse
     */
    transactions?: Array<Transaction>;
    /**
     * 
     * @type {string}
     * @memberof GetTransactionsResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetTransactionsResponse interface.
 */
export function instanceOfGetTransactionsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetTransactionsResponseFromJSON(json: any): GetTransactionsResponse {
    return GetTransactionsResponseFromJSONTyped(json, false);
}

export function GetTransactionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetTransactionsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(TransactionFromJSON)),
        'nextPageNumber': !exists(json, 'nextPageNumber') ? undefined : json['nextPageNumber'],
    };
}

export function GetTransactionsResponseToJSON(value?: GetTransactionsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(TransactionToJSON)),
        'nextPageNumber': value.nextPageNumber,
    };
}

