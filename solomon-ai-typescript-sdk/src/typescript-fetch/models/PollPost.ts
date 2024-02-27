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
import type { ContentInsights } from './ContentInsights';
import {
    ContentInsightsFromJSON,
    ContentInsightsFromJSONTyped,
    ContentInsightsToJSON,
} from './ContentInsights';
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
import type { PollResponse } from './PollResponse';
import {
    PollResponseFromJSON,
    PollResponseFromJSONTyped,
    PollResponseToJSON,
} from './PollResponse';
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
import type { ThreadParticipantType } from './ThreadParticipantType';
import {
    ThreadParticipantTypeFromJSON,
    ThreadParticipantTypeFromJSONTyped,
    ThreadParticipantTypeToJSON,
} from './ThreadParticipantType';

/**
 * 
 * @export
 * @interface PollPost
 */
export interface PollPost {
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    createdAt?: string;
    /**
     * 
     * @type {PostType}
     * @memberof PollPost
     */
    action: PostType;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    content: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof PollPost
     */
    mentions?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof PollPost
     */
    hashtags?: Array<string>;
    /**
     * 
     * @type {Media}
     * @memberof PollPost
     */
    media?: Media;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof PollPost
     */
    extra?: { [key: string]: string; };
    /**
     * 
     * @type {Array<Comment>}
     * @memberof PollPost
     */
    comments?: Array<Comment>;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    backendPlatformUserId?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    profileId?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    title?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof PollPost
     */
    tags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    topicName?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    authorUsername?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    authorProfileImage?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    affinityScore?: string;
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    qualityScore?: string;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof PollPost
     */
    userIdToAffinityScoreMap?: { [key: string]: string; };
    /**
     * 
     * @type {ContentInsights}
     * @memberof PollPost
     */
    insights?: ContentInsights;
    /**
     * 
     * @type {{ [key: string]: string; }}
     * @memberof PollPost
     */
    userIdToReportsMap?: { [key: string]: string; };
    /**
     * 
     * @type {string}
     * @memberof PollPost
     */
    backgroundImageUrl?: string;
    /**
     * 
     * @type {AccountType}
     * @memberof PollPost
     */
    authorAccountType?: AccountType;
    /**
     * 
     * @type {{ [key: string]: PollResponse; }}
     * @memberof PollPost
     */
    userIdToPollResponsesMap?: { [key: string]: PollResponse; };
    /**
     * 
     * @type {Array<string>}
     * @memberof PollPost
     */
    pollOptions: Array<string>;
    /**
     * 
     * @type {{ [key: string]: number; }}
     * @memberof PollPost
     */
    pollDistribution?: { [key: string]: number; };
    /**
     * 
     * @type {Date}
     * @memberof PollPost
     */
    pollEndDate?: Date;
    /**
     * 
     * @type {Array<Note>}
     * @memberof PollPost
     */
    notes?: Array<Note>;
    /**
     * 
     * @type {Thread}
     * @memberof PollPost
     */
    thread?: Thread;
    /**
     * 
     * @type {ThreadParticipantType}
     * @memberof PollPost
     */
    threadParticipantType?: ThreadParticipantType;
    /**
     * 
     * @type {{ [key: string]: Reaction; }}
     * @memberof PollPost
     */
    userIdToReactionMap?: { [key: string]: Reaction; };
    /**
     * 
     * @type {Category}
     * @memberof PollPost
     */
    category?: Category;
}

/**
 * Check if a given object implements the PollPost interface.
 */
export function instanceOfPollPost(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "action" in value;
    isInstance = isInstance && "content" in value;
    isInstance = isInstance && "pollOptions" in value;

    return isInstance;
}

export function PollPostFromJSON(json: any): PollPost {
    return PollPostFromJSONTyped(json, false);
}

export function PollPostFromJSONTyped(json: any, ignoreDiscriminator: boolean): PollPost {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'createdAt': !exists(json, 'createdAt') ? undefined : json['createdAt'],
        'action': PostTypeFromJSON(json['action']),
        'content': json['content'],
        'mentions': !exists(json, 'mentions') ? undefined : json['mentions'],
        'hashtags': !exists(json, 'hashtags') ? undefined : json['hashtags'],
        'media': !exists(json, 'media') ? undefined : MediaFromJSON(json['media']),
        'extra': !exists(json, 'extra') ? undefined : json['extra'],
        'comments': !exists(json, 'comments') ? undefined : ((json['comments'] as Array<any>).map(CommentFromJSON)),
        'backendPlatformUserId': !exists(json, 'backendPlatformUserId') ? undefined : json['backendPlatformUserId'],
        'profileId': !exists(json, 'profileId') ? undefined : json['profileId'],
        'title': !exists(json, 'title') ? undefined : json['title'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'topicName': !exists(json, 'topicName') ? undefined : json['topicName'],
        'authorUsername': !exists(json, 'authorUsername') ? undefined : json['authorUsername'],
        'authorProfileImage': !exists(json, 'authorProfileImage') ? undefined : json['authorProfileImage'],
        'affinityScore': !exists(json, 'affinityScore') ? undefined : json['affinityScore'],
        'qualityScore': !exists(json, 'qualityScore') ? undefined : json['qualityScore'],
        'userIdToAffinityScoreMap': !exists(json, 'userIdToAffinityScoreMap') ? undefined : json['userIdToAffinityScoreMap'],
        'insights': !exists(json, 'insights') ? undefined : ContentInsightsFromJSON(json['insights']),
        'userIdToReportsMap': !exists(json, 'userIdToReportsMap') ? undefined : json['userIdToReportsMap'],
        'backgroundImageUrl': !exists(json, 'backgroundImageUrl') ? undefined : json['backgroundImageUrl'],
        'authorAccountType': !exists(json, 'authorAccountType') ? undefined : AccountTypeFromJSON(json['authorAccountType']),
        'userIdToPollResponsesMap': !exists(json, 'userIdToPollResponsesMap') ? undefined : (mapValues(json['userIdToPollResponsesMap'], PollResponseFromJSON)),
        'pollOptions': json['pollOptions'],
        'pollDistribution': !exists(json, 'pollDistribution') ? undefined : json['pollDistribution'],
        'pollEndDate': !exists(json, 'pollEndDate') ? undefined : (new Date(json['pollEndDate'])),
        'notes': !exists(json, 'notes') ? undefined : ((json['notes'] as Array<any>).map(NoteFromJSON)),
        'thread': !exists(json, 'thread') ? undefined : ThreadFromJSON(json['thread']),
        'threadParticipantType': !exists(json, 'threadParticipantType') ? undefined : ThreadParticipantTypeFromJSON(json['threadParticipantType']),
        'userIdToReactionMap': !exists(json, 'userIdToReactionMap') ? undefined : (mapValues(json['userIdToReactionMap'], ReactionFromJSON)),
        'category': !exists(json, 'category') ? undefined : CategoryFromJSON(json['category']),
    };
}

export function PollPostToJSON(value?: PollPost | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'createdAt': value.createdAt,
        'action': PostTypeToJSON(value.action),
        'content': value.content,
        'mentions': value.mentions,
        'hashtags': value.hashtags,
        'media': MediaToJSON(value.media),
        'extra': value.extra,
        'comments': value.comments === undefined ? undefined : ((value.comments as Array<any>).map(CommentToJSON)),
        'backendPlatformUserId': value.backendPlatformUserId,
        'profileId': value.profileId,
        'title': value.title,
        'tags': value.tags,
        'topicName': value.topicName,
        'authorUsername': value.authorUsername,
        'authorProfileImage': value.authorProfileImage,
        'affinityScore': value.affinityScore,
        'qualityScore': value.qualityScore,
        'userIdToAffinityScoreMap': value.userIdToAffinityScoreMap,
        'insights': ContentInsightsToJSON(value.insights),
        'userIdToReportsMap': value.userIdToReportsMap,
        'backgroundImageUrl': value.backgroundImageUrl,
        'authorAccountType': AccountTypeToJSON(value.authorAccountType),
        'userIdToPollResponsesMap': value.userIdToPollResponsesMap === undefined ? undefined : (mapValues(value.userIdToPollResponsesMap, PollResponseToJSON)),
        'pollOptions': value.pollOptions,
        'pollDistribution': value.pollDistribution,
        'pollEndDate': value.pollEndDate === undefined ? undefined : (value.pollEndDate.toISOString()),
        'notes': value.notes === undefined ? undefined : ((value.notes as Array<any>).map(NoteToJSON)),
        'thread': ThreadToJSON(value.thread),
        'threadParticipantType': ThreadParticipantTypeToJSON(value.threadParticipantType),
        'userIdToReactionMap': value.userIdToReactionMap === undefined ? undefined : (mapValues(value.userIdToReactionMap, ReactionToJSON)),
        'category': CategoryToJSON(value.category),
    };
}

