# PlaidInitiateTokenExchangeRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **userId** | **String** |  | [default to null] |
| **fullName** | **String** | The user&#39;s full legal name. This is an optional field used in  the [returning user experience](https://plaid.com/docs/link/returning-user) to associate Items to the user. | [default to null] |
| **email** | **String** | The user&#39;s email address. This field is optional, but required to enable the  [pre-authenticated returning user flow](https://plaid.com/docs/link/returning-user/#enabling-the-returning-user-experience). | [default to null] |
| **phoneNumber** | **String** | The user&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. This field is optional, but required to enable the [returning user experience](https://plaid.com/docs/link/returning-user). | [default to null] |
| **profileType** | [**FinancialUserProfileType**](FinancialUserProfileType.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

