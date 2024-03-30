# AccountingServiceApi

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createPayrollProfile**](AccountingServiceApi.md#createPayrollProfile) | **POST** /accounting-microservice/api/v1/profile | Create Payroll Profile |
| [**deleteProfile**](AccountingServiceApi.md#deleteProfile) | **DELETE** /accounting-microservice/api/v1/profile/{authZeroUserId} | Delete Payroll Profile |
| [**exchangePublicLinkTokenForAccountTokenResponse**](AccountingServiceApi.md#exchangePublicLinkTokenForAccountTokenResponse) | **POST** /accounting-microservice/api/v1/merge/exchange-token | Exchange Link Token |
| [**getLinkToken**](AccountingServiceApi.md#getLinkToken) | **POST** /accounting-microservice/api/v1/merge/initiate-token-exchange | Get Link Token |
| [**healthCheck3**](AccountingServiceApi.md#healthCheck3) | **GET** /accounting-microservice/api/v1/health | Health Check |
| [**readAccountingProfileResponse**](AccountingServiceApi.md#readAccountingProfileResponse) | **GET** /accounting-microservice/api/v1/profile/{authZeroUserId} | Get Business Account Profile |
| [**readBalanceSheets**](AccountingServiceApi.md#readBalanceSheets) | **POST** /accounting-microservice/api/v1/balance-sheets | Gets Balance Sheets |
| [**readBusinessChartOfAccounts**](AccountingServiceApi.md#readBusinessChartOfAccounts) | **POST** /accounting-microservice/api/v1/chart-of-accounts | Gets Chart of Accounts |
| [**readBusinessTransactions**](AccountingServiceApi.md#readBusinessTransactions) | **POST** /accounting-microservice/api/v1/business-transactions | Gets Business Transactions |
| [**readCashFlowStatements**](AccountingServiceApi.md#readCashFlowStatements) | **POST** /accounting-microservice/api/v1/cashflow-statements | Gets Cashfloe Sheets |
| [**readIncomeStatements**](AccountingServiceApi.md#readIncomeStatements) | **POST** /accounting-microservice/api/v1/income-statements | Gets Cashfloe Sheets |
| [**readinessCheck**](AccountingServiceApi.md#readinessCheck) | **GET** /accounting-microservice/api/v1/ready | Readiness Check |
| [**updatePayrollProfile**](AccountingServiceApi.md#updatePayrollProfile) | **PUT** /accounting-microservice/api/v1/profile | Update Payroll Profile |


<a name="createPayrollProfile"></a>
# **createPayrollProfile**
> CreateAccountingProfileResponse createPayrollProfile(CreateAccountingProfileRequest)

Create Payroll Profile

    Creates a payroll profile.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateAccountingProfileRequest** | [**CreateAccountingProfileRequest**](../Models/CreateAccountingProfileRequest.md)| Defines a message named CreateAccountingProfileRequest. | |

### Return type

[**CreateAccountingProfileResponse**](../Models/CreateAccountingProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteProfile"></a>
# **deleteProfile**
> DeleteAccountingProfileResponse deleteProfile(authZeroUserId)

Delete Payroll Profile

    Deletes a business payroll profile.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |

### Return type

[**DeleteAccountingProfileResponse**](../Models/DeleteAccountingProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="exchangePublicLinkTokenForAccountTokenResponse"></a>
# **exchangePublicLinkTokenForAccountTokenResponse**
> ExchangePublicLinkTokenForAccountTokenResponse exchangePublicLinkTokenForAccountTokenResponse(ExchangePublicLinkTokenForAccountTokenRequest)

Exchange Link Token

    Exchanges a public link token for an account token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ExchangePublicLinkTokenForAccountTokenRequest** | [**ExchangePublicLinkTokenForAccountTokenRequest**](../Models/ExchangePublicLinkTokenForAccountTokenRequest.md)|  | |

### Return type

[**ExchangePublicLinkTokenForAccountTokenResponse**](../Models/ExchangePublicLinkTokenForAccountTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getLinkToken"></a>
# **getLinkToken**
> GetMergeLinkTokenResponse getLinkToken(GetMergeLinkTokenRequest)

Get Link Token

    Provides a link token to initialize a Link session for the user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **GetMergeLinkTokenRequest** | [**GetMergeLinkTokenRequest**](../Models/GetMergeLinkTokenRequest.md)| Defines a message named GetMergeLinkTokenRequest. | |

### Return type

[**GetMergeLinkTokenResponse**](../Models/GetMergeLinkTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="healthCheck3"></a>
# **healthCheck3**
> HealthCheckResponse2 healthCheck3()

Health Check

    Performs a health check on the service.

### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheckResponse2**](../Models/HealthCheckResponse2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readAccountingProfileResponse"></a>
# **readAccountingProfileResponse**
> ReadAccountingProfileResponse readAccountingProfileResponse(authZeroUserId, linkedAccountingAccountId)

Get Business Account Profile

    Retrieves a business account profile.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **authZeroUserId** | **String**|  | [default to null] |
| **linkedAccountingAccountId** | **String**|  | [default to null] |

### Return type

[**ReadAccountingProfileResponse**](../Models/ReadAccountingProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readBalanceSheets"></a>
# **readBalanceSheets**
> ReadBalanceSheetsResponse readBalanceSheets(ReadBalanceSheetsRequest)

Gets Balance Sheets

    Queries balance sheets for a business.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ReadBalanceSheetsRequest** | [**ReadBalanceSheetsRequest**](../Models/ReadBalanceSheetsRequest.md)|  | |

### Return type

[**ReadBalanceSheetsResponse**](../Models/ReadBalanceSheetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readBusinessChartOfAccounts"></a>
# **readBusinessChartOfAccounts**
> ReadBusinessChartOfAccountsResponse readBusinessChartOfAccounts(ReadBusinessChartOfAccountsRequest)

Gets Chart of Accounts

    Queries chart of accounts for a business.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ReadBusinessChartOfAccountsRequest** | [**ReadBusinessChartOfAccountsRequest**](../Models/ReadBusinessChartOfAccountsRequest.md)|  | |

### Return type

[**ReadBusinessChartOfAccountsResponse**](../Models/ReadBusinessChartOfAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readBusinessTransactions"></a>
# **readBusinessTransactions**
> ReadBusinessTransactionsResponse readBusinessTransactions(ReadBusinessTransactionsRequest)

Gets Business Transactions

    Queries transactions for a business.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ReadBusinessTransactionsRequest** | [**ReadBusinessTransactionsRequest**](../Models/ReadBusinessTransactionsRequest.md)| Request for reading business transactions with pagination and time filtering. | |

### Return type

[**ReadBusinessTransactionsResponse**](../Models/ReadBusinessTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readCashFlowStatements"></a>
# **readCashFlowStatements**
> ReadCashFlowStatementsResponse readCashFlowStatements(ReadCashFlowStatementsRequest)

Gets Cashfloe Sheets

    Queries cashflow sheets for a business.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ReadCashFlowStatementsRequest** | [**ReadCashFlowStatementsRequest**](../Models/ReadCashFlowStatementsRequest.md)|  | |

### Return type

[**ReadCashFlowStatementsResponse**](../Models/ReadCashFlowStatementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readIncomeStatements"></a>
# **readIncomeStatements**
> ReadIncomeStatementsResponse readIncomeStatements(ReadIncomeStatementsRequest)

Gets Cashfloe Sheets

    Queries income sheets for a business.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ReadIncomeStatementsRequest** | [**ReadIncomeStatementsRequest**](../Models/ReadIncomeStatementsRequest.md)|  | |

### Return type

[**ReadIncomeStatementsResponse**](../Models/ReadIncomeStatementsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readinessCheck"></a>
# **readinessCheck**
> ReadynessCheckResponse2 readinessCheck()

Readiness Check

    Performs a readiness check on the service.

### Parameters
This endpoint does not need any parameter.

### Return type

[**ReadynessCheckResponse2**](../Models/ReadynessCheckResponse2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updatePayrollProfile"></a>
# **updatePayrollProfile**
> UpdateAccountingProfileResponse updatePayrollProfile(UpdateAccountingProfileRequest)

Update Payroll Profile

    Updates a payroll profile.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateAccountingProfileRequest** | [**UpdateAccountingProfileRequest**](../Models/UpdateAccountingProfileRequest.md)| Defines a message named UpdateAccountingProfileRequest. | |

### Return type

[**UpdateAccountingProfileResponse**](../Models/UpdateAccountingProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

