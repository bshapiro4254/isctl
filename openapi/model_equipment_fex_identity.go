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

// EquipmentFexIdentity FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.
type EquipmentFexIdentity struct {
	EquipmentPhysicalIdentity `yaml:"EquipmentPhysicalIdentity,inline"`
	// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
	SwitchId       *int64                      `json:"SwitchId,omitempty" yaml:"SwitchId,omitempty"`
	Fex            *EquipmentFexRelationship   `json:"Fex,omitempty" yaml:"Fex,omitempty"`
	NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty" yaml:"NetworkElement,omitempty"`
}

// NewEquipmentFexIdentity instantiates a new EquipmentFexIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFexIdentity() *EquipmentFexIdentity {
	this := EquipmentFexIdentity{}
	return &this
}

// NewEquipmentFexIdentityWithDefaults instantiates a new EquipmentFexIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFexIdentityWithDefaults() *EquipmentFexIdentity {
	this := EquipmentFexIdentity{}
	return &this
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentFexIdentity) GetSwitchId() int64 {
	if o == nil || o.SwitchId == nil {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetSwitchIdOk() (*int64, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentFexIdentity) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetFex returns the Fex field value if set, zero value otherwise.
func (o *EquipmentFexIdentity) GetFex() EquipmentFexRelationship {
	if o == nil || o.Fex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.Fex
}

// GetFexOk returns a tuple with the Fex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.Fex == nil {
		return nil, false
	}
	return o.Fex, true
}

// HasFex returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasFex() bool {
	if o != nil && o.Fex != nil {
		return true
	}

	return false
}

// SetFex gets a reference to the given EquipmentFexRelationship and assigns it to the Fex field.
func (o *EquipmentFexIdentity) SetFex(v EquipmentFexRelationship) {
	o.Fex = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentFexIdentity) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexIdentity) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentFexIdentity) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentFexIdentity) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o EquipmentFexIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Fex != nil {
		toSerialize["Fex"] = o.Fex
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentFexIdentity struct {
	value *EquipmentFexIdentity
	isSet bool
}

func (v NullableEquipmentFexIdentity) Get() *EquipmentFexIdentity {
	return v.value
}

func (v *NullableEquipmentFexIdentity) Set(val *EquipmentFexIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFexIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFexIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFexIdentity(val *EquipmentFexIdentity) *NullableEquipmentFexIdentity {
	return &NullableEquipmentFexIdentity{value: val, isSet: true}
}

func (v NullableEquipmentFexIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFexIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
