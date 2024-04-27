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
import type { BankAccount } from './BankAccount';
import {
    BankAccountFromJSON,
    BankAccountFromJSONTyped,
    BankAccountToJSON,
} from './BankAccount';
import type { CreditAccount } from './CreditAccount';
import {
    CreditAccountFromJSON,
    CreditAccountFromJSONTyped,
    CreditAccountToJSON,
} from './CreditAccount';
import type { InvestmentAccount } from './InvestmentAccount';
import {
    InvestmentAccountFromJSON,
    InvestmentAccountFromJSONTyped,
    InvestmentAccountToJSON,
} from './InvestmentAccount';
import type { LinkStatus } from './LinkStatus';
import {
    LinkStatusFromJSON,
    LinkStatusFromJSONTyped,
    LinkStatusToJSON,
} from './LinkStatus';
import type { LinkType } from './LinkType';
import {
    LinkTypeFromJSON,
    LinkTypeFromJSONTyped,
    LinkTypeToJSON,
} from './LinkType';
import type { MortgageAccount } from './MortgageAccount';
import {
    MortgageAccountFromJSON,
    MortgageAccountFromJSONTyped,
    MortgageAccountToJSON,
} from './MortgageAccount';
import type { PlaidLink } from './PlaidLink';
import {
    PlaidLinkFromJSON,
    PlaidLinkFromJSONTyped,
    PlaidLinkToJSON,
} from './PlaidLink';
import type { PlaidSync } from './PlaidSync';
import {
    PlaidSyncFromJSON,
    PlaidSyncFromJSONTyped,
    PlaidSyncToJSON,
} from './PlaidSync';
import type { StudentLoanAccount } from './StudentLoanAccount';
import {
    StudentLoanAccountFromJSON,
    StudentLoanAccountFromJSONTyped,
    StudentLoanAccountToJSON,
} from './StudentLoanAccount';
import type { Token } from './Token';
import {
    TokenFromJSON,
    TokenFromJSONTyped,
    TokenToJSON,
} from './Token';

/**
 * A Link represents a login at a financial institution. A single end-user of your application might have accounts at different financial
 * institutions, which means they would have multiple different Items. An Item is not the same as a financial institution account,
 * although every account will be associated with an Item. For example, if a user has one login at their bank that allows them to access
 * both their checking account and their savings account, a single Item would be associated with both of those accounts. Each Item 
 * linked within your application will have a corresponding access_token, which is a token that you can use to make API requests related
 * to that specific Item.
 * Two Items created for the same set of credentials at the same institution will be considered different and not share the same item_id.
 * @export
 * @interface Link
 */
export interface Link {
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    id?: string;
    /**
     * 
     * @type {PlaidSync}
     * @memberof Link
     */
    plaidSync?: PlaidSync;
    /**
     * 
     * @type {LinkStatus}
     * @memberof Link
     */
    linkStatus?: LinkStatus;
    /**
     * 
     * @type {PlaidLink}
     * @memberof Link
     */
    plaidLink?: PlaidLink;
    /**
     * 
     * @type {boolean}
     * @memberof Link
     */
    plaidNewAccountsAvailable?: boolean;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    expirationDate?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    institutionName?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    customInstitutionName?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    lastManualSync?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    lastSuccessfulUpdate?: string;
    /**
     * 
     * @type {Token}
     * @memberof Link
     */
    token?: Token;
    /**
     * 
     * @type {Array<BankAccount>}
     * @memberof Link
     */
    bankAccounts?: Array<BankAccount>;
    /**
     * 
     * @type {Array<InvestmentAccount>}
     * @memberof Link
     */
    investmentAccounts?: Array<InvestmentAccount>;
    /**
     * 
     * @type {Array<CreditAccount>}
     * @memberof Link
     */
    creditAccounts?: Array<CreditAccount>;
    /**
     * 
     * @type {Array<MortgageAccount>}
     * @memberof Link
     */
    mortgageAccounts?: Array<MortgageAccount>;
    /**
     * 
     * @type {Array<StudentLoanAccount>}
     * @memberof Link
     */
    studentLoanAccounts?: Array<StudentLoanAccount>;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    plaidInstitutionId?: string;
    /**
     * 
     * @type {LinkType}
     * @memberof Link
     */
    linkType?: LinkType;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    errorCode?: string;
    /**
     * 
     * @type {string}
     * @memberof Link
     */
    updatedAt?: string;
    /**
     * 
     * @type {boolean}
     * @memberof Link
     */
    newAccountsAvailable?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof Link
     */
    shouldBeUpdated?: boolean;
}

/**
 * Check if a given object implements the Link interface.
 */
export function instanceOfLink(value: object): boolean {
    return true;
}

export function LinkFromJSON(json: any): Link {
    return LinkFromJSONTyped(json, false);
}

export function LinkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Link {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'plaidSync': json['plaidSync'] == null ? undefined : PlaidSyncFromJSON(json['plaidSync']),
        'linkStatus': json['linkStatus'] == null ? undefined : LinkStatusFromJSON(json['linkStatus']),
        'plaidLink': json['plaidLink'] == null ? undefined : PlaidLinkFromJSON(json['plaidLink']),
        'plaidNewAccountsAvailable': json['plaidNewAccountsAvailable'] == null ? undefined : json['plaidNewAccountsAvailable'],
        'expirationDate': json['expirationDate'] == null ? undefined : json['expirationDate'],
        'institutionName': json['institutionName'] == null ? undefined : json['institutionName'],
        'customInstitutionName': json['customInstitutionName'] == null ? undefined : json['customInstitutionName'],
        'description': json['description'] == null ? undefined : json['description'],
        'lastManualSync': json['lastManualSync'] == null ? undefined : json['lastManualSync'],
        'lastSuccessfulUpdate': json['lastSuccessfulUpdate'] == null ? undefined : json['lastSuccessfulUpdate'],
        'token': json['token'] == null ? undefined : TokenFromJSON(json['token']),
        'bankAccounts': json['bankAccounts'] == null ? undefined : ((json['bankAccounts'] as Array<any>).map(BankAccountFromJSON)),
        'investmentAccounts': json['investmentAccounts'] == null ? undefined : ((json['investmentAccounts'] as Array<any>).map(InvestmentAccountFromJSON)),
        'creditAccounts': json['creditAccounts'] == null ? undefined : ((json['creditAccounts'] as Array<any>).map(CreditAccountFromJSON)),
        'mortgageAccounts': json['mortgageAccounts'] == null ? undefined : ((json['mortgageAccounts'] as Array<any>).map(MortgageAccountFromJSON)),
        'studentLoanAccounts': json['studentLoanAccounts'] == null ? undefined : ((json['studentLoanAccounts'] as Array<any>).map(StudentLoanAccountFromJSON)),
        'plaidInstitutionId': json['plaidInstitutionId'] == null ? undefined : json['plaidInstitutionId'],
        'linkType': json['linkType'] == null ? undefined : LinkTypeFromJSON(json['linkType']),
        'errorCode': json['errorCode'] == null ? undefined : json['errorCode'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
        'newAccountsAvailable': json['newAccountsAvailable'] == null ? undefined : json['newAccountsAvailable'],
        'shouldBeUpdated': json['shouldBeUpdated'] == null ? undefined : json['shouldBeUpdated'],
    };
}

export function LinkToJSON(value?: Link | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'plaidSync': PlaidSyncToJSON(value['plaidSync']),
        'linkStatus': LinkStatusToJSON(value['linkStatus']),
        'plaidLink': PlaidLinkToJSON(value['plaidLink']),
        'plaidNewAccountsAvailable': value['plaidNewAccountsAvailable'],
        'expirationDate': value['expirationDate'],
        'institutionName': value['institutionName'],
        'customInstitutionName': value['customInstitutionName'],
        'description': value['description'],
        'lastManualSync': value['lastManualSync'],
        'lastSuccessfulUpdate': value['lastSuccessfulUpdate'],
        'token': TokenToJSON(value['token']),
        'bankAccounts': value['bankAccounts'] == null ? undefined : ((value['bankAccounts'] as Array<any>).map(BankAccountToJSON)),
        'investmentAccounts': value['investmentAccounts'] == null ? undefined : ((value['investmentAccounts'] as Array<any>).map(InvestmentAccountToJSON)),
        'creditAccounts': value['creditAccounts'] == null ? undefined : ((value['creditAccounts'] as Array<any>).map(CreditAccountToJSON)),
        'mortgageAccounts': value['mortgageAccounts'] == null ? undefined : ((value['mortgageAccounts'] as Array<any>).map(MortgageAccountToJSON)),
        'studentLoanAccounts': value['studentLoanAccounts'] == null ? undefined : ((value['studentLoanAccounts'] as Array<any>).map(StudentLoanAccountToJSON)),
        'plaidInstitutionId': value['plaidInstitutionId'],
        'linkType': LinkTypeToJSON(value['linkType']),
        'errorCode': value['errorCode'],
        'updatedAt': value['updatedAt'],
        'newAccountsAvailable': value['newAccountsAvailable'],
        'shouldBeUpdated': value['shouldBeUpdated'],
    };
}

