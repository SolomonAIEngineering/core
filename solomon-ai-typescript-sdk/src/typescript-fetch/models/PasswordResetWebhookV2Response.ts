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
 * @interface PasswordResetWebhookV2Response
 */
export interface PasswordResetWebhookV2Response {
    /**
     * 
     * @type {boolean}
     * @memberof PasswordResetWebhookV2Response
     */
    success?: boolean;
}

/**
 * Check if a given object implements the PasswordResetWebhookV2Response interface.
 */
export function instanceOfPasswordResetWebhookV2Response(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PasswordResetWebhookV2ResponseFromJSON(json: any): PasswordResetWebhookV2Response {
    return PasswordResetWebhookV2ResponseFromJSONTyped(json, false);
}

export function PasswordResetWebhookV2ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PasswordResetWebhookV2Response {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'success': !exists(json, 'success') ? undefined : json['success'],
    };
}

export function PasswordResetWebhookV2ResponseToJSON(value?: PasswordResetWebhookV2Response | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'success': value.success,
    };
}

