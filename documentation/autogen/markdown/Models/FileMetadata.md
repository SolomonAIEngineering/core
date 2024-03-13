# FileMetadata
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **createdAt** | **Date** | Timestamp when the file was created. | [optional] [default to null] |
| **updatedAt** | **Date** | Timestamp when the file was last updated. | [optional] [default to null] |
| **size** | **String** | Size of the file in bytes. | [optional] [default to null] |
| **fileType** | **String** | Type of the file (e.g., &#39;text&#39;, &#39;image&#39;, &#39;video&#39;). | [optional] [default to null] |
| **tags** | **List** | Tags associated with the file. | [optional] [default to null] |
| **isDeleted** | **Boolean** | Flag indicating if the file is marked as deleted. | [optional] [default to null] |
| **version** | **Integer** | Version of the file metadata format. | [optional] [default to null] |
| **s3Key** | **String** | s3 key path S3 key (path within the S3 bucket) for the file. | [optional] [default to null] |
| **s3BucketName** | **String** | Name of the S3 bucket where the file is stored. | [optional] [default to null] |
| **s3Region** | **String** | AWS region where the S3 bucket is located. | [optional] [default to null] |
| **s3VersionId** | **String** | Version ID of the file, used when versioning is enabled in the S3 bucket. | [optional] [default to null] |
| **s3Etag** | **String** | Entity tag (ETag) of the file, a hash of the file used for change detection. | [optional] [default to null] |
| **s3ContentType** | **String** | MIME type of the file. | [optional] [default to null] |
| **s3ContentLength** | **String** | Size of the file in bytes. | [optional] [default to null] |
| **s3ContentEncoding** | **String** | Encoding format used on the file, if any (e.g., gzip). | [optional] [default to null] |
| **s3ContentDisposition** | **String** | How the file is to be presented in a web browser (attachment, inline). | [optional] [default to null] |
| **s3LastModified** | **Date** | The date and time when the file was last modified in S3. | [optional] [default to null] |
| **s3StorageClass** | **String** | S3 storage class of the file (e.g., STANDARD, INTELLIGENT_TIERING, GLACIER). | [optional] [default to null] |
| **s3ServerSideEncryption** | **String** | Details of server-side encryption used on the file, if any (e.g., AES256, aws:kms). | [optional] [default to null] |
| **s3Acl** | **String** | Access control list (ACL) permissions for the file in S3. | [optional] [default to null] |
| **s3Metadata** | **Map** | Custom metadata added to the file in S3 as key-value pairs. | [optional] [default to null] |
| **versionId** | **String** |  | [optional] [default to null] |
| **uploadId** | **String** |  | [optional] [default to null] |
| **location** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

