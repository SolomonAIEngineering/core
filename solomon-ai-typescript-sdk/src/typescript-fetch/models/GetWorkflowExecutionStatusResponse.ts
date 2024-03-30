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
 * @interface GetWorkflowExecutionStatusResponse
 */
export interface GetWorkflowExecutionStatusResponse {
    /**
     * 
     * @type {string}
     * @memberof GetWorkflowExecutionStatusResponse
     */
    workflowId?: string;
    /**
     * 
     * @type {string}
     * @memberof GetWorkflowExecutionStatusResponse
     */
    status?: string;
    /**
     * 
     * @type {string}
     * @memberof GetWorkflowExecutionStatusResponse
     */
    runId?: string;
}

/**
 * Check if a given object implements the GetWorkflowExecutionStatusResponse interface.
 */
export function instanceOfGetWorkflowExecutionStatusResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetWorkflowExecutionStatusResponseFromJSON(json: any): GetWorkflowExecutionStatusResponse {
    return GetWorkflowExecutionStatusResponseFromJSONTyped(json, false);
}

export function GetWorkflowExecutionStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetWorkflowExecutionStatusResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'workflowId': !exists(json, 'workflowId') ? undefined : json['workflowId'],
        'status': !exists(json, 'status') ? undefined : json['status'],
        'runId': !exists(json, 'runId') ? undefined : json['runId'],
    };
}

export function GetWorkflowExecutionStatusResponseToJSON(value?: GetWorkflowExecutionStatusResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'workflowId': value.workflowId,
        'status': value.status,
        'runId': value.runId,
    };
}

