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
 * @interface UpdateRecurringTransactionRequest
 */
export interface UpdateRecurringTransactionRequest {
    /**
     * 
     * @type {PlaidAccountRecurringTransaction}
     * @memberof UpdateRecurringTransactionRequest
     */
    transaction: PlaidAccountRecurringTransaction;
}

/**
 * Check if a given object implements the UpdateRecurringTransactionRequest interface.
 */
export function instanceOfUpdateRecurringTransactionRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "transaction" in value;

    return isInstance;
}

export function UpdateRecurringTransactionRequestFromJSON(json: any): UpdateRecurringTransactionRequest {
    return UpdateRecurringTransactionRequestFromJSONTyped(json, false);
}

export function UpdateRecurringTransactionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateRecurringTransactionRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transaction': PlaidAccountRecurringTransactionFromJSON(json['transaction']),
    };
}

export function UpdateRecurringTransactionRequestToJSON(value?: UpdateRecurringTransactionRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transaction': PlaidAccountRecurringTransactionToJSON(value.transaction),
    };
}
