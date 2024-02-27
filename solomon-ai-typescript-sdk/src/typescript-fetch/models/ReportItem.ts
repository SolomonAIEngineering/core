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
 * The ReportItem object is used to represent a report item for a Balance Sheet, 
 * Cash Flow Statement or Profit and Loss Report.
 * @export
 * @interface ReportItem
 */
export interface ReportItem {
    /**
     * 
     * @type {string}
     * @memberof ReportItem
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ReportItem
     */
    remoteId?: string;
    /**
     * The report item's name.
     * @type {string}
     * @memberof ReportItem
     */
    name?: string;
    /**
     * The report item's value.
     * @type {string}
     * @memberof ReportItem
     */
    value?: string;
    /**
     * 
     * @type {string}
     * @memberof ReportItem
     */
    company?: string;
    /**
     * Consider using google.protobuf.Timestamp
     * @type {Date}
     * @memberof ReportItem
     */
    modifiedAt?: Date;
}

/**
 * Check if a given object implements the ReportItem interface.
 */
export function instanceOfReportItem(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReportItemFromJSON(json: any): ReportItem {
    return ReportItemFromJSONTyped(json, false);
}

export function ReportItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReportItem {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'value': !exists(json, 'value') ? undefined : json['value'],
        'company': !exists(json, 'company') ? undefined : json['company'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
    };
}

export function ReportItemToJSON(value?: ReportItem | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'name': value.name,
        'value': value.value,
        'company': value.company,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
    };
}

