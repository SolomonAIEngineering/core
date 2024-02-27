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
export const LinkType = {
    Unspecified: 'LINK_TYPE_UNSPECIFIED',
    Plaid: 'LINK_TYPE_PLAID',
    Manual: 'LINK_TYPE_MANUAL'
} as const;
export type LinkType = typeof LinkType[keyof typeof LinkType];


export function LinkTypeFromJSON(json: any): LinkType {
    return LinkTypeFromJSONTyped(json, false);
}

export function LinkTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): LinkType {
    return json as LinkType;
}

export function LinkTypeToJSON(value?: LinkType | null): any {
    return value as any;
}

