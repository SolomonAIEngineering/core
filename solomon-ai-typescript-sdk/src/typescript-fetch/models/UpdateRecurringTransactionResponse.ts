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
 * @interface UpdateRecurringTransactionResponse
 */
export interface UpdateRecurringTransactionResponse {
    /**
     * 
     * @type {PlaidAccountRecurringTransaction}
     * @memberof UpdateRecurringTransactionResponse
     */
    transaction?: PlaidAccountRecurringTransaction;
}

/**
 * Check if a given object implements the UpdateRecurringTransactionResponse interface.
 */
export function instanceOfUpdateRecurringTransactionResponse(value: object): boolean {
    return true;
}

export function UpdateRecurringTransactionResponseFromJSON(json: any): UpdateRecurringTransactionResponse {
    return UpdateRecurringTransactionResponseFromJSONTyped(json, false);
}

export function UpdateRecurringTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateRecurringTransactionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'transaction': json['transaction'] == null ? undefined : PlaidAccountRecurringTransactionFromJSON(json['transaction']),
    };
}

export function UpdateRecurringTransactionResponseToJSON(value?: UpdateRecurringTransactionResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'transaction': PlaidAccountRecurringTransactionToJSON(value['transaction']),
    };
}

