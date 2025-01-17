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

// checks if the NestedOwnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedOwnerDTO{}

// NestedOwnerDTO struct for NestedOwnerDTO
type NestedOwnerDTO struct {
	Id *string `json:"id,omitempty"`
	Key *string `json:"key,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Href *string `json:"href,omitempty"`
	ContentAccessMode *string `json:"contentAccessMode,omitempty"`
}

// NewNestedOwnerDTO instantiates a new NestedOwnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedOwnerDTO() *NestedOwnerDTO {
	this := NestedOwnerDTO{}
	return &this
}

// NewNestedOwnerDTOWithDefaults instantiates a new NestedOwnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedOwnerDTOWithDefaults() *NestedOwnerDTO {
	this := NestedOwnerDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NestedOwnerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOwnerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NestedOwnerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NestedOwnerDTO) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *NestedOwnerDTO) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOwnerDTO) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *NestedOwnerDTO) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *NestedOwnerDTO) SetKey(v string) {
	o.Key = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NestedOwnerDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOwnerDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NestedOwnerDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NestedOwnerDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *NestedOwnerDTO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOwnerDTO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *NestedOwnerDTO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *NestedOwnerDTO) SetHref(v string) {
	o.Href = &v
}

// GetContentAccessMode returns the ContentAccessMode field value if set, zero value otherwise.
func (o *NestedOwnerDTO) GetContentAccessMode() string {
	if o == nil || IsNil(o.ContentAccessMode) {
		var ret string
		return ret
	}
	return *o.ContentAccessMode
}

// GetContentAccessModeOk returns a tuple with the ContentAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOwnerDTO) GetContentAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentAccessMode) {
		return nil, false
	}
	return o.ContentAccessMode, true
}

// HasContentAccessMode returns a boolean if a field has been set.
func (o *NestedOwnerDTO) HasContentAccessMode() bool {
	if o != nil && !IsNil(o.ContentAccessMode) {
		return true
	}

	return false
}

// SetContentAccessMode gets a reference to the given string and assigns it to the ContentAccessMode field.
func (o *NestedOwnerDTO) SetContentAccessMode(v string) {
	o.ContentAccessMode = &v
}

func (o NestedOwnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedOwnerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.ContentAccessMode) {
		toSerialize["contentAccessMode"] = o.ContentAccessMode
	}
	return toSerialize, nil
}

type NullableNestedOwnerDTO struct {
	value *NestedOwnerDTO
	isSet bool
}

func (v NullableNestedOwnerDTO) Get() *NestedOwnerDTO {
	return v.value
}

func (v *NullableNestedOwnerDTO) Set(val *NestedOwnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedOwnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedOwnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedOwnerDTO(val *NestedOwnerDTO) *NullableNestedOwnerDTO {
	return &NullableNestedOwnerDTO{value: val, isSet: true}
}

func (v NullableNestedOwnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedOwnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


