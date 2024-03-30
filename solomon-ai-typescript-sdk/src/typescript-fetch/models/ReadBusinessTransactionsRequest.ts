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
 * Request for reading business transactions with pagination and time filtering.
 * @export
 * @interface ReadBusinessTransactionsRequest
 */
export interface ReadBusinessTransactionsRequest {
    /**
     * The ID of the user whose transactions are to be fetched.
     * @type {string}
     * @memberof ReadBusinessTransactionsRequest
     */
    authZeroUserId: string;
    /**
     * The page number of the paginated results.
     * @type {string}
     * @memberof ReadBusinessTransactionsRequest
     */
    pageNumber: string;
    /**
     * The number of items to be returned per page.
     * @type {string}
     * @memberof ReadBusinessTransactionsRequest
     */
    pageSize: string;
    /**
     * The start of the time range for the transactions (inclusive).
     * @type {string}
     * @memberof ReadBusinessTransactionsRequest
     */
    startTime?: string;
    /**
     * The end of the time range for the transactions (exclusive).
     * @type {string}
     * @memberof ReadBusinessTransactionsRequest
     */
    endTime?: string;
}

/**
 * Check if a given object implements the ReadBusinessTransactionsRequest interface.
 */
export function instanceOfReadBusinessTransactionsRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "authZeroUserId" in value;
    isInstance = isInstance && "pageNumber" in value;
    isInstance = isInstance && "pageSize" in value;

    return isInstance;
}

export function ReadBusinessTransactionsRequestFromJSON(json: any): ReadBusinessTransactionsRequest {
    return ReadBusinessTransactionsRequestFromJSONTyped(json, false);
}

export function ReadBusinessTransactionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReadBusinessTransactionsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authZeroUserId': json['authZeroUserId'],
        'pageNumber': json['pageNumber'],
        'pageSize': json['pageSize'],
        'startTime': !exists(json, 'startTime') ? undefined : json['startTime'],
        'endTime': !exists(json, 'endTime') ? undefined : json['endTime'],
    };
}

export function ReadBusinessTransactionsRequestToJSON(value?: ReadBusinessTransactionsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authZeroUserId': value.authZeroUserId,
        'pageNumber': value.pageNumber,
        'pageSize': value.pageSize,
        'startTime': value.startTime,
        'endTime': value.endTime,
    };
}

