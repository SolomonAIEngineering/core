# UpdateRole_request
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | **String** | Name of the role. | [optional] [default to null] |
| **type** | [**RoleType**](RoleType.md) |  | [optional] [default to null] |
| **canCreateUsers** | **Boolean** | Permissions related to user management. | [optional] [default to null] |
| **canReadUsers** | **Boolean** |  | [optional] [default to null] |
| **canUpdateUsers** | **Boolean** |  | [optional] [default to null] |
| **canDeleteUsers** | **Boolean** |  | [optional] [default to null] |
| **canCreateProjects** | **Boolean** | Permissions related to project management. | [optional] [default to null] |
| **canReadProjects** | **Boolean** |  | [optional] [default to null] |
| **canUpdateProjects** | **Boolean** |  | [optional] [default to null] |
| **canDeleteProjects** | **Boolean** |  | [optional] [default to null] |
| **canCreateReports** | **Boolean** | Permissions related to report management. | [optional] [default to null] |
| **canReadReports** | **Boolean** |  | [optional] [default to null] |
| **canUpdateReports** | **Boolean** |  | [optional] [default to null] |
| **canDeleteReports** | **Boolean** |  | [optional] [default to null] |
| **createdAt** | **Date** | Add more permissions as necessary for other modules or features. Timestamps for tracking creation and modification times. | [optional] [default to null] |
| **updatedAt** | **Date** |  | [optional] [default to null] |
| **auditLog** | [**List**](RoleAuditEvents.md) | Audit log for this role. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

