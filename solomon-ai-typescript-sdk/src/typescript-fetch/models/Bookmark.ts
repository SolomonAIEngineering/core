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
import type { Publication } from './Publication';
import {
    PublicationFromJSON,
    PublicationFromJSONTyped,
    PublicationToJSON,
} from './Publication';

/**
 * 
 * @export
 * @interface Bookmark
 */
export interface Bookmark {
    /**
     * 
     * @type {string}
     * @memberof Bookmark
     */
    id?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof Bookmark
     */
    postIds?: Array<string>;
    /**
     * 
     * @type {Array<Publication>}
     * @memberof Bookmark
     */
    publications?: Array<Publication>;
}

/**
 * Check if a given object implements the Bookmark interface.
 */
export function instanceOfBookmark(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BookmarkFromJSON(json: any): Bookmark {
    return BookmarkFromJSONTyped(json, false);
}

export function BookmarkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Bookmark {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'postIds': !exists(json, 'postIds') ? undefined : json['postIds'],
        'publications': !exists(json, 'publications') ? undefined : ((json['publications'] as Array<any>).map(PublicationFromJSON)),
    };
}

export function BookmarkToJSON(value?: Bookmark | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'postIds': value.postIds,
        'publications': value.publications === undefined ? undefined : ((value.publications as Array<any>).map(PublicationToJSON)),
    };
}

