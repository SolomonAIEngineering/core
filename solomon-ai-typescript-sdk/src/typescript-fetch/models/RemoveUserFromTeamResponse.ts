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
import type { Team } from './Team';
import {
    TeamFromJSON,
    TeamFromJSONTyped,
    TeamToJSON,
} from './Team';

/**
 * 
 * @export
 * @interface RemoveUserFromTeamResponse
 */
export interface RemoveUserFromTeamResponse {
    /**
     * 
     * @type {Team}
     * @memberof RemoveUserFromTeamResponse
     */
    team?: Team;
}

/**
 * Check if a given object implements the RemoveUserFromTeamResponse interface.
 */
export function instanceOfRemoveUserFromTeamResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RemoveUserFromTeamResponseFromJSON(json: any): RemoveUserFromTeamResponse {
    return RemoveUserFromTeamResponseFromJSONTyped(json, false);
}

export function RemoveUserFromTeamResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RemoveUserFromTeamResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'team': !exists(json, 'team') ? undefined : TeamFromJSON(json['team']),
    };
}

export function RemoveUserFromTeamResponseToJSON(value?: RemoveUserFromTeamResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'team': TeamToJSON(value.team),
    };
}

