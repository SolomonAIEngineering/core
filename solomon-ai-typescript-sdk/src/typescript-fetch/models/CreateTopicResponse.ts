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
    return true;
}

export function CreateTopicResponseFromJSON(json: any): CreateTopicResponse {
    return CreateTopicResponseFromJSONTyped(json, false);
}

export function CreateTopicResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateTopicResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'topicId': json['topicId'] == null ? undefined : json['topicId'],
    };
}

export function CreateTopicResponseToJSON(value?: CreateTopicResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'topicId': value['topicId'],
    };
}

