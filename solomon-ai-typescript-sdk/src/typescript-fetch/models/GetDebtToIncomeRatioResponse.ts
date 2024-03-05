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
import type { DebtToIncomeRatio } from './DebtToIncomeRatio';
import {
    DebtToIncomeRatioFromJSON,
    DebtToIncomeRatioFromJSONTyped,
    DebtToIncomeRatioToJSON,
} from './DebtToIncomeRatio';

/**
 * 
 * @export
 * @interface GetDebtToIncomeRatioResponse
 */
export interface GetDebtToIncomeRatioResponse {
    /**
     * 
     * @type {Array<DebtToIncomeRatio>}
     * @memberof GetDebtToIncomeRatioResponse
     */
    debtToIncomeRatios?: Array<DebtToIncomeRatio>;
    /**
     * 
     * @type {string}
     * @memberof GetDebtToIncomeRatioResponse
     */
    nextPageNumber?: string;
}

/**
 * Check if a given object implements the GetDebtToIncomeRatioResponse interface.
 */
export function instanceOfGetDebtToIncomeRatioResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetDebtToIncomeRatioResponseFromJSON(json: any): GetDebtToIncomeRatioResponse {
    return GetDebtToIncomeRatioResponseFromJSONTyped(json, false);
}

export function GetDebtToIncomeRatioResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetDebtToIncomeRatioResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'debtToIncomeRatios': !exists(json, 'debtToIncomeRatios') ? undefined : ((json['debtToIncomeRatios'] as Array<any>).map(DebtToIncomeRatioFromJSON)),
        'nextPageNumber': !exists(json, 'nextPageNumber') ? undefined : json['nextPageNumber'],
    };
}

export function GetDebtToIncomeRatioResponseToJSON(value?: GetDebtToIncomeRatioResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'debtToIncomeRatios': value.debtToIncomeRatios === undefined ? undefined : ((value.debtToIncomeRatios as Array<any>).map(DebtToIncomeRatioToJSON)),
        'nextPageNumber': value.nextPageNumber,
    };
}
