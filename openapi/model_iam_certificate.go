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

// IamCertificate Holds a certificate, signed by a CAcert.
type IamCertificate struct {
	MoBaseMo    `yaml:"MoBaseMo,inline"`
	Certificate *X509Certificate `json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
	// Status of the certificate. * `PendingValidation` - The certificate has not been validated. * `Valid` - The certificate is valid. * `Invalid` - Ther certificate is invalid.
	Status             *string                            `json:"Status,omitempty" yaml:"Status,omitempty"`
	CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty" yaml:"CertificateRequest,omitempty"`
}

// NewIamCertificate instantiates a new IamCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamCertificate() *IamCertificate {
	this := IamCertificate{}
	var status string = "PendingValidation"
	this.Status = &status
	return &this
}

// NewIamCertificateWithDefaults instantiates a new IamCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamCertificateWithDefaults() *IamCertificate {
	this := IamCertificate{}
	var status string = "PendingValidation"
	this.Status = &status
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *IamCertificate) GetCertificate() X509Certificate {
	if o == nil || o.Certificate == nil {
		var ret X509Certificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetCertificateOk() (*X509Certificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *IamCertificate) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given X509Certificate and assigns it to the Certificate field.
func (o *IamCertificate) SetCertificate(v X509Certificate) {
	o.Certificate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamCertificate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamCertificate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamCertificate) SetStatus(v string) {
	o.Status = &v
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamCertificate) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamCertificate) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamCertificate) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}
	return json.Marshal(toSerialize)
}

type NullableIamCertificate struct {
	value *IamCertificate
	isSet bool
}

func (v NullableIamCertificate) Get() *IamCertificate {
	return v.value
}

func (v *NullableIamCertificate) Set(val *IamCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamCertificate(val *IamCertificate) *NullableIamCertificate {
	return &NullableIamCertificate{value: val, isSet: true}
}

func (v NullableIamCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
