# CreditNote
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **transactionDate** | **Date** | The credit note&#39;s transaction date.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **status** | **String** | The credit note&#39;s status. | [optional] [default to null] |
| **number** | **String** | The credit note&#39;s number. | [optional] [default to null] |
| **contact** | **String** | The credit note&#39;s contact. | [optional] [default to null] |
| **company** | **String** | The company the credit note belongs to. | [optional] [default to null] |
| **exchangeRate** | **String** | The credit note&#39;s exchange rate.  Consider using double or float if this represents a number | [optional] [default to null] |
| **totalAmount** | **Double** | The credit note&#39;s total amount. | [optional] [default to null] |
| **remainingCredit** | **Double** | The amount of value remaining in the credit note that the customer can use. | [optional] [default to null] |
| **lineItems** | [**List**](CreditNoteLineItem.md) |  | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **remoteCreatedAt** | **Date** |  | [optional] [default to null] |
| **remoteUpdatedAt** | **Date** |  | [optional] [default to null] |
| **paymentIds** | **List** | These are IDs | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the CreditNote was generated in. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

