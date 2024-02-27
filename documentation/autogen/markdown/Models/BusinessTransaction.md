# BusinessTransaction
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **transactionType** | **String** | The type of transaction, which can by any transaction object not already included in Merge’s common model. | [optional] [default to null] |
| **number** | **String** | The transaction&#39;s number used for identifying purposes. | [optional] [default to null] |
| **transactionDate** | **Date** | The date upon which the transaction occurred. | [optional] [default to null] |
| **account** | **String** | The transaction&#39;s account. | [optional] [default to null] |
| **contact** | **String** | The contact to whom the transaction relates to. | [optional] [default to null] |
| **totalAmount** | **String** | The total amount being paid after taxes.  Might be better as double. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **exchangeRate** | **String** | The transaction&#39;s exchange rate.  Assuming string, but might be better as float or double. | [optional] [default to null] |
| **company** | **String** | The company the transaction belongs to. | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **lineItems** | [**List**](TransactionLineItem.md) |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted in the third party platform. | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the Transaction was generated in. | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** |  | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

