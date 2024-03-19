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
import type { PlaidAccountTransaction } from './PlaidAccountTransaction';
import {
    PlaidAccountTransactionFromJSON,
    PlaidAccountTransactionFromJSONTyped,
    PlaidAccountTransactionToJSON,
} from './PlaidAccountTransaction';

/**
 * 
 * @export
 * @interface SearchTransactionsResponse
 */
export interface SearchTransactionsResponse {
    /**
     * 
     * @type {Array<PlaidAccountTransaction>}
     * @memberof SearchTransactionsResponse
     */
    transactions?: Array<PlaidAccountTransaction>;
    /**
     * 
     * @type {string}
     * @memberof SearchTransactionsResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the SearchTransactionsResponse interface.
 */
export function instanceOfSearchTransactionsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SearchTransactionsResponseFromJSON(json: any): SearchTransactionsResponse {
    return SearchTransactionsResponseFromJSONTyped(json, false);
}

export function SearchTransactionsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchTransactionsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountTransactionFromJSON)),
        'nextPageNumber': !exists(json, 'nextPageNumber') ? undefined : json['nextPageNumber'],
    };
}

export function SearchTransactionsResponseToJSON(value?: SearchTransactionsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(PlaidAccountTransactionToJSON)),
        'nextPageNumber': value.nextPageNumber,
    };
}
