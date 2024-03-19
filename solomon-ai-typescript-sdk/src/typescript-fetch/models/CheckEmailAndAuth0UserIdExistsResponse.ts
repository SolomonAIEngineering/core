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
import type { BusinessAccount } from './BusinessAccount';
import {
    BusinessAccountFromJSON,
    BusinessAccountFromJSONTyped,
    BusinessAccountToJSON,
} from './BusinessAccount';
import type { UserAccount } from './UserAccount';
import {
    UserAccountFromJSON,
    UserAccountFromJSONTyped,
    UserAccountToJSON,
} from './UserAccount';

/**
 * 
 * @export
 * @interface CheckEmailAndAuth0UserIdExistsResponse
 */
export interface CheckEmailAndAuth0UserIdExistsResponse {
    /**
     * 
     * @type {UserAccount}
     * @memberof CheckEmailAndAuth0UserIdExistsResponse
     */
    userAccount?: UserAccount;
    /**
     * 
     * @type {BusinessAccount}
     * @memberof CheckEmailAndAuth0UserIdExistsResponse
     */
    businessAccount?: BusinessAccount;
    /**
     * 
     * @type {boolean}
     * @memberof CheckEmailAndAuth0UserIdExistsResponse
     */
    _exists?: boolean;
}

/**
 * Check if a given object implements the CheckEmailAndAuth0UserIdExistsResponse interface.
 */
export function instanceOfCheckEmailAndAuth0UserIdExistsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CheckEmailAndAuth0UserIdExistsResponseFromJSON(json: any): CheckEmailAndAuth0UserIdExistsResponse {
    return CheckEmailAndAuth0UserIdExistsResponseFromJSONTyped(json, false);
}

export function CheckEmailAndAuth0UserIdExistsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CheckEmailAndAuth0UserIdExistsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userAccount': !exists(json, 'userAccount') ? undefined : UserAccountFromJSON(json['userAccount']),
        'businessAccount': !exists(json, 'businessAccount') ? undefined : BusinessAccountFromJSON(json['businessAccount']),
        '_exists': !exists(json, 'exists') ? undefined : json['exists'],
    };
}

export function CheckEmailAndAuth0UserIdExistsResponseToJSON(value?: CheckEmailAndAuth0UserIdExistsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userAccount': UserAccountToJSON(value.userAccount),
        'businessAccount': BusinessAccountToJSON(value.businessAccount),
        'exists': value._exists,
    };
}
