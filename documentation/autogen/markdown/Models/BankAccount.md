# BankAccount
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **userId** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **number** | **String** |  | [optional] [default to null] |
| **type** | [**BankAccountType**](BankAccountType.md) |  | [default to null] |
| **balance** | **Float** |  | [default to null] |
| **currency** | **String** |  | [default to null] |
| **currentFunds** | **Double** |  | [default to null] |
| **balanceLimit** | **String** |  | [optional] [default to null] |
| **pockets** | [**List**](Pocket.md) |  | [optional] [default to null] |
| **plaidAccountId** | **String** |  | [optional] [default to null] |
| **subtype** | **String** |  | [optional] [default to null] |
| **status** | [**BankAccountStatus**](BankAccountStatus.md) |  | [optional] [default to null] |
| **transactions** | [**List**](PlaidAccountTransaction.md) |  | [optional] [default to null] |
| **recurringTransactions** | [**List**](PlaidAccountRecurringTransaction.md) |  | [optional] [default to null] |
| **statements** | [**List**](AccountStatements.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

