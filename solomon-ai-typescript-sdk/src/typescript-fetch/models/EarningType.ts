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
 * The type of earning. Possible values include: SALARY, REIMBURSEMENT, OVERTIME, BONUS. 
 * In cases where there is no clear mapping, the original value passed through will be returned.
 * @export
 */
export const EarningType = {
    Unspecified: 'EARNING_TYPE_UNSPECIFIED',
    Salary: 'EARNING_TYPE_SALARY',
    Reimbursement: 'EARNING_TYPE_REIMBURSEMENT',
    Overtime: 'EARNING_TYPE_OVERTIME',
    Bonus: 'EARNING_TYPE_BONUS'
} as const;
export type EarningType = typeof EarningType[keyof typeof EarningType];


export function EarningTypeFromJSON(json: any): EarningType {
    return EarningTypeFromJSONTyped(json, false);
}

export function EarningTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EarningType {
    return json as EarningType;
}

export function EarningTypeToJSON(value?: EarningType | null): any {
    return value as any;
}

