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

// CapabilityIoCardCapabilityDefAllOf Definition of the list of properties defined in 'capability.IoCardCapabilityDef', excluding properties defined in parent classes.
type CapabilityIoCardCapabilityDefAllOf struct {
	// Device connector support on Iocard.
	DcSupported *bool `json:"DcSupported,omitempty" yaml:"DcSupported,omitempty"`
}

// NewCapabilityIoCardCapabilityDefAllOf instantiates a new CapabilityIoCardCapabilityDefAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityIoCardCapabilityDefAllOf() *CapabilityIoCardCapabilityDefAllOf {
	this := CapabilityIoCardCapabilityDefAllOf{}
	return &this
}

// NewCapabilityIoCardCapabilityDefAllOfWithDefaults instantiates a new CapabilityIoCardCapabilityDefAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityIoCardCapabilityDefAllOfWithDefaults() *CapabilityIoCardCapabilityDefAllOf {
	this := CapabilityIoCardCapabilityDefAllOf{}
	return &this
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *CapabilityIoCardCapabilityDefAllOf) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityIoCardCapabilityDefAllOf) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *CapabilityIoCardCapabilityDefAllOf) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *CapabilityIoCardCapabilityDefAllOf) SetDcSupported(v bool) {
	o.DcSupported = &v
}

func (o CapabilityIoCardCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}
	return json.Marshal(toSerialize)
}

type NullableCapabilityIoCardCapabilityDefAllOf struct {
	value *CapabilityIoCardCapabilityDefAllOf
	isSet bool
}

func (v NullableCapabilityIoCardCapabilityDefAllOf) Get() *CapabilityIoCardCapabilityDefAllOf {
	return v.value
}

func (v *NullableCapabilityIoCardCapabilityDefAllOf) Set(val *CapabilityIoCardCapabilityDefAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityIoCardCapabilityDefAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityIoCardCapabilityDefAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityIoCardCapabilityDefAllOf(val *CapabilityIoCardCapabilityDefAllOf) *NullableCapabilityIoCardCapabilityDefAllOf {
	return &NullableCapabilityIoCardCapabilityDefAllOf{value: val, isSet: true}
}

func (v NullableCapabilityIoCardCapabilityDefAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityIoCardCapabilityDefAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
