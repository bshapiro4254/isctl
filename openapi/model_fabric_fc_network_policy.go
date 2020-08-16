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

// FabricFcNetworkPolicy A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.
type FabricFcNetworkPolicy struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	// Enable or Disable Trunking on all of configured FC uplink ports.
	EnableTrunking *bool                                 `json:"EnableTrunking,omitempty" yaml:"EnableTrunking,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles []FabricSwitchProfileRelationship `json:"Profiles,omitempty" yaml:"Profiles,omitempty"`
}

// NewFabricFcNetworkPolicy instantiates a new FabricFcNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcNetworkPolicy() *FabricFcNetworkPolicy {
	this := FabricFcNetworkPolicy{}
	return &this
}

// NewFabricFcNetworkPolicyWithDefaults instantiates a new FabricFcNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcNetworkPolicyWithDefaults() *FabricFcNetworkPolicy {
	this := FabricFcNetworkPolicy{}
	return &this
}

// GetEnableTrunking returns the EnableTrunking field value if set, zero value otherwise.
func (o *FabricFcNetworkPolicy) GetEnableTrunking() bool {
	if o == nil || o.EnableTrunking == nil {
		var ret bool
		return ret
	}
	return *o.EnableTrunking
}

// GetEnableTrunkingOk returns a tuple with the EnableTrunking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetEnableTrunkingOk() (*bool, bool) {
	if o == nil || o.EnableTrunking == nil {
		return nil, false
	}
	return o.EnableTrunking, true
}

// HasEnableTrunking returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasEnableTrunking() bool {
	if o != nil && o.EnableTrunking != nil {
		return true
	}

	return false
}

// SetEnableTrunking gets a reference to the given bool and assigns it to the EnableTrunking field.
func (o *FabricFcNetworkPolicy) SetEnableTrunking(v bool) {
	o.EnableTrunking = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricFcNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricFcNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricFcNetworkPolicy) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricFcNetworkPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricFcNetworkPolicy) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.EnableTrunking != nil {
		toSerialize["EnableTrunking"] = o.EnableTrunking
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	return json.Marshal(toSerialize)
}

type NullableFabricFcNetworkPolicy struct {
	value *FabricFcNetworkPolicy
	isSet bool
}

func (v NullableFabricFcNetworkPolicy) Get() *FabricFcNetworkPolicy {
	return v.value
}

func (v *NullableFabricFcNetworkPolicy) Set(val *FabricFcNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcNetworkPolicy(val *FabricFcNetworkPolicy) *NullableFabricFcNetworkPolicy {
	return &NullableFabricFcNetworkPolicy{value: val, isSet: true}
}

func (v NullableFabricFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}