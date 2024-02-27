# AccountApi

All URIs are relative to *http://social-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCannyUserSSOToken**](AccountApi.md#getCannyUserSSOToken) | **GET** /api/v1/user/{userId}/canny | Retrieves user sso token for canny |


<a name="getCannyUserSSOToken"></a>
# **getCannyUserSSOToken**
> GetCannyUserSSOTokenResponse getCannyUserSSOToken(userId)

Retrieves user sso token for canny

    Fetches a user sso token for canny

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |

### Return type

[**GetCannyUserSSOTokenResponse**](../Models/GetCannyUserSSOTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

