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
import type { Post } from './Post';
import {
    PostFromJSON,
    PostFromJSONTyped,
    PostToJSON,
} from './Post';

/**
 * 
 * @export
 * @interface GetPostsByTopicResponse
 */
export interface GetPostsByTopicResponse {
    /**
     * 
     * @type {Array<Post>}
     * @memberof GetPostsByTopicResponse
     */
    posts?: Array<Post>;
}

/**
 * Check if a given object implements the GetPostsByTopicResponse interface.
 */
export function instanceOfGetPostsByTopicResponse(value: object): boolean {
    return true;
}

export function GetPostsByTopicResponseFromJSON(json: any): GetPostsByTopicResponse {
    return GetPostsByTopicResponseFromJSONTyped(json, false);
}

export function GetPostsByTopicResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetPostsByTopicResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'posts': json['posts'] == null ? undefined : ((json['posts'] as Array<any>).map(PostFromJSON)),
    };
}

export function GetPostsByTopicResponseToJSON(value?: GetPostsByTopicResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'posts': value['posts'] == null ? undefined : ((value['posts'] as Array<any>).map(PostToJSON)),
    };
}

