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
import type { Address } from './Address';
import {
    AddressFromJSON,
    AddressFromJSONTyped,
    AddressToJSON,
} from './Address';
import type { ProfileType } from './ProfileType';
import {
    ProfileTypeFromJSON,
    ProfileTypeFromJSONTyped,
    ProfileTypeToJSON,
} from './ProfileType';
import type { Role } from './Role';
import {
    RoleFromJSON,
    RoleFromJSONTyped,
    RoleToJSON,
} from './Role';
import type { Settings } from './Settings';
import {
    SettingsFromJSON,
    SettingsFromJSONTyped,
    SettingsToJSON,
} from './Settings';
import type { Tags } from './Tags';
import {
    TagsFromJSON,
    TagsFromJSONTyped,
    TagsToJSON,
} from './Tags';

/**
 * BusinessAccount represents a business account within the context of solomon-ai.
 * @export
 * @interface BusinessAccount
 */
export interface BusinessAccount {
    /**
     * Unique identifier for the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    id?: string;
    /**
     * Email associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    email?: string;
    /**
     * 
     * @type {Address}
     * @memberof BusinessAccount
     */
    address?: Address;
    /**
     * Short description of the business account. Maximum of 200 characters.
     * @type {string}
     * @memberof BusinessAccount
     */
    bio?: string;
    /**
     * Headline for the profile of the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    headline?: string;
    /**
     * Phone number associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    phoneNumber?: string;
    /**
     * Tags associated with the business account. Between 1 and 10 tags are allowed.
     * @type {Array<Tags>}
     * @memberof BusinessAccount
     */
    tags?: Array<Tags>;
    /**
     * Identifier for the associated authentication service account.
     * @type {string}
     * @memberof BusinessAccount
     */
    authnAccountId?: string;
    /**
     * Indicates whether the business account is active.
     * @type {boolean}
     * @memberof BusinessAccount
     */
    isActive?: boolean;
    /**
     * Username for the business account. Must be at least 10 characters long.
     * @type {string}
     * @memberof BusinessAccount
     */
    username?: string;
    /**
     * Indicates whether the business account is private.
     * @type {boolean}
     * @memberof BusinessAccount
     */
    isPrivate?: boolean;
    /**
     * Indicates whether the email associated with the business account has been verified.
     * @type {boolean}
     * @memberof BusinessAccount
     */
    isEmailVerified?: boolean;
    /**
     * Timestamp indicating when the business account was created.
     * @type {Date}
     * @memberof BusinessAccount
     */
    createdAt?: Date;
    /**
     * Timestamp indicating when the email for the business account was verified.
     * @type {Date}
     * @memberof BusinessAccount
     */
    verifiedAt?: Date;
    /**
     * Date when the company associated with the business account was established.
     * @type {string}
     * @memberof BusinessAccount
     */
    companyEstablishedDate?: string;
    /**
     * Industry type of the company associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    companyIndustryType?: string;
    /**
     * Website URL of the company associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    companyWebsiteUrl?: string;
    /**
     * Description of the company associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    companyDescription?: string;
    /**
     * Name of the company associated with the business account.
     * @type {string}
     * @memberof BusinessAccount
     */
    companyName?: string;
    /**
     * 
     * @type {Settings}
     * @memberof BusinessAccount
     */
    settings?: Settings;
    /**
     * 
     * @type {ProfileType}
     * @memberof BusinessAccount
     */
    accountType?: ProfileType;
    /**
     * Profile image associated with the user account.
     * @type {string}
     * @memberof BusinessAccount
     */
    profileImageUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof BusinessAccount
     */
    auth0UserId?: string;
    /**
     * 
     * @type {Role}
     * @memberof BusinessAccount
     */
    role?: Role;
    /**
     * 
     * @type {string}
     * @memberof BusinessAccount
     */
    algoliaUserId?: string;
}

/**
 * Check if a given object implements the BusinessAccount interface.
 */
export function instanceOfBusinessAccount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BusinessAccountFromJSON(json: any): BusinessAccount {
    return BusinessAccountFromJSONTyped(json, false);
}

export function BusinessAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): BusinessAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'address': !exists(json, 'address') ? undefined : AddressFromJSON(json['address']),
        'bio': !exists(json, 'bio') ? undefined : json['bio'],
        'headline': !exists(json, 'headline') ? undefined : json['headline'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'tags': !exists(json, 'tags') ? undefined : ((json['tags'] as Array<any>).map(TagsFromJSON)),
        'authnAccountId': !exists(json, 'authnAccountId') ? undefined : json['authnAccountId'],
        'isActive': !exists(json, 'isActive') ? undefined : json['isActive'],
        'username': !exists(json, 'username') ? undefined : json['username'],
        'isPrivate': !exists(json, 'isPrivate') ? undefined : json['isPrivate'],
        'isEmailVerified': !exists(json, 'isEmailVerified') ? undefined : json['isEmailVerified'],
        'createdAt': !exists(json, 'createdAt') ? undefined : (new Date(json['createdAt'])),
        'verifiedAt': !exists(json, 'verifiedAt') ? undefined : (new Date(json['verifiedAt'])),
        'companyEstablishedDate': !exists(json, 'companyEstablishedDate') ? undefined : json['companyEstablishedDate'],
        'companyIndustryType': !exists(json, 'companyIndustryType') ? undefined : json['companyIndustryType'],
        'companyWebsiteUrl': !exists(json, 'companyWebsiteUrl') ? undefined : json['companyWebsiteUrl'],
        'companyDescription': !exists(json, 'companyDescription') ? undefined : json['companyDescription'],
        'companyName': !exists(json, 'companyName') ? undefined : json['companyName'],
        'settings': !exists(json, 'settings') ? undefined : SettingsFromJSON(json['settings']),
        'accountType': !exists(json, 'accountType') ? undefined : ProfileTypeFromJSON(json['accountType']),
        'profileImageUrl': !exists(json, 'profileImageUrl') ? undefined : json['profileImageUrl'],
        'auth0UserId': !exists(json, 'auth0UserId') ? undefined : json['auth0UserId'],
        'role': !exists(json, 'role') ? undefined : RoleFromJSON(json['role']),
        'algoliaUserId': !exists(json, 'algoliaUserId') ? undefined : json['algoliaUserId'],
    };
}

export function BusinessAccountToJSON(value?: BusinessAccount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'email': value.email,
        'address': AddressToJSON(value.address),
        'bio': value.bio,
        'headline': value.headline,
        'phoneNumber': value.phoneNumber,
        'tags': value.tags === undefined ? undefined : ((value.tags as Array<any>).map(TagsToJSON)),
        'authnAccountId': value.authnAccountId,
        'isActive': value.isActive,
        'username': value.username,
        'isPrivate': value.isPrivate,
        'isEmailVerified': value.isEmailVerified,
        'createdAt': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'verifiedAt': value.verifiedAt === undefined ? undefined : (value.verifiedAt.toISOString()),
        'companyEstablishedDate': value.companyEstablishedDate,
        'companyIndustryType': value.companyIndustryType,
        'companyWebsiteUrl': value.companyWebsiteUrl,
        'companyDescription': value.companyDescription,
        'companyName': value.companyName,
        'settings': SettingsToJSON(value.settings),
        'accountType': ProfileTypeToJSON(value.accountType),
        'profileImageUrl': value.profileImageUrl,
        'auth0UserId': value.auth0UserId,
        'role': RoleToJSON(value.role),
        'algoliaUserId': value.algoliaUserId,
    };
}

