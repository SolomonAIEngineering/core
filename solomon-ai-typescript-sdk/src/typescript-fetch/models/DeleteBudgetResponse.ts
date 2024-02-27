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
/**
 * 
 * @export
 * @interface DeleteBudgetResponse
 */
export interface DeleteBudgetResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteBudgetResponse
     */
    deleted?: boolean;
}

/**
 * Check if a given object implements the DeleteBudgetResponse interface.
 */
export function instanceOfDeleteBudgetResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DeleteBudgetResponseFromJSON(json: any): DeleteBudgetResponse {
    return DeleteBudgetResponseFromJSONTyped(json, false);
}

export function DeleteBudgetResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteBudgetResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deleted': !exists(json, 'deleted') ? undefined : json['deleted'],
    };
}

export function DeleteBudgetResponseToJSON(value?: DeleteBudgetResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deleted': value.deleted,
    };
}

