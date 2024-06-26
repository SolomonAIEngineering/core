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
export const PersonalActionableInsightName = {
    Unspecified: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_UNSPECIFIED',
    Expense: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_EXPENSE',
    EmergencyFund: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_EMERGENCY_FUND',
    DebtPrioritization: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_DEBT_PRIORITIZATION',
    NonEssentialExpenses: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_NON_ESSENTIAL_EXPENSES',
    NonSubscriptions: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_NON_SUBSCRIPTIONS',
    DiscretionarySpending: 'PERSONAL_ACTIONABLE_INSIGHT_NAME_DISCRETIONARY_SPENDING'
} as const;
export type PersonalActionableInsightName = typeof PersonalActionableInsightName[keyof typeof PersonalActionableInsightName];


export function instanceOfPersonalActionableInsightName(value: any): boolean {
    return Object.values(PersonalActionableInsightName).includes(value);
}

export function PersonalActionableInsightNameFromJSON(json: any): PersonalActionableInsightName {
    return PersonalActionableInsightNameFromJSONTyped(json, false);
}

export function PersonalActionableInsightNameFromJSONTyped(json: any, ignoreDiscriminator: boolean): PersonalActionableInsightName {
    return json as PersonalActionableInsightName;
}

export function PersonalActionableInsightNameToJSON(value?: PersonalActionableInsightName | null): any {
    return value as any;
}

