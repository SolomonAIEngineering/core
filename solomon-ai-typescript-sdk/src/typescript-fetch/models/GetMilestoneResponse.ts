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
import type { Milestone } from './Milestone';
import {
    MilestoneFromJSON,
    MilestoneFromJSONTyped,
    MilestoneToJSON,
} from './Milestone';

/**
 * 
 * @export
 * @interface GetMilestoneResponse
 */
export interface GetMilestoneResponse {
    /**
     * 
     * @type {Milestone}
     * @memberof GetMilestoneResponse
     */
    milestone?: Milestone;
}

/**
 * Check if a given object implements the GetMilestoneResponse interface.
 */
export function instanceOfGetMilestoneResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetMilestoneResponseFromJSON(json: any): GetMilestoneResponse {
    return GetMilestoneResponseFromJSONTyped(json, false);
}

export function GetMilestoneResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMilestoneResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'milestone': !exists(json, 'milestone') ? undefined : MilestoneFromJSON(json['milestone']),
    };
}

export function GetMilestoneResponseToJSON(value?: GetMilestoneResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'milestone': MilestoneToJSON(value.milestone),
    };
}

