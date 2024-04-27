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
 * @interface Topic
 */
export interface Topic {
    /**
     * 
     * @type {string}
     * @memberof Topic
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Topic
     */
    topicName: string;
    /**
     * 
     * @type {string}
     * @memberof Topic
     */
    description: string;
    /**
     * 
     * @type {string}
     * @memberof Topic
     */
    imageUrl: string;
}

/**
 * Check if a given object implements the Topic interface.
 */
export function instanceOfTopic(value: object): boolean {
    if (!('topicName' in value)) return false;
    if (!('description' in value)) return false;
    if (!('imageUrl' in value)) return false;
    return true;
}

export function TopicFromJSON(json: any): Topic {
    return TopicFromJSONTyped(json, false);
}

export function TopicFromJSONTyped(json: any, ignoreDiscriminator: boolean): Topic {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'topicName': json['topicName'],
        'description': json['description'],
        'imageUrl': json['imageUrl'],
    };
}

export function TopicToJSON(value?: Topic | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'topicName': value['topicName'],
        'description': value['description'],
        'imageUrl': value['imageUrl'],
    };
}

