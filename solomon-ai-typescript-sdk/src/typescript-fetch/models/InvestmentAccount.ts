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
import type { InvesmentHolding } from './InvesmentHolding';
import {
    InvesmentHoldingFromJSON,
    InvesmentHoldingFromJSONTyped,
    InvesmentHoldingToJSON,
} from './InvesmentHolding';
import type { InvestmentSecurity } from './InvestmentSecurity';
import {
    InvestmentSecurityFromJSON,
    InvestmentSecurityFromJSONTyped,
    InvestmentSecurityToJSON,
} from './InvestmentSecurity';
import type { PlaidAccountInvestmentTransaction } from './PlaidAccountInvestmentTransaction';
import {
    PlaidAccountInvestmentTransactionFromJSON,
    PlaidAccountInvestmentTransactionFromJSONTyped,
    PlaidAccountInvestmentTransactionToJSON,
} from './PlaidAccountInvestmentTransaction';

/**
 * 
 * @export
 * @interface InvestmentAccount
 */
export interface InvestmentAccount {
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    userId?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    number?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    type?: string;
    /**
     * 
     * @type {number}
     * @memberof InvestmentAccount
     */
    balance?: number;
    /**
     * 
     * @type {number}
     * @memberof InvestmentAccount
     */
    currentFunds?: number;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    balanceLimit?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    plaidAccountId?: string;
    /**
     * 
     * @type {string}
     * @memberof InvestmentAccount
     */
    subtype?: string;
    /**
     * 
     * @type {Array<InvesmentHolding>}
     * @memberof InvestmentAccount
     */
    holdings?: Array<InvesmentHolding>;
    /**
     * 
     * @type {Array<InvestmentSecurity>}
     * @memberof InvestmentAccount
     */
    securities?: Array<InvestmentSecurity>;
    /**
     * 
     * @type {BankAccountStatus}
     * @memberof InvestmentAccount
     */
    status?: BankAccountStatus;
    /**
     * 
     * @type {Array<PlaidAccountInvestmentTransaction>}
     * @memberof InvestmentAccount
     */
    transactions?: Array<PlaidAccountInvestmentTransaction>;
    /**
     * 
     * @type {Array<AccountStatements>}
     * @memberof InvestmentAccount
     */
    statements?: Array<AccountStatements>;
}

/**
 * Check if a given object implements the InvestmentAccount interface.
 */
export function instanceOfInvestmentAccount(value: object): boolean {
    return true;
}

export function InvestmentAccountFromJSON(json: any): InvestmentAccount {
    return InvestmentAccountFromJSONTyped(json, false);
}

export function InvestmentAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvestmentAccount {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'userId': json['userId'] == null ? undefined : json['userId'],
        'name': json['name'] == null ? undefined : json['name'],
        'number': json['number'] == null ? undefined : json['number'],
        'type': json['type'] == null ? undefined : json['type'],
        'balance': json['balance'] == null ? undefined : json['balance'],
        'currentFunds': json['currentFunds'] == null ? undefined : json['currentFunds'],
        'balanceLimit': json['balanceLimit'] == null ? undefined : json['balanceLimit'],
        'plaidAccountId': json['plaidAccountId'] == null ? undefined : json['plaidAccountId'],
        'subtype': json['subtype'] == null ? undefined : json['subtype'],
        'holdings': json['holdings'] == null ? undefined : ((json['holdings'] as Array<any>).map(InvesmentHoldingFromJSON)),
        'securities': json['securities'] == null ? undefined : ((json['securities'] as Array<any>).map(InvestmentSecurityFromJSON)),
        'status': json['status'] == null ? undefined : BankAccountStatusFromJSON(json['status']),
        'transactions': json['transactions'] == null ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountInvestmentTransactionFromJSON)),
        'statements': json['statements'] == null ? undefined : ((json['statements'] as Array<any>).map(AccountStatementsFromJSON)),
    };
}

export function InvestmentAccountToJSON(value?: InvestmentAccount | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'userId': value['userId'],
        'name': value['name'],
        'number': value['number'],
        'type': value['type'],
        'balance': value['balance'],
        'currentFunds': value['currentFunds'],
        'balanceLimit': value['balanceLimit'],
        'plaidAccountId': value['plaidAccountId'],
        'subtype': value['subtype'],
        'holdings': value['holdings'] == null ? undefined : ((value['holdings'] as Array<any>).map(InvesmentHoldingToJSON)),
        'securities': value['securities'] == null ? undefined : ((value['securities'] as Array<any>).map(InvestmentSecurityToJSON)),
        'status': BankAccountStatusToJSON(value['status']),
        'transactions': value['transactions'] == null ? undefined : ((value['transactions'] as Array<any>).map(PlaidAccountInvestmentTransactionToJSON)),
        'statements': value['statements'] == null ? undefined : ((value['statements'] as Array<any>).map(AccountStatementsToJSON)),
    };
}

