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
 * The position's pay frequency. Possible values include: WEEKLY, BIWEEKLY, MONTHLY, 
 * QUARTERLY, SEMIANNUALLY, ANNUALLY, THIRTEEN-MONTHLY, PRO_RATA, SEMIMONTHLY. 
 * In cases where there is no clear mapping, the original value passed through will be returned.
 * @export
 */
export const PayFrequency = {
    Unspecified: 'PAY_FREQUENCY_UNSPECIFIED',
    Weekly: 'PAY_FREQUENCY_WEEKLY',
    Biweekly: 'PAY_FREQUENCY_BIWEEKLY',
    Monthly: 'PAY_FREQUENCY_MONTHLY',
    Quarterly: 'PAY_FREQUENCY_QUARTERLY',
    Semiannually: 'PAY_FREQUENCY_SEMIANNUALLY',
    Annually: 'PAY_FREQUENCY_ANNUALLY',
    ThirteenMonthly: 'PAY_FREQUENCY_THIRTEEN_MONTHLY',
    ProRata: 'PAY_FREQUENCY_PRO_RATA',
    Semimonthly: 'PAY_FREQUENCY_SEMIMONTHLY'
} as const;
export type PayFrequency = typeof PayFrequency[keyof typeof PayFrequency];


export function instanceOfPayFrequency(value: any): boolean {
    return Object.values(PayFrequency).includes(value);
}

export function PayFrequencyFromJSON(json: any): PayFrequency {
    return PayFrequencyFromJSONTyped(json, false);
}

export function PayFrequencyFromJSONTyped(json: any, ignoreDiscriminator: boolean): PayFrequency {
    return json as PayFrequency;
}

export function PayFrequencyToJSON(value?: PayFrequency | null): any {
    return value as any;
}

