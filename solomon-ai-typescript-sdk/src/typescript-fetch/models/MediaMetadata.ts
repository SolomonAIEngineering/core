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
import type { MediaCrop } from './MediaCrop';
import {
    MediaCropFromJSON,
    MediaCropFromJSONTyped,
    MediaCropToJSON,
} from './MediaCrop';
import type { MediaResize } from './MediaResize';
import {
    MediaResizeFromJSON,
    MediaResizeFromJSONTyped,
    MediaResizeToJSON,
} from './MediaResize';
import type { MediaType } from './MediaType';
import {
    MediaTypeFromJSON,
    MediaTypeFromJSONTyped,
    MediaTypeToJSON,
} from './MediaType';

/**
 * 
 * @export
 * @interface MediaMetadata
 */
export interface MediaMetadata {
    /**
     * 
     * @type {string}
     * @memberof MediaMetadata
     */
    id?: string;
    /**
     * 
     * @type {MediaResize}
     * @memberof MediaMetadata
     */
    resize?: MediaResize;
    /**
     * 
     * @type {MediaCrop}
     * @memberof MediaMetadata
     */
    crop?: MediaCrop;
    /**
     * 
     * @type {string}
     * @memberof MediaMetadata
     */
    imageWidth?: string;
    /**
     * 
     * @type {string}
     * @memberof MediaMetadata
     */
    imageHeight?: string;
    /**
     * 
     * @type {MediaType}
     * @memberof MediaMetadata
     */
    type?: MediaType;
}

/**
 * Check if a given object implements the MediaMetadata interface.
 */
export function instanceOfMediaMetadata(value: object): boolean {
    return true;
}

export function MediaMetadataFromJSON(json: any): MediaMetadata {
    return MediaMetadataFromJSONTyped(json, false);
}

export function MediaMetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): MediaMetadata {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'resize': json['resize'] == null ? undefined : MediaResizeFromJSON(json['resize']),
        'crop': json['crop'] == null ? undefined : MediaCropFromJSON(json['crop']),
        'imageWidth': json['imageWidth'] == null ? undefined : json['imageWidth'],
        'imageHeight': json['imageHeight'] == null ? undefined : json['imageHeight'],
        'type': json['type'] == null ? undefined : MediaTypeFromJSON(json['type']),
    };
}

export function MediaMetadataToJSON(value?: MediaMetadata | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'resize': MediaResizeToJSON(value['resize']),
        'crop': MediaCropToJSON(value['crop']),
        'imageWidth': value['imageWidth'],
        'imageHeight': value['imageHeight'],
        'type': MediaTypeToJSON(value['type']),
    };
}

