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
import type { PlaidAccountRecurringTransaction } from './PlaidAccountRecurringTransaction';
import {
    PlaidAccountRecurringTransactionFromJSON,
    PlaidAccountRecurringTransactionFromJSONTyped,
    PlaidAccountRecurringTransactionToJSON,
} from './PlaidAccountRecurringTransaction';

/**
 * 
 * @export
 * @interface BulkUpdateRecurringTransactionRequest
 */
export interface BulkUpdateRecurringTransactionRequest {
    /**
     * 
     * @type {Array<PlaidAccountRecurringTransaction>}
     * @memberof BulkUpdateRecurringTransactionRequest
     */
    transactions: Array<PlaidAccountRecurringTransaction>;
}

/**
 * Check if a given object implements the BulkUpdateRecurringTransactionRequest interface.
 */
export function instanceOfBulkUpdateRecurringTransactionRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "transactions" in value;

    return isInstance;
}

export function BulkUpdateRecurringTransactionRequestFromJSON(json: any): BulkUpdateRecurringTransactionRequest {
    return BulkUpdateRecurringTransactionRequestFromJSONTyped(json, false);
}

export function BulkUpdateRecurringTransactionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): BulkUpdateRecurringTransactionRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transactions': ((json['transactions'] as Array<any>).map(PlaidAccountRecurringTransactionFromJSON)),
    };
}

export function BulkUpdateRecurringTransactionRequestToJSON(value?: BulkUpdateRecurringTransactionRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transactions': ((value.transactions as Array<any>).map(PlaidAccountRecurringTransactionToJSON)),
    };
}
