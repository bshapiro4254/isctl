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

// PoolAbstractPool Pool represents a  collection of identifiers that can be allocated to server profiles.
type PoolAbstractPool struct {
	PolicyAbstractPolicy `yaml:"PolicyAbstractPolicy,inline"`
	// Number of IDs that are currently assigned.
	Assigned *int64 `json:"Assigned,omitempty" yaml:"Assigned,omitempty"`
	// Assignment order decides the order in which the next identifier is allocated. * `sequential` - Identifiers are assigned in a sequential order. * `default` - Assignment order is decided by the system.
	AssignmentOrder *string `json:"AssignmentOrder,omitempty" yaml:"AssignmentOrder,omitempty"`
	// Total number of identifiers in this pool.
	Size *int64 `json:"Size,omitempty" yaml:"Size,omitempty"`
}

// NewPoolAbstractPool instantiates a new PoolAbstractPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolAbstractPool() *PoolAbstractPool {
	this := PoolAbstractPool{}
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewPoolAbstractPoolWithDefaults instantiates a new PoolAbstractPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolAbstractPoolWithDefaults() *PoolAbstractPool {
	this := PoolAbstractPool{}
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetAssigned() int64 {
	if o == nil || o.Assigned == nil {
		var ret int64
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetAssignedOk() (*int64, bool) {
	if o == nil || o.Assigned == nil {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasAssigned() bool {
	if o != nil && o.Assigned != nil {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int64 and assigns it to the Assigned field.
func (o *PoolAbstractPool) SetAssigned(v int64) {
	o.Assigned = &v
}

// GetAssignmentOrder returns the AssignmentOrder field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetAssignmentOrder() string {
	if o == nil || o.AssignmentOrder == nil {
		var ret string
		return ret
	}
	return *o.AssignmentOrder
}

// GetAssignmentOrderOk returns a tuple with the AssignmentOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetAssignmentOrderOk() (*string, bool) {
	if o == nil || o.AssignmentOrder == nil {
		return nil, false
	}
	return o.AssignmentOrder, true
}

// HasAssignmentOrder returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasAssignmentOrder() bool {
	if o != nil && o.AssignmentOrder != nil {
		return true
	}

	return false
}

// SetAssignmentOrder gets a reference to the given string and assigns it to the AssignmentOrder field.
func (o *PoolAbstractPool) SetAssignmentOrder(v string) {
	o.AssignmentOrder = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PoolAbstractPool) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolAbstractPool) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PoolAbstractPool) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *PoolAbstractPool) SetSize(v int64) {
	o.Size = &v
}

func (o PoolAbstractPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if o.Assigned != nil {
		toSerialize["Assigned"] = o.Assigned
	}
	if o.AssignmentOrder != nil {
		toSerialize["AssignmentOrder"] = o.AssignmentOrder
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullablePoolAbstractPool struct {
	value *PoolAbstractPool
	isSet bool
}

func (v NullablePoolAbstractPool) Get() *PoolAbstractPool {
	return v.value
}

func (v *NullablePoolAbstractPool) Set(val *PoolAbstractPool) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolAbstractPool) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolAbstractPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolAbstractPool(val *PoolAbstractPool) *NullablePoolAbstractPool {
	return &NullablePoolAbstractPool{value: val, isSet: true}
}

func (v NullablePoolAbstractPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolAbstractPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
