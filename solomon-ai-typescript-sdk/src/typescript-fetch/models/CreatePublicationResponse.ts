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
 * @interface CreatePublicationResponse
 */
export interface CreatePublicationResponse {
    /**
     * 
     * @type {string}
     * @memberof CreatePublicationResponse
     */
    id?: string;
}

/**
 * Check if a given object implements the CreatePublicationResponse interface.
 */
export function instanceOfCreatePublicationResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreatePublicationResponseFromJSON(json: any): CreatePublicationResponse {
    return CreatePublicationResponseFromJSONTyped(json, false);
}

export function CreatePublicationResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreatePublicationResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function CreatePublicationResponseToJSON(value?: CreatePublicationResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
    };
}

