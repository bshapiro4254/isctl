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

// WorkflowFileTransfer Message to transfer a file from Intersight connected device to remote server.
type WorkflowFileTransfer struct {
	ConnectorBaseMessage `yaml:"ConnectorBaseMessage,inline"`
	// Destination file path on the target server.
	DestinationFilePath *string `json:"DestinationFilePath,omitempty" yaml:"DestinationFilePath,omitempty"`
	// File permission to set on the transferred file.
	FileMode *int64 `json:"FileMode,omitempty" yaml:"FileMode,omitempty"`
	// Source file path on the Intersight connected device.
	SourceFilePath *string `json:"SourceFilePath,omitempty" yaml:"SourceFilePath,omitempty"`
}

// NewWorkflowFileTransfer instantiates a new WorkflowFileTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFileTransfer() *WorkflowFileTransfer {
	this := WorkflowFileTransfer{}
	return &this
}

// NewWorkflowFileTransferWithDefaults instantiates a new WorkflowFileTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFileTransferWithDefaults() *WorkflowFileTransfer {
	this := WorkflowFileTransfer{}
	return &this
}

// GetDestinationFilePath returns the DestinationFilePath field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetDestinationFilePath() string {
	if o == nil || o.DestinationFilePath == nil {
		var ret string
		return ret
	}
	return *o.DestinationFilePath
}

// GetDestinationFilePathOk returns a tuple with the DestinationFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetDestinationFilePathOk() (*string, bool) {
	if o == nil || o.DestinationFilePath == nil {
		return nil, false
	}
	return o.DestinationFilePath, true
}

// HasDestinationFilePath returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasDestinationFilePath() bool {
	if o != nil && o.DestinationFilePath != nil {
		return true
	}

	return false
}

// SetDestinationFilePath gets a reference to the given string and assigns it to the DestinationFilePath field.
func (o *WorkflowFileTransfer) SetDestinationFilePath(v string) {
	o.DestinationFilePath = &v
}

// GetFileMode returns the FileMode field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetFileMode() int64 {
	if o == nil || o.FileMode == nil {
		var ret int64
		return ret
	}
	return *o.FileMode
}

// GetFileModeOk returns a tuple with the FileMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetFileModeOk() (*int64, bool) {
	if o == nil || o.FileMode == nil {
		return nil, false
	}
	return o.FileMode, true
}

// HasFileMode returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasFileMode() bool {
	if o != nil && o.FileMode != nil {
		return true
	}

	return false
}

// SetFileMode gets a reference to the given int64 and assigns it to the FileMode field.
func (o *WorkflowFileTransfer) SetFileMode(v int64) {
	o.FileMode = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *WorkflowFileTransfer) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

func (o WorkflowFileTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if o.DestinationFilePath != nil {
		toSerialize["DestinationFilePath"] = o.DestinationFilePath
	}
	if o.FileMode != nil {
		toSerialize["FileMode"] = o.FileMode
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowFileTransfer struct {
	value *WorkflowFileTransfer
	isSet bool
}

func (v NullableWorkflowFileTransfer) Get() *WorkflowFileTransfer {
	return v.value
}

func (v *NullableWorkflowFileTransfer) Set(val *WorkflowFileTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFileTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFileTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFileTransfer(val *WorkflowFileTransfer) *NullableWorkflowFileTransfer {
	return &NullableWorkflowFileTransfer{value: val, isSet: true}
}

func (v NullableWorkflowFileTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFileTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
