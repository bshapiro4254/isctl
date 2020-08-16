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

// BootPchStorage Device type used when booting from a PCH Storage device.
type BootPchStorage struct {
	BootDeviceBase `yaml:"BootDeviceBase,inline"`
	Bootloader     *BootBootloader `json:"Bootloader,omitempty" yaml:"Bootloader,omitempty"`
	// The Logical Unit Number (LUN) of the device. It is the Virtual Drive number for Cisco UCS C-Series Servers. Virtual Drive number is found in storage inventory.
	Lun *int64 `json:"Lun,omitempty" yaml:"Lun,omitempty"`
}

// NewBootPchStorage instantiates a new BootPchStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootPchStorage() *BootPchStorage {
	this := BootPchStorage{}
	return &this
}

// NewBootPchStorageWithDefaults instantiates a new BootPchStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootPchStorageWithDefaults() *BootPchStorage {
	this := BootPchStorage{}
	return &this
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise.
func (o *BootPchStorage) GetBootloader() BootBootloader {
	if o == nil || o.Bootloader == nil {
		var ret BootBootloader
		return ret
	}
	return *o.Bootloader
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPchStorage) GetBootloaderOk() (*BootBootloader, bool) {
	if o == nil || o.Bootloader == nil {
		return nil, false
	}
	return o.Bootloader, true
}

// HasBootloader returns a boolean if a field has been set.
func (o *BootPchStorage) HasBootloader() bool {
	if o != nil && o.Bootloader != nil {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given BootBootloader and assigns it to the Bootloader field.
func (o *BootPchStorage) SetBootloader(v BootBootloader) {
	o.Bootloader = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *BootPchStorage) GetLun() int64 {
	if o == nil || o.Lun == nil {
		var ret int64
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPchStorage) GetLunOk() (*int64, bool) {
	if o == nil || o.Lun == nil {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *BootPchStorage) HasLun() bool {
	if o != nil && o.Lun != nil {
		return true
	}

	return false
}

// SetLun gets a reference to the given int64 and assigns it to the Lun field.
func (o *BootPchStorage) SetLun(v int64) {
	o.Lun = &v
}

func (o BootPchStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBootDeviceBase, errBootDeviceBase := json.Marshal(o.BootDeviceBase)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	errBootDeviceBase = json.Unmarshal([]byte(serializedBootDeviceBase), &toSerialize)
	if errBootDeviceBase != nil {
		return []byte{}, errBootDeviceBase
	}
	if o.Bootloader != nil {
		toSerialize["Bootloader"] = o.Bootloader
	}
	if o.Lun != nil {
		toSerialize["Lun"] = o.Lun
	}
	return json.Marshal(toSerialize)
}

type NullableBootPchStorage struct {
	value *BootPchStorage
	isSet bool
}

func (v NullableBootPchStorage) Get() *BootPchStorage {
	return v.value
}

func (v *NullableBootPchStorage) Set(val *BootPchStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPchStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPchStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPchStorage(val *BootPchStorage) *NullableBootPchStorage {
	return &NullableBootPchStorage{value: val, isSet: true}
}

func (v NullableBootPchStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPchStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
