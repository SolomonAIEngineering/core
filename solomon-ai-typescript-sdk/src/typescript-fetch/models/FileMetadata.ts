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
 * @interface FileMetadata
 */
export interface FileMetadata {
    /**
     * 
     * @type {string}
     * @memberof FileMetadata
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof FileMetadata
     */
    name?: string;
    /**
     * Timestamp when the file was created.
     * @type {Date}
     * @memberof FileMetadata
     */
    createdAt?: Date;
    /**
     * Timestamp when the file was last updated.
     * @type {Date}
     * @memberof FileMetadata
     */
    updatedAt?: Date;
    /**
     * Size of the file in bytes.
     * @type {string}
     * @memberof FileMetadata
     */
    size?: string;
    /**
     * Type of the file (e.g., 'text', 'image', 'video').
     * @type {string}
     * @memberof FileMetadata
     */
    fileType?: string;
    /**
     * Tags associated with the file.
     * @type {Array<string>}
     * @memberof FileMetadata
     */
    tags?: Array<string>;
    /**
     * Flag indicating if the file is marked as deleted.
     * @type {boolean}
     * @memberof FileMetadata
     */
    isDeleted?: boolean;
    /**
     * Version of the file metadata format.
     * @type {number}
     * @memberof FileMetadata
     */
    version?: number;
    /**
     * s3 key path
     * S3 key (path within the S3 bucket) for the file.
     * @type {string}
     * @memberof FileMetadata
     */
    s3Key?: string;
    /**
     * Name of the S3 bucket where the file is stored.
     * @type {string}
     * @memberof FileMetadata
     */
    s3BucketName?: string;
    /**
     * AWS region where the S3 bucket is located.
     * @type {string}
     * @memberof FileMetadata
     */
    s3Region?: string;
    /**
     * Version ID of the file, used when versioning is enabled in the S3 bucket.
     * @type {string}
     * @memberof FileMetadata
     */
    s3VersionId?: string;
    /**
     * Entity tag (ETag) of the file, a hash of the file used for change detection.
     * @type {string}
     * @memberof FileMetadata
     */
    s3Etag?: string;
    /**
     * MIME type of the file.
     * @type {string}
     * @memberof FileMetadata
     */
    s3ContentType?: string;
    /**
     * Size of the file in bytes.
     * @type {string}
     * @memberof FileMetadata
     */
    s3ContentLength?: string;
    /**
     * Encoding format used on the file, if any (e.g., gzip).
     * @type {string}
     * @memberof FileMetadata
     */
    s3ContentEncoding?: string;
    /**
     * How the file is to be presented in a web browser (attachment, inline).
     * @type {string}
     * @memberof FileMetadata
     */
    s3ContentDisposition?: string;
    /**
     * The date and time when the file was last modified in S3.
     * @type {Date}
     * @memberof FileMetadata
     */
    s3LastModified?: Date;
    /**
     * S3 storage class of the file (e.g., STANDARD, INTELLIGENT_TIERING, GLACIER).
     * @type {string}
     * @memberof FileMetadata
     */
    s3StorageClass?: string;
    /**
     * Details of server-side encryption used on the file, if any (e.g., AES256, aws:kms).
     * @type {string}
     * @memberof FileMetadata
     */
    s3ServerSideEncryption?: string;
    /**
     * Access control list (ACL) permissions for the file in S3.
     * @type {string}
     * @memberof FileMetadata
     */
    s3Acl?: string;
    /**
     * Custom metadata added to the file in S3 as key-value pairs.
     * @type {{ [key: string]: string; }}
     * @memberof FileMetadata
     */
    s3Metadata?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof FileMetadata
     */
    versionId?: string;
    /**
     * 
     * @type {string}
     * @memberof FileMetadata
     */
    uploadId?: string;
    /**
     * 
     * @type {string}
     * @memberof FileMetadata
     */
    location?: string;
}

/**
 * Check if a given object implements the FileMetadata interface.
 */
export function instanceOfFileMetadata(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FileMetadataFromJSON(json: any): FileMetadata {
    return FileMetadataFromJSONTyped(json, false);
}

export function FileMetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): FileMetadata {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'updatedAt': !exists(json, 'updatedAt') ? undefined : (new Date(json['updatedAt'])),
        'size': !exists(json, 'size') ? undefined : json['size'],
        'fileType': !exists(json, 'fileType') ? undefined : json['fileType'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'isDeleted': !exists(json, 'isDeleted') ? undefined : json['isDeleted'],
        'version': !exists(json, 'version') ? undefined : json['version'],
        's3Key': !exists(json, 's3Key') ? undefined : json['s3Key'],
        's3BucketName': !exists(json, 's3BucketName') ? undefined : json['s3BucketName'],
        's3Region': !exists(json, 's3Region') ? undefined : json['s3Region'],
        's3VersionId': !exists(json, 's3VersionId') ? undefined : json['s3VersionId'],
        's3Etag': !exists(json, 's3Etag') ? undefined : json['s3Etag'],
        's3ContentType': !exists(json, 's3ContentType') ? undefined : json['s3ContentType'],
        's3ContentLength': !exists(json, 's3ContentLength') ? undefined : json['s3ContentLength'],
        's3ContentEncoding': !exists(json, 's3ContentEncoding') ? undefined : json['s3ContentEncoding'],
        's3ContentDisposition': !exists(json, 's3ContentDisposition') ? undefined : json['s3ContentDisposition'],
        's3LastModified': !exists(json, 's3LastModified') ? undefined : (new Date(json['s3LastModified'])),
        's3StorageClass': !exists(json, 's3StorageClass') ? undefined : json['s3StorageClass'],
        's3ServerSideEncryption': !exists(json, 's3ServerSideEncryption') ? undefined : json['s3ServerSideEncryption'],
        's3Acl': !exists(json, 's3Acl') ? undefined : json['s3Acl'],
        's3Metadata': !exists(json, 's3Metadata') ? undefined : json['s3Metadata'],
        'versionId': !exists(json, 'versionId') ? undefined : json['versionId'],
        'uploadId': !exists(json, 'uploadId') ? undefined : json['uploadId'],
        'location': !exists(json, 'location') ? undefined : json['location'],
    };
}

export function FileMetadataToJSON(value?: FileMetadata | null): any {
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
        'size': value.size,
        'fileType': value.fileType,
        'tags': value.tags,
        'isDeleted': value.isDeleted,
        'version': value.version,
        's3Key': value.s3Key,
        's3BucketName': value.s3BucketName,
        's3Region': value.s3Region,
        's3VersionId': value.s3VersionId,
        's3Etag': value.s3Etag,
        's3ContentType': value.s3ContentType,
        's3ContentLength': value.s3ContentLength,
        's3ContentEncoding': value.s3ContentEncoding,
        's3ContentDisposition': value.s3ContentDisposition,
        's3LastModified': value.s3LastModified === undefined ? undefined : (value.s3LastModified.toISOString()),
        's3StorageClass': value.s3StorageClass,
        's3ServerSideEncryption': value.s3ServerSideEncryption,
        's3Acl': value.s3Acl,
        's3Metadata': value.s3Metadata,
        'versionId': value.versionId,
        'uploadId': value.uploadId,
        'location': value.location,
    };
}

