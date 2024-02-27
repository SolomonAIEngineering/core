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
 * @interface CreateMilestoneRequest
 */
export interface CreateMilestoneRequest {
    /**
     * 
     * @type {string}
     * @memberof CreateMilestoneRequest
     */
    smartGoalId: string;
    /**
     * 
     * @type {Milestone}
     * @memberof CreateMilestoneRequest
     */
    milestone: Milestone;
}

/**
 * Check if a given object implements the CreateMilestoneRequest interface.
 */
export function instanceOfCreateMilestoneRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "smartGoalId" in value;
    isInstance = isInstance && "milestone" in value;

    return isInstance;
}

export function CreateMilestoneRequestFromJSON(json: any): CreateMilestoneRequest {
    return CreateMilestoneRequestFromJSONTyped(json, false);
}

export function CreateMilestoneRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateMilestoneRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'smartGoalId': json['smartGoalId'],
        'milestone': MilestoneFromJSON(json['milestone']),
    };
}

export function CreateMilestoneRequestToJSON(value?: CreateMilestoneRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'smartGoalId': value.smartGoalId,
        'milestone': MilestoneToJSON(value.milestone),
    };
}

