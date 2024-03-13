# FolderMetadata
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **childFolder** | [**List**](FolderMetadata.md) |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |
| **updatedAt** | **Date** |  | [optional] [default to null] |
| **files** | [**List**](FileMetadata.md) |  | [optional] [default to null] |
| **isDeleted** | **Boolean** |  | [optional] [default to null] |
| **s3BucketName** | **String** | The S3 bucket name where the folder is located. | [optional] [default to null] |
| **s3FolderPath** | **String** | The prefix path representing the folder in the S3 bucket. | [optional] [default to null] |
| **s3Region** | **String** | AWS region where the S3 bucket containing the folder is located. | [optional] [default to null] |
| **s3Metadata** | **Map** | Custom metadata for the folder, represented as key-value pairs. | [optional] [default to null] |
| **s3Acl** | **String** | Access control list (ACL) permissions for the folder in S3. | [optional] [default to null] |
| **s3LastModified** | **Date** | The date and time when the folder was last modified in S3. This might represent the last time a file was added, removed, or changed in the folder. | [optional] [default to null] |
| **versionId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

