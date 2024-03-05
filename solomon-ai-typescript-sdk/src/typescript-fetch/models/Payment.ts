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
/**
 * The Payment object represents general payments made towards a specific transaction.
 * @export
 * @interface Payment
 */
export interface Payment {
    /**
     * 
     * @type {string}
     * @memberof Payment
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Payment
     */
    remoteId?: string;
    /**
     * The payment's transaction date.
     * @type {Date}
     * @memberof Payment
     */
    transactionDate?: Date;
    /**
     * The supplier, or customer involved in the payment.
     * @type {string}
     * @memberof Payment
     */
    contact?: string;
    /**
     * The supplier’s or customer’s account in which the payment is made.
     * @type {string}
     * @memberof Payment
     */
    account?: string;
    /**
     * 
     * @type {string}
     * @memberof Payment
     */
    currency?: string;
    /**
     * The payment's exchange rate.
     * 
     * Assuming string due to the example provided, but could be float or double.
     * @type {string}
     * @memberof Payment
     */
    exchangeRate?: string;
    /**
     * The company the payment belongs to.
     * @type {string}
     * @memberof Payment
     */
    company?: string;
    /**
     * The total amount of money being paid to the supplier, or customer, after taxes.
     * 
     * Might want to use double or a more precise type
     * @type {number}
     * @memberof Payment
     */
    totalAmount?: number;
    /**
     * 
     * @type {Array<string>}
     * @memberof Payment
     */
    trackingCategories?: Array<string>;
    /**
     * 
     * @type {Date}
     * @memberof Payment
     */
    remoteUpdatedAt?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof Payment
     */
    remoteWasDeleted?: boolean;
    /**
     * The accounting period that the Payment was generated in.
     * @type {string}
     * @memberof Payment
     */
    accountingPeriod?: string;
    /**
     * 
     * @type {Date}
     * @memberof Payment
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof Payment
     */
    mergeRecordId?: string;
    /**
     * 
     * @type {Date}
     * @memberof Payment
     */
    createdAt?: Date;
}

/**
 * Check if a given object implements the Payment interface.
 */
export function instanceOfPayment(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PaymentFromJSON(json: any): Payment {
    return PaymentFromJSONTyped(json, false);
}

export function PaymentFromJSONTyped(json: any, ignoreDiscriminator: boolean): Payment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'transactionDate': !exists(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'contact': !exists(json, 'contact') ? undefined : json['contact'],
        'account': !exists(json, 'account') ? undefined : json['account'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : json['exchangeRate'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'totalAmount': !exists(json, 'totalAmount') ? undefined : json['totalAmount'],
        'trackingCategories': !exists(json, 'trackingCategories') ? undefined : json['trackingCategories'],
        'remoteUpdatedAt': !exists(json, 'remoteUpdatedAt') ? undefined : (new Date(json['remoteUpdatedAt'])),
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'accountingPeriod': !exists(json, 'accountingPeriod') ? undefined : json['accountingPeriod'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'mergeRecordId': !exists(json, 'mergeRecordId') ? undefined : json['mergeRecordId'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
    };
}

export function PaymentToJSON(value?: Payment | null): any {
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
        'contact': value.contact,
        'account': value.account,
        'currency': value.currency,
        'exchangeRate': value.exchangeRate,
        'company': value.company,
        'totalAmount': value.totalAmount,
        'trackingCategories': value.trackingCategories,
        'remoteUpdatedAt': value.remoteUpdatedAt === undefined ? undefined : (value.remoteUpdatedAt.toISOString()),
        'remoteWasDeleted': value.remoteWasDeleted,
        'accountingPeriod': value.accountingPeriod,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'mergeRecordId': value.mergeRecordId,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
    };
}
