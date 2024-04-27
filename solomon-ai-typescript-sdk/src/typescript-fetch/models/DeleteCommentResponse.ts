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
 * @interface DeleteCommentResponse
 */
export interface DeleteCommentResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DeleteCommentResponse
     */
    sucess?: boolean;
}

/**
 * Check if a given object implements the DeleteCommentResponse interface.
 */
export function instanceOfDeleteCommentResponse(value: object): boolean {
    return true;
}

export function DeleteCommentResponseFromJSON(json: any): DeleteCommentResponse {
    return DeleteCommentResponseFromJSONTyped(json, false);
}

export function DeleteCommentResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeleteCommentResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'sucess': json['sucess'] == null ? undefined : json['sucess'],
    };
}

export function DeleteCommentResponseToJSON(value?: DeleteCommentResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'sucess': value['sucess'],
    };
}

