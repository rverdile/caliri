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

// checks if the SystemPurposeAttributesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemPurposeAttributesDTO{}

// SystemPurposeAttributesDTO Represents system purpose attribute details
type SystemPurposeAttributesDTO struct {
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
	SystemPurposeAttributes *map[string][]string `json:"systemPurposeAttributes,omitempty"`
}

// NewSystemPurposeAttributesDTO instantiates a new SystemPurposeAttributesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemPurposeAttributesDTO() *SystemPurposeAttributesDTO {
	this := SystemPurposeAttributesDTO{}
	return &this
}

// NewSystemPurposeAttributesDTOWithDefaults instantiates a new SystemPurposeAttributesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemPurposeAttributesDTOWithDefaults() *SystemPurposeAttributesDTO {
	this := SystemPurposeAttributesDTO{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *SystemPurposeAttributesDTO) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeAttributesDTO) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *SystemPurposeAttributesDTO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *SystemPurposeAttributesDTO) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

// GetSystemPurposeAttributes returns the SystemPurposeAttributes field value if set, zero value otherwise.
func (o *SystemPurposeAttributesDTO) GetSystemPurposeAttributes() map[string][]string {
	if o == nil || IsNil(o.SystemPurposeAttributes) {
		var ret map[string][]string
		return ret
	}
	return *o.SystemPurposeAttributes
}

// GetSystemPurposeAttributesOk returns a tuple with the SystemPurposeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemPurposeAttributesDTO) GetSystemPurposeAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.SystemPurposeAttributes) {
		return nil, false
	}
	return o.SystemPurposeAttributes, true
}

// HasSystemPurposeAttributes returns a boolean if a field has been set.
func (o *SystemPurposeAttributesDTO) HasSystemPurposeAttributes() bool {
	if o != nil && !IsNil(o.SystemPurposeAttributes) {
		return true
	}

	return false
}

// SetSystemPurposeAttributes gets a reference to the given map[string][]string and assigns it to the SystemPurposeAttributes field.
func (o *SystemPurposeAttributesDTO) SetSystemPurposeAttributes(v map[string][]string) {
	o.SystemPurposeAttributes = &v
}

func (o SystemPurposeAttributesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemPurposeAttributesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.SystemPurposeAttributes) {
		toSerialize["systemPurposeAttributes"] = o.SystemPurposeAttributes
	}
	return toSerialize, nil
}

type NullableSystemPurposeAttributesDTO struct {
	value *SystemPurposeAttributesDTO
	isSet bool
}

func (v NullableSystemPurposeAttributesDTO) Get() *SystemPurposeAttributesDTO {
	return v.value
}

func (v *NullableSystemPurposeAttributesDTO) Set(val *SystemPurposeAttributesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemPurposeAttributesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemPurposeAttributesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemPurposeAttributesDTO(val *SystemPurposeAttributesDTO) *NullableSystemPurposeAttributesDTO {
	return &NullableSystemPurposeAttributesDTO{value: val, isSet: true}
}

func (v NullableSystemPurposeAttributesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemPurposeAttributesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


