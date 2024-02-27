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
 * @interface NotificationActivity
 */
export interface NotificationActivity {
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    actorName?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    foreignId?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    activityId?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    verb?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    time?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    target?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    origin?: string;
    /**
     * 
     * @type {string}
     * @memberof NotificationActivity
     */
    object?: string;
}

/**
 * Check if a given object implements the NotificationActivity interface.
 */
export function instanceOfNotificationActivity(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function NotificationActivityFromJSON(json: any): NotificationActivity {
    return NotificationActivityFromJSONTyped(json, false);
}

export function NotificationActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): NotificationActivity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'actorName': !exists(json, 'actorName') ? undefined : json['actorName'],
        'foreignId': !exists(json, 'foreignId') ? undefined : json['foreignId'],
        'activityId': !exists(json, 'activityId') ? undefined : json['activityId'],
        'verb': !exists(json, 'verb') ? undefined : json['verb'],
        'time': !exists(json, 'time') ? undefined : json['time'],
        'target': !exists(json, 'target') ? undefined : json['target'],
        'origin': !exists(json, 'origin') ? undefined : json['origin'],
        'object': !exists(json, 'object') ? undefined : json['object'],
    };
}

export function NotificationActivityToJSON(value?: NotificationActivity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'actorName': value.actorName,
        'foreignId': value.foreignId,
        'activityId': value.activityId,
        'verb': value.verb,
        'time': value.time,
        'target': value.target,
        'origin': value.origin,
        'object': value.object,
    };
}

