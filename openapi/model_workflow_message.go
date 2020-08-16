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

// WorkflowMessage Intermediate Task or Workflow message with severity.
type WorkflowMessage struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// An i18n message that can be translated in multiple languages to support internationalization.
	Message *string `json:"Message,omitempty" yaml:"Message,omitempty"`
	// The severity of the Task or Workflow message warning/error/info etc. * `Info` - The enum represents the log level to be used to convey info message. * `Warning` - The enum represents the log level to be used to convey warning message. * `Debug` - The enum represents the log level to be used to convey debug message. * `Error` - The enum represents the log level to be used to convey error message.
	Severity *string `json:"Severity,omitempty" yaml:"Severity,omitempty"`
}

// NewWorkflowMessage instantiates a new WorkflowMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMessage() *WorkflowMessage {
	this := WorkflowMessage{}
	var severity string = "Info"
	this.Severity = &severity
	return &this
}

// NewWorkflowMessageWithDefaults instantiates a new WorkflowMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMessageWithDefaults() *WorkflowMessage {
	this := WorkflowMessage{}
	var severity string = "Info"
	this.Severity = &severity
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WorkflowMessage) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WorkflowMessage) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WorkflowMessage) SetMessage(v string) {
	o.Message = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *WorkflowMessage) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMessage) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *WorkflowMessage) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *WorkflowMessage) SetSeverity(v string) {
	o.Severity = &v
}

func (o WorkflowMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowMessage struct {
	value *WorkflowMessage
	isSet bool
}

func (v NullableWorkflowMessage) Get() *WorkflowMessage {
	return v.value
}

func (v *NullableWorkflowMessage) Set(val *WorkflowMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMessage(val *WorkflowMessage) *NullableWorkflowMessage {
	return &NullableWorkflowMessage{value: val, isSet: true}
}

func (v NullableWorkflowMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
