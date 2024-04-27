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
import type { PostType } from './PostType';
import {
    PostTypeFromJSON,
    PostTypeFromJSONTyped,
    PostTypeToJSON,
} from './PostType';

/**
 * 
 * @export
 * @interface ReportCommentBody
 */
export interface ReportCommentBody {
    /**
     * 
     * @type {PostType}
     * @memberof ReportCommentBody
     */
    postType: PostType;
}

/**
 * Check if a given object implements the ReportCommentBody interface.
 */
export function instanceOfReportCommentBody(value: object): boolean {
    if (!('postType' in value)) return false;
    return true;
}

export function ReportCommentBodyFromJSON(json: any): ReportCommentBody {
    return ReportCommentBodyFromJSONTyped(json, false);
}

export function ReportCommentBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReportCommentBody {
    if (json == null) {
        return json;
    }
    return {
        
        'postType': PostTypeFromJSON(json['postType']),
    };
}

export function ReportCommentBodyToJSON(value?: ReportCommentBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'postType': PostTypeToJSON(value['postType']),
    };
}

