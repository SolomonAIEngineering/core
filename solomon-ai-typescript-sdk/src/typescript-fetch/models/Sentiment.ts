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
 * @interface Sentiment
 */
export interface Sentiment {
    /**
     * 
     * @type {number}
     * @memberof Sentiment
     */
    negative?: number;
    /**
     * 
     * @type {number}
     * @memberof Sentiment
     */
    neutral?: number;
    /**
     * 
     * @type {number}
     * @memberof Sentiment
     */
    positive?: number;
    /**
     * 
     * @type {number}
     * @memberof Sentiment
     */
    compound?: number;
}

/**
 * Check if a given object implements the Sentiment interface.
 */
export function instanceOfSentiment(value: object): boolean {
    return true;
}

export function SentimentFromJSON(json: any): Sentiment {
    return SentimentFromJSONTyped(json, false);
}

export function SentimentFromJSONTyped(json: any, ignoreDiscriminator: boolean): Sentiment {
    if (json == null) {
        return json;
    }
    return {
        
        'negative': json['negative'] == null ? undefined : json['negative'],
        'neutral': json['neutral'] == null ? undefined : json['neutral'],
        'positive': json['positive'] == null ? undefined : json['positive'],
        'compound': json['compound'] == null ? undefined : json['compound'],
    };
}

export function SentimentToJSON(value?: Sentiment | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'negative': value['negative'],
        'neutral': value['neutral'],
        'positive': value['positive'],
        'compound': value['compound'],
    };
}

