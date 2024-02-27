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
import type { PocketType } from './PocketType';
import {
    PocketTypeFromJSON,
    PocketTypeFromJSONTyped,
    PocketTypeToJSON,
} from './PocketType';
import type { SmartGoal } from './SmartGoal';
import {
    SmartGoalFromJSON,
    SmartGoalFromJSONTyped,
    SmartGoalToJSON,
} from './SmartGoal';

/**
 * 
 * @export
 * @interface Pocket
 */
export interface Pocket {
    /**
     * 
     * @type {string}
     * @memberof Pocket
     */
    id?: string;
    /**
     * 
     * @type {Array<SmartGoal>}
     * @memberof Pocket
     */
    goals?: Array<SmartGoal>;
    /**
     * 
     * @type {PocketType}
     * @memberof Pocket
     */
    type?: PocketType;
    /**
     * 
     * @type {Array<string>}
     * @memberof Pocket
     */
    tags?: Array<string>;
}

/**
 * Check if a given object implements the Pocket interface.
 */
export function instanceOfPocket(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PocketFromJSON(json: any): Pocket {
    return PocketFromJSONTyped(json, false);
}

export function PocketFromJSONTyped(json: any, ignoreDiscriminator: boolean): Pocket {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'goals': !exists(json, 'goals') ? undefined : ((json['goals'] as Array<any>).map(SmartGoalFromJSON)),
        'type': !exists(json, 'type') ? undefined : PocketTypeFromJSON(json['type']),
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
    };
}

export function PocketToJSON(value?: Pocket | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'goals': value.goals === undefined ? undefined : ((value.goals as Array<any>).map(SmartGoalToJSON)),
        'type': PocketTypeToJSON(value.type),
        'tags': value.tags,
    };
}

