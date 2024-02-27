# PurchaseOrder
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |
| **status** | **String** |  | [optional] [default to null] |
| **issueDate** | **Date** | The purchase order&#39;s issue date. | [optional] [default to null] |
| **purchaseOrderNumber** | **String** | The human-readable number of the purchase order. | [optional] [default to null] |
| **deliveryDate** | **Date** | The purchase order&#39;s delivery date. | [optional] [default to null] |
| **deliveryAddress** | [**CompanyAddress**](CompanyAddress.md) |  | [optional] [default to null] |
| **customer** | **String** | The contact making the purchase order. | [optional] [default to null] |
| **vendor** | **String** | The party fulfilling the purchase order. | [optional] [default to null] |
| **memo** | **String** | A memo attached to the purchase order. | [optional] [default to null] |
| **company** | **String** | The company the purchase order belongs to. | [optional] [default to null] |
| **totalAmount** | **Float** | The purchase order&#39;s total amount.  Might be better as double. | [optional] [default to null] |
| **currency** | **String** | The purchase order&#39;s currency. | [optional] [default to null] |
| **exchangeRate** | **String** | Assuming string, but might be better as float or double. | [optional] [default to null] |
| **lineItems** | [**List**](PurchaseOrderLineItem.md) |  | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **remoteCreatedAt** | **Date** |  | [optional] [default to null] |
| **remoteUpdatedAt** | **Date** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |
| **accountingPeriod** | **String** | The accounting period that the PurchaseOrder was generated in. | [optional] [default to null] |
| **remoteId** | **String** |  | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

