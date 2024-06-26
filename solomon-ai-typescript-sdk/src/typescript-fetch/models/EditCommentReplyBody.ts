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
import type { CommentReply } from './CommentReply';
import {
    CommentReplyFromJSON,
    CommentReplyFromJSONTyped,
    CommentReplyToJSON,
} from './CommentReply';
import type { PostType } from './PostType';
import {
    PostTypeFromJSON,
    PostTypeFromJSONTyped,
    PostTypeToJSON,
} from './PostType';

/**
 * 
 * @export
 * @interface EditCommentReplyBody
 */
export interface EditCommentReplyBody {
    /**
     * 
     * @type {CommentReply}
     * @memberof EditCommentReplyBody
     */
    reply?: CommentReply;
    /**
     * 
     * @type {PostType}
     * @memberof EditCommentReplyBody
     */
    postType: PostType;
}

/**
 * Check if a given object implements the EditCommentReplyBody interface.
 */
export function instanceOfEditCommentReplyBody(value: object): boolean {
    if (!('postType' in value)) return false;
    return true;
}

export function EditCommentReplyBodyFromJSON(json: any): EditCommentReplyBody {
    return EditCommentReplyBodyFromJSONTyped(json, false);
}

export function EditCommentReplyBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): EditCommentReplyBody {
    if (json == null) {
        return json;
    }
    return {
        
        'reply': json['reply'] == null ? undefined : CommentReplyFromJSON(json['reply']),
        'postType': PostTypeFromJSON(json['postType']),
    };
}

export function EditCommentReplyBodyToJSON(value?: EditCommentReplyBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'reply': CommentReplyToJSON(value['reply']),
        'postType': PostTypeToJSON(value['postType']),
    };
}

