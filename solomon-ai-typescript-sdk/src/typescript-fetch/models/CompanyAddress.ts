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
/**
 * The CompanyAddress object is used to represent a contact's or company's address.
 * @export
 * @interface CompanyAddress
 */
export interface CompanyAddress {
    /**
     * 
     * @type {string}
     * @memberof CompanyAddress
     */
    id?: string;
    /**
     * 
     * @type {string}
     * @memberof CompanyAddress
     */
    type?: string;
    /**
     * Line 1 of the address's street.
     * @type {string}
     * @memberof CompanyAddress
     */
    street1?: string;
    /**
     * Line 2 of the address's street.
     * @type {string}
     * @memberof CompanyAddress
     */
    street2?: string;
    /**
     * The address's city.
     * @type {string}
     * @memberof CompanyAddress
     */
    city?: string;
    /**
     * The address's state or region.
     * @type {string}
     * @memberof CompanyAddress
     */
    state?: string;
    /**
     * Typically, this might just be 'state' but used your field name to keep it consistent with the JSON
     * @type {string}
     * @memberof CompanyAddress
     */
    countrySubdivision?: string;
    /**
     * The address's country.
     * @type {string}
     * @memberof CompanyAddress
     */
    country?: string;
    /**
     * The address's zip code.
     * @type {string}
     * @memberof CompanyAddress
     */
    zipCode?: string;
    /**
     * Consider using google.protobuf.Timestamp if precise time manipulation is required
     * @type {Date}
     * @memberof CompanyAddress
     */
    modifiedAt?: Date;
}

/**
 * Check if a given object implements the CompanyAddress interface.
 */
export function instanceOfCompanyAddress(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CompanyAddressFromJSON(json: any): CompanyAddress {
    return CompanyAddressFromJSONTyped(json, false);
}

export function CompanyAddressFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompanyAddress {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'street1': !exists(json, 'street1') ? undefined : json['street1'],
        'street2': !exists(json, 'street2') ? undefined : json['street2'],
        'city': !exists(json, 'city') ? undefined : json['city'],
        'state': !exists(json, 'state') ? undefined : json['state'],
        'countrySubdivision': !exists(json, 'countrySubdivision') ? undefined : json['countrySubdivision'],
        'country': !exists(json, 'country') ? undefined : json['country'],
        'zipCode': !exists(json, 'zipCode') ? undefined : json['zipCode'],
        'modifiedAt': !exists(json, 'modifiedAt') ? undefined : (new Date(json['modifiedAt'])),
    };
}

export function CompanyAddressToJSON(value?: CompanyAddress | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'type': value.type,
        'street1': value.street1,
        'street2': value.street2,
        'city': value.city,
        'state': value.state,
        'countrySubdivision': value.countrySubdivision,
        'country': value.country,
        'zipCode': value.zipCode,
        'modifiedAt': value.modifiedAt === undefined ? undefined : (value.modifiedAt.toISOString()),
    };
}

