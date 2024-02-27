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
 * Defines a message named GetLinkTokenResponse.
 * @export
 * @interface GetMergeLinkTokenResponse
 */
export interface GetMergeLinkTokenResponse {
    /**
     * A string field named "link_token" with field number 1.
     * @type {string}
     * @memberof GetMergeLinkTokenResponse
     */
    linkToken?: string;
    /**
     * A string field named "integration_name" with field number 2.
     * @type {string}
     * @memberof GetMergeLinkTokenResponse
     */
    integrationName?: string;
    /**
     * A string field named "magic_link_url" with field number 3.
     * @type {string}
     * @memberof GetMergeLinkTokenResponse
     */
    magicLinkUrl?: string;
    /**
     * A string field named "end_user_origin_id" with field number 4.
     * @type {string}
     * @memberof GetMergeLinkTokenResponse
     */
    endUserOriginId?: string;
    /**
     * A string field named "organization_name" with field number 5.
     * @type {string}
     * @memberof GetMergeLinkTokenResponse
     */
    organizationName?: string;
}

/**
 * Check if a given object implements the GetMergeLinkTokenResponse interface.
 */
export function instanceOfGetMergeLinkTokenResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GetMergeLinkTokenResponseFromJSON(json: any): GetMergeLinkTokenResponse {
    return GetMergeLinkTokenResponseFromJSONTyped(json, false);
}

export function GetMergeLinkTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetMergeLinkTokenResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'linkToken': !exists(json, 'linkToken') ? undefined : json['linkToken'],
        'integrationName': !exists(json, 'integrationName') ? undefined : json['integrationName'],
        'magicLinkUrl': !exists(json, 'magicLinkUrl') ? undefined : json['magicLinkUrl'],
        'endUserOriginId': !exists(json, 'endUserOriginId') ? undefined : json['endUserOriginId'],
        'organizationName': !exists(json, 'organizationName') ? undefined : json['organizationName'],
    };
}

export function GetMergeLinkTokenResponseToJSON(value?: GetMergeLinkTokenResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'linkToken': value.linkToken,
        'integrationName': value.integrationName,
        'magicLinkUrl': value.magicLinkUrl,
        'endUserOriginId': value.endUserOriginId,
        'organizationName': value.organizationName,
    };
}

