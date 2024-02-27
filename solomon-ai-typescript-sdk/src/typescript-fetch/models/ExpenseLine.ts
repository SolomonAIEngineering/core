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
 * The ExpenseLine object is used to represent an expense's line items.
 * @export
 * @interface ExpenseLine
 */
export interface ExpenseLine {
    /**
     * 
     * @type {string}
     * @memberof ExpenseLine
     */
    id?: string;
    /**
     * The line's net amount.
     * @type {number}
     * @memberof ExpenseLine
     */
    netAmount?: number;
    /**
     * 
     * @type {string}
     * @memberof ExpenseLine
     */
    trackingCategory?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ExpenseLine
     */
    trackingCategories?: Array<string>;
    /**
     * The company the line belongs to.
     * @type {string}
     * @memberof ExpenseLine
     */
    company?: string;
    /**
     * The line's item.
     * 
     * This seems to be an ID
     * @type {string}
     * @memberof ExpenseLine
     */
    item?: string;
    /**
     * The expense's payment account.
     * @type {string}
     * @memberof ExpenseLine
     */
    account?: string;
    /**
     * The expense's contact.
     * 
     * Optional based on provided JSON
     * @type {string}
     * @memberof ExpenseLine
     */
    contact?: string;
    /**
     * The description of the item that was purchased by the company.
     * @type {string}
     * @memberof ExpenseLine
     */
    description?: string;
    /**
     * The expense line item's exchange rate.
     * 
     * Consider using double or float if this represents a number
     * @type {string}
     * @memberof ExpenseLine
     */
    exchangeRate?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof ExpenseLine
     */
    remoteId?: string;
    /**
     * 
     * @type {string}
     * @memberof ExpenseLine
     */
    currency?: string;
    /**
     * Consider using google.protobuf.Timestamp
     * @type {Date}
     * @memberof ExpenseLine
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {boolean}
     * @memberof ExpenseLine
     */
    remoteWasDeleted?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof ExpenseLine
     */
    createdAt?: Date;
}

/**
 * Check if a given object implements the ExpenseLine interface.
 */
export function instanceOfExpenseLine(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ExpenseLineFromJSON(json: any): ExpenseLine {
    return ExpenseLineFromJSONTyped(json, false);
}

export function ExpenseLineFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExpenseLine {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'netAmount': !exists(json, 'netAmount') ? undefined : json['netAmount'],
        'trackingCategory': !exists(json, 'trackingCategory') ? undefined : json['trackingCategory'],
        'trackingCategories': !exists(json, 'trackingCategories') ? undefined : json['trackingCategories'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'item': !exists(json, 'item') ? undefined : json['item'],
        'account': !exists(json, 'account') ? undefined : json['account'],
        'contact': !exists(json, 'contact') ? undefined : json['contact'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : json['exchangeRate'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'currency': !exists(json, 'currency') ? undefined : json['currency'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
    };
}

export function ExpenseLineToJSON(value?: ExpenseLine | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'netAmount': value.netAmount,
        'trackingCategory': value.trackingCategory,
        'trackingCategories': value.trackingCategories,
        'company': value.company,
        'item': value.item,
        'account': value.account,
        'contact': value.contact,
        'description': value.description,
        'exchangeRate': value.exchangeRate,
        'remoteId': value.remoteId,
        'currency': value.currency,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'remoteWasDeleted': value.remoteWasDeleted,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
    };
}

