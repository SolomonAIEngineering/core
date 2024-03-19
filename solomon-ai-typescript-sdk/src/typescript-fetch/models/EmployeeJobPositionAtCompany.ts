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
import type { FlsaStatus } from './FlsaStatus';
import {
    FlsaStatusFromJSON,
    FlsaStatusFromJSONTyped,
    FlsaStatusToJSON,
} from './FlsaStatus';
import type { PayFrequency } from './PayFrequency';
import {
    PayFrequencyFromJSON,
    PayFrequencyFromJSONTyped,
    PayFrequencyToJSON,
} from './PayFrequency';
import type { PayPeriod } from './PayPeriod';
import {
    PayPeriodFromJSON,
    PayPeriodFromJSONTyped,
    PayPeriodToJSON,
} from './PayPeriod';

/**
 * The Employment object is used to represent a job position at a company.
 * NOTE: When there is a change in pay or title, integrations with historical 
 * data will create new Employment objects while integrations without 
 * historical data will update existing ones.
 * @export
 * @interface EmployeeJobPositionAtCompany
 */
export interface EmployeeJobPositionAtCompany {
    /**
     * 
     * @type {string}
     * @memberof EmployeeJobPositionAtCompany
     */
    id?: string;
    /**
     * The third-party API ID of the matching object.
     * @type {string}
     * @memberof EmployeeJobPositionAtCompany
     */
    remoteId?: string;
    /**
     * 
     * @type {string}
     * @memberof EmployeeJobPositionAtCompany
     */
    jobTitle?: string;
    /**
     * 
     * @type {number}
     * @memberof EmployeeJobPositionAtCompany
     */
    payRate?: number;
    /**
     * 
     * @type {PayPeriod}
     * @memberof EmployeeJobPositionAtCompany
     */
    payPeriod?: PayPeriod;
    /**
     * 
     * @type {PayFrequency}
     * @memberof EmployeeJobPositionAtCompany
     */
    payFrequency?: PayFrequency;
    /**
     * 
     * @type {string}
     * @memberof EmployeeJobPositionAtCompany
     */
    payCurrency?: string;
    /**
     * 
     * @type {FlsaStatus}
     * @memberof EmployeeJobPositionAtCompany
     */
    flsaStatus?: FlsaStatus;
    /**
     * 
     * @type {Date}
     * @memberof EmployeeJobPositionAtCompany
     */
    effectiveDate?: Date;
    /**
     * Indicates whether or not this object has been deleted in the third party platform.
     * @type {boolean}
     * @memberof EmployeeJobPositionAtCompany
     */
    remoteWasDeleted?: boolean;
    /**
     * 
     * @type {Date}
     * @memberof EmployeeJobPositionAtCompany
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof EmployeeJobPositionAtCompany
     */
    modifiedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof EmployeeJobPositionAtCompany
     */
    mergeAccountId?: string;
}

/**
 * Check if a given object implements the EmployeeJobPositionAtCompany interface.
 */
export function instanceOfEmployeeJobPositionAtCompany(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EmployeeJobPositionAtCompanyFromJSON(json: any): EmployeeJobPositionAtCompany {
    return EmployeeJobPositionAtCompanyFromJSONTyped(json, false);
}

export function EmployeeJobPositionAtCompanyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmployeeJobPositionAtCompany {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'remoteId': !exists(json, 'remoteId') ? undefined : json['remoteId'],
        'jobTitle': !exists(json, 'jobTitle') ? undefined : json['jobTitle'],
        'payRate': !exists(json, 'payRate') ? undefined : json['payRate'],
        'payPeriod': !exists(json, 'payPeriod') ? undefined : PayPeriodFromJSON(json['payPeriod']),
        'payFrequency': !exists(json, 'payFrequency') ? undefined : PayFrequencyFromJSON(json['payFrequency']),
        'payCurrency': !exists(json, 'payCurrency') ? undefined : json['payCurrency'],
        'flsaStatus': !exists(json, 'flsaStatus') ? undefined : FlsaStatusFromJSON(json['flsaStatus']),
        'effectiveDate': !exists(json, 'effectiveDate') ? undefined : (new Date(json['effectiveDate'])),
        'remoteWasDeleted': !exists(json, 'remoteWasDeleted') ? undefined : json['remoteWasDeleted'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
        'mergeAccountId': !exists(json, 'mergeAccountId') ? undefined : json['mergeAccountId'],
    };
}

export function EmployeeJobPositionAtCompanyToJSON(value?: EmployeeJobPositionAtCompany | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'remoteId': value.remoteId,
        'jobTitle': value.jobTitle,
        'payRate': value.payRate,
        'payPeriod': PayPeriodToJSON(value.payPeriod),
        'payFrequency': PayFrequencyToJSON(value.payFrequency),
        'payCurrency': value.payCurrency,
        'flsaStatus': FlsaStatusToJSON(value.flsaStatus),
        'effectiveDate': value.effectiveDate === undefined ? undefined : (value.effectiveDate.toISOString()),
        'remoteWasDeleted': value.remoteWasDeleted,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
        'mergeAccountId': value.mergeAccountId,
    };
}
