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
import type { Category1 } from './Category1';
import {
    Category1FromJSON,
    Category1FromJSONTyped,
    Category1ToJSON,
} from './Category1';

/**
 * The Budgets table stores information about each budget created by the user,
 * including the name of the budget, the start and end dates, and the user ID.
 * @export
 * @interface Budget
 */
export interface Budget {
    /**
     * 
     * @type {string}
     * @memberof Budget
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Budget
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof Budget
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof Budget
     */
    startDate?: string;
    /**
     * 
     * @type {string}
     * @memberof Budget
     */
    endDate?: string;
    /**
     * 
     * @type {Category1}
     * @memberof Budget
     */
    category?: Category1;
}

/**
 * Check if a given object implements the Budget interface.
 */
export function instanceOfBudget(value: object): boolean {
    return true;
}

export function BudgetFromJSON(json: any): Budget {
    return BudgetFromJSONTyped(json, false);
}

export function BudgetFromJSONTyped(json: any, ignoreDiscriminator: boolean): Budget {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'description': json['description'] == null ? undefined : json['description'],
        'startDate': json['startDate'] == null ? undefined : json['startDate'],
        'endDate': json['endDate'] == null ? undefined : json['endDate'],
        'category': json['category'] == null ? undefined : Category1FromJSON(json['category']),
    };
}

export function BudgetToJSON(value?: Budget | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'name': value['name'],
        'description': value['description'],
        'startDate': value['startDate'],
        'endDate': value['endDate'],
        'category': Category1ToJSON(value['category']),
    };
}

