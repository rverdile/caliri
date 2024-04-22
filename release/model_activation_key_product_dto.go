/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActivationKeyProductDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivationKeyProductDTO{}

// ActivationKeyProductDTO struct for ActivationKeyProductDTO
type ActivationKeyProductDTO struct {
	ProductId string `json:"productId"`
}

type _ActivationKeyProductDTO ActivationKeyProductDTO

// NewActivationKeyProductDTO instantiates a new ActivationKeyProductDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationKeyProductDTO(productId string) *ActivationKeyProductDTO {
	this := ActivationKeyProductDTO{}
	this.ProductId = productId
	return &this
}

// NewActivationKeyProductDTOWithDefaults instantiates a new ActivationKeyProductDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationKeyProductDTOWithDefaults() *ActivationKeyProductDTO {
	this := ActivationKeyProductDTO{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *ActivationKeyProductDTO) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *ActivationKeyProductDTO) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *ActivationKeyProductDTO) SetProductId(v string) {
	o.ProductId = v
}

func (o ActivationKeyProductDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivationKeyProductDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	return toSerialize, nil
}

func (o *ActivationKeyProductDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productId",
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

	varActivationKeyProductDTO := _ActivationKeyProductDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivationKeyProductDTO)

	if err != nil {
		return err
	}

	*o = ActivationKeyProductDTO(varActivationKeyProductDTO)

	return err
}

type NullableActivationKeyProductDTO struct {
	value *ActivationKeyProductDTO
	isSet bool
}

func (v NullableActivationKeyProductDTO) Get() *ActivationKeyProductDTO {
	return v.value
}

func (v *NullableActivationKeyProductDTO) Set(val *ActivationKeyProductDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationKeyProductDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationKeyProductDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationKeyProductDTO(val *ActivationKeyProductDTO) *NullableActivationKeyProductDTO {
	return &NullableActivationKeyProductDTO{value: val, isSet: true}
}

func (v NullableActivationKeyProductDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationKeyProductDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


