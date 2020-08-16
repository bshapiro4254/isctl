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

// ApplianceUpgradeAllOf Definition of the list of properties defined in 'appliance.Upgrade', excluding properties defined in parent classes.
type ApplianceUpgradeAllOf struct {
	// Indicates if the software upgrade is active or not.
	Active *bool `json:"Active,omitempty" yaml:"Active,omitempty"`
	// Indicates that the request was automatically created by the system.
	AutoCreated     *bool                 `json:"AutoCreated,omitempty" yaml:"AutoCreated,omitempty"`
	CompletedPhases *[]OnpremUpgradePhase `json:"CompletedPhases,omitempty" yaml:"CompletedPhases,omitempty"`
	CurrentPhase    *OnpremUpgradePhase   `json:"CurrentPhase,omitempty" yaml:"CurrentPhase,omitempty"`
	// Description of the software upgrade.
	Description *string `json:"Description,omitempty" yaml:"Description,omitempty"`
	// Elapsed time in seconds during the software upgrade.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" yaml:"ElapsedTime,omitempty"`
	// End date of the software upgrade.
	EndTime *time.Time `json:"EndTime,omitempty" yaml:"EndTime,omitempty"`
	// Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried.
	ErrorCode *int64 `json:"ErrorCode,omitempty" yaml:"ErrorCode,omitempty"`
	// Software upgrade manifest's fingerprint.
	Fingerprint *string `json:"Fingerprint,omitempty" yaml:"Fingerprint,omitempty"`
	// Track if software upgrade is upgrading or rolling back.
	IsRollingBack *bool     `json:"IsRollingBack,omitempty" yaml:"IsRollingBack,omitempty"`
	Messages      *[]string `json:"Messages,omitempty" yaml:"Messages,omitempty"`
	// Track if rollback is needed.
	RollbackNeeded *bool                 `json:"RollbackNeeded,omitempty" yaml:"RollbackNeeded,omitempty"`
	RollbackPhases *[]OnpremUpgradePhase `json:"RollbackPhases,omitempty" yaml:"RollbackPhases,omitempty"`
	// Status of the Intersight Appliance's software rollback status.
	RollbackStatus *string   `json:"RollbackStatus,omitempty" yaml:"RollbackStatus,omitempty"`
	Services       *[]string `json:"Services,omitempty" yaml:"Services,omitempty"`
	// Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.
	StartTime *time.Time `json:"StartTime,omitempty" yaml:"StartTime,omitempty"`
	// Status of the Intersight Appliance's software upgrade.
	Status *string `json:"Status,omitempty" yaml:"Status,omitempty"`
	// TotalPhase represents the total number of the upgradePhases for one upgrade.
	TotalPhases *int64    `json:"TotalPhases,omitempty" yaml:"TotalPhases,omitempty"`
	UiPackages  *[]string `json:"UiPackages,omitempty" yaml:"UiPackages,omitempty"`
	// Software upgrade manifest's version.
	Version     *string                           `json:"Version,omitempty" yaml:"Version,omitempty"`
	Account     *IamAccountRelationship           `json:"Account,omitempty" yaml:"Account,omitempty"`
	ImageBundle *ApplianceImageBundleRelationship `json:"ImageBundle,omitempty" yaml:"ImageBundle,omitempty"`
}

// NewApplianceUpgradeAllOf instantiates a new ApplianceUpgradeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceUpgradeAllOf() *ApplianceUpgradeAllOf {
	this := ApplianceUpgradeAllOf{}
	return &this
}

// NewApplianceUpgradeAllOfWithDefaults instantiates a new ApplianceUpgradeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceUpgradeAllOfWithDefaults() *ApplianceUpgradeAllOf {
	this := ApplianceUpgradeAllOf{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApplianceUpgradeAllOf) SetActive(v bool) {
	o.Active = &v
}

// GetAutoCreated returns the AutoCreated field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetAutoCreated() bool {
	if o == nil || o.AutoCreated == nil {
		var ret bool
		return ret
	}
	return *o.AutoCreated
}

// GetAutoCreatedOk returns a tuple with the AutoCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetAutoCreatedOk() (*bool, bool) {
	if o == nil || o.AutoCreated == nil {
		return nil, false
	}
	return o.AutoCreated, true
}

// HasAutoCreated returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasAutoCreated() bool {
	if o != nil && o.AutoCreated != nil {
		return true
	}

	return false
}

// SetAutoCreated gets a reference to the given bool and assigns it to the AutoCreated field.
func (o *ApplianceUpgradeAllOf) SetAutoCreated(v bool) {
	o.AutoCreated = &v
}

// GetCompletedPhases returns the CompletedPhases field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetCompletedPhases() []OnpremUpgradePhase {
	if o == nil || o.CompletedPhases == nil {
		var ret []OnpremUpgradePhase
		return ret
	}
	return *o.CompletedPhases
}

// GetCompletedPhasesOk returns a tuple with the CompletedPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetCompletedPhasesOk() (*[]OnpremUpgradePhase, bool) {
	if o == nil || o.CompletedPhases == nil {
		return nil, false
	}
	return o.CompletedPhases, true
}

// HasCompletedPhases returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasCompletedPhases() bool {
	if o != nil && o.CompletedPhases != nil {
		return true
	}

	return false
}

// SetCompletedPhases gets a reference to the given []OnpremUpgradePhase and assigns it to the CompletedPhases field.
func (o *ApplianceUpgradeAllOf) SetCompletedPhases(v []OnpremUpgradePhase) {
	o.CompletedPhases = &v
}

// GetCurrentPhase returns the CurrentPhase field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetCurrentPhase() OnpremUpgradePhase {
	if o == nil || o.CurrentPhase == nil {
		var ret OnpremUpgradePhase
		return ret
	}
	return *o.CurrentPhase
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetCurrentPhaseOk() (*OnpremUpgradePhase, bool) {
	if o == nil || o.CurrentPhase == nil {
		return nil, false
	}
	return o.CurrentPhase, true
}

// HasCurrentPhase returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasCurrentPhase() bool {
	if o != nil && o.CurrentPhase != nil {
		return true
	}

	return false
}

// SetCurrentPhase gets a reference to the given OnpremUpgradePhase and assigns it to the CurrentPhase field.
func (o *ApplianceUpgradeAllOf) SetCurrentPhase(v OnpremUpgradePhase) {
	o.CurrentPhase = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceUpgradeAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *ApplianceUpgradeAllOf) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceUpgradeAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetErrorCode() int64 {
	if o == nil || o.ErrorCode == nil {
		var ret int64
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetErrorCodeOk() (*int64, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int64 and assigns it to the ErrorCode field.
func (o *ApplianceUpgradeAllOf) SetErrorCode(v int64) {
	o.ErrorCode = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *ApplianceUpgradeAllOf) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetIsRollingBack returns the IsRollingBack field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetIsRollingBack() bool {
	if o == nil || o.IsRollingBack == nil {
		var ret bool
		return ret
	}
	return *o.IsRollingBack
}

// GetIsRollingBackOk returns a tuple with the IsRollingBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetIsRollingBackOk() (*bool, bool) {
	if o == nil || o.IsRollingBack == nil {
		return nil, false
	}
	return o.IsRollingBack, true
}

// HasIsRollingBack returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasIsRollingBack() bool {
	if o != nil && o.IsRollingBack != nil {
		return true
	}

	return false
}

// SetIsRollingBack gets a reference to the given bool and assigns it to the IsRollingBack field.
func (o *ApplianceUpgradeAllOf) SetIsRollingBack(v bool) {
	o.IsRollingBack = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetMessages() []string {
	if o == nil || o.Messages == nil {
		var ret []string
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetMessagesOk() (*[]string, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []string and assigns it to the Messages field.
func (o *ApplianceUpgradeAllOf) SetMessages(v []string) {
	o.Messages = &v
}

// GetRollbackNeeded returns the RollbackNeeded field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetRollbackNeeded() bool {
	if o == nil || o.RollbackNeeded == nil {
		var ret bool
		return ret
	}
	return *o.RollbackNeeded
}

// GetRollbackNeededOk returns a tuple with the RollbackNeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetRollbackNeededOk() (*bool, bool) {
	if o == nil || o.RollbackNeeded == nil {
		return nil, false
	}
	return o.RollbackNeeded, true
}

// HasRollbackNeeded returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasRollbackNeeded() bool {
	if o != nil && o.RollbackNeeded != nil {
		return true
	}

	return false
}

// SetRollbackNeeded gets a reference to the given bool and assigns it to the RollbackNeeded field.
func (o *ApplianceUpgradeAllOf) SetRollbackNeeded(v bool) {
	o.RollbackNeeded = &v
}

// GetRollbackPhases returns the RollbackPhases field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetRollbackPhases() []OnpremUpgradePhase {
	if o == nil || o.RollbackPhases == nil {
		var ret []OnpremUpgradePhase
		return ret
	}
	return *o.RollbackPhases
}

// GetRollbackPhasesOk returns a tuple with the RollbackPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetRollbackPhasesOk() (*[]OnpremUpgradePhase, bool) {
	if o == nil || o.RollbackPhases == nil {
		return nil, false
	}
	return o.RollbackPhases, true
}

// HasRollbackPhases returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasRollbackPhases() bool {
	if o != nil && o.RollbackPhases != nil {
		return true
	}

	return false
}

// SetRollbackPhases gets a reference to the given []OnpremUpgradePhase and assigns it to the RollbackPhases field.
func (o *ApplianceUpgradeAllOf) SetRollbackPhases(v []OnpremUpgradePhase) {
	o.RollbackPhases = &v
}

// GetRollbackStatus returns the RollbackStatus field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetRollbackStatus() string {
	if o == nil || o.RollbackStatus == nil {
		var ret string
		return ret
	}
	return *o.RollbackStatus
}

// GetRollbackStatusOk returns a tuple with the RollbackStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetRollbackStatusOk() (*string, bool) {
	if o == nil || o.RollbackStatus == nil {
		return nil, false
	}
	return o.RollbackStatus, true
}

// HasRollbackStatus returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasRollbackStatus() bool {
	if o != nil && o.RollbackStatus != nil {
		return true
	}

	return false
}

// SetRollbackStatus gets a reference to the given string and assigns it to the RollbackStatus field.
func (o *ApplianceUpgradeAllOf) SetRollbackStatus(v string) {
	o.RollbackStatus = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ApplianceUpgradeAllOf) SetServices(v []string) {
	o.Services = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceUpgradeAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceUpgradeAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTotalPhases returns the TotalPhases field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetTotalPhases() int64 {
	if o == nil || o.TotalPhases == nil {
		var ret int64
		return ret
	}
	return *o.TotalPhases
}

// GetTotalPhasesOk returns a tuple with the TotalPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetTotalPhasesOk() (*int64, bool) {
	if o == nil || o.TotalPhases == nil {
		return nil, false
	}
	return o.TotalPhases, true
}

// HasTotalPhases returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasTotalPhases() bool {
	if o != nil && o.TotalPhases != nil {
		return true
	}

	return false
}

// SetTotalPhases gets a reference to the given int64 and assigns it to the TotalPhases field.
func (o *ApplianceUpgradeAllOf) SetTotalPhases(v int64) {
	o.TotalPhases = &v
}

// GetUiPackages returns the UiPackages field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetUiPackages() []string {
	if o == nil || o.UiPackages == nil {
		var ret []string
		return ret
	}
	return *o.UiPackages
}

// GetUiPackagesOk returns a tuple with the UiPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetUiPackagesOk() (*[]string, bool) {
	if o == nil || o.UiPackages == nil {
		return nil, false
	}
	return o.UiPackages, true
}

// HasUiPackages returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasUiPackages() bool {
	if o != nil && o.UiPackages != nil {
		return true
	}

	return false
}

// SetUiPackages gets a reference to the given []string and assigns it to the UiPackages field.
func (o *ApplianceUpgradeAllOf) SetUiPackages(v []string) {
	o.UiPackages = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceUpgradeAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceUpgradeAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetImageBundle returns the ImageBundle field value if set, zero value otherwise.
func (o *ApplianceUpgradeAllOf) GetImageBundle() ApplianceImageBundleRelationship {
	if o == nil || o.ImageBundle == nil {
		var ret ApplianceImageBundleRelationship
		return ret
	}
	return *o.ImageBundle
}

// GetImageBundleOk returns a tuple with the ImageBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceUpgradeAllOf) GetImageBundleOk() (*ApplianceImageBundleRelationship, bool) {
	if o == nil || o.ImageBundle == nil {
		return nil, false
	}
	return o.ImageBundle, true
}

// HasImageBundle returns a boolean if a field has been set.
func (o *ApplianceUpgradeAllOf) HasImageBundle() bool {
	if o != nil && o.ImageBundle != nil {
		return true
	}

	return false
}

// SetImageBundle gets a reference to the given ApplianceImageBundleRelationship and assigns it to the ImageBundle field.
func (o *ApplianceUpgradeAllOf) SetImageBundle(v ApplianceImageBundleRelationship) {
	o.ImageBundle = &v
}

func (o ApplianceUpgradeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["Active"] = o.Active
	}
	if o.AutoCreated != nil {
		toSerialize["AutoCreated"] = o.AutoCreated
	}
	if o.CompletedPhases != nil {
		toSerialize["CompletedPhases"] = o.CompletedPhases
	}
	if o.CurrentPhase != nil {
		toSerialize["CurrentPhase"] = o.CurrentPhase
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Fingerprint != nil {
		toSerialize["Fingerprint"] = o.Fingerprint
	}
	if o.IsRollingBack != nil {
		toSerialize["IsRollingBack"] = o.IsRollingBack
	}
	if o.Messages != nil {
		toSerialize["Messages"] = o.Messages
	}
	if o.RollbackNeeded != nil {
		toSerialize["RollbackNeeded"] = o.RollbackNeeded
	}
	if o.RollbackPhases != nil {
		toSerialize["RollbackPhases"] = o.RollbackPhases
	}
	if o.RollbackStatus != nil {
		toSerialize["RollbackStatus"] = o.RollbackStatus
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TotalPhases != nil {
		toSerialize["TotalPhases"] = o.TotalPhases
	}
	if o.UiPackages != nil {
		toSerialize["UiPackages"] = o.UiPackages
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ImageBundle != nil {
		toSerialize["ImageBundle"] = o.ImageBundle
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceUpgradeAllOf struct {
	value *ApplianceUpgradeAllOf
	isSet bool
}

func (v NullableApplianceUpgradeAllOf) Get() *ApplianceUpgradeAllOf {
	return v.value
}

func (v *NullableApplianceUpgradeAllOf) Set(val *ApplianceUpgradeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceUpgradeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceUpgradeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceUpgradeAllOf(val *ApplianceUpgradeAllOf) *NullableApplianceUpgradeAllOf {
	return &NullableApplianceUpgradeAllOf{value: val, isSet: true}
}

func (v NullableApplianceUpgradeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceUpgradeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
