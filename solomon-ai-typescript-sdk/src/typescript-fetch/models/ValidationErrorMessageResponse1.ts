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
import type { ErrorCode1 } from './ErrorCode1';
import {
    ErrorCode1FromJSON,
    ErrorCode1FromJSONTyped,
    ErrorCode1ToJSON,
} from './ErrorCode1';

/**
 * 
 * @export
 * @interface ValidationErrorMessageResponse1
 */
export interface ValidationErrorMessageResponse1 {
    /**
     * 
     * @type {ErrorCode1}
     * @memberof ValidationErrorMessageResponse1
     */
    code?: ErrorCode1;
    /**
     * 
     * @type {string}
     * @memberof ValidationErrorMessageResponse1
     */
    message?: string;
}

/**
 * Check if a given object implements the ValidationErrorMessageResponse1 interface.
 */
export function instanceOfValidationErrorMessageResponse1(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ValidationErrorMessageResponse1FromJSON(json: any): ValidationErrorMessageResponse1 {
    return ValidationErrorMessageResponse1FromJSONTyped(json, false);
}

export function ValidationErrorMessageResponse1FromJSONTyped(json: any, ignoreDiscriminator: boolean): ValidationErrorMessageResponse1 {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : ErrorCode1FromJSON(json['code']),
        'message': !exists(json, 'message') ? undefined : json['message'],
    };
}

export function ValidationErrorMessageResponse1ToJSON(value?: ValidationErrorMessageResponse1 | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': ErrorCode1ToJSON(value.code),
        'message': value.message,
    };
}

