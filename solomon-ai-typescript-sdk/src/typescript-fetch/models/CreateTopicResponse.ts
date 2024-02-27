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
 * @interface CreateTopicResponse
 */
export interface CreateTopicResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateTopicResponse
     */
    topicId?: string;
}

/**
 * Check if a given object implements the CreateTopicResponse interface.
 */
export function instanceOfCreateTopicResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateTopicResponseFromJSON(json: any): CreateTopicResponse {
    return CreateTopicResponseFromJSONTyped(json, false);
}

export function CreateTopicResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateTopicResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'topicId': !exists(json, 'topicId') ? undefined : json['topicId'],
    };
}

export function CreateTopicResponseToJSON(value?: CreateTopicResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'topicId': value.topicId,
    };
}

