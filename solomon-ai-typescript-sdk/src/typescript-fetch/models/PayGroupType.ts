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
 * The group type Possible values include: TEAM, DEPARTMENT, COST_CENTER, BUSINESS_UNIT, GROUP. 
 * In cases where there is no clear mapping, the original value passed through will be returned.
 * @export
 */
export const PayGroupType = {
    Unspecified: 'PAY_GROUP_TYPE_UNSPECIFIED',
    Team: 'PAY_GROUP_TYPE_TEAM',
    Department: 'PAY_GROUP_TYPE_DEPARTMENT',
    CostCenter: 'PAY_GROUP_TYPE_COST_CENTER',
    BusinessUnit: 'PAY_GROUP_TYPE_BUSINESS_UNIT',
    Group: 'PAY_GROUP_TYPE_GROUP'
} as const;
export type PayGroupType = typeof PayGroupType[keyof typeof PayGroupType];


export function PayGroupTypeFromJSON(json: any): PayGroupType {
    return PayGroupTypeFromJSONTyped(json, false);
}

export function PayGroupTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PayGroupType {
    return json as PayGroupType;
}

export function PayGroupTypeToJSON(value?: PayGroupType | null): any {
    return value as any;
}

