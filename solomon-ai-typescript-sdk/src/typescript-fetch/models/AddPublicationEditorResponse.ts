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
 * @interface AddPublicationEditorResponse
 */
export interface AddPublicationEditorResponse {
    /**
     * 
     * @type {Publication}
     * @memberof AddPublicationEditorResponse
     */
    publication?: Publication;
}

/**
 * Check if a given object implements the AddPublicationEditorResponse interface.
 */
export function instanceOfAddPublicationEditorResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AddPublicationEditorResponseFromJSON(json: any): AddPublicationEditorResponse {
    return AddPublicationEditorResponseFromJSONTyped(json, false);
}

export function AddPublicationEditorResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddPublicationEditorResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'publication': !exists(json, 'publication') ? undefined : PublicationFromJSON(json['publication']),
    };
}

export function AddPublicationEditorResponseToJSON(value?: AddPublicationEditorResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'publication': PublicationToJSON(value.publication),
    };
}

