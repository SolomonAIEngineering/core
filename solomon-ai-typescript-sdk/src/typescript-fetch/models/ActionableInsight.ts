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
 * 
 * @export
 * @interface ActionableInsight
 */
export interface ActionableInsight {
    /**
     * 
     * @type {string}
     * @memberof ActionableInsight
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof ActionableInsight
     */
    detailedAction?: string;
    /**
     * 
     * @type {string}
     * @memberof ActionableInsight
     */
    summarizedAction?: string;
    /**
     * 
     * @type {Date}
     * @memberof ActionableInsight
     */
    generatedTime?: Date;
    /**
     * 
     * @type {Array<string>}
     * @memberof ActionableInsight
     */
    tags?: Array<string>;
}

/**
 * Check if a given object implements the ActionableInsight interface.
 */
export function instanceOfActionableInsight(value: object): boolean {
    return true;
}

export function ActionableInsightFromJSON(json: any): ActionableInsight {
    return ActionableInsightFromJSONTyped(json, false);
}

export function ActionableInsightFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActionableInsight {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'detailedAction': json['detailedAction'] == null ? undefined : json['detailedAction'],
        'summarizedAction': json['summarizedAction'] == null ? undefined : json['summarizedAction'],
        'generatedTime': json['generatedTime'] == null ? undefined : (new Date(json['generatedTime'])),
        'tags': json['tags'] == null ? undefined : json['tags'],
    };
}

export function ActionableInsightToJSON(value?: ActionableInsight | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'detailedAction': value['detailedAction'],
        'summarizedAction': value['summarizedAction'],
        'generatedTime': value['generatedTime'] == null ? undefined : ((value['generatedTime']).toISOString()),
        'tags': value['tags'],
    };
}

