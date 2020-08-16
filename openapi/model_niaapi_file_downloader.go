/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// NiaapiFileDownloader Provide a presigned url to download the metadata file from server.
type NiaapiFileDownloader struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// Filename of this Metadata package file, folder will be handled by api.
	FileName *string `json:"FileName,omitempty" yaml:"FileName,omitempty"`
	// The presigned URL from server to download this file.
	PresignedUrl *string `json:"PresignedUrl,omitempty" yaml:"PresignedUrl,omitempty"`
}

// NewNiaapiFileDownloader instantiates a new NiaapiFileDownloader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiFileDownloader() *NiaapiFileDownloader {
	this := NiaapiFileDownloader{}
	return &this
}

// NewNiaapiFileDownloaderWithDefaults instantiates a new NiaapiFileDownloader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiFileDownloaderWithDefaults() *NiaapiFileDownloader {
	this := NiaapiFileDownloader{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NiaapiFileDownloader) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloader) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NiaapiFileDownloader) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NiaapiFileDownloader) SetFileName(v string) {
	o.FileName = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *NiaapiFileDownloader) GetPresignedUrl() string {
	if o == nil || o.PresignedUrl == nil {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFileDownloader) GetPresignedUrlOk() (*string, bool) {
	if o == nil || o.PresignedUrl == nil {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *NiaapiFileDownloader) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl != nil {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *NiaapiFileDownloader) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

func (o NiaapiFileDownloader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.PresignedUrl != nil {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}
	return json.Marshal(toSerialize)
}

type NullableNiaapiFileDownloader struct {
	value *NiaapiFileDownloader
	isSet bool
}

func (v NullableNiaapiFileDownloader) Get() *NiaapiFileDownloader {
	return v.value
}

func (v *NullableNiaapiFileDownloader) Set(val *NiaapiFileDownloader) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiFileDownloader) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiFileDownloader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiFileDownloader(val *NiaapiFileDownloader) *NullableNiaapiFileDownloader {
	return &NullableNiaapiFileDownloader{value: val, isSet: true}
}

func (v NullableNiaapiFileDownloader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiFileDownloader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
