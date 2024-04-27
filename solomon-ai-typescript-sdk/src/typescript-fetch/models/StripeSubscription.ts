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
import type { StripeSubscriptionStatus } from './StripeSubscriptionStatus';
import {
    StripeSubscriptionStatusFromJSON,
    StripeSubscriptionStatusFromJSONTyped,
    StripeSubscriptionStatusToJSON,
} from './StripeSubscriptionStatus';

/**
 * 
 * @export
 * @interface StripeSubscription
 */
export interface StripeSubscription {
    /**
     * 
     * @type {string}
     * @memberof StripeSubscription
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof StripeSubscription
     */
    stripeSubscriptionId?: string;
    /**
     * 
     * @type {StripeSubscriptionStatus}
     * @memberof StripeSubscription
     */
    stripeSubscriptionStatus?: StripeSubscriptionStatus;
    /**
     * 
     * @type {string}
     * @memberof StripeSubscription
     */
    stripeSubscriptionActiveUntil?: string;
    /**
     * 
     * @type {string}
     * @memberof StripeSubscription
     */
    stripeWebhookLatestTimestamp?: string;
    /**
     * 
     * @type {boolean}
     * @memberof StripeSubscription
     */
    isTrialing?: boolean;
}

/**
 * Check if a given object implements the StripeSubscription interface.
 */
export function instanceOfStripeSubscription(value: object): boolean {
    return true;
}

export function StripeSubscriptionFromJSON(json: any): StripeSubscription {
    return StripeSubscriptionFromJSONTyped(json, false);
}

export function StripeSubscriptionFromJSONTyped(json: any, ignoreDiscriminator: boolean): StripeSubscription {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'stripeSubscriptionId': json['stripeSubscriptionId'] == null ? undefined : json['stripeSubscriptionId'],
        'stripeSubscriptionStatus': json['stripeSubscriptionStatus'] == null ? undefined : StripeSubscriptionStatusFromJSON(json['stripeSubscriptionStatus']),
        'stripeSubscriptionActiveUntil': json['stripeSubscriptionActiveUntil'] == null ? undefined : json['stripeSubscriptionActiveUntil'],
        'stripeWebhookLatestTimestamp': json['stripeWebhookLatestTimestamp'] == null ? undefined : json['stripeWebhookLatestTimestamp'],
        'isTrialing': json['isTrialing'] == null ? undefined : json['isTrialing'],
    };
}

export function StripeSubscriptionToJSON(value?: StripeSubscription | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'stripeSubscriptionId': value['stripeSubscriptionId'],
        'stripeSubscriptionStatus': StripeSubscriptionStatusToJSON(value['stripeSubscriptionStatus']),
        'stripeSubscriptionActiveUntil': value['stripeSubscriptionActiveUntil'],
        'stripeWebhookLatestTimestamp': value['stripeWebhookLatestTimestamp'],
        'isTrialing': value['isTrialing'],
    };
}

