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

// UuidpoolUuidLease UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
type UuidpoolUuidLease struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// UUID Prefix+Suffix numbers.
	Uuid             *string                         `json:"Uuid,omitempty" yaml:"Uuid,omitempty"`
	AssignedToEntity *MoBaseMoRelationship           `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	Pool             *UuidpoolPoolRelationship       `json:"Pool,omitempty" yaml:"Pool,omitempty"`
	PoolMember       *UuidpoolPoolMemberRelationship `json:"PoolMember,omitempty" yaml:"PoolMember,omitempty"`
	Universe         *UuidpoolUniverseRelationship   `json:"Universe,omitempty" yaml:"Universe,omitempty"`
}

// NewUuidpoolUuidLease instantiates a new UuidpoolUuidLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolUuidLease() *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	return &this
}

// NewUuidpoolUuidLeaseWithDefaults instantiates a new UuidpoolUuidLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolUuidLeaseWithDefaults() *UuidpoolUuidLease {
	this := UuidpoolUuidLease{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *UuidpoolUuidLease) SetUuid(v string) {
	o.Uuid = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *UuidpoolUuidLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetPool() UuidpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolUuidLease) SetPool(v UuidpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetPoolMember() UuidpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret UuidpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetPoolMemberOk() (*UuidpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given UuidpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *UuidpoolUuidLease) SetPoolMember(v UuidpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *UuidpoolUuidLease) GetUniverse() UuidpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret UuidpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolUuidLease) GetUniverseOk() (*UuidpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *UuidpoolUuidLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given UuidpoolUniverseRelationship and assigns it to the Universe field.
func (o *UuidpoolUuidLease) SetUniverse(v UuidpoolUniverseRelationship) {
	o.Universe = &v
}

func (o UuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
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
	return json.Marshal(toSerialize)
}

type NullableUuidpoolUuidLease struct {
	value *UuidpoolUuidLease
	isSet bool
}

func (v NullableUuidpoolUuidLease) Get() *UuidpoolUuidLease {
	return v.value
}

func (v *NullableUuidpoolUuidLease) Set(val *UuidpoolUuidLease) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolUuidLease) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolUuidLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolUuidLease(val *UuidpoolUuidLease) *NullableUuidpoolUuidLease {
	return &NullableUuidpoolUuidLease{value: val, isSet: true}
}

func (v NullableUuidpoolUuidLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolUuidLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}