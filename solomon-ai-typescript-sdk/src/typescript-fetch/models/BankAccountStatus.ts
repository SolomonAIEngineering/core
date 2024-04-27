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


/**
 * 
 * @export
 */
export const BankAccountStatus = {
    Unspecified: 'BANK_ACCOUNT_STATUS_UNSPECIFIED',
    Active: 'BANK_ACCOUNT_STATUS_ACTIVE',
    Inactive: 'BANK_ACCOUNT_STATUS_INACTIVE'
} as const;
export type BankAccountStatus = typeof BankAccountStatus[keyof typeof BankAccountStatus];


export function instanceOfBankAccountStatus(value: any): boolean {
    return Object.values(BankAccountStatus).includes(value);
}

export function BankAccountStatusFromJSON(json: any): BankAccountStatus {
    return BankAccountStatusFromJSONTyped(json, false);
}

export function BankAccountStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): BankAccountStatus {
    return json as BankAccountStatus;
}

export function BankAccountStatusToJSON(value?: BankAccountStatus | null): any {
    return value as any;
}

