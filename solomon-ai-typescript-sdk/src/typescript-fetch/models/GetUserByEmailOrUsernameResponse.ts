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
import type { UserAccount } from './UserAccount';
import {
    UserAccountFromJSON,
    UserAccountFromJSONTyped,
    UserAccountToJSON,
} from './UserAccount';

/**
 * 
 * @export
 * @interface GetUserByEmailOrUsernameResponse
 */
export interface GetUserByEmailOrUsernameResponse {
    /**
     * 
     * @type {UserAccount}
     * @memberof GetUserByEmailOrUsernameResponse
     */
    account?: UserAccount;
}

/**
 * Check if a given object implements the GetUserByEmailOrUsernameResponse interface.
 */
export function instanceOfGetUserByEmailOrUsernameResponse(value: object): boolean {
    return true;
}

export function GetUserByEmailOrUsernameResponseFromJSON(json: any): GetUserByEmailOrUsernameResponse {
    return GetUserByEmailOrUsernameResponseFromJSONTyped(json, false);
}

export function GetUserByEmailOrUsernameResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetUserByEmailOrUsernameResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'account': json['account'] == null ? undefined : UserAccountFromJSON(json['account']),
    };
}

export function GetUserByEmailOrUsernameResponseToJSON(value?: GetUserByEmailOrUsernameResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'account': UserAccountToJSON(value['account']),
    };
}

