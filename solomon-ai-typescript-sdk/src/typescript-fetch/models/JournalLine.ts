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
 * The JournalLine object is used to represent a journal entry's line items.
 * @export
 * @interface JournalLine
 */
export interface JournalLine {
    /**
     * 
     * @type {string}
     * @memberof JournalLine
     */
    id?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof JournalLine
     */
    remoteId?: string;
    /**
     * 
     * @type {string}
     * @memberof JournalLine
     */
    account?: string;
    /**
     * The value of the line item including taxes and other fees.
     * 
     * Might want to use double or a more precise type
     * @type {number}
     * @memberof JournalLine
     */
    netAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof JournalLine
     */
    trackingCategory?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof JournalLine
     */
    trackingCategories?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof JournalLine
     */
    contact?: string;
    /**
     * The company the journal entry belongs to.
     * @type {string}
     * @memberof JournalLine
     */
    company?: string;
    /**
     * The line's description.
     * @type {string}
     * @memberof JournalLine
     */
    description?: string;
    /**
     * The journal line item's exchange rate.
     * 
     * Assuming string due to the example provided, but could be float or double.
     * @type {string}
     * @memberof JournalLine
     */
    exchangeRate?: string;
    /**
     * 
     * @type {Date}
     * @memberof JournalLine
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof JournalLine
     */
    createdAt?: Date;
}

/**
 * Check if a given object implements the JournalLine interface.
 */
export function instanceOfJournalLine(value: object): boolean {
    return true;
}

export function JournalLineFromJSON(json: any): JournalLine {
    return JournalLineFromJSONTyped(json, false);
}

export function JournalLineFromJSONTyped(json: any, ignoreDiscriminator: boolean): JournalLine {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'remoteId': json['remoteId'] == null ? undefined : json['remoteId'],
        'account': json['account'] == null ? undefined : json['account'],
        'netAmount': json['netAmount'] == null ? undefined : json['netAmount'],
        'trackingCategory': json['trackingCategory'] == null ? undefined : json['trackingCategory'],
        'trackingCategories': json['trackingCategories'] == null ? undefined : json['trackingCategories'],
        'contact': json['contact'] == null ? undefined : json['contact'],
        'company': json['company'] == null ? undefined : json['company'],
        'description': json['description'] == null ? undefined : json['description'],
        'exchangeRate': json['exchangeRate'] == null ? undefined : json['exchangeRate'],
        'modifiedAt': json['modifiedAt'] == null ? undefined : (new Date(json['modifiedAt'])),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
    };
}

export function JournalLineToJSON(value?: JournalLine | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'remoteId': value['remoteId'],
        'account': value['account'],
        'netAmount': value['netAmount'],
        'trackingCategory': value['trackingCategory'],
        'trackingCategories': value['trackingCategories'],
        'contact': value['contact'],
        'company': value['company'],
        'description': value['description'],
        'exchangeRate': value['exchangeRate'],
        'modifiedAt': value['modifiedAt'] == null ? undefined : ((value['modifiedAt']).toISOString()),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
    };
}

