# BusinessChartOfAccount
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [default to null] |
| **remoteId** | **String** |  | [optional] [default to null] |
| **name** | **String** | The account&#39;s name. | [optional] [default to null] |
| **description** | **String** | The account&#39;s description. | [optional] [default to null] |
| **classification** | **String** | The account&#39;s broadest grouping. Possible values include: ASSET, EQUITY, EXPENSE,  LIABILITY, REVENUE. In cases where there is no clear mapping, the original  value passed through will be returned. | [optional] [default to null] |
| **type** | **String** | The account&#39;s type is a narrower and more specific grouping within the account&#39;s classification. | [optional] [default to null] |
| **status** | **String** | The account&#39;s status. Possible values include: ACTIVE, PENDING, INACTIVE. In cases where there is  no clear mapping, the original value passed through will be returned. | [optional] [default to null] |
| **currentBalance** | **Double** | The account&#39;s current balance. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **accountNumber** | **String** | The account&#39;s number. | [optional] [default to null] |
| **parentAccountId** | **String** | ID of the parent account. | [optional] [default to null] |
| **company** | **String** | The company the account belongs to. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

