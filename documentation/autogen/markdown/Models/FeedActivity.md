# FeedActivity
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **actor** | [**Actor**](Actor.md) |  | [default to null] |
| **verb** | [**PostType**](PostType.md) |  | [optional] [default to null] |
| **regularPost** | [**Post**](Post.md) |  | [optional] [default to null] |
| **sharedPost** | [**SharedPost**](SharedPost.md) |  | [optional] [default to null] |
| **pollPost** | [**PollPost**](PollPost.md) |  | [optional] [default to null] |
| **foreignId** | **String** |  | [optional] [default to null] |
| **target** | **String** | Object Describes the target of the activity. The precise meaning of the activity&#39;s target is dependent on the activities verb, but will often be the object the English preposition \&quot;to\&quot;. For instance, in the activity, \&quot;John saved a movie to his wishlist\&quot;, the target of the activity is \&quot;wishlist\&quot;. | [optional] [default to null] |
| **time** | **String** |  | [optional] [default to null] |
| **origin** | **String** |  | [optional] [default to null] |
| **to** | **List** | The TO field allows you to specify a list of feeds to which the activity should be copied. One way to think about it is as the CC functionality of email. | [optional] [default to null] |
| **score** | **String** |  | [optional] [default to null] |
| **extra** | **Map** |  | [optional] [default to null] |
| **getstreamActivityId** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

