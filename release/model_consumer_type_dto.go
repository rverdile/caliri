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

// checks if the ConsumerTypeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumerTypeDTO{}

// ConsumerTypeDTO Represents a consumer type used to differentiate various types of consumers
type ConsumerTypeDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Manifest *bool `json:"manifest,omitempty"`
}

// NewConsumerTypeDTO instantiates a new ConsumerTypeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerTypeDTO() *ConsumerTypeDTO {
	this := ConsumerTypeDTO{}
	return &this
}

// NewConsumerTypeDTOWithDefaults instantiates a new ConsumerTypeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerTypeDTOWithDefaults() *ConsumerTypeDTO {
	this := ConsumerTypeDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ConsumerTypeDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerTypeDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ConsumerTypeDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ConsumerTypeDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *ConsumerTypeDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerTypeDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *ConsumerTypeDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *ConsumerTypeDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConsumerTypeDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerTypeDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConsumerTypeDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConsumerTypeDTO) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsumerTypeDTO) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerTypeDTO) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsumerTypeDTO) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ConsumerTypeDTO) SetLabel(v string) {
	o.Label = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *ConsumerTypeDTO) GetManifest() bool {
	if o == nil || IsNil(o.Manifest) {
		var ret bool
		return ret
	}
	return *o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerTypeDTO) GetManifestOk() (*bool, bool) {
	if o == nil || IsNil(o.Manifest) {
		return nil, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *ConsumerTypeDTO) HasManifest() bool {
	if o != nil && !IsNil(o.Manifest) {
		return true
	}

	return false
}

// SetManifest gets a reference to the given bool and assigns it to the Manifest field.
func (o *ConsumerTypeDTO) SetManifest(v bool) {
	o.Manifest = &v
}

func (o ConsumerTypeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumerTypeDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Manifest) {
		toSerialize["manifest"] = o.Manifest
	}
	return toSerialize, nil
}

type NullableConsumerTypeDTO struct {
	value *ConsumerTypeDTO
	isSet bool
}

func (v NullableConsumerTypeDTO) Get() *ConsumerTypeDTO {
	return v.value
}

func (v *NullableConsumerTypeDTO) Set(val *ConsumerTypeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerTypeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerTypeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerTypeDTO(val *ConsumerTypeDTO) *NullableConsumerTypeDTO {
	return &NullableConsumerTypeDTO{value: val, isSet: true}
}

func (v NullableConsumerTypeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerTypeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


