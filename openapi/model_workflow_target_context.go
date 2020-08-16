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

// WorkflowTargetContext Target details like moid, type and name of the Intersight managed object.
type WorkflowTargetContext struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Moid of the target Intersight managed object.
	TargetMoid *string `json:"TargetMoid,omitempty" yaml:"TargetMoid,omitempty"`
	// Name of the target instance.
	TargetName *string `json:"TargetName,omitempty" yaml:"TargetName,omitempty"`
	// Object type of the target Intersight managed object.
	TargetType *string `json:"TargetType,omitempty" yaml:"TargetType,omitempty"`
}

// NewWorkflowTargetContext instantiates a new WorkflowTargetContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetContext() *WorkflowTargetContext {
	this := WorkflowTargetContext{}
	return &this
}

// NewWorkflowTargetContextWithDefaults instantiates a new WorkflowTargetContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetContextWithDefaults() *WorkflowTargetContext {
	this := WorkflowTargetContext{}
	return &this
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *WorkflowTargetContext) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *WorkflowTargetContext) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowTargetContext) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetContext) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowTargetContext) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowTargetContext) SetTargetType(v string) {
	o.TargetType = &v
}

func (o WorkflowTargetContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTargetContext struct {
	value *WorkflowTargetContext
	isSet bool
}

func (v NullableWorkflowTargetContext) Get() *WorkflowTargetContext {
	return v.value
}

func (v *NullableWorkflowTargetContext) Set(val *WorkflowTargetContext) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetContext) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetContext(val *WorkflowTargetContext) *NullableWorkflowTargetContext {
	return &NullableWorkflowTargetContext{value: val, isSet: true}
}

func (v NullableWorkflowTargetContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
