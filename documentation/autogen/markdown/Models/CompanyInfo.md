# CompanyInfo
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **name** | **String** | The company&#39;s name. | [optional] [default to null] |
| **legalName** | **String** | The company&#39;s legal name. | [optional] [default to null] |
| **taxNumber** | **String** | The company&#39;s tax number. | [optional] [default to null] |
| **fiscalYearEndMonth** | **Integer** | The company&#39;s fiscal year end month. | [optional] [default to null] |
| **fiscalYearEndDay** | **Integer** | The company&#39;s fiscal year end day. | [optional] [default to null] |
| **currency** | **String** | The currency set in the company&#39;s accounting platform. | [optional] [default to null] |
| **remoteCreatedAt** | **Date** | When the third party&#39;s company was created.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **urls** | **List** | The company&#39;s urls. | [optional] [default to null] |
| **addresses** | [**List**](CompanyAddress.md) |  | [optional] [default to null] |
| **phoneNumbers** | **List** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

