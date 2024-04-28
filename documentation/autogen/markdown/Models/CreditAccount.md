# CreditAccount
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** |  | [optional] [default to null] |
| **userId** | **String** |  | [optional] [default to null] |
| **name** | **String** |  | [optional] [default to null] |
| **number** | **String** |  | [optional] [default to null] |
| **type** | [**BankAccountType**](BankAccountType.md) |  | [optional] [default to null] |
| **balance** | **Float** |  | [optional] [default to null] |
| **currentFunds** | **Double** |  | [optional] [default to null] |
| **balanceLimit** | **String** |  | [optional] [default to null] |
| **plaidAccountId** | **String** |  | [optional] [default to null] |
| **subtype** | **String** |  | [optional] [default to null] |
| **isOverdue** | **Boolean** |  | [optional] [default to null] |
| **lastPaymentAmount** | **Double** |  | [optional] [default to null] |
| **lastPaymentDate** | **String** |  | [optional] [default to null] |
| **lastStatementIssueDate** | **String** |  | [optional] [default to null] |
| **minimumAmountDueDate** | **Double** |  | [optional] [default to null] |
| **nextPaymentDate** | **String** |  | [optional] [default to null] |
| **aprs** | [**List**](Apr.md) |  | [optional] [default to null] |
| **lastStatementBalance** | **Double** |  | [optional] [default to null] |
| **minimumPaymentAmount** | **Double** |  | [optional] [default to null] |
| **nextPaymentDueDate** | **String** |  | [optional] [default to null] |
| **status** | [**BankAccountStatus**](BankAccountStatus.md) |  | [optional] [default to null] |
| **transactions** | [**List**](PlaidAccountTransaction.md) |  | [optional] [default to null] |
| **recurringTransactions** | [**List**](PlaidAccountRecurringTransaction.md) |  | [optional] [default to null] |
| **pockets** | [**List**](Pocket.md) |  | [optional] [default to null] |
| **statements** | [**List**](AccountStatements.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

