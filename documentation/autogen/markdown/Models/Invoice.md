# Invoice
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **type** | **String** | Whether the invoice is an accounts receivable or accounts payable  If type is accounts_payable, the invoice is a bill. If type is  accounts_receivable, it is an invoice. Possible values include: ACCOUNTS_RECEIVABLE, ACCOUNTS_PAYABLE. | [optional] [default to null] |
| **contact** | **String** | The invoice&#39;s contact. | [optional] [default to null] |
| **number** | **String** | The invoice&#39;s number. | [optional] [default to null] |
| **issueDate** | **Date** | The invoice&#39;s issue date. | [optional] [default to null] |
| **dueDate** | **Date** | The invoice&#39;s due date. | [optional] [default to null] |
| **paidOnDate** | **Date** | The invoice&#39;s paid date. | [optional] [default to null] |
| **memo** | **String** | The invoice&#39;s private note. | [optional] [default to null] |
| **company** | **String** | The company the invoice belongs to. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **exchangeRate** | **String** | The invoice&#39;s exchange rate. | [optional] [default to null] |
| **totalDiscount** | **Float** | The total discounts applied to the total cost. | [optional] [default to null] |
| **subTotal** | **Float** | The total amount being paid before taxes. | [optional] [default to null] |
| **status** | **String** |  | [optional] [default to null] |
| **totalTaxAmount** | **Float** | The total amount being paid in taxes. | [optional] [default to null] |
| **totalAmount** | **Float** | The invoice&#39;s total amount. | [optional] [default to null] |
| **balance** | **Float** | The invoice&#39;s remaining balance. | [optional] [default to null] |
| **remoteUpdatedAt** | **Date** | When the third party&#39;s invoice entry was updated. | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **payments** | **List** | Array of Payment object IDs. | [optional] [default to null] |
| **lineItems** | [**List**](InvoiceLineItem.md) |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the Invoice was generated in. | [optional] [default to null] |
| **purchaseOrders** | **List** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

