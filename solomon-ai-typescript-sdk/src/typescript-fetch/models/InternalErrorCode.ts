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
export const InternalErrorCode = {
    NoInternalError: 'no_internal_error',
    InternalError: 'internal_error',
    Cancelled: 'cancelled',
    DeadlineExceeded: 'deadline_exceeded',
    AlreadyExists: 'already_exists',
    ResourceExhausted: 'resource_exhausted',
    FailedPrecondition: 'failed_precondition',
    Aborted: 'aborted',
    OutOfRange: 'out_of_range',
    Unavailable: 'unavailable',
    DataLoss: 'data_loss'
} as const;
export type InternalErrorCode = typeof InternalErrorCode[keyof typeof InternalErrorCode];


export function instanceOfInternalErrorCode(value: any): boolean {
    return Object.values(InternalErrorCode).includes(value);
}

export function InternalErrorCodeFromJSON(json: any): InternalErrorCode {
    return InternalErrorCodeFromJSONTyped(json, false);
}

export function InternalErrorCodeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InternalErrorCode {
    return json as InternalErrorCode;
}

export function InternalErrorCodeToJSON(value?: InternalErrorCode | null): any {
    return value as any;
}

