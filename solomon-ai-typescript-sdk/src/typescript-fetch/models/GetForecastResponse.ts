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
import type { Forecast } from './Forecast';
import {
    ForecastFromJSON,
    ForecastFromJSONTyped,
    ForecastToJSON,
} from './Forecast';

/**
 * 
 * @export
 * @interface GetForecastResponse
 */
export interface GetForecastResponse {
    /**
     * 
     * @type {Forecast}
     * @memberof GetForecastResponse
     */
    forecast?: Forecast;
}

/**
 * Check if a given object implements the GetForecastResponse interface.
 */
export function instanceOfGetForecastResponse(value: object): boolean {
    return true;
}

export function GetForecastResponseFromJSON(json: any): GetForecastResponse {
    return GetForecastResponseFromJSONTyped(json, false);
}

export function GetForecastResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetForecastResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'forecast': json['forecast'] == null ? undefined : ForecastFromJSON(json['forecast']),
    };
}

export function GetForecastResponseToJSON(value?: GetForecastResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'forecast': ForecastToJSON(value['forecast']),
    };
}

