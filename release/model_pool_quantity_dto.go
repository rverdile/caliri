/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the PoolQuantityDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolQuantityDTO{}

// PoolQuantityDTO DTO representing pool quantity & pool.
type PoolQuantityDTO struct {
	Quantity *int32 `json:"quantity,omitempty"`
	Pool *PoolDTO `json:"pool,omitempty"`
}

// NewPoolQuantityDTO instantiates a new PoolQuantityDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolQuantityDTO() *PoolQuantityDTO {
	this := PoolQuantityDTO{}
	return &this
}

// NewPoolQuantityDTOWithDefaults instantiates a new PoolQuantityDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolQuantityDTOWithDefaults() *PoolQuantityDTO {
	this := PoolQuantityDTO{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PoolQuantityDTO) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolQuantityDTO) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PoolQuantityDTO) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PoolQuantityDTO) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *PoolQuantityDTO) GetPool() PoolDTO {
	if o == nil || IsNil(o.Pool) {
		var ret PoolDTO
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolQuantityDTO) GetPoolOk() (*PoolDTO, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *PoolQuantityDTO) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given PoolDTO and assigns it to the Pool field.
func (o *PoolQuantityDTO) SetPool(v PoolDTO) {
	o.Pool = &v
}

func (o PoolQuantityDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolQuantityDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	return toSerialize, nil
}

type NullablePoolQuantityDTO struct {
	value *PoolQuantityDTO
	isSet bool
}

func (v NullablePoolQuantityDTO) Get() *PoolQuantityDTO {
	return v.value
}

func (v *NullablePoolQuantityDTO) Set(val *PoolQuantityDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolQuantityDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolQuantityDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolQuantityDTO(val *PoolQuantityDTO) *NullablePoolQuantityDTO {
	return &NullablePoolQuantityDTO{value: val, isSet: true}
}

func (v NullablePoolQuantityDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolQuantityDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


