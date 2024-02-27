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
import type { BankAccountStatus } from './BankAccountStatus';
import {
    BankAccountStatusFromJSON,
    BankAccountStatusFromJSONTyped,
    BankAccountStatusToJSON,
} from './BankAccountStatus';
import type { BankAccountType } from './BankAccountType';
import {
    BankAccountTypeFromJSON,
    BankAccountTypeFromJSONTyped,
    BankAccountTypeToJSON,
} from './BankAccountType';
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
 * @interface BankAccount
 */
export interface BankAccount {
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    userId?: string;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    number?: string;
    /**
     * 
     * @type {BankAccountType}
     * @memberof BankAccount
     */
    type: BankAccountType;
    /**
     * 
     * @type {number}
     * @memberof BankAccount
     */
    balance: number;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    currency: string;
    /**
     * 
     * @type {number}
     * @memberof BankAccount
     */
    currentFunds: number;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    balanceLimit?: string;
    /**
     * 
     * @type {Array<Pocket>}
     * @memberof BankAccount
     */
    pockets?: Array<Pocket>;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    plaidAccountId?: string;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    subtype?: string;
    /**
     * 
     * @type {BankAccountStatus}
     * @memberof BankAccount
     */
    status?: BankAccountStatus;
    /**
     * 
     * @type {Array<PlaidAccountTransaction>}
     * @memberof BankAccount
     */
    transactions?: Array<PlaidAccountTransaction>;
    /**
     * 
     * @type {Array<PlaidAccountRecurringTransaction>}
     * @memberof BankAccount
     */
    recurringTransactions?: Array<PlaidAccountRecurringTransaction>;
}

/**
 * Check if a given object implements the BankAccount interface.
 */
export function instanceOfBankAccount(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "balance" in value;
    isInstance = isInstance && "currency" in value;
    isInstance = isInstance && "currentFunds" in value;

    return isInstance;
}

export function BankAccountFromJSON(json: any): BankAccount {
    return BankAccountFromJSONTyped(json, false);
}

export function BankAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): BankAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'userId': !exists(json, 'userId') ? undefined : json['userId'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'number': !exists(json, 'number') ? undefined : json['number'],
        'type': BankAccountTypeFromJSON(json['type']),
        'balance': json['balance'],
        'currency': json['currency'],
        'currentFunds': json['currentFunds'],
        'balanceLimit': !exists(json, 'balanceLimit') ? undefined : json['balanceLimit'],
        'pockets': !exists(json, 'pockets') ? undefined : ((json['pockets'] as Array<any>).map(PocketFromJSON)),
        'plaidAccountId': !exists(json, 'plaidAccountId') ? undefined : json['plaidAccountId'],
        'subtype': !exists(json, 'subtype') ? undefined : json['subtype'],
        'status': !exists(json, 'status') ? undefined : BankAccountStatusFromJSON(json['status']),
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountTransactionFromJSON)),
        'recurringTransactions': !exists(json, 'recurringTransactions') ? undefined : ((json['recurringTransactions'] as Array<any>).map(PlaidAccountRecurringTransactionFromJSON)),
    };
}

export function BankAccountToJSON(value?: BankAccount | null): any {
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
        'type': BankAccountTypeToJSON(value.type),
        'balance': value.balance,
        'currency': value.currency,
        'currentFunds': value.currentFunds,
        'balanceLimit': value.balanceLimit,
        'pockets': value.pockets === undefined ? undefined : ((value.pockets as Array<any>).map(PocketToJSON)),
        'plaidAccountId': value.plaidAccountId,
        'subtype': value.subtype,
        'status': BankAccountStatusToJSON(value.status),
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(PlaidAccountTransactionToJSON)),
        'recurringTransactions': value.recurringTransactions === undefined ? undefined : ((value.recurringTransactions as Array<any>).map(PlaidAccountRecurringTransactionToJSON)),
    };
}

