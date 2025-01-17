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

// checks if the ConsumerActivationKeyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumerActivationKeyDTO{}

// ConsumerActivationKeyDTO Represent activation keys used by consumer in registration
type ConsumerActivationKeyDTO struct {
	ActivationKeyName *string `json:"activationKeyName,omitempty"`
	ActivationKeyId *string `json:"activationKeyId,omitempty"`
}

// NewConsumerActivationKeyDTO instantiates a new ConsumerActivationKeyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerActivationKeyDTO() *ConsumerActivationKeyDTO {
	this := ConsumerActivationKeyDTO{}
	return &this
}

// NewConsumerActivationKeyDTOWithDefaults instantiates a new ConsumerActivationKeyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerActivationKeyDTOWithDefaults() *ConsumerActivationKeyDTO {
	this := ConsumerActivationKeyDTO{}
	return &this
}

// GetActivationKeyName returns the ActivationKeyName field value if set, zero value otherwise.
func (o *ConsumerActivationKeyDTO) GetActivationKeyName() string {
	if o == nil || IsNil(o.ActivationKeyName) {
		var ret string
		return ret
	}
	return *o.ActivationKeyName
}

// GetActivationKeyNameOk returns a tuple with the ActivationKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerActivationKeyDTO) GetActivationKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActivationKeyName) {
		return nil, false
	}
	return o.ActivationKeyName, true
}

// HasActivationKeyName returns a boolean if a field has been set.
func (o *ConsumerActivationKeyDTO) HasActivationKeyName() bool {
	if o != nil && !IsNil(o.ActivationKeyName) {
		return true
	}

	return false
}

// SetActivationKeyName gets a reference to the given string and assigns it to the ActivationKeyName field.
func (o *ConsumerActivationKeyDTO) SetActivationKeyName(v string) {
	o.ActivationKeyName = &v
}

// GetActivationKeyId returns the ActivationKeyId field value if set, zero value otherwise.
func (o *ConsumerActivationKeyDTO) GetActivationKeyId() string {
	if o == nil || IsNil(o.ActivationKeyId) {
		var ret string
		return ret
	}
	return *o.ActivationKeyId
}

// GetActivationKeyIdOk returns a tuple with the ActivationKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerActivationKeyDTO) GetActivationKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivationKeyId) {
		return nil, false
	}
	return o.ActivationKeyId, true
}

// HasActivationKeyId returns a boolean if a field has been set.
func (o *ConsumerActivationKeyDTO) HasActivationKeyId() bool {
	if o != nil && !IsNil(o.ActivationKeyId) {
		return true
	}

	return false
}

// SetActivationKeyId gets a reference to the given string and assigns it to the ActivationKeyId field.
func (o *ConsumerActivationKeyDTO) SetActivationKeyId(v string) {
	o.ActivationKeyId = &v
}

func (o ConsumerActivationKeyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumerActivationKeyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivationKeyName) {
		toSerialize["activationKeyName"] = o.ActivationKeyName
	}
	if !IsNil(o.ActivationKeyId) {
		toSerialize["activationKeyId"] = o.ActivationKeyId
	}
	return toSerialize, nil
}

type NullableConsumerActivationKeyDTO struct {
	value *ConsumerActivationKeyDTO
	isSet bool
}

func (v NullableConsumerActivationKeyDTO) Get() *ConsumerActivationKeyDTO {
	return v.value
}

func (v *NullableConsumerActivationKeyDTO) Set(val *ConsumerActivationKeyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerActivationKeyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerActivationKeyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerActivationKeyDTO(val *ConsumerActivationKeyDTO) *NullableConsumerActivationKeyDTO {
	return &NullableConsumerActivationKeyDTO{value: val, isSet: true}
}

func (v NullableConsumerActivationKeyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerActivationKeyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


