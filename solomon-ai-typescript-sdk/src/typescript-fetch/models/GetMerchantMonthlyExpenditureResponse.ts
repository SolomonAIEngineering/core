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

import { exists, mapValues } from '../runtime';
import type { MerchantMonthlyExpenditure } from './MerchantMonthlyExpenditure';
import {
    MerchantMonthlyExpenditureFromJSON,
    MerchantMonthlyExpenditureFromJSONTyped,
    MerchantMonthlyExpenditureToJSON,
} from './MerchantMonthlyExpenditure';

/**
 * 
 * @export
 * @interface GetMerchantMonthlyExpenditureResponse
 */
export interface GetMerchantMonthlyExpenditureResponse {
    /**
     * 
     * @type {Array<MerchantMonthlyExpenditure>}
     * @memberof GetMerchantMonthlyExpenditureResponse
     */
    merchantMonthlyExpenditures?: Array<MerchantMonthlyExpenditure>;
    /**
     * 
     * @type {string}
     * @memberof GetMerchantMonthlyExpenditureResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetMerchantMonthlyExpenditureResponse interface.
 */
export function instanceOfGetMerchantMonthlyExpenditureResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetMerchantMonthlyExpenditureResponseFromJSON(json: any): GetMerchantMonthlyExpenditureResponse {
    return GetMerchantMonthlyExpenditureResponseFromJSONTyped(json, false);
}

export function GetMerchantMonthlyExpenditureResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMerchantMonthlyExpenditureResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'merchantMonthlyExpenditures': !exists(json, 'merchantMonthlyExpenditures') ? undefined : ((json['merchantMonthlyExpenditures'] as Array<any>).map(MerchantMonthlyExpenditureFromJSON)),
        'nextPageNumber': !exists(json, 'nextPageNumber') ? undefined : json['nextPageNumber'],
    };
}

export function GetMerchantMonthlyExpenditureResponseToJSON(value?: GetMerchantMonthlyExpenditureResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'merchantMonthlyExpenditures': value.merchantMonthlyExpenditures === undefined ? undefined : ((value.merchantMonthlyExpenditures as Array<any>).map(MerchantMonthlyExpenditureToJSON)),
        'nextPageNumber': value.nextPageNumber,
    };
}
