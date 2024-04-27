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
import type { MerchantMetricsFinancialSubProfile } from './MerchantMetricsFinancialSubProfile';
import {
    MerchantMetricsFinancialSubProfileFromJSON,
    MerchantMetricsFinancialSubProfileFromJSONTyped,
    MerchantMetricsFinancialSubProfileToJSON,
} from './MerchantMetricsFinancialSubProfile';

/**
 * 
 * @export
 * @interface GetMerchantMetricsFinancialSubProfileOverTimeResponse
 */
export interface GetMerchantMetricsFinancialSubProfileOverTimeResponse {
    /**
     * 
     * @type {Array<MerchantMetricsFinancialSubProfile>}
     * @memberof GetMerchantMetricsFinancialSubProfileOverTimeResponse
     */
    data?: Array<MerchantMetricsFinancialSubProfile>;
}

/**
 * Check if a given object implements the GetMerchantMetricsFinancialSubProfileOverTimeResponse interface.
 */
export function instanceOfGetMerchantMetricsFinancialSubProfileOverTimeResponse(value: object): boolean {
    return true;
}

export function GetMerchantMetricsFinancialSubProfileOverTimeResponseFromJSON(json: any): GetMerchantMetricsFinancialSubProfileOverTimeResponse {
    return GetMerchantMetricsFinancialSubProfileOverTimeResponseFromJSONTyped(json, false);
}

export function GetMerchantMetricsFinancialSubProfileOverTimeResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMerchantMetricsFinancialSubProfileOverTimeResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'] == null ? undefined : ((json['data'] as Array<any>).map(MerchantMetricsFinancialSubProfileFromJSON)),
    };
}

export function GetMerchantMetricsFinancialSubProfileOverTimeResponseToJSON(value?: GetMerchantMetricsFinancialSubProfileOverTimeResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'data': value['data'] == null ? undefined : ((value['data'] as Array<any>).map(MerchantMetricsFinancialSubProfileToJSON)),
    };
}
