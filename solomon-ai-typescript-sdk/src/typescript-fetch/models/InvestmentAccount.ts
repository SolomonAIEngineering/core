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
}

/**
 * Check if a given object implements the InvestmentAccount interface.
 */
export function instanceOfInvestmentAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InvestmentAccountFromJSON(json: any): InvestmentAccount {
    return InvestmentAccountFromJSONTyped(json, false);
}

export function InvestmentAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvestmentAccount {
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
        'holdings': !exists(json, 'holdings') ? undefined : ((json['holdings'] as Array<any>).map(InvesmentHoldingFromJSON)),
        'securities': !exists(json, 'securities') ? undefined : ((json['securities'] as Array<any>).map(InvestmentSecurityFromJSON)),
        'status': !exists(json, 'status') ? undefined : BankAccountStatusFromJSON(json['status']),
        'transactions': !exists(json, 'transactions') ? undefined : ((json['transactions'] as Array<any>).map(PlaidAccountInvestmentTransactionFromJSON)),
    };
}

export function InvestmentAccountToJSON(value?: InvestmentAccount | null): any {
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
        'holdings': value.holdings === undefined ? undefined : ((value.holdings as Array<any>).map(InvesmentHoldingToJSON)),
        'securities': value.securities === undefined ? undefined : ((value.securities as Array<any>).map(InvestmentSecurityToJSON)),
        'status': BankAccountStatusToJSON(value.status),
        'transactions': value.transactions === undefined ? undefined : ((value.transactions as Array<any>).map(PlaidAccountInvestmentTransactionToJSON)),
    };
}
