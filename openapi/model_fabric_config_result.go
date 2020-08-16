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

// FabricConfigResult This provides overall state and detailed information for the deploy and validation profile configuration results.
type FabricConfigResult struct {
	PolicyAbstractConfigResult `yaml:"PolicyAbstractConfigResult,inline"`
	Profile                    *FabricSwitchProfileRelationship `json:"Profile,omitempty" yaml:"Profile,omitempty"`
	// An array of relationships to fabricConfigResultEntry resources.
	ResultEntries []FabricConfigResultEntryRelationship `json:"ResultEntries,omitempty" yaml:"ResultEntries,omitempty"`
}

// NewFabricConfigResult instantiates a new FabricConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricConfigResult() *FabricConfigResult {
	this := FabricConfigResult{}
	return &this
}

// NewFabricConfigResultWithDefaults instantiates a new FabricConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricConfigResultWithDefaults() *FabricConfigResult {
	this := FabricConfigResult{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *FabricConfigResult) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricConfigResult) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricConfigResult) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given FabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricConfigResult) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricConfigResult) GetResultEntries() []FabricConfigResultEntryRelationship {
	if o == nil {
		var ret []FabricConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricConfigResult) GetResultEntriesOk() (*[]FabricConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return &o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *FabricConfigResult) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []FabricConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *FabricConfigResult) SetResultEntries(v []FabricConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o FabricConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return []byte{}, errPolicyAbstractConfigResult
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}
	return json.Marshal(toSerialize)
}

type NullableFabricConfigResult struct {
	value *FabricConfigResult
	isSet bool
}

func (v NullableFabricConfigResult) Get() *FabricConfigResult {
	return v.value
}

func (v *NullableFabricConfigResult) Set(val *FabricConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricConfigResult(val *FabricConfigResult) *NullableFabricConfigResult {
	return &NullableFabricConfigResult{value: val, isSet: true}
}

func (v NullableFabricConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}