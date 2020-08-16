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

// FcpoolPoolMember PoolMember represents a single WWN ID that is part of a pool.
type FcpoolPoolMember struct {
	PoolAbstractPoolMember `yaml:"PoolAbstractPoolMember,inline"`
	// WWN ID of this pool member.
	WwnId            *string                    `json:"WwnId,omitempty" yaml:"WwnId,omitempty"`
	AssignedToEntity *MoBaseMoRelationship      `json:"AssignedToEntity,omitempty" yaml:"AssignedToEntity,omitempty"`
	BlockHead        *FcpoolFcBlockRelationship `json:"BlockHead,omitempty" yaml:"BlockHead,omitempty"`
	Peer             *FcpoolLeaseRelationship   `json:"Peer,omitempty" yaml:"Peer,omitempty"`
	Pool             *FcpoolPoolRelationship    `json:"Pool,omitempty" yaml:"Pool,omitempty"`
}

// NewFcpoolPoolMember instantiates a new FcpoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolPoolMember() *FcpoolPoolMember {
	this := FcpoolPoolMember{}
	return &this
}

// NewFcpoolPoolMemberWithDefaults instantiates a new FcpoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolPoolMemberWithDefaults() *FcpoolPoolMember {
	this := FcpoolPoolMember{}
	return &this
}

// GetWwnId returns the WwnId field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetWwnId() string {
	if o == nil || o.WwnId == nil {
		var ret string
		return ret
	}
	return *o.WwnId
}

// GetWwnIdOk returns a tuple with the WwnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetWwnIdOk() (*string, bool) {
	if o == nil || o.WwnId == nil {
		return nil, false
	}
	return o.WwnId, true
}

// HasWwnId returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasWwnId() bool {
	if o != nil && o.WwnId != nil {
		return true
	}

	return false
}

// SetWwnId gets a reference to the given string and assigns it to the WwnId field.
func (o *FcpoolPoolMember) SetWwnId(v string) {
	o.WwnId = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *FcpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetBlockHead() FcpoolFcBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret FcpoolFcBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetBlockHeadOk() (*FcpoolFcBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given FcpoolFcBlockRelationship and assigns it to the BlockHead field.
func (o *FcpoolPoolMember) SetBlockHead(v FcpoolFcBlockRelationship) {
	o.BlockHead = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetPeer() FcpoolLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret FcpoolLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetPeerOk() (*FcpoolLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given FcpoolLeaseRelationship and assigns it to the Peer field.
func (o *FcpoolPoolMember) SetPeer(v FcpoolLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolPoolMember) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPoolMember) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolPoolMember) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolPoolMember) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

func (o FcpoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	if o.WwnId != nil {
		toSerialize["WwnId"] = o.WwnId
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	return json.Marshal(toSerialize)
}

type NullableFcpoolPoolMember struct {
	value *FcpoolPoolMember
	isSet bool
}

func (v NullableFcpoolPoolMember) Get() *FcpoolPoolMember {
	return v.value
}

func (v *NullableFcpoolPoolMember) Set(val *FcpoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPoolMember(val *FcpoolPoolMember) *NullableFcpoolPoolMember {
	return &NullableFcpoolPoolMember{value: val, isSet: true}
}

func (v NullableFcpoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
