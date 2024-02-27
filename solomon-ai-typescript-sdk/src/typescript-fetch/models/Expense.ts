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
import type { ExpenseLine } from './ExpenseLine';
import {
    ExpenseLineFromJSON,
    ExpenseLineFromJSONTyped,
    ExpenseLineToJSON,
} from './ExpenseLine';

/**
 * The Expense object is used to represent a direct purchase by a business, typically made with a check, credit card, or cash. 
 * Each Expense object is dedicated to a grouping of expenses, with each expense recorded in the lines object.
 * 
 * The Expense object is used also used to represent refunds to direct purchases. Refunds can be distinguished from purchases 
 * by the amount sign of the records. Expense objects with a negative amount are purchases and Expense objects 
 * with a positive amount are refunds to those purchases.
 * @export
 * @interface Expense
 */
export interface Expense {
    /**
     * 
     * @type {string}
     * @memberof Expense
     */
    id?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof Expense
     */
    remoteId?: string;
    /**
     * When the transaction occurred.
     * 
     * Consider using google.protobuf.Timestamp
     * @type {Date}
     * @memberof Expense
     */
    transactionDate?: Date;
    /**
     * When the expense was created.
     * 
     * Consider using google.protobuf.Timestamp
     * @type {Date}
     * @memberof Expense
     */
    remoteCreatedAt?: Date;
    /**
     * The expense's payment account.
     * @type {string}
     * @memberof Expense
     */
    account?: string;
    /**
     * The expense's contact.
     * @type {string}
     * @memberof Expense
     */
    contact?: string;
    /**
     * The expense's total amount.
     * @type {number}
     * @memberof Expense
     */
    totalAmount?: number;
    /**
     * The expense's total amount before tax.
     * @type {number}
     * @memberof Expense
     */
    subTotal?: number;
    /**
     * The expense's total tax amount.
     * @type {number}
     * @memberof Expense
     */
    totalTaxAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof Expense
     */
    currency?: string;
    /**
     * The expense's exchange rate.
     * 
     * Consider using double or float if this represents a number
     * @type {string}
     * @memberof Expense
     */
    exchangeRate?: string;
    /**
     * The company the expense belongs to.
     * @type {string}
     * @memberof Expense
     */
    company?: string;
    /**
     * The expense's private note.
     * @type {string}
     * @memberof Expense
     */
    memo?: string;
    /**
     * The ExpenseLine object is used to represent an expense's line items.
     * @type {Array<ExpenseLine>}
     * @memberof Expense
     */
    lines?: Array<ExpenseLine>;
    /**
     * 
     * @type {Array<string>}
     * @memberof Expense
     */
    trackingCategories?: Array<string>;
    /**
     * Indicates whether or not this object has been deleted by third party webhooks.
     * @type {boolean}
     * @memberof Expense
     */
    remoteWasDeleted?: boolean;
    /**
     * The accounting period that the Expense was generated in.
     * @type {string}
     * @memberof Expense
     */
    accountingPeriod?: string;
    /**
     * Consider using google.protobuf.Timestamp
     * @type {Date}
     * @memberof Expense
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof Expense
     */
    mergeRecordId?: string;
}

/**
 * Check if a given object implements the Expense interface.
 */
export function instanceOfExpense(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ExpenseFromJSON(json: any): Expense {
    return ExpenseFromJSONTyped(json, false);
}

export function ExpenseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Expense {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'transactionDate': !exists(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'remoteCreatedAt': !exists(json, 'remoteCreatedAt') ? undefined : (new Date(json['remoteCreatedAt'])),
        'account': !exists(json, 'account') ? undefined : json['account'],
        'contact': !exists(json, 'contact') ? undefined : json['contact'],
        'totalAmount': !exists(json, 'totalAmount') ? undefined : json['totalAmount'],
        'subTotal': !exists(json, 'subTotal') ? undefined : json['subTotal'],
        'totalTaxAmount': !exists(json, 'totalTaxAmount') ? undefined : json['totalTaxAmount'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : json['exchangeRate'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'memo': !exists(json, 'memo') ? undefined : json['memo'],
        'lines': !exists(json, 'lines') ? undefined : ((json['lines'] as Array<any>).map(ExpenseLineFromJSON)),
        'trackingCategories': !exists(json, 'trackingCategories') ? undefined : json['trackingCategories'],
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'accountingPeriod': !exists(json, 'accountingPeriod') ? undefined : json['accountingPeriod'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'mergeRecordId': !exists(json, 'mergeRecordId') ? undefined : json['mergeRecordId'],
    };
}

export function ExpenseToJSON(value?: Expense | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'transactionDate': value.transactionDate === undefined ? undefined : (value.transactionDate.toISOString()),
        'remoteCreatedAt': value.remoteCreatedAt === undefined ? undefined : (value.remoteCreatedAt.toISOString()),
        'account': value.account,
        'contact': value.contact,
        'totalAmount': value.totalAmount,
        'subTotal': value.subTotal,
        'totalTaxAmount': value.totalTaxAmount,
        'currency': value.currency,
        'exchangeRate': value.exchangeRate,
        'company': value.company,
        'memo': value.memo,
        'lines': value.lines === undefined ? undefined : ((value.lines as Array<any>).map(ExpenseLineToJSON)),
        'trackingCategories': value.trackingCategories,
        'remoteWasDeleted': value.remoteWasDeleted,
        'accountingPeriod': value.accountingPeriod,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'mergeRecordId': value.mergeRecordId,
    };
}

