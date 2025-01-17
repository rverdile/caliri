/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the CertificateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateDTO{}

// CertificateDTO Represents certificate details
type CertificateDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Cert *string `json:"cert,omitempty"`
	Serial *CertificateSerialDTO `json:"serial,omitempty"`
}

// NewCertificateDTO instantiates a new CertificateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDTO() *CertificateDTO {
	this := CertificateDTO{}
	return &this
}

// NewCertificateDTOWithDefaults instantiates a new CertificateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDTOWithDefaults() *CertificateDTO {
	this := CertificateDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CertificateDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CertificateDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CertificateDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CertificateDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CertificateDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *CertificateDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateDTO) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CertificateDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CertificateDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CertificateDTO) SetKey(v string) {
	o.Key = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *CertificateDTO) GetCert() string {
	if o == nil || IsNil(o.Cert) {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetCertOk() (*string, bool) {
	if o == nil || IsNil(o.Cert) {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *CertificateDTO) HasCert() bool {
	if o != nil && !IsNil(o.Cert) {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *CertificateDTO) SetCert(v string) {
	o.Cert = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CertificateDTO) GetSerial() CertificateSerialDTO {
	if o == nil || IsNil(o.Serial) {
		var ret CertificateSerialDTO
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDTO) GetSerialOk() (*CertificateSerialDTO, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CertificateDTO) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given CertificateSerialDTO and assigns it to the Serial field.
func (o *CertificateDTO) SetSerial(v CertificateSerialDTO) {
	o.Serial = &v
}

func (o CertificateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Cert) {
		toSerialize["cert"] = o.Cert
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableCertificateDTO struct {
	value *CertificateDTO
	isSet bool
}

func (v NullableCertificateDTO) Get() *CertificateDTO {
	return v.value
}

func (v *NullableCertificateDTO) Set(val *CertificateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDTO(val *CertificateDTO) *NullableCertificateDTO {
	return &NullableCertificateDTO{value: val, isSet: true}
}

func (v NullableCertificateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


