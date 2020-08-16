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

// VnicScsiQueueSettingsAllOf Definition of the list of properties defined in 'vnic.ScsiQueueSettings', excluding properties defined in parent classes.
type VnicScsiQueueSettingsAllOf struct {
	// The number of SCSI I/O queue resources the system should allocate.
	Count *int64 `json:"Count,omitempty" yaml:"Count,omitempty"`
	// The number of descriptors in each SCSI I/O queue.
	RingSize *int64 `json:"RingSize,omitempty" yaml:"RingSize,omitempty"`
}

// NewVnicScsiQueueSettingsAllOf instantiates a new VnicScsiQueueSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicScsiQueueSettingsAllOf() *VnicScsiQueueSettingsAllOf {
	this := VnicScsiQueueSettingsAllOf{}
	return &this
}

// NewVnicScsiQueueSettingsAllOfWithDefaults instantiates a new VnicScsiQueueSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicScsiQueueSettingsAllOfWithDefaults() *VnicScsiQueueSettingsAllOf {
	this := VnicScsiQueueSettingsAllOf{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VnicScsiQueueSettingsAllOf) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScsiQueueSettingsAllOf) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VnicScsiQueueSettingsAllOf) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *VnicScsiQueueSettingsAllOf) SetCount(v int64) {
	o.Count = &v
}

// GetRingSize returns the RingSize field value if set, zero value otherwise.
func (o *VnicScsiQueueSettingsAllOf) GetRingSize() int64 {
	if o == nil || o.RingSize == nil {
		var ret int64
		return ret
	}
	return *o.RingSize
}

// GetRingSizeOk returns a tuple with the RingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicScsiQueueSettingsAllOf) GetRingSizeOk() (*int64, bool) {
	if o == nil || o.RingSize == nil {
		return nil, false
	}
	return o.RingSize, true
}

// HasRingSize returns a boolean if a field has been set.
func (o *VnicScsiQueueSettingsAllOf) HasRingSize() bool {
	if o != nil && o.RingSize != nil {
		return true
	}

	return false
}

// SetRingSize gets a reference to the given int64 and assigns it to the RingSize field.
func (o *VnicScsiQueueSettingsAllOf) SetRingSize(v int64) {
	o.RingSize = &v
}

func (o VnicScsiQueueSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.RingSize != nil {
		toSerialize["RingSize"] = o.RingSize
	}
	return json.Marshal(toSerialize)
}

type NullableVnicScsiQueueSettingsAllOf struct {
	value *VnicScsiQueueSettingsAllOf
	isSet bool
}

func (v NullableVnicScsiQueueSettingsAllOf) Get() *VnicScsiQueueSettingsAllOf {
	return v.value
}

func (v *NullableVnicScsiQueueSettingsAllOf) Set(val *VnicScsiQueueSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicScsiQueueSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicScsiQueueSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicScsiQueueSettingsAllOf(val *VnicScsiQueueSettingsAllOf) *NullableVnicScsiQueueSettingsAllOf {
	return &NullableVnicScsiQueueSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicScsiQueueSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicScsiQueueSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
