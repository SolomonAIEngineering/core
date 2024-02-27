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
import type { PersonalActionableInsightName } from './PersonalActionableInsightName';
import {
    PersonalActionableInsightNameFromJSON,
    PersonalActionableInsightNameFromJSONTyped,
    PersonalActionableInsightNameToJSON,
} from './PersonalActionableInsightName';

/**
 * 
 * @export
 * @interface PersonalActionableInsight
 */
export interface PersonalActionableInsight {
    /**
     * 
     * @type {string}
     * @memberof PersonalActionableInsight
     */
    id?: string;
    /**
     * 
     * @type {PersonalActionableInsightName}
     * @memberof PersonalActionableInsight
     */
    insightName?: PersonalActionableInsightName;
    /**
     * 
     * @type {string}
     * @memberof PersonalActionableInsight
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof PersonalActionableInsight
     */
    takeaway?: string;
    /**
     * 
     * @type {string}
     * @memberof PersonalActionableInsight
     */
    action?: string;
    /**
     * 
     * @type {string}
     * @memberof PersonalActionableInsight
     */
    expectedBenefit?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof PersonalActionableInsight
     */
    tags?: Array<string>;
    /**
     * 
     * @type {Date}
     * @memberof PersonalActionableInsight
     */
    generatedTime?: Date;
    /**
     * 
     * @type {Array<string>}
     * @memberof PersonalActionableInsight
     */
    metricsToOptimizeFor?: Array<string>;
}

/**
 * Check if a given object implements the PersonalActionableInsight interface.
 */
export function instanceOfPersonalActionableInsight(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PersonalActionableInsightFromJSON(json: any): PersonalActionableInsight {
    return PersonalActionableInsightFromJSONTyped(json, false);
}

export function PersonalActionableInsightFromJSONTyped(json: any, ignoreDiscriminator: boolean): PersonalActionableInsight {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'insightName': !exists(json, 'insightName') ? undefined : PersonalActionableInsightNameFromJSON(json['insightName']),
        'description': !exists(json, 'description') ? undefined : json['description'],
        'takeaway': !exists(json, 'takeaway') ? undefined : json['takeaway'],
        'action': !exists(json, 'action') ? undefined : json['action'],
        'expectedBenefit': !exists(json, 'expectedBenefit') ? undefined : json['expectedBenefit'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'generatedTime': !exists(json, 'generatedTime') ? undefined : (new Date(json['generatedTime'])),
        'metricsToOptimizeFor': !exists(json, 'metricsToOptimizeFor') ? undefined : json['metricsToOptimizeFor'],
    };
}

export function PersonalActionableInsightToJSON(value?: PersonalActionableInsight | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'insightName': PersonalActionableInsightNameToJSON(value.insightName),
        'description': value.description,
        'takeaway': value.takeaway,
        'action': value.action,
        'expectedBenefit': value.expectedBenefit,
        'tags': value.tags,
        'generatedTime': value.generatedTime === undefined ? undefined : (value.generatedTime.toISOString()),
        'metricsToOptimizeFor': value.metricsToOptimizeFor,
    };
}

