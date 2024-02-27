# ExpenseLine
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **netAmount** | **Double** | The line&#39;s net amount. | [optional] [default to null] |
| **trackingCategory** | **String** |  | [optional] [default to null] |
| **trackingCategories** | **List** |  | [optional] [default to null] |
| **company** | **String** | The company the line belongs to. | [optional] [default to null] |
| **item** | **String** | The line&#39;s item.  This seems to be an ID | [optional] [default to null] |
| **account** | **String** | The expense&#39;s payment account. | [optional] [default to null] |
| **contact** | **String** | The expense&#39;s contact.  Optional based on provided JSON | [optional] [default to null] |
| **description** | **String** | The description of the item that was purchased by the company. | [optional] [default to null] |
| **exchangeRate** | **String** | The expense line item&#39;s exchange rate.  Consider using double or float if this represents a number | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **modifiedAt** | **Date** | Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

