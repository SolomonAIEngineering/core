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


/**
 * 
 * @export
 */
export const StripeSubscriptionStatus = {
    Unspecified: 'STRIPE_SUBSCRIPTION_STATUS_UNSPECIFIED',
    Trialing: 'STRIPE_SUBSCRIPTION_STATUS_TRIALING',
    Active: 'STRIPE_SUBSCRIPTION_STATUS_ACTIVE',
    PastDue: 'STRIPE_SUBSCRIPTION_STATUS_PAST_DUE',
    Canceled: 'STRIPE_SUBSCRIPTION_STATUS_CANCELED',
    Unpaid: 'STRIPE_SUBSCRIPTION_STATUS_UNPAID',
    Complete: 'STRIPE_SUBSCRIPTION_STATUS_COMPLETE',
    Incomplete: 'STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE',
    IncompleteExpired: 'STRIPE_SUBSCRIPTION_STATUS_INCOMPLETE_EXPIRED',
    Created: 'STRIPE_SUBSCRIPTION_STATUS_CREATED',
    Paused: 'STRIPE_SUBSCRIPTION_STATUS_PAUSED'
} as const;
export type StripeSubscriptionStatus = typeof StripeSubscriptionStatus[keyof typeof StripeSubscriptionStatus];


export function instanceOfStripeSubscriptionStatus(value: any): boolean {
    return Object.values(StripeSubscriptionStatus).includes(value);
}

export function StripeSubscriptionStatusFromJSON(json: any): StripeSubscriptionStatus {
    return StripeSubscriptionStatusFromJSONTyped(json, false);
}

export function StripeSubscriptionStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): StripeSubscriptionStatus {
    return json as StripeSubscriptionStatus;
}

export function StripeSubscriptionStatusToJSON(value?: StripeSubscriptionStatus | null): any {
    return value as any;
}

