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
 * @interface ReadIncomeStatementsRequest
 */
export interface ReadIncomeStatementsRequest {
    /**
     * 
     * @type {string}
     * @memberof ReadIncomeStatementsRequest
     */
    authZeroUserId: string;
    /**
     * 
     * @type {string}
     * @memberof ReadIncomeStatementsRequest
     */
    acountingIntegrationMergeLinkId: string;
    /**
     * 
     * @type {Date}
     * @memberof ReadIncomeStatementsRequest
     */
    startDate: Date;
    /**
     * 
     * @type {Date}
     * @memberof ReadIncomeStatementsRequest
     */
    endDate: Date;
    /**
     * 
     * @type {number}
     * @memberof ReadIncomeStatementsRequest
     */
    pageSize: number;
    /**
     * 
     * @type {number}
     * @memberof ReadIncomeStatementsRequest
     */
    pageNumber: number;
}

/**
 * Check if a given object implements the ReadIncomeStatementsRequest interface.
 */
export function instanceOfReadIncomeStatementsRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "authZeroUserId" in value;
    isInstance = isInstance && "acountingIntegrationMergeLinkId" in value;
    isInstance = isInstance && "startDate" in value;
    isInstance = isInstance && "endDate" in value;
    isInstance = isInstance && "pageSize" in value;
    isInstance = isInstance && "pageNumber" in value;

    return isInstance;
}

export function ReadIncomeStatementsRequestFromJSON(json: any): ReadIncomeStatementsRequest {
    return ReadIncomeStatementsRequestFromJSONTyped(json, false);
}

export function ReadIncomeStatementsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReadIncomeStatementsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authZeroUserId': json['authZeroUserId'],
        'acountingIntegrationMergeLinkId': json['acountingIntegrationMergeLinkId'],
        'startDate': (new Date(json['startDate'])),
        'endDate': (new Date(json['endDate'])),
        'pageSize': json['pageSize'],
        'pageNumber': json['pageNumber'],
    };
}

export function ReadIncomeStatementsRequestToJSON(value?: ReadIncomeStatementsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authZeroUserId': value.authZeroUserId,
        'acountingIntegrationMergeLinkId': value.acountingIntegrationMergeLinkId,
        'startDate': (value.startDate.toISOString()),
        'endDate': (value.endDate.toISOString()),
        'pageSize': value.pageSize,
        'pageNumber': value.pageNumber,
    };
}
