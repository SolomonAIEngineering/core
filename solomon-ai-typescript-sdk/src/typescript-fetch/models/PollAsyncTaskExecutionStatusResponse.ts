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
import type { TaskState } from './TaskState';
import {
    TaskStateFromJSON,
    TaskStateFromJSONTyped,
    TaskStateToJSON,
} from './TaskState';

/**
 * 
 * @export
 * @interface PollAsyncTaskExecutionStatusResponse
 */
export interface PollAsyncTaskExecutionStatusResponse {
    /**
     * 
     * @type {string}
     * @memberof PollAsyncTaskExecutionStatusResponse
     */
    workflowId?: string;
    /**
     * 
     * @type {TaskState}
     * @memberof PollAsyncTaskExecutionStatusResponse
     */
    status?: TaskState;
    /**
     * 
     * @type {string}
     * @memberof PollAsyncTaskExecutionStatusResponse
     */
    runId?: string;
}

/**
 * Check if a given object implements the PollAsyncTaskExecutionStatusResponse interface.
 */
export function instanceOfPollAsyncTaskExecutionStatusResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PollAsyncTaskExecutionStatusResponseFromJSON(json: any): PollAsyncTaskExecutionStatusResponse {
    return PollAsyncTaskExecutionStatusResponseFromJSONTyped(json, false);
}

export function PollAsyncTaskExecutionStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PollAsyncTaskExecutionStatusResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'workflowId': !exists(json, 'workflowId') ? undefined : json['workflowId'],
        'status': !exists(json, 'status') ? undefined : TaskStateFromJSON(json['status']),
        'runId': !exists(json, 'runId') ? undefined : json['runId'],
    };
}

export function PollAsyncTaskExecutionStatusResponseToJSON(value?: PollAsyncTaskExecutionStatusResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'workflowId': value.workflowId,
        'status': TaskStateToJSON(value.status),
        'runId': value.runId,
    };
}

