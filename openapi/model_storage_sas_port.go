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

// StorageSasPort Sas Port details of the SAS endpoint.
type StorageSasPort struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// The SAS Address assigned to storage port.
	Address *string `json:"Address,omitempty" yaml:"Address,omitempty"`
	// The unique disk identifier.
	DiskId *int64 `json:"DiskId,omitempty" yaml:"DiskId,omitempty"`
	// The end-point Id assigned to storage port.
	EndPointId *int64 `json:"EndPointId,omitempty" yaml:"EndPointId,omitempty"`
	// The description for the link.
	LinkDescription *string `json:"LinkDescription,omitempty" yaml:"LinkDescription,omitempty"`
	// The link speed negotiated for communication.
	LinkSpeed           *string                              `json:"LinkSpeed,omitempty" yaml:"LinkSpeed,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty" yaml:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	StoragePhysicalDisk *StoragePhysicalDiskRelationship     `json:"StoragePhysicalDisk,omitempty" yaml:"StoragePhysicalDisk,omitempty"`
}

// NewStorageSasPort instantiates a new StorageSasPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSasPort() *StorageSasPort {
	this := StorageSasPort{}
	return &this
}

// NewStorageSasPortWithDefaults instantiates a new StorageSasPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSasPortWithDefaults() *StorageSasPort {
	this := StorageSasPort{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *StorageSasPort) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *StorageSasPort) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *StorageSasPort) SetAddress(v string) {
	o.Address = &v
}

// GetDiskId returns the DiskId field value if set, zero value otherwise.
func (o *StorageSasPort) GetDiskId() int64 {
	if o == nil || o.DiskId == nil {
		var ret int64
		return ret
	}
	return *o.DiskId
}

// GetDiskIdOk returns a tuple with the DiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetDiskIdOk() (*int64, bool) {
	if o == nil || o.DiskId == nil {
		return nil, false
	}
	return o.DiskId, true
}

// HasDiskId returns a boolean if a field has been set.
func (o *StorageSasPort) HasDiskId() bool {
	if o != nil && o.DiskId != nil {
		return true
	}

	return false
}

// SetDiskId gets a reference to the given int64 and assigns it to the DiskId field.
func (o *StorageSasPort) SetDiskId(v int64) {
	o.DiskId = &v
}

// GetEndPointId returns the EndPointId field value if set, zero value otherwise.
func (o *StorageSasPort) GetEndPointId() int64 {
	if o == nil || o.EndPointId == nil {
		var ret int64
		return ret
	}
	return *o.EndPointId
}

// GetEndPointIdOk returns a tuple with the EndPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetEndPointIdOk() (*int64, bool) {
	if o == nil || o.EndPointId == nil {
		return nil, false
	}
	return o.EndPointId, true
}

// HasEndPointId returns a boolean if a field has been set.
func (o *StorageSasPort) HasEndPointId() bool {
	if o != nil && o.EndPointId != nil {
		return true
	}

	return false
}

// SetEndPointId gets a reference to the given int64 and assigns it to the EndPointId field.
func (o *StorageSasPort) SetEndPointId(v int64) {
	o.EndPointId = &v
}

// GetLinkDescription returns the LinkDescription field value if set, zero value otherwise.
func (o *StorageSasPort) GetLinkDescription() string {
	if o == nil || o.LinkDescription == nil {
		var ret string
		return ret
	}
	return *o.LinkDescription
}

// GetLinkDescriptionOk returns a tuple with the LinkDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetLinkDescriptionOk() (*string, bool) {
	if o == nil || o.LinkDescription == nil {
		return nil, false
	}
	return o.LinkDescription, true
}

// HasLinkDescription returns a boolean if a field has been set.
func (o *StorageSasPort) HasLinkDescription() bool {
	if o != nil && o.LinkDescription != nil {
		return true
	}

	return false
}

// SetLinkDescription gets a reference to the given string and assigns it to the LinkDescription field.
func (o *StorageSasPort) SetLinkDescription(v string) {
	o.LinkDescription = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *StorageSasPort) GetLinkSpeed() string {
	if o == nil || o.LinkSpeed == nil {
		var ret string
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetLinkSpeedOk() (*string, bool) {
	if o == nil || o.LinkSpeed == nil {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *StorageSasPort) HasLinkSpeed() bool {
	if o != nil && o.LinkSpeed != nil {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given string and assigns it to the LinkSpeed field.
func (o *StorageSasPort) SetLinkSpeed(v string) {
	o.LinkSpeed = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageSasPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageSasPort) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageSasPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageSasPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageSasPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageSasPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStoragePhysicalDisk returns the StoragePhysicalDisk field value if set, zero value otherwise.
func (o *StorageSasPort) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.StoragePhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.StoragePhysicalDisk
}

// GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasPort) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.StoragePhysicalDisk == nil {
		return nil, false
	}
	return o.StoragePhysicalDisk, true
}

// HasStoragePhysicalDisk returns a boolean if a field has been set.
func (o *StorageSasPort) HasStoragePhysicalDisk() bool {
	if o != nil && o.StoragePhysicalDisk != nil {
		return true
	}

	return false
}

// SetStoragePhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the StoragePhysicalDisk field.
func (o *StorageSasPort) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.StoragePhysicalDisk = &v
}

func (o StorageSasPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.DiskId != nil {
		toSerialize["DiskId"] = o.DiskId
	}
	if o.EndPointId != nil {
		toSerialize["EndPointId"] = o.EndPointId
	}
	if o.LinkDescription != nil {
		toSerialize["LinkDescription"] = o.LinkDescription
	}
	if o.LinkSpeed != nil {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StoragePhysicalDisk != nil {
		toSerialize["StoragePhysicalDisk"] = o.StoragePhysicalDisk
	}
	return json.Marshal(toSerialize)
}

type NullableStorageSasPort struct {
	value *StorageSasPort
	isSet bool
}

func (v NullableStorageSasPort) Get() *StorageSasPort {
	return v.value
}

func (v *NullableStorageSasPort) Set(val *StorageSasPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSasPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSasPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSasPort(val *StorageSasPort) *NullableStorageSasPort {
	return &NullableStorageSasPort{value: val, isSet: true}
}

func (v NullableStorageSasPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSasPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
