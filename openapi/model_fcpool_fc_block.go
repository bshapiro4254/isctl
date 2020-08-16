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

// FcpoolFcBlock A block of contiguous WWN addresses that are part of a pool.
type FcpoolFcBlock struct {
	PoolAbstractBlock `yaml:"PoolAbstractBlock,inline"`
	IdBlock           *FcpoolBlock            `json:"IdBlock,omitempty" yaml:"IdBlock,omitempty"`
	Pool              *FcpoolPoolRelationship `json:"Pool,omitempty" yaml:"Pool,omitempty"`
}

// NewFcpoolFcBlock instantiates a new FcpoolFcBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolFcBlock() *FcpoolFcBlock {
	this := FcpoolFcBlock{}
	return &this
}

// NewFcpoolFcBlockWithDefaults instantiates a new FcpoolFcBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolFcBlockWithDefaults() *FcpoolFcBlock {
	this := FcpoolFcBlock{}
	return &this
}

// GetIdBlock returns the IdBlock field value if set, zero value otherwise.
func (o *FcpoolFcBlock) GetIdBlock() FcpoolBlock {
	if o == nil || o.IdBlock == nil {
		var ret FcpoolBlock
		return ret
	}
	return *o.IdBlock
}

// GetIdBlockOk returns a tuple with the IdBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetIdBlockOk() (*FcpoolBlock, bool) {
	if o == nil || o.IdBlock == nil {
		return nil, false
	}
	return o.IdBlock, true
}

// HasIdBlock returns a boolean if a field has been set.
func (o *FcpoolFcBlock) HasIdBlock() bool {
	if o != nil && o.IdBlock != nil {
		return true
	}

	return false
}

// SetIdBlock gets a reference to the given FcpoolBlock and assigns it to the IdBlock field.
func (o *FcpoolFcBlock) SetIdBlock(v FcpoolBlock) {
	o.IdBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolFcBlock) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolFcBlock) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolFcBlock) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

func (o FcpoolFcBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlock, errPoolAbstractBlock := json.Marshal(o.PoolAbstractBlock)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	errPoolAbstractBlock = json.Unmarshal([]byte(serializedPoolAbstractBlock), &toSerialize)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	if o.IdBlock != nil {
		toSerialize["IdBlock"] = o.IdBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	return json.Marshal(toSerialize)
}

type NullableFcpoolFcBlock struct {
	value *FcpoolFcBlock
	isSet bool
}

func (v NullableFcpoolFcBlock) Get() *FcpoolFcBlock {
	return v.value
}

func (v *NullableFcpoolFcBlock) Set(val *FcpoolFcBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolFcBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolFcBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolFcBlock(val *FcpoolFcBlock) *NullableFcpoolFcBlock {
	return &NullableFcpoolFcBlock{value: val, isSet: true}
}

func (v NullableFcpoolFcBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolFcBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}