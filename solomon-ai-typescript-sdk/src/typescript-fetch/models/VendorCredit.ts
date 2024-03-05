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
import type { VendorCreditLine } from './VendorCreditLine';
import {
    VendorCreditLineFromJSON,
    VendorCreditLineFromJSONTyped,
    VendorCreditLineToJSON,
} from './VendorCreditLine';

/**
 * The VendorCredit object represents an accounts receivable transaction indicating that a customer is owed a gift or refund. 
 * It includes details such as the amount of credit, the vendor responsible, 
 * the associated account, and other relevant information.
 * @export
 * @interface VendorCredit
 */
export interface VendorCredit {
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    mergeRecordId?: string;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    remoteId?: string;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    number?: string;
    /**
     * 
     * @type {Date}
     * @memberof VendorCredit
     */
    transactionDate?: Date;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    vendor?: string;
    /**
     * 
     * @type {number}
     * @memberof VendorCredit
     */
    totalAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    currency?: string;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    exchangeRate?: string;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    company?: string;
    /**
     * 
     * @type {Array<VendorCreditLine>}
     * @memberof VendorCredit
     */
    lines?: Array<VendorCreditLine>;
    /**
     * 
     * @type {Array<string>}
     * @memberof VendorCredit
     */
    trackingCategories?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof VendorCredit
     */
    remoteWasDeleted?: boolean;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    accountingPeriod?: string;
    /**
     * 
     * @type {Date}
     * @memberof VendorCredit
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof VendorCredit
     */
    id?: string;
}

/**
 * Check if a given object implements the VendorCredit interface.
 */
export function instanceOfVendorCredit(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function VendorCreditFromJSON(json: any): VendorCredit {
    return VendorCreditFromJSONTyped(json, false);
}

export function VendorCreditFromJSONTyped(json: any, ignoreDiscriminator: boolean): VendorCredit {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mergeRecordId': !exists(json, 'mergeRecordId') ? undefined : json['mergeRecordId'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'number': !exists(json, 'number') ? undefined : json['number'],
        'transactionDate': !exists(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'vendor': !exists(json, 'vendor') ? undefined : json['vendor'],
        'totalAmount': !exists(json, 'totalAmount') ? undefined : json['totalAmount'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : json['exchangeRate'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'lines': !exists(json, 'lines') ? undefined : ((json['lines'] as Array<any>).map(VendorCreditLineFromJSON)),
        'trackingCategories': !exists(json, 'trackingCategories') ? undefined : json['trackingCategories'],
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'accountingPeriod': !exists(json, 'accountingPeriod') ? undefined : json['accountingPeriod'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function VendorCreditToJSON(value?: VendorCredit | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'mergeRecordId': value.mergeRecordId,
        'remoteId': value.remoteId,
        'number': value.number,
        'transactionDate': value.transactionDate === undefined ? undefined : (value.transactionDate.toISOString()),
        'vendor': value.vendor,
        'totalAmount': value.totalAmount,
        'currency': value.currency,
        'exchangeRate': value.exchangeRate,
        'company': value.company,
        'lines': value.lines === undefined ? undefined : ((value.lines as Array<any>).map(VendorCreditLineToJSON)),
        'trackingCategories': value.trackingCategories,
        'remoteWasDeleted': value.remoteWasDeleted,
        'accountingPeriod': value.accountingPeriod,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'id': value.id,
    };
}
