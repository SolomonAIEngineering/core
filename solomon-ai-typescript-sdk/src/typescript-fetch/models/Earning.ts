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
    let isInstance = true;

    return isInstance;
}

export function EarningFromJSON(json: any): Earning {
    return EarningFromJSONTyped(json, false);
}

export function EarningFromJSONTyped(json: any, ignoreDiscriminator: boolean): Earning {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'type': !exists(json, 'type') ? undefined : EarningTypeFromJSON(json['type']),
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'payrollRunMergeAccountId': !exists(json, 'payrollRunMergeAccountId') ? undefined : json['payrollRunMergeAccountId'],
        'mergeAccountId': !exists(json, 'mergeAccountId') ? undefined : json['mergeAccountId'],
    };
}

export function EarningToJSON(value?: Earning | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'amount': value.amount,
        'type': EarningTypeToJSON(value.type),
        'remoteWasDeleted': value.remoteWasDeleted,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'payrollRunMergeAccountId': value.payrollRunMergeAccountId,
        'mergeAccountId': value.mergeAccountId,
    };
}
