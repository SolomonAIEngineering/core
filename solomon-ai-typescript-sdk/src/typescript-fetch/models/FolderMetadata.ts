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
import type { FileMetadata } from './FileMetadata';
import {
    FileMetadataFromJSON,
    FileMetadataFromJSONTyped,
    FileMetadataToJSON,
} from './FileMetadata';

/**
 * 
 * @export
 * @interface FolderMetadata
 */
export interface FolderMetadata {
    /**
     * 
     * @type {string}
     * @memberof FolderMetadata
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof FolderMetadata
     */
    name?: string;
    /**
     * 
     * @type {Array<FolderMetadata>}
     * @memberof FolderMetadata
     */
    childFolder?: Array<FolderMetadata>;
    /**
     * 
     * @type {Date}
     * @memberof FolderMetadata
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof FolderMetadata
     */
    updatedAt?: Date;
    /**
     * 
     * @type {Array<FileMetadata>}
     * @memberof FolderMetadata
     */
    files?: Array<FileMetadata>;
    /**
     * 
     * @type {boolean}
     * @memberof FolderMetadata
     */
    isDeleted?: boolean;
    /**
     * The S3 bucket name where the folder is located.
     * @type {string}
     * @memberof FolderMetadata
     */
    s3BucketName?: string;
    /**
     * The prefix path representing the folder in the S3 bucket.
     * @type {string}
     * @memberof FolderMetadata
     */
    s3FolderPath?: string;
    /**
     * AWS region where the S3 bucket containing the folder is located.
     * @type {string}
     * @memberof FolderMetadata
     */
    s3Region?: string;
    /**
     * Custom metadata for the folder, represented as key-value pairs.
     * @type {{ [key: string]: string; }}
     * @memberof FolderMetadata
     */
    s3Metadata?: { [key: string]: string; };
    /**
     * Access control list (ACL) permissions for the folder in S3.
     * @type {string}
     * @memberof FolderMetadata
     */
    s3Acl?: string;
    /**
     * The date and time when the folder was last modified in S3.
     * This might represent the last time a file was added, removed, or changed in the folder.
     * @type {Date}
     * @memberof FolderMetadata
     */
    s3LastModified?: Date;
    /**
     * 
     * @type {string}
     * @memberof FolderMetadata
     */
    versionId?: string;
}

/**
 * Check if a given object implements the FolderMetadata interface.
 */
export function instanceOfFolderMetadata(value: object): boolean {
    return true;
}

export function FolderMetadataFromJSON(json: any): FolderMetadata {
    return FolderMetadataFromJSONTyped(json, false);
}

export function FolderMetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): FolderMetadata {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'childFolder': json['childFolder'] == null ? undefined : ((json['childFolder'] as Array<any>).map(FolderMetadataFromJSON)),
        'createdAt': json['createdAt'] == null ? undefined : (new Date(json['createdAt'])),
        'updatedAt': json['updatedAt'] == null ? undefined : (new Date(json['updatedAt'])),
        'files': json['files'] == null ? undefined : ((json['files'] as Array<any>).map(FileMetadataFromJSON)),
        'isDeleted': json['isDeleted'] == null ? undefined : json['isDeleted'],
        's3BucketName': json['s3BucketName'] == null ? undefined : json['s3BucketName'],
        's3FolderPath': json['s3FolderPath'] == null ? undefined : json['s3FolderPath'],
        's3Region': json['s3Region'] == null ? undefined : json['s3Region'],
        's3Metadata': json['s3Metadata'] == null ? undefined : json['s3Metadata'],
        's3Acl': json['s3Acl'] == null ? undefined : json['s3Acl'],
        's3LastModified': json['s3LastModified'] == null ? undefined : (new Date(json['s3LastModified'])),
        'versionId': json['versionId'] == null ? undefined : json['versionId'],
    };
}

export function FolderMetadataToJSON(value?: FolderMetadata | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'name': value['name'],
        'childFolder': value['childFolder'] == null ? undefined : ((value['childFolder'] as Array<any>).map(FolderMetadataToJSON)),
        'createdAt': value['createdAt'] == null ? undefined : ((value['createdAt']).toISOString()),
        'updatedAt': value['updatedAt'] == null ? undefined : ((value['updatedAt']).toISOString()),
        'files': value['files'] == null ? undefined : ((value['files'] as Array<any>).map(FileMetadataToJSON)),
        'isDeleted': value['isDeleted'],
        's3BucketName': value['s3BucketName'],
        's3FolderPath': value['s3FolderPath'],
        's3Region': value['s3Region'],
        's3Metadata': value['s3Metadata'],
        's3Acl': value['s3Acl'],
        's3LastModified': value['s3LastModified'] == null ? undefined : ((value['s3LastModified']).toISOString()),
        'versionId': value['versionId'],
    };
}

