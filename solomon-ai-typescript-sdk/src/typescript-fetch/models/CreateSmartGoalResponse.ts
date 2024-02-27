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
 * @interface CreateSmartGoalResponse
 */
export interface CreateSmartGoalResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateSmartGoalResponse
     */
    smartGoalId?: string;
}

/**
 * Check if a given object implements the CreateSmartGoalResponse interface.
 */
export function instanceOfCreateSmartGoalResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateSmartGoalResponseFromJSON(json: any): CreateSmartGoalResponse {
    return CreateSmartGoalResponseFromJSONTyped(json, false);
}

export function CreateSmartGoalResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateSmartGoalResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'smartGoalId': !exists(json, 'smartGoalId') ? undefined : json['smartGoalId'],
    };
}

export function CreateSmartGoalResponseToJSON(value?: CreateSmartGoalResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'smartGoalId': value.smartGoalId,
    };
}

