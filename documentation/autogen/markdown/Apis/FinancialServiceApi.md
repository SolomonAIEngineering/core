# FinancialServiceApi

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addDefaultPocketsToBankAccount**](FinancialServiceApi.md#addDefaultPocketsToBankAccount) | **POST** /financial-microservice/api/v1/pocket/bank-account | adds a default set of pockets to a specific bank account of interest |
| [**addNoteToFinancialUserProfile**](FinancialServiceApi.md#addNoteToFinancialUserProfile) | **POST** /financial-microservice/api/v1/financial-profile/business/note | Adds a note to a business account |
| [**addNoteToRecurringTransaction**](FinancialServiceApi.md#addNoteToRecurringTransaction) | **POST** /financial-microservice/api/v1/transactions/recurring/note | adds a note to a transaction |
| [**addNoteToSmartGoal**](FinancialServiceApi.md#addNoteToSmartGoal) | **POST** /financial-microservice/api/v1/smart-goal/note | adds a note to a smart goal |
| [**addNoteToTransaction**](FinancialServiceApi.md#addNoteToTransaction) | **POST** /financial-microservice/api/v1/transactions/transaction/note | adds a note to a transaction |
| [**addTransactionsToManuallyLinkedAccount**](FinancialServiceApi.md#addTransactionsToManuallyLinkedAccount) | **POST** /financial-microservice/api/v1/manual-linked-account/transactions | Adds transactions to a manually linked account |
| [**askCopilotQuestion**](FinancialServiceApi.md#askCopilotQuestion) | **POST** /financial-microservice/api/v1/copilot/quota/question | Ask a question to copilot |
| [**bulkUpdateRecurringTransaction**](FinancialServiceApi.md#bulkUpdateRecurringTransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring/bulk | update a transaction |
| [**bulkUpdateTransaction**](FinancialServiceApi.md#bulkUpdateTransaction) | **PUT** /financial-microservice/api/v1/transactions/transaction/bulk | update a transaction |
| [**checkIfQuotaExceeded**](FinancialServiceApi.md#checkIfQuotaExceeded) | **GET** /financial-microservice/api/v1/copilot/quota/exceeded/{userId} | Checks if the question quota has been exceeded |
| [**createBankAccount**](FinancialServiceApi.md#createBankAccount) | **POST** /financial-microservice/api/v1/bank-account/profile | create a bank account for a given user profile |
| [**createBudget**](FinancialServiceApi.md#createBudget) | **POST** /financial-microservice/api/v1/budget | create a budget |
| [**createLink**](FinancialServiceApi.md#createLink) | **POST** /financial-microservice/api/v1/link | create link |
| [**createMilestone**](FinancialServiceApi.md#createMilestone) | **POST** /financial-microservice/api/v1/milestone | create a milestone |
| [**createSmartGoal**](FinancialServiceApi.md#createSmartGoal) | **POST** /financial-microservice/api/v1/smart-goal | create a smart goal |
| [**createSubscription**](FinancialServiceApi.md#createSubscription) | **POST** /financial-microservice/api/v1/stripe/subscription | Creates a new subscription for a given customer against stripe |
| [**createUserProfile1**](FinancialServiceApi.md#createUserProfile1) | **POST** /financial-microservice/api/v1/profile | create a user profile |
| [**deleteBudget**](FinancialServiceApi.md#deleteBudget) | **DELETE** /financial-microservice/api/v1/budget/{budgetId} | delete a budget |
| [**deleteLink**](FinancialServiceApi.md#deleteLink) | **DELETE** /financial-microservice/api/v1/link/{linkId}/user/{userId} | delete link by id |
| [**deleteMilestone**](FinancialServiceApi.md#deleteMilestone) | **DELETE** /financial-microservice/api/v1/milestone/{milestoneId} | delete a milestone |
| [**deleteNoteFromRecurringTransaction**](FinancialServiceApi.md#deleteNoteFromRecurringTransaction) | **DELETE** /financial-microservice/api/v1/transactions/recurring/{transactionId}/note/{noteId} | deletes a note from a transaction |
| [**deleteNoteFromSmartGoal**](FinancialServiceApi.md#deleteNoteFromSmartGoal) | **DELETE** /financial-microservice/api/v1/smart-goal/note/{noteId} | deletes a note from a smart goal |
| [**deleteNoteFromTransaction**](FinancialServiceApi.md#deleteNoteFromTransaction) | **DELETE** /financial-microservice/api/v1/transactions/transaction/{transactionId}/note/{noteId} | deletes a note from a transaction |
| [**deletePocket**](FinancialServiceApi.md#deletePocket) | **DELETE** /financial-microservice/api/v1/pocket/{pocketId} | deletes a pocket |
| [**deleteSmartGoal**](FinancialServiceApi.md#deleteSmartGoal) | **DELETE** /financial-microservice/api/v1/smart-goal/{smartGoalId} | delete a smart goal |
| [**deleteTransaction**](FinancialServiceApi.md#deleteTransaction) | **DELETE** /financial-microservice/api/v1/transactions/recurring/{transactionId} | deletes a transaction by id |
| [**deleteTransaction1**](FinancialServiceApi.md#deleteTransaction1) | **DELETE** /financial-microservice/api/v1/transactions/transaction/{transactionId} | deletes a transaction by id |
| [**deleteUserProfile1**](FinancialServiceApi.md#deleteUserProfile1) | **DELETE** /financial-microservice/api/v1/bank-account/{bankAccountId} | deletes a bank account for a given user profile |
| [**deleteUserProfile2**](FinancialServiceApi.md#deleteUserProfile2) | **DELETE** /financial-microservice/api/v1/profile/{userId} | deletes a user profile |
| [**exchangePlaidToken**](FinancialServiceApi.md#exchangePlaidToken) | **POST** /financial-microservice/api/v1/plaid/exchange-token | exchange plaid token |
| [**getAccountBalance**](FinancialServiceApi.md#getAccountBalance) | **GET** /financial-microservice/api/v1/historical-account-balance/user/{userId}/plaid-account-id/{plaidAccountId} | gets account balance of an account |
| [**getAccountBalanceHistory**](FinancialServiceApi.md#getAccountBalanceHistory) | **GET** /financial-microservice/api/v1/analytics/balance-history/account/{plaidAccountId}/pagenumber/{pageNumber}/pagesize/{pageSize} | Returns the account balance history for an account |
| [**getAllBudgets**](FinancialServiceApi.md#getAllBudgets) | **GET** /financial-microservice/api/v1/budget | get all budgets |
| [**getBankAccount**](FinancialServiceApi.md#getBankAccount) | **GET** /financial-microservice/api/v1/bank-account/{bankAccountId} | get a bank account for a given user profile |
| [**getBudget**](FinancialServiceApi.md#getBudget) | **GET** /financial-microservice/api/v1/budget/{budgetId} | get budget by id |
| [**getCategoryMetricsFinancialSubProfileOverTime**](FinancialServiceApi.md#getCategoryMetricsFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/category-metrics | Gets category metrics for a financial sub profile over time |
| [**getCategoryMonthlyTransactionCount**](FinancialServiceApi.md#getCategoryMonthlyTransactionCount) | **GET** /financial-microservice/api/v1/analytics/category-monthly-transaction-count/user/{userId} | Get monthly transaction count by user, month, and category |
| [**getDebtToIncomeRatio**](FinancialServiceApi.md#getDebtToIncomeRatio) | **GET** /financial-microservice/api/v1/analytics/debt-to-income-ratio/user/{userId} | Get Debt-to-Income ratio by user and month |
| [**getExpenseMetrics**](FinancialServiceApi.md#getExpenseMetrics) | **GET** /financial-microservice/api/v1/analytics/expenses/user/{userId} | Get Expense Metrics by user, month and category |
| [**getExpenseMetricsFinancialSubProfileOverTime**](FinancialServiceApi.md#getExpenseMetricsFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/expense-metrics | Gets expense metrics for a financial sub profile over time |
| [**getFinancialProfile**](FinancialServiceApi.md#getFinancialProfile) | **GET** /financial-microservice/api/v1/analytics/finance-profile/user/{userId} | Get Financial Profile by user and month |
| [**getForecast**](FinancialServiceApi.md#getForecast) | **GET** /financial-microservice/api/v1/forecast/{smartGoalId} | get forecast by id |
| [**getIncomeExpenseRatio**](FinancialServiceApi.md#getIncomeExpenseRatio) | **GET** /financial-microservice/api/v1/analytics/income-expense-ratio/user/{userId} | Get Income Expense Ratio by user and month |
| [**getIncomeMetrics**](FinancialServiceApi.md#getIncomeMetrics) | **GET** /financial-microservice/api/v1/analytics/income/user/{userId} | Get Income Metrics by user, month and category |
| [**getIncomeMetricsFinancialSubProfileOverTime**](FinancialServiceApi.md#getIncomeMetricsFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/income-metrics | Gets income metrics for a financial sub profile over time |
| [**getInvestmentAccount**](FinancialServiceApi.md#getInvestmentAccount) | **GET** /financial-microservice/api/v1/account/{userId}/investment/{investmentAccountId} | get investment account by id |
| [**getLiabilityAccount**](FinancialServiceApi.md#getLiabilityAccount) | **GET** /financial-microservice/api/v1/account/{userId}/liability/{liabilityAccountId} | get liability account by id |
| [**getLink**](FinancialServiceApi.md#getLink) | **GET** /financial-microservice/api/v1/link/{linkId} | get link by id |
| [**getLinks**](FinancialServiceApi.md#getLinks) | **GET** /financial-microservice/api/v1/links/{userId} | get links |
| [**getLocationMetricsFinancialSubProfileOverTime**](FinancialServiceApi.md#getLocationMetricsFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/location-metrics | Gets income metrics for a financial sub profile over time |
| [**getMelodyFinancialContext**](FinancialServiceApi.md#getMelodyFinancialContext) | **GET** /financial-microservice/api/v1/analytics/melody-financial-context/user/{userId} | Get Melody Financial Context |
| [**getMerchantMetricsFinancialSubProfileOverTime**](FinancialServiceApi.md#getMerchantMetricsFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/merchant-metrics | Gets merchant metrics for a financial sub profile over time |
| [**getMerchantMonthlyExpenditure**](FinancialServiceApi.md#getMerchantMonthlyExpenditure) | **GET** /financial-microservice/api/v1/analytics/merchant-monthly-expenditure/user/{userId} | Get Merchant Monthly Expenditure by user, month and merchant name |
| [**getMilestone**](FinancialServiceApi.md#getMilestone) | **GET** /financial-microservice/api/v1/milestone/{milestoneId} | get milestone by id |
| [**getMilestones**](FinancialServiceApi.md#getMilestones) | **GET** /financial-microservice/api/v1/milestone/smart-goal/{smartGoalId} | get milestones by smart goal id |
| [**getMonthlyBalance**](FinancialServiceApi.md#getMonthlyBalance) | **GET** /financial-microservice/api/v1/analytics/monthly-balance/user/{userId} | Get Monthly Balance by user and month |
| [**getMonthlyExpenditure**](FinancialServiceApi.md#getMonthlyExpenditure) | **GET** /financial-microservice/api/v1/analytics/monthly-expenditure/user/{userId} | Get Monthly Expenditure by user and month |
| [**getMonthlyIncome**](FinancialServiceApi.md#getMonthlyIncome) | **GET** /financial-microservice/api/v1/analytics/monthly-income/user/{userId} | Get Monthly Income by user and month |
| [**getMonthlySavings**](FinancialServiceApi.md#getMonthlySavings) | **GET** /financial-microservice/api/v1/analytics/monthly-savings/{userId} | Get Monthly Savings by user and month |
| [**getMonthlyTotalQuantityBySecurityAndUser**](FinancialServiceApi.md#getMonthlyTotalQuantityBySecurityAndUser) | **GET** /financial-microservice/api/v1/analytics/monthly-total-quantity-by-security-and-user/user/{userId} | Get Monthly Total Quantity of Security by user, month and security |
| [**getMonthlyTransactionCount**](FinancialServiceApi.md#getMonthlyTransactionCount) | **GET** /financial-microservice/api/v1/analytics/monthly-transaction-count/user/{userId} | Get Monthly Transaction Count by user and month |
| [**getMortageAccount**](FinancialServiceApi.md#getMortageAccount) | **GET** /financial-microservice/api/v1/account/{userId}/mortgage/{mortgageAccountId} | get mortgage account by id |
| [**getNoteFromSmartGoal**](FinancialServiceApi.md#getNoteFromSmartGoal) | **GET** /financial-microservice/api/v1/smart-goal/note/{noteId} | gets a note from a smart goal |
| [**getNoteFromTransaction**](FinancialServiceApi.md#getNoteFromTransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/note/{noteId} | gets a note from a transaction |
| [**getNotesFromFinancialUserProfile**](FinancialServiceApi.md#getNotesFromFinancialUserProfile) | **GET** /financial-microservice/api/v1/financial-profile/business/{businessAccountUserId}/{profileType}/note | Gets notes from a business account |
| [**getNotesFromSmartGoal**](FinancialServiceApi.md#getNotesFromSmartGoal) | **GET** /financial-microservice/api/v1/smart-goal/{smartGoalId}/note | gets notes from a smart goal |
| [**getPaymentChannelFinancialSubProfileOverTime**](FinancialServiceApi.md#getPaymentChannelFinancialSubProfileOverTime) | **GET** /financial-microservice/api/v1/financial-profile/payment-channel-metrics | Gets payment metrics for a financial sub profile over time |
| [**getPaymentChannelMonthlyExpenditure**](FinancialServiceApi.md#getPaymentChannelMonthlyExpenditure) | **GET** /financial-microservice/api/v1/analytics/payment-channel-expenditure/user/{userId} | Get Payment Channel Monthly Expenditure by user, month, and payment channel |
| [**getPocket**](FinancialServiceApi.md#getPocket) | **GET** /financial-microservice/api/v1/pocket/{pocketId} | get a pocket |
| [**getRecurringTransaction**](FinancialServiceApi.md#getRecurringTransaction) | **GET** /financial-microservice/api/v1/transactions/recurring/{transactionId} | lists a set of transactions against a given account of interest |
| [**getRecurringTransactionsForUser**](FinancialServiceApi.md#getRecurringTransactionsForUser) | **GET** /financial-microservice/api/v1/transactions/recurring-transactions/{userId} | get recurring transactions |
| [**getSmartGoalsByPocketId**](FinancialServiceApi.md#getSmartGoalsByPocketId) | **GET** /financial-microservice/api/v1/smart-goal/pocket/{pocketId} | get smart goals by pocket id |
| [**getSplitTransaction**](FinancialServiceApi.md#getSplitTransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/split | gets a split transaction |
| [**getStudentLoanAccount**](FinancialServiceApi.md#getStudentLoanAccount) | **GET** /financial-microservice/api/v1/account/{userId}/student-loan/{studentLoanAccountId} | get student loan account by id |
| [**getTotalInvestmentBySecurity**](FinancialServiceApi.md#getTotalInvestmentBySecurity) | **GET** /financial-microservice/api/v1/analytics/total-investment/user/{userId} | Get Total Investment by user and security |
| [**getTransaction**](FinancialServiceApi.md#getTransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId} | lists a set of transactions against a given account of interest |
| [**getTransactions**](FinancialServiceApi.md#getTransactions) | **GET** /financial-microservice/api/v1/transactions/user/{userId}/plaid-account-id/{plaidAccountId}/pageNumber/{pageNumber}/pageSize/{pageSize} | get transactions tied to a bank account and account id |
| [**getTransactions1**](FinancialServiceApi.md#getTransactions1) | **GET** /financial-microservice/api/v1/transactions/{userId}/pageNumber/{pageNumber}/pageSize/{pageSize} | get transactions |
| [**getTransactionsByTime**](FinancialServiceApi.md#getTransactionsByTime) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/range | get transactions by time |
| [**getTransactionsForPastMonth**](FinancialServiceApi.md#getTransactionsForPastMonth) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/month | Get transactions for the past month |
| [**getTransactionsForPastWeek**](FinancialServiceApi.md#getTransactionsForPastWeek) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/week | get transactions for the past week |
| [**getUserAccountBalanceHistory**](FinancialServiceApi.md#getUserAccountBalanceHistory) | **GET** /financial-microservice/api/v1/analytics/balance-history/user/{userId}/pagenumber/{pageNumber}/pagesize/{pageSize} | Returns the account balance history for a user |
| [**getUserCategoryMonthlyExpenditure**](FinancialServiceApi.md#getUserCategoryMonthlyExpenditure) | **GET** /financial-microservice/api/v1/analytics/category-monthly-expenditure/user/{userId} | Returns the monthly category expenditure for a user |
| [**getUserCategoryMonthlyIncome**](FinancialServiceApi.md#getUserCategoryMonthlyIncome) | **GET** /financial-microservice/api/v1/analytics/category-monthly-income/user/{userId} | Get monthly income by user for a specific category |
| [**getUserProfile1**](FinancialServiceApi.md#getUserProfile1) | **GET** /financial-microservice/api/v1/profile/{userId} | Gets a user profile |
| [**healthCheck2**](FinancialServiceApi.md#healthCheck2) | **GET** /financial-microservice/api/v1/health | health check |
| [**initiatePlaidSetup**](FinancialServiceApi.md#initiatePlaidSetup) | **POST** /financial-microservice/api/v1/plaid/initiate-token-exchange | initiate plaid setup |
| [**initiatePlaidTokenUpdate**](FinancialServiceApi.md#initiatePlaidTokenUpdate) | **POST** /financial-microservice/api/v1/plaid/initiate-token-update | initiate plaid link token update |
| [**listRecurringTransactionNotes**](FinancialServiceApi.md#listRecurringTransactionNotes) | **GET** /financial-microservice/api/v1/transactions/recurring/{transactionId}/notes | lists notes from a transaction |
| [**listRecurringTransactionsForUserAndAccount**](FinancialServiceApi.md#listRecurringTransactionsForUserAndAccount) | **GET** /financial-microservice/api/v1/transactions/recurrings | lists a set of transactions against a given account of interest |
| [**listTransactionNotes**](FinancialServiceApi.md#listTransactionNotes) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/notes | lists notes from a transaction |
| [**listTransactions**](FinancialServiceApi.md#listTransactions) | **GET** /financial-microservice/api/v1/transactions | lists a set of transactions against a given account of interest |
| [**listTransactions1**](FinancialServiceApi.md#listTransactions1) | **GET** /financial-microservice/api/v1/transactions/all_accounts | lists a set of transactions across all connected accounts |
| [**pollAsyncTaskExecutionStatus**](FinancialServiceApi.md#pollAsyncTaskExecutionStatus) | **GET** /financial-microservice/api/v1/async-task/{workflowId}/run/{runId} | polls the status of an async task |
| [**readynessCheck2**](FinancialServiceApi.md#readynessCheck2) | **GET** /financial-microservice/api/v1/ready | readyness check |
| [**searchTransactions**](FinancialServiceApi.md#searchTransactions) | **POST** /financial-microservice/api/v1/transactions/transaction/search | searches transactions |
| [**splitTransaction**](FinancialServiceApi.md#splitTransaction) | **POST** /financial-microservice/api/v1/transactions/transaction/split | splits a transaction |
| [**transactionAggregates**](FinancialServiceApi.md#transactionAggregates) | **GET** /financial-microservice/api/v1/analytics/transaction-aggregates/{userId} | Returns the aggregated transactions for a user and month |
| [**triggerSync**](FinancialServiceApi.md#triggerSync) | **POST** /financial-microservice/api/v1/sync/trigger | Triggers a sync |
| [**unsplitTransactions**](FinancialServiceApi.md#unsplitTransactions) | **POST** /financial-microservice/api/v1/transactions/transaction/unsplit | unsplit a transaction |
| [**updateBankAccount**](FinancialServiceApi.md#updateBankAccount) | **PUT** /financial-microservice/api/v1/bank-account | update a bank account for a given user profile |
| [**updateNoteToRecurringTransaction**](FinancialServiceApi.md#updateNoteToRecurringTransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring/note | Updates a note to a transaction |
| [**updateNoteToSmartGoal**](FinancialServiceApi.md#updateNoteToSmartGoal) | **PUT** /financial-microservice/api/v1/smart-goal/note | updates a note to a smart goal |
| [**updateNoteToTransaction**](FinancialServiceApi.md#updateNoteToTransaction) | **PUT** /financial-microservice/api/v1/transactions/transaction/note | Updates a note to a transaction |
| [**updatePocket**](FinancialServiceApi.md#updatePocket) | **PUT** /financial-microservice/api/v1/pocket | updates a pocket |
| [**updateSingleTransaction**](FinancialServiceApi.md#updateSingleTransaction) | **PUT** /financial-microservice/api/v1/transactions/single-transaction | update a transaction |
| [**updateSmartGoal**](FinancialServiceApi.md#updateSmartGoal) | **PUT** /financial-microservice/api/v1/smart-goal | update a smart goal |
| [**updateTransaction**](FinancialServiceApi.md#updateTransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring | update a transaction |
| [**updateUserProfile**](FinancialServiceApi.md#updateUserProfile) | **PUT** /financial-microservice/api/v1/profile | update a user profile |
| [**updatesBudget**](FinancialServiceApi.md#updatesBudget) | **PUT** /financial-microservice/api/v1/budget | updates a budget |
| [**updatesMilestone**](FinancialServiceApi.md#updatesMilestone) | **PUT** /financial-microservice/api/v1/milestone | updates a milestone |


<a name="addDefaultPocketsToBankAccount"></a>
# **addDefaultPocketsToBankAccount**
> AddDefaultPocketsToBankAccountResponse addDefaultPocketsToBankAccount(AddDefaultPocketsToBankAccountRequest)

adds a default set of pockets to a specific bank account of interest

    This endpoint adds a default pocket to a specific bank account of interest

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddDefaultPocketsToBankAccountRequest** | [**AddDefaultPocketsToBankAccountRequest**](../Models/AddDefaultPocketsToBankAccountRequest.md)|  | |

### Return type

[**AddDefaultPocketsToBankAccountResponse**](../Models/AddDefaultPocketsToBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="addNoteToFinancialUserProfile"></a>
# **addNoteToFinancialUserProfile**
> AddNoteToFinancialUserProfileResponse addNoteToFinancialUserProfile(AddNoteToFinancialUserProfileRequest)

Adds a note to a business account

    This endpoint adds a note to a business account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddNoteToFinancialUserProfileRequest** | [**AddNoteToFinancialUserProfileRequest**](../Models/AddNoteToFinancialUserProfileRequest.md)|  | |

### Return type

[**AddNoteToFinancialUserProfileResponse**](../Models/AddNoteToFinancialUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="addNoteToRecurringTransaction"></a>
# **addNoteToRecurringTransaction**
> AddNoteToRecurringTransactionResponse addNoteToRecurringTransaction(AddNoteToRecurringTransactionRequest)

adds a note to a transaction

    This endpoint adds a note to a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddNoteToRecurringTransactionRequest** | [**AddNoteToRecurringTransactionRequest**](../Models/AddNoteToRecurringTransactionRequest.md)|  | |

### Return type

[**AddNoteToRecurringTransactionResponse**](../Models/AddNoteToRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="addNoteToSmartGoal"></a>
# **addNoteToSmartGoal**
> AddNoteToSmartGoalResponse addNoteToSmartGoal(AddNoteToSmartGoalRequest)

adds a note to a smart goal

    This endpoint adds a note to a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddNoteToSmartGoalRequest** | [**AddNoteToSmartGoalRequest**](../Models/AddNoteToSmartGoalRequest.md)|  | |

### Return type

[**AddNoteToSmartGoalResponse**](../Models/AddNoteToSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="addNoteToTransaction"></a>
# **addNoteToTransaction**
> AddNoteToTransactionResponse addNoteToTransaction(AddNoteToTransactionRequest)

adds a note to a transaction

    This endpoint adds a note to a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddNoteToTransactionRequest** | [**AddNoteToTransactionRequest**](../Models/AddNoteToTransactionRequest.md)|  | |

### Return type

[**AddNoteToTransactionResponse**](../Models/AddNoteToTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="addTransactionsToManuallyLinkedAccount"></a>
# **addTransactionsToManuallyLinkedAccount**
> AddTransactionsToManuallyLinkedAccountResponse addTransactionsToManuallyLinkedAccount(AddTransactionsToManuallyLinkedAccountRequest)

Adds transactions to a manually linked account

    This endpoint adds transactions to a manually linked account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **AddTransactionsToManuallyLinkedAccountRequest** | [**AddTransactionsToManuallyLinkedAccountRequest**](../Models/AddTransactionsToManuallyLinkedAccountRequest.md)| Add transactions to an account. | |

### Return type

[**AddTransactionsToManuallyLinkedAccountResponse**](../Models/AddTransactionsToManuallyLinkedAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="askCopilotQuestion"></a>
# **askCopilotQuestion**
> RecordAskCopilotQuestionResponse askCopilotQuestion(RecordAskCopilotQuestionRequest)

Ask a question to copilot

    This endpoint checks if a user can ask his/her copilot a question

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **RecordAskCopilotQuestionRequest** | [**RecordAskCopilotQuestionRequest**](../Models/RecordAskCopilotQuestionRequest.md)|  | |

### Return type

[**RecordAskCopilotQuestionResponse**](../Models/RecordAskCopilotQuestionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="bulkUpdateRecurringTransaction"></a>
# **bulkUpdateRecurringTransaction**
> BulkUpdateRecurringTransactionResponse bulkUpdateRecurringTransaction(BulkUpdateRecurringTransactionRequest)

update a transaction

    This endpoint updates a set of transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **BulkUpdateRecurringTransactionRequest** | [**BulkUpdateRecurringTransactionRequest**](../Models/BulkUpdateRecurringTransactionRequest.md)|  | |

### Return type

[**BulkUpdateRecurringTransactionResponse**](../Models/BulkUpdateRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="bulkUpdateTransaction"></a>
# **bulkUpdateTransaction**
> BulkUpdateTransactionResponse bulkUpdateTransaction(BulkUpdateTransactionRequest)

update a transaction

    This endpoint updates a set of transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **BulkUpdateTransactionRequest** | [**BulkUpdateTransactionRequest**](../Models/BulkUpdateTransactionRequest.md)|  | |

### Return type

[**BulkUpdateTransactionResponse**](../Models/BulkUpdateTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="checkIfQuotaExceeded"></a>
# **checkIfQuotaExceeded**
> CheckIfQuotaExceededResponse checkIfQuotaExceeded(userId, profileType)

Checks if the question quota has been exceeded

    This endpoint checks if a user has exceeded his/her copilot question quota

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| the account id associated with the user | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**CheckIfQuotaExceededResponse**](../Models/CheckIfQuotaExceededResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="createBankAccount"></a>
# **createBankAccount**
> CreateBankAccountResponse createBankAccount(CreateBankAccountRequest)

create a bank account for a given user profile

    This endpoint creates a bank account for a given user profile

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateBankAccountRequest** | [**CreateBankAccountRequest**](../Models/CreateBankAccountRequest.md)|  | |

### Return type

[**CreateBankAccountResponse**](../Models/CreateBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createBudget"></a>
# **createBudget**
> CreateBudgetResponse createBudget(CreateBudgetRequest)

create a budget

    This endpoint creates a budget

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateBudgetRequest** | [**CreateBudgetRequest**](../Models/CreateBudgetRequest.md)|  | |

### Return type

[**CreateBudgetResponse**](../Models/CreateBudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createLink"></a>
# **createLink**
> CreateManualLinkResponse createLink(CreateManualLinkRequest)

create link

    This endpoint creates a link

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateManualLinkRequest** | [**CreateManualLinkRequest**](../Models/CreateManualLinkRequest.md)|  | |

### Return type

[**CreateManualLinkResponse**](../Models/CreateManualLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createMilestone"></a>
# **createMilestone**
> CreateMilestoneResponse createMilestone(CreateMilestoneRequest)

create a milestone

    This endpoint creates a milestone

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateMilestoneRequest** | [**CreateMilestoneRequest**](../Models/CreateMilestoneRequest.md)|  | |

### Return type

[**CreateMilestoneResponse**](../Models/CreateMilestoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createSmartGoal"></a>
# **createSmartGoal**
> CreateSmartGoalResponse createSmartGoal(CreateSmartGoalRequest)

create a smart goal

    This endpoint creates a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateSmartGoalRequest** | [**CreateSmartGoalRequest**](../Models/CreateSmartGoalRequest.md)|  | |

### Return type

[**CreateSmartGoalResponse**](../Models/CreateSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createSubscription"></a>
# **createSubscription**
> CreateSubscriptionResponse createSubscription(CreateSubscriptionRequest)

Creates a new subscription for a given customer against stripe

    This endpoint enabled a user to create a new subscription against stripe

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateSubscriptionRequest** | [**CreateSubscriptionRequest**](../Models/CreateSubscriptionRequest.md)|  | |

### Return type

[**CreateSubscriptionResponse**](../Models/CreateSubscriptionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createUserProfile1"></a>
# **createUserProfile1**
> CreateUserProfileResponse1 createUserProfile1(CreateUserProfileRequest1)

create a user profile

    This endpoint performs an a creation operation of a user profile based on the provided parametersThis operation is implemented as a distributed transactions as this operation can span multiple services

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateUserProfileRequest1** | [**CreateUserProfileRequest1**](../Models/CreateUserProfileRequest1.md)|  | |

### Return type

[**CreateUserProfileResponse1**](../Models/CreateUserProfileResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteBudget"></a>
# **deleteBudget**
> DeleteBudgetResponse deleteBudget(budgetId)

delete a budget

    This endpoint deletes a budget

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **budgetId** | **String**| The budget id Validations: - budget_id must be greater than 0 | [default to null] |

### Return type

[**DeleteBudgetResponse**](../Models/DeleteBudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteLink"></a>
# **deleteLink**
> DeleteLinkResponse deleteLink(linkId, userId, profileType)

delete link by id

    This endpoint deletes the link if the link exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **linkId** | **String**| The link id Validations: - link_id must be greater than 0 | [default to null] |
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**DeleteLinkResponse**](../Models/DeleteLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteMilestone"></a>
# **deleteMilestone**
> DeleteMilestoneResponse deleteMilestone(milestoneId)

delete a milestone

    This endpoint deletes a milestone

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **milestoneId** | **String**| The milestone id Validations: - milestone_id must be greater than 0 | [default to null] |

### Return type

[**DeleteMilestoneResponse**](../Models/DeleteMilestoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteNoteFromRecurringTransaction"></a>
# **deleteNoteFromRecurringTransaction**
> DeleteNoteFromRecurringTransactionResponse deleteNoteFromRecurringTransaction(transactionId, noteId)

deletes a note from a transaction

    This endpoint deletes a note from a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |
| **noteId** | **String**| The note id Validations: - note_id must be greater than 0 | [default to null] |

### Return type

[**DeleteNoteFromRecurringTransactionResponse**](../Models/DeleteNoteFromRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteNoteFromSmartGoal"></a>
# **deleteNoteFromSmartGoal**
> DeleteNoteFromSmartGoalResponse deleteNoteFromSmartGoal(noteId, smartGoalId)

deletes a note from a smart goal

    This endpoint deletes a note from a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **noteId** | **String**| The note id Validations: - note_id must be greater than 0 | [default to null] |
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |

### Return type

[**DeleteNoteFromSmartGoalResponse**](../Models/DeleteNoteFromSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteNoteFromTransaction"></a>
# **deleteNoteFromTransaction**
> DeleteNoteFromTransactionResponse deleteNoteFromTransaction(transactionId, noteId)

deletes a note from a transaction

    This endpoint deletes a note from a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |
| **noteId** | **String**| The note id Validations: - note_id must be greater than 0 | [default to null] |

### Return type

[**DeleteNoteFromTransactionResponse**](../Models/DeleteNoteFromTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deletePocket"></a>
# **deletePocket**
> DeletePocketResponse deletePocket(pocketId)

deletes a pocket

    This endpoint deletes a pocket

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pocketId** | **String**| The pocket id Validations: - pocket_id must be greater than 0 | [default to null] |

### Return type

[**DeletePocketResponse**](../Models/DeletePocketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteSmartGoal"></a>
# **deleteSmartGoal**
> DeleteSmartGoalResponse deleteSmartGoal(smartGoalId)

delete a smart goal

    This endpoint deletes a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |

### Return type

[**DeleteSmartGoalResponse**](../Models/DeleteSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteTransaction"></a>
# **deleteTransaction**
> DeleteRecurringTransactionResponse deleteTransaction(transactionId)

deletes a transaction by id

    This endpoint deletes a specific transaction based on the transaction id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| the transaction of interest we aim to delete | [default to null] |

### Return type

[**DeleteRecurringTransactionResponse**](../Models/DeleteRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteTransaction1"></a>
# **deleteTransaction1**
> DeleteTransactionResponse deleteTransaction1(transactionId)

deletes a transaction by id

    This endpoint deletes a specific transaction based on the transaction id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| the transaction of interest we aim to delete | [default to null] |

### Return type

[**DeleteTransactionResponse**](../Models/DeleteTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteUserProfile1"></a>
# **deleteUserProfile1**
> DeleteBankAccountResponse deleteUserProfile1(bankAccountId, userId, profileType)

deletes a bank account for a given user profile

    This endpoint performs a delete operation on a user profile based on the provided parametersThis deletion operation spans multiple services (plaid) as user details are stored across a suite of our backend servicesThe operation itself is an atomic one of nature. Either all services successfully delete the recod or we fail the requestDivergent state is not expected to be encountered with this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **bankAccountId** | **String**| The bank account id Validations: - bank_account_id must be greater than 0 | [default to null] |
| **userId** | **String**| The account ID associated with the user Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**DeleteBankAccountResponse**](../Models/DeleteBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteUserProfile2"></a>
# **deleteUserProfile2**
> DeleteUserProfileResponse1 deleteUserProfile2(userId, profileType)

deletes a user profile

    This endpoint performs a delete operation on a user profile based on the provided parameters

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user. NOTE: This user_id is the simfiny backend platform wide user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**DeleteUserProfileResponse1**](../Models/DeleteUserProfileResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="exchangePlaidToken"></a>
# **exchangePlaidToken**
> PlaidExchangeTokenResponse exchangePlaidToken(PlaidExchangeTokenRequest)

exchange plaid token

    This endpoint exchanges a plaid link token for a plaid access token

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PlaidExchangeTokenRequest** | [**PlaidExchangeTokenRequest**](../Models/PlaidExchangeTokenRequest.md)|  | |

### Return type

[**PlaidExchangeTokenResponse**](../Models/PlaidExchangeTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="getAccountBalance"></a>
# **getAccountBalance**
> GetHistoricalAccountBalanceResponse getAccountBalance(userId, plaidAccountId, profileType)

gets account balance of an account

    This endpoint returns the historical balance for a given account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **plaidAccountId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetHistoricalAccountBalanceResponse**](../Models/GetHistoricalAccountBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAccountBalanceHistory"></a>
# **getAccountBalanceHistory**
> GetAccountBalanceHistoryResponse getAccountBalanceHistory(plaidAccountId, pageNumber, pageSize)

Returns the account balance history for an account

    This API could accept account_id as input parameters and return the account balance history for that account.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **plaidAccountId** | **String**| Account ID | [default to null] |
| **pageNumber** | **String**|  | [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [default to null] |

### Return type

[**GetAccountBalanceHistoryResponse**](../Models/GetAccountBalanceHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllBudgets"></a>
# **getAllBudgets**
> GetAllBudgetsResponse getAllBudgets(pocketId, smartGoalId, milestoneId)

get all budgets

    This endpoint returns all budgets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pocketId** | **String**| The pocket account id Validations: - pocket_account_id must be greater than 0 | [default to null] |
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |
| **milestoneId** | **String**| The milestone id Validations: - milestone_id must be greater than 0 | [default to null] |

### Return type

[**GetAllBudgetsResponse**](../Models/GetAllBudgetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getBankAccount"></a>
# **getBankAccount**
> GetBankAccountResponse getBankAccount(bankAccountId)

get a bank account for a given user profile

    This endpoint returns the bank account if the user record id and bank account id exists example: /api/v1/bank-account?user_id&#x3D;xxxxxx&amp;&amp;bank_account_id&#x3D;xxxxxxx

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **bankAccountId** | **String**| The bank account id Validations: - bank_account_id must be greater than 0 | [default to null] |

### Return type

[**GetBankAccountResponse**](../Models/GetBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getBudget"></a>
# **getBudget**
> GetBudgetResponse getBudget(budgetId)

get budget by id

    This endpoint returns the budget if the budget exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **budgetId** | **String**| The budget id Validations: - budget_id must be greater than 0 | [default to null] |

### Return type

[**GetBudgetResponse**](../Models/GetBudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCategoryMetricsFinancialSubProfileOverTime"></a>
# **getCategoryMetricsFinancialSubProfileOverTime**
> GetCategoryMetricsFinancialSubProfileOverTimeResponse getCategoryMetricsFinancialSubProfileOverTime(userId, profileType, personalFinanceCategoryPrimary, month, pageNumber, pageSize)

Gets category metrics for a financial sub profile over time

    This endpoint gets category metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **personalFinanceCategoryPrimary** | **String**|  | [optional] [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetCategoryMetricsFinancialSubProfileOverTimeResponse**](../Models/GetCategoryMetricsFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCategoryMonthlyTransactionCount"></a>
# **getCategoryMonthlyTransactionCount**
> GetCategoryMonthlyTransactionCountResponse getCategoryMonthlyTransactionCount(userId, month, personalFinanceCategoryPrimary, pageNumber, pageSize, profileType)

Get monthly transaction count by user, month, and category

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| has to be present and defined | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **personalFinanceCategoryPrimary** | **String**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetCategoryMonthlyTransactionCountResponse**](../Models/GetCategoryMonthlyTransactionCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDebtToIncomeRatio"></a>
# **getDebtToIncomeRatio**
> GetDebtToIncomeRatioResponse getDebtToIncomeRatio(userId, month, pageNumber, pageSize, profileType)

Get Debt-to-Income ratio by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetDebtToIncomeRatioResponse**](../Models/GetDebtToIncomeRatioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getExpenseMetrics"></a>
# **getExpenseMetrics**
> GetExpenseMetricsResponse getExpenseMetrics(userId, month, personalFinanceCategoryPrimary, pageNumber, pageSize, profileType)

Get Expense Metrics by user, month and category

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optonal | [optional] [default to null] |
| **personalFinanceCategoryPrimary** | **String**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetExpenseMetricsResponse**](../Models/GetExpenseMetricsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getExpenseMetricsFinancialSubProfileOverTime"></a>
# **getExpenseMetricsFinancialSubProfileOverTime**
> GetExpenseMetricsFinancialSubProfileOverTimeResponse getExpenseMetricsFinancialSubProfileOverTime(userId, profileType, month, pageNumber, pageSize)

Gets expense metrics for a financial sub profile over time

    This endpoint gets expense metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetExpenseMetricsFinancialSubProfileOverTimeResponse**](../Models/GetExpenseMetricsFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getFinancialProfile"></a>
# **getFinancialProfile**
> GetFinancialProfileResponse getFinancialProfile(userId, month, pageNumber, pageSize, profileType)

Get Financial Profile by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetFinancialProfileResponse**](../Models/GetFinancialProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getForecast"></a>
# **getForecast**
> GetForecastResponse getForecast(smartGoalId)

get forecast by id

    This endpoint returns the forecast if the forecast exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |

### Return type

[**GetForecastResponse**](../Models/GetForecastResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getIncomeExpenseRatio"></a>
# **getIncomeExpenseRatio**
> GetIncomeExpenseRatioResponse getIncomeExpenseRatio(userId, month, pageNumber, pageSize, profileType)

Get Income Expense Ratio by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetIncomeExpenseRatioResponse**](../Models/GetIncomeExpenseRatioResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getIncomeMetrics"></a>
# **getIncomeMetrics**
> GetIncomeMetricsResponse getIncomeMetrics(userId, month, personalFinanceCategoryPrimary, pageNumber, pageSize, profileType)

Get Income Metrics by user, month and category

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **personalFinanceCategoryPrimary** | **String**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetIncomeMetricsResponse**](../Models/GetIncomeMetricsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getIncomeMetricsFinancialSubProfileOverTime"></a>
# **getIncomeMetricsFinancialSubProfileOverTime**
> GetIncomeMetricsFinancialSubProfileOverTimeResponse getIncomeMetricsFinancialSubProfileOverTime(userId, profileType, month, pageNumber, pageSize)

Gets income metrics for a financial sub profile over time

    This endpoint gets expense metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetIncomeMetricsFinancialSubProfileOverTimeResponse**](../Models/GetIncomeMetricsFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInvestmentAccount"></a>
# **getInvestmentAccount**
> GetInvestmentAcccountResponse getInvestmentAccount(userId, investmentAccountId, profileType)

get investment account by id

    This endpoint returns the investment account if the investment account exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **investmentAccountId** | **String**| The investment account id Validations: - investment_account_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetInvestmentAcccountResponse**](../Models/GetInvestmentAcccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLiabilityAccount"></a>
# **getLiabilityAccount**
> GetLiabilityAccountResponse getLiabilityAccount(userId, liabilityAccountId, profileType)

get liability account by id

    This endpoint returns the liability account if the liability account exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **liabilityAccountId** | **String**| The liability account id Validations: - liability_account_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetLiabilityAccountResponse**](../Models/GetLiabilityAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLink"></a>
# **getLink**
> GetLinkResponse getLink(linkId, userId, profileType)

get link by id

    This endpoint returns the link if the link exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **linkId** | **String**| The link id Validations: - link_id must be greater than 0 | [default to null] |
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetLinkResponse**](../Models/GetLinkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLinks"></a>
# **getLinks**
> GetLinksResponse getLinks(userId, profileType)

get links

    This endpoint returns the links

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetLinksResponse**](../Models/GetLinksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLocationMetricsFinancialSubProfileOverTime"></a>
# **getLocationMetricsFinancialSubProfileOverTime**
> GetLocationMetricsFinancialSubProfileOverTimeResponse getLocationMetricsFinancialSubProfileOverTime(userId, profileType, month, locationCity, pageNumber, pageSize)

Gets income metrics for a financial sub profile over time

    This endpoint gets location metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **month** | **Long**|  | [optional] [default to null] |
| **locationCity** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetLocationMetricsFinancialSubProfileOverTimeResponse**](../Models/GetLocationMetricsFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMelodyFinancialContext"></a>
# **getMelodyFinancialContext**
> GetMelodyFinancialContextResponse getMelodyFinancialContext(userId, profileType)

Get Melody Financial Context

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMelodyFinancialContextResponse**](../Models/GetMelodyFinancialContextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMerchantMetricsFinancialSubProfileOverTime"></a>
# **getMerchantMetricsFinancialSubProfileOverTime**
> GetMerchantMetricsFinancialSubProfileOverTimeResponse getMerchantMetricsFinancialSubProfileOverTime(userId, profileType, month, merchantName, pageNumber, pageSize)

Gets merchant metrics for a financial sub profile over time

    This endpoint gets merchant metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **month** | **Long**|  | [optional] [default to null] |
| **merchantName** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetMerchantMetricsFinancialSubProfileOverTimeResponse**](../Models/GetMerchantMetricsFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMerchantMonthlyExpenditure"></a>
# **getMerchantMonthlyExpenditure**
> GetMerchantMonthlyExpenditureResponse getMerchantMonthlyExpenditure(userId, month, merchantName, pageNumber, pageSize, profileType)

Get Merchant Monthly Expenditure by user, month and merchant name

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **merchantName** | **String**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMerchantMonthlyExpenditureResponse**](../Models/GetMerchantMonthlyExpenditureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMilestone"></a>
# **getMilestone**
> GetMilestoneResponse getMilestone(milestoneId)

get milestone by id

    This endpoint returns the milestone if the milestone exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **milestoneId** | **String**| The milestone id Validations: - milestone_id must be greater than 0 | [default to null] |

### Return type

[**GetMilestoneResponse**](../Models/GetMilestoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMilestones"></a>
# **getMilestones**
> GetMilestonesBySmartGoalIdResponse getMilestones(smartGoalId)

get milestones by smart goal id

    This endpoint returns the milestones if the smart goal exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |

### Return type

[**GetMilestonesBySmartGoalIdResponse**](../Models/GetMilestonesBySmartGoalIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlyBalance"></a>
# **getMonthlyBalance**
> GetMonthlyBalanceResponse getMonthlyBalance(userId, month, pageNumber, pageSize, profileType)

Get Monthly Balance by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**| optional | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlyBalanceResponse**](../Models/GetMonthlyBalanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlyExpenditure"></a>
# **getMonthlyExpenditure**
> GetMonthlyExpenditureResponse getMonthlyExpenditure(userId, month, pageNumber, pageSize, profileType)

Get Monthly Expenditure by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlyExpenditureResponse**](../Models/GetMonthlyExpenditureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlyIncome"></a>
# **getMonthlyIncome**
> GetMonthlyIncomeResponse getMonthlyIncome(userId, month, pageNumber, pageSize, profileType)

Get Monthly Income by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlyIncomeResponse**](../Models/GetMonthlyIncomeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlySavings"></a>
# **getMonthlySavings**
> GetMonthlySavingsResponse getMonthlySavings(userId, month, pageNumber, pageSize, profileType)

Get Monthly Savings by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlySavingsResponse**](../Models/GetMonthlySavingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlyTotalQuantityBySecurityAndUser"></a>
# **getMonthlyTotalQuantityBySecurityAndUser**
> GetMonthlyTotalQuantityBySecurityAndUserResponse getMonthlyTotalQuantityBySecurityAndUser(userId, month, securityId, pageNumber, pageSize, profileType)

Get Monthly Total Quantity of Security by user, month and security

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **securityId** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlyTotalQuantityBySecurityAndUserResponse**](../Models/GetMonthlyTotalQuantityBySecurityAndUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMonthlyTransactionCount"></a>
# **getMonthlyTransactionCount**
> GetMonthlyTransactionCountResponse getMonthlyTransactionCount(userId, month, pageNumber, pageSize, profileType)

Get Monthly Transaction Count by user and month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMonthlyTransactionCountResponse**](../Models/GetMonthlyTransactionCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMortageAccount"></a>
# **getMortageAccount**
> GetMortgageAccountResponse getMortageAccount(userId, mortgageAccountId, profileType)

get mortgage account by id

    This endpoint returns the mortgage account if the mortgage account exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **mortgageAccountId** | **String**| The mortage account id Validations: - mortage_account_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetMortgageAccountResponse**](../Models/GetMortgageAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNoteFromSmartGoal"></a>
# **getNoteFromSmartGoal**
> GetNoteFromSmartGoalResponse getNoteFromSmartGoal(noteId)

gets a note from a smart goal

    This endpoint gets a note from a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **noteId** | **String**| The note id Validations: - note_id must be greater than 0 | [default to null] |

### Return type

[**GetNoteFromSmartGoalResponse**](../Models/GetNoteFromSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNoteFromTransaction"></a>
# **getNoteFromTransaction**
> GetNoteFromTransactionResponse getNoteFromTransaction(transactionId, noteId)

gets a note from a transaction

    This endpoint gets a note from a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |
| **noteId** | **String**| The note id Validations: - note_id must be greater than 0 | [default to null] |

### Return type

[**GetNoteFromTransactionResponse**](../Models/GetNoteFromTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNotesFromFinancialUserProfile"></a>
# **getNotesFromFinancialUserProfile**
> GetNotesFromFinancialUserProfileResponse getNotesFromFinancialUserProfile(businessAccountUserId, profileType)

Gets notes from a business account

    This endpoint gets notes from a business account

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **businessAccountUserId** | **String**| The business account id Validations: - business_account_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to null] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetNotesFromFinancialUserProfileResponse**](../Models/GetNotesFromFinancialUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNotesFromSmartGoal"></a>
# **getNotesFromSmartGoal**
> GetNotesFromSmartGoalResponse getNotesFromSmartGoal(smartGoalId)

gets notes from a smart goal

    This endpoint gets notes from a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **smartGoalId** | **String**| The smart goal id Validations: - smart_goal_id must be greater than 0 | [default to null] |

### Return type

[**GetNotesFromSmartGoalResponse**](../Models/GetNotesFromSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPaymentChannelFinancialSubProfileOverTime"></a>
# **getPaymentChannelFinancialSubProfileOverTime**
> GetPaymentChannelFinancialSubProfileOverTimeResponse getPaymentChannelFinancialSubProfileOverTime(userId, profileType, month, paymentChannel, pageNumber, pageSize)

Gets payment metrics for a financial sub profile over time

    This endpoint gets payment metrics for a financial sub profile over time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **month** | **Long**|  | [optional] [default to null] |
| **paymentChannel** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |

### Return type

[**GetPaymentChannelFinancialSubProfileOverTimeResponse**](../Models/GetPaymentChannelFinancialSubProfileOverTimeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPaymentChannelMonthlyExpenditure"></a>
# **getPaymentChannelMonthlyExpenditure**
> GetPaymentChannelMonthlyExpenditureResponse getPaymentChannelMonthlyExpenditure(userId, month, paymentChannel, pageNumber, pageSize, profileType)

Get Payment Channel Monthly Expenditure by user, month, and payment channel

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **paymentChannel** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetPaymentChannelMonthlyExpenditureResponse**](../Models/GetPaymentChannelMonthlyExpenditureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPocket"></a>
# **getPocket**
> GetPocketResponse getPocket(pocketId)

get a pocket

    This endpoint returns the pocket if the pocket exists example: /api/v1/pocket/xxxxxxx

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pocketId** | **String**| The pocket account id Validations: - pocket_account_id must be greater than 0 | [default to null] |

### Return type

[**GetPocketResponse**](../Models/GetPocketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRecurringTransaction"></a>
# **getRecurringTransaction**
> GetSingleRecurringTransactionResponse getRecurringTransaction(transactionId)

lists a set of transactions against a given account of interest

    This endpoint gets a specific transaction based on the transaction id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| the transaction of interest we aim to obtain | [default to null] |

### Return type

[**GetSingleRecurringTransactionResponse**](../Models/GetSingleRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getRecurringTransactionsForUser"></a>
# **getRecurringTransactionsForUser**
> GetRecurringTransactionsForUserResponse getRecurringTransactionsForUser(userId, profileType, pageNumber, pageSize)

get recurring transactions

    This endpoint returns the recurring transactions for a given user id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**|  | [optional] [default to null] |

### Return type

[**GetRecurringTransactionsForUserResponse**](../Models/GetRecurringTransactionsForUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSmartGoalsByPocketId"></a>
# **getSmartGoalsByPocketId**
> GetSmartGoalsByPocketIdResponse getSmartGoalsByPocketId(pocketId)

get smart goals by pocket id

    This endpoint returns the smart goals if the pocket exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pocketId** | **String**| The pocket account id Validations: - pocket_account_id must be greater than 0 | [default to null] |

### Return type

[**GetSmartGoalsByPocketIdResponse**](../Models/GetSmartGoalsByPocketIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSplitTransaction"></a>
# **getSplitTransaction**
> GetSplitTransactionResponse getSplitTransaction(transactionId)

gets a split transaction

    This endpoint gets a split transaction that was split previously into multiple transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |

### Return type

[**GetSplitTransactionResponse**](../Models/GetSplitTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getStudentLoanAccount"></a>
# **getStudentLoanAccount**
> GetStudentLoanAccountResponse getStudentLoanAccount(userId, studentLoanAccountId, profileType)

get student loan account by id

    This endpoint returns the student loan account if the student loan account exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **studentLoanAccountId** | **String**| The student loan account id Validations: - student_loan_account_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetStudentLoanAccountResponse**](../Models/GetStudentLoanAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTotalInvestmentBySecurity"></a>
# **getTotalInvestmentBySecurity**
> GetTotalInvestmentBySecurityResponse getTotalInvestmentBySecurity(userId, securityId, pageNumber, pageSize, profileType)

Get Total Investment by user and security

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **securityId** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetTotalInvestmentBySecurityResponse**](../Models/GetTotalInvestmentBySecurityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransaction"></a>
# **getTransaction**
> GetTransactionResponse getTransaction(transactionId)

lists a set of transactions against a given account of interest

    This endpoint gets a specific transaction based on the transaction id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| the transaction of interest we aim to obtain | [default to null] |

### Return type

[**GetTransactionResponse**](../Models/GetTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactions"></a>
# **getTransactions**
> GetTransactionsForBankAccountResponse getTransactions(userId, plaidAccountId, pageNumber, pageSize, profileType)

get transactions tied to a bank account and account id

    This endpoint returns the transactions for a given user id and a business account id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **plaidAccountId** | **String**|  | [default to null] |
| **pageNumber** | **String**|  | [default to null] |
| **pageSize** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetTransactionsForBankAccountResponse**](../Models/GetTransactionsForBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactions1"></a>
# **getTransactions1**
> GetTransactionsResponse getTransactions1(userId, pageNumber, pageSize, profileType)

get transactions

    This endpoint returns the transactions for a given user id

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **pageNumber** | **String**|  | [default to null] |
| **pageSize** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetTransactionsResponse**](../Models/GetTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactionsByTime"></a>
# **getTransactionsByTime**
> GetTransactionsBetweenTimeRangesResponse getTransactionsByTime(userId, plaidAccountId, profileType, startDate, endDate, page, limit, financialAccountType)

get transactions by time

    This endpoint returns the transactions for a given user id and time example: /api/v1/users/{user_id}/accounts/{plaid_account_id}/transactions?start_time&#x3D;{start_time}&amp;end_time&#x3D;{end_time}

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The user id Validations: - user_id must be greater than 0 | [default to null] |
| **plaidAccountId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **startDate** | **Date**| start date | [optional] [default to null] |
| **endDate** | **Date**| end date | [optional] [default to null] |
| **page** | **Long**| Current page number | [optional] [default to null] |
| **limit** | **Long**| Number of transactions per page | [optional] [default to null] |
| **financialAccountType** | **String**|  | [optional] [default to FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED] [enum: FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED, FINANCIAL_ACCOUNT_TYPE_BANK, FINANCIAL_ACCOUNT_TYPE_INVESTMENT, FINANCIAL_ACCOUNT_TYPE_CREDIT, FINANCIAL_ACCOUNT_TYPE_MORTGAGE, FINANCIAL_ACCOUNT_TYPE_STUDENT_LOAN] |

### Return type

[**GetTransactionsBetweenTimeRangesResponse**](../Models/GetTransactionsBetweenTimeRangesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactionsForPastMonth"></a>
# **getTransactionsForPastMonth**
> GetTransactionsForPastMonthResponse getTransactionsForPastMonth(userId, plaidAccountId, profileType, page, limit, financialAccountType)

Get transactions for the past month

    This endpoint returns the transactions for a given user id over the past month example: /api/v1/users/{user_id}/accounts/{plaid_account_id}/transactions/month

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **plaidAccountId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **page** | **Long**| Current page number | [optional] [default to null] |
| **limit** | **Long**| Number of transactions per page | [optional] [default to null] |
| **financialAccountType** | **String**|  | [optional] [default to FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED] [enum: FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED, FINANCIAL_ACCOUNT_TYPE_BANK, FINANCIAL_ACCOUNT_TYPE_INVESTMENT, FINANCIAL_ACCOUNT_TYPE_CREDIT, FINANCIAL_ACCOUNT_TYPE_MORTGAGE, FINANCIAL_ACCOUNT_TYPE_STUDENT_LOAN] |

### Return type

[**GetTransactionsForPastMonthResponse**](../Models/GetTransactionsForPastMonthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactionsForPastWeek"></a>
# **getTransactionsForPastWeek**
> GetTransactionsForPastWeekResponse getTransactionsForPastWeek(userId, plaidAccountId, profileType, page, limit, financialAccountType)

get transactions for the past week

    This endpoint returns the transactions for a given user id and time example: /api/v1/users/{user_id}/accounts/{plaid_account_id}/transactions/week

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **plaidAccountId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **page** | **Long**| Current page number | [optional] [default to null] |
| **limit** | **Long**| Number of transactions per page | [optional] [default to null] |
| **financialAccountType** | **String**|  | [optional] [default to FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED] [enum: FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED, FINANCIAL_ACCOUNT_TYPE_BANK, FINANCIAL_ACCOUNT_TYPE_INVESTMENT, FINANCIAL_ACCOUNT_TYPE_CREDIT, FINANCIAL_ACCOUNT_TYPE_MORTGAGE, FINANCIAL_ACCOUNT_TYPE_STUDENT_LOAN] |

### Return type

[**GetTransactionsForPastWeekResponse**](../Models/GetTransactionsForPastWeekResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserAccountBalanceHistory"></a>
# **getUserAccountBalanceHistory**
> GetUserAccountBalanceHistoryResponse getUserAccountBalanceHistory(userId, pageNumber, pageSize, profileType)

Returns the account balance history for a user

    This API could accept user_id as input parameters and return the account balance history for that user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| User ID | [default to null] |
| **pageNumber** | **String**|  | [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserAccountBalanceHistoryResponse**](../Models/GetUserAccountBalanceHistoryResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserCategoryMonthlyExpenditure"></a>
# **getUserCategoryMonthlyExpenditure**
> GetUserCategoryMonthlyExpenditureResponse getUserCategoryMonthlyExpenditure(userId, personalFinanceCategoryPrimary, month, pageNumber, pageSize, profileType)

Returns the monthly category expenditure for a user

    This API could accept user_id as an input parameter and return the monthly category expenditure for that user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| User ID | [default to null] |
| **personalFinanceCategoryPrimary** | **String**| Personal finance category | [optional] [default to null] |
| **month** | **Long**| Month in the format of YYYYMM | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserCategoryMonthlyExpenditureResponse**](../Models/GetUserCategoryMonthlyExpenditureResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserCategoryMonthlyIncome"></a>
# **getUserCategoryMonthlyIncome**
> GetUserCategoryMonthlyIncomeResponse getUserCategoryMonthlyIncome(userId, personalFinanceCategoryPrimary, month, pageNumber, pageSize, profileType)

Get monthly income by user for a specific category

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **personalFinanceCategoryPrimary** | **String**|  | [optional] [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetUserCategoryMonthlyIncomeResponse**](../Models/GetUserCategoryMonthlyIncomeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserProfile1"></a>
# **getUserProfile1**
> GetUserProfileResponse1 getUserProfile1(userId, profileType, bypassCache)

Gets a user profile

    Queries and obtains a user profile based on the provided parameters

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The account ID associated with the user. NOTE: This user_id is the simfiny backend platform wide user id Validations: - user_id must be greater than 0 | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **bypassCache** | **Boolean**| bypass_cache is a flag to bypass the cache and fetch the data from the source this is very valuable when we want to get the latest data from the source | [optional] [default to null] |

### Return type

[**GetUserProfileResponse1**](../Models/GetUserProfileResponse1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="healthCheck2"></a>
# **healthCheck2**
> HealthCheckResponse healthCheck2()

health check

    This endpoint performs a healc check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**HealthCheckResponse**](../Models/HealthCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="initiatePlaidSetup"></a>
# **initiatePlaidSetup**
> PlaidInitiateTokenExchangeResponse initiatePlaidSetup(PlaidInitiateTokenExchangeRequest)

initiate plaid setup

    This endpoint initiates a plaid link token creation

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PlaidInitiateTokenExchangeRequest** | [**PlaidInitiateTokenExchangeRequest**](../Models/PlaidInitiateTokenExchangeRequest.md)|  | |

### Return type

[**PlaidInitiateTokenExchangeResponse**](../Models/PlaidInitiateTokenExchangeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="initiatePlaidTokenUpdate"></a>
# **initiatePlaidTokenUpdate**
> PlaidInitiateTokenUpdateResponse initiatePlaidTokenUpdate(PlaidInitiateTokenUpdateRequest)

initiate plaid link token update

    This endpoint initiates a plaid link token update

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PlaidInitiateTokenUpdateRequest** | [**PlaidInitiateTokenUpdateRequest**](../Models/PlaidInitiateTokenUpdateRequest.md)|  | |

### Return type

[**PlaidInitiateTokenUpdateResponse**](../Models/PlaidInitiateTokenUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="listRecurringTransactionNotes"></a>
# **listRecurringTransactionNotes**
> ListRecurringTransactionNotesResponse listRecurringTransactionNotes(transactionId)

lists notes from a transaction

    This endpoint lists notes from a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |

### Return type

[**ListRecurringTransactionNotesResponse**](../Models/ListRecurringTransactionNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRecurringTransactionsForUserAndAccount"></a>
# **listRecurringTransactionsForUserAndAccount**
> ListRecurringTransactionsForUserAndAccountResponse listRecurringTransactionsForUserAndAccount(accountId, userId, profileType, pageNumber, pageSize, financialAccountType)

lists a set of transactions against a given account of interest

    This endpoint gets a list of transactions against a given account of interest

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**|  | [default to null] |
| **userId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**|  | [optional] [default to null] |
| **financialAccountType** | **String**|  | [optional] [default to FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED] [enum: FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED, FINANCIAL_ACCOUNT_TYPE_BANK, FINANCIAL_ACCOUNT_TYPE_INVESTMENT, FINANCIAL_ACCOUNT_TYPE_CREDIT, FINANCIAL_ACCOUNT_TYPE_MORTGAGE, FINANCIAL_ACCOUNT_TYPE_STUDENT_LOAN] |

### Return type

[**ListRecurringTransactionsForUserAndAccountResponse**](../Models/ListRecurringTransactionsForUserAndAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listTransactionNotes"></a>
# **listTransactionNotes**
> ListTransactionNotesResponse listTransactionNotes(transactionId)

lists notes from a transaction

    This endpoint lists notes from a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transactionId** | **String**| The transaction id Validations: - transaction_id must be greater than 0 | [default to null] |

### Return type

[**ListTransactionNotesResponse**](../Models/ListTransactionNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listTransactions"></a>
# **listTransactions**
> ListTransactionsResponse listTransactions(accountId, userId, profileType, pageNumber, pageSize, financialAccountType)

lists a set of transactions against a given account of interest

    This endpoint gets a list of transactions against a given account of interest

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**|  | [default to null] |
| **userId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**|  | [optional] [default to null] |
| **financialAccountType** | **String**|  | [optional] [default to FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED] [enum: FINANCIAL_ACCOUNT_TYPE_UNSPECIFIED, FINANCIAL_ACCOUNT_TYPE_BANK, FINANCIAL_ACCOUNT_TYPE_INVESTMENT, FINANCIAL_ACCOUNT_TYPE_CREDIT, FINANCIAL_ACCOUNT_TYPE_MORTGAGE, FINANCIAL_ACCOUNT_TYPE_STUDENT_LOAN] |

### Return type

[**ListTransactionsResponse**](../Models/ListTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listTransactions1"></a>
# **listTransactions1**
> ListTransactionsAcrossAllAccountsResponse listTransactions1(userId, profileType, pageNumber, pageSize)

lists a set of transactions across all connected accounts

    This endpoint gets a list of transactions across all connected accounts

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **profileType** | **String**|  | [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**|  | [optional] [default to null] |

### Return type

[**ListTransactionsAcrossAllAccountsResponse**](../Models/ListTransactionsAcrossAllAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="pollAsyncTaskExecutionStatus"></a>
# **pollAsyncTaskExecutionStatus**
> PollAsyncTaskExecutionStatusResponse pollAsyncTaskExecutionStatus(workflowId, runId)

polls the status of an async task

    This endpoint polls the status of an async task

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **workflowId** | **String**| The task id Validations: - workflow id cannot be empty | [default to null] |
| **runId** | **String**| The run id Validations: - run id cannot be empty | [default to null] |

### Return type

[**PollAsyncTaskExecutionStatusResponse**](../Models/PollAsyncTaskExecutionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readynessCheck2"></a>
# **readynessCheck2**
> ReadynessCheckResponse readynessCheck2()

readyness check

    This endpoint performs a readiness check on the service

### Parameters
This endpoint does not need any parameter.

### Return type

[**ReadynessCheckResponse**](../Models/ReadynessCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="searchTransactions"></a>
# **searchTransactions**
> SearchTransactionsResponse searchTransactions(SearchTransactionsRequest)

searches transactions

    This endpoint searches transactions based on a search query and returns a list of transactions that match the search query

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **SearchTransactionsRequest** | [**SearchTransactionsRequest**](../Models/SearchTransactionsRequest.md)|  | |

### Return type

[**SearchTransactionsResponse**](../Models/SearchTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="splitTransaction"></a>
# **splitTransaction**
> SplitTransactionResponse splitTransaction(SplitTransactionRequest)

splits a transaction

    This endpoint splits a transaction into multiple transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **SplitTransactionRequest** | [**SplitTransactionRequest**](../Models/SplitTransactionRequest.md)|  | |

### Return type

[**SplitTransactionResponse**](../Models/SplitTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="transactionAggregates"></a>
# **transactionAggregates**
> GetTransactionAggregatesResponse transactionAggregates(userId, month, personalFinanceCategoryPrimary, locationCity, paymentChannel, merchantName, pageNumber, pageSize, profileType)

Returns the aggregated transactions for a user and month

    This API could accept user_id and month as input parameters and return the aggregated transactions for that user and month.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**|  | [default to null] |
| **month** | **Long**|  | [optional] [default to null] |
| **personalFinanceCategoryPrimary** | **String**|  | [optional] [default to null] |
| **locationCity** | **String**|  | [optional] [default to null] |
| **paymentChannel** | **String**|  | [optional] [default to null] |
| **merchantName** | **String**|  | [optional] [default to null] |
| **pageNumber** | **String**|  | [optional] [default to null] |
| **pageSize** | **String**| Number of items to return per page. | [optional] [default to null] |
| **profileType** | **String**|  | [optional] [default to FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED] [enum: FINANCIAL_USER_PROFILE_TYPE_UNSPECIFIED, FINANCIAL_USER_PROFILE_TYPE_USER, FINANCIAL_USER_PROFILE_TYPE_BUSINESS] |

### Return type

[**GetTransactionAggregatesResponse**](../Models/GetTransactionAggregatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="triggerSync"></a>
# **triggerSync**
> TriggerSyncResponse triggerSync(TriggerSyncRequest)

Triggers a sync

    This endpoint triggers a sync

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **TriggerSyncRequest** | [**TriggerSyncRequest**](../Models/TriggerSyncRequest.md)|  | |

### Return type

[**TriggerSyncResponse**](../Models/TriggerSyncResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="unsplitTransactions"></a>
# **unsplitTransactions**
> UnSplitTransactionsResponse unsplitTransactions(UnSplitTransactionsRequest)

unsplit a transaction

    This endpoint unsplit a transaction that was split previously into multiple transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UnSplitTransactionsRequest** | [**UnSplitTransactionsRequest**](../Models/UnSplitTransactionsRequest.md)|  | |

### Return type

[**UnSplitTransactionsResponse**](../Models/UnSplitTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateBankAccount"></a>
# **updateBankAccount**
> UpdateBankAccountResponse updateBankAccount(UpdateBankAccountRequest)

update a bank account for a given user profile

    This endpoint updates a bank account if the user record id and bank account id exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateBankAccountRequest** | [**UpdateBankAccountRequest**](../Models/UpdateBankAccountRequest.md)|  | |

### Return type

[**UpdateBankAccountResponse**](../Models/UpdateBankAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateNoteToRecurringTransaction"></a>
# **updateNoteToRecurringTransaction**
> UpdateNoteToRecurringTransactionResponse updateNoteToRecurringTransaction(UpdateNoteToRecurringTransactionRequest)

Updates a note to a transaction

    This endpoint Updates a note to a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateNoteToRecurringTransactionRequest** | [**UpdateNoteToRecurringTransactionRequest**](../Models/UpdateNoteToRecurringTransactionRequest.md)|  | |

### Return type

[**UpdateNoteToRecurringTransactionResponse**](../Models/UpdateNoteToRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateNoteToSmartGoal"></a>
# **updateNoteToSmartGoal**
> UpdateNoteToSmartGoalResponse updateNoteToSmartGoal(UpdateNoteToSmartGoalRequest)

updates a note to a smart goal

    This endpoint updates a note to a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateNoteToSmartGoalRequest** | [**UpdateNoteToSmartGoalRequest**](../Models/UpdateNoteToSmartGoalRequest.md)|  | |

### Return type

[**UpdateNoteToSmartGoalResponse**](../Models/UpdateNoteToSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateNoteToTransaction"></a>
# **updateNoteToTransaction**
> UpdateNoteToTransactionResponse updateNoteToTransaction(UpdateNoteToTransactionRequest)

Updates a note to a transaction

    This endpoint Updates a note to a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateNoteToTransactionRequest** | [**UpdateNoteToTransactionRequest**](../Models/UpdateNoteToTransactionRequest.md)|  | |

### Return type

[**UpdateNoteToTransactionResponse**](../Models/UpdateNoteToTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updatePocket"></a>
# **updatePocket**
> UpdatePocketResponse updatePocket(UpdatePocketRequest)

updates a pocket

    This endpoint updates a pocket

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdatePocketRequest** | [**UpdatePocketRequest**](../Models/UpdatePocketRequest.md)|  | |

### Return type

[**UpdatePocketResponse**](../Models/UpdatePocketResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateSingleTransaction"></a>
# **updateSingleTransaction**
> UpdateSingleTransactionResponse updateSingleTransaction(UpdateSingleTransactionRequest)

update a transaction

    This endpoint updates a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateSingleTransactionRequest** | [**UpdateSingleTransactionRequest**](../Models/UpdateSingleTransactionRequest.md)|  | |

### Return type

[**UpdateSingleTransactionResponse**](../Models/UpdateSingleTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateSmartGoal"></a>
# **updateSmartGoal**
> UpdateSmartGoalResponse updateSmartGoal(UpdateSmartGoalRequest)

update a smart goal

    This endpoint updates a smart goal

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateSmartGoalRequest** | [**UpdateSmartGoalRequest**](../Models/UpdateSmartGoalRequest.md)|  | |

### Return type

[**UpdateSmartGoalResponse**](../Models/UpdateSmartGoalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateTransaction"></a>
# **updateTransaction**
> UpdateRecurringTransactionResponse updateTransaction(UpdateRecurringTransactionRequest)

update a transaction

    This endpoint updates a transaction

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateRecurringTransactionRequest** | [**UpdateRecurringTransactionRequest**](../Models/UpdateRecurringTransactionRequest.md)|  | |

### Return type

[**UpdateRecurringTransactionResponse**](../Models/UpdateRecurringTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserProfile"></a>
# **updateUserProfile**
> UpdateUserProfileResponse updateUserProfile(UpdateUserProfileRequest)

update a user profile

    This endpoint performs an updates operation on a user profile based on the provided parameters

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateUserProfileRequest** | [**UpdateUserProfileRequest**](../Models/UpdateUserProfileRequest.md)|  | |

### Return type

[**UpdateUserProfileResponse**](../Models/UpdateUserProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updatesBudget"></a>
# **updatesBudget**
> UpdateBudgetResponse updatesBudget(UpdateBudgetRequest)

updates a budget

    This endpoint updates a budget

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateBudgetRequest** | [**UpdateBudgetRequest**](../Models/UpdateBudgetRequest.md)|  | |

### Return type

[**UpdateBudgetResponse**](../Models/UpdateBudgetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updatesMilestone"></a>
# **updatesMilestone**
> UpdateMilestoneResponse updatesMilestone(UpdateMilestoneRequest)

updates a milestone

    This endpoint updates a milestone

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateMilestoneRequest** | [**UpdateMilestoneRequest**](../Models/UpdateMilestoneRequest.md)|  | |

### Return type

[**UpdateMilestoneResponse**](../Models/UpdateMilestoneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

