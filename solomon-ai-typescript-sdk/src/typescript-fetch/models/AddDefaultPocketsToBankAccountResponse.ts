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
import type { BankAccount } from './BankAccount';
import {
    BankAccountFromJSON,
    BankAccountFromJSONTyped,
    BankAccountToJSON,
} from './BankAccount';
import type { CreditAccount } from './CreditAccount';
import {
    CreditAccountFromJSON,
    CreditAccountFromJSONTyped,
    CreditAccountToJSON,
} from './CreditAccount';

/**
 * 
 * @export
 * @interface AddDefaultPocketsToBankAccountResponse
 */
export interface AddDefaultPocketsToBankAccountResponse {
    /**
     * 
     * @type {BankAccount}
     * @memberof AddDefaultPocketsToBankAccountResponse
     */
    bankAccount?: BankAccount;
    /**
     * 
     * @type {CreditAccount}
     * @memberof AddDefaultPocketsToBankAccountResponse
     */
    creditAccount?: CreditAccount;
}

/**
 * Check if a given object implements the AddDefaultPocketsToBankAccountResponse interface.
 */
export function instanceOfAddDefaultPocketsToBankAccountResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AddDefaultPocketsToBankAccountResponseFromJSON(json: any): AddDefaultPocketsToBankAccountResponse {
    return AddDefaultPocketsToBankAccountResponseFromJSONTyped(json, false);
}

export function AddDefaultPocketsToBankAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddDefaultPocketsToBankAccountResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bankAccount': !exists(json, 'bankAccount') ? undefined : BankAccountFromJSON(json['bankAccount']),
        'creditAccount': !exists(json, 'creditAccount') ? undefined : CreditAccountFromJSON(json['creditAccount']),
    };
}

export function AddDefaultPocketsToBankAccountResponseToJSON(value?: AddDefaultPocketsToBankAccountResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bankAccount': BankAccountToJSON(value.bankAccount),
        'creditAccount': CreditAccountToJSON(value.creditAccount),
    };
}
