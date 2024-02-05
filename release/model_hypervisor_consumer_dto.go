/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the HypervisorConsumerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HypervisorConsumerDTO{}

// HypervisorConsumerDTO struct for HypervisorConsumerDTO
type HypervisorConsumerDTO struct {
	Uuid *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	Owner *NestedOwnerDTO `json:"owner,omitempty"`
}

// NewHypervisorConsumerDTO instantiates a new HypervisorConsumerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHypervisorConsumerDTO() *HypervisorConsumerDTO {
	this := HypervisorConsumerDTO{}
	return &this
}

// NewHypervisorConsumerDTOWithDefaults instantiates a new HypervisorConsumerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHypervisorConsumerDTOWithDefaults() *HypervisorConsumerDTO {
	this := HypervisorConsumerDTO{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HypervisorConsumerDTO) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorConsumerDTO) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HypervisorConsumerDTO) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HypervisorConsumerDTO) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HypervisorConsumerDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorConsumerDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HypervisorConsumerDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HypervisorConsumerDTO) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *HypervisorConsumerDTO) GetOwner() NestedOwnerDTO {
	if o == nil || IsNil(o.Owner) {
		var ret NestedOwnerDTO
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HypervisorConsumerDTO) GetOwnerOk() (*NestedOwnerDTO, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *HypervisorConsumerDTO) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NestedOwnerDTO and assigns it to the Owner field.
func (o *HypervisorConsumerDTO) SetOwner(v NestedOwnerDTO) {
	o.Owner = &v
}

func (o HypervisorConsumerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HypervisorConsumerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableHypervisorConsumerDTO struct {
	value *HypervisorConsumerDTO
	isSet bool
}

func (v NullableHypervisorConsumerDTO) Get() *HypervisorConsumerDTO {
	return v.value
}

func (v *NullableHypervisorConsumerDTO) Set(val *HypervisorConsumerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableHypervisorConsumerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableHypervisorConsumerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHypervisorConsumerDTO(val *HypervisorConsumerDTO) *NullableHypervisorConsumerDTO {
	return &NullableHypervisorConsumerDTO{value: val, isSet: true}
}

func (v NullableHypervisorConsumerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHypervisorConsumerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

