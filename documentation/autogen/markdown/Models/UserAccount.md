# UserAccount
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | Unique identifier for the account. | [optional] [default to null] |
| **email** | **String** | Email associated with the user account. | [optional] [default to null] |
| **address** | [**Address**](Address.md) |  | [optional] [default to null] |
| **bio** | **String** | Brief description about the user, up to 200 characters. | [optional] [default to null] |
| **headline** | **String** | Short headline for the user&#39;s profile. | [optional] [default to null] |
| **phoneNumber** | **String** | Phone number associated with the account. | [optional] [default to null] |
| **tags** | [**List**](Tags.md) | Tags associated with the user account, between 1 and 10. | [optional] [default to null] |
| **authnAccountId** | **String** | ID for the authentication service linked to this account. | [optional] [default to null] |
| **isActive** | **Boolean** | Indicates if the account is currently active. | [optional] [default to null] |
| **firstname** | **String** | User&#39;s first name. | [optional] [default to null] |
| **lastname** | **String** | User&#39;s last name. | [optional] [default to null] |
| **username** | **String** | Username associated with the account, minimum of 10 characters. | [optional] [default to null] |
| **isPrivate** | **Boolean** | Indicates if the account is set to private. | [optional] [default to null] |
| **isEmailVerified** | **Boolean** | Indicates if the user&#39;s email has been verified. | [optional] [default to null] |
| **createdAt** | **Date** | Timestamp for when the account was created. | [optional] [default to null] |
| **verifiedAt** | **Date** | Timestamp for when the email was verified. | [optional] [default to null] |
| **settings** | [**Settings**](Settings.md) |  | [optional] [default to null] |
| **accountType** | [**ProfileType**](ProfileType.md) |  | [optional] [default to null] |
| **profileImageUrl** | **String** | Profile image associated with the user account. | [optional] [default to null] |
| **auth0UserId** | **String** |  | [optional] [default to null] |
| **role** | [**Role**](Role.md) |  | [optional] [default to null] |
| **algoliaUserId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

