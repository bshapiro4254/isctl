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

// TechsupportmanagementTechSupportBundleAllOf Definition of the list of properties defined in 'techsupportmanagement.TechSupportBundle', excluding properties defined in parent classes.
type TechsupportmanagementTechSupportBundleAllOf struct {
	// The device identifier used to uniquely identify an individual device.
	DeviceIdentifier *string `json:"DeviceIdentifier,omitempty" yaml:"DeviceIdentifier,omitempty"`
	// The device type obtained from the inventory.
	DeviceType *string `json:"DeviceType,omitempty" yaml:"DeviceType,omitempty"`
	// Product identification of the device.
	Pid           *string                     `json:"Pid,omitempty" yaml:"Pid,omitempty"`
	PlatformParam *ConnectorPlatformParamBase `json:"PlatformParam,omitempty" yaml:"PlatformParam,omitempty"`
	// The platform type of the device. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - Intersight on-premise appliance. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `NetAppOntap` - A NetApp ONTAP storage system. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft HyperV system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A Dynatrace controller that monitors applications. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `MicrosoftAzure` - A Microsoft Azure target. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `IMCBlade` - An Intersight managed UCS Blade Server.
	PlatformType *string `json:"PlatformType,omitempty" yaml:"PlatformType,omitempty"`
	// Serial number of the device.
	Serial             *string                                             `json:"Serial,omitempty" yaml:"Serial,omitempty"`
	DeviceRegistration *AssetDeviceRegistrationRelationship                `json:"DeviceRegistration,omitempty" yaml:"DeviceRegistration,omitempty"`
	TargetResource     *MoBaseMoRelationship                               `json:"TargetResource,omitempty" yaml:"TargetResource,omitempty"`
	TechSupportStatus  *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty" yaml:"TechSupportStatus,omitempty"`
}

// NewTechsupportmanagementTechSupportBundleAllOf instantiates a new TechsupportmanagementTechSupportBundleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementTechSupportBundleAllOf() *TechsupportmanagementTechSupportBundleAllOf {
	this := TechsupportmanagementTechSupportBundleAllOf{}
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewTechsupportmanagementTechSupportBundleAllOfWithDefaults instantiates a new TechsupportmanagementTechSupportBundleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementTechSupportBundleAllOfWithDefaults() *TechsupportmanagementTechSupportBundleAllOf {
	this := TechsupportmanagementTechSupportBundleAllOf{}
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetDeviceIdentifier returns the DeviceIdentifier field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifier() string {
	if o == nil || o.DeviceIdentifier == nil {
		var ret string
		return ret
	}
	return *o.DeviceIdentifier
}

// GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifierOk() (*string, bool) {
	if o == nil || o.DeviceIdentifier == nil {
		return nil, false
	}
	return o.DeviceIdentifier, true
}

// HasDeviceIdentifier returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceIdentifier() bool {
	if o != nil && o.DeviceIdentifier != nil {
		return true
	}

	return false
}

// SetDeviceIdentifier gets a reference to the given string and assigns it to the DeviceIdentifier field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceIdentifier(v string) {
	o.DeviceIdentifier = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetPlatformParam returns the PlatformParam field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParam() ConnectorPlatformParamBase {
	if o == nil || o.PlatformParam == nil {
		var ret ConnectorPlatformParamBase
		return ret
	}
	return *o.PlatformParam
}

// GetPlatformParamOk returns a tuple with the PlatformParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool) {
	if o == nil || o.PlatformParam == nil {
		return nil, false
	}
	return o.PlatformParam, true
}

// HasPlatformParam returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformParam() bool {
	if o != nil && o.PlatformParam != nil {
		return true
	}

	return false
}

// SetPlatformParam gets a reference to the given ConnectorPlatformParamBase and assigns it to the PlatformParam field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformParam(v ConnectorPlatformParamBase) {
	o.PlatformParam = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetTargetResource returns the TargetResource field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResource() MoBaseMoRelationship {
	if o == nil || o.TargetResource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetResource
}

// GetTargetResourceOk returns a tuple with the TargetResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetResource == nil {
		return nil, false
	}
	return o.TargetResource, true
}

// HasTargetResource returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasTargetResource() bool {
	if o != nil && o.TargetResource != nil {
		return true
	}

	return false
}

// SetTargetResource gets a reference to the given MoBaseMoRelationship and assigns it to the TargetResource field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetTargetResource(v MoBaseMoRelationship) {
	o.TargetResource = &v
}

// GetTechSupportStatus returns the TechSupportStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship {
	if o == nil || o.TechSupportStatus == nil {
		var ret TechsupportmanagementTechSupportStatusRelationship
		return ret
	}
	return *o.TechSupportStatus
}

// GetTechSupportStatusOk returns a tuple with the TechSupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool) {
	if o == nil || o.TechSupportStatus == nil {
		return nil, false
	}
	return o.TechSupportStatus, true
}

// HasTechSupportStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementTechSupportBundleAllOf) HasTechSupportStatus() bool {
	if o != nil && o.TechSupportStatus != nil {
		return true
	}

	return false
}

// SetTechSupportStatus gets a reference to the given TechsupportmanagementTechSupportStatusRelationship and assigns it to the TechSupportStatus field.
func (o *TechsupportmanagementTechSupportBundleAllOf) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship) {
	o.TechSupportStatus = &v
}

func (o TechsupportmanagementTechSupportBundleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceIdentifier != nil {
		toSerialize["DeviceIdentifier"] = o.DeviceIdentifier
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.PlatformParam != nil {
		toSerialize["PlatformParam"] = o.PlatformParam
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.TargetResource != nil {
		toSerialize["TargetResource"] = o.TargetResource
	}
	if o.TechSupportStatus != nil {
		toSerialize["TechSupportStatus"] = o.TechSupportStatus
	}
	return json.Marshal(toSerialize)
}

type NullableTechsupportmanagementTechSupportBundleAllOf struct {
	value *TechsupportmanagementTechSupportBundleAllOf
	isSet bool
}

func (v NullableTechsupportmanagementTechSupportBundleAllOf) Get() *TechsupportmanagementTechSupportBundleAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementTechSupportBundleAllOf) Set(val *TechsupportmanagementTechSupportBundleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementTechSupportBundleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementTechSupportBundleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementTechSupportBundleAllOf(val *TechsupportmanagementTechSupportBundleAllOf) *NullableTechsupportmanagementTechSupportBundleAllOf {
	return &NullableTechsupportmanagementTechSupportBundleAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementTechSupportBundleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementTechSupportBundleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
