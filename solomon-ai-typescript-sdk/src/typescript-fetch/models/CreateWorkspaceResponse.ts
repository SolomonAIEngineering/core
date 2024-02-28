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
import type { Workspace } from './Workspace';
import {
    WorkspaceFromJSON,
    WorkspaceFromJSONTyped,
    WorkspaceToJSON,
} from './Workspace';

/**
 * 
 * @export
 * @interface CreateWorkspaceResponse
 */
export interface CreateWorkspaceResponse {
    /**
     * 
     * @type {Workspace}
     * @memberof CreateWorkspaceResponse
     */
    workspace?: Workspace;
}

/**
 * Check if a given object implements the CreateWorkspaceResponse interface.
 */
export function instanceOfCreateWorkspaceResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CreateWorkspaceResponseFromJSON(json: any): CreateWorkspaceResponse {
    return CreateWorkspaceResponseFromJSONTyped(json, false);
}

export function CreateWorkspaceResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateWorkspaceResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'workspace': !exists(json, 'workspace') ? undefined : WorkspaceFromJSON(json['workspace']),
    };
}

export function CreateWorkspaceResponseToJSON(value?: CreateWorkspaceResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'workspace': WorkspaceToJSON(value.workspace),
    };
}

