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

// VirtualizationVmwareResourceConsumptionAllOf Definition of the list of properties defined in 'virtualization.VmwareResourceConsumption', excluding properties defined in parent classes.
type VirtualizationVmwareResourceConsumptionAllOf struct {
	// The amount of CPU consumed in Hz.
	CpuConsumed *int64 `json:"CpuConsumed,omitempty" yaml:"CpuConsumed,omitempty"`
	// Memory consumed by this host in bytes.
	MemoryConsumed *int64 `json:"MemoryConsumed,omitempty" yaml:"MemoryConsumed,omitempty"`
}

// NewVirtualizationVmwareResourceConsumptionAllOf instantiates a new VirtualizationVmwareResourceConsumptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareResourceConsumptionAllOf() *VirtualizationVmwareResourceConsumptionAllOf {
	this := VirtualizationVmwareResourceConsumptionAllOf{}
	return &this
}

// NewVirtualizationVmwareResourceConsumptionAllOfWithDefaults instantiates a new VirtualizationVmwareResourceConsumptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareResourceConsumptionAllOfWithDefaults() *VirtualizationVmwareResourceConsumptionAllOf {
	this := VirtualizationVmwareResourceConsumptionAllOf{}
	return &this
}

// GetCpuConsumed returns the CpuConsumed field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceConsumptionAllOf) GetCpuConsumed() int64 {
	if o == nil || o.CpuConsumed == nil {
		var ret int64
		return ret
	}
	return *o.CpuConsumed
}

// GetCpuConsumedOk returns a tuple with the CpuConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceConsumptionAllOf) GetCpuConsumedOk() (*int64, bool) {
	if o == nil || o.CpuConsumed == nil {
		return nil, false
	}
	return o.CpuConsumed, true
}

// HasCpuConsumed returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceConsumptionAllOf) HasCpuConsumed() bool {
	if o != nil && o.CpuConsumed != nil {
		return true
	}

	return false
}

// SetCpuConsumed gets a reference to the given int64 and assigns it to the CpuConsumed field.
func (o *VirtualizationVmwareResourceConsumptionAllOf) SetCpuConsumed(v int64) {
	o.CpuConsumed = &v
}

// GetMemoryConsumed returns the MemoryConsumed field value if set, zero value otherwise.
func (o *VirtualizationVmwareResourceConsumptionAllOf) GetMemoryConsumed() int64 {
	if o == nil || o.MemoryConsumed == nil {
		var ret int64
		return ret
	}
	return *o.MemoryConsumed
}

// GetMemoryConsumedOk returns a tuple with the MemoryConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareResourceConsumptionAllOf) GetMemoryConsumedOk() (*int64, bool) {
	if o == nil || o.MemoryConsumed == nil {
		return nil, false
	}
	return o.MemoryConsumed, true
}

// HasMemoryConsumed returns a boolean if a field has been set.
func (o *VirtualizationVmwareResourceConsumptionAllOf) HasMemoryConsumed() bool {
	if o != nil && o.MemoryConsumed != nil {
		return true
	}

	return false
}

// SetMemoryConsumed gets a reference to the given int64 and assigns it to the MemoryConsumed field.
func (o *VirtualizationVmwareResourceConsumptionAllOf) SetMemoryConsumed(v int64) {
	o.MemoryConsumed = &v
}

func (o VirtualizationVmwareResourceConsumptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuConsumed != nil {
		toSerialize["CpuConsumed"] = o.CpuConsumed
	}
	if o.MemoryConsumed != nil {
		toSerialize["MemoryConsumed"] = o.MemoryConsumed
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualizationVmwareResourceConsumptionAllOf struct {
	value *VirtualizationVmwareResourceConsumptionAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareResourceConsumptionAllOf) Get() *VirtualizationVmwareResourceConsumptionAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareResourceConsumptionAllOf) Set(val *VirtualizationVmwareResourceConsumptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareResourceConsumptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareResourceConsumptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareResourceConsumptionAllOf(val *VirtualizationVmwareResourceConsumptionAllOf) *NullableVirtualizationVmwareResourceConsumptionAllOf {
	return &NullableVirtualizationVmwareResourceConsumptionAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareResourceConsumptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareResourceConsumptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
