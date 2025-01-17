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

// checks if the UeberCertificateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeberCertificateDTO{}

// UeberCertificateDTO Represents an ueber certificate
type UeberCertificateDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	Cert *string `json:"cert,omitempty"`
	Serial *CertificateSerialDTO `json:"serial,omitempty"`
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
}

// NewUeberCertificateDTO instantiates a new UeberCertificateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeberCertificateDTO() *UeberCertificateDTO {
	this := UeberCertificateDTO{}
	return &this
}

// NewUeberCertificateDTOWithDefaults instantiates a new UeberCertificateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeberCertificateDTOWithDefaults() *UeberCertificateDTO {
	this := UeberCertificateDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *UeberCertificateDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *UeberCertificateDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UeberCertificateDTO) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *UeberCertificateDTO) SetKey(v string) {
	o.Key = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetCert() string {
	if o == nil || IsNil(o.Cert) {
		var ret string
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetCertOk() (*string, bool) {
	if o == nil || IsNil(o.Cert) {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasCert() bool {
	if o != nil && !IsNil(o.Cert) {
		return true
	}

	return false
}

// SetCert gets a reference to the given string and assigns it to the Cert field.
func (o *UeberCertificateDTO) SetCert(v string) {
	o.Cert = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetSerial() CertificateSerialDTO {
	if o == nil || IsNil(o.Serial) {
		var ret CertificateSerialDTO
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetSerialOk() (*CertificateSerialDTO, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given CertificateSerialDTO and assigns it to the Serial field.
func (o *UeberCertificateDTO) SetSerial(v CertificateSerialDTO) {
	o.Serial = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UeberCertificateDTO) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeberCertificateDTO) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UeberCertificateDTO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *UeberCertificateDTO) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

func (o UeberCertificateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeberCertificateDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableUeberCertificateDTO struct {
	value *UeberCertificateDTO
	isSet bool
}

func (v NullableUeberCertificateDTO) Get() *UeberCertificateDTO {
	return v.value
}

func (v *NullableUeberCertificateDTO) Set(val *UeberCertificateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUeberCertificateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUeberCertificateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeberCertificateDTO(val *UeberCertificateDTO) *NullableUeberCertificateDTO {
	return &NullableUeberCertificateDTO{value: val, isSet: true}
}

func (v NullableUeberCertificateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeberCertificateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


