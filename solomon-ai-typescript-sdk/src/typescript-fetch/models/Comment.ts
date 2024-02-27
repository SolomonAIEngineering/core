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

import { exists, mapValues } from '../runtime';
import type { AccountType } from './AccountType';
import {
    AccountTypeFromJSON,
    AccountTypeFromJSONTyped,
    AccountTypeToJSON,
} from './AccountType';
import type { CommentReply } from './CommentReply';
import {
    CommentReplyFromJSON,
    CommentReplyFromJSONTyped,
    CommentReplyToJSON,
} from './CommentReply';
import type { Media } from './Media';
import {
    MediaFromJSON,
    MediaFromJSONTyped,
    MediaToJSON,
} from './Media';
import type { Note } from './Note';
import {
    NoteFromJSON,
    NoteFromJSONTyped,
    NoteToJSON,
} from './Note';
import type { Reaction } from './Reaction';
import {
    ReactionFromJSON,
    ReactionFromJSONTyped,
    ReactionToJSON,
} from './Reaction';

/**
 * 
 * @export
 * @interface Comment
 */
export interface Comment {
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    backendPlatformUserId?: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    profileId?: string;
    /**
     * 
     * @type {Media}
     * @memberof Comment
     */
    media?: Media;
    /**
     * 
     * @type {Array<string>}
     * @memberof Comment
     */
    mentions?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof Comment
     */
    hashtags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    content?: string;
    /**
     * 
     * @type {Array<CommentReply>}
     * @memberof Comment
     */
    replies?: Array<CommentReply>;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof Comment
     */
    extra?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    authorUsername: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    authorProfileImage: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    affinityScore?: string;
    /**
     * 
     * @type {string}
     * @memberof Comment
     */
    qualityScore?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof Comment
     */
    userIdToAffinityScoreMap?: { [key: string]: string; };
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof Comment
     */
    userIdToReportsMap?: { [key: string]: string; };
    /**
     * 
     * @type {AccountType}
     * @memberof Comment
     */
    authorAccountType?: AccountType;
    /**
     * 
     * @type {{ [key: string]: Reaction; }}
     * @memberof Comment
     */
    userIdToReactionMap?: { [key: string]: Reaction; };
    /**
     * 
     * @type {Array<Note>}
     * @memberof Comment
     */
    notes?: Array<Note>;
}

/**
 * Check if a given object implements the Comment interface.
 */
export function instanceOfComment(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "authorUsername" in value;
    isInstance = isInstance && "authorProfileImage" in value;

    return isInstance;
}

export function CommentFromJSON(json: any): Comment {
    return CommentFromJSONTyped(json, false);
}

export function CommentFromJSONTyped(json: any, ignoreDiscriminator: boolean): Comment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'backendPlatformUserId': !exists(json, 'backendPlatformUserId') ? undefined : json['backendPlatformUserId'],
        'profileId': !exists(json, 'profileId') ? undefined : json['profileId'],
        'media': !exists(json, 'media') ? undefined : MediaFromJSON(json['media']),
        'mentions': !exists(json, 'mentions') ? undefined : json['mentions'],
        'hashtags': !exists(json, 'hashtags') ? undefined : json['hashtags'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'content': !exists(json, 'content') ? undefined : json['content'],
        'replies': !exists(json, 'replies') ? undefined : ((json['replies'] as Array<any>).map(CommentReplyFromJSON)),
        'extra': !exists(json, 'extra') ? undefined : json['extra'],
        'authorUsername': json['authorUsername'],
        'authorProfileImage': json['authorProfileImage'],
        'affinityScore': !exists(json, 'affinityScore') ? undefined : json['affinityScore'],
        'qualityScore': !exists(json, 'qualityScore') ? undefined : json['qualityScore'],
        'userIdToAffinityScoreMap': !exists(json, 'userIdToAffinityScoreMap') ? undefined : json['userIdToAffinityScoreMap'],
        'userIdToReportsMap': !exists(json, 'userIdToReportsMap') ? undefined : json['userIdToReportsMap'],
        'authorAccountType': !exists(json, 'authorAccountType') ? undefined : AccountTypeFromJSON(json['authorAccountType']),
        'userIdToReactionMap': !exists(json, 'userIdToReactionMap') ? undefined : (mapValues(json['userIdToReactionMap'], ReactionFromJSON)),
        'notes': !exists(json, 'notes') ? undefined : ((json['notes'] as Array<any>).map(NoteFromJSON)),
    };
}

export function CommentToJSON(value?: Comment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'backendPlatformUserId': value.backendPlatformUserId,
        'profileId': value.profileId,
        'media': MediaToJSON(value.media),
        'mentions': value.mentions,
        'hashtags': value.hashtags,
        'createdAt': value.createdAt,
        'content': value.content,
        'replies': value.replies === undefined ? undefined : ((value.replies as Array<any>).map(CommentReplyToJSON)),
        'extra': value.extra,
        'authorUsername': value.authorUsername,
        'authorProfileImage': value.authorProfileImage,
        'affinityScore': value.affinityScore,
        'qualityScore': value.qualityScore,
        'userIdToAffinityScoreMap': value.userIdToAffinityScoreMap,
        'userIdToReportsMap': value.userIdToReportsMap,
        'authorAccountType': AccountTypeToJSON(value.authorAccountType),
        'userIdToReactionMap': value.userIdToReactionMap === undefined ? undefined : (mapValues(value.userIdToReactionMap, ReactionToJSON)),
        'notes': value.notes === undefined ? undefined : ((value.notes as Array<any>).map(NoteToJSON)),
    };
}

