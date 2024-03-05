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
 * @interface BulkUpdateTransactionResponse
 */
export interface BulkUpdateTransactionResponse {
    /**
     * 
     * @type {Array<PlaidAccountTransaction>}
     * @memberof BulkUpdateTransactionResponse
     */
    transactions?: Array<PlaidAccountTransaction>;
}

/**
 * Check if a given object implements the BulkUpdateTransactionResponse interface.
 */
export function instanceOfBulkUpdateTransactionResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BulkUpdateTransactionResponseFromJSON(json: any): BulkUpdateTransactionResponse {
    return BulkUpdateTransactionResponseFromJSONTyped(json, false);
}

export function BulkUpdateTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BulkUpdateTransactionResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountTransactionFromJSON)),
    };
}

export function BulkUpdateTransactionResponseToJSON(value?: BulkUpdateTransactionResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(PlaidAccountTransactionToJSON)),
    };
}
