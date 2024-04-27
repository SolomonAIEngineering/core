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
 * Defines a message named GetMergeLinkTokenRequest.
 * @export
 * @interface GetMergeLinkTokenRequest
 */
export interface GetMergeLinkTokenRequest {
    /**
     * This unique identifier typically represents the ID for your end user across all services.
     * This value must be distinct from other Linked Accounts' unique identifiers.
     * @type {string}
     * @memberof GetMergeLinkTokenRequest
     */
    authZeroUserId: string;
    /**
     * Your end user's organization.
     * @type {string}
     * @memberof GetMergeLinkTokenRequest
     */
    organizationName: string;
    /**
     * Your end user's email address. This is purely for
     * identification purposes - setting this value will not cause any emails to be sent.
     * @type {string}
     * @memberof GetMergeLinkTokenRequest
     */
    email: string;
}

/**
 * Check if a given object implements the GetMergeLinkTokenRequest interface.
 */
export function instanceOfGetMergeLinkTokenRequest(value: object): boolean {
    if (!('authZeroUserId' in value)) return false;
    if (!('organizationName' in value)) return false;
    if (!('email' in value)) return false;
    return true;
}

export function GetMergeLinkTokenRequestFromJSON(json: any): GetMergeLinkTokenRequest {
    return GetMergeLinkTokenRequestFromJSONTyped(json, false);
}

export function GetMergeLinkTokenRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMergeLinkTokenRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'authZeroUserId': json['authZeroUserId'],
        'organizationName': json['organizationName'],
        'email': json['email'],
    };
}

export function GetMergeLinkTokenRequestToJSON(value?: GetMergeLinkTokenRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'authZeroUserId': value['authZeroUserId'],
        'organizationName': value['organizationName'],
        'email': value['email'],
    };
}

