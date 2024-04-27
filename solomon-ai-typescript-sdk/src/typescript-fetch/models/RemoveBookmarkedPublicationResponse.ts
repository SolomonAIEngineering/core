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
import type { Bookmark } from './Bookmark';
import {
    BookmarkFromJSON,
    BookmarkFromJSONTyped,
    BookmarkToJSON,
} from './Bookmark';

/**
 * 
 * @export
 * @interface RemoveBookmarkedPublicationResponse
 */
export interface RemoveBookmarkedPublicationResponse {
    /**
     * 
     * @type {Bookmark}
     * @memberof RemoveBookmarkedPublicationResponse
     */
    bookmark?: Bookmark;
}

/**
 * Check if a given object implements the RemoveBookmarkedPublicationResponse interface.
 */
export function instanceOfRemoveBookmarkedPublicationResponse(value: object): boolean {
    return true;
}

export function RemoveBookmarkedPublicationResponseFromJSON(json: any): RemoveBookmarkedPublicationResponse {
    return RemoveBookmarkedPublicationResponseFromJSONTyped(json, false);
}

export function RemoveBookmarkedPublicationResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RemoveBookmarkedPublicationResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'bookmark': json['bookmark'] == null ? undefined : BookmarkFromJSON(json['bookmark']),
    };
}

export function RemoveBookmarkedPublicationResponseToJSON(value?: RemoveBookmarkedPublicationResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'bookmark': BookmarkToJSON(value['bookmark']),
    };
}

