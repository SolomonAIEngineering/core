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
import type { FinancialAccountType } from './FinancialAccountType';
import {
    FinancialAccountTypeFromJSON,
    FinancialAccountTypeFromJSONTyped,
    FinancialAccountTypeToJSON,
} from './FinancialAccountType';
import type { FinancialUserProfileType } from './FinancialUserProfileType';
import {
    FinancialUserProfileTypeFromJSON,
    FinancialUserProfileTypeFromJSONTyped,
    FinancialUserProfileTypeToJSON,
} from './FinancialUserProfileType';

/**
 * 
 * @export
 * @interface AddDefaultPocketsToBankAccountRequest
 */
export interface AddDefaultPocketsToBankAccountRequest {
    /**
     * 
     * @type {string}
     * @memberof AddDefaultPocketsToBankAccountRequest
     */
    userId: string;
    /**
     * 
     * @type {string}
     * @memberof AddDefaultPocketsToBankAccountRequest
     */
    plaidAccountId: string;
    /**
     * 
     * @type {FinancialUserProfileType}
     * @memberof AddDefaultPocketsToBankAccountRequest
     */
    profileType: FinancialUserProfileType;
    /**
     * 
     * @type {FinancialAccountType}
     * @memberof AddDefaultPocketsToBankAccountRequest
     */
    financialAccountType: FinancialAccountType;
}

/**
 * Check if a given object implements the AddDefaultPocketsToBankAccountRequest interface.
 */
export function instanceOfAddDefaultPocketsToBankAccountRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "userId" in value;
    isInstance = isInstance && "plaidAccountId" in value;
    isInstance = isInstance && "profileType" in value;
    isInstance = isInstance && "financialAccountType" in value;

    return isInstance;
}

export function AddDefaultPocketsToBankAccountRequestFromJSON(json: any): AddDefaultPocketsToBankAccountRequest {
    return AddDefaultPocketsToBankAccountRequestFromJSONTyped(json, false);
}

export function AddDefaultPocketsToBankAccountRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddDefaultPocketsToBankAccountRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userId': json['userId'],
        'plaidAccountId': json['plaidAccountId'],
        'profileType': FinancialUserProfileTypeFromJSON(json['profileType']),
        'financialAccountType': FinancialAccountTypeFromJSON(json['financialAccountType']),
    };
}

export function AddDefaultPocketsToBankAccountRequestToJSON(value?: AddDefaultPocketsToBankAccountRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'userId': value.userId,
        'plaidAccountId': value.plaidAccountId,
        'profileType': FinancialUserProfileTypeToJSON(value.profileType),
        'financialAccountType': FinancialAccountTypeToJSON(value.financialAccountType),
    };
}
