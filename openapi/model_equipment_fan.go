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

// EquipmentFan Fan in a Fabric Interconnect / Chassis / RackUnit.
type EquipmentFan struct {
	EquipmentBase `yaml:"EquipmentBase,inline"`
	// This field is to provide description for the fan.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// This field acts as the identifier for this particular Fan, within the Fabric Interconnect.
	FanId *int64 `json:"FanId,omitempty" yaml:"FanId,omitempty"`
	// This field is used to identify the Fan Module to which this Fan belongs.
	FanModuleId *int64 `json:"FanModuleId,omitempty" yaml:"FanModuleId,omitempty"`
	// Fan module Identifier for the fan.
	ModuleId *int64 `json:"ModuleId,omitempty" yaml:"ModuleId,omitempty"`
	// This field is used to indicate this fan unit's operational state.
	OperState *string `json:"OperState,omitempty" yaml:"OperState,omitempty"`
	// This field identifies the Part Number for this Fan Unit.
	PartNumber *string `json:"PartNumber,omitempty" yaml:"PartNumber,omitempty"`
	// This field identifies the Product ID for the fans.
	Pid *string `json:"Pid,omitempty" yaml:"Pid,omitempty"`
	// This field is used to indicate this fan unit's presence.
	Presence *string `json:"Presence,omitempty" yaml:"Presence,omitempty"`
	// This field identifies the Stockkeeping Unit for this Fan Unit.
	Sku *string `json:"Sku,omitempty" yaml:"Sku,omitempty"`
	// Tray identifier for the fan module.
	TrayId *int64 `json:"TrayId,omitempty" yaml:"TrayId,omitempty"`
	// This field identifies the Vendor ID for this Fan Unit.
	Vid                 *string                              `json:"Vid,omitempty" yaml:"Vid,omitempty"`
	EquipmentFanModule  *EquipmentFanModuleRelationship      `json:"EquipmentFanModule,omitempty" yaml:"EquipmentFanModule,omitempty"`
	EquipmentFex        *EquipmentFexRelationship            `json:"EquipmentFex,omitempty" yaml:"EquipmentFex,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewEquipmentFan instantiates a new EquipmentFan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFan() *EquipmentFan {
	this := EquipmentFan{}
	return &this
}

// NewEquipmentFanWithDefaults instantiates a new EquipmentFan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFanWithDefaults() *EquipmentFan {
	this := EquipmentFan{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentFan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentFan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentFan) SetDescription(v string) {
	o.Description = &v
}

// GetFanId returns the FanId field value if set, zero value otherwise.
func (o *EquipmentFan) GetFanId() int64 {
	if o == nil || o.FanId == nil {
		var ret int64
		return ret
	}
	return *o.FanId
}

// GetFanIdOk returns a tuple with the FanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetFanIdOk() (*int64, bool) {
	if o == nil || o.FanId == nil {
		return nil, false
	}
	return o.FanId, true
}

// HasFanId returns a boolean if a field has been set.
func (o *EquipmentFan) HasFanId() bool {
	if o != nil && o.FanId != nil {
		return true
	}

	return false
}

// SetFanId gets a reference to the given int64 and assigns it to the FanId field.
func (o *EquipmentFan) SetFanId(v int64) {
	o.FanId = &v
}

// GetFanModuleId returns the FanModuleId field value if set, zero value otherwise.
func (o *EquipmentFan) GetFanModuleId() int64 {
	if o == nil || o.FanModuleId == nil {
		var ret int64
		return ret
	}
	return *o.FanModuleId
}

// GetFanModuleIdOk returns a tuple with the FanModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetFanModuleIdOk() (*int64, bool) {
	if o == nil || o.FanModuleId == nil {
		return nil, false
	}
	return o.FanModuleId, true
}

// HasFanModuleId returns a boolean if a field has been set.
func (o *EquipmentFan) HasFanModuleId() bool {
	if o != nil && o.FanModuleId != nil {
		return true
	}

	return false
}

// SetFanModuleId gets a reference to the given int64 and assigns it to the FanModuleId field.
func (o *EquipmentFan) SetFanModuleId(v int64) {
	o.FanModuleId = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentFan) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentFan) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentFan) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentFan) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentFan) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentFan) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentFan) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentFan) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentFan) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentFan) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentFan) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentFan) SetPid(v string) {
	o.Pid = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentFan) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentFan) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentFan) SetPresence(v string) {
	o.Presence = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentFan) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentFan) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentFan) SetSku(v string) {
	o.Sku = &v
}

// GetTrayId returns the TrayId field value if set, zero value otherwise.
func (o *EquipmentFan) GetTrayId() int64 {
	if o == nil || o.TrayId == nil {
		var ret int64
		return ret
	}
	return *o.TrayId
}

// GetTrayIdOk returns a tuple with the TrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetTrayIdOk() (*int64, bool) {
	if o == nil || o.TrayId == nil {
		return nil, false
	}
	return o.TrayId, true
}

// HasTrayId returns a boolean if a field has been set.
func (o *EquipmentFan) HasTrayId() bool {
	if o != nil && o.TrayId != nil {
		return true
	}

	return false
}

// SetTrayId gets a reference to the given int64 and assigns it to the TrayId field.
func (o *EquipmentFan) SetTrayId(v int64) {
	o.TrayId = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentFan) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentFan) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentFan) SetVid(v string) {
	o.Vid = &v
}

// GetEquipmentFanModule returns the EquipmentFanModule field value if set, zero value otherwise.
func (o *EquipmentFan) GetEquipmentFanModule() EquipmentFanModuleRelationship {
	if o == nil || o.EquipmentFanModule == nil {
		var ret EquipmentFanModuleRelationship
		return ret
	}
	return *o.EquipmentFanModule
}

// GetEquipmentFanModuleOk returns a tuple with the EquipmentFanModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetEquipmentFanModuleOk() (*EquipmentFanModuleRelationship, bool) {
	if o == nil || o.EquipmentFanModule == nil {
		return nil, false
	}
	return o.EquipmentFanModule, true
}

// HasEquipmentFanModule returns a boolean if a field has been set.
func (o *EquipmentFan) HasEquipmentFanModule() bool {
	if o != nil && o.EquipmentFanModule != nil {
		return true
	}

	return false
}

// SetEquipmentFanModule gets a reference to the given EquipmentFanModuleRelationship and assigns it to the EquipmentFanModule field.
func (o *EquipmentFan) SetEquipmentFanModule(v EquipmentFanModuleRelationship) {
	o.EquipmentFanModule = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentFan) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentFan) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentFan) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentFan) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentFan) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentFan) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentFan) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFan) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentFan) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentFan) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentFan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.FanId != nil {
		toSerialize["FanId"] = o.FanId
	}
	if o.FanModuleId != nil {
		toSerialize["FanModuleId"] = o.FanModuleId
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.TrayId != nil {
		toSerialize["TrayId"] = o.TrayId
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.EquipmentFanModule != nil {
		toSerialize["EquipmentFanModule"] = o.EquipmentFanModule
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentFan struct {
	value *EquipmentFan
	isSet bool
}

func (v NullableEquipmentFan) Get() *EquipmentFan {
	return v.value
}

func (v *NullableEquipmentFan) Set(val *EquipmentFan) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFan) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFan(val *EquipmentFan) *NullableEquipmentFan {
	return &NullableEquipmentFan{value: val, isSet: true}
}

func (v NullableEquipmentFan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
