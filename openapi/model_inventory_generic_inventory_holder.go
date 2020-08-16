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

// InventoryGenericInventoryHolder A container class for generic inventory.
type InventoryGenericInventoryHolder struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// The endpoint represented by this holder.
	Endpoint        *string                      `json:"Endpoint,omitempty" yaml:"Endpoint,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty" yaml:"ComputeBlade,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty" yaml:"ComputeRackUnit,omitempty"`
	// An array of relationships to inventoryGenericInventory resources.
	GenericInventory    []InventoryGenericInventoryRelationship `json:"GenericInventory,omitempty" yaml:"GenericInventory,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewInventoryGenericInventoryHolder instantiates a new InventoryGenericInventoryHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGenericInventoryHolder() *InventoryGenericInventoryHolder {
	this := InventoryGenericInventoryHolder{}
	return &this
}

// NewInventoryGenericInventoryHolderWithDefaults instantiates a new InventoryGenericInventoryHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGenericInventoryHolderWithDefaults() *InventoryGenericInventoryHolder {
	this := InventoryGenericInventoryHolder{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *InventoryGenericInventoryHolder) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *InventoryGenericInventoryHolder) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *InventoryGenericInventoryHolder) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetGenericInventory returns the GenericInventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryGenericInventoryHolder) GetGenericInventory() []InventoryGenericInventoryRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryRelationship
		return ret
	}
	return o.GenericInventory
}

// GetGenericInventoryOk returns a tuple with the GenericInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryGenericInventoryHolder) GetGenericInventoryOk() (*[]InventoryGenericInventoryRelationship, bool) {
	if o == nil || o.GenericInventory == nil {
		return nil, false
	}
	return &o.GenericInventory, true
}

// HasGenericInventory returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasGenericInventory() bool {
	if o != nil && o.GenericInventory != nil {
		return true
	}

	return false
}

// SetGenericInventory gets a reference to the given []InventoryGenericInventoryRelationship and assigns it to the GenericInventory field.
func (o *InventoryGenericInventoryHolder) SetGenericInventory(v []InventoryGenericInventoryRelationship) {
	o.GenericInventory = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *InventoryGenericInventoryHolder) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolder) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolder) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolder) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryGenericInventoryHolder) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryGenericInventoryHolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.Endpoint != nil {
		toSerialize["Endpoint"] = o.Endpoint
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.GenericInventory != nil {
		toSerialize["GenericInventory"] = o.GenericInventory
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryGenericInventoryHolder struct {
	value *InventoryGenericInventoryHolder
	isSet bool
}

func (v NullableInventoryGenericInventoryHolder) Get() *InventoryGenericInventoryHolder {
	return v.value
}

func (v *NullableInventoryGenericInventoryHolder) Set(val *InventoryGenericInventoryHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventoryHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventoryHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventoryHolder(val *InventoryGenericInventoryHolder) *NullableInventoryGenericInventoryHolder {
	return &NullableInventoryGenericInventoryHolder{value: val, isSet: true}
}

func (v NullableInventoryGenericInventoryHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventoryHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
