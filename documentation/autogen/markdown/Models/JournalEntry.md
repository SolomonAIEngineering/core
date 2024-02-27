# JournalEntry
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** |  | [optional] [default to null] |
| **transactionDate** | **Date** | The journal entry&#39;s transaction date. | [optional] [default to null] |
| **remoteCreatedAt** | **Date** | When the third party&#39;s journal entry was created. | [optional] [default to null] |
| **remoteUpdatedAt** | **Date** |  | [optional] [default to null] |
| **payments** | **List** |  | [optional] [default to null] |
| **memo** | **String** | The journal entry&#39;s private note. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **exchangeRate** | **String** | The journal entry&#39;s exchange rate.  Assuming string due to the example provided, but could be float or double. | [optional] [default to null] |
| **company** | **String** | The company the journal entry belongs to. | [optional] [default to null] |
| **lines** | [**List**](JournalLine.md) | The JournalLine object is used to represent a journal entry&#39;s line items. | [optional] [default to null] |
| **journalNumber** | **String** | Reference number for identifying journal entries. | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |
| **postingStatus** | **String** |  | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the JournalEntry was generated in. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |
| **appliedPayments** | **List** | A list of the Payment Applied to Lines common models  related to a given Invoice, Credit Note, or Journal Entry. | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

