# BusinessAccount
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | Unique identifier for the business account. | [optional] [default to null] |
| **email** | **String** | Email associated with the business account. | [optional] [default to null] |
| **address** | [**Address**](Address.md) |  | [optional] [default to null] |
| **bio** | **String** | Short description of the business account. Maximum of 200 characters. | [optional] [default to null] |
| **headline** | **String** | Headline for the profile of the business account. | [optional] [default to null] |
| **phoneNumber** | **String** | Phone number associated with the business account. | [optional] [default to null] |
| **tags** | [**List**](Tags.md) | Tags associated with the business account. Between 1 and 10 tags are allowed. | [optional] [default to null] |
| **authnAccountId** | **String** | Identifier for the associated authentication service account. | [optional] [default to null] |
| **isActive** | **Boolean** | Indicates whether the business account is active. | [optional] [default to null] |
| **username** | **String** | Username for the business account. Must be at least 10 characters long. | [optional] [default to null] |
| **isPrivate** | **Boolean** | Indicates whether the business account is private. | [optional] [default to null] |
| **isEmailVerified** | **Boolean** | Indicates whether the email associated with the business account has been verified. | [optional] [default to null] |
| **createdAt** | **Date** | Timestamp indicating when the business account was created. | [optional] [default to null] |
| **verifiedAt** | **Date** | Timestamp indicating when the email for the business account was verified. | [optional] [default to null] |
| **companyEstablishedDate** | **String** | Date when the company associated with the business account was established. | [optional] [default to null] |
| **companyIndustryType** | **String** | Industry type of the company associated with the business account. | [optional] [default to null] |
| **companyWebsiteUrl** | **String** | Website URL of the company associated with the business account. | [optional] [default to null] |
| **companyDescription** | **String** | Description of the company associated with the business account. | [optional] [default to null] |
| **companyName** | **String** | Name of the company associated with the business account. | [optional] [default to null] |
| **settings** | [**Settings**](Settings.md) |  | [optional] [default to null] |
| **accountType** | [**ProfileType**](ProfileType.md) |  | [optional] [default to null] |
| **profileImageUrl** | **String** | Profile image associated with the user account. | [optional] [default to null] |
| **auth0UserId** | **String** |  | [optional] [default to null] |
| **role** | [**Role**](Role.md) |  | [optional] [default to null] |
| **algoliaUserId** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

