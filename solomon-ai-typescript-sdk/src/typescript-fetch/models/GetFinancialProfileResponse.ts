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
import type { FinancialProfile } from './FinancialProfile';
import {
    FinancialProfileFromJSON,
    FinancialProfileFromJSONTyped,
    FinancialProfileToJSON,
} from './FinancialProfile';

/**
 * 
 * @export
 * @interface GetFinancialProfileResponse
 */
export interface GetFinancialProfileResponse {
    /**
     * 
     * @type {Array<FinancialProfile>}
     * @memberof GetFinancialProfileResponse
     */
    financialProfiles?: Array<FinancialProfile>;
    /**
     * 
     * @type {string}
     * @memberof GetFinancialProfileResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetFinancialProfileResponse interface.
 */
export function instanceOfGetFinancialProfileResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetFinancialProfileResponseFromJSON(json: any): GetFinancialProfileResponse {
    return GetFinancialProfileResponseFromJSONTyped(json, false);
}

export function GetFinancialProfileResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetFinancialProfileResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'financialProfiles': !exists(json, 'financialProfiles') ? undefined : ((json['financialProfiles'] as Array<any>).map(FinancialProfileFromJSON)),
        'nextPageNumber': !exists(json, 'nextPageNumber') ? undefined : json['nextPageNumber'],
    };
}

export function GetFinancialProfileResponseToJSON(value?: GetFinancialProfileResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'financialProfiles': value.financialProfiles === undefined ? undefined : ((value.financialProfiles as Array<any>).map(FinancialProfileToJSON)),
        'nextPageNumber': value.nextPageNumber,
    };
}

