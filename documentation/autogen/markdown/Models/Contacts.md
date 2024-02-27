# Contacts
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **remoteId** | **String** | The third-party API ID of the matching object. | [optional] [default to null] |
| **name** | **String** | The contact&#39;s name. | [optional] [default to null] |
| **isSupplier** | **Boolean** | Whether the contact is a supplier. | [optional] [default to null] |
| **isCustomer** | **Boolean** | Whether the contact is a customer. | [optional] [default to null] |
| **emailAddress** | **String** | The contact&#39;s email address. | [optional] [default to null] |
| **taxNumber** | **String** | The contact&#39;s tax number. | [optional] [default to null] |
| **status** | **String** |  | [optional] [default to null] |
| **currency** | **String** | The currency the contact&#39;s transactions are in. | [optional] [default to null] |
| **remoteUpdatedAt** | **Date** | When the third party&#39;s contact was updated.  Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **company** | **String** | The company the contact belongs to. | [optional] [default to null] |
| **addressesIds** | **List** | Address object IDs for the given Contacts object.  These are IDs, not the Address structure itself | [optional] [default to null] |
| **phoneNumbers** | **List** |  | [optional] [default to null] |
| **remoteWasDeleted** | **Boolean** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [default to null] |
| **modifiedAt** | **Date** | Consider using google.protobuf.Timestamp | [optional] [default to null] |
| **mergeRecordId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

