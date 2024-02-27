# PlaidAccountTransaction
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **accountId** | **String** | The bank account ID associated with the transaction. | [optional] [default to null] |
| **amount** | **Double** | The amount of the transaction.  @gotag: ch:\&quot;amount\&quot; | [optional] [default to null] |
| **isoCurrencyCode** | **String** | The currency code of the transaction.  @gotag: ch:\&quot;iso_currency_code\&quot; | [optional] [default to null] |
| **unofficialCurrencyCode** | **String** | The unofficial currency code of the transaction.  @gotag: ch:\&quot;unofficial_currency_code\&quot; | [optional] [default to null] |
| **transactionId** | **String** | The transaction ID of interest.  @gotag: ch:\&quot;transaction_id\&quot; | [optional] [default to null] |
| **transactionCode** | **String** | The transaction code.  @gotag: ch:\&quot;transaction_code\&quot; | [optional] [default to null] |
| **currentDate** | **Date** | The date of the transaction.  @gotag: ch:\&quot;date\&quot; | [optional] [default to null] |
| **currentDatetime** | **Date** | The current datetime of the transaction.  @gotag: ch:\&quot;datetime\&quot; | [optional] [default to null] |
| **authorizedDate** | **Date** | The time at which the transaction was authorized.  @gotag: ch:\&quot;authorized_date\&quot; | [optional] [default to null] |
| **authorizedDatetime** | **Date** | The date-time when the transaction was authorized.  @gotag: ch:\&quot;authorized_datetime\&quot; | [optional] [default to null] |
| **categoryId** | **String** | The category ID of the transaction.  @gotag: ch:\&quot;category_id\&quot; | [optional] [default to null] |
| **categories** | **List** | The set of categories that the transaction belongs to. | [optional] [default to null] |
| **personalFinanceCategoryPrimary** | **String** | The primary personal finance category of the transaction.  @gotag: ch:\&quot;personal_finance_category_primary\&quot; | [optional] [default to null] |
| **personalFinanceCategoryDetailed** | **String** | The detailed personal finance category of the transaction.  @gotag: ch:\&quot;personal_finance_category_detailed\&quot; | [optional] [default to null] |
| **transactionName** | **String** | The name of the transaction.  @gotag: ch:\&quot;name\&quot; | [optional] [default to null] |
| **merchantName** | **String** | The merchant name of the transaction.  @gotag: ch:\&quot;merchant_name\&quot; | [optional] [default to null] |
| **checkNumber** | **String** | The check number associated with the transaction.  @gotag: ch:\&quot;check_number\&quot; | [optional] [default to null] |
| **paymentChannel** | **String** | The payment channel for the transaction.  @gotag: ch:\&quot;payment_channel\&quot; | [optional] [default to null] |
| **pending** | **Boolean** | Indicates whether the transaction is pending.  @gotag: ch:\&quot;pending\&quot; | [optional] [default to null] |
| **pendingTransactionId** | **String** | The ID of the pending transaction, if applicable.  @gotag: ch:\&quot;pending_transaction_id\&quot; | [optional] [default to null] |
| **accountOwner** | **String** | The account owner associated with the transaction.  @gotag: ch:\&quot;account_owner\&quot; | [optional] [default to null] |
| **paymentMetaByOrderOf** | **String** | Information about the entity to whom the payment is made (if available). | [optional] [default to null] |
| **paymentMetaPayee** | **String** | Information about the payee (if available). | [optional] [default to null] |
| **paymentMetaPayer** | **String** | Information about the payer (if available). | [optional] [default to null] |
| **paymentMetaPaymentMethod** | **String** | The payment method used for the transaction (if available). | [optional] [default to null] |
| **paymentMetaPaymentProcessor** | **String** | The payment processor involved in the transaction (if available). | [optional] [default to null] |
| **paymentMetaPpdId** | **String** | The Prearranged Payment and Deposit (PPD) ID (if available). | [optional] [default to null] |
| **paymentMetaReason** | **String** | The reason for the payment (if available). | [optional] [default to null] |
| **paymentMetaReferenceNumber** | **String** | The reference number associated with the payment (if available). | [optional] [default to null] |
| **locationAddress** | **String** | The street address of the transaction location (if available). | [optional] [default to null] |
| **locationCity** | **String** | The city of the transaction location (if available). | [optional] [default to null] |
| **locationRegion** | **String** | The region or state of the transaction location (if available). | [optional] [default to null] |
| **locationPostalCode** | **String** | The postal code of the transaction location (if available). | [optional] [default to null] |
| **locationCountry** | **String** | The country of the transaction location (if available). | [optional] [default to null] |
| **locationLat** | **Double** | The latitude of the transaction location (if available). | [optional] [default to null] |
| **locationLon** | **Double** | The longitude of the transaction location (if available). | [optional] [default to null] |
| **locationStoreNumber** | **String** | The store number associated with the transaction location (if available). | [optional] [default to null] |
| **time** | **Date** | The timestamp associated with the transaction. | [optional] [default to null] |
| **additionalProperties** | [**Any1**](Any1.md) |  | [optional] [default to null] |
| **id** | **String** | The unique ID for this transaction. | [optional] [default to null] |
| **userId** | **String** | The user ID associated with this transaction. | [optional] [default to null] |
| **linkId** | **String** | The link ID associated with this transaction. | [optional] [default to null] |
| **needsReview** | **Boolean** | Indicates whether this transaction needs review. | [optional] [default to null] |
| **hideTransaction** | **Boolean** | Indicates whether this transaction should be hidden. | [optional] [default to null] |
| **tags** | **List** | Tags associated with this transaction. | [optional] [default to null] |
| **notes** | [**List**](SmartNote.md) | Notes associated with this transaction. | [optional] [default to null] |
| **splits** | [**List**](TransactionSplit.md) | The number of splits associated with this transaction. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

