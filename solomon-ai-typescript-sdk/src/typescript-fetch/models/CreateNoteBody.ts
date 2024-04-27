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
import type { Note } from './Note';
import {
    NoteFromJSON,
    NoteFromJSONTyped,
    NoteToJSON,
} from './Note';
import type { PostType } from './PostType';
import {
    PostTypeFromJSON,
    PostTypeFromJSONTyped,
    PostTypeToJSON,
} from './PostType';

/**
 * 
 * @export
 * @interface CreateNoteBody
 */
export interface CreateNoteBody {
    /**
     * 
     * @type {PostType}
     * @memberof CreateNoteBody
     */
    postType: PostType;
    /**
     * 
     * @type {Note}
     * @memberof CreateNoteBody
     */
    note?: Note;
}

/**
 * Check if a given object implements the CreateNoteBody interface.
 */
export function instanceOfCreateNoteBody(value: object): boolean {
    if (!('postType' in value)) return false;
    return true;
}

export function CreateNoteBodyFromJSON(json: any): CreateNoteBody {
    return CreateNoteBodyFromJSONTyped(json, false);
}

export function CreateNoteBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateNoteBody {
    if (json == null) {
        return json;
    }
    return {
        
        'postType': PostTypeFromJSON(json['postType']),
        'note': json['note'] == null ? undefined : NoteFromJSON(json['note']),
    };
}

export function CreateNoteBodyToJSON(value?: CreateNoteBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'postType': PostTypeToJSON(value['postType']),
        'note': NoteToJSON(value['note']),
    };
}

