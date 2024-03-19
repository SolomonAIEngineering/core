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
export const BankAccountType1 = {
    Unspecified: 'BANK_ACCOUNT_TYPE_UNSPECIFIED',
    Checking: 'BANK_ACCOUNT_TYPE_CHECKING',
    Savings: 'BANK_ACCOUNT_TYPE_SAVINGS'
} as const;
export type BankAccountType1 = typeof BankAccountType1[keyof typeof BankAccountType1];


export function BankAccountType1FromJSON(json: any): BankAccountType1 {
    return BankAccountType1FromJSONTyped(json, false);
}

export function BankAccountType1FromJSONTyped(json: any, ignoreDiscriminator: boolean): BankAccountType1 {
    return json as BankAccountType1;
}

export function BankAccountType1ToJSON(value?: BankAccountType1 | null): any {
    return value as any;
}
