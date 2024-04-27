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
import type { Entities } from './Entities';
import {
    EntitiesFromJSON,
    EntitiesFromJSONTyped,
    EntitiesToJSON,
} from './Entities';
import type { Sentiment } from './Sentiment';
import {
    SentimentFromJSON,
    SentimentFromJSONTyped,
    SentimentToJSON,
} from './Sentiment';

/**
 * 
 * @export
 * @interface ContentInsights
 */
export interface ContentInsights {
    /**
     * 
     * @type {string}
     * @memberof ContentInsights
     */
    sentenceCount?: string;
    /**
     * 
     * @type {string}
     * @memberof ContentInsights
     */
    wordCount?: string;
    /**
     * 
     * @type {string}
     * @memberof ContentInsights
     */
    language?: string;
    /**
     * 
     * @type {number}
     * @memberof ContentInsights
     */
    languageConfidence?: number;
    /**
     * 
     * @type {Array<Entities>}
     * @memberof ContentInsights
     */
    entities?: Array<Entities>;
    /**
     * 
     * @type {Sentiment}
     * @memberof ContentInsights
     */
    sentiment?: Sentiment;
}

/**
 * Check if a given object implements the ContentInsights interface.
 */
export function instanceOfContentInsights(value: object): boolean {
    return true;
}

export function ContentInsightsFromJSON(json: any): ContentInsights {
    return ContentInsightsFromJSONTyped(json, false);
}

export function ContentInsightsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContentInsights {
    if (json == null) {
        return json;
    }
    return {
        
        'sentenceCount': json['sentenceCount'] == null ? undefined : json['sentenceCount'],
        'wordCount': json['wordCount'] == null ? undefined : json['wordCount'],
        'language': json['language'] == null ? undefined : json['language'],
        'languageConfidence': json['languageConfidence'] == null ? undefined : json['languageConfidence'],
        'entities': json['entities'] == null ? undefined : ((json['entities'] as Array<any>).map(EntitiesFromJSON)),
        'sentiment': json['sentiment'] == null ? undefined : SentimentFromJSON(json['sentiment']),
    };
}

export function ContentInsightsToJSON(value?: ContentInsights | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'sentenceCount': value['sentenceCount'],
        'wordCount': value['wordCount'],
        'language': value['language'],
        'languageConfidence': value['languageConfidence'],
        'entities': value['entities'] == null ? undefined : ((value['entities'] as Array<any>).map(EntitiesToJSON)),
        'sentiment': SentimentToJSON(value['sentiment']),
    };
}

