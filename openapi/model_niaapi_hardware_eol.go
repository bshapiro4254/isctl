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
	"time"
)

// NiaapiHardwareEol This contains the end of life notice of hardware.
type NiaapiHardwareEol struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// String contains the PID of hardwares affected by this notice, seperated by comma.
	AffectedPids *string `json:"AffectedPids,omitempty" yaml:"AffectedPids,omitempty"`
	// When this notice is announced.
	AnnouncementDate *time.Time `json:"AnnouncementDate,omitempty" yaml:"AnnouncementDate,omitempty"`
	// Epoch time of Announcement Date.
	AnnouncementDateEpoch *int64 `json:"AnnouncementDateEpoch,omitempty" yaml:"AnnouncementDateEpoch,omitempty"`
	// The bulletinno of this hardware end of life notice.
	BulletinNo *string `json:"BulletinNo,omitempty" yaml:"BulletinNo,omitempty"`
	// The description of this hardware end of life notice.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// Date time of end of new services attachment  .
	EndofNewServiceAttachmentDate *time.Time `json:"EndofNewServiceAttachmentDate,omitempty" yaml:"EndofNewServiceAttachmentDate,omitempty"`
	// Epoch time of New service attachment Date .
	EndofNewServiceAttachmentDateEpoch *int64 `json:"EndofNewServiceAttachmentDateEpoch,omitempty" yaml:"EndofNewServiceAttachmentDateEpoch,omitempty"`
	// Date time of end of routinefailure analysis.
	EndofRoutineFailureAnalysisDate *time.Time `json:"EndofRoutineFailureAnalysisDate,omitempty" yaml:"EndofRoutineFailureAnalysisDate,omitempty"`
	// Epoch time of End of Routine Failure Analysis Date.
	EndofRoutineFailureAnalysisDateEpoch *int64 `json:"EndofRoutineFailureAnalysisDateEpoch,omitempty" yaml:"EndofRoutineFailureAnalysisDateEpoch,omitempty"`
	// When this hardware will end sale.
	EndofSaleDate *time.Time `json:"EndofSaleDate,omitempty" yaml:"EndofSaleDate,omitempty"`
	// Epoch time of End of Sale Date.
	EndofSaleDateEpoch *int64 `json:"EndofSaleDateEpoch,omitempty" yaml:"EndofSaleDateEpoch,omitempty"`
	// Date time of end of security support .
	EndofSecuritySupport *time.Time `json:"EndofSecuritySupport,omitempty" yaml:"EndofSecuritySupport,omitempty"`
	// Epoch time of End of Security Support Date .
	EndofSecuritySupportEpoch *int64 `json:"EndofSecuritySupportEpoch,omitempty" yaml:"EndofSecuritySupportEpoch,omitempty"`
	// Date time of end of service contract renew .
	EndofServiceContractRenewalDate *time.Time `json:"EndofServiceContractRenewalDate,omitempty" yaml:"EndofServiceContractRenewalDate,omitempty"`
	// Epoch time of End of Renewal service contract.
	EndofServiceContractRenewalDateEpoch *int64 `json:"EndofServiceContractRenewalDateEpoch,omitempty" yaml:"EndofServiceContractRenewalDateEpoch,omitempty"`
	// The date of end of maintainance.
	EndofSwMaintenanceDate *time.Time `json:"EndofSwMaintenanceDate,omitempty" yaml:"EndofSwMaintenanceDate,omitempty"`
	// Epoch time of End of maintenance Date.
	EndofSwMaintenanceDateEpoch *int64 `json:"EndofSwMaintenanceDateEpoch,omitempty" yaml:"EndofSwMaintenanceDateEpoch,omitempty"`
	// Hardware end of sale URL link to the notice webpage.
	HardwareEolUrl *string `json:"HardwareEolUrl,omitempty" yaml:"HardwareEolUrl,omitempty"`
	// The title of this hardware end of life notice.
	Headline *string `json:"Headline,omitempty" yaml:"Headline,omitempty"`
	// Date time of end of support .
	LastDateofSupport *time.Time `json:"LastDateofSupport,omitempty" yaml:"LastDateofSupport,omitempty"`
	// Epoch time of last date of support .
	LastDateofSupportEpoch *int64 `json:"LastDateofSupportEpoch,omitempty" yaml:"LastDateofSupportEpoch,omitempty"`
	// Date time of Lastship Date.
	LastShipDate *time.Time `json:"LastShipDate,omitempty" yaml:"LastShipDate,omitempty"`
	// Epoch time of last ship Date.
	LastShipDateEpoch *int64 `json:"LastShipDateEpoch,omitempty" yaml:"LastShipDateEpoch,omitempty"`
	// The URL of this migration notice.
	MigrationUrl *string `json:"MigrationUrl,omitempty" yaml:"MigrationUrl,omitempty"`
}

// NewNiaapiHardwareEol instantiates a new NiaapiHardwareEol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiHardwareEol() *NiaapiHardwareEol {
	this := NiaapiHardwareEol{}
	return &this
}

// NewNiaapiHardwareEolWithDefaults instantiates a new NiaapiHardwareEol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiHardwareEolWithDefaults() *NiaapiHardwareEol {
	this := NiaapiHardwareEol{}
	return &this
}

// GetAffectedPids returns the AffectedPids field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetAffectedPids() string {
	if o == nil || o.AffectedPids == nil {
		var ret string
		return ret
	}
	return *o.AffectedPids
}

// GetAffectedPidsOk returns a tuple with the AffectedPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetAffectedPidsOk() (*string, bool) {
	if o == nil || o.AffectedPids == nil {
		return nil, false
	}
	return o.AffectedPids, true
}

// HasAffectedPids returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasAffectedPids() bool {
	if o != nil && o.AffectedPids != nil {
		return true
	}

	return false
}

// SetAffectedPids gets a reference to the given string and assigns it to the AffectedPids field.
func (o *NiaapiHardwareEol) SetAffectedPids(v string) {
	o.AffectedPids = &v
}

// GetAnnouncementDate returns the AnnouncementDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetAnnouncementDate() time.Time {
	if o == nil || o.AnnouncementDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AnnouncementDate
}

// GetAnnouncementDateOk returns a tuple with the AnnouncementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetAnnouncementDateOk() (*time.Time, bool) {
	if o == nil || o.AnnouncementDate == nil {
		return nil, false
	}
	return o.AnnouncementDate, true
}

// HasAnnouncementDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasAnnouncementDate() bool {
	if o != nil && o.AnnouncementDate != nil {
		return true
	}

	return false
}

// SetAnnouncementDate gets a reference to the given time.Time and assigns it to the AnnouncementDate field.
func (o *NiaapiHardwareEol) SetAnnouncementDate(v time.Time) {
	o.AnnouncementDate = &v
}

// GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetAnnouncementDateEpoch() int64 {
	if o == nil || o.AnnouncementDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.AnnouncementDateEpoch
}

// GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetAnnouncementDateEpochOk() (*int64, bool) {
	if o == nil || o.AnnouncementDateEpoch == nil {
		return nil, false
	}
	return o.AnnouncementDateEpoch, true
}

// HasAnnouncementDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasAnnouncementDateEpoch() bool {
	if o != nil && o.AnnouncementDateEpoch != nil {
		return true
	}

	return false
}

// SetAnnouncementDateEpoch gets a reference to the given int64 and assigns it to the AnnouncementDateEpoch field.
func (o *NiaapiHardwareEol) SetAnnouncementDateEpoch(v int64) {
	o.AnnouncementDateEpoch = &v
}

// GetBulletinNo returns the BulletinNo field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetBulletinNo() string {
	if o == nil || o.BulletinNo == nil {
		var ret string
		return ret
	}
	return *o.BulletinNo
}

// GetBulletinNoOk returns a tuple with the BulletinNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetBulletinNoOk() (*string, bool) {
	if o == nil || o.BulletinNo == nil {
		return nil, false
	}
	return o.BulletinNo, true
}

// HasBulletinNo returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasBulletinNo() bool {
	if o != nil && o.BulletinNo != nil {
		return true
	}

	return false
}

// SetBulletinNo gets a reference to the given string and assigns it to the BulletinNo field.
func (o *NiaapiHardwareEol) SetBulletinNo(v string) {
	o.BulletinNo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiaapiHardwareEol) SetDescription(v string) {
	o.Description = &v
}

// GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDate() time.Time {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofNewServiceAttachmentDate
}

// GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool) {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDate, true
}

// HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofNewServiceAttachmentDate() bool {
	if o != nil && o.EndofNewServiceAttachmentDate != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDate gets a reference to the given time.Time and assigns it to the EndofNewServiceAttachmentDate field.
func (o *NiaapiHardwareEol) SetEndofNewServiceAttachmentDate(v time.Time) {
	o.EndofNewServiceAttachmentDate = &v
}

// GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateEpoch() int64 {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofNewServiceAttachmentDateEpoch
}

// GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDateEpoch, true
}

// HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofNewServiceAttachmentDateEpoch() bool {
	if o != nil && o.EndofNewServiceAttachmentDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDateEpoch gets a reference to the given int64 and assigns it to the EndofNewServiceAttachmentDateEpoch field.
func (o *NiaapiHardwareEol) SetEndofNewServiceAttachmentDateEpoch(v int64) {
	o.EndofNewServiceAttachmentDateEpoch = &v
}

// GetEndofRoutineFailureAnalysisDate returns the EndofRoutineFailureAnalysisDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDate() time.Time {
	if o == nil || o.EndofRoutineFailureAnalysisDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofRoutineFailureAnalysisDate
}

// GetEndofRoutineFailureAnalysisDateOk returns a tuple with the EndofRoutineFailureAnalysisDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateOk() (*time.Time, bool) {
	if o == nil || o.EndofRoutineFailureAnalysisDate == nil {
		return nil, false
	}
	return o.EndofRoutineFailureAnalysisDate, true
}

// HasEndofRoutineFailureAnalysisDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofRoutineFailureAnalysisDate() bool {
	if o != nil && o.EndofRoutineFailureAnalysisDate != nil {
		return true
	}

	return false
}

// SetEndofRoutineFailureAnalysisDate gets a reference to the given time.Time and assigns it to the EndofRoutineFailureAnalysisDate field.
func (o *NiaapiHardwareEol) SetEndofRoutineFailureAnalysisDate(v time.Time) {
	o.EndofRoutineFailureAnalysisDate = &v
}

// GetEndofRoutineFailureAnalysisDateEpoch returns the EndofRoutineFailureAnalysisDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateEpoch() int64 {
	if o == nil || o.EndofRoutineFailureAnalysisDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofRoutineFailureAnalysisDateEpoch
}

// GetEndofRoutineFailureAnalysisDateEpochOk returns a tuple with the EndofRoutineFailureAnalysisDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofRoutineFailureAnalysisDateEpoch == nil {
		return nil, false
	}
	return o.EndofRoutineFailureAnalysisDateEpoch, true
}

// HasEndofRoutineFailureAnalysisDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofRoutineFailureAnalysisDateEpoch() bool {
	if o != nil && o.EndofRoutineFailureAnalysisDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofRoutineFailureAnalysisDateEpoch gets a reference to the given int64 and assigns it to the EndofRoutineFailureAnalysisDateEpoch field.
func (o *NiaapiHardwareEol) SetEndofRoutineFailureAnalysisDateEpoch(v int64) {
	o.EndofRoutineFailureAnalysisDateEpoch = &v
}

// GetEndofSaleDate returns the EndofSaleDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSaleDate() time.Time {
	if o == nil || o.EndofSaleDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSaleDate
}

// GetEndofSaleDateOk returns a tuple with the EndofSaleDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSaleDateOk() (*time.Time, bool) {
	if o == nil || o.EndofSaleDate == nil {
		return nil, false
	}
	return o.EndofSaleDate, true
}

// HasEndofSaleDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSaleDate() bool {
	if o != nil && o.EndofSaleDate != nil {
		return true
	}

	return false
}

// SetEndofSaleDate gets a reference to the given time.Time and assigns it to the EndofSaleDate field.
func (o *NiaapiHardwareEol) SetEndofSaleDate(v time.Time) {
	o.EndofSaleDate = &v
}

// GetEndofSaleDateEpoch returns the EndofSaleDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSaleDateEpoch() int64 {
	if o == nil || o.EndofSaleDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSaleDateEpoch
}

// GetEndofSaleDateEpochOk returns a tuple with the EndofSaleDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSaleDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofSaleDateEpoch == nil {
		return nil, false
	}
	return o.EndofSaleDateEpoch, true
}

// HasEndofSaleDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSaleDateEpoch() bool {
	if o != nil && o.EndofSaleDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofSaleDateEpoch gets a reference to the given int64 and assigns it to the EndofSaleDateEpoch field.
func (o *NiaapiHardwareEol) SetEndofSaleDateEpoch(v int64) {
	o.EndofSaleDateEpoch = &v
}

// GetEndofSecuritySupport returns the EndofSecuritySupport field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSecuritySupport() time.Time {
	if o == nil || o.EndofSecuritySupport == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSecuritySupport
}

// GetEndofSecuritySupportOk returns a tuple with the EndofSecuritySupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSecuritySupportOk() (*time.Time, bool) {
	if o == nil || o.EndofSecuritySupport == nil {
		return nil, false
	}
	return o.EndofSecuritySupport, true
}

// HasEndofSecuritySupport returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSecuritySupport() bool {
	if o != nil && o.EndofSecuritySupport != nil {
		return true
	}

	return false
}

// SetEndofSecuritySupport gets a reference to the given time.Time and assigns it to the EndofSecuritySupport field.
func (o *NiaapiHardwareEol) SetEndofSecuritySupport(v time.Time) {
	o.EndofSecuritySupport = &v
}

// GetEndofSecuritySupportEpoch returns the EndofSecuritySupportEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSecuritySupportEpoch() int64 {
	if o == nil || o.EndofSecuritySupportEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSecuritySupportEpoch
}

// GetEndofSecuritySupportEpochOk returns a tuple with the EndofSecuritySupportEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSecuritySupportEpochOk() (*int64, bool) {
	if o == nil || o.EndofSecuritySupportEpoch == nil {
		return nil, false
	}
	return o.EndofSecuritySupportEpoch, true
}

// HasEndofSecuritySupportEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSecuritySupportEpoch() bool {
	if o != nil && o.EndofSecuritySupportEpoch != nil {
		return true
	}

	return false
}

// SetEndofSecuritySupportEpoch gets a reference to the given int64 and assigns it to the EndofSecuritySupportEpoch field.
func (o *NiaapiHardwareEol) SetEndofSecuritySupportEpoch(v int64) {
	o.EndofSecuritySupportEpoch = &v
}

// GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDate() time.Time {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofServiceContractRenewalDate
}

// GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateOk() (*time.Time, bool) {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDate, true
}

// HasEndofServiceContractRenewalDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofServiceContractRenewalDate() bool {
	if o != nil && o.EndofServiceContractRenewalDate != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDate gets a reference to the given time.Time and assigns it to the EndofServiceContractRenewalDate field.
func (o *NiaapiHardwareEol) SetEndofServiceContractRenewalDate(v time.Time) {
	o.EndofServiceContractRenewalDate = &v
}

// GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateEpoch() int64 {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofServiceContractRenewalDateEpoch
}

// GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDateEpoch, true
}

// HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofServiceContractRenewalDateEpoch() bool {
	if o != nil && o.EndofServiceContractRenewalDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDateEpoch gets a reference to the given int64 and assigns it to the EndofServiceContractRenewalDateEpoch field.
func (o *NiaapiHardwareEol) SetEndofServiceContractRenewalDateEpoch(v int64) {
	o.EndofServiceContractRenewalDateEpoch = &v
}

// GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDate() time.Time {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSwMaintenanceDate
}

// GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateOk() (*time.Time, bool) {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDate, true
}

// HasEndofSwMaintenanceDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSwMaintenanceDate() bool {
	if o != nil && o.EndofSwMaintenanceDate != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDate gets a reference to the given time.Time and assigns it to the EndofSwMaintenanceDate field.
func (o *NiaapiHardwareEol) SetEndofSwMaintenanceDate(v time.Time) {
	o.EndofSwMaintenanceDate = &v
}

// GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateEpoch() int64 {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSwMaintenanceDateEpoch
}

// GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDateEpoch, true
}

// HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasEndofSwMaintenanceDateEpoch() bool {
	if o != nil && o.EndofSwMaintenanceDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDateEpoch gets a reference to the given int64 and assigns it to the EndofSwMaintenanceDateEpoch field.
func (o *NiaapiHardwareEol) SetEndofSwMaintenanceDateEpoch(v int64) {
	o.EndofSwMaintenanceDateEpoch = &v
}

// GetHardwareEolUrl returns the HardwareEolUrl field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetHardwareEolUrl() string {
	if o == nil || o.HardwareEolUrl == nil {
		var ret string
		return ret
	}
	return *o.HardwareEolUrl
}

// GetHardwareEolUrlOk returns a tuple with the HardwareEolUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetHardwareEolUrlOk() (*string, bool) {
	if o == nil || o.HardwareEolUrl == nil {
		return nil, false
	}
	return o.HardwareEolUrl, true
}

// HasHardwareEolUrl returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasHardwareEolUrl() bool {
	if o != nil && o.HardwareEolUrl != nil {
		return true
	}

	return false
}

// SetHardwareEolUrl gets a reference to the given string and assigns it to the HardwareEolUrl field.
func (o *NiaapiHardwareEol) SetHardwareEolUrl(v string) {
	o.HardwareEolUrl = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *NiaapiHardwareEol) SetHeadline(v string) {
	o.Headline = &v
}

// GetLastDateofSupport returns the LastDateofSupport field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetLastDateofSupport() time.Time {
	if o == nil || o.LastDateofSupport == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDateofSupport
}

// GetLastDateofSupportOk returns a tuple with the LastDateofSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetLastDateofSupportOk() (*time.Time, bool) {
	if o == nil || o.LastDateofSupport == nil {
		return nil, false
	}
	return o.LastDateofSupport, true
}

// HasLastDateofSupport returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasLastDateofSupport() bool {
	if o != nil && o.LastDateofSupport != nil {
		return true
	}

	return false
}

// SetLastDateofSupport gets a reference to the given time.Time and assigns it to the LastDateofSupport field.
func (o *NiaapiHardwareEol) SetLastDateofSupport(v time.Time) {
	o.LastDateofSupport = &v
}

// GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetLastDateofSupportEpoch() int64 {
	if o == nil || o.LastDateofSupportEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastDateofSupportEpoch
}

// GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetLastDateofSupportEpochOk() (*int64, bool) {
	if o == nil || o.LastDateofSupportEpoch == nil {
		return nil, false
	}
	return o.LastDateofSupportEpoch, true
}

// HasLastDateofSupportEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasLastDateofSupportEpoch() bool {
	if o != nil && o.LastDateofSupportEpoch != nil {
		return true
	}

	return false
}

// SetLastDateofSupportEpoch gets a reference to the given int64 and assigns it to the LastDateofSupportEpoch field.
func (o *NiaapiHardwareEol) SetLastDateofSupportEpoch(v int64) {
	o.LastDateofSupportEpoch = &v
}

// GetLastShipDate returns the LastShipDate field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetLastShipDate() time.Time {
	if o == nil || o.LastShipDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastShipDate
}

// GetLastShipDateOk returns a tuple with the LastShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetLastShipDateOk() (*time.Time, bool) {
	if o == nil || o.LastShipDate == nil {
		return nil, false
	}
	return o.LastShipDate, true
}

// HasLastShipDate returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasLastShipDate() bool {
	if o != nil && o.LastShipDate != nil {
		return true
	}

	return false
}

// SetLastShipDate gets a reference to the given time.Time and assigns it to the LastShipDate field.
func (o *NiaapiHardwareEol) SetLastShipDate(v time.Time) {
	o.LastShipDate = &v
}

// GetLastShipDateEpoch returns the LastShipDateEpoch field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetLastShipDateEpoch() int64 {
	if o == nil || o.LastShipDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastShipDateEpoch
}

// GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetLastShipDateEpochOk() (*int64, bool) {
	if o == nil || o.LastShipDateEpoch == nil {
		return nil, false
	}
	return o.LastShipDateEpoch, true
}

// HasLastShipDateEpoch returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasLastShipDateEpoch() bool {
	if o != nil && o.LastShipDateEpoch != nil {
		return true
	}

	return false
}

// SetLastShipDateEpoch gets a reference to the given int64 and assigns it to the LastShipDateEpoch field.
func (o *NiaapiHardwareEol) SetLastShipDateEpoch(v int64) {
	o.LastShipDateEpoch = &v
}

// GetMigrationUrl returns the MigrationUrl field value if set, zero value otherwise.
func (o *NiaapiHardwareEol) GetMigrationUrl() string {
	if o == nil || o.MigrationUrl == nil {
		var ret string
		return ret
	}
	return *o.MigrationUrl
}

// GetMigrationUrlOk returns a tuple with the MigrationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiHardwareEol) GetMigrationUrlOk() (*string, bool) {
	if o == nil || o.MigrationUrl == nil {
		return nil, false
	}
	return o.MigrationUrl, true
}

// HasMigrationUrl returns a boolean if a field has been set.
func (o *NiaapiHardwareEol) HasMigrationUrl() bool {
	if o != nil && o.MigrationUrl != nil {
		return true
	}

	return false
}

// SetMigrationUrl gets a reference to the given string and assigns it to the MigrationUrl field.
func (o *NiaapiHardwareEol) SetMigrationUrl(v string) {
	o.MigrationUrl = &v
}

func (o NiaapiHardwareEol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.AffectedPids != nil {
		toSerialize["AffectedPids"] = o.AffectedPids
	}
	if o.AnnouncementDate != nil {
		toSerialize["AnnouncementDate"] = o.AnnouncementDate
	}
	if o.AnnouncementDateEpoch != nil {
		toSerialize["AnnouncementDateEpoch"] = o.AnnouncementDateEpoch
	}
	if o.BulletinNo != nil {
		toSerialize["BulletinNo"] = o.BulletinNo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndofNewServiceAttachmentDate != nil {
		toSerialize["EndofNewServiceAttachmentDate"] = o.EndofNewServiceAttachmentDate
	}
	if o.EndofNewServiceAttachmentDateEpoch != nil {
		toSerialize["EndofNewServiceAttachmentDateEpoch"] = o.EndofNewServiceAttachmentDateEpoch
	}
	if o.EndofRoutineFailureAnalysisDate != nil {
		toSerialize["EndofRoutineFailureAnalysisDate"] = o.EndofRoutineFailureAnalysisDate
	}
	if o.EndofRoutineFailureAnalysisDateEpoch != nil {
		toSerialize["EndofRoutineFailureAnalysisDateEpoch"] = o.EndofRoutineFailureAnalysisDateEpoch
	}
	if o.EndofSaleDate != nil {
		toSerialize["EndofSaleDate"] = o.EndofSaleDate
	}
	if o.EndofSaleDateEpoch != nil {
		toSerialize["EndofSaleDateEpoch"] = o.EndofSaleDateEpoch
	}
	if o.EndofSecuritySupport != nil {
		toSerialize["EndofSecuritySupport"] = o.EndofSecuritySupport
	}
	if o.EndofSecuritySupportEpoch != nil {
		toSerialize["EndofSecuritySupportEpoch"] = o.EndofSecuritySupportEpoch
	}
	if o.EndofServiceContractRenewalDate != nil {
		toSerialize["EndofServiceContractRenewalDate"] = o.EndofServiceContractRenewalDate
	}
	if o.EndofServiceContractRenewalDateEpoch != nil {
		toSerialize["EndofServiceContractRenewalDateEpoch"] = o.EndofServiceContractRenewalDateEpoch
	}
	if o.EndofSwMaintenanceDate != nil {
		toSerialize["EndofSwMaintenanceDate"] = o.EndofSwMaintenanceDate
	}
	if o.EndofSwMaintenanceDateEpoch != nil {
		toSerialize["EndofSwMaintenanceDateEpoch"] = o.EndofSwMaintenanceDateEpoch
	}
	if o.HardwareEolUrl != nil {
		toSerialize["HardwareEolUrl"] = o.HardwareEolUrl
	}
	if o.Headline != nil {
		toSerialize["Headline"] = o.Headline
	}
	if o.LastDateofSupport != nil {
		toSerialize["LastDateofSupport"] = o.LastDateofSupport
	}
	if o.LastDateofSupportEpoch != nil {
		toSerialize["LastDateofSupportEpoch"] = o.LastDateofSupportEpoch
	}
	if o.LastShipDate != nil {
		toSerialize["LastShipDate"] = o.LastShipDate
	}
	if o.LastShipDateEpoch != nil {
		toSerialize["LastShipDateEpoch"] = o.LastShipDateEpoch
	}
	if o.MigrationUrl != nil {
		toSerialize["MigrationUrl"] = o.MigrationUrl
	}
	return json.Marshal(toSerialize)
}

type NullableNiaapiHardwareEol struct {
	value *NiaapiHardwareEol
	isSet bool
}

func (v NullableNiaapiHardwareEol) Get() *NiaapiHardwareEol {
	return v.value
}

func (v *NullableNiaapiHardwareEol) Set(val *NiaapiHardwareEol) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiHardwareEol) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiHardwareEol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiHardwareEol(val *NiaapiHardwareEol) *NullableNiaapiHardwareEol {
	return &NullableNiaapiHardwareEol{value: val, isSet: true}
}

func (v NullableNiaapiHardwareEol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiHardwareEol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
