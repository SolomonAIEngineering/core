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

import { exists, mapValues } from '../runtime';
/**
 * Defines a message named ExchangePublicLinkTokenForAccountTokenResponse.
 * @export
 * @interface ExchangePublicLinkTokenForAccountTokenResponse
 */
export interface ExchangePublicLinkTokenForAccountTokenResponse {
    /**
     * 
     * @type {boolean}
     * @memberof ExchangePublicLinkTokenForAccountTokenResponse
     */
    success: boolean;
    /**
     * A string field named "workflow_id" with field number 2.
     * @type {string}
     * @memberof ExchangePublicLinkTokenForAccountTokenResponse
     */
    workflowId?: string;
    /**
     * A string field named "workflow_run_id" with field number 3.
     * @type {string}
     * @memberof ExchangePublicLinkTokenForAccountTokenResponse
     */
    workflowRunId?: string;
}

/**
 * Check if a given object implements the ExchangePublicLinkTokenForAccountTokenResponse interface.
 */
export function instanceOfExchangePublicLinkTokenForAccountTokenResponse(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "success" in value;

    return isInstance;
}

export function ExchangePublicLinkTokenForAccountTokenResponseFromJSON(json: any): ExchangePublicLinkTokenForAccountTokenResponse {
    return ExchangePublicLinkTokenForAccountTokenResponseFromJSONTyped(json, false);
}

export function ExchangePublicLinkTokenForAccountTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ExchangePublicLinkTokenForAccountTokenResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'success': json['success'],
        'workflowId': !exists(json, 'workflowId') ? undefined : json['workflowId'],
        'workflowRunId': !exists(json, 'workflowRunId') ? undefined : json['workflowRunId'],
    };
}

export function ExchangePublicLinkTokenForAccountTokenResponseToJSON(value?: ExchangePublicLinkTokenForAccountTokenResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'success': value.success,
        'workflowId': value.workflowId,
        'workflowRunId': value.workflowRunId,
    };
}

