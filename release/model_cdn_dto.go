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

// checks if the CdnDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnDTO{}

// CdnDTO DTO representing a Cdn
type CdnDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Certificate *CertificateDTO `json:"certificate,omitempty"`
}

// NewCdnDTO instantiates a new CdnDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnDTO() *CdnDTO {
	this := CdnDTO{}
	return &this
}

// NewCdnDTOWithDefaults instantiates a new CdnDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnDTOWithDefaults() *CdnDTO {
	this := CdnDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CdnDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CdnDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *CdnDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CdnDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CdnDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *CdnDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CdnDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CdnDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CdnDTO) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CdnDTO) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CdnDTO) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CdnDTO) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CdnDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CdnDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CdnDTO) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CdnDTO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CdnDTO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CdnDTO) SetUrl(v string) {
	o.Url = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CdnDTO) GetCertificate() CertificateDTO {
	if o == nil || IsNil(o.Certificate) {
		var ret CertificateDTO
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnDTO) GetCertificateOk() (*CertificateDTO, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CdnDTO) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateDTO and assigns it to the Certificate field.
func (o *CdnDTO) SetCertificate(v CertificateDTO) {
	o.Certificate = &v
}

func (o CdnDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return toSerialize, nil
}

type NullableCdnDTO struct {
	value *CdnDTO
	isSet bool
}

func (v NullableCdnDTO) Get() *CdnDTO {
	return v.value
}

func (v *NullableCdnDTO) Set(val *CdnDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnDTO(val *CdnDTO) *NullableCdnDTO {
	return &NullableCdnDTO{value: val, isSet: true}
}

func (v NullableCdnDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


