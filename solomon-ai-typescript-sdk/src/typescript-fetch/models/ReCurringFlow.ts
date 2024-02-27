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
export const ReCurringFlow = {
    Unspecified: 'RE_CURRING_FLOW_UNSPECIFIED',
    Inflow: 'RE_CURRING_FLOW_INFLOW',
    Outflow: 'RE_CURRING_FLOW_OUTFLOW'
} as const;
export type ReCurringFlow = typeof ReCurringFlow[keyof typeof ReCurringFlow];


export function ReCurringFlowFromJSON(json: any): ReCurringFlow {
    return ReCurringFlowFromJSONTyped(json, false);
}

export function ReCurringFlowFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReCurringFlow {
    return json as ReCurringFlow;
}

export function ReCurringFlowToJSON(value?: ReCurringFlow | null): any {
    return value as any;
}

