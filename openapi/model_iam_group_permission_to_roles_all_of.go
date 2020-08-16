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

// IamGroupPermissionToRolesAllOf Definition of the list of properties defined in 'iam.GroupPermissionToRoles', excluding properties defined in parent classes.
type IamGroupPermissionToRolesAllOf struct {
	Group *CmrfCmRf   `json:"Group,omitempty" yaml:"Group,omitempty"`
	Orgs  *[]CmrfCmRf `json:"Orgs,omitempty" yaml:"Orgs,omitempty"`
}

// NewIamGroupPermissionToRolesAllOf instantiates a new IamGroupPermissionToRolesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamGroupPermissionToRolesAllOf() *IamGroupPermissionToRolesAllOf {
	this := IamGroupPermissionToRolesAllOf{}
	return &this
}

// NewIamGroupPermissionToRolesAllOfWithDefaults instantiates a new IamGroupPermissionToRolesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamGroupPermissionToRolesAllOfWithDefaults() *IamGroupPermissionToRolesAllOf {
	this := IamGroupPermissionToRolesAllOf{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IamGroupPermissionToRolesAllOf) GetGroup() CmrfCmRf {
	if o == nil || o.Group == nil {
		var ret CmrfCmRf
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamGroupPermissionToRolesAllOf) GetGroupOk() (*CmrfCmRf, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IamGroupPermissionToRolesAllOf) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given CmrfCmRf and assigns it to the Group field.
func (o *IamGroupPermissionToRolesAllOf) SetGroup(v CmrfCmRf) {
	o.Group = &v
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *IamGroupPermissionToRolesAllOf) GetOrgs() []CmrfCmRf {
	if o == nil || o.Orgs == nil {
		var ret []CmrfCmRf
		return ret
	}
	return *o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamGroupPermissionToRolesAllOf) GetOrgsOk() (*[]CmrfCmRf, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *IamGroupPermissionToRolesAllOf) HasOrgs() bool {
	if o != nil && o.Orgs != nil {
		return true
	}

	return false
}

// SetOrgs gets a reference to the given []CmrfCmRf and assigns it to the Orgs field.
func (o *IamGroupPermissionToRolesAllOf) SetOrgs(v []CmrfCmRf) {
	o.Orgs = &v
}

func (o IamGroupPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.Orgs != nil {
		toSerialize["Orgs"] = o.Orgs
	}
	return json.Marshal(toSerialize)
}

type NullableIamGroupPermissionToRolesAllOf struct {
	value *IamGroupPermissionToRolesAllOf
	isSet bool
}

func (v NullableIamGroupPermissionToRolesAllOf) Get() *IamGroupPermissionToRolesAllOf {
	return v.value
}

func (v *NullableIamGroupPermissionToRolesAllOf) Set(val *IamGroupPermissionToRolesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamGroupPermissionToRolesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamGroupPermissionToRolesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamGroupPermissionToRolesAllOf(val *IamGroupPermissionToRolesAllOf) *NullableIamGroupPermissionToRolesAllOf {
	return &NullableIamGroupPermissionToRolesAllOf{value: val, isSet: true}
}

func (v NullableIamGroupPermissionToRolesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamGroupPermissionToRolesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
