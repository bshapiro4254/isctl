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

// TelemetryDruidFieldAccessorPostAggregator These post-aggregators return the value produced by the specified aggregator. 'fieldName' refers to the output name of the aggregator given in the aggregations portion of the query. For complex aggregators, like \"cardinality\" and \"hyperUnique\", the type of the post-aggregator determines what the post-aggregator will return. Use type \"fieldAccess\" to return the raw aggregation object, or use type \"finalizingFieldAccess\" to return a finalized value, such as an estimated cardinality.
type TelemetryDruidFieldAccessorPostAggregator struct {
	// The post-aggregator type.
	Type string `json:"type" yaml:"type"`
	// Output name for the post-aggregator.
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	// Name of the metric column.
	FieldName *string `json:"fieldName,omitempty" yaml:"fieldName,omitempty"`
}

// NewTelemetryDruidFieldAccessorPostAggregator instantiates a new TelemetryDruidFieldAccessorPostAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidFieldAccessorPostAggregator(type_ string) *TelemetryDruidFieldAccessorPostAggregator {
	this := TelemetryDruidFieldAccessorPostAggregator{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidFieldAccessorPostAggregatorWithDefaults instantiates a new TelemetryDruidFieldAccessorPostAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidFieldAccessorPostAggregatorWithDefaults() *TelemetryDruidFieldAccessorPostAggregator {
	this := TelemetryDruidFieldAccessorPostAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidFieldAccessorPostAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFieldAccessorPostAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidFieldAccessorPostAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidFieldAccessorPostAggregator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFieldAccessorPostAggregator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidFieldAccessorPostAggregator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidFieldAccessorPostAggregator) SetName(v string) {
	o.Name = &v
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *TelemetryDruidFieldAccessorPostAggregator) GetFieldName() string {
	if o == nil || o.FieldName == nil {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidFieldAccessorPostAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil || o.FieldName == nil {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *TelemetryDruidFieldAccessorPostAggregator) HasFieldName() bool {
	if o != nil && o.FieldName != nil {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *TelemetryDruidFieldAccessorPostAggregator) SetFieldName(v string) {
	o.FieldName = &v
}

func (o TelemetryDruidFieldAccessorPostAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FieldName != nil {
		toSerialize["fieldName"] = o.FieldName
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryDruidFieldAccessorPostAggregator struct {
	value *TelemetryDruidFieldAccessorPostAggregator
	isSet bool
}

func (v NullableTelemetryDruidFieldAccessorPostAggregator) Get() *TelemetryDruidFieldAccessorPostAggregator {
	return v.value
}

func (v *NullableTelemetryDruidFieldAccessorPostAggregator) Set(val *TelemetryDruidFieldAccessorPostAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFieldAccessorPostAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFieldAccessorPostAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFieldAccessorPostAggregator(val *TelemetryDruidFieldAccessorPostAggregator) *NullableTelemetryDruidFieldAccessorPostAggregator {
	return &NullableTelemetryDruidFieldAccessorPostAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidFieldAccessorPostAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFieldAccessorPostAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
