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
import type { Budget } from './Budget';
import {
    BudgetFromJSON,
    BudgetFromJSONTyped,
    BudgetToJSON,
} from './Budget';

/**
 * 
 * @export
 * @interface GetAllBudgetsResponse
 */
export interface GetAllBudgetsResponse {
    /**
     * 
     * @type {Array<Budget>}
     * @memberof GetAllBudgetsResponse
     */
    budgets?: Array<Budget>;
}

/**
 * Check if a given object implements the GetAllBudgetsResponse interface.
 */
export function instanceOfGetAllBudgetsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetAllBudgetsResponseFromJSON(json: any): GetAllBudgetsResponse {
    return GetAllBudgetsResponseFromJSONTyped(json, false);
}

export function GetAllBudgetsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetAllBudgetsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'budgets': !exists(json, 'budgets') ? undefined : ((json['budgets'] as Array<any>).map(BudgetFromJSON)),
    };
}

export function GetAllBudgetsResponseToJSON(value?: GetAllBudgetsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'budgets': value.budgets === undefined ? undefined : ((value.budgets as Array<any>).map(BudgetToJSON)),
    };
}

