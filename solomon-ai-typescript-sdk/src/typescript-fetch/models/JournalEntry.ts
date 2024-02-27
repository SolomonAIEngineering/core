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
import type { JournalLine } from './JournalLine';
import {
    JournalLineFromJSON,
    JournalLineFromJSONTyped,
    JournalLineToJSON,
} from './JournalLine';

/**
 * A JournalEntry is a record of a transaction or event that is entered into a company's accounting system.
 * 
 * The JournalEntry common model contains records that are automatically created as a result of a 
 * certain type of transaction, like an Invoice, and records that are manually created against a company’s ledger.
 * 
 * The lines of a given JournalEntry object should always sum to 0. A positive net_amount means
 * the line represents a debit and a negative net_amount represents a credit.
 * @export
 * @interface JournalEntry
 */
export interface JournalEntry {
    /**
     * 
     * @type {string}
     * @memberof JournalEntry
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof JournalEntry
     */
    remoteId?: string;
    /**
     * The journal entry's transaction date.
     * @type {Date}
     * @memberof JournalEntry
     */
    transactionDate?: Date;
    /**
     * When the third party's journal entry was created.
     * @type {Date}
     * @memberof JournalEntry
     */
    remoteCreatedAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof JournalEntry
     */
    remoteUpdatedAt?: Date;
    /**
     * 
     * @type {Array<string>}
     * @memberof JournalEntry
     */
    payments?: Array<string>;
    /**
     * The journal entry's private note.
     * @type {string}
     * @memberof JournalEntry
     */
    memo?: string;
    /**
     * 
     * @type {string}
     * @memberof JournalEntry
     */
    currency?: string;
    /**
     * The journal entry's exchange rate.
     * 
     * Assuming string due to the example provided, but could be float or double.
     * @type {string}
     * @memberof JournalEntry
     */
    exchangeRate?: string;
    /**
     * The company the journal entry belongs to.
     * @type {string}
     * @memberof JournalEntry
     */
    company?: string;
    /**
     * The JournalLine object is used to represent a journal entry's line items.
     * @type {Array<JournalLine>}
     * @memberof JournalEntry
     */
    lines?: Array<JournalLine>;
    /**
     * Reference number for identifying journal entries.
     * @type {string}
     * @memberof JournalEntry
     */
    journalNumber?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof JournalEntry
     */
    trackingCategories?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof JournalEntry
     */
    remoteWasDeleted?: boolean;
    /**
     * 
     * @type {string}
     * @memberof JournalEntry
     */
    postingStatus?: string;
    /**
     * The accounting period that the JournalEntry was generated in.
     * @type {string}
     * @memberof JournalEntry
     */
    accountingPeriod?: string;
    /**
     * 
     * @type {Date}
     * @memberof JournalEntry
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof JournalEntry
     */
    mergeRecordId?: string;
    /**
     * A list of the Payment Applied to Lines common models 
     * related to a given Invoice, Credit Note, or Journal Entry.
     * @type {Array<string>}
     * @memberof JournalEntry
     */
    appliedPayments?: Array<string>;
    /**
     * 
     * @type {Date}
     * @memberof JournalEntry
     */
    createdAt?: Date;
}

/**
 * Check if a given object implements the JournalEntry interface.
 */
export function instanceOfJournalEntry(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function JournalEntryFromJSON(json: any): JournalEntry {
    return JournalEntryFromJSONTyped(json, false);
}

export function JournalEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): JournalEntry {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'transactionDate': !exists(json, 'transactionDate') ? undefined : (new Date(json['transactionDate'])),
        'remoteCreatedAt': !exists(json, 'remoteCreatedAt') ? undefined : (new Date(json['remoteCreatedAt'])),
        'remoteUpdatedAt': !exists(json, 'remoteUpdatedAt') ? undefined : (new Date(json['remoteUpdatedAt'])),
        'payments': !exists(json, 'payments') ? undefined : json['payments'],
        'memo': !exists(json, 'memo') ? undefined : json['memo'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : json['exchangeRate'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'lines': !exists(json, 'lines') ? undefined : ((json['lines'] as Array<any>).map(JournalLineFromJSON)),
        'journalNumber': !exists(json, 'journalNumber') ? undefined : json['journalNumber'],
        'trackingCategories': !exists(json, 'trackingCategories') ? undefined : json['trackingCategories'],
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'postingStatus': !exists(json, 'postingStatus') ? undefined : json['postingStatus'],
        'accountingPeriod': !exists(json, 'accountingPeriod') ? undefined : json['accountingPeriod'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'mergeRecordId': !exists(json, 'mergeRecordId') ? undefined : json['mergeRecordId'],
        'appliedPayments': !exists(json, 'appliedPayments') ? undefined : json['appliedPayments'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
    };
}

export function JournalEntryToJSON(value?: JournalEntry | null): any {
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
        'remoteUpdatedAt': value.remoteUpdatedAt === undefined ? undefined : (value.remoteUpdatedAt.toISOString()),
        'payments': value.payments,
        'memo': value.memo,
        'currency': value.currency,
        'exchangeRate': value.exchangeRate,
        'company': value.company,
        'lines': value.lines === undefined ? undefined : ((value.lines as Array<any>).map(JournalLineToJSON)),
        'journalNumber': value.journalNumber,
        'trackingCategories': value.trackingCategories,
        'remoteWasDeleted': value.remoteWasDeleted,
        'postingStatus': value.postingStatus,
        'accountingPeriod': value.accountingPeriod,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'mergeRecordId': value.mergeRecordId,
        'appliedPayments': value.appliedPayments,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
    };
}

