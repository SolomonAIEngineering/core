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
export const BankAccountType = {
    Unspecified: 'BANK_ACCOUNT_TYPE_UNSPECIFIED',
    Plaid: 'BANK_ACCOUNT_TYPE_PLAID',
    Manual: 'BANK_ACCOUNT_TYPE_MANUAL'
} as const;
export type BankAccountType = typeof BankAccountType[keyof typeof BankAccountType];


export function instanceOfBankAccountType(value: any): boolean {
    return Object.values(BankAccountType).includes(value);
}

export function BankAccountTypeFromJSON(json: any): BankAccountType {
    return BankAccountTypeFromJSONTyped(json, false);
}

export function BankAccountTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BankAccountType {
    return json as BankAccountType;
}

export function BankAccountTypeToJSON(value?: BankAccountType | null): any {
    return value as any;
}

