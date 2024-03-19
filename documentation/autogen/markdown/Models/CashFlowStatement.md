# CashFlowStatement
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **name** | **String** | The cash flow statement&#39;s name. | [optional] [default to null] |
| **currency** | **String** |  | [optional] [default to null] |
| **company** | **String** | The company the cash flow statement belongs to. | [optional] [default to null] |
| **startPeriod** | **Date** | The cash flow statement&#39;s start period.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **endPeriod** | **Date** | The cash flow statement&#39;s end period.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **cashAtBeginningOfPeriod** | **Double** | Cash and cash equivalents at the beginning of the cash flow statement&#39;s period. | [optional] [default to null] |
| **cashAtEndOfPeriod** | **Double** | Cash and cash equivalents at the beginning of the cash flow statement&#39;s period. | [optional] [default to null] |
| **operatingActivities** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **investingActivities** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **financingActivities** | [**List**](ReportItem.md) |  | [optional] [default to null] |
| **remoteGeneratedAt** | **Date** | The time that cash flow statement was generated by the accounting system.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **modifiedAt** | **Date** | Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
