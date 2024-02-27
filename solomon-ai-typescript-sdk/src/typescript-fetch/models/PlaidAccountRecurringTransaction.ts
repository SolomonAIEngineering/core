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
import type { Any1 } from './Any1';
import {
    Any1FromJSON,
    Any1FromJSONTyped,
    Any1ToJSON,
} from './Any1';
import type { SmartNote } from './SmartNote';
import {
    SmartNoteFromJSON,
    SmartNoteFromJSONTyped,
    SmartNoteToJSON,
} from './SmartNote';

/**
 * Message representing recurring transactions associated with a Plaid account.
 * @export
 * @interface PlaidAccountRecurringTransaction
 */
export interface PlaidAccountRecurringTransaction {
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    accountId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    streamId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    categoryId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    merchantName?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    personalFinanceCategoryPrimary?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    personalFinanceCategoryDetailed?: string;
    /**
     * 
     * @type {Date}
     * @memberof PlaidAccountRecurringTransaction
     */
    firstDate?: Date;
    /**
     * 
     * @type {Date}
     * @memberof PlaidAccountRecurringTransaction
     */
    lastDate?: Date;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    frequency?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    transactionIds?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    averageAmount?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    averageAmountIsoCurrencyCode?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    lastAmount?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    lastAmountIsoCurrencyCode?: string;
    /**
     * 
     * @type {boolean}
     * @memberof PlaidAccountRecurringTransaction
     */
    isActive?: boolean;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    status?: string;
    /**
     * 
     * @type {Date}
     * @memberof PlaidAccountRecurringTransaction
     */
    updatedTime?: Date;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    userId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    linkId?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof PlaidAccountRecurringTransaction
     */
    flow?: string;
    /**
     * The timestamp associated with this recurring transaction.
     * @type {Date}
     * @memberof PlaidAccountRecurringTransaction
     */
    time?: Date;
    /**
     * 
     * @type {Any1}
     * @memberof PlaidAccountRecurringTransaction
     */
    additionalProperties?: Any1;
    /**
     * Notes associated with this recurring transaction.
     * @type {Array<SmartNote>}
     * @memberof PlaidAccountRecurringTransaction
     */
    notes?: Array<SmartNote>;
}

/**
 * Check if a given object implements the PlaidAccountRecurringTransaction interface.
 */
export function instanceOfPlaidAccountRecurringTransaction(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PlaidAccountRecurringTransactionFromJSON(json: any): PlaidAccountRecurringTransaction {
    return PlaidAccountRecurringTransactionFromJSONTyped(json, false);
}

export function PlaidAccountRecurringTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlaidAccountRecurringTransaction {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accountId': !exists(json, 'accountId') ? undefined : json['accountId'],
        'streamId': !exists(json, 'streamId') ? undefined : json['streamId'],
        'categoryId': !exists(json, 'categoryId') ? undefined : json['categoryId'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'merchantName': !exists(json, 'merchantName') ? undefined : json['merchantName'],
        'personalFinanceCategoryPrimary': !exists(json, 'personalFinanceCategoryPrimary') ? undefined : json['personalFinanceCategoryPrimary'],
        'personalFinanceCategoryDetailed': !exists(json, 'personalFinanceCategoryDetailed') ? undefined : json['personalFinanceCategoryDetailed'],
        'firstDate': !exists(json, 'firstDate') ? undefined : (new Date(json['firstDate'])),
        'lastDate': !exists(json, 'lastDate') ? undefined : (new Date(json['lastDate'])),
        'frequency': !exists(json, 'frequency') ? undefined : json['frequency'],
        'transactionIds': !exists(json, 'transactionIds') ? undefined : json['transactionIds'],
        'averageAmount': !exists(json, 'averageAmount') ? undefined : json['averageAmount'],
        'averageAmountIsoCurrencyCode': !exists(json, 'averageAmountIsoCurrencyCode') ? undefined : json['averageAmountIsoCurrencyCode'],
        'lastAmount': !exists(json, 'lastAmount') ? undefined : json['lastAmount'],
        'lastAmountIsoCurrencyCode': !exists(json, 'lastAmountIsoCurrencyCode') ? undefined : json['lastAmountIsoCurrencyCode'],
        'isActive': !exists(json, 'isActive') ? undefined : json['isActive'],
        'status': !exists(json, 'status') ? undefined : json['status'],
        'updatedTime': !exists(json, 'updatedTime') ? undefined : (new Date(json['updatedTime'])),
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'linkId': !exists(json, 'linkId') ? undefined : json['linkId'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'flow': !exists(json, 'flow') ? undefined : json['flow'],
        'time': !exists(json, 'time') ? undefined : (new Date(json['time'])),
        'additionalProperties': !exists(json, 'additionalProperties') ? undefined : Any1FromJSON(json['additionalProperties']),
        'notes': !exists(json, 'notes') ? undefined : ((json['notes'] as Array<any>).map(SmartNoteFromJSON)),
    };
}

export function PlaidAccountRecurringTransactionToJSON(value?: PlaidAccountRecurringTransaction | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accountId': value.accountId,
        'streamId': value.streamId,
        'categoryId': value.categoryId,
        'description': value.description,
        'merchantName': value.merchantName,
        'personalFinanceCategoryPrimary': value.personalFinanceCategoryPrimary,
        'personalFinanceCategoryDetailed': value.personalFinanceCategoryDetailed,
        'firstDate': value.firstDate === undefined ? undefined : (value.firstDate.toISOString()),
        'lastDate': value.lastDate === undefined ? undefined : (value.lastDate.toISOString()),
        'frequency': value.frequency,
        'transactionIds': value.transactionIds,
        'averageAmount': value.averageAmount,
        'averageAmountIsoCurrencyCode': value.averageAmountIsoCurrencyCode,
        'lastAmount': value.lastAmount,
        'lastAmountIsoCurrencyCode': value.lastAmountIsoCurrencyCode,
        'isActive': value.isActive,
        'status': value.status,
        'updatedTime': value.updatedTime === undefined ? undefined : (value.updatedTime.toISOString()),
        'userId': value.userId,
        'linkId': value.linkId,
        'id': value.id,
        'flow': value.flow,
        'time': value.time === undefined ? undefined : (value.time.toISOString()),
        'additionalProperties': Any1ToJSON(value.additionalProperties),
        'notes': value.notes === undefined ? undefined : ((value.notes as Array<any>).map(SmartNoteToJSON)),
    };
}

