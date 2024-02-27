# IncomeStatement
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **name** | **String** | The income statement&#39;s name. | [optional] [default to null] |
| **currency** | **String** | The income statement&#39;s currency. | [optional] [default to null] |
| **company** | **String** | The company the income statement belongs to. | [optional] [default to null] |
| **startPeriod** | **Date** | The income statement&#39;s start period. | [optional] [default to null] |
| **endPeriod** | **Date** | The income statement&#39;s end period. | [optional] [default to null] |
| **income** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **costOfSales** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **grossProfit** | **Double** | The revenue minus the cost of sale. | [optional] [default to null] |
| **operatingExpenses** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **netOperatingIncome** | **Double** | The revenue minus the operating expenses. | [optional] [default to null] |
| **nonOperatingExpenses** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **netIncome** | **Double** | The gross profit minus the total expenses. | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **modifiedAt** | **Date** |  | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |
| **createdAt** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

