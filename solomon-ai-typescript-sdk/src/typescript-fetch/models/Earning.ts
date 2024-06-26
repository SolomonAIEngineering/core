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
import type { EarningType } from './EarningType';
import {
    EarningTypeFromJSON,
    EarningTypeFromJSONTyped,
    EarningTypeToJSON,
} from './EarningType';

/**
 * The Earning object is used to represent an array of different compensations 
 * that an employee receives within specific wage categories.
 * @export
 * @interface Earning
 */
export interface Earning {
    /**
     * 
     * @type {string}
     * @memberof Earning
     */
    id?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof Earning
     */
    remoteId?: string;
    /**
     * 
     * @type {number}
     * @memberof Earning
     */
    amount?: number;
    /**
     * 
     * @type {EarningType}
     * @memberof Earning
     */
    type?: EarningType;
    /**
     * Indicates whether or not this object has been deleted in the third party platform.
     * @type {boolean}
     * @memberof Earning
     */
    remoteWasDeleted?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof Earning
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Earning
     */
    modifiedAt?: Date;
    /**
     * The payroll being run.
     * @type {string}
     * @memberof Earning
     */
    payrollRunMergeAccountId?: string;
    /**
     * 
     * @type {string}
     * @memberof Earning
     */
    mergeAccountId?: string;
}

/**
 * Check if a given object implements the Earning interface.
 */
export function instanceOfEarning(value: object): boolean {
    return true;
}

export function EarningFromJSON(json: any): Earning {
    return EarningFromJSONTyped(json, false);
}

export function EarningFromJSONTyped(json: any, ignoreDiscriminator: boolean): Earning {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'remoteId': json['remoteId'] == null ? undefined : json['remoteId'],
        'amount': json['amount'] == null ? undefined : json['amount'],
        'type': json['type'] == null ? undefined : EarningTypeFromJSON(json['type']),
        'remoteWasDeleted': json['remoteWasDeleted'] == null ? undefined : json['remoteWasDeleted'],
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'modifiedAt': json['modifiedAt'] == null ? undefined : (new Date(json['modifiedAt'])),
        'payrollRunMergeAccountId': json['payrollRunMergeAccountId'] == null ? undefined : json['payrollRunMergeAccountId'],
        'mergeAccountId': json['mergeAccountId'] == null ? undefined : json['mergeAccountId'],
    };
}

export function EarningToJSON(value?: Earning | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'remoteId': value['remoteId'],
        'amount': value['amount'],
        'type': EarningTypeToJSON(value['type']),
        'remoteWasDeleted': value['remoteWasDeleted'],
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'modifiedAt': value['modifiedAt'] == null ? undefined : ((value['modifiedAt']).toISOString()),
        'payrollRunMergeAccountId': value['payrollRunMergeAccountId'],
        'mergeAccountId': value['mergeAccountId'],
    };
}

