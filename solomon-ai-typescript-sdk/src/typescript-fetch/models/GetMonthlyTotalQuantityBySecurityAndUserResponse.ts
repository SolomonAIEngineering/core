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

import { mapValues } from '../runtime';
import type { MonthlyTotalQuantityBySecurityAndUser } from './MonthlyTotalQuantityBySecurityAndUser';
import {
    MonthlyTotalQuantityBySecurityAndUserFromJSON,
    MonthlyTotalQuantityBySecurityAndUserFromJSONTyped,
    MonthlyTotalQuantityBySecurityAndUserToJSON,
} from './MonthlyTotalQuantityBySecurityAndUser';

/**
 * 
 * @export
 * @interface GetMonthlyTotalQuantityBySecurityAndUserResponse
 */
export interface GetMonthlyTotalQuantityBySecurityAndUserResponse {
    /**
     * 
     * @type {Array<MonthlyTotalQuantityBySecurityAndUser>}
     * @memberof GetMonthlyTotalQuantityBySecurityAndUserResponse
     */
    monthlyTotalQuantityBySecurityAndUser?: Array<MonthlyTotalQuantityBySecurityAndUser>;
    /**
     * 
     * @type {string}
     * @memberof GetMonthlyTotalQuantityBySecurityAndUserResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetMonthlyTotalQuantityBySecurityAndUserResponse interface.
 */
export function instanceOfGetMonthlyTotalQuantityBySecurityAndUserResponse(value: object): boolean {
    return true;
}

export function GetMonthlyTotalQuantityBySecurityAndUserResponseFromJSON(json: any): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    return GetMonthlyTotalQuantityBySecurityAndUserResponseFromJSONTyped(json, false);
}

export function GetMonthlyTotalQuantityBySecurityAndUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMonthlyTotalQuantityBySecurityAndUserResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'monthlyTotalQuantityBySecurityAndUser': json['monthlyTotalQuantityBySecurityAndUser'] == null ? undefined : ((json['monthlyTotalQuantityBySecurityAndUser'] as Array<any>).map(MonthlyTotalQuantityBySecurityAndUserFromJSON)),
        'nextPageNumber': json['nextPageNumber'] == null ? undefined : json['nextPageNumber'],
    };
}

export function GetMonthlyTotalQuantityBySecurityAndUserResponseToJSON(value?: GetMonthlyTotalQuantityBySecurityAndUserResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'monthlyTotalQuantityBySecurityAndUser': value['monthlyTotalQuantityBySecurityAndUser'] == null ? undefined : ((value['monthlyTotalQuantityBySecurityAndUser'] as Array<any>).map(MonthlyTotalQuantityBySecurityAndUserToJSON)),
        'nextPageNumber': value['nextPageNumber'],
    };
}

