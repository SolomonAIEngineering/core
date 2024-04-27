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
import type { InvestmentAccount } from './InvestmentAccount';
import {
    InvestmentAccountFromJSON,
    InvestmentAccountFromJSONTyped,
    InvestmentAccountToJSON,
} from './InvestmentAccount';

/**
 * 
 * @export
 * @interface GetInvestmentAcccountResponse
 */
export interface GetInvestmentAcccountResponse {
    /**
     * 
     * @type {InvestmentAccount}
     * @memberof GetInvestmentAcccountResponse
     */
    investmentAccount?: InvestmentAccount;
}

/**
 * Check if a given object implements the GetInvestmentAcccountResponse interface.
 */
export function instanceOfGetInvestmentAcccountResponse(value: object): boolean {
    return true;
}

export function GetInvestmentAcccountResponseFromJSON(json: any): GetInvestmentAcccountResponse {
    return GetInvestmentAcccountResponseFromJSONTyped(json, false);
}

export function GetInvestmentAcccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetInvestmentAcccountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'investmentAccount': json['investmentAccount'] == null ? undefined : InvestmentAccountFromJSON(json['investmentAccount']),
    };
}

export function GetInvestmentAcccountResponseToJSON(value?: GetInvestmentAcccountResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'investmentAccount': InvestmentAccountToJSON(value['investmentAccount']),
    };
}

