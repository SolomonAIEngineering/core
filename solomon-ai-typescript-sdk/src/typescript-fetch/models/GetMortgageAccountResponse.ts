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
import type { MortgageAccount } from './MortgageAccount';
import {
    MortgageAccountFromJSON,
    MortgageAccountFromJSONTyped,
    MortgageAccountToJSON,
} from './MortgageAccount';

/**
 * 
 * @export
 * @interface GetMortgageAccountResponse
 */
export interface GetMortgageAccountResponse {
    /**
     * 
     * @type {MortgageAccount}
     * @memberof GetMortgageAccountResponse
     */
    mortageAccount?: MortgageAccount;
}

/**
 * Check if a given object implements the GetMortgageAccountResponse interface.
 */
export function instanceOfGetMortgageAccountResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetMortgageAccountResponseFromJSON(json: any): GetMortgageAccountResponse {
    return GetMortgageAccountResponseFromJSONTyped(json, false);
}

export function GetMortgageAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMortgageAccountResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mortageAccount': !exists(json, 'mortageAccount') ? undefined : MortgageAccountFromJSON(json['mortageAccount']),
    };
}

export function GetMortgageAccountResponseToJSON(value?: GetMortgageAccountResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'mortageAccount': MortgageAccountToJSON(value.mortageAccount),
    };
}

