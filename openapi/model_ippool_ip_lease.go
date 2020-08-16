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

// IppoolIpLease IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
type IppoolIpLease struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// IPv4 Address given as a lease to an external entity like server profiles.
	IpV4Address      *string                       `json:"IpV4Address,omitempty" yaml:"IpV4Address,omitempty"`
	IpV4Config       *IppoolIpV4Config             `json:"IpV4Config,omitempty" yaml:"IpV4Config,omitempty"`
	AssignedToEntity *MoBaseMoRelationship         `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	Pool             *IppoolPoolRelationship       `json:"Pool,omitempty" yaml:"Pool,omitempty"`
	PoolMember       *IppoolPoolMemberRelationship `json:"PoolMember,omitempty" yaml:"PoolMember,omitempty"`
	Universe         *IppoolUniverseRelationship   `json:"Universe,omitempty" yaml:"Universe,omitempty"`
	Vrf              *VrfVrfRelationship           `json:"Vrf,omitempty" yaml:"Vrf,omitempty"`
}

// NewIppoolIpLease instantiates a new IppoolIpLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpLease() *IppoolIpLease {
	this := IppoolIpLease{}
	return &this
}

// NewIppoolIpLeaseWithDefaults instantiates a new IppoolIpLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpLeaseWithDefaults() *IppoolIpLease {
	this := IppoolIpLease{}
	return &this
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolIpLease) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolIpLease) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV4Config returns the IpV4Config field value if set, zero value otherwise.
func (o *IppoolIpLease) GetIpV4Config() IppoolIpV4Config {
	if o == nil || o.IpV4Config == nil {
		var ret IppoolIpV4Config
		return ret
	}
	return *o.IpV4Config
}

// GetIpV4ConfigOk returns a tuple with the IpV4Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetIpV4ConfigOk() (*IppoolIpV4Config, bool) {
	if o == nil || o.IpV4Config == nil {
		return nil, false
	}
	return o.IpV4Config, true
}

// HasIpV4Config returns a boolean if a field has been set.
func (o *IppoolIpLease) HasIpV4Config() bool {
	if o != nil && o.IpV4Config != nil {
		return true
	}

	return false
}

// SetIpV4Config gets a reference to the given IppoolIpV4Config and assigns it to the IpV4Config field.
func (o *IppoolIpLease) SetIpV4Config(v IppoolIpV4Config) {
	o.IpV4Config = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IppoolIpLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolIpLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolIpLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolIpLease) GetPool() IppoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolIpLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolPoolRelationship and assigns it to the Pool field.
func (o *IppoolIpLease) SetPool(v IppoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *IppoolIpLease) GetPoolMember() IppoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret IppoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetPoolMemberOk() (*IppoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *IppoolIpLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given IppoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *IppoolIpLease) SetPoolMember(v IppoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *IppoolIpLease) GetUniverse() IppoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret IppoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetUniverseOk() (*IppoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *IppoolIpLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given IppoolUniverseRelationship and assigns it to the Universe field.
func (o *IppoolIpLease) SetUniverse(v IppoolUniverseRelationship) {
	o.Universe = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *IppoolIpLease) GetVrf() VrfVrfRelationship {
	if o == nil || o.Vrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpLease) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolIpLease) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given VrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolIpLease) SetVrf(v VrfVrfRelationship) {
	o.Vrf = &v
}

func (o IppoolIpLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.IpV4Config != nil {
		toSerialize["IpV4Config"] = o.IpV4Config
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}
	return json.Marshal(toSerialize)
}

type NullableIppoolIpLease struct {
	value *IppoolIpLease
	isSet bool
}

func (v NullableIppoolIpLease) Get() *IppoolIpLease {
	return v.value
}

func (v *NullableIppoolIpLease) Set(val *IppoolIpLease) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpLease) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpLease(val *IppoolIpLease) *NullableIppoolIpLease {
	return &NullableIppoolIpLease{value: val, isSet: true}
}

func (v NullableIppoolIpLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}