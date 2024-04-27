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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface InvesmentHolding
 */
export interface InvesmentHolding {
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    plaidAccountId?: string;
    /**
     * 
     * @type {number}
     * @memberof InvesmentHolding
     */
    costBasis?: number;
    /**
     * 
     * @type {number}
     * @memberof InvesmentHolding
     */
    institutionPrice?: number;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    institutionPriceAsOf?: string;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    institutionPriceDatetime?: string;
    /**
     * 
     * @type {number}
     * @memberof InvesmentHolding
     */
    institutionValue?: number;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    isoCurrencyCode?: string;
    /**
     * 
     * @type {number}
     * @memberof InvesmentHolding
     */
    quantity?: number;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    securityId?: string;
    /**
     * 
     * @type {string}
     * @memberof InvesmentHolding
     */
    unofficialCurrencyCode?: string;
}

/**
 * Check if a given object implements the InvesmentHolding interface.
 */
export function instanceOfInvesmentHolding(value: object): boolean {
    return true;
}

export function InvesmentHoldingFromJSON(json: any): InvesmentHolding {
    return InvesmentHoldingFromJSONTyped(json, false);
}

export function InvesmentHoldingFromJSONTyped(json: any, ignoreDiscriminator: boolean): InvesmentHolding {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'plaidAccountId': json['plaidAccountId'] == null ? undefined : json['plaidAccountId'],
        'costBasis': json['costBasis'] == null ? undefined : json['costBasis'],
        'institutionPrice': json['institutionPrice'] == null ? undefined : json['institutionPrice'],
        'institutionPriceAsOf': json['institutionPriceAsOf'] == null ? undefined : json['institutionPriceAsOf'],
        'institutionPriceDatetime': json['institutionPriceDatetime'] == null ? undefined : json['institutionPriceDatetime'],
        'institutionValue': json['institutionValue'] == null ? undefined : json['institutionValue'],
        'isoCurrencyCode': json['isoCurrencyCode'] == null ? undefined : json['isoCurrencyCode'],
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
        'securityId': json['securityId'] == null ? undefined : json['securityId'],
        'unofficialCurrencyCode': json['unofficialCurrencyCode'] == null ? undefined : json['unofficialCurrencyCode'],
    };
}

export function InvesmentHoldingToJSON(value?: InvesmentHolding | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'name': value['name'],
        'plaidAccountId': value['plaidAccountId'],
        'costBasis': value['costBasis'],
        'institutionPrice': value['institutionPrice'],
        'institutionPriceAsOf': value['institutionPriceAsOf'],
        'institutionPriceDatetime': value['institutionPriceDatetime'],
        'institutionValue': value['institutionValue'],
        'isoCurrencyCode': value['isoCurrencyCode'],
        'quantity': value['quantity'],
        'securityId': value['securityId'],
        'unofficialCurrencyCode': value['unofficialCurrencyCode'],
    };
}

