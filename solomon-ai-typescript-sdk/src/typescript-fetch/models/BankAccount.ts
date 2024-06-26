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
import type { AccountStatements } from './AccountStatements';
import {
    AccountStatementsFromJSON,
    AccountStatementsFromJSONTyped,
    AccountStatementsToJSON,
} from './AccountStatements';
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
    /**
     * 
     * @type {Array<AccountStatements>}
     * @memberof BankAccount
     */
    statements?: Array<AccountStatements>;
    /**
     * 
     * @type {string}
     * @memberof BankAccount
     */
    plaidAccountType?: string;
}

/**
 * Check if a given object implements the BankAccount interface.
 */
export function instanceOfBankAccount(value: object): boolean {
    if (!('type' in value)) return false;
    if (!('balance' in value)) return false;
    if (!('currency' in value)) return false;
    if (!('currentFunds' in value)) return false;
    return true;
}

export function BankAccountFromJSON(json: any): BankAccount {
    return BankAccountFromJSONTyped(json, false);
}

export function BankAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): BankAccount {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'userId': json['userId'] == null ? undefined : json['userId'],
        'name': json['name'] == null ? undefined : json['name'],
        'number': json['number'] == null ? undefined : json['number'],
        'type': BankAccountTypeFromJSON(json['type']),
        'balance': json['balance'],
        'currency': json['currency'],
        'currentFunds': json['currentFunds'],
        'balanceLimit': json['balanceLimit'] == null ? undefined : json['balanceLimit'],
        'pockets': json['pockets'] == null ? undefined : ((json['pockets'] as Array<any>).map(PocketFromJSON)),
        'plaidAccountId': json['plaidAccountId'] == null ? undefined : json['plaidAccountId'],
        'subtype': json['subtype'] == null ? undefined : json['subtype'],
        'status': json['status'] == null ? undefined : BankAccountStatusFromJSON(json['status']),
        'transactions': json['transactions'] == null ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountTransactionFromJSON)),
        'recurringTransactions': json['recurringTransactions'] == null ? undefined : ((json['recurringTransactions'] as Array<any>).map(PlaidAccountRecurringTransactionFromJSON)),
        'statements': json['statements'] == null ? undefined : ((json['statements'] as Array<any>).map(AccountStatementsFromJSON)),
        'plaidAccountType': json['plaidAccountType'] == null ? undefined : json['plaidAccountType'],
    };
}

export function BankAccountToJSON(value?: BankAccount | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'userId': value['userId'],
        'name': value['name'],
        'number': value['number'],
        'type': BankAccountTypeToJSON(value['type']),
        'balance': value['balance'],
        'currency': value['currency'],
        'currentFunds': value['currentFunds'],
        'balanceLimit': value['balanceLimit'],
        'pockets': value['pockets'] == null ? undefined : ((value['pockets'] as Array<any>).map(PocketToJSON)),
        'plaidAccountId': value['plaidAccountId'],
        'subtype': value['subtype'],
        'status': BankAccountStatusToJSON(value['status']),
        'transactions': value['transactions'] == null ? undefined : ((value['transactions'] as Array<any>).map(PlaidAccountTransactionToJSON)),
        'recurringTransactions': value['recurringTransactions'] == null ? undefined : ((value['recurringTransactions'] as Array<any>).map(PlaidAccountRecurringTransactionToJSON)),
        'statements': value['statements'] == null ? undefined : ((value['statements'] as Array<any>).map(AccountStatementsToJSON)),
        'plaidAccountType': value['plaidAccountType'],
    };
}

