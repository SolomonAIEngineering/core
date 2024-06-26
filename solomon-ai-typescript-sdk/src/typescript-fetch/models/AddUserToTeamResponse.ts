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
 * @interface AddUserToTeamResponse
 */
export interface AddUserToTeamResponse {
    /**
     * 
     * @type {Team}
     * @memberof AddUserToTeamResponse
     */
    team?: Team;
}

/**
 * Check if a given object implements the AddUserToTeamResponse interface.
 */
export function instanceOfAddUserToTeamResponse(value: object): boolean {
    return true;
}

export function AddUserToTeamResponseFromJSON(json: any): AddUserToTeamResponse {
    return AddUserToTeamResponseFromJSONTyped(json, false);
}

export function AddUserToTeamResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddUserToTeamResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'team': json['team'] == null ? undefined : TeamFromJSON(json['team']),
    };
}

export function AddUserToTeamResponseToJSON(value?: AddUserToTeamResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'team': TeamToJSON(value['team']),
    };
}

