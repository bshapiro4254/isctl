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

// PkixEddsaKeySpecAllOf Definition of the list of properties defined in 'pkix.EddsaKeySpec', excluding properties defined in parent classes.
type PkixEddsaKeySpecAllOf struct {
	// The EdDSA algorithm, as defined in RFC 8032. * `Ed25519` - The edwards25519 curve, as defined in RFC 8032 section 5.1. * `Ed25519ph` - The edwards25519 curve for the PureEdDSA variant. * `Ed25519ctx` - The edwards25519 curve for the HashEdDSA variant.
	Algorithm *string `json:"Algorithm,omitempty" yaml:"Algorithm,omitempty"`
}

// NewPkixEddsaKeySpecAllOf instantiates a new PkixEddsaKeySpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkixEddsaKeySpecAllOf() *PkixEddsaKeySpecAllOf {
	this := PkixEddsaKeySpecAllOf{}
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// NewPkixEddsaKeySpecAllOfWithDefaults instantiates a new PkixEddsaKeySpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkixEddsaKeySpecAllOfWithDefaults() *PkixEddsaKeySpecAllOf {
	this := PkixEddsaKeySpecAllOf{}
	var algorithm string = "Ed25519"
	this.Algorithm = &algorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PkixEddsaKeySpecAllOf) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkixEddsaKeySpecAllOf) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PkixEddsaKeySpecAllOf) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *PkixEddsaKeySpecAllOf) SetAlgorithm(v string) {
	o.Algorithm = &v
}

func (o PkixEddsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["Algorithm"] = o.Algorithm
	}
	return json.Marshal(toSerialize)
}

type NullablePkixEddsaKeySpecAllOf struct {
	value *PkixEddsaKeySpecAllOf
	isSet bool
}

func (v NullablePkixEddsaKeySpecAllOf) Get() *PkixEddsaKeySpecAllOf {
	return v.value
}

func (v *NullablePkixEddsaKeySpecAllOf) Set(val *PkixEddsaKeySpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePkixEddsaKeySpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePkixEddsaKeySpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkixEddsaKeySpecAllOf(val *PkixEddsaKeySpecAllOf) *NullablePkixEddsaKeySpecAllOf {
	return &NullablePkixEddsaKeySpecAllOf{value: val, isSet: true}
}

func (v NullablePkixEddsaKeySpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkixEddsaKeySpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
