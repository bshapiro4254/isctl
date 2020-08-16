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

// ComputeServerSetting Models the configurable properties of a server in Intersight.
type ComputeServerSetting struct {
	InventoryBase `yaml:"InventoryBase,inline"`
	// User configured state of the locator LED for the server. * `None` - No operation property for locator led. * `On` - The Locator Led is turned on. * `Off` - The Locator Led is turned off.
	AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty" yaml:"AdminLocatorLedState,omitempty"`
	// User configured power state of the server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
	AdminPowerState *string `json:"AdminPowerState,omitempty" yaml:"AdminPowerState,omitempty"`
	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty" yaml:"ConfigState,omitempty"`
	// The name of the device chosen by user for configuring One-Time Boot device.
	OneTimeBootDevice         *string                              `json:"OneTimeBootDevice,omitempty" yaml:"OneTimeBootDevice,omitempty"`
	PersistentMemoryOperation *ComputePersistentMemoryOperation    `json:"PersistentMemoryOperation,omitempty" yaml:"PersistentMemoryOperation,omitempty"`
	ServerConfig              *ComputeServerConfig                 `json:"ServerConfig,omitempty" yaml:"ServerConfig,omitempty"`
	LocatorLed                *EquipmentLocatorLedRelationship     `json:"LocatorLed,omitempty" yaml:"LocatorLed,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	RunningWorkflow           *WorkflowWorkflowInfoRelationship    `json:"RunningWorkflow,omitempty" yaml:"RunningWorkflow,omitempty"`
	Server                    *ComputePhysicalRelationship         `json:"Server,omitempty" yaml:"Server,omitempty"`
}

// NewComputeServerSetting instantiates a new ComputeServerSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerSetting() *ComputeServerSetting {
	this := ComputeServerSetting{}
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// NewComputeServerSettingWithDefaults instantiates a new ComputeServerSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerSettingWithDefaults() *ComputeServerSetting {
	this := ComputeServerSetting{}
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// GetAdminLocatorLedState returns the AdminLocatorLedState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetAdminLocatorLedState() string {
	if o == nil || o.AdminLocatorLedState == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedState
}

// GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetAdminLocatorLedStateOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedState == nil {
		return nil, false
	}
	return o.AdminLocatorLedState, true
}

// HasAdminLocatorLedState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasAdminLocatorLedState() bool {
	if o != nil && o.AdminLocatorLedState != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedState gets a reference to the given string and assigns it to the AdminLocatorLedState field.
func (o *ComputeServerSetting) SetAdminLocatorLedState(v string) {
	o.AdminLocatorLedState = &v
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *ComputeServerSetting) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *ComputeServerSetting) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetOneTimeBootDevice returns the OneTimeBootDevice field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetOneTimeBootDevice() string {
	if o == nil || o.OneTimeBootDevice == nil {
		var ret string
		return ret
	}
	return *o.OneTimeBootDevice
}

// GetOneTimeBootDeviceOk returns a tuple with the OneTimeBootDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetOneTimeBootDeviceOk() (*string, bool) {
	if o == nil || o.OneTimeBootDevice == nil {
		return nil, false
	}
	return o.OneTimeBootDevice, true
}

// HasOneTimeBootDevice returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasOneTimeBootDevice() bool {
	if o != nil && o.OneTimeBootDevice != nil {
		return true
	}

	return false
}

// SetOneTimeBootDevice gets a reference to the given string and assigns it to the OneTimeBootDevice field.
func (o *ComputeServerSetting) SetOneTimeBootDevice(v string) {
	o.OneTimeBootDevice = &v
}

// GetPersistentMemoryOperation returns the PersistentMemoryOperation field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetPersistentMemoryOperation() ComputePersistentMemoryOperation {
	if o == nil || o.PersistentMemoryOperation == nil {
		var ret ComputePersistentMemoryOperation
		return ret
	}
	return *o.PersistentMemoryOperation
}

// GetPersistentMemoryOperationOk returns a tuple with the PersistentMemoryOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetPersistentMemoryOperationOk() (*ComputePersistentMemoryOperation, bool) {
	if o == nil || o.PersistentMemoryOperation == nil {
		return nil, false
	}
	return o.PersistentMemoryOperation, true
}

// HasPersistentMemoryOperation returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasPersistentMemoryOperation() bool {
	if o != nil && o.PersistentMemoryOperation != nil {
		return true
	}

	return false
}

// SetPersistentMemoryOperation gets a reference to the given ComputePersistentMemoryOperation and assigns it to the PersistentMemoryOperation field.
func (o *ComputeServerSetting) SetPersistentMemoryOperation(v ComputePersistentMemoryOperation) {
	o.PersistentMemoryOperation = &v
}

// GetServerConfig returns the ServerConfig field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetServerConfig() ComputeServerConfig {
	if o == nil || o.ServerConfig == nil {
		var ret ComputeServerConfig
		return ret
	}
	return *o.ServerConfig
}

// GetServerConfigOk returns a tuple with the ServerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetServerConfigOk() (*ComputeServerConfig, bool) {
	if o == nil || o.ServerConfig == nil {
		return nil, false
	}
	return o.ServerConfig, true
}

// HasServerConfig returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasServerConfig() bool {
	if o != nil && o.ServerConfig != nil {
		return true
	}

	return false
}

// SetServerConfig gets a reference to the given ComputeServerConfig and assigns it to the ServerConfig field.
func (o *ComputeServerSetting) SetServerConfig(v ComputeServerConfig) {
	o.ServerConfig = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeServerSetting) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeServerSetting) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningWorkflow returns the RunningWorkflow field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetRunningWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.RunningWorkflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.RunningWorkflow
}

// GetRunningWorkflowOk returns a tuple with the RunningWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetRunningWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflow == nil {
		return nil, false
	}
	return o.RunningWorkflow, true
}

// HasRunningWorkflow returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasRunningWorkflow() bool {
	if o != nil && o.RunningWorkflow != nil {
		return true
	}

	return false
}

// SetRunningWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflow field.
func (o *ComputeServerSetting) SetRunningWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflow = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ComputeServerSetting) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSetting) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ComputeServerSetting) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *ComputeServerSetting) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o ComputeServerSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if o.AdminLocatorLedState != nil {
		toSerialize["AdminLocatorLedState"] = o.AdminLocatorLedState
	}
	if o.AdminPowerState != nil {
		toSerialize["AdminPowerState"] = o.AdminPowerState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.OneTimeBootDevice != nil {
		toSerialize["OneTimeBootDevice"] = o.OneTimeBootDevice
	}
	if o.PersistentMemoryOperation != nil {
		toSerialize["PersistentMemoryOperation"] = o.PersistentMemoryOperation
	}
	if o.ServerConfig != nil {
		toSerialize["ServerConfig"] = o.ServerConfig
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningWorkflow != nil {
		toSerialize["RunningWorkflow"] = o.RunningWorkflow
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableComputeServerSetting struct {
	value *ComputeServerSetting
	isSet bool
}

func (v NullableComputeServerSetting) Get() *ComputeServerSetting {
	return v.value
}

func (v *NullableComputeServerSetting) Set(val *ComputeServerSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerSetting(val *ComputeServerSetting) *NullableComputeServerSetting {
	return &NullableComputeServerSetting{value: val, isSet: true}
}

func (v NullableComputeServerSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
