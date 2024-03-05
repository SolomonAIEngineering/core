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
import type { NotFoundErrorCode } from './NotFoundErrorCode';
import {
    NotFoundErrorCodeFromJSON,
    NotFoundErrorCodeFromJSONTyped,
    NotFoundErrorCodeToJSON,
} from './NotFoundErrorCode';

/**
 * 
 * @export
 * @interface PathUnknownErrorMessageResponse
 */
export interface PathUnknownErrorMessageResponse {
    /**
     * 
     * @type {NotFoundErrorCode}
     * @memberof PathUnknownErrorMessageResponse
     */
    code?: NotFoundErrorCode;
    /**
     * 
     * @type {string}
     * @memberof PathUnknownErrorMessageResponse
     */
    message?: string;
}

/**
 * Check if a given object implements the PathUnknownErrorMessageResponse interface.
 */
export function instanceOfPathUnknownErrorMessageResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PathUnknownErrorMessageResponseFromJSON(json: any): PathUnknownErrorMessageResponse {
    return PathUnknownErrorMessageResponseFromJSONTyped(json, false);
}

export function PathUnknownErrorMessageResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PathUnknownErrorMessageResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : NotFoundErrorCodeFromJSON(json['code']),
        'message': !exists(json, 'message') ? undefined : json['message'],
    };
}

export function PathUnknownErrorMessageResponseToJSON(value?: PathUnknownErrorMessageResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': NotFoundErrorCodeToJSON(value.code),
        'message': value.message,
    };
}
