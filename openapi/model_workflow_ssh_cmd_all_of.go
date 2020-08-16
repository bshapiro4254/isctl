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

// WorkflowSshCmdAllOf Definition of the list of properties defined in 'workflow.SshCmd', excluding properties defined in parent classes.
type WorkflowSshCmdAllOf struct {
	// SSH command to execute on the remote server.
	Command *string `json:"Command,omitempty" yaml:"Command,omitempty"`
	// SSH command type to execute on the remote server. * `NonInteractiveCmd` - Execute a non-interactive SSH command on the remote server. * `InteractiveCmd` - Execute an interactive SSH command on the remote server.
	CommandType   *string                  `json:"CommandType,omitempty" yaml:"CommandType,omitempty"`
	ExpectPrompts *[]ConnectorExpectPrompt `json:"ExpectPrompts,omitempty" yaml:"ExpectPrompts,omitempty"`
	// Regex of the remote server's shell prompt.
	ShellPrompt *string `json:"ShellPrompt,omitempty" yaml:"ShellPrompt,omitempty"`
	// Expect timeout value in seconds for the shell prompt.
	ShellPromptTimeout *int64 `json:"ShellPromptTimeout,omitempty" yaml:"ShellPromptTimeout,omitempty"`
}

// NewWorkflowSshCmdAllOf instantiates a new WorkflowSshCmdAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSshCmdAllOf() *WorkflowSshCmdAllOf {
	this := WorkflowSshCmdAllOf{}
	var commandType string = "NonInteractiveCmd"
	this.CommandType = &commandType
	return &this
}

// NewWorkflowSshCmdAllOfWithDefaults instantiates a new WorkflowSshCmdAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSshCmdAllOfWithDefaults() *WorkflowSshCmdAllOf {
	this := WorkflowSshCmdAllOf{}
	var commandType string = "NonInteractiveCmd"
	this.CommandType = &commandType
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *WorkflowSshCmdAllOf) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshCmdAllOf) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *WorkflowSshCmdAllOf) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *WorkflowSshCmdAllOf) SetCommand(v string) {
	o.Command = &v
}

// GetCommandType returns the CommandType field value if set, zero value otherwise.
func (o *WorkflowSshCmdAllOf) GetCommandType() string {
	if o == nil || o.CommandType == nil {
		var ret string
		return ret
	}
	return *o.CommandType
}

// GetCommandTypeOk returns a tuple with the CommandType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshCmdAllOf) GetCommandTypeOk() (*string, bool) {
	if o == nil || o.CommandType == nil {
		return nil, false
	}
	return o.CommandType, true
}

// HasCommandType returns a boolean if a field has been set.
func (o *WorkflowSshCmdAllOf) HasCommandType() bool {
	if o != nil && o.CommandType != nil {
		return true
	}

	return false
}

// SetCommandType gets a reference to the given string and assigns it to the CommandType field.
func (o *WorkflowSshCmdAllOf) SetCommandType(v string) {
	o.CommandType = &v
}

// GetExpectPrompts returns the ExpectPrompts field value if set, zero value otherwise.
func (o *WorkflowSshCmdAllOf) GetExpectPrompts() []ConnectorExpectPrompt {
	if o == nil || o.ExpectPrompts == nil {
		var ret []ConnectorExpectPrompt
		return ret
	}
	return *o.ExpectPrompts
}

// GetExpectPromptsOk returns a tuple with the ExpectPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshCmdAllOf) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool) {
	if o == nil || o.ExpectPrompts == nil {
		return nil, false
	}
	return o.ExpectPrompts, true
}

// HasExpectPrompts returns a boolean if a field has been set.
func (o *WorkflowSshCmdAllOf) HasExpectPrompts() bool {
	if o != nil && o.ExpectPrompts != nil {
		return true
	}

	return false
}

// SetExpectPrompts gets a reference to the given []ConnectorExpectPrompt and assigns it to the ExpectPrompts field.
func (o *WorkflowSshCmdAllOf) SetExpectPrompts(v []ConnectorExpectPrompt) {
	o.ExpectPrompts = &v
}

// GetShellPrompt returns the ShellPrompt field value if set, zero value otherwise.
func (o *WorkflowSshCmdAllOf) GetShellPrompt() string {
	if o == nil || o.ShellPrompt == nil {
		var ret string
		return ret
	}
	return *o.ShellPrompt
}

// GetShellPromptOk returns a tuple with the ShellPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshCmdAllOf) GetShellPromptOk() (*string, bool) {
	if o == nil || o.ShellPrompt == nil {
		return nil, false
	}
	return o.ShellPrompt, true
}

// HasShellPrompt returns a boolean if a field has been set.
func (o *WorkflowSshCmdAllOf) HasShellPrompt() bool {
	if o != nil && o.ShellPrompt != nil {
		return true
	}

	return false
}

// SetShellPrompt gets a reference to the given string and assigns it to the ShellPrompt field.
func (o *WorkflowSshCmdAllOf) SetShellPrompt(v string) {
	o.ShellPrompt = &v
}

// GetShellPromptTimeout returns the ShellPromptTimeout field value if set, zero value otherwise.
func (o *WorkflowSshCmdAllOf) GetShellPromptTimeout() int64 {
	if o == nil || o.ShellPromptTimeout == nil {
		var ret int64
		return ret
	}
	return *o.ShellPromptTimeout
}

// GetShellPromptTimeoutOk returns a tuple with the ShellPromptTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSshCmdAllOf) GetShellPromptTimeoutOk() (*int64, bool) {
	if o == nil || o.ShellPromptTimeout == nil {
		return nil, false
	}
	return o.ShellPromptTimeout, true
}

// HasShellPromptTimeout returns a boolean if a field has been set.
func (o *WorkflowSshCmdAllOf) HasShellPromptTimeout() bool {
	if o != nil && o.ShellPromptTimeout != nil {
		return true
	}

	return false
}

// SetShellPromptTimeout gets a reference to the given int64 and assigns it to the ShellPromptTimeout field.
func (o *WorkflowSshCmdAllOf) SetShellPromptTimeout(v int64) {
	o.ShellPromptTimeout = &v
}

func (o WorkflowSshCmdAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Command != nil {
		toSerialize["Command"] = o.Command
	}
	if o.CommandType != nil {
		toSerialize["CommandType"] = o.CommandType
	}
	if o.ExpectPrompts != nil {
		toSerialize["ExpectPrompts"] = o.ExpectPrompts
	}
	if o.ShellPrompt != nil {
		toSerialize["ShellPrompt"] = o.ShellPrompt
	}
	if o.ShellPromptTimeout != nil {
		toSerialize["ShellPromptTimeout"] = o.ShellPromptTimeout
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSshCmdAllOf struct {
	value *WorkflowSshCmdAllOf
	isSet bool
}

func (v NullableWorkflowSshCmdAllOf) Get() *WorkflowSshCmdAllOf {
	return v.value
}

func (v *NullableWorkflowSshCmdAllOf) Set(val *WorkflowSshCmdAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSshCmdAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSshCmdAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSshCmdAllOf(val *WorkflowSshCmdAllOf) *NullableWorkflowSshCmdAllOf {
	return &NullableWorkflowSshCmdAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSshCmdAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSshCmdAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}