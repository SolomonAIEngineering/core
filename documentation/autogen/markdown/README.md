# Documentation for User Service

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://user-service.platform.svc.cluster.local:9896*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AccountingServiceApi* | [**createPayrollProfile**](Apis/AccountingServiceApi.md#createpayrollprofile) | **POST** /accounting-microservice/api/v1/profile | Create Payroll Profile |
*AccountingServiceApi* | [**deleteProfile**](Apis/AccountingServiceApi.md#deleteprofile) | **DELETE** /accounting-microservice/api/v1/profile/{authZeroUserId} | Delete Payroll Profile |
*AccountingServiceApi* | [**exchangePublicLinkTokenForAccountTokenResponse**](Apis/AccountingServiceApi.md#exchangepubliclinktokenforaccounttokenresponse) | **POST** /accounting-microservice/api/v1/merge/exchange-token | Exchange Link Token |
*AccountingServiceApi* | [**getLinkToken**](Apis/AccountingServiceApi.md#getlinktoken) | **POST** /accounting-microservice/api/v1/merge/initiate-token-exchange | Get Link Token |
*AccountingServiceApi* | [**healthCheck3**](Apis/AccountingServiceApi.md#healthcheck3) | **GET** /accounting-microservice/api/v1/health | Health Check |
*AccountingServiceApi* | [**readAccountingProfileResponse**](Apis/AccountingServiceApi.md#readaccountingprofileresponse) | **GET** /accounting-microservice/api/v1/profile/{authZeroUserId} | Get Business Account Profile |
*AccountingServiceApi* | [**readBalanceSheets**](Apis/AccountingServiceApi.md#readbalancesheets) | **POST** /accounting-microservice/api/v1/balance-sheets | Gets Balance Sheets |
*AccountingServiceApi* | [**readBusinessChartOfAccounts**](Apis/AccountingServiceApi.md#readbusinesschartofaccounts) | **POST** /accounting-microservice/api/v1/chart-of-accounts | Gets Chart of Accounts |
*AccountingServiceApi* | [**readBusinessTransactions**](Apis/AccountingServiceApi.md#readbusinesstransactions) | **POST** /accounting-microservice/api/v1/business-transactions | Gets Business Transactions |
*AccountingServiceApi* | [**readCashFlowStatements**](Apis/AccountingServiceApi.md#readcashflowstatements) | **POST** /accounting-microservice/api/v1/cashflow-statements | Gets Cashfloe Sheets |
*AccountingServiceApi* | [**readIncomeStatements**](Apis/AccountingServiceApi.md#readincomestatements) | **POST** /accounting-microservice/api/v1/income-statements | Gets Cashfloe Sheets |
*AccountingServiceApi* | [**readinessCheck**](Apis/AccountingServiceApi.md#readinesscheck) | **GET** /accounting-microservice/api/v1/ready | Readiness Check |
*AccountingServiceApi* | [**updatePayrollProfile**](Apis/AccountingServiceApi.md#updatepayrollprofile) | **PUT** /accounting-microservice/api/v1/profile | Update Payroll Profile |
| *FinancialServiceApi* | [**addDefaultPocketsToBankAccount**](Apis/FinancialServiceApi.md#adddefaultpocketstobankaccount) | **POST** /financial-microservice/api/v1/pocket/bank-account | adds a default set of pockets to a specific bank account of interest |
*FinancialServiceApi* | [**addNoteToFinancialUserProfile**](Apis/FinancialServiceApi.md#addnotetofinancialuserprofile) | **POST** /financial-microservice/api/v1/financial-profile/business/note | Adds a note to a business account |
*FinancialServiceApi* | [**addNoteToRecurringTransaction**](Apis/FinancialServiceApi.md#addnotetorecurringtransaction) | **POST** /financial-microservice/api/v1/transactions/recurring/note | adds a note to a transaction |
*FinancialServiceApi* | [**addNoteToSmartGoal**](Apis/FinancialServiceApi.md#addnotetosmartgoal) | **POST** /financial-microservice/api/v1/smart-goal/note | adds a note to a smart goal |
*FinancialServiceApi* | [**addNoteToTransaction**](Apis/FinancialServiceApi.md#addnotetotransaction) | **POST** /financial-microservice/api/v1/transactions/transaction/note | adds a note to a transaction |
*FinancialServiceApi* | [**askCopilotQuestion**](Apis/FinancialServiceApi.md#askcopilotquestion) | **POST** /financial-microservice/api/v1/copilot/quota/question | Ask a question to copilot |
*FinancialServiceApi* | [**bulkUpdateRecurringTransaction**](Apis/FinancialServiceApi.md#bulkupdaterecurringtransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring/bulk | update a transaction |
*FinancialServiceApi* | [**bulkUpdateTransaction**](Apis/FinancialServiceApi.md#bulkupdatetransaction) | **PUT** /financial-microservice/api/v1/transactions/transaction/bulk | update a transaction |
*FinancialServiceApi* | [**checkIfQuotaExceeded**](Apis/FinancialServiceApi.md#checkifquotaexceeded) | **GET** /financial-microservice/api/v1/copilot/quota/exceeded/{userId} | Checks if the question quota has been exceeded |
*FinancialServiceApi* | [**createBankAccount**](Apis/FinancialServiceApi.md#createbankaccount) | **POST** /financial-microservice/api/v1/bank-account/profile | create a bank account for a given user profile |
*FinancialServiceApi* | [**createBudget**](Apis/FinancialServiceApi.md#createbudget) | **POST** /financial-microservice/api/v1/budget | create a budget |
*FinancialServiceApi* | [**createLink**](Apis/FinancialServiceApi.md#createlink) | **POST** /financial-microservice/api/v1/link | create link |
*FinancialServiceApi* | [**createMilestone**](Apis/FinancialServiceApi.md#createmilestone) | **POST** /financial-microservice/api/v1/milestone | create a milestone |
*FinancialServiceApi* | [**createSmartGoal**](Apis/FinancialServiceApi.md#createsmartgoal) | **POST** /financial-microservice/api/v1/smart-goal | create a smart goal |
*FinancialServiceApi* | [**createSubscription**](Apis/FinancialServiceApi.md#createsubscription) | **POST** /financial-microservice/api/v1/stripe/subscription | Creates a new subscription for a given customer against stripe |
*FinancialServiceApi* | [**createUserProfile1**](Apis/FinancialServiceApi.md#createuserprofile1) | **POST** /financial-microservice/api/v1/profile | create a user profile |
*FinancialServiceApi* | [**deleteBudget**](Apis/FinancialServiceApi.md#deletebudget) | **DELETE** /financial-microservice/api/v1/budget/{budgetId} | delete a budget |
*FinancialServiceApi* | [**deleteLink**](Apis/FinancialServiceApi.md#deletelink) | **DELETE** /financial-microservice/api/v1/link/{linkId}/user/{userId} | delete link by id |
*FinancialServiceApi* | [**deleteMilestone**](Apis/FinancialServiceApi.md#deletemilestone) | **DELETE** /financial-microservice/api/v1/milestone/{milestoneId} | delete a milestone |
*FinancialServiceApi* | [**deleteNoteFromRecurringTransaction**](Apis/FinancialServiceApi.md#deletenotefromrecurringtransaction) | **DELETE** /financial-microservice/api/v1/transactions/recurring/{transactionId}/note/{noteId} | deletes a note from a transaction |
*FinancialServiceApi* | [**deleteNoteFromSmartGoal**](Apis/FinancialServiceApi.md#deletenotefromsmartgoal) | **DELETE** /financial-microservice/api/v1/smart-goal/note/{noteId} | deletes a note from a smart goal |
*FinancialServiceApi* | [**deleteNoteFromTransaction**](Apis/FinancialServiceApi.md#deletenotefromtransaction) | **DELETE** /financial-microservice/api/v1/transactions/transaction/{transactionId}/note/{noteId} | deletes a note from a transaction |
*FinancialServiceApi* | [**deletePocket**](Apis/FinancialServiceApi.md#deletepocket) | **DELETE** /financial-microservice/api/v1/pocket/{pocketId} | deletes a pocket |
*FinancialServiceApi* | [**deleteSmartGoal**](Apis/FinancialServiceApi.md#deletesmartgoal) | **DELETE** /financial-microservice/api/v1/smart-goal/{smartGoalId} | delete a smart goal |
*FinancialServiceApi* | [**deleteTransaction**](Apis/FinancialServiceApi.md#deletetransaction) | **DELETE** /financial-microservice/api/v1/transactions/recurring/{transactionId} | deletes a transaction by id |
*FinancialServiceApi* | [**deleteTransaction1**](Apis/FinancialServiceApi.md#deletetransaction1) | **DELETE** /financial-microservice/api/v1/transactions/transaction/{transactionId} | deletes a transaction by id |
*FinancialServiceApi* | [**deleteUserProfile1**](Apis/FinancialServiceApi.md#deleteuserprofile1) | **DELETE** /financial-microservice/api/v1/bank-account/{bankAccountId} | deletes a bank account for a given user profile |
*FinancialServiceApi* | [**deleteUserProfile2**](Apis/FinancialServiceApi.md#deleteuserprofile2) | **DELETE** /financial-microservice/api/v1/profile/{userId} | deletes a user profile |
*FinancialServiceApi* | [**exchangePlaidToken**](Apis/FinancialServiceApi.md#exchangeplaidtoken) | **POST** /financial-microservice/api/v1/plaid/exchange-token | exchange plaid token |
*FinancialServiceApi* | [**getAccountBalance**](Apis/FinancialServiceApi.md#getaccountbalance) | **GET** /financial-microservice/api/v1/historical-account-balance/user/{userId}/plaid-account-id/{plaidAccountId} | gets account balance of an account |
*FinancialServiceApi* | [**getAccountBalanceHistory**](Apis/FinancialServiceApi.md#getaccountbalancehistory) | **GET** /financial-microservice/api/v1/analytics/balance-history/account/{plaidAccountId}/pagenumber/{pageNumber}/pagesize/{pageSize} | Returns the account balance history for an account |
*FinancialServiceApi* | [**getAllBudgets**](Apis/FinancialServiceApi.md#getallbudgets) | **GET** /financial-microservice/api/v1/budget | get all budgets |
*FinancialServiceApi* | [**getBankAccount**](Apis/FinancialServiceApi.md#getbankaccount) | **GET** /financial-microservice/api/v1/bank-account/{bankAccountId} | get a bank account for a given user profile |
*FinancialServiceApi* | [**getBudget**](Apis/FinancialServiceApi.md#getbudget) | **GET** /financial-microservice/api/v1/budget/{budgetId} | get budget by id |
*FinancialServiceApi* | [**getCategoryMonthlyTransactionCount**](Apis/FinancialServiceApi.md#getcategorymonthlytransactioncount) | **GET** /financial-microservice/api/v1/analytics/category-monthly-transaction-count/user/{userId} | Get monthly transaction count by user, month, and category |
*FinancialServiceApi* | [**getDebtToIncomeRatio**](Apis/FinancialServiceApi.md#getdebttoincomeratio) | **GET** /financial-microservice/api/v1/analytics/debt-to-income-ratio/user/{userId} | Get Debt-to-Income ratio by user and month |
*FinancialServiceApi* | [**getExpenseMetrics**](Apis/FinancialServiceApi.md#getexpensemetrics) | **GET** /financial-microservice/api/v1/analytics/expenses/user/{userId} | Get Expense Metrics by user, month and category |
*FinancialServiceApi* | [**getFinancialProfile**](Apis/FinancialServiceApi.md#getfinancialprofile) | **GET** /financial-microservice/api/v1/analytics/finance-profile/user/{userId} | Get Financial Profile by user and month |
*FinancialServiceApi* | [**getForecast**](Apis/FinancialServiceApi.md#getforecast) | **GET** /financial-microservice/api/v1/forecast/{smartGoalId} | get forecast by id |
*FinancialServiceApi* | [**getIncomeExpenseRatio**](Apis/FinancialServiceApi.md#getincomeexpenseratio) | **GET** /financial-microservice/api/v1/analytics/income-expense-ratio/user/{userId} | Get Income Expense Ratio by user and month |
*FinancialServiceApi* | [**getIncomeMetrics**](Apis/FinancialServiceApi.md#getincomemetrics) | **GET** /financial-microservice/api/v1/analytics/income/user/{userId} | Get Income Metrics by user, month and category |
*FinancialServiceApi* | [**getInvestmentAccount**](Apis/FinancialServiceApi.md#getinvestmentaccount) | **GET** /financial-microservice/api/v1/account/{userId}/investment/{investmentAccountId} | get investment account by id |
*FinancialServiceApi* | [**getLiabilityAccount**](Apis/FinancialServiceApi.md#getliabilityaccount) | **GET** /financial-microservice/api/v1/account/{userId}/liability/{liabilityAccountId} | get liability account by id |
*FinancialServiceApi* | [**getLink**](Apis/FinancialServiceApi.md#getlink) | **GET** /financial-microservice/api/v1/link/{linkId} | get link by id |
*FinancialServiceApi* | [**getLinks**](Apis/FinancialServiceApi.md#getlinks) | **GET** /financial-microservice/api/v1/links/{userId} | get links |
*FinancialServiceApi* | [**getMelodyFinancialContext**](Apis/FinancialServiceApi.md#getmelodyfinancialcontext) | **GET** /financial-microservice/api/v1/analytics/melody-financial-context/user/{userId} | Get Melody Financial Context |
*FinancialServiceApi* | [**getMerchantMonthlyExpenditure**](Apis/FinancialServiceApi.md#getmerchantmonthlyexpenditure) | **GET** /financial-microservice/api/v1/analytics/merchant-monthly-expenditure/user/{userId} | Get Merchant Monthly Expenditure by user, month and merchant name |
*FinancialServiceApi* | [**getMilestone**](Apis/FinancialServiceApi.md#getmilestone) | **GET** /financial-microservice/api/v1/milestone/{milestoneId} | get milestone by id |
*FinancialServiceApi* | [**getMilestones**](Apis/FinancialServiceApi.md#getmilestones) | **GET** /financial-microservice/api/v1/milestone/smart-goal/{smartGoalId} | get milestones by smart goal id |
*FinancialServiceApi* | [**getMonthlyBalance**](Apis/FinancialServiceApi.md#getmonthlybalance) | **GET** /financial-microservice/api/v1/analytics/monthly-balance/user/{userId} | Get Monthly Balance by user and month |
*FinancialServiceApi* | [**getMonthlyExpenditure**](Apis/FinancialServiceApi.md#getmonthlyexpenditure) | **GET** /financial-microservice/api/v1/analytics/monthly-expenditure/user/{userId} | Get Monthly Expenditure by user and month |
*FinancialServiceApi* | [**getMonthlyIncome**](Apis/FinancialServiceApi.md#getmonthlyincome) | **GET** /financial-microservice/api/v1/analytics/monthly-income/user/{userId} | Get Monthly Income by user and month |
*FinancialServiceApi* | [**getMonthlySavings**](Apis/FinancialServiceApi.md#getmonthlysavings) | **GET** /financial-microservice/api/v1/analytics/monthly-savings/{userId} | Get Monthly Savings by user and month |
*FinancialServiceApi* | [**getMonthlyTotalQuantityBySecurityAndUser**](Apis/FinancialServiceApi.md#getmonthlytotalquantitybysecurityanduser) | **GET** /financial-microservice/api/v1/analytics/monthly-total-quantity-by-security-and-user/user/{userId} | Get Monthly Total Quantity of Security by user, month and security |
*FinancialServiceApi* | [**getMonthlyTransactionCount**](Apis/FinancialServiceApi.md#getmonthlytransactioncount) | **GET** /financial-microservice/api/v1/analytics/monthly-transaction-count/user/{userId} | Get Monthly Transaction Count by user and month |
*FinancialServiceApi* | [**getMortageAccount**](Apis/FinancialServiceApi.md#getmortageaccount) | **GET** /financial-microservice/api/v1/account/{userId}/mortgage/{mortgageAccountId} | get mortgage account by id |
*FinancialServiceApi* | [**getNoteFromSmartGoal**](Apis/FinancialServiceApi.md#getnotefromsmartgoal) | **GET** /financial-microservice/api/v1/smart-goal/note/{noteId} | gets a note from a smart goal |
*FinancialServiceApi* | [**getNoteFromTransaction**](Apis/FinancialServiceApi.md#getnotefromtransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/note/{noteId} | gets a note from a transaction |
*FinancialServiceApi* | [**getNotesFromFinancialUserProfile**](Apis/FinancialServiceApi.md#getnotesfromfinancialuserprofile) | **GET** /financial-microservice/api/v1/financial-profile/business/{businessAccountUserId}/{profileType}/note | Gets notes from a business account |
*FinancialServiceApi* | [**getNotesFromSmartGoal**](Apis/FinancialServiceApi.md#getnotesfromsmartgoal) | **GET** /financial-microservice/api/v1/smart-goal/{smartGoalId}/note | gets notes from a smart goal |
*FinancialServiceApi* | [**getPaymentChannelMonthlyExpenditure**](Apis/FinancialServiceApi.md#getpaymentchannelmonthlyexpenditure) | **GET** /financial-microservice/api/v1/analytics/payment-channel-expenditure/user/{userId} | Get Payment Channel Monthly Expenditure by user, month, and payment channel |
*FinancialServiceApi* | [**getPocket**](Apis/FinancialServiceApi.md#getpocket) | **GET** /financial-microservice/api/v1/pocket/{pocketId} | get a pocket |
*FinancialServiceApi* | [**getRecurringTransaction**](Apis/FinancialServiceApi.md#getrecurringtransaction) | **GET** /financial-microservice/api/v1/transactions/recurring/{transactionId} | lists a set of transactions against a given account of interest |
*FinancialServiceApi* | [**getRecurringTransactionsForUser**](Apis/FinancialServiceApi.md#getrecurringtransactionsforuser) | **GET** /financial-microservice/api/v1/transactions/recurring-transactions/{userId} | get recurring transactions |
*FinancialServiceApi* | [**getSmartGoalsByPocketId**](Apis/FinancialServiceApi.md#getsmartgoalsbypocketid) | **GET** /financial-microservice/api/v1/smart-goal/pocket/{pocketId} | get smart goals by pocket id |
*FinancialServiceApi* | [**getSplitTransaction**](Apis/FinancialServiceApi.md#getsplittransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/split | gets a split transaction |
*FinancialServiceApi* | [**getStudentLoanAccount**](Apis/FinancialServiceApi.md#getstudentloanaccount) | **GET** /financial-microservice/api/v1/account/{userId}/student-loan/{studentLoanAccountId} | get student loan account by id |
*FinancialServiceApi* | [**getTotalInvestmentBySecurity**](Apis/FinancialServiceApi.md#gettotalinvestmentbysecurity) | **GET** /financial-microservice/api/v1/analytics/total-investment/user/{userId} | Get Total Investment by user and security |
*FinancialServiceApi* | [**getTransaction**](Apis/FinancialServiceApi.md#gettransaction) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId} | lists a set of transactions against a given account of interest |
*FinancialServiceApi* | [**getTransactions**](Apis/FinancialServiceApi.md#gettransactions) | **GET** /financial-microservice/api/v1/transactions/user/{userId}/plaid-account-id/{plaidAccountId}/pageNumber/{pageNumber}/pageSize/{pageSize} | get transactions tied to a bank account and account id |
*FinancialServiceApi* | [**getTransactions1**](Apis/FinancialServiceApi.md#gettransactions1) | **GET** /financial-microservice/api/v1/transactions/{userId}/pageNumber/{pageNumber}/pageSize/{pageSize} | get transactions |
*FinancialServiceApi* | [**getTransactionsByTime**](Apis/FinancialServiceApi.md#gettransactionsbytime) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/range | get transactions by time |
*FinancialServiceApi* | [**getTransactionsForPastMonth**](Apis/FinancialServiceApi.md#gettransactionsforpastmonth) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/month | Get transactions for the past month |
*FinancialServiceApi* | [**getTransactionsForPastWeek**](Apis/FinancialServiceApi.md#gettransactionsforpastweek) | **GET** /financial-microservice/api/v1/users/{userId}/accounts/{plaidAccountId}/transactions/week | get transactions for the past week |
*FinancialServiceApi* | [**getUserAccountBalanceHistory**](Apis/FinancialServiceApi.md#getuseraccountbalancehistory) | **GET** /financial-microservice/api/v1/analytics/balance-history/user/{userId}/pagenumber/{pageNumber}/pagesize/{pageSize} | Returns the account balance history for a user |
*FinancialServiceApi* | [**getUserCategoryMonthlyExpenditure**](Apis/FinancialServiceApi.md#getusercategorymonthlyexpenditure) | **GET** /financial-microservice/api/v1/analytics/category-monthly-expenditure/user/{userId} | Returns the monthly category expenditure for a user |
*FinancialServiceApi* | [**getUserCategoryMonthlyIncome**](Apis/FinancialServiceApi.md#getusercategorymonthlyincome) | **GET** /financial-microservice/api/v1/analytics/category-monthly-income/user/{userId} | Get monthly income by user for a specific category |
*FinancialServiceApi* | [**getUserProfile1**](Apis/FinancialServiceApi.md#getuserprofile1) | **GET** /financial-microservice/api/v1/profile/{userId} | Gets a user profile |
*FinancialServiceApi* | [**healthCheck2**](Apis/FinancialServiceApi.md#healthcheck2) | **GET** /financial-microservice/api/v1/health | health check |
*FinancialServiceApi* | [**initiatePlaidSetup**](Apis/FinancialServiceApi.md#initiateplaidsetup) | **POST** /financial-microservice/api/v1/plaid/initiate-token-exchange | initiate plaid setup |
*FinancialServiceApi* | [**initiatePlaidTokenUpdate**](Apis/FinancialServiceApi.md#initiateplaidtokenupdate) | **POST** /financial-microservice/api/v1/plaid/initiate-token-update | initiate plaid link token update |
*FinancialServiceApi* | [**listRecurringTransactionNotes**](Apis/FinancialServiceApi.md#listrecurringtransactionnotes) | **GET** /financial-microservice/api/v1/transactions/recurring/{transactionId}/notes | lists notes from a transaction |
*FinancialServiceApi* | [**listRecurringTransactionsForUserAndAccount**](Apis/FinancialServiceApi.md#listrecurringtransactionsforuserandaccount) | **GET** /financial-microservice/api/v1/transactions/recurrings | lists a set of transactions against a given account of interest |
*FinancialServiceApi* | [**listTransactionNotes**](Apis/FinancialServiceApi.md#listtransactionnotes) | **GET** /financial-microservice/api/v1/transactions/transaction/{transactionId}/notes | lists notes from a transaction |
*FinancialServiceApi* | [**listTransactions**](Apis/FinancialServiceApi.md#listtransactions) | **GET** /financial-microservice/api/v1/transactions | lists a set of transactions against a given account of interest |
*FinancialServiceApi* | [**listTransactions1**](Apis/FinancialServiceApi.md#listtransactions1) | **GET** /financial-microservice/api/v1/transactions/all_accounts | lists a set of transactions across all connected accounts |
*FinancialServiceApi* | [**pollAsyncTaskExecutionStatus**](Apis/FinancialServiceApi.md#pollasynctaskexecutionstatus) | **GET** /financial-microservice/api/v1/async-task/{workflowId}/run/{runId} | polls the status of an async task |
*FinancialServiceApi* | [**readynessCheck2**](Apis/FinancialServiceApi.md#readynesscheck2) | **GET** /financial-microservice/api/v1/ready | readyness check |
*FinancialServiceApi* | [**searchTransactions**](Apis/FinancialServiceApi.md#searchtransactions) | **POST** /financial-microservice/api/v1/transactions/transaction/search | searches transactions |
*FinancialServiceApi* | [**splitTransaction**](Apis/FinancialServiceApi.md#splittransaction) | **POST** /financial-microservice/api/v1/transactions/transaction/split | splits a transaction |
*FinancialServiceApi* | [**transactionAggregates**](Apis/FinancialServiceApi.md#transactionaggregates) | **GET** /financial-microservice/api/v1/analytics/transaction-aggregates/{userId} | Returns the aggregated transactions for a user and month |
*FinancialServiceApi* | [**triggerSync**](Apis/FinancialServiceApi.md#triggersync) | **POST** /financial-microservice/api/v1/sync/trigger | Triggers a sync |
*FinancialServiceApi* | [**unsplitTransactions**](Apis/FinancialServiceApi.md#unsplittransactions) | **POST** /financial-microservice/api/v1/transactions/transaction/unsplit | unsplit a transaction |
*FinancialServiceApi* | [**updateBankAccount**](Apis/FinancialServiceApi.md#updatebankaccount) | **PUT** /financial-microservice/api/v1/bank-account | update a bank account for a given user profile |
*FinancialServiceApi* | [**updateNoteToRecurringTransaction**](Apis/FinancialServiceApi.md#updatenotetorecurringtransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring/note | Updates a note to a transaction |
*FinancialServiceApi* | [**updateNoteToSmartGoal**](Apis/FinancialServiceApi.md#updatenotetosmartgoal) | **PUT** /financial-microservice/api/v1/smart-goal/note | updates a note to a smart goal |
*FinancialServiceApi* | [**updateNoteToTransaction**](Apis/FinancialServiceApi.md#updatenotetotransaction) | **PUT** /financial-microservice/api/v1/transactions/transaction/note | Updates a note to a transaction |
*FinancialServiceApi* | [**updatePocket**](Apis/FinancialServiceApi.md#updatepocket) | **PUT** /financial-microservice/api/v1/pocket | updates a pocket |
*FinancialServiceApi* | [**updateSingleTransaction**](Apis/FinancialServiceApi.md#updatesingletransaction) | **PUT** /financial-microservice/api/v1/transactions/single-transaction | update a transaction |
*FinancialServiceApi* | [**updateSmartGoal**](Apis/FinancialServiceApi.md#updatesmartgoal) | **PUT** /financial-microservice/api/v1/smart-goal | update a smart goal |
*FinancialServiceApi* | [**updateTransaction**](Apis/FinancialServiceApi.md#updatetransaction) | **PUT** /financial-microservice/api/v1/transactions/recurring | update a transaction |
*FinancialServiceApi* | [**updateUserProfile**](Apis/FinancialServiceApi.md#updateuserprofile) | **PUT** /financial-microservice/api/v1/profile | update a user profile |
*FinancialServiceApi* | [**updatesBudget**](Apis/FinancialServiceApi.md#updatesbudget) | **PUT** /financial-microservice/api/v1/budget | updates a budget |
*FinancialServiceApi* | [**updatesMilestone**](Apis/FinancialServiceApi.md#updatesmilestone) | **PUT** /financial-microservice/api/v1/milestone | updates a milestone |
| *SocialServiceApi* | [**acceptFollowProfile**](Apis/SocialServiceApi.md#acceptfollowprofile) | **POST** /social-microservice/api/v1/follow-requests/{followRecordId}/accept | Accepts a user's follow request |
*SocialServiceApi* | [**addCommentQualityScore**](Apis/SocialServiceApi.md#addcommentqualityscore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/quality | Add Comment Quality Score |
*SocialServiceApi* | [**addPostQualityScore**](Apis/SocialServiceApi.md#addpostqualityscore) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/quality | Adds a quality score to a post |
*SocialServiceApi* | [**addPostToPublication**](Apis/SocialServiceApi.md#addposttopublication) | **POST** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId} | Add a post to a publication |
*SocialServiceApi* | [**addPostToThread**](Apis/SocialServiceApi.md#addposttothread) | **POST** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType} | Adds A Post To A Thread |
*SocialServiceApi* | [**addPublicationEditor**](Apis/SocialServiceApi.md#addpublicationeditor) | **PUT** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Adds an editor to a publication |
*SocialServiceApi* | [**blockUserProfile**](Apis/SocialServiceApi.md#blockuserprofile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/block/target/{targetUserId} | blocks a user profile |
*SocialServiceApi* | [**bookmarkPost**](Apis/SocialServiceApi.md#bookmarkpost) | **POST** /social-microservice/api/v1/users/{userId}/post/bookmark/{postId} | Bookmarks a post |
*SocialServiceApi* | [**bookmarkPublication**](Apis/SocialServiceApi.md#bookmarkpublication) | **POST** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Bookmarks a publication |
*SocialServiceApi* | [**createComment**](Apis/SocialServiceApi.md#createcomment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment | Create A Commnet |
*SocialServiceApi* | [**createCommentReply**](Apis/SocialServiceApi.md#createcommentreply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply | Reply To A Comment |
*SocialServiceApi* | [**createCommunityProfile**](Apis/SocialServiceApi.md#createcommunityprofile) | **POST** /social-microservice/api/v1/community-profiles/{userId} | Create a community Profile |
*SocialServiceApi* | [**createNote**](Apis/SocialServiceApi.md#createnote) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/note | Creates and associates a note to a given post |
*SocialServiceApi* | [**createPoll**](Apis/SocialServiceApi.md#createpoll) | **POST** /social-microservice/api/v1/users/{userId}/poll | Create a poll |
*SocialServiceApi* | [**createPost**](Apis/SocialServiceApi.md#createpost) | **POST** /social-microservice/api/v1/users/{userId}/post | Create a post |
*SocialServiceApi* | [**createPublication**](Apis/SocialServiceApi.md#createpublication) | **POST** /social-microservice/api/v1/users/{userId}/publication | Creates a publication |
*SocialServiceApi* | [**createTopic**](Apis/SocialServiceApi.md#createtopic) | **POST** /social-microservice/api/v1/users/{userId}/community/{communityProfileId}/topic | Create A Topic |
*SocialServiceApi* | [**createUserProfile**](Apis/SocialServiceApi.md#createuserprofile) | **POST** /social-microservice/api/v1/users | creates a user profile |
*SocialServiceApi* | [**deleteComment**](Apis/SocialServiceApi.md#deletecomment) | **DELETE** /social-microservice/api/v1/post/{postId}/comment/{commentId} | Delete A Comment |
*SocialServiceApi* | [**deleteCommentReply**](Apis/SocialServiceApi.md#deletecommentreply) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Delete A Comment Reply |
*SocialServiceApi* | [**deleteCommunityProfile**](Apis/SocialServiceApi.md#deletecommunityprofile) | **DELETE** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Delete Community Profile |
*SocialServiceApi* | [**deleteNote**](Apis/SocialServiceApi.md#deletenote) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Delete a note |
*SocialServiceApi* | [**deletePoll**](Apis/SocialServiceApi.md#deletepoll) | **DELETE** /social-microservice/api/v1/users/{userId}/poll/{postId} | Delete a poll |
*SocialServiceApi* | [**deletePost**](Apis/SocialServiceApi.md#deletepost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Delete a post |
*SocialServiceApi* | [**deletePostFromPublication**](Apis/SocialServiceApi.md#deletepostfrompublication) | **DELETE** /social-microservice/api/v1/users/editor/{editorUserId}/publication/{publicationId}/post/{postId} | Deletes a post from a publication |
*SocialServiceApi* | [**deletePublication**](Apis/SocialServiceApi.md#deletepublication) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId} | Deletes a publication |
*SocialServiceApi* | [**deletePublicationEditor**](Apis/SocialServiceApi.md#deletepublicationeditor) | **DELETE** /social-microservice/api/v1/users/admin/{adminUserId}/publication/{publicationId}/editor/{editorUserId} | Deletes an editor to a publication |
*SocialServiceApi* | [**deleteUserProfile**](Apis/SocialServiceApi.md#deleteuserprofile) | **DELETE** /social-microservice/api/v1/users/{userId} | deletes a user profile |
*SocialServiceApi* | [**discoverProfiles**](Apis/SocialServiceApi.md#discoverprofiles) | **GET** /social-microservice/api/v1/users/{userId}/discover/limit/{limit} | Discover Profiles |
*SocialServiceApi* | [**editCommentReply**](Apis/SocialServiceApi.md#editcommentreply) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId} | Edit A Comment Reply |
*SocialServiceApi* | [**editCommunityProfile**](Apis/SocialServiceApi.md#editcommunityprofile) | **PUT** /social-microservice/api/v1/users/{userId}/community-profiles/{communityProfileId} | Edit a community Profile |
*SocialServiceApi* | [**editNote**](Apis/SocialServiceApi.md#editnote) | **PUT** /social-microservice/api/v1/users/{userId}/post/{postId}/note/{noteId} | Update a note |
*SocialServiceApi* | [**editPost**](Apis/SocialServiceApi.md#editpost) | **PUT** /social-microservice/api/v1/post/{postId}/type/{postType} | Edits a post by id |
*SocialServiceApi* | [**editUserProfile**](Apis/SocialServiceApi.md#edituserprofile) | **PUT** /social-microservice/api/v1/users/{userId} | update a user profile |
*SocialServiceApi* | [**followCommunityProfile**](Apis/SocialServiceApi.md#followcommunityprofile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/community-profiles/{targetCommunityProfileId} | Follows A Community Profile |
*SocialServiceApi* | [**followProfile**](Apis/SocialServiceApi.md#followprofile) | **POST** /social-microservice/api/v1/users/source/{sourceUserId}/follow/target/{targetUserId} | follow a user profile |
*SocialServiceApi* | [**getAccountsFollowing**](Apis/SocialServiceApi.md#getaccountsfollowing) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/account-type/{accountType}/following | Get Communities and users you are following |
*SocialServiceApi* | [**getBlogPostsByTag**](Apis/SocialServiceApi.md#getblogpostsbytag) | **GET** /social-microservice/api/v1/posts/blog/tag/{tag} | Get blog posts by tag |
*SocialServiceApi* | [**getBookmarkedPosts**](Apis/SocialServiceApi.md#getbookmarkedposts) | **GET** /social-microservice/api/v1/users/bookmarks/{userId} | Get Bookmarked Posts |
*SocialServiceApi* | [**getCannyUserSSOToken1**](Apis/SocialServiceApi.md#getcannyuserssotoken1) | **GET** /social-microservice/api/v1/user/{userId}/canny/email/{email} | Retrieves user sso token for canny |
*SocialServiceApi* | [**getCommentReplies**](Apis/SocialServiceApi.md#getcommentreplies) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/replies | Get Comment Replies |
*SocialServiceApi* | [**getCommunitiesUserFollows**](Apis/SocialServiceApi.md#getcommunitiesuserfollows) | **GET** /social-microservice/api/v1/users/{userId}/communities-followed | Gets all the communities a user follows |
*SocialServiceApi* | [**getCommunityBlogPosts**](Apis/SocialServiceApi.md#getcommunityblogposts) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/posts/blog | Get community blog posts |
*SocialServiceApi* | [**getCommunityFeed**](Apis/SocialServiceApi.md#getcommunityfeed) | **GET** /social-microservice/api/v1/community-profiles/{communityProfileId}/timeline | Gets A Community Feed |
*SocialServiceApi* | [**getCommunityProfile**](Apis/SocialServiceApi.md#getcommunityprofile) | **GET** /social-microservice/api/v1/social/community-profiles/{communityId}/requestor/{requestorProfileId}/profile-type/{requestorProfileType} | Get a community Profile |
*SocialServiceApi* | [**getCommunityProfiles**](Apis/SocialServiceApi.md#getcommunityprofiles) | **GET** /social-microservice/api/v1/community-profiles/page-size/{pageSize}/page-number/{pageNumber} | Get Community Profiles |
*SocialServiceApi* | [**getFollowers**](Apis/SocialServiceApi.md#getfollowers) | **GET** /social-microservice/api/v1/users/{userId}/profile/{profileId}/followers | Get Users Following you |
*SocialServiceApi* | [**getPendingFollows**](Apis/SocialServiceApi.md#getpendingfollows) | **GET** /social-microservice/api/v1/users/{userId}/follow/pending-requests | Get Pending Follow Requests |
*SocialServiceApi* | [**getPoll**](Apis/SocialServiceApi.md#getpoll) | **GET** /social-microservice/api/v1/users/{userId}/poll/{postId} | Get a poll |
*SocialServiceApi* | [**getPolls**](Apis/SocialServiceApi.md#getpolls) | **GET** /social-microservice/api/v1/users/{userId}/polls | Get all the polls of a given user |
*SocialServiceApi* | [**getPost**](Apis/SocialServiceApi.md#getpost) | **GET** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType} | Get a post |
*SocialServiceApi* | [**getPostThread**](Apis/SocialServiceApi.md#getpostthread) | **GET** /social-microservice/api/v1/users/{userId}/post/thread/{postId} | Gets A Post's Thread |
*SocialServiceApi* | [**getPosts**](Apis/SocialServiceApi.md#getposts) | **GET** /social-microservice/api/v1/users/{userId}/posts | Get all the posts of a given user |
*SocialServiceApi* | [**getPostsByCategory**](Apis/SocialServiceApi.md#getpostsbycategory) | **GET** /social-microservice/api/v1/user/{userId}/category/{category}/posts/{postType}/limit/{limit}/offset/{offset} | Get all posts associated with a category |
*SocialServiceApi* | [**getPostsByTopic**](Apis/SocialServiceApi.md#getpostsbytopic) | **GET** /social-microservice/api/v1/community/{communityProfileId}/topic/{topicName}/posts | Get all posts associated with a topic |
*SocialServiceApi* | [**getPublication**](Apis/SocialServiceApi.md#getpublication) | **GET** /social-microservice/api/v1/users/{userId}/publication/{publicationId} | Gets a publication |
*SocialServiceApi* | [**getTopicsOfCommunitiesUserFollows**](Apis/SocialServiceApi.md#gettopicsofcommunitiesuserfollows) | **GET** /social-microservice/api/v1/users/{userId}/topics | Get Topics Of Communities User Follows |
*SocialServiceApi* | [**getUserFeed**](Apis/SocialServiceApi.md#getuserfeed) | **GET** /social-microservice/api/v1/users/{userId}/timeline | Gets A Userfeed |
*SocialServiceApi* | [**getUserProfile**](Apis/SocialServiceApi.md#getuserprofile) | **GET** /social-microservice/api/v1/users/{userId} | gets a user profile |
*SocialServiceApi* | [**getUserProfiles**](Apis/SocialServiceApi.md#getuserprofiles) | **GET** /social-microservice/api/v1/users/page-size/{pageSize}/page-number/{pageNumber} | Gets a set of user profiles |
*SocialServiceApi* | [**healthCheck1**](Apis/SocialServiceApi.md#healthcheck1) | **GET** /social-microservice/api/v1/health | health check |
*SocialServiceApi* | [**reactToComment**](Apis/SocialServiceApi.md#reacttocomment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment |
*SocialServiceApi* | [**reactToCommentReply**](Apis/SocialServiceApi.md#reacttocommentreply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/account-type/{accountType}/reaction/{reaction} | Reacts to a comment reply |
*SocialServiceApi* | [**reactToPost**](Apis/SocialServiceApi.md#reacttopost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/account-type/{accountType}/reaction/{reaction} | Reacts to a post |
*SocialServiceApi* | [**readynessCheck1**](Apis/SocialServiceApi.md#readynesscheck1) | **GET** /social-microservice/api/v1/ready | readyness check |
*SocialServiceApi* | [**removeBookmarkedPost**](Apis/SocialServiceApi.md#removebookmarkedpost) | **DELETE** /social-microservice/api/v1/users/{userId}/post/{postId}/bookmark | Deletes A Bookmarked Post |
*SocialServiceApi* | [**removeBookmarkedPublication**](Apis/SocialServiceApi.md#removebookmarkedpublication) | **DELETE** /social-microservice/api/v1/users/{userId}/publication/{publicationId}/bookmark | Deletes A Bookmarked Publication |
*SocialServiceApi* | [**removePostFromThread**](Apis/SocialServiceApi.md#removepostfromthread) | **DELETE** /social-microservice/api/v1/users/{userId}/post/thread/{parentPostId}/type/{postType}/target/{participantPostId} | Deletes A Post From A Thread |
*SocialServiceApi* | [**reportComment**](Apis/SocialServiceApi.md#reportcomment) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/report | Report A Comment |
*SocialServiceApi* | [**reportCommentReply**](Apis/SocialServiceApi.md#reportcommentreply) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/comment/{commentId}/reply/{replyId}/report | Report A Comment Reply |
*SocialServiceApi* | [**reportPost**](Apis/SocialServiceApi.md#reportpost) | **POST** /social-microservice/api/v1/users/{userId}/post/{postId}/type/{postType}/report | Report a post |
*SocialServiceApi* | [**respondToPoll**](Apis/SocialServiceApi.md#respondtopoll) | **POST** /social-microservice/api/v1/users/{userId}/poll/{pollId} | Adds a user response to a given poll by a user |
*SocialServiceApi* | [**sharePost**](Apis/SocialServiceApi.md#sharepost) | **POST** /social-microservice/api/v1/users/{userId}/post/share/{parentPostId}/type/{parentPostType} | Share a post |
| *UserServiceV1Api* | [**checkEmailExists**](Apis/UserServiceV1Api.md#checkemailexists) | **GET** /user-microservice/api/v1/user-service/user/email/{email}/exists | Checks that an email exists or not |
*UserServiceV1Api* | [**checkUsernameExists**](Apis/UserServiceV1Api.md#checkusernameexists) | **GET** /user-microservice/api/v1/user-service/user/username/{username}/exists | Checks that a username exists or not |
*UserServiceV1Api* | [**deleteUser**](Apis/UserServiceV1Api.md#deleteuser) | **DELETE** /user-microservice/api/v1/user-service/user/{userId} | deletes a user account |
*UserServiceV1Api* | [**getUser**](Apis/UserServiceV1Api.md#getuser) | **GET** /user-microservice/api/v1/user-service/user/{userId} | Gets a user account |
*UserServiceV1Api* | [**getUserByEmail**](Apis/UserServiceV1Api.md#getuserbyemail) | **GET** /user-microservice/api/v1/user-service/user/email/{email} | Gets a user account by email |
*UserServiceV1Api* | [**getUserByEmailOrUsername**](Apis/UserServiceV1Api.md#getuserbyemailorusername) | **GET** /user-microservice/api/v1/user-service/user/query-account-by-email-or-username | gets a user account by either email or username |
*UserServiceV1Api* | [**getUserByUsername**](Apis/UserServiceV1Api.md#getuserbyusername) | **GET** /user-microservice/api/v1/user-service/user/username/{username} | Gets a user account by user name |
*UserServiceV1Api* | [**getUserId**](Apis/UserServiceV1Api.md#getuserid) | **GET** /user-microservice/api/v1/user-service/user/query-id | get a user account id |
*UserServiceV1Api* | [**healthCheck**](Apis/UserServiceV1Api.md#healthcheck) | **GET** /user-microservice/api/v1/user-service/user/health | health check |
*UserServiceV1Api* | [**passwordReset**](Apis/UserServiceV1Api.md#passwordreset) | **POST** /user-microservice/api/v1/user-service/user/webhook/password-reset | password reset |
*UserServiceV1Api* | [**readynessCheck**](Apis/UserServiceV1Api.md#readynesscheck) | **GET** /user-microservice/api/v1/user-service/user/ready | readyness check |
*UserServiceV1Api* | [**updateUser**](Apis/UserServiceV1Api.md#updateuser) | **PUT** /user-microservice/api/v1/user-service/user | update a user account |
*UserServiceV1Api* | [**verification**](Apis/UserServiceV1Api.md#verification) | **POST** /user-microservice/api/v1/user-service/user/verification/{userId} | user verification |
| *UserServiceV2Api* | [**checkEmailAndAuth0UserIdExists**](Apis/UserServiceV2Api.md#checkemailandauth0useridexists) | **GET** /user-microservice/api/v2/user/email/{email}/auth0/{auth0UserId}/exists | Checks that an email and auth0 user id exists or not |
*UserServiceV2Api* | [**checkEmailExistsV2**](Apis/UserServiceV2Api.md#checkemailexistsv2) | **GET** /user-microservice/api/v2/user-service/user/email/{email}/exists | Checks that an email exists or not |
*UserServiceV2Api* | [**checkUsernameExistsV2**](Apis/UserServiceV2Api.md#checkusernameexistsv2) | **GET** /user-microservice/api/v2/user-service/user/username/{username}/exists | Checks that a username exists or not |
*UserServiceV2Api* | [**createRole**](Apis/UserServiceV2Api.md#createrole) | **POST** /user-microservice/api/v2/user-service/user/role | Creates a new role |
*UserServiceV2Api* | [**createUserV2**](Apis/UserServiceV2Api.md#createuserv2) | **POST** /user-microservice/api/v2/user-service/user | create a user account |
*UserServiceV2Api* | [**deleteRole**](Apis/UserServiceV2Api.md#deleterole) | **DELETE** /user-microservice/api/v2/user-service/user/role/{id} | Deletes a role |
*UserServiceV2Api* | [**deleteUserV2**](Apis/UserServiceV2Api.md#deleteuserv2) | **DELETE** /user-microservice/api/v2/user-service/user/{userId} | deletes a user account |
*UserServiceV2Api* | [**getCannyUserSSOToken**](Apis/UserServiceV2Api.md#getcannyuserssotoken) | **GET** /user-microservice/api/v2/user-service/user/canny/{userId} | Retrieves user sso token for canny |
*UserServiceV2Api* | [**getRole**](Apis/UserServiceV2Api.md#getrole) | **GET** /user-microservice/api/v2/user-service/user/role/{id} | Retrieves a role |
*UserServiceV2Api* | [**getUserByAuth0ID**](Apis/UserServiceV2Api.md#getuserbyauth0id) | **GET** /user-microservice/api/v2/user-service/user/auth-zero/{auth0UserId} | Retrieve user account details by auth0 id and profile type |
*UserServiceV2Api* | [**getUserByAuthnIDV2**](Apis/UserServiceV2Api.md#getuserbyauthnidv2) | **GET** /user-microservice/api/v2/user-service/user/authn/{authnId} | Retrieve user account details by authn id |
*UserServiceV2Api* | [**getUserByEmailOrUsernameV2**](Apis/UserServiceV2Api.md#getuserbyemailorusernamev2) | **GET** /user-microservice/api/v2/user-service/user/account/query-by-email-or-username | Retrieve user account by email or username |
*UserServiceV2Api* | [**getUserByEmailV2**](Apis/UserServiceV2Api.md#getuserbyemailv2) | **GET** /user-microservice/api/v2/user-service/user/email/{email} | Retrieve user details by email |
*UserServiceV2Api* | [**getUserByUsernameV2**](Apis/UserServiceV2Api.md#getuserbyusernamev2) | **GET** /user-microservice/api/v2/user-service/user/username/{username} | Retrieve user details by username |
*UserServiceV2Api* | [**getUserIdV2**](Apis/UserServiceV2Api.md#getuseridv2) | **GET** /user-microservice/api/v2/user-service/user/query-id | get a user account id |
*UserServiceV2Api* | [**getUserV2**](Apis/UserServiceV2Api.md#getuserv2) | **GET** /user-microservice/api/v2/user-service/user/{userId} | Retrieve user account details |
*UserServiceV2Api* | [**listRoles**](Apis/UserServiceV2Api.md#listroles) | **GET** /user-microservice/api/v2/user-service/user/roles | Lists all roles |
*UserServiceV2Api* | [**passwordResetWebhookV2**](Apis/UserServiceV2Api.md#passwordresetwebhookv2) | **POST** /user-microservice/api/v2/user-service/user/webhook/password-reset | Webhook for Password Reset |
*UserServiceV2Api* | [**retrieveBusinessSettings**](Apis/UserServiceV2Api.md#retrievebusinesssettings) | **GET** /user-microservice/api/v2/user-service/user/business/settings/{userId} | Retrieve Business Account Settings |
*UserServiceV2Api* | [**updateRole**](Apis/UserServiceV2Api.md#updaterole) | **PATCH** /user-microservice/api/v2/user/role | Updates an existing role |
*UserServiceV2Api* | [**updateUserV2**](Apis/UserServiceV2Api.md#updateuserv2) | **PUT** /user-microservice/api/v2/user-service/user | update a user account |
*UserServiceV2Api* | [**verifyUserV2**](Apis/UserServiceV2Api.md#verifyuserv2) | **POST** /user-microservice/api/v2/user-service/user/verification/{userId}/profile-type/{profileType} | User Email Verification |
| *WorkspaceServiceApi* | [**createAccount**](Apis/WorkspaceServiceApi.md#createaccount) | **POST** /workspace-microservice/api/v1/accounts | Create a new account |
*WorkspaceServiceApi* | [**createFolder**](Apis/WorkspaceServiceApi.md#createfolder) | **POST** /workspace-microservice/api/v1/folders | Create a folder |
*WorkspaceServiceApi* | [**createWorkspace**](Apis/WorkspaceServiceApi.md#createworkspace) | **POST** /workspace-microservice/api/v1/workspaces | Create a workspace |
*WorkspaceServiceApi* | [**deleteAccount**](Apis/WorkspaceServiceApi.md#deleteaccount) | **DELETE** /workspace-microservice/api/v1/accounts/{authZeroUserId} | Delete an account |
*WorkspaceServiceApi* | [**deleteFile**](Apis/WorkspaceServiceApi.md#deletefile) | **DELETE** /workspace-microservice/api/v1/files/{fileId} | Delete a file |
*WorkspaceServiceApi* | [**deleteFolder**](Apis/WorkspaceServiceApi.md#deletefolder) | **DELETE** /workspace-microservice/api/v1/folders/{folderId} | Delete a folder |
*WorkspaceServiceApi* | [**deleteWorkspace**](Apis/WorkspaceServiceApi.md#deleteworkspace) | **DELETE** /workspace-microservice/api/v1/workspaces/{workspaceId} | Delete a workspace |
*WorkspaceServiceApi* | [**getAccount**](Apis/WorkspaceServiceApi.md#getaccount) | **GET** /workspace-microservice/api/v1/accounts/{authZeroUserId} | Get account by ID |
*WorkspaceServiceApi* | [**listFolder**](Apis/WorkspaceServiceApi.md#listfolder) | **GET** /workspace-microservice/api/v1/folders | List folders |
*WorkspaceServiceApi* | [**listWorkspace**](Apis/WorkspaceServiceApi.md#listworkspace) | **GET** /workspace-microservice/api/v1/workspaces | List workspaces |
*WorkspaceServiceApi* | [**updateFile**](Apis/WorkspaceServiceApi.md#updatefile) | **PUT** /workspace-microservice/api/v1/files | Update a file |
*WorkspaceServiceApi* | [**updateFolder**](Apis/WorkspaceServiceApi.md#updatefolder) | **PUT** /workspace-microservice/api/v1/folders | Update a folder |
*WorkspaceServiceApi* | [**updateWorkspace**](Apis/WorkspaceServiceApi.md#updateworkspace) | **PUT** /workspace-microservice/api/v1/workspaces | Update a workspace |
| *WorkspaceServiceRestApi* | [**workspaceMicroserviceRestApiV1FileUploadPost**](Apis/WorkspaceServiceRestApi.md#workspacemicroservicerestapiv1fileuploadpost) | **POST** /workspace-microservice/rest-api/v1/file/upload | Uploads a file to the server |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [AcceptFollowProfileResponse](./Models/AcceptFollowProfileResponse.md)
 - [Account](./Models/Account.md)
 - [AccountBalanceHistory](./Models/AccountBalanceHistory.md)
 - [AccountType](./Models/AccountType.md)
 - [AccountingAttachment](./Models/AccountingAttachment.md)
 - [AccountingIntegrationMergeLink](./Models/AccountingIntegrationMergeLink.md)
 - [ActionableInsight](./Models/ActionableInsight.md)
 - [Actor](./Models/Actor.md)
 - [AddCommentQualityScoreResponse](./Models/AddCommentQualityScoreResponse.md)
 - [AddDefaultPocketsToBankAccountRequest](./Models/AddDefaultPocketsToBankAccountRequest.md)
 - [AddDefaultPocketsToBankAccountResponse](./Models/AddDefaultPocketsToBankAccountResponse.md)
 - [AddNoteToFinancialUserProfileRequest](./Models/AddNoteToFinancialUserProfileRequest.md)
 - [AddNoteToFinancialUserProfileResponse](./Models/AddNoteToFinancialUserProfileResponse.md)
 - [AddNoteToRecurringTransactionRequest](./Models/AddNoteToRecurringTransactionRequest.md)
 - [AddNoteToRecurringTransactionResponse](./Models/AddNoteToRecurringTransactionResponse.md)
 - [AddNoteToSmartGoalRequest](./Models/AddNoteToSmartGoalRequest.md)
 - [AddNoteToSmartGoalResponse](./Models/AddNoteToSmartGoalResponse.md)
 - [AddNoteToTransactionRequest](./Models/AddNoteToTransactionRequest.md)
 - [AddNoteToTransactionResponse](./Models/AddNoteToTransactionResponse.md)
 - [AddPostQualityScoreResponse](./Models/AddPostQualityScoreResponse.md)
 - [AddPostToPublicationResponse](./Models/AddPostToPublicationResponse.md)
 - [AddPostToThreadResponse](./Models/AddPostToThreadResponse.md)
 - [AddPublicationEditorResponse](./Models/AddPublicationEditorResponse.md)
 - [Address](./Models/Address.md)
 - [Any](./Models/Any.md)
 - [Any1](./Models/Any1.md)
 - [ApplicationTheme](./Models/ApplicationTheme.md)
 - [Apr](./Models/Apr.md)
 - [AuditAction](./Models/AuditAction.md)
 - [BalanceSheet](./Models/BalanceSheet.md)
 - [BankAccount](./Models/BankAccount.md)
 - [BankAccountStatus](./Models/BankAccountStatus.md)
 - [BankAccountType](./Models/BankAccountType.md)
 - [BankAccountType1](./Models/BankAccountType1.md)
 - [BankInfo](./Models/BankInfo.md)
 - [BaseTimeline](./Models/BaseTimeline.md)
 - [BenefitPlanType](./Models/BenefitPlanType.md)
 - [BlockUserProfileResponse](./Models/BlockUserProfileResponse.md)
 - [Bookmark](./Models/Bookmark.md)
 - [BookmarkPostResponse](./Models/BookmarkPostResponse.md)
 - [BookmarkPublicationResponse](./Models/BookmarkPublicationResponse.md)
 - [Budget](./Models/Budget.md)
 - [BulkUpdateRecurringTransactionRequest](./Models/BulkUpdateRecurringTransactionRequest.md)
 - [BulkUpdateRecurringTransactionResponse](./Models/BulkUpdateRecurringTransactionResponse.md)
 - [BulkUpdateTransactionRequest](./Models/BulkUpdateTransactionRequest.md)
 - [BulkUpdateTransactionResponse](./Models/BulkUpdateTransactionResponse.md)
 - [BusinessAccount](./Models/BusinessAccount.md)
 - [BusinessActionableInsight](./Models/BusinessActionableInsight.md)
 - [BusinessActionableInsightType](./Models/BusinessActionableInsightType.md)
 - [BusinessChartOfAccount](./Models/BusinessChartOfAccount.md)
 - [BusinessTransaction](./Models/BusinessTransaction.md)
 - [CashFlowStatement](./Models/CashFlowStatement.md)
 - [Category](./Models/Category.md)
 - [Category1](./Models/Category1.md)
 - [CategoryMetricsFinancialSubProfile](./Models/CategoryMetricsFinancialSubProfile.md)
 - [CategoryMonthlyExpenditure](./Models/CategoryMonthlyExpenditure.md)
 - [CategoryMonthlyIncome](./Models/CategoryMonthlyIncome.md)
 - [CategoryMonthlyTransactionCount](./Models/CategoryMonthlyTransactionCount.md)
 - [CheckEmailAndAuth0UserIdExistsResponse](./Models/CheckEmailAndAuth0UserIdExistsResponse.md)
 - [CheckEmailExistsResponse](./Models/CheckEmailExistsResponse.md)
 - [CheckEmailExistsV2Response](./Models/CheckEmailExistsV2Response.md)
 - [CheckIfQuotaExceededResponse](./Models/CheckIfQuotaExceededResponse.md)
 - [CheckUsernameExistsResponse](./Models/CheckUsernameExistsResponse.md)
 - [CheckUsernameExistsV2Response](./Models/CheckUsernameExistsV2Response.md)
 - [Comment](./Models/Comment.md)
 - [CommentReply](./Models/CommentReply.md)
 - [CommunityProfile](./Models/CommunityProfile.md)
 - [CompanyAddress](./Models/CompanyAddress.md)
 - [CompanyInfo](./Models/CompanyInfo.md)
 - [CompanyProfile](./Models/CompanyProfile.md)
 - [Contacts](./Models/Contacts.md)
 - [ContentInsights](./Models/ContentInsights.md)
 - [CreateAccountRequest](./Models/CreateAccountRequest.md)
 - [CreateAccountResponse](./Models/CreateAccountResponse.md)
 - [CreateAccountingProfileRequest](./Models/CreateAccountingProfileRequest.md)
 - [CreateAccountingProfileResponse](./Models/CreateAccountingProfileResponse.md)
 - [CreateBankAccountRequest](./Models/CreateBankAccountRequest.md)
 - [CreateBankAccountResponse](./Models/CreateBankAccountResponse.md)
 - [CreateBudgetRequest](./Models/CreateBudgetRequest.md)
 - [CreateBudgetResponse](./Models/CreateBudgetResponse.md)
 - [CreateCommentReplyBody](./Models/CreateCommentReplyBody.md)
 - [CreateCommentReplyResponse](./Models/CreateCommentReplyResponse.md)
 - [CreateCommentResponse](./Models/CreateCommentResponse.md)
 - [CreateCommunityProfileBody](./Models/CreateCommunityProfileBody.md)
 - [CreateCommunityProfileResponse](./Models/CreateCommunityProfileResponse.md)
 - [CreateFolderRequest](./Models/CreateFolderRequest.md)
 - [CreateFolderResponse](./Models/CreateFolderResponse.md)
 - [CreateManualLinkRequest](./Models/CreateManualLinkRequest.md)
 - [CreateManualLinkResponse](./Models/CreateManualLinkResponse.md)
 - [CreateMilestoneRequest](./Models/CreateMilestoneRequest.md)
 - [CreateMilestoneResponse](./Models/CreateMilestoneResponse.md)
 - [CreateNoteBody](./Models/CreateNoteBody.md)
 - [CreateNoteResponse](./Models/CreateNoteResponse.md)
 - [CreatePollResponse](./Models/CreatePollResponse.md)
 - [CreatePostResponse](./Models/CreatePostResponse.md)
 - [CreatePublicationResponse](./Models/CreatePublicationResponse.md)
 - [CreateRoleResponse](./Models/CreateRoleResponse.md)
 - [CreateSmartGoalRequest](./Models/CreateSmartGoalRequest.md)
 - [CreateSmartGoalResponse](./Models/CreateSmartGoalResponse.md)
 - [CreateSubscriptionRequest](./Models/CreateSubscriptionRequest.md)
 - [CreateSubscriptionResponse](./Models/CreateSubscriptionResponse.md)
 - [CreateTopicResponse](./Models/CreateTopicResponse.md)
 - [CreateUserProfileRequest](./Models/CreateUserProfileRequest.md)
 - [CreateUserProfileRequest1](./Models/CreateUserProfileRequest1.md)
 - [CreateUserProfileResponse](./Models/CreateUserProfileResponse.md)
 - [CreateUserProfileResponse1](./Models/CreateUserProfileResponse1.md)
 - [CreateUserV2Request](./Models/CreateUserV2Request.md)
 - [CreateUserV2Response](./Models/CreateUserV2Response.md)
 - [CreateWorkspaceRequest](./Models/CreateWorkspaceRequest.md)
 - [CreateWorkspaceResponse](./Models/CreateWorkspaceResponse.md)
 - [CreditAccount](./Models/CreditAccount.md)
 - [CreditNote](./Models/CreditNote.md)
 - [CreditNoteLineItem](./Models/CreditNoteLineItem.md)
 - [DebtToIncomeRatio](./Models/DebtToIncomeRatio.md)
 - [Deduction](./Models/Deduction.md)
 - [DeleteAccountResponse](./Models/DeleteAccountResponse.md)
 - [DeleteAccountingProfileResponse](./Models/DeleteAccountingProfileResponse.md)
 - [DeleteBankAccountResponse](./Models/DeleteBankAccountResponse.md)
 - [DeleteBudgetResponse](./Models/DeleteBudgetResponse.md)
 - [DeleteCommentReplyResponse](./Models/DeleteCommentReplyResponse.md)
 - [DeleteCommentResponse](./Models/DeleteCommentResponse.md)
 - [DeleteCommunityProfileResponse](./Models/DeleteCommunityProfileResponse.md)
 - [DeleteFileResponse](./Models/DeleteFileResponse.md)
 - [DeleteFolderResponse](./Models/DeleteFolderResponse.md)
 - [DeleteLinkResponse](./Models/DeleteLinkResponse.md)
 - [DeleteMilestoneResponse](./Models/DeleteMilestoneResponse.md)
 - [DeleteNoteFromRecurringTransactionResponse](./Models/DeleteNoteFromRecurringTransactionResponse.md)
 - [DeleteNoteFromSmartGoalResponse](./Models/DeleteNoteFromSmartGoalResponse.md)
 - [DeleteNoteFromTransactionResponse](./Models/DeleteNoteFromTransactionResponse.md)
 - [DeleteNoteResponse](./Models/DeleteNoteResponse.md)
 - [DeletePocketResponse](./Models/DeletePocketResponse.md)
 - [DeletePollResponse](./Models/DeletePollResponse.md)
 - [DeletePostFromPublicationResponse](./Models/DeletePostFromPublicationResponse.md)
 - [DeletePostResponse](./Models/DeletePostResponse.md)
 - [DeletePublicationEditorResponse](./Models/DeletePublicationEditorResponse.md)
 - [DeletePublicationResponse](./Models/DeletePublicationResponse.md)
 - [DeleteRecurringTransactionResponse](./Models/DeleteRecurringTransactionResponse.md)
 - [DeleteRoleResponse](./Models/DeleteRoleResponse.md)
 - [DeleteSmartGoalResponse](./Models/DeleteSmartGoalResponse.md)
 - [DeleteTransactionResponse](./Models/DeleteTransactionResponse.md)
 - [DeleteUserProfileResponse](./Models/DeleteUserProfileResponse.md)
 - [DeleteUserProfileResponse1](./Models/DeleteUserProfileResponse1.md)
 - [DeleteUserResponse](./Models/DeleteUserResponse.md)
 - [DeleteUserV2Response](./Models/DeleteUserV2Response.md)
 - [DeleteWorkspaceResponse](./Models/DeleteWorkspaceResponse.md)
 - [DependentRelationship](./Models/DependentRelationship.md)
 - [Dependents](./Models/Dependents.md)
 - [DigitalWorkerSettings](./Models/DigitalWorkerSettings.md)
 - [DiscoverProfilesResponse](./Models/DiscoverProfilesResponse.md)
 - [Earning](./Models/Earning.md)
 - [EarningType](./Models/EarningType.md)
 - [EditCommentReplyBody](./Models/EditCommentReplyBody.md)
 - [EditCommentReplyResponse](./Models/EditCommentReplyResponse.md)
 - [EditCommunityProfileRequest](./Models/EditCommunityProfileRequest.md)
 - [EditCommunityProfileResponse](./Models/EditCommunityProfileResponse.md)
 - [EditNoteResponse](./Models/EditNoteResponse.md)
 - [EditPostResponse](./Models/EditPostResponse.md)
 - [EditUserProfileResponse](./Models/EditUserProfileResponse.md)
 - [EmployeTimeOffBalance](./Models/EmployeTimeOffBalance.md)
 - [Employee](./Models/Employee.md)
 - [EmployeeBenefits](./Models/EmployeeBenefits.md)
 - [EmployeeJobPositionAtCompany](./Models/EmployeeJobPositionAtCompany.md)
 - [EmployeePayrollRun](./Models/EmployeePayrollRun.md)
 - [EmployerBenefits](./Models/EmployerBenefits.md)
 - [EmployerPayrollRun](./Models/EmployerPayrollRun.md)
 - [EmploymentStatus](./Models/EmploymentStatus.md)
 - [Entities](./Models/Entities.md)
 - [ErrorCode](./Models/ErrorCode.md)
 - [ErrorCode1](./Models/ErrorCode1.md)
 - [Ethnicity](./Models/Ethnicity.md)
 - [ExchangePublicLinkTokenForAccountTokenRequest](./Models/ExchangePublicLinkTokenForAccountTokenRequest.md)
 - [ExchangePublicLinkTokenForAccountTokenResponse](./Models/ExchangePublicLinkTokenForAccountTokenResponse.md)
 - [Expense](./Models/Expense.md)
 - [ExpenseLine](./Models/ExpenseLine.md)
 - [ExpenseMetrics](./Models/ExpenseMetrics.md)
 - [ExpenseMetricsFinancialSubProfileMetrics](./Models/ExpenseMetricsFinancialSubProfileMetrics.md)
 - [FeedActivity](./Models/FeedActivity.md)
 - [FeedType](./Models/FeedType.md)
 - [FileMetadata](./Models/FileMetadata.md)
 - [FinancialAccountType](./Models/FinancialAccountType.md)
 - [FinancialPreferences](./Models/FinancialPreferences.md)
 - [FinancialProfile](./Models/FinancialProfile.md)
 - [FinancialUserProfile](./Models/FinancialUserProfile.md)
 - [FinancialUserProfileType](./Models/FinancialUserProfileType.md)
 - [FlsaStatus](./Models/FlsaStatus.md)
 - [FolderMetadata](./Models/FolderMetadata.md)
 - [FollowCommunityProfileResponse](./Models/FollowCommunityProfileResponse.md)
 - [FollowProfileResponse](./Models/FollowProfileResponse.md)
 - [Forecast](./Models/Forecast.md)
 - [Gender](./Models/Gender.md)
 - [GetAccountBalanceHistoryResponse](./Models/GetAccountBalanceHistoryResponse.md)
 - [GetAccountResponse](./Models/GetAccountResponse.md)
 - [GetAccountsFollowingResponse](./Models/GetAccountsFollowingResponse.md)
 - [GetAllBudgetsResponse](./Models/GetAllBudgetsResponse.md)
 - [GetBankAccountResponse](./Models/GetBankAccountResponse.md)
 - [GetBlogPostsByTagResponse](./Models/GetBlogPostsByTagResponse.md)
 - [GetBookmarkedPostsResponse](./Models/GetBookmarkedPostsResponse.md)
 - [GetBudgetResponse](./Models/GetBudgetResponse.md)
 - [GetBusinessSettingsResponse](./Models/GetBusinessSettingsResponse.md)
 - [GetCannyUserSSOTokenResponse](./Models/GetCannyUserSSOTokenResponse.md)
 - [GetCategoryMonthlyTransactionCountResponse](./Models/GetCategoryMonthlyTransactionCountResponse.md)
 - [GetCommentRepliesResponse](./Models/GetCommentRepliesResponse.md)
 - [GetCommunitiesUserFollowsResponse](./Models/GetCommunitiesUserFollowsResponse.md)
 - [GetCommunityBlogPostsResponse](./Models/GetCommunityBlogPostsResponse.md)
 - [GetCommunityFeedResponse](./Models/GetCommunityFeedResponse.md)
 - [GetCommunityProfileResponse](./Models/GetCommunityProfileResponse.md)
 - [GetCommunityProfilesResponse](./Models/GetCommunityProfilesResponse.md)
 - [GetDebtToIncomeRatioResponse](./Models/GetDebtToIncomeRatioResponse.md)
 - [GetExpenseMetricsResponse](./Models/GetExpenseMetricsResponse.md)
 - [GetFinancialProfileResponse](./Models/GetFinancialProfileResponse.md)
 - [GetFollowersResponse](./Models/GetFollowersResponse.md)
 - [GetForecastResponse](./Models/GetForecastResponse.md)
 - [GetHistoricalAccountBalanceResponse](./Models/GetHistoricalAccountBalanceResponse.md)
 - [GetIncomeExpenseRatioResponse](./Models/GetIncomeExpenseRatioResponse.md)
 - [GetIncomeMetricsResponse](./Models/GetIncomeMetricsResponse.md)
 - [GetInvestmentAcccountResponse](./Models/GetInvestmentAcccountResponse.md)
 - [GetLiabilityAccountResponse](./Models/GetLiabilityAccountResponse.md)
 - [GetLinkResponse](./Models/GetLinkResponse.md)
 - [GetLinksResponse](./Models/GetLinksResponse.md)
 - [GetMelodyFinancialContextResponse](./Models/GetMelodyFinancialContextResponse.md)
 - [GetMerchantMonthlyExpenditureResponse](./Models/GetMerchantMonthlyExpenditureResponse.md)
 - [GetMergeLinkTokenRequest](./Models/GetMergeLinkTokenRequest.md)
 - [GetMergeLinkTokenResponse](./Models/GetMergeLinkTokenResponse.md)
 - [GetMilestoneResponse](./Models/GetMilestoneResponse.md)
 - [GetMilestonesBySmartGoalIdResponse](./Models/GetMilestonesBySmartGoalIdResponse.md)
 - [GetMonthlyBalanceResponse](./Models/GetMonthlyBalanceResponse.md)
 - [GetMonthlyExpenditureResponse](./Models/GetMonthlyExpenditureResponse.md)
 - [GetMonthlyIncomeResponse](./Models/GetMonthlyIncomeResponse.md)
 - [GetMonthlySavingsResponse](./Models/GetMonthlySavingsResponse.md)
 - [GetMonthlyTotalQuantityBySecurityAndUserResponse](./Models/GetMonthlyTotalQuantityBySecurityAndUserResponse.md)
 - [GetMonthlyTransactionCountResponse](./Models/GetMonthlyTransactionCountResponse.md)
 - [GetMortgageAccountResponse](./Models/GetMortgageAccountResponse.md)
 - [GetNoteFromSmartGoalResponse](./Models/GetNoteFromSmartGoalResponse.md)
 - [GetNoteFromTransactionResponse](./Models/GetNoteFromTransactionResponse.md)
 - [GetNotesFromFinancialUserProfileResponse](./Models/GetNotesFromFinancialUserProfileResponse.md)
 - [GetNotesFromSmartGoalResponse](./Models/GetNotesFromSmartGoalResponse.md)
 - [GetPaymentChannelMonthlyExpenditureResponse](./Models/GetPaymentChannelMonthlyExpenditureResponse.md)
 - [GetPendingFollowsResponse](./Models/GetPendingFollowsResponse.md)
 - [GetPocketResponse](./Models/GetPocketResponse.md)
 - [GetPollResponse](./Models/GetPollResponse.md)
 - [GetPollsResponse](./Models/GetPollsResponse.md)
 - [GetPostResponse](./Models/GetPostResponse.md)
 - [GetPostThreadResponse](./Models/GetPostThreadResponse.md)
 - [GetPostsByCategoryResponse](./Models/GetPostsByCategoryResponse.md)
 - [GetPostsByTopicResponse](./Models/GetPostsByTopicResponse.md)
 - [GetPostsResponse](./Models/GetPostsResponse.md)
 - [GetPublicationResponse](./Models/GetPublicationResponse.md)
 - [GetRecurringTransactionsForUserResponse](./Models/GetRecurringTransactionsForUserResponse.md)
 - [GetRoleResponse](./Models/GetRoleResponse.md)
 - [GetSingleRecurringTransactionResponse](./Models/GetSingleRecurringTransactionResponse.md)
 - [GetSmartGoalsByPocketIdResponse](./Models/GetSmartGoalsByPocketIdResponse.md)
 - [GetSplitTransactionResponse](./Models/GetSplitTransactionResponse.md)
 - [GetStudentLoanAccountResponse](./Models/GetStudentLoanAccountResponse.md)
 - [GetTopicsOfCommunitiesUserFollowsResponse](./Models/GetTopicsOfCommunitiesUserFollowsResponse.md)
 - [GetTotalInvestmentBySecurityResponse](./Models/GetTotalInvestmentBySecurityResponse.md)
 - [GetTransactionAggregatesResponse](./Models/GetTransactionAggregatesResponse.md)
 - [GetTransactionResponse](./Models/GetTransactionResponse.md)
 - [GetTransactionsBetweenTimeRangesResponse](./Models/GetTransactionsBetweenTimeRangesResponse.md)
 - [GetTransactionsForBankAccountResponse](./Models/GetTransactionsForBankAccountResponse.md)
 - [GetTransactionsForPastMonthResponse](./Models/GetTransactionsForPastMonthResponse.md)
 - [GetTransactionsForPastWeekResponse](./Models/GetTransactionsForPastWeekResponse.md)
 - [GetTransactionsResponse](./Models/GetTransactionsResponse.md)
 - [GetUserAccountBalanceHistoryResponse](./Models/GetUserAccountBalanceHistoryResponse.md)
 - [GetUserByAuth0IDResponse](./Models/GetUserByAuth0IDResponse.md)
 - [GetUserByAuthnIDV2Response](./Models/GetUserByAuthnIDV2Response.md)
 - [GetUserByEmailOrUsernameResponse](./Models/GetUserByEmailOrUsernameResponse.md)
 - [GetUserByEmailOrUsernameV2Response](./Models/GetUserByEmailOrUsernameV2Response.md)
 - [GetUserByEmailResponse](./Models/GetUserByEmailResponse.md)
 - [GetUserByEmailV2Response](./Models/GetUserByEmailV2Response.md)
 - [GetUserByUsernameResponse](./Models/GetUserByUsernameResponse.md)
 - [GetUserByUsernameV2Response](./Models/GetUserByUsernameV2Response.md)
 - [GetUserCategoryMonthlyExpenditureResponse](./Models/GetUserCategoryMonthlyExpenditureResponse.md)
 - [GetUserCategoryMonthlyIncomeResponse](./Models/GetUserCategoryMonthlyIncomeResponse.md)
 - [GetUserFeedResponse](./Models/GetUserFeedResponse.md)
 - [GetUserIdResponse](./Models/GetUserIdResponse.md)
 - [GetUserIdV2Response](./Models/GetUserIdV2Response.md)
 - [GetUserProfileResponse](./Models/GetUserProfileResponse.md)
 - [GetUserProfileResponse1](./Models/GetUserProfileResponse1.md)
 - [GetUserProfilesResponse](./Models/GetUserProfilesResponse.md)
 - [GetUserResponse](./Models/GetUserResponse.md)
 - [GetUserV2Response](./Models/GetUserV2Response.md)
 - [GoalType](./Models/GoalType.md)
 - [Group](./Models/Group.md)
 - [HealthCheckResponse](./Models/HealthCheckResponse.md)
 - [HealthCheckResponse1](./Models/HealthCheckResponse1.md)
 - [HealthCheckResponse2](./Models/HealthCheckResponse2.md)
 - [HrisIntegrationMergeLink](./Models/HrisIntegrationMergeLink.md)
 - [HrisLinkedAccount](./Models/HrisLinkedAccount.md)
 - [IncomeExpenseRatio](./Models/IncomeExpenseRatio.md)
 - [IncomeMetrics](./Models/IncomeMetrics.md)
 - [IncomeMetricsFinancialSubProfile](./Models/IncomeMetricsFinancialSubProfile.md)
 - [IncomeStatement](./Models/IncomeStatement.md)
 - [InternalErrorCode](./Models/InternalErrorCode.md)
 - [InternalErrorMessageResponse](./Models/InternalErrorMessageResponse.md)
 - [InvesmentHolding](./Models/InvesmentHolding.md)
 - [InvestmentAccount](./Models/InvestmentAccount.md)
 - [InvestmentSecurity](./Models/InvestmentSecurity.md)
 - [Invoice](./Models/Invoice.md)
 - [InvoiceLineItem](./Models/InvoiceLineItem.md)
 - [Item](./Models/Item.md)
 - [JournalEntry](./Models/JournalEntry.md)
 - [JournalLine](./Models/JournalLine.md)
 - [LikedDashboardPanels](./Models/LikedDashboardPanels.md)
 - [Link](./Models/Link.md)
 - [LinkStatus](./Models/LinkStatus.md)
 - [LinkType](./Models/LinkType.md)
 - [LinkedAccountingAccount](./Models/LinkedAccountingAccount.md)
 - [ListFolderResponse](./Models/ListFolderResponse.md)
 - [ListRecurringTransactionNotesResponse](./Models/ListRecurringTransactionNotesResponse.md)
 - [ListRecurringTransactionsForUserAndAccountResponse](./Models/ListRecurringTransactionsForUserAndAccountResponse.md)
 - [ListRolesResponse](./Models/ListRolesResponse.md)
 - [ListTransactionNotesResponse](./Models/ListTransactionNotesResponse.md)
 - [ListTransactionsAcrossAllAccountsResponse](./Models/ListTransactionsAcrossAllAccountsResponse.md)
 - [ListTransactionsResponse](./Models/ListTransactionsResponse.md)
 - [ListWorkspaceResponse](./Models/ListWorkspaceResponse.md)
 - [LocationAddress](./Models/LocationAddress.md)
 - [LocationFinancialSubProfile](./Models/LocationFinancialSubProfile.md)
 - [LocationType](./Models/LocationType.md)
 - [MaritalStatus](./Models/MaritalStatus.md)
 - [Media](./Models/Media.md)
 - [MediaCrop](./Models/MediaCrop.md)
 - [MediaMetadata](./Models/MediaMetadata.md)
 - [MediaResize](./Models/MediaResize.md)
 - [MediaType](./Models/MediaType.md)
 - [MelodyFinancialContext](./Models/MelodyFinancialContext.md)
 - [MerchantMetricsFinancialSubProfile](./Models/MerchantMetricsFinancialSubProfile.md)
 - [MerchantMonthlyExpenditure](./Models/MerchantMonthlyExpenditure.md)
 - [MergeBusinessProfile](./Models/MergeBusinessProfile.md)
 - [MergeLinkedAccountToken](./Models/MergeLinkedAccountToken.md)
 - [Milestone](./Models/Milestone.md)
 - [MonthlyBalance](./Models/MonthlyBalance.md)
 - [MonthlyExpenditure](./Models/MonthlyExpenditure.md)
 - [MonthlyIncome](./Models/MonthlyIncome.md)
 - [MonthlySavings](./Models/MonthlySavings.md)
 - [MonthlyTotalQuantityBySecurityAndUser](./Models/MonthlyTotalQuantityBySecurityAndUser.md)
 - [MonthlyTransactionCount](./Models/MonthlyTransactionCount.md)
 - [MortgageAccount](./Models/MortgageAccount.md)
 - [NotFoundErrorCode](./Models/NotFoundErrorCode.md)
 - [Note](./Models/Note.md)
 - [NotificationActivity](./Models/NotificationActivity.md)
 - [NotificationFeedGroup](./Models/NotificationFeedGroup.md)
 - [NotificationSettings](./Models/NotificationSettings.md)
 - [NotificationTimeline](./Models/NotificationTimeline.md)
 - [NotificationType](./Models/NotificationType.md)
 - [PasswordResetWebhookResponse](./Models/PasswordResetWebhookResponse.md)
 - [PasswordResetWebhookV2Response](./Models/PasswordResetWebhookV2Response.md)
 - [PathUnknownErrorMessageResponse](./Models/PathUnknownErrorMessageResponse.md)
 - [PayFrequency](./Models/PayFrequency.md)
 - [PayGroupType](./Models/PayGroupType.md)
 - [PayPeriod](./Models/PayPeriod.md)
 - [Payment](./Models/Payment.md)
 - [PaymentChannelMetricsFinancialSubProfile](./Models/PaymentChannelMetricsFinancialSubProfile.md)
 - [PaymentChannelMonthlyExpenditure](./Models/PaymentChannelMonthlyExpenditure.md)
 - [PayrollRunState](./Models/PayrollRunState.md)
 - [PayrollRunType](./Models/PayrollRunType.md)
 - [PendingFollowRequest](./Models/PendingFollowRequest.md)
 - [PersonalActionableInsight](./Models/PersonalActionableInsight.md)
 - [PersonalActionableInsightName](./Models/PersonalActionableInsightName.md)
 - [PlaidAccountInvestmentTransaction](./Models/PlaidAccountInvestmentTransaction.md)
 - [PlaidAccountRecurringTransaction](./Models/PlaidAccountRecurringTransaction.md)
 - [PlaidAccountTransaction](./Models/PlaidAccountTransaction.md)
 - [PlaidExchangeTokenRequest](./Models/PlaidExchangeTokenRequest.md)
 - [PlaidExchangeTokenResponse](./Models/PlaidExchangeTokenResponse.md)
 - [PlaidInitiateTokenExchangeRequest](./Models/PlaidInitiateTokenExchangeRequest.md)
 - [PlaidInitiateTokenExchangeResponse](./Models/PlaidInitiateTokenExchangeResponse.md)
 - [PlaidInitiateTokenUpdateRequest](./Models/PlaidInitiateTokenUpdateRequest.md)
 - [PlaidInitiateTokenUpdateResponse](./Models/PlaidInitiateTokenUpdateResponse.md)
 - [PlaidLink](./Models/PlaidLink.md)
 - [PlaidSync](./Models/PlaidSync.md)
 - [Pocket](./Models/Pocket.md)
 - [PocketType](./Models/PocketType.md)
 - [PolicyType](./Models/PolicyType.md)
 - [PollAsyncTaskExecutionStatusResponse](./Models/PollAsyncTaskExecutionStatusResponse.md)
 - [PollPost](./Models/PollPost.md)
 - [PollResponse](./Models/PollResponse.md)
 - [Post](./Models/Post.md)
 - [PostType](./Models/PostType.md)
 - [ProfileType](./Models/ProfileType.md)
 - [Publication](./Models/Publication.md)
 - [PublicationType](./Models/PublicationType.md)
 - [PurchaseOrder](./Models/PurchaseOrder.md)
 - [PurchaseOrderLineItem](./Models/PurchaseOrderLineItem.md)
 - [ReactToCommentReplyResponse](./Models/ReactToCommentReplyResponse.md)
 - [ReactToCommentResponse](./Models/ReactToCommentResponse.md)
 - [ReactToPostResponse](./Models/ReactToPostResponse.md)
 - [Reaction](./Models/Reaction.md)
 - [ReadAccountingProfileResponse](./Models/ReadAccountingProfileResponse.md)
 - [ReadBalanceSheetsRequest](./Models/ReadBalanceSheetsRequest.md)
 - [ReadBalanceSheetsResponse](./Models/ReadBalanceSheetsResponse.md)
 - [ReadBusinessChartOfAccountsRequest](./Models/ReadBusinessChartOfAccountsRequest.md)
 - [ReadBusinessChartOfAccountsResponse](./Models/ReadBusinessChartOfAccountsResponse.md)
 - [ReadBusinessTransactionsRequest](./Models/ReadBusinessTransactionsRequest.md)
 - [ReadBusinessTransactionsResponse](./Models/ReadBusinessTransactionsResponse.md)
 - [ReadCashFlowStatementsRequest](./Models/ReadCashFlowStatementsRequest.md)
 - [ReadCashFlowStatementsResponse](./Models/ReadCashFlowStatementsResponse.md)
 - [ReadIncomeStatementsRequest](./Models/ReadIncomeStatementsRequest.md)
 - [ReadIncomeStatementsResponse](./Models/ReadIncomeStatementsResponse.md)
 - [ReadynessCheckResponse](./Models/ReadynessCheckResponse.md)
 - [ReadynessCheckResponse1](./Models/ReadynessCheckResponse1.md)
 - [ReadynessCheckResponse2](./Models/ReadynessCheckResponse2.md)
 - [RecordAskCopilotQuestionRequest](./Models/RecordAskCopilotQuestionRequest.md)
 - [RecordAskCopilotQuestionResponse](./Models/RecordAskCopilotQuestionResponse.md)
 - [RemoveBookmarkedPostResponse](./Models/RemoveBookmarkedPostResponse.md)
 - [RemoveBookmarkedPublicationResponse](./Models/RemoveBookmarkedPublicationResponse.md)
 - [RemovePostFromThreadResponse](./Models/RemovePostFromThreadResponse.md)
 - [ReportCommentBody](./Models/ReportCommentBody.md)
 - [ReportCommentReplyBody](./Models/ReportCommentReplyBody.md)
 - [ReportCommentReplyResponse](./Models/ReportCommentReplyResponse.md)
 - [ReportCommentResponse](./Models/ReportCommentResponse.md)
 - [ReportItem](./Models/ReportItem.md)
 - [ReportPostResponse](./Models/ReportPostResponse.md)
 - [RespondToPollBody](./Models/RespondToPollBody.md)
 - [RespondToPollResponse](./Models/RespondToPollResponse.md)
 - [RiskToleranceSettings](./Models/RiskToleranceSettings.md)
 - [Role](./Models/Role.md)
 - [RoleAuditEvents](./Models/RoleAuditEvents.md)
 - [RoleType](./Models/RoleType.md)
 - [SearchTransactionsRequest](./Models/SearchTransactionsRequest.md)
 - [SearchTransactionsResponse](./Models/SearchTransactionsResponse.md)
 - [Sentiment](./Models/Sentiment.md)
 - [Settings](./Models/Settings.md)
 - [SharePostResponse](./Models/SharePostResponse.md)
 - [SharedPost](./Models/SharedPost.md)
 - [SmartGoal](./Models/SmartGoal.md)
 - [SmartNote](./Models/SmartNote.md)
 - [SocialProfileMetadata](./Models/SocialProfileMetadata.md)
 - [SocialRelationshipMetadata](./Models/SocialRelationshipMetadata.md)
 - [SplitTransactionRequest](./Models/SplitTransactionRequest.md)
 - [SplitTransactionResponse](./Models/SplitTransactionResponse.md)
 - [Status](./Models/Status.md)
 - [Status1](./Models/Status1.md)
 - [StripeSubscription](./Models/StripeSubscription.md)
 - [StripeSubscriptionStatus](./Models/StripeSubscriptionStatus.md)
 - [StudentLoanAccount](./Models/StudentLoanAccount.md)
 - [Tags](./Models/Tags.md)
 - [TaskState](./Models/TaskState.md)
 - [Tax](./Models/Tax.md)
 - [TaxRate](./Models/TaxRate.md)
 - [Thread](./Models/Thread.md)
 - [ThreadParticipantType](./Models/ThreadParticipantType.md)
 - [Token](./Models/Token.md)
 - [Topic](./Models/Topic.md)
 - [TotalInvestmentBySecurity](./Models/TotalInvestmentBySecurity.md)
 - [Transaction](./Models/Transaction.md)
 - [TransactionAggregatesByMonth](./Models/TransactionAggregatesByMonth.md)
 - [TransactionLineItem](./Models/TransactionLineItem.md)
 - [TransactionSplit](./Models/TransactionSplit.md)
 - [TriggerSyncRequest](./Models/TriggerSyncRequest.md)
 - [TriggerSyncResponse](./Models/TriggerSyncResponse.md)
 - [UnSplitTransactionsRequest](./Models/UnSplitTransactionsRequest.md)
 - [UnSplitTransactionsResponse](./Models/UnSplitTransactionsResponse.md)
 - [UpdateAccountingProfileRequest](./Models/UpdateAccountingProfileRequest.md)
 - [UpdateAccountingProfileResponse](./Models/UpdateAccountingProfileResponse.md)
 - [UpdateBankAccountRequest](./Models/UpdateBankAccountRequest.md)
 - [UpdateBankAccountResponse](./Models/UpdateBankAccountResponse.md)
 - [UpdateBudgetRequest](./Models/UpdateBudgetRequest.md)
 - [UpdateBudgetResponse](./Models/UpdateBudgetResponse.md)
 - [UpdateFileRequest](./Models/UpdateFileRequest.md)
 - [UpdateFileResponse](./Models/UpdateFileResponse.md)
 - [UpdateFolderRequest](./Models/UpdateFolderRequest.md)
 - [UpdateFolderResponse](./Models/UpdateFolderResponse.md)
 - [UpdateMilestoneRequest](./Models/UpdateMilestoneRequest.md)
 - [UpdateMilestoneResponse](./Models/UpdateMilestoneResponse.md)
 - [UpdateNoteToRecurringTransactionRequest](./Models/UpdateNoteToRecurringTransactionRequest.md)
 - [UpdateNoteToRecurringTransactionResponse](./Models/UpdateNoteToRecurringTransactionResponse.md)
 - [UpdateNoteToSmartGoalRequest](./Models/UpdateNoteToSmartGoalRequest.md)
 - [UpdateNoteToSmartGoalResponse](./Models/UpdateNoteToSmartGoalResponse.md)
 - [UpdateNoteToTransactionRequest](./Models/UpdateNoteToTransactionRequest.md)
 - [UpdateNoteToTransactionResponse](./Models/UpdateNoteToTransactionResponse.md)
 - [UpdatePocketRequest](./Models/UpdatePocketRequest.md)
 - [UpdatePocketResponse](./Models/UpdatePocketResponse.md)
 - [UpdateRecurringTransactionRequest](./Models/UpdateRecurringTransactionRequest.md)
 - [UpdateRecurringTransactionResponse](./Models/UpdateRecurringTransactionResponse.md)
 - [UpdateRoleResponse](./Models/UpdateRoleResponse.md)
 - [UpdateSingleTransactionRequest](./Models/UpdateSingleTransactionRequest.md)
 - [UpdateSingleTransactionResponse](./Models/UpdateSingleTransactionResponse.md)
 - [UpdateSmartGoalRequest](./Models/UpdateSmartGoalRequest.md)
 - [UpdateSmartGoalResponse](./Models/UpdateSmartGoalResponse.md)
 - [UpdateUserProfileRequest](./Models/UpdateUserProfileRequest.md)
 - [UpdateUserProfileResponse](./Models/UpdateUserProfileResponse.md)
 - [UpdateUserRequest](./Models/UpdateUserRequest.md)
 - [UpdateUserResponse](./Models/UpdateUserResponse.md)
 - [UpdateUserV2Request](./Models/UpdateUserV2Request.md)
 - [UpdateUserV2Response](./Models/UpdateUserV2Response.md)
 - [UpdateWorkspaceRequest](./Models/UpdateWorkspaceRequest.md)
 - [UpdateWorkspaceResponse](./Models/UpdateWorkspaceResponse.md)
 - [UserAccount](./Models/UserAccount.md)
 - [UserProfile](./Models/UserProfile.md)
 - [UserTags](./Models/UserTags.md)
 - [ValidationErrorMessageResponse](./Models/ValidationErrorMessageResponse.md)
 - [ValidationErrorMessageResponse1](./Models/ValidationErrorMessageResponse1.md)
 - [VendorCredit](./Models/VendorCredit.md)
 - [VendorCreditLine](./Models/VendorCreditLine.md)
 - [VerifyUserResponse](./Models/VerifyUserResponse.md)
 - [VerifyUserV2Response](./Models/VerifyUserV2Response.md)
 - [Workspace](./Models/Workspace.md)
 - [workspaceservicehttp.FileUploadResponse](./Models/workspaceservicehttp.FileUploadResponse.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

All endpoints do not require authorization.
