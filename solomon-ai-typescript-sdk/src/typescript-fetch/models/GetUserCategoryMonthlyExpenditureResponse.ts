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
import type { CategoryMonthlyExpenditure } from './CategoryMonthlyExpenditure';
import {
    CategoryMonthlyExpenditureFromJSON,
    CategoryMonthlyExpenditureFromJSONTyped,
    CategoryMonthlyExpenditureToJSON,
} from './CategoryMonthlyExpenditure';

/**
 * 
 * @export
 * @interface GetUserCategoryMonthlyExpenditureResponse
 */
export interface GetUserCategoryMonthlyExpenditureResponse {
    /**
     * 
     * @type {Array<CategoryMonthlyExpenditure>}
     * @memberof GetUserCategoryMonthlyExpenditureResponse
     */
    categoryMonthlyExpenditure?: Array<CategoryMonthlyExpenditure>;
    /**
     * 
     * @type {string}
     * @memberof GetUserCategoryMonthlyExpenditureResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetUserCategoryMonthlyExpenditureResponse interface.
 */
export function instanceOfGetUserCategoryMonthlyExpenditureResponse(value: object): boolean {
    return true;
}

export function GetUserCategoryMonthlyExpenditureResponseFromJSON(json: any): GetUserCategoryMonthlyExpenditureResponse {
    return GetUserCategoryMonthlyExpenditureResponseFromJSONTyped(json, false);
}

export function GetUserCategoryMonthlyExpenditureResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetUserCategoryMonthlyExpenditureResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'categoryMonthlyExpenditure': json['categoryMonthlyExpenditure'] == null ? undefined : ((json['categoryMonthlyExpenditure'] as Array<any>).map(CategoryMonthlyExpenditureFromJSON)),
        'nextPageNumber': json['nextPageNumber'] == null ? undefined : json['nextPageNumber'],
    };
}

export function GetUserCategoryMonthlyExpenditureResponseToJSON(value?: GetUserCategoryMonthlyExpenditureResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'categoryMonthlyExpenditure': value['categoryMonthlyExpenditure'] == null ? undefined : ((value['categoryMonthlyExpenditure'] as Array<any>).map(CategoryMonthlyExpenditureToJSON)),
        'nextPageNumber': value['nextPageNumber'],
    };
}

