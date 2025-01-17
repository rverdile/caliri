/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ActivationKeyPoolDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivationKeyPoolDTO{}

// ActivationKeyPoolDTO struct for ActivationKeyPoolDTO
type ActivationKeyPoolDTO struct {
	PoolId string `json:"poolId"`
	Quantity *int64 `json:"quantity,omitempty"`
}

type _ActivationKeyPoolDTO ActivationKeyPoolDTO

// NewActivationKeyPoolDTO instantiates a new ActivationKeyPoolDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationKeyPoolDTO(poolId string) *ActivationKeyPoolDTO {
	this := ActivationKeyPoolDTO{}
	this.PoolId = poolId
	return &this
}

// NewActivationKeyPoolDTOWithDefaults instantiates a new ActivationKeyPoolDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationKeyPoolDTOWithDefaults() *ActivationKeyPoolDTO {
	this := ActivationKeyPoolDTO{}
	return &this
}

// GetPoolId returns the PoolId field value
func (o *ActivationKeyPoolDTO) GetPoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *ActivationKeyPoolDTO) GetPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *ActivationKeyPoolDTO) SetPoolId(v string) {
	o.PoolId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ActivationKeyPoolDTO) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivationKeyPoolDTO) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ActivationKeyPoolDTO) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *ActivationKeyPoolDTO) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o ActivationKeyPoolDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivationKeyPoolDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["poolId"] = o.PoolId
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

func (o *ActivationKeyPoolDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"poolId",
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

	varActivationKeyPoolDTO := _ActivationKeyPoolDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivationKeyPoolDTO)

	if err != nil {
		return err
	}

	*o = ActivationKeyPoolDTO(varActivationKeyPoolDTO)

	return err
}

type NullableActivationKeyPoolDTO struct {
	value *ActivationKeyPoolDTO
	isSet bool
}

func (v NullableActivationKeyPoolDTO) Get() *ActivationKeyPoolDTO {
	return v.value
}

func (v *NullableActivationKeyPoolDTO) Set(val *ActivationKeyPoolDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationKeyPoolDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationKeyPoolDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationKeyPoolDTO(val *ActivationKeyPoolDTO) *NullableActivationKeyPoolDTO {
	return &NullableActivationKeyPoolDTO{value: val, isSet: true}
}

func (v NullableActivationKeyPoolDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationKeyPoolDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


