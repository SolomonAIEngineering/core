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
import type { Apr } from './Apr';
import {
    AprFromJSON,
    AprFromJSONTyped,
    AprToJSON,
} from './Apr';
import type { BankAccountStatus } from './BankAccountStatus';
import {
    BankAccountStatusFromJSON,
    BankAccountStatusFromJSONTyped,
    BankAccountStatusToJSON,
} from './BankAccountStatus';
import type { PlaidAccountRecurringTransaction } from './PlaidAccountRecurringTransaction';
import {
    PlaidAccountRecurringTransactionFromJSON,
    PlaidAccountRecurringTransactionFromJSONTyped,
    PlaidAccountRecurringTransactionToJSON,
} from './PlaidAccountRecurringTransaction';
import type { PlaidAccountTransaction } from './PlaidAccountTransaction';
import {
    PlaidAccountTransactionFromJSON,
    PlaidAccountTransactionFromJSONTyped,
    PlaidAccountTransactionToJSON,
} from './PlaidAccountTransaction';
import type { Pocket } from './Pocket';
import {
    PocketFromJSON,
    PocketFromJSONTyped,
    PocketToJSON,
} from './Pocket';

/**
 * 
 * @export
 * @interface CreditAccount
 */
export interface CreditAccount {
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    userId?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    number?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    type?: string;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    balance?: number;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    currentFunds?: number;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    balanceLimit?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    plaidAccountId?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    subtype?: string;
    /**
     * 
     * @type {boolean}
     * @memberof CreditAccount
     */
    isOverdue?: boolean;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    lastPaymentAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    lastPaymentDate?: string;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    lastStatementIssueDate?: string;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    minimumAmountDueDate?: number;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    nextPaymentDate?: string;
    /**
     * 
     * @type {Array<Apr>}
     * @memberof CreditAccount
     */
    aprs?: Array<Apr>;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    lastStatementBalance?: number;
    /**
     * 
     * @type {number}
     * @memberof CreditAccount
     */
    minimumPaymentAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof CreditAccount
     */
    nextPaymentDueDate?: string;
    /**
     * 
     * @type {BankAccountStatus}
     * @memberof CreditAccount
     */
    status?: BankAccountStatus;
    /**
     * 
     * @type {Array<PlaidAccountTransaction>}
     * @memberof CreditAccount
     */
    transactions?: Array<PlaidAccountTransaction>;
    /**
     * 
     * @type {Array<PlaidAccountRecurringTransaction>}
     * @memberof CreditAccount
     */
    recurringTransactions?: Array<PlaidAccountRecurringTransaction>;
    /**
     * 
     * @type {Array<Pocket>}
     * @memberof CreditAccount
     */
    pockets?: Array<Pocket>;
}

/**
 * Check if a given object implements the CreditAccount interface.
 */
export function instanceOfCreditAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreditAccountFromJSON(json: any): CreditAccount {
    return CreditAccountFromJSONTyped(json, false);
}

export function CreditAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreditAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'number': !exists(json, 'number') ? undefined : json['number'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'balance': !exists(json, 'balance') ? undefined : json['balance'],
        'currentFunds': !exists(json, 'currentFunds') ? undefined : json['currentFunds'],
        'balanceLimit': !exists(json, 'balanceLimit') ? undefined : json['balanceLimit'],
        'plaidAccountId': !exists(json, 'plaidAccountId') ? undefined : json['plaidAccountId'],
        'subtype': !exists(json, 'subtype') ? undefined : json['subtype'],
        'isOverdue': !exists(json, 'isOverdue') ? undefined : json['isOverdue'],
        'lastPaymentAmount': !exists(json, 'lastPaymentAmount') ? undefined : json['lastPaymentAmount'],
        'lastPaymentDate': !exists(json, 'lastPaymentDate') ? undefined : json['lastPaymentDate'],
        'lastStatementIssueDate': !exists(json, 'lastStatementIssueDate') ? undefined : json['lastStatementIssueDate'],
        'minimumAmountDueDate': !exists(json, 'minimumAmountDueDate') ? undefined : json['minimumAmountDueDate'],
        'nextPaymentDate': !exists(json, 'nextPaymentDate') ? undefined : json['nextPaymentDate'],
        'aprs': !exists(json, 'aprs') ? undefined : ((json['aprs'] as Array<any>).map(AprFromJSON)),
        'lastStatementBalance': !exists(json, 'lastStatementBalance') ? undefined : json['lastStatementBalance'],
        'minimumPaymentAmount': !exists(json, 'minimumPaymentAmount') ? undefined : json['minimumPaymentAmount'],
        'nextPaymentDueDate': !exists(json, 'nextPaymentDueDate') ? undefined : json['nextPaymentDueDate'],
        'status': !exists(json, 'status') ? undefined : BankAccountStatusFromJSON(json['status']),
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountTransactionFromJSON)),
        'recurringTransactions': !exists(json, 'recurringTransactions') ? undefined : ((json['recurringTransactions'] as Array<any>).map(PlaidAccountRecurringTransactionFromJSON)),
        'pockets': !exists(json, 'pockets') ? undefined : ((json['pockets'] as Array<any>).map(PocketFromJSON)),
    };
}

export function CreditAccountToJSON(value?: CreditAccount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'userId': value.userId,
        'name': value.name,
        'number': value.number,
        'type': value.type,
        'balance': value.balance,
        'currentFunds': value.currentFunds,
        'balanceLimit': value.balanceLimit,
        'plaidAccountId': value.plaidAccountId,
        'subtype': value.subtype,
        'isOverdue': value.isOverdue,
        'lastPaymentAmount': value.lastPaymentAmount,
        'lastPaymentDate': value.lastPaymentDate,
        'lastStatementIssueDate': value.lastStatementIssueDate,
        'minimumAmountDueDate': value.minimumAmountDueDate,
        'nextPaymentDate': value.nextPaymentDate,
        'aprs': value.aprs === undefined ? undefined : ((value.aprs as Array<any>).map(AprToJSON)),
        'lastStatementBalance': value.lastStatementBalance,
        'minimumPaymentAmount': value.minimumPaymentAmount,
        'nextPaymentDueDate': value.nextPaymentDueDate,
        'status': BankAccountStatusToJSON(value.status),
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(PlaidAccountTransactionToJSON)),
        'recurringTransactions': value.recurringTransactions === undefined ? undefined : ((value.recurringTransactions as Array<any>).map(PlaidAccountRecurringTransactionToJSON)),
        'pockets': value.pockets === undefined ? undefined : ((value.pockets as Array<any>).map(PocketToJSON)),
    };
}
