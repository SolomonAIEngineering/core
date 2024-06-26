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
import type { Team } from './Team';
import {
    TeamFromJSON,
    TeamFromJSONTyped,
    TeamToJSON,
} from './Team';

/**
 * 
 * @export
 * @interface UpdateTeamResponse
 */
export interface UpdateTeamResponse {
    /**
     * 
     * @type {Team}
     * @memberof UpdateTeamResponse
     */
    team?: Team;
}

/**
 * Check if a given object implements the UpdateTeamResponse interface.
 */
export function instanceOfUpdateTeamResponse(value: object): boolean {
    return true;
}

export function UpdateTeamResponseFromJSON(json: any): UpdateTeamResponse {
    return UpdateTeamResponseFromJSONTyped(json, false);
}

export function UpdateTeamResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateTeamResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'team': json['team'] == null ? undefined : TeamFromJSON(json['team']),
    };
}

export function UpdateTeamResponseToJSON(value?: UpdateTeamResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'team': TeamToJSON(value['team']),
    };
}

