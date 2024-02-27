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
import type { StudentLoanAccount } from './StudentLoanAccount';
import {
    StudentLoanAccountFromJSON,
    StudentLoanAccountFromJSONTyped,
    StudentLoanAccountToJSON,
} from './StudentLoanAccount';

/**
 * 
 * @export
 * @interface GetStudentLoanAccountResponse
 */
export interface GetStudentLoanAccountResponse {
    /**
     * 
     * @type {StudentLoanAccount}
     * @memberof GetStudentLoanAccountResponse
     */
    studentLoanAccount?: StudentLoanAccount;
}

/**
 * Check if a given object implements the GetStudentLoanAccountResponse interface.
 */
export function instanceOfGetStudentLoanAccountResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetStudentLoanAccountResponseFromJSON(json: any): GetStudentLoanAccountResponse {
    return GetStudentLoanAccountResponseFromJSONTyped(json, false);
}

export function GetStudentLoanAccountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetStudentLoanAccountResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'studentLoanAccount': !exists(json, 'studentLoanAccount') ? undefined : StudentLoanAccountFromJSON(json['studentLoanAccount']),
    };
}

export function GetStudentLoanAccountResponseToJSON(value?: GetStudentLoanAccountResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'studentLoanAccount': StudentLoanAccountToJSON(value.studentLoanAccount),
    };
}

