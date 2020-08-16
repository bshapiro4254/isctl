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

// WorkflowWorkflowMeta Contains a workflow definition which is a sequence of tasks to execute.
type WorkflowWorkflowMeta struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The description for the workflow.
	Description     *string   `json:"Description,omitempty" yaml:"Description,omitempty"`
	InputParameters *[]string `json:"InputParameters,omitempty" yaml:"InputParameters,omitempty"`
	// The name given to the workflow.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// The workflow output parameters.
	OutputParameters interface{} `json:"OutputParameters,omitempty" yaml:"OutputParameters,omitempty"`
	// When true, this workflow can be retried for 2 weeks since the last modification of the workflow.
	Retryable *bool `json:"Retryable,omitempty" yaml:"Retryable,omitempty"`
	// The Conductor schema version that decides what attribute can be supported.
	SchemaVersion *int64 `json:"SchemaVersion,omitempty" yaml:"SchemaVersion,omitempty"`
	// The src is workflow owner service.
	Src *string `json:"Src,omitempty" yaml:"Src,omitempty"`
	// The tasks contained inside of the workflow.
	Tasks interface{} `json:"Tasks,omitempty" yaml:"Tasks,omitempty"`
	// The type of workflow definition. * `SystemDefined` - System defined workflow definition. * `UserDefined` - User defined workflow definition. * `Dynamic` - Dynamically defined workflow definition.
	Type *string `json:"Type,omitempty" yaml:"Type,omitempty"`
	// The version for the workflow so we can support multiple versions for the same workflow name.
	Version *int64 `json:"Version,omitempty" yaml:"Version,omitempty"`
	// Parameter decides if workflows will wait for a duplicate to finish before starting a new one.
	WaitOnDuplicate *bool                                 `json:"WaitOnDuplicate,omitempty" yaml:"WaitOnDuplicate,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty" yaml:"Organization,omitempty"`
}

// NewWorkflowWorkflowMeta instantiates a new WorkflowWorkflowMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowMeta() *WorkflowWorkflowMeta {
	this := WorkflowWorkflowMeta{}
	var type_ string = "SystemDefined"
	this.Type = &type_
	return &this
}

// NewWorkflowWorkflowMetaWithDefaults instantiates a new WorkflowWorkflowMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowMetaWithDefaults() *WorkflowWorkflowMeta {
	this := WorkflowWorkflowMeta{}
	var type_ string = "SystemDefined"
	this.Type = &type_
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWorkflowMeta) SetDescription(v string) {
	o.Description = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetInputParameters() []string {
	if o == nil || o.InputParameters == nil {
		var ret []string
		return ret
	}
	return *o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetInputParametersOk() (*[]string, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given []string and assigns it to the InputParameters field.
func (o *WorkflowWorkflowMeta) SetInputParameters(v []string) {
	o.InputParameters = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWorkflowMeta) SetName(v string) {
	o.Name = &v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowMeta) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowMeta) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || o.OutputParameters == nil {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasOutputParameters() bool {
	if o != nil && o.OutputParameters != nil {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *WorkflowWorkflowMeta) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowMeta) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetSchemaVersion() int64 {
	if o == nil || o.SchemaVersion == nil {
		var ret int64
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetSchemaVersionOk() (*int64, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int64 and assigns it to the SchemaVersion field.
func (o *WorkflowWorkflowMeta) SetSchemaVersion(v int64) {
	o.SchemaVersion = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *WorkflowWorkflowMeta) SetSrc(v string) {
	o.Src = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowMeta) GetTasks() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowMeta) GetTasksOk() (*interface{}, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given interface{} and assigns it to the Tasks field.
func (o *WorkflowWorkflowMeta) SetTasks(v interface{}) {
	o.Tasks = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowWorkflowMeta) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowWorkflowMeta) SetVersion(v int64) {
	o.Version = &v
}

// GetWaitOnDuplicate returns the WaitOnDuplicate field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetWaitOnDuplicate() bool {
	if o == nil || o.WaitOnDuplicate == nil {
		var ret bool
		return ret
	}
	return *o.WaitOnDuplicate
}

// GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetWaitOnDuplicateOk() (*bool, bool) {
	if o == nil || o.WaitOnDuplicate == nil {
		return nil, false
	}
	return o.WaitOnDuplicate, true
}

// HasWaitOnDuplicate returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasWaitOnDuplicate() bool {
	if o != nil && o.WaitOnDuplicate != nil {
		return true
	}

	return false
}

// SetWaitOnDuplicate gets a reference to the given bool and assigns it to the WaitOnDuplicate field.
func (o *WorkflowWorkflowMeta) SetWaitOnDuplicate(v bool) {
	o.WaitOnDuplicate = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowWorkflowMeta) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowMeta) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowWorkflowMeta) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *WorkflowWorkflowMeta) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o WorkflowWorkflowMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OutputParameters != nil {
		toSerialize["OutputParameters"] = o.OutputParameters
	}
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	if o.SchemaVersion != nil {
		toSerialize["SchemaVersion"] = o.SchemaVersion
	}
	if o.Src != nil {
		toSerialize["Src"] = o.Src
	}
	if o.Tasks != nil {
		toSerialize["Tasks"] = o.Tasks
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WaitOnDuplicate != nil {
		toSerialize["WaitOnDuplicate"] = o.WaitOnDuplicate
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowWorkflowMeta struct {
	value *WorkflowWorkflowMeta
	isSet bool
}

func (v NullableWorkflowWorkflowMeta) Get() *WorkflowWorkflowMeta {
	return v.value
}

func (v *NullableWorkflowWorkflowMeta) Set(val *WorkflowWorkflowMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowMeta(val *WorkflowWorkflowMeta) *NullableWorkflowWorkflowMeta {
	return &NullableWorkflowWorkflowMeta{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
