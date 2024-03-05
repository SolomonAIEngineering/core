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
 * @brief Represents an account in the context of the service.
 * @export
 * @interface Account
 */
export interface Account {
    /**
     * Unique identifier for the account.
     * @type {string}
     * @memberof Account
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Account
     */
    auth0UserId?: string;
    /**
     * 
     * @type {Array<Workspace>}
     * @memberof Account
     */
    workspace?: Array<Workspace>;
}

/**
 * Check if a given object implements the Account interface.
 */
export function instanceOfAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AccountFromJSON(json: any): Account {
    return AccountFromJSONTyped(json, false);
}

export function AccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): Account {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'auth0UserId': !exists(json, 'auth0UserId') ? undefined : json['auth0UserId'],
        'workspace': !exists(json, 'workspace') ? undefined : ((json['workspace'] as Array<any>).map(WorkspaceFromJSON)),
    };
}

export function AccountToJSON(value?: Account | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'auth0UserId': value.auth0UserId,
        'workspace': value.workspace === undefined ? undefined : ((value.workspace as Array<any>).map(WorkspaceToJSON)),
    };
}
