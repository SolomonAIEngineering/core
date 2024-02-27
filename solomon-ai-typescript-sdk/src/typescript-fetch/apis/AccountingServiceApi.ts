/* tslint:disable */
/* eslint-disable */
/**
 * User Service
 * Solomon AI User Service API
 *
 * The version of the OpenAPI document: 0.1
 * Contact: yoanyomba@solomon-ai.co
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  CreateAccountingProfileRequest,
  CreateAccountingProfileResponse,
  DeleteAccountingProfileResponse,
  ExchangePublicLinkTokenForAccountTokenRequest,
  ExchangePublicLinkTokenForAccountTokenResponse,
  GetMergeLinkTokenRequest,
  GetMergeLinkTokenResponse,
  HealthCheckResponse2,
  InternalErrorMessageResponse,
  PathUnknownErrorMessageResponse,
  ReadAccountingProfileResponse,
  ReadBalanceSheetsRequest,
  ReadBalanceSheetsResponse,
  ReadBusinessChartOfAccountsRequest,
  ReadBusinessChartOfAccountsResponse,
  ReadCashFlowStatementsRequest,
  ReadCashFlowStatementsResponse,
  ReadIncomeStatementsRequest,
  ReadIncomeStatementsResponse,
  ReadynessCheckResponse2,
  Status,
  UpdateAccountingProfileRequest,
  UpdateAccountingProfileResponse,
  ValidationErrorMessageResponse1,
} from '../models/index';
import {
    CreateAccountingProfileRequestFromJSON,
    CreateAccountingProfileRequestToJSON,
    CreateAccountingProfileResponseFromJSON,
    CreateAccountingProfileResponseToJSON,
    DeleteAccountingProfileResponseFromJSON,
    DeleteAccountingProfileResponseToJSON,
    ExchangePublicLinkTokenForAccountTokenRequestFromJSON,
    ExchangePublicLinkTokenForAccountTokenRequestToJSON,
    ExchangePublicLinkTokenForAccountTokenResponseFromJSON,
    ExchangePublicLinkTokenForAccountTokenResponseToJSON,
    GetMergeLinkTokenRequestFromJSON,
    GetMergeLinkTokenRequestToJSON,
    GetMergeLinkTokenResponseFromJSON,
    GetMergeLinkTokenResponseToJSON,
    HealthCheckResponse2FromJSON,
    HealthCheckResponse2ToJSON,
    InternalErrorMessageResponseFromJSON,
    InternalErrorMessageResponseToJSON,
    PathUnknownErrorMessageResponseFromJSON,
    PathUnknownErrorMessageResponseToJSON,
    ReadAccountingProfileResponseFromJSON,
    ReadAccountingProfileResponseToJSON,
    ReadBalanceSheetsRequestFromJSON,
    ReadBalanceSheetsRequestToJSON,
    ReadBalanceSheetsResponseFromJSON,
    ReadBalanceSheetsResponseToJSON,
    ReadBusinessChartOfAccountsRequestFromJSON,
    ReadBusinessChartOfAccountsRequestToJSON,
    ReadBusinessChartOfAccountsResponseFromJSON,
    ReadBusinessChartOfAccountsResponseToJSON,
    ReadCashFlowStatementsRequestFromJSON,
    ReadCashFlowStatementsRequestToJSON,
    ReadCashFlowStatementsResponseFromJSON,
    ReadCashFlowStatementsResponseToJSON,
    ReadIncomeStatementsRequestFromJSON,
    ReadIncomeStatementsRequestToJSON,
    ReadIncomeStatementsResponseFromJSON,
    ReadIncomeStatementsResponseToJSON,
    ReadynessCheckResponse2FromJSON,
    ReadynessCheckResponse2ToJSON,
    StatusFromJSON,
    StatusToJSON,
    UpdateAccountingProfileRequestFromJSON,
    UpdateAccountingProfileRequestToJSON,
    UpdateAccountingProfileResponseFromJSON,
    UpdateAccountingProfileResponseToJSON,
    ValidationErrorMessageResponse1FromJSON,
    ValidationErrorMessageResponse1ToJSON,
} from '../models/index';

export interface CreatePayrollProfileRequest {
    createAccountingProfileRequest: CreateAccountingProfileRequest;
}

export interface DeleteProfileRequest {
    authZeroUserId: string;
}

export interface ExchangePublicLinkTokenForAccountTokenResponseRequest {
    exchangePublicLinkTokenForAccountTokenRequest: ExchangePublicLinkTokenForAccountTokenRequest;
}

export interface GetLinkTokenRequest {
    getMergeLinkTokenRequest: GetMergeLinkTokenRequest;
}

export interface ReadAccountingProfileResponseRequest {
    authZeroUserId: string;
    linkedAccountingAccountId: string;
}

export interface ReadBalanceSheetsOperationRequest {
    readBalanceSheetsRequest: ReadBalanceSheetsRequest;
}

export interface ReadBusinessChartOfAccountsOperationRequest {
    readBusinessChartOfAccountsRequest: ReadBusinessChartOfAccountsRequest;
}

export interface ReadCashFlowStatementsOperationRequest {
    readCashFlowStatementsRequest: ReadCashFlowStatementsRequest;
}

export interface ReadIncomeStatementsOperationRequest {
    readIncomeStatementsRequest: ReadIncomeStatementsRequest;
}

export interface UpdatePayrollProfileRequest {
    updateAccountingProfileRequest: UpdateAccountingProfileRequest;
}

/**
 * 
 */
export class AccountingServiceApi extends runtime.BaseAPI {

    /**
     * Creates a payroll profile.
     * Create Payroll Profile
     */
    async createPayrollProfileRaw(requestParameters: CreatePayrollProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateAccountingProfileResponse>> {
        if (requestParameters.createAccountingProfileRequest === null || requestParameters.createAccountingProfileRequest === undefined) {
            throw new runtime.RequiredError('createAccountingProfileRequest','Required parameter requestParameters.createAccountingProfileRequest was null or undefined when calling createPayrollProfile.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/profile`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateAccountingProfileRequestToJSON(requestParameters.createAccountingProfileRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateAccountingProfileResponseFromJSON(jsonValue));
    }

    /**
     * Creates a payroll profile.
     * Create Payroll Profile
     */
    async createPayrollProfile(requestParameters: CreatePayrollProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateAccountingProfileResponse> {
        const response = await this.createPayrollProfileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Deletes a business payroll profile.
     * Delete Payroll Profile
     */
    async deleteProfileRaw(requestParameters: DeleteProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DeleteAccountingProfileResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling deleteProfile.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/accounting-microservice/api/v1/profile/{authZeroUserId}`.replace(`{${"authZeroUserId"}}`, encodeURIComponent(String(requestParameters.authZeroUserId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DeleteAccountingProfileResponseFromJSON(jsonValue));
    }

    /**
     * Deletes a business payroll profile.
     * Delete Payroll Profile
     */
    async deleteProfile(requestParameters: DeleteProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DeleteAccountingProfileResponse> {
        const response = await this.deleteProfileRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Exchanges a public link token for an account token.
     * Exchange Link Token
     */
    async exchangePublicLinkTokenForAccountTokenResponseRaw(requestParameters: ExchangePublicLinkTokenForAccountTokenResponseRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ExchangePublicLinkTokenForAccountTokenResponse>> {
        if (requestParameters.exchangePublicLinkTokenForAccountTokenRequest === null || requestParameters.exchangePublicLinkTokenForAccountTokenRequest === undefined) {
            throw new runtime.RequiredError('exchangePublicLinkTokenForAccountTokenRequest','Required parameter requestParameters.exchangePublicLinkTokenForAccountTokenRequest was null or undefined when calling exchangePublicLinkTokenForAccountTokenResponse.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/merge/exchange-token`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ExchangePublicLinkTokenForAccountTokenRequestToJSON(requestParameters.exchangePublicLinkTokenForAccountTokenRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ExchangePublicLinkTokenForAccountTokenResponseFromJSON(jsonValue));
    }

    /**
     * Exchanges a public link token for an account token.
     * Exchange Link Token
     */
    async exchangePublicLinkTokenForAccountTokenResponse(requestParameters: ExchangePublicLinkTokenForAccountTokenResponseRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ExchangePublicLinkTokenForAccountTokenResponse> {
        const response = await this.exchangePublicLinkTokenForAccountTokenResponseRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Provides a link token to initialize a Link session for the user.
     * Get Link Token
     */
    async getLinkTokenRaw(requestParameters: GetLinkTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetMergeLinkTokenResponse>> {
        if (requestParameters.getMergeLinkTokenRequest === null || requestParameters.getMergeLinkTokenRequest === undefined) {
            throw new runtime.RequiredError('getMergeLinkTokenRequest','Required parameter requestParameters.getMergeLinkTokenRequest was null or undefined when calling getLinkToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/merge/initiate-token-exchange`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: GetMergeLinkTokenRequestToJSON(requestParameters.getMergeLinkTokenRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetMergeLinkTokenResponseFromJSON(jsonValue));
    }

    /**
     * Provides a link token to initialize a Link session for the user.
     * Get Link Token
     */
    async getLinkToken(requestParameters: GetLinkTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetMergeLinkTokenResponse> {
        const response = await this.getLinkTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Performs a health check on the service.
     * Health Check
     */
    async healthCheck3Raw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<HealthCheckResponse2>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/accounting-microservice/api/v1/health`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => HealthCheckResponse2FromJSON(jsonValue));
    }

    /**
     * Performs a health check on the service.
     * Health Check
     */
    async healthCheck3(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<HealthCheckResponse2> {
        const response = await this.healthCheck3Raw(initOverrides);
        return await response.value();
    }

    /**
     * Retrieves a business account profile.
     * Get Business Account Profile
     */
    async readAccountingProfileResponseRaw(requestParameters: ReadAccountingProfileResponseRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadAccountingProfileResponse>> {
        if (requestParameters.authZeroUserId === null || requestParameters.authZeroUserId === undefined) {
            throw new runtime.RequiredError('authZeroUserId','Required parameter requestParameters.authZeroUserId was null or undefined when calling readAccountingProfileResponse.');
        }

        if (requestParameters.linkedAccountingAccountId === null || requestParameters.linkedAccountingAccountId === undefined) {
            throw new runtime.RequiredError('linkedAccountingAccountId','Required parameter requestParameters.linkedAccountingAccountId was null or undefined when calling readAccountingProfileResponse.');
        }

        const queryParameters: any = {};

        if (requestParameters.linkedAccountingAccountId !== undefined) {
            queryParameters['linkedAccountingAccountId'] = requestParameters.linkedAccountingAccountId;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/accounting-microservice/api/v1/profile/{authZeroUserId}`.replace(`{${"authZeroUserId"}}`, encodeURIComponent(String(requestParameters.authZeroUserId))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadAccountingProfileResponseFromJSON(jsonValue));
    }

    /**
     * Retrieves a business account profile.
     * Get Business Account Profile
     */
    async readAccountingProfileResponse(requestParameters: ReadAccountingProfileResponseRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadAccountingProfileResponse> {
        const response = await this.readAccountingProfileResponseRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Queries balance sheets for a business.
     * Gets Balance Sheets
     */
    async readBalanceSheetsRaw(requestParameters: ReadBalanceSheetsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadBalanceSheetsResponse>> {
        if (requestParameters.readBalanceSheetsRequest === null || requestParameters.readBalanceSheetsRequest === undefined) {
            throw new runtime.RequiredError('readBalanceSheetsRequest','Required parameter requestParameters.readBalanceSheetsRequest was null or undefined when calling readBalanceSheets.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/balance-sheets`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReadBalanceSheetsRequestToJSON(requestParameters.readBalanceSheetsRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadBalanceSheetsResponseFromJSON(jsonValue));
    }

    /**
     * Queries balance sheets for a business.
     * Gets Balance Sheets
     */
    async readBalanceSheets(requestParameters: ReadBalanceSheetsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadBalanceSheetsResponse> {
        const response = await this.readBalanceSheetsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Queries chart of accounts for a business.
     * Gets Chart of Accounts
     */
    async readBusinessChartOfAccountsRaw(requestParameters: ReadBusinessChartOfAccountsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadBusinessChartOfAccountsResponse>> {
        if (requestParameters.readBusinessChartOfAccountsRequest === null || requestParameters.readBusinessChartOfAccountsRequest === undefined) {
            throw new runtime.RequiredError('readBusinessChartOfAccountsRequest','Required parameter requestParameters.readBusinessChartOfAccountsRequest was null or undefined when calling readBusinessChartOfAccounts.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/chart-of-accounts`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReadBusinessChartOfAccountsRequestToJSON(requestParameters.readBusinessChartOfAccountsRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadBusinessChartOfAccountsResponseFromJSON(jsonValue));
    }

    /**
     * Queries chart of accounts for a business.
     * Gets Chart of Accounts
     */
    async readBusinessChartOfAccounts(requestParameters: ReadBusinessChartOfAccountsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadBusinessChartOfAccountsResponse> {
        const response = await this.readBusinessChartOfAccountsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Queries cashflow sheets for a business.
     * Gets Cashfloe Sheets
     */
    async readCashFlowStatementsRaw(requestParameters: ReadCashFlowStatementsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadCashFlowStatementsResponse>> {
        if (requestParameters.readCashFlowStatementsRequest === null || requestParameters.readCashFlowStatementsRequest === undefined) {
            throw new runtime.RequiredError('readCashFlowStatementsRequest','Required parameter requestParameters.readCashFlowStatementsRequest was null or undefined when calling readCashFlowStatements.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/cashflow-statements`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReadCashFlowStatementsRequestToJSON(requestParameters.readCashFlowStatementsRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadCashFlowStatementsResponseFromJSON(jsonValue));
    }

    /**
     * Queries cashflow sheets for a business.
     * Gets Cashfloe Sheets
     */
    async readCashFlowStatements(requestParameters: ReadCashFlowStatementsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadCashFlowStatementsResponse> {
        const response = await this.readCashFlowStatementsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Queries income sheets for a business.
     * Gets Cashfloe Sheets
     */
    async readIncomeStatementsRaw(requestParameters: ReadIncomeStatementsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadIncomeStatementsResponse>> {
        if (requestParameters.readIncomeStatementsRequest === null || requestParameters.readIncomeStatementsRequest === undefined) {
            throw new runtime.RequiredError('readIncomeStatementsRequest','Required parameter requestParameters.readIncomeStatementsRequest was null or undefined when calling readIncomeStatements.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/income-statements`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReadIncomeStatementsRequestToJSON(requestParameters.readIncomeStatementsRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadIncomeStatementsResponseFromJSON(jsonValue));
    }

    /**
     * Queries income sheets for a business.
     * Gets Cashfloe Sheets
     */
    async readIncomeStatements(requestParameters: ReadIncomeStatementsOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadIncomeStatementsResponse> {
        const response = await this.readIncomeStatementsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Performs a readiness check on the service.
     * Readiness Check
     */
    async readinessCheckRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ReadynessCheckResponse2>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/accounting-microservice/api/v1/ready`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ReadynessCheckResponse2FromJSON(jsonValue));
    }

    /**
     * Performs a readiness check on the service.
     * Readiness Check
     */
    async readinessCheck(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ReadynessCheckResponse2> {
        const response = await this.readinessCheckRaw(initOverrides);
        return await response.value();
    }

    /**
     * Updates a payroll profile.
     * Update Payroll Profile
     */
    async updatePayrollProfileRaw(requestParameters: UpdatePayrollProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UpdateAccountingProfileResponse>> {
        if (requestParameters.updateAccountingProfileRequest === null || requestParameters.updateAccountingProfileRequest === undefined) {
            throw new runtime.RequiredError('updateAccountingProfileRequest','Required parameter requestParameters.updateAccountingProfileRequest was null or undefined when calling updatePayrollProfile.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/accounting-microservice/api/v1/profile`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateAccountingProfileRequestToJSON(requestParameters.updateAccountingProfileRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UpdateAccountingProfileResponseFromJSON(jsonValue));
    }

    /**
     * Updates a payroll profile.
     * Update Payroll Profile
     */
    async updatePayrollProfile(requestParameters: UpdatePayrollProfileRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UpdateAccountingProfileResponse> {
        const response = await this.updatePayrollProfileRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
