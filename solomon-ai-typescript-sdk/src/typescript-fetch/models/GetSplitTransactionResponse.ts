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
import type { TransactionSplit } from './TransactionSplit';
import {
    TransactionSplitFromJSON,
    TransactionSplitFromJSONTyped,
    TransactionSplitToJSON,
} from './TransactionSplit';

/**
 * 
 * @export
 * @interface GetSplitTransactionResponse
 */
export interface GetSplitTransactionResponse {
    /**
     * 
     * @type {Array<TransactionSplit>}
     * @memberof GetSplitTransactionResponse
     */
    splitTransactions?: Array<TransactionSplit>;
}

/**
 * Check if a given object implements the GetSplitTransactionResponse interface.
 */
export function instanceOfGetSplitTransactionResponse(value: object): boolean {
    return true;
}

export function GetSplitTransactionResponseFromJSON(json: any): GetSplitTransactionResponse {
    return GetSplitTransactionResponseFromJSONTyped(json, false);
}

export function GetSplitTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetSplitTransactionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'splitTransactions': json['splitTransactions'] == null ? undefined : ((json['splitTransactions'] as Array<any>).map(TransactionSplitFromJSON)),
    };
}

export function GetSplitTransactionResponseToJSON(value?: GetSplitTransactionResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'splitTransactions': value['splitTransactions'] == null ? undefined : ((value['splitTransactions'] as Array<any>).map(TransactionSplitToJSON)),
    };
}

