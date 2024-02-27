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
import type { Category } from './Category';
import {
    CategoryFromJSON,
    CategoryFromJSONTyped,
    CategoryToJSON,
} from './Category';
import type { Comment } from './Comment';
import {
    CommentFromJSON,
    CommentFromJSONTyped,
    CommentToJSON,
} from './Comment';
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
import type { Reaction } from './Reaction';
import {
    ReactionFromJSON,
    ReactionFromJSONTyped,
    ReactionToJSON,
} from './Reaction';
import type { Thread } from './Thread';
import {
    ThreadFromJSON,
    ThreadFromJSONTyped,
    ThreadToJSON,
} from './Thread';

/**
 * 
 * @export
 * @interface SharedPost
 */
export interface SharedPost {
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    originalPostId: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    originalAuthorUsername?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    createdAt?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    content: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SharedPost
     */
    mentions?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof SharedPost
     */
    hashtags?: Array<string>;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof SharedPost
     */
    extra?: { [key: string]: string; };
    /**
     * 
     * @type {Array<Comment>}
     * @memberof SharedPost
     */
    comments?: Array<Comment>;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    backendPlatformUserId?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    profileId?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof SharedPost
     */
    tags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    authorUsername?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    affinityScore?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    qualityScore?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof SharedPost
     */
    userIdToAffinityScoreMap?: { [key: string]: string; };
    /**
     * AccountType is the account type of the creator of this piece of
     *  content
     * @type {{ [key: string]: string; }}
     * @memberof SharedPost
     */
    userIdToReportsMap?: { [key: string]: string; };
    /**
     * 
     * @type {Array<Note>}
     * @memberof SharedPost
     */
    notes?: Array<Note>;
    /**
     * 
     * @type {Thread}
     * @memberof SharedPost
     */
    thread?: Thread;
    /**
     * 
     * @type {AccountType}
     * @memberof SharedPost
     */
    authorAccountType?: AccountType;
    /**
     * 
     * @type {{ [key: string]: Reaction; }}
     * @memberof SharedPost
     */
    userIdToReactionMap?: { [key: string]: Reaction; };
    /**
     * 
     * @type {PostType}
     * @memberof SharedPost
     */
    action: PostType;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    originalPostUserProfileId?: string;
    /**
     * 
     * @type {string}
     * @memberof SharedPost
     */
    originalPostUserbackendPlaformId?: string;
    /**
     * 
     * @type {PostType}
     * @memberof SharedPost
     */
    originalPostAction: PostType;
    /**
     * 
     * @type {Category}
     * @memberof SharedPost
     */
    category?: Category;
}

/**
 * Check if a given object implements the SharedPost interface.
 */
export function instanceOfSharedPost(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "originalPostId" in value;
    isInstance = isInstance && "content" in value;
    isInstance = isInstance && "action" in value;
    isInstance = isInstance && "originalPostAction" in value;

    return isInstance;
}

export function SharedPostFromJSON(json: any): SharedPost {
    return SharedPostFromJSONTyped(json, false);
}

export function SharedPostFromJSONTyped(json: any, ignoreDiscriminator: boolean): SharedPost {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'originalPostId': json['originalPostId'],
        'originalAuthorUsername': !exists(json, 'originalAuthorUsername') ? undefined : json['originalAuthorUsername'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'content': json['content'],
        'mentions': !exists(json, 'mentions') ? undefined : json['mentions'],
        'hashtags': !exists(json, 'hashtags') ? undefined : json['hashtags'],
        'extra': !exists(json, 'extra') ? undefined : json['extra'],
        'comments': !exists(json, 'comments') ? undefined : ((json['comments'] as Array<any>).map(CommentFromJSON)),
        'backendPlatformUserId': !exists(json, 'backendPlatformUserId') ? undefined : json['backendPlatformUserId'],
        'profileId': !exists(json, 'profileId') ? undefined : json['profileId'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'authorUsername': !exists(json, 'authorUsername') ? undefined : json['authorUsername'],
        'affinityScore': !exists(json, 'affinityScore') ? undefined : json['affinityScore'],
        'qualityScore': !exists(json, 'qualityScore') ? undefined : json['qualityScore'],
        'userIdToAffinityScoreMap': !exists(json, 'userIdToAffinityScoreMap') ? undefined : json['userIdToAffinityScoreMap'],
        'userIdToReportsMap': !exists(json, 'userIdToReportsMap') ? undefined : json['userIdToReportsMap'],
        'notes': !exists(json, 'notes') ? undefined : ((json['notes'] as Array<any>).map(NoteFromJSON)),
        'thread': !exists(json, 'thread') ? undefined : ThreadFromJSON(json['thread']),
        'authorAccountType': !exists(json, 'authorAccountType') ? undefined : AccountTypeFromJSON(json['authorAccountType']),
        'userIdToReactionMap': !exists(json, 'userIdToReactionMap') ? undefined : (mapValues(json['userIdToReactionMap'], ReactionFromJSON)),
        'action': PostTypeFromJSON(json['action']),
        'originalPostUserProfileId': !exists(json, 'originalPostUserProfileId') ? undefined : json['originalPostUserProfileId'],
        'originalPostUserbackendPlaformId': !exists(json, 'originalPostUserbackendPlaformId') ? undefined : json['originalPostUserbackendPlaformId'],
        'originalPostAction': PostTypeFromJSON(json['originalPostAction']),
        'category': !exists(json, 'category') ? undefined : CategoryFromJSON(json['category']),
    };
}

export function SharedPostToJSON(value?: SharedPost | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'originalPostId': value.originalPostId,
        'originalAuthorUsername': value.originalAuthorUsername,
        'createdAt': value.createdAt,
        'content': value.content,
        'mentions': value.mentions,
        'hashtags': value.hashtags,
        'extra': value.extra,
        'comments': value.comments === undefined ? undefined : ((value.comments as Array<any>).map(CommentToJSON)),
        'backendPlatformUserId': value.backendPlatformUserId,
        'profileId': value.profileId,
        'tags': value.tags,
        'authorUsername': value.authorUsername,
        'affinityScore': value.affinityScore,
        'qualityScore': value.qualityScore,
        'userIdToAffinityScoreMap': value.userIdToAffinityScoreMap,
        'userIdToReportsMap': value.userIdToReportsMap,
        'notes': value.notes === undefined ? undefined : ((value.notes as Array<any>).map(NoteToJSON)),
        'thread': ThreadToJSON(value.thread),
        'authorAccountType': AccountTypeToJSON(value.authorAccountType),
        'userIdToReactionMap': value.userIdToReactionMap === undefined ? undefined : (mapValues(value.userIdToReactionMap, ReactionToJSON)),
        'action': PostTypeToJSON(value.action),
        'originalPostUserProfileId': value.originalPostUserProfileId,
        'originalPostUserbackendPlaformId': value.originalPostUserbackendPlaformId,
        'originalPostAction': PostTypeToJSON(value.originalPostAction),
        'category': CategoryToJSON(value.category),
    };
}

