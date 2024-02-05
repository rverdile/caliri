/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ProductContentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductContentDTO{}

// ProductContentDTO Product content mapping exposed to the API
type ProductContentDTO struct {
	Content ContentDTO `json:"content"`
	Enabled *bool `json:"enabled,omitempty"`
}

type _ProductContentDTO ProductContentDTO

// NewProductContentDTO instantiates a new ProductContentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductContentDTO(content ContentDTO) *ProductContentDTO {
	this := ProductContentDTO{}
	this.Content = content
	return &this
}

// NewProductContentDTOWithDefaults instantiates a new ProductContentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductContentDTOWithDefaults() *ProductContentDTO {
	this := ProductContentDTO{}
	return &this
}

// GetContent returns the Content field value
func (o *ProductContentDTO) GetContent() ContentDTO {
	if o == nil {
		var ret ContentDTO
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ProductContentDTO) GetContentOk() (*ContentDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ProductContentDTO) SetContent(v ContentDTO) {
	o.Content = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProductContentDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductContentDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProductContentDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProductContentDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ProductContentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductContentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *ProductContentDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProductContentDTO := _ProductContentDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductContentDTO)

	if err != nil {
		return err
	}

	*o = ProductContentDTO(varProductContentDTO)

	return err
}

type NullableProductContentDTO struct {
	value *ProductContentDTO
	isSet bool
}

func (v NullableProductContentDTO) Get() *ProductContentDTO {
	return v.value
}

func (v *NullableProductContentDTO) Set(val *ProductContentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProductContentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProductContentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductContentDTO(val *ProductContentDTO) *NullableProductContentDTO {
	return &NullableProductContentDTO{value: val, isSet: true}
}

func (v NullableProductContentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductContentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

