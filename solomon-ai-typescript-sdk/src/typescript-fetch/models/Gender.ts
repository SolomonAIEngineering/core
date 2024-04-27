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
 * The employee's gender. Possible values include: MALE, FEMALE, NON-BINARY, OTHER,
 * PREFER_NOT_TO_DISCLOSE. In cases where there is no clear mapping, 
 * the original value passed through will be returned.
 * @export
 */
export const Gender = {
    Unspecified: 'GENDER_UNSPECIFIED',
    Male: 'GENDER_MALE',
    Female: 'GENDER_FEMALE',
    NonBinary: 'GENDER_NON_BINARY',
    Other: 'GENDER_OTHER',
    PreferNotToDisclose: 'GENDER_PREFER_NOT_TO_DISCLOSE'
} as const;
export type Gender = typeof Gender[keyof typeof Gender];


export function instanceOfGender(value: any): boolean {
    return Object.values(Gender).includes(value);
}

export function GenderFromJSON(json: any): Gender {
    return GenderFromJSONTyped(json, false);
}

export function GenderFromJSONTyped(json: any, ignoreDiscriminator: boolean): Gender {
    return json as Gender;
}

export function GenderToJSON(value?: Gender | null): any {
    return value as any;
}

