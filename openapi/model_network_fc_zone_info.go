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

// NetworkFcZoneInfo FC Zone information of a Fabric Interconnect.
type NetworkFcZoneInfo struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// The number of Fibre Channel user zones defined on a Fabric Interconnect.
	UserZoneCount *int64 `json:"UserZoneCount,omitempty" yaml:"UserZoneCount,omitempty"`
	// The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect.
	UserZoneLimit *int64 `json:"UserZoneLimit,omitempty" yaml:"UserZoneLimit,omitempty"`
	// The number of Fibre Channel zones defined on a Fabric Interconnect.
	ZoneCount *int64 `json:"ZoneCount,omitempty" yaml:"ZoneCount,omitempty"`
	// The maximum number of Fibre Channel zones allowed on a Fabric Interconnect.
	ZoneLimit           *int64                               `json:"ZoneLimit,omitempty" yaml:"ZoneLimit,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	NetworkElement      *NetworkElementRelationship          `json:"NetworkElement,omitempty" yaml:"NetworkElement,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewNetworkFcZoneInfo instantiates a new NetworkFcZoneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFcZoneInfo() *NetworkFcZoneInfo {
	this := NetworkFcZoneInfo{}
	return &this
}

// NewNetworkFcZoneInfoWithDefaults instantiates a new NetworkFcZoneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFcZoneInfoWithDefaults() *NetworkFcZoneInfo {
	this := NetworkFcZoneInfo{}
	return &this
}

// GetUserZoneCount returns the UserZoneCount field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetUserZoneCount() int64 {
	if o == nil || o.UserZoneCount == nil {
		var ret int64
		return ret
	}
	return *o.UserZoneCount
}

// GetUserZoneCountOk returns a tuple with the UserZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetUserZoneCountOk() (*int64, bool) {
	if o == nil || o.UserZoneCount == nil {
		return nil, false
	}
	return o.UserZoneCount, true
}

// HasUserZoneCount returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasUserZoneCount() bool {
	if o != nil && o.UserZoneCount != nil {
		return true
	}

	return false
}

// SetUserZoneCount gets a reference to the given int64 and assigns it to the UserZoneCount field.
func (o *NetworkFcZoneInfo) SetUserZoneCount(v int64) {
	o.UserZoneCount = &v
}

// GetUserZoneLimit returns the UserZoneLimit field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetUserZoneLimit() int64 {
	if o == nil || o.UserZoneLimit == nil {
		var ret int64
		return ret
	}
	return *o.UserZoneLimit
}

// GetUserZoneLimitOk returns a tuple with the UserZoneLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetUserZoneLimitOk() (*int64, bool) {
	if o == nil || o.UserZoneLimit == nil {
		return nil, false
	}
	return o.UserZoneLimit, true
}

// HasUserZoneLimit returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasUserZoneLimit() bool {
	if o != nil && o.UserZoneLimit != nil {
		return true
	}

	return false
}

// SetUserZoneLimit gets a reference to the given int64 and assigns it to the UserZoneLimit field.
func (o *NetworkFcZoneInfo) SetUserZoneLimit(v int64) {
	o.UserZoneLimit = &v
}

// GetZoneCount returns the ZoneCount field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetZoneCount() int64 {
	if o == nil || o.ZoneCount == nil {
		var ret int64
		return ret
	}
	return *o.ZoneCount
}

// GetZoneCountOk returns a tuple with the ZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetZoneCountOk() (*int64, bool) {
	if o == nil || o.ZoneCount == nil {
		return nil, false
	}
	return o.ZoneCount, true
}

// HasZoneCount returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasZoneCount() bool {
	if o != nil && o.ZoneCount != nil {
		return true
	}

	return false
}

// SetZoneCount gets a reference to the given int64 and assigns it to the ZoneCount field.
func (o *NetworkFcZoneInfo) SetZoneCount(v int64) {
	o.ZoneCount = &v
}

// GetZoneLimit returns the ZoneLimit field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetZoneLimit() int64 {
	if o == nil || o.ZoneLimit == nil {
		var ret int64
		return ret
	}
	return *o.ZoneLimit
}

// GetZoneLimitOk returns a tuple with the ZoneLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetZoneLimitOk() (*int64, bool) {
	if o == nil || o.ZoneLimit == nil {
		return nil, false
	}
	return o.ZoneLimit, true
}

// HasZoneLimit returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasZoneLimit() bool {
	if o != nil && o.ZoneLimit != nil {
		return true
	}

	return false
}

// SetZoneLimit gets a reference to the given int64 and assigns it to the ZoneLimit field.
func (o *NetworkFcZoneInfo) SetZoneLimit(v int64) {
	o.ZoneLimit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *NetworkFcZoneInfo) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkFcZoneInfo) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkFcZoneInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkFcZoneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.UserZoneCount != nil {
		toSerialize["UserZoneCount"] = o.UserZoneCount
	}
	if o.UserZoneLimit != nil {
		toSerialize["UserZoneLimit"] = o.UserZoneLimit
	}
	if o.ZoneCount != nil {
		toSerialize["ZoneCount"] = o.ZoneCount
	}
	if o.ZoneLimit != nil {
		toSerialize["ZoneLimit"] = o.ZoneLimit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkFcZoneInfo struct {
	value *NetworkFcZoneInfo
	isSet bool
}

func (v NullableNetworkFcZoneInfo) Get() *NetworkFcZoneInfo {
	return v.value
}

func (v *NullableNetworkFcZoneInfo) Set(val *NetworkFcZoneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFcZoneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFcZoneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFcZoneInfo(val *NetworkFcZoneInfo) *NullableNetworkFcZoneInfo {
	return &NullableNetworkFcZoneInfo{value: val, isSet: true}
}

func (v NullableNetworkFcZoneInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFcZoneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
