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

// checks if the ClaimantOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimantOwner{}

// ClaimantOwner Represents an owner claiming the consumers of an anonymous owner
type ClaimantOwner struct {
	ClaimantOwnerKey string `json:"claimant_owner_key"`
}

type _ClaimantOwner ClaimantOwner

// NewClaimantOwner instantiates a new ClaimantOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimantOwner(claimantOwnerKey string) *ClaimantOwner {
	this := ClaimantOwner{}
	this.ClaimantOwnerKey = claimantOwnerKey
	return &this
}

// NewClaimantOwnerWithDefaults instantiates a new ClaimantOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimantOwnerWithDefaults() *ClaimantOwner {
	this := ClaimantOwner{}
	return &this
}

// GetClaimantOwnerKey returns the ClaimantOwnerKey field value
func (o *ClaimantOwner) GetClaimantOwnerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimantOwnerKey
}

// GetClaimantOwnerKeyOk returns a tuple with the ClaimantOwnerKey field value
// and a boolean to check if the value has been set.
func (o *ClaimantOwner) GetClaimantOwnerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimantOwnerKey, true
}

// SetClaimantOwnerKey sets field value
func (o *ClaimantOwner) SetClaimantOwnerKey(v string) {
	o.ClaimantOwnerKey = v
}

func (o ClaimantOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimantOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["claimant_owner_key"] = o.ClaimantOwnerKey
	return toSerialize, nil
}

func (o *ClaimantOwner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"claimant_owner_key",
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

	varClaimantOwner := _ClaimantOwner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClaimantOwner)

	if err != nil {
		return err
	}

	*o = ClaimantOwner(varClaimantOwner)

	return err
}

type NullableClaimantOwner struct {
	value *ClaimantOwner
	isSet bool
}

func (v NullableClaimantOwner) Get() *ClaimantOwner {
	return v.value
}

func (v *NullableClaimantOwner) Set(val *ClaimantOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimantOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimantOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimantOwner(val *ClaimantOwner) *NullableClaimantOwner {
	return &NullableClaimantOwner{value: val, isSet: true}
}

func (v NullableClaimantOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimantOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


