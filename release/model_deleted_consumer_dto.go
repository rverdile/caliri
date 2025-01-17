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

// checks if the DeletedConsumerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletedConsumerDTO{}

// DeletedConsumerDTO Represents a deleted consumer
type DeletedConsumerDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	ConsumerUuid *string `json:"consumerUuid,omitempty"`
	ConsumerName *string `json:"consumerName,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerKey *string `json:"ownerKey,omitempty"`
	OwnerDisplayName *string `json:"ownerDisplayName,omitempty"`
	PrincipalName *string `json:"principalName,omitempty"`
}

// NewDeletedConsumerDTO instantiates a new DeletedConsumerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedConsumerDTO() *DeletedConsumerDTO {
	this := DeletedConsumerDTO{}
	return &this
}

// NewDeletedConsumerDTOWithDefaults instantiates a new DeletedConsumerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedConsumerDTOWithDefaults() *DeletedConsumerDTO {
	this := DeletedConsumerDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *DeletedConsumerDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *DeletedConsumerDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeletedConsumerDTO) SetId(v string) {
	o.Id = &v
}

// GetConsumerUuid returns the ConsumerUuid field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetConsumerUuid() string {
	if o == nil || IsNil(o.ConsumerUuid) {
		var ret string
		return ret
	}
	return *o.ConsumerUuid
}

// GetConsumerUuidOk returns a tuple with the ConsumerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetConsumerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerUuid) {
		return nil, false
	}
	return o.ConsumerUuid, true
}

// HasConsumerUuid returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasConsumerUuid() bool {
	if o != nil && !IsNil(o.ConsumerUuid) {
		return true
	}

	return false
}

// SetConsumerUuid gets a reference to the given string and assigns it to the ConsumerUuid field.
func (o *DeletedConsumerDTO) SetConsumerUuid(v string) {
	o.ConsumerUuid = &v
}

// GetConsumerName returns the ConsumerName field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetConsumerName() string {
	if o == nil || IsNil(o.ConsumerName) {
		var ret string
		return ret
	}
	return *o.ConsumerName
}

// GetConsumerNameOk returns a tuple with the ConsumerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetConsumerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumerName) {
		return nil, false
	}
	return o.ConsumerName, true
}

// HasConsumerName returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasConsumerName() bool {
	if o != nil && !IsNil(o.ConsumerName) {
		return true
	}

	return false
}

// SetConsumerName gets a reference to the given string and assigns it to the ConsumerName field.
func (o *DeletedConsumerDTO) SetConsumerName(v string) {
	o.ConsumerName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *DeletedConsumerDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerKey returns the OwnerKey field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetOwnerKey() string {
	if o == nil || IsNil(o.OwnerKey) {
		var ret string
		return ret
	}
	return *o.OwnerKey
}

// GetOwnerKeyOk returns a tuple with the OwnerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetOwnerKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerKey) {
		return nil, false
	}
	return o.OwnerKey, true
}

// HasOwnerKey returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasOwnerKey() bool {
	if o != nil && !IsNil(o.OwnerKey) {
		return true
	}

	return false
}

// SetOwnerKey gets a reference to the given string and assigns it to the OwnerKey field.
func (o *DeletedConsumerDTO) SetOwnerKey(v string) {
	o.OwnerKey = &v
}

// GetOwnerDisplayName returns the OwnerDisplayName field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetOwnerDisplayName() string {
	if o == nil || IsNil(o.OwnerDisplayName) {
		var ret string
		return ret
	}
	return *o.OwnerDisplayName
}

// GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetOwnerDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerDisplayName) {
		return nil, false
	}
	return o.OwnerDisplayName, true
}

// HasOwnerDisplayName returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasOwnerDisplayName() bool {
	if o != nil && !IsNil(o.OwnerDisplayName) {
		return true
	}

	return false
}

// SetOwnerDisplayName gets a reference to the given string and assigns it to the OwnerDisplayName field.
func (o *DeletedConsumerDTO) SetOwnerDisplayName(v string) {
	o.OwnerDisplayName = &v
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise.
func (o *DeletedConsumerDTO) GetPrincipalName() string {
	if o == nil || IsNil(o.PrincipalName) {
		var ret string
		return ret
	}
	return *o.PrincipalName
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletedConsumerDTO) GetPrincipalNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalName) {
		return nil, false
	}
	return o.PrincipalName, true
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *DeletedConsumerDTO) HasPrincipalName() bool {
	if o != nil && !IsNil(o.PrincipalName) {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given string and assigns it to the PrincipalName field.
func (o *DeletedConsumerDTO) SetPrincipalName(v string) {
	o.PrincipalName = &v
}

func (o DeletedConsumerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletedConsumerDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConsumerUuid) {
		toSerialize["consumerUuid"] = o.ConsumerUuid
	}
	if !IsNil(o.ConsumerName) {
		toSerialize["consumerName"] = o.ConsumerName
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerKey) {
		toSerialize["ownerKey"] = o.OwnerKey
	}
	if !IsNil(o.OwnerDisplayName) {
		toSerialize["ownerDisplayName"] = o.OwnerDisplayName
	}
	if !IsNil(o.PrincipalName) {
		toSerialize["principalName"] = o.PrincipalName
	}
	return toSerialize, nil
}

type NullableDeletedConsumerDTO struct {
	value *DeletedConsumerDTO
	isSet bool
}

func (v NullableDeletedConsumerDTO) Get() *DeletedConsumerDTO {
	return v.value
}

func (v *NullableDeletedConsumerDTO) Set(val *DeletedConsumerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedConsumerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedConsumerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedConsumerDTO(val *DeletedConsumerDTO) *NullableDeletedConsumerDTO {
	return &NullableDeletedConsumerDTO{value: val, isSet: true}
}

func (v NullableDeletedConsumerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedConsumerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


