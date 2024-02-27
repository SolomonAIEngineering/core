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
export const NotificationType = {
    Unspecified: 'NOTIFICATION_TYPE_UNSPECIFIED',
    Email: 'NOTIFICATION_TYPE_EMAIL',
    Sms: 'NOTIFICATION_TYPE_SMS',
    InApp: 'NOTIFICATION_TYPE_IN_APP',
    Slack: 'NOTIFICATION_TYPE_SLACK'
} as const;
export type NotificationType = typeof NotificationType[keyof typeof NotificationType];


export function NotificationTypeFromJSON(json: any): NotificationType {
    return NotificationTypeFromJSONTyped(json, false);
}

export function NotificationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NotificationType {
    return json as NotificationType;
}

export function NotificationTypeToJSON(value?: NotificationType | null): any {
    return value as any;
}

