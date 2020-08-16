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

// IppoolPoolAllOf Definition of the list of properties defined in 'ippool.Pool', excluding properties defined in parent classes.
type IppoolPoolAllOf struct {
	IpV4Blocks *[]IppoolIpBlock  `json:"IpV4Blocks,omitempty" yaml:"IpV4Blocks,omitempty"`
	IpV4Config *IppoolIpV4Config `json:"IpV4Config,omitempty" yaml:"IpV4Config,omitempty"`
	// Number of IPv4 addresses currently in use.
	V4Assigned *int64 `json:"V4Assigned,omitempty" yaml:"V4Assigned,omitempty"`
	// Number of IPv4 addresses in this pool.
	V4Size *int64 `json:"V4Size,omitempty" yaml:"V4Size,omitempty"`
	// Number of IPv6 addresses currently in use.
	V6Assigned *int64 `json:"V6Assigned,omitempty" yaml:"V6Assigned,omitempty"`
	// Number of IPv6 addresses in this pool.
	V6Size       *int64                                `json:"V6Size,omitempty" yaml:"V6Size,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
	// An array of relationships to ippoolShadowPool resources.
	ShadowPools []IppoolShadowPoolRelationship `json:"ShadowPools,omitempty" yaml:"ShadowPools,omitempty"`
}

// NewIppoolPoolAllOf instantiates a new IppoolPoolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolPoolAllOf() *IppoolPoolAllOf {
	this := IppoolPoolAllOf{}
	return &this
}

// NewIppoolPoolAllOfWithDefaults instantiates a new IppoolPoolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolPoolAllOfWithDefaults() *IppoolPoolAllOf {
	this := IppoolPoolAllOf{}
	return &this
}

// GetIpV4Blocks returns the IpV4Blocks field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetIpV4Blocks() []IppoolIpBlock {
	if o == nil || o.IpV4Blocks == nil {
		var ret []IppoolIpBlock
		return ret
	}
	return *o.IpV4Blocks
}

// GetIpV4BlocksOk returns a tuple with the IpV4Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetIpV4BlocksOk() (*[]IppoolIpBlock, bool) {
	if o == nil || o.IpV4Blocks == nil {
		return nil, false
	}
	return o.IpV4Blocks, true
}

// HasIpV4Blocks returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasIpV4Blocks() bool {
	if o != nil && o.IpV4Blocks != nil {
		return true
	}

	return false
}

// SetIpV4Blocks gets a reference to the given []IppoolIpBlock and assigns it to the IpV4Blocks field.
func (o *IppoolPoolAllOf) SetIpV4Blocks(v []IppoolIpBlock) {
	o.IpV4Blocks = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given IppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolPoolAllOf) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config = &v
}

// GetV4Assigned returns the V4Assigned field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetV4Assigned() int64 {
	if o == nil || o.V4Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V4Assigned
}

// GetV4AssignedOk returns a tuple with the V4Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetV4AssignedOk() (*int64, bool) {
	if o == nil || o.V4Assigned == nil {
		return nil, false
	}
	return o.V4Assigned, true
}

// HasV4Assigned returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasV4Assigned() bool {
	if o != nil && o.V4Assigned != nil {
		return true
	}

	return false
}

// SetV4Assigned gets a reference to the given int64 and assigns it to the V4Assigned field.
func (o *IppoolPoolAllOf) SetV4Assigned(v int64) {
	o.V4Assigned = &v
}

// GetV4Size returns the V4Size field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetV4Size() int64 {
	if o == nil || o.V4Size == nil {
		var ret int64
		return ret
	}
	return *o.V4Size
}

// GetV4SizeOk returns a tuple with the V4Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetV4SizeOk() (*int64, bool) {
	if o == nil || o.V4Size == nil {
		return nil, false
	}
	return o.V4Size, true
}

// HasV4Size returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasV4Size() bool {
	if o != nil && o.V4Size != nil {
		return true
	}

	return false
}

// SetV4Size gets a reference to the given int64 and assigns it to the V4Size field.
func (o *IppoolPoolAllOf) SetV4Size(v int64) {
	o.V4Size = &v
}

// GetV6Assigned returns the V6Assigned field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetV6Assigned() int64 {
	if o == nil || o.V6Assigned == nil {
		var ret int64
		return ret
	}
	return *o.V6Assigned
}

// GetV6AssignedOk returns a tuple with the V6Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetV6AssignedOk() (*int64, bool) {
	if o == nil || o.V6Assigned == nil {
		return nil, false
	}
	return o.V6Assigned, true
}

// HasV6Assigned returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasV6Assigned() bool {
	if o != nil && o.V6Assigned != nil {
		return true
	}

	return false
}

// SetV6Assigned gets a reference to the given int64 and assigns it to the V6Assigned field.
func (o *IppoolPoolAllOf) SetV6Assigned(v int64) {
	o.V6Assigned = &v
}

// GetV6Size returns the V6Size field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetV6Size() int64 {
	if o == nil || o.V6Size == nil {
		var ret int64
		return ret
	}
	return *o.V6Size
}

// GetV6SizeOk returns a tuple with the V6Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetV6SizeOk() (*int64, bool) {
	if o == nil || o.V6Size == nil {
		return nil, false
	}
	return o.V6Size, true
}

// HasV6Size returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasV6Size() bool {
	if o != nil && o.V6Size != nil {
		return true
	}

	return false
}

// SetV6Size gets a reference to the given int64 and assigns it to the V6Size field.
func (o *IppoolPoolAllOf) SetV6Size(v int64) {
	o.V6Size = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *IppoolPoolAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *IppoolPoolAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetShadowPools returns the ShadowPools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolPoolAllOf) GetShadowPools() []IppoolShadowPoolRelationship {
	if o == nil {
		var ret []IppoolShadowPoolRelationship
		return ret
	}
	return o.ShadowPools
}

// GetShadowPoolsOk returns a tuple with the ShadowPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolPoolAllOf) GetShadowPoolsOk() (*[]IppoolShadowPoolRelationship, bool) {
	if o == nil || o.ShadowPools == nil {
		return nil, false
	}
	return &o.ShadowPools, true
}

// HasShadowPools returns a boolean if a field has been set.
func (o *IppoolPoolAllOf) HasShadowPools() bool {
	if o != nil && o.ShadowPools != nil {
		return true
	}

	return false
}

// SetShadowPools gets a reference to the given []IppoolShadowPoolRelationship and assigns it to the ShadowPools field.
func (o *IppoolPoolAllOf) SetShadowPools(v []IppoolShadowPoolRelationship) {
	o.ShadowPools = v
}

func (o IppoolPoolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpV4Blocks != nil {
		toSerialize["IpV4Blocks"] = o.IpV4Blocks
	}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}
	if o.V4Assigned != nil {
		toSerialize["V4Assigned"] = o.V4Assigned
	}
	if o.V4Size != nil {
		toSerialize["V4Size"] = o.V4Size
	}
	if o.V6Assigned != nil {
		toSerialize["V6Assigned"] = o.V6Assigned
	}
	if o.V6Size != nil {
		toSerialize["V6Size"] = o.V6Size
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ShadowPools != nil {
		toSerialize["ShadowPools"] = o.ShadowPools
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolPoolAllOf struct {
	value *IppoolPoolAllOf
	isSet bool
}

func (v NullableIppoolPoolAllOf) Get() *IppoolPoolAllOf {
	return v.value
}

func (v *NullableIppoolPoolAllOf) Set(val *IppoolPoolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolPoolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolPoolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolPoolAllOf(val *IppoolPoolAllOf) *NullableIppoolPoolAllOf {
	return &NullableIppoolPoolAllOf{value: val, isSet: true}
}

func (v NullableIppoolPoolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolPoolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
