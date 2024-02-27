# Expense
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **transactionDate** | **Date** | When the transaction occurred.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **remoteCreatedAt** | **Date** | When the expense was created.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **account** | **String** | The expense&#39;s payment account. | [optional] [default to null] |
| **contact** | **String** | The expense&#39;s contact. | [optional] [default to null] |
| **totalAmount** | **Double** | The expense&#39;s total amount. | [optional] [default to null] |
| **subTotal** | **Double** | The expense&#39;s total amount before tax. | [optional] [default to null] |
| **totalTaxAmount** | **Double** | The expense&#39;s total tax amount. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **exchangeRate** | **String** | The expense&#39;s exchange rate.  Consider using double or float if this represents a number | [optional] [default to null] |
| **company** | **String** | The company the expense belongs to. | [optional] [default to null] |
| **memo** | **String** | The expense&#39;s private note. | [optional] [default to null] |
| **lines** | [**List**](ExpenseLine.md) | The ExpenseLine object is used to represent an expense&#39;s line items. | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the Expense was generated in. | [optional] [default to null] |
| **modifiedAt** | **Date** | Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

