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
import type { FolderMetadata } from './FolderMetadata';
import {
    FolderMetadataFromJSON,
    FolderMetadataFromJSONTyped,
    FolderMetadataToJSON,
} from './FolderMetadata';

/**
 * 
 * @export
 * @interface Workspace
 */
export interface Workspace {
    /**
     * 
     * @type {string}
     * @memberof Workspace
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Workspace
     */
    name?: string;
    /**
     * 
     * @type {Date}
     * @memberof Workspace
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Workspace
     */
    updatedAt?: Date;
    /**
     * 
     * @type {Array<string>}
     * @memberof Workspace
     */
    tags?: Array<string>;
    /**
     * 
     * @type {Array<FolderMetadata>}
     * @memberof Workspace
     */
    folders?: Array<FolderMetadata>;
    /**
     * 
     * @type {number}
     * @memberof Workspace
     */
    version?: number;
    /**
     * 
     * @type {boolean}
     * @memberof Workspace
     */
    isDeleted?: boolean;
    /**
     * The S3 bucket name where the folder is located.
     * @type {string}
     * @memberof Workspace
     */
    s3BucketName?: string;
    /**
     * The prefix path representing the folder in the S3 bucket.
     * @type {string}
     * @memberof Workspace
     */
    s3FolderPath?: string;
    /**
     * AWS region where the S3 bucket containing the folder is located.
     * @type {string}
     * @memberof Workspace
     */
    s3Region?: string;
    /**
     * Custom metadata for the folder, represented as key-value pairs.
     * @type {{ [key: string]: string; }}
     * @memberof Workspace
     */
    s3Metadata?: { [key: string]: string; };
    /**
     * Access control list (ACL) permissions for the folder in S3.
     * @type {string}
     * @memberof Workspace
     */
    s3Acl?: string;
    /**
     * The date and time when the folder was last modified in S3.
     * This might represent the last time a file was added, removed, or changed in the folder.
     * @type {Date}
     * @memberof Workspace
     */
    s3LastModified?: Date;
    /**
     * 
     * @type {string}
     * @memberof Workspace
     */
    uniqueIdentifier?: string;
    /**
     * 
     * @type {string}
     * @memberof Workspace
     */
    versionId?: string;
}

/**
 * Check if a given object implements the Workspace interface.
 */
export function instanceOfWorkspace(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function WorkspaceFromJSON(json: any): Workspace {
    return WorkspaceFromJSONTyped(json, false);
}

export function WorkspaceFromJSONTyped(json: any, ignoreDiscriminator: boolean): Workspace {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'folders': !exists(json, 'folders') ? undefined : ((json['folders'] as Array<any>).map(FolderMetadataFromJSON)),
        'version': !exists(json, 'version') ? undefined : json['version'],
        'isDeleted': !exists(json, 'isDeleted') ? undefined : json['isDeleted'],
        's3BucketName': !exists(json, 's3BucketName') ? undefined : json['s3BucketName'],
        's3FolderPath': !exists(json, 's3FolderPath') ? undefined : json['s3FolderPath'],
        's3Region': !exists(json, 's3Region') ? undefined : json['s3Region'],
        's3Metadata': !exists(json, 's3Metadata') ? undefined : json['s3Metadata'],
        's3Acl': !exists(json, 's3Acl') ? undefined : json['s3Acl'],
        's3LastModified': !exists(json, 's3LastModified') ? undefined : (new Date(json['s3LastModified'])),
        'uniqueIdentifier': !exists(json, 'uniqueIdentifier') ? undefined : json['uniqueIdentifier'],
        'versionId': !exists(json, 'versionId') ? undefined : json['versionId'],
    };
}

export function WorkspaceToJSON(value?: Workspace | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updatedAt': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
        'tags': value.tags,
        'folders': value.folders === undefined ? undefined : ((value.folders as Array<any>).map(FolderMetadataToJSON)),
        'version': value.version,
        'isDeleted': value.isDeleted,
        's3BucketName': value.s3BucketName,
        's3FolderPath': value.s3FolderPath,
        's3Region': value.s3Region,
        's3Metadata': value.s3Metadata,
        's3Acl': value.s3Acl,
        's3LastModified': value.s3LastModified === undefined ? undefined : (value.s3LastModified.toISOString()),
        'uniqueIdentifier': value.uniqueIdentifier,
        'versionId': value.versionId,
    };
}

