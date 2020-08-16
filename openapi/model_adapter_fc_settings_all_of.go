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

// AdapterFcSettingsAllOf Definition of the list of properties defined in 'adapter.FcSettings', excluding properties defined in parent classes.
type AdapterFcSettingsAllOf struct {
	// Status of FIP protocol on the adapter interfaces.
	FipEnabled *bool `json:"FipEnabled,omitempty" yaml:"FipEnabled,omitempty"`
}

// NewAdapterFcSettingsAllOf instantiates a new AdapterFcSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterFcSettingsAllOf() *AdapterFcSettingsAllOf {
	this := AdapterFcSettingsAllOf{}
	return &this
}

// NewAdapterFcSettingsAllOfWithDefaults instantiates a new AdapterFcSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterFcSettingsAllOfWithDefaults() *AdapterFcSettingsAllOf {
	this := AdapterFcSettingsAllOf{}
	return &this
}

// GetFipEnabled returns the FipEnabled field value if set, zero value otherwise.
func (o *AdapterFcSettingsAllOf) GetFipEnabled() bool {
	if o == nil || o.FipEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FipEnabled
}

// GetFipEnabledOk returns a tuple with the FipEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterFcSettingsAllOf) GetFipEnabledOk() (*bool, bool) {
	if o == nil || o.FipEnabled == nil {
		return nil, false
	}
	return o.FipEnabled, true
}

// HasFipEnabled returns a boolean if a field has been set.
func (o *AdapterFcSettingsAllOf) HasFipEnabled() bool {
	if o != nil && o.FipEnabled != nil {
		return true
	}

	return false
}

// SetFipEnabled gets a reference to the given bool and assigns it to the FipEnabled field.
func (o *AdapterFcSettingsAllOf) SetFipEnabled(v bool) {
	o.FipEnabled = &v
}

func (o AdapterFcSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FipEnabled != nil {
		toSerialize["FipEnabled"] = o.FipEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableAdapterFcSettingsAllOf struct {
	value *AdapterFcSettingsAllOf
	isSet bool
}

func (v NullableAdapterFcSettingsAllOf) Get() *AdapterFcSettingsAllOf {
	return v.value
}

func (v *NullableAdapterFcSettingsAllOf) Set(val *AdapterFcSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterFcSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterFcSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterFcSettingsAllOf(val *AdapterFcSettingsAllOf) *NullableAdapterFcSettingsAllOf {
	return &NullableAdapterFcSettingsAllOf{value: val, isSet: true}
}

func (v NullableAdapterFcSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterFcSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
