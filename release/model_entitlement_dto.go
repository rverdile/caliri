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

// checks if the EntitlementDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDTO{}

// EntitlementDTO A DTO representation of the Entitlement entity
type EntitlementDTO struct {
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Id *string `json:"id,omitempty"`
	Consumer *NestedConsumerDTO `json:"consumer,omitempty"`
	Pool *PoolDTO `json:"pool,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	Certificates []CertificateDTO `json:"certificates,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewEntitlementDTO instantiates a new EntitlementDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDTO() *EntitlementDTO {
	this := EntitlementDTO{}
	return &this
}

// NewEntitlementDTOWithDefaults instantiates a new EntitlementDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDTOWithDefaults() *EntitlementDTO {
	this := EntitlementDTO{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EntitlementDTO) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EntitlementDTO) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *EntitlementDTO) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *EntitlementDTO) GetUpdated() string {
	if o == nil || IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *EntitlementDTO) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *EntitlementDTO) SetUpdated(v string) {
	o.Updated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementDTO) SetId(v string) {
	o.Id = &v
}

// GetConsumer returns the Consumer field value if set, zero value otherwise.
func (o *EntitlementDTO) GetConsumer() NestedConsumerDTO {
	if o == nil || IsNil(o.Consumer) {
		var ret NestedConsumerDTO
		return ret
	}
	return *o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetConsumerOk() (*NestedConsumerDTO, bool) {
	if o == nil || IsNil(o.Consumer) {
		return nil, false
	}
	return o.Consumer, true
}

// HasConsumer returns a boolean if a field has been set.
func (o *EntitlementDTO) HasConsumer() bool {
	if o != nil && !IsNil(o.Consumer) {
		return true
	}

	return false
}

// SetConsumer gets a reference to the given NestedConsumerDTO and assigns it to the Consumer field.
func (o *EntitlementDTO) SetConsumer(v NestedConsumerDTO) {
	o.Consumer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *EntitlementDTO) GetPool() PoolDTO {
	if o == nil || IsNil(o.Pool) {
		var ret PoolDTO
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetPoolOk() (*PoolDTO, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *EntitlementDTO) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given PoolDTO and assigns it to the Pool field.
func (o *EntitlementDTO) SetPool(v PoolDTO) {
	o.Pool = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *EntitlementDTO) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *EntitlementDTO) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *EntitlementDTO) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *EntitlementDTO) GetCertificates() []CertificateDTO {
	if o == nil || IsNil(o.Certificates) {
		var ret []CertificateDTO
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetCertificatesOk() ([]CertificateDTO, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *EntitlementDTO) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateDTO and assigns it to the Certificates field.
func (o *EntitlementDTO) SetCertificates(v []CertificateDTO) {
	o.Certificates = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *EntitlementDTO) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *EntitlementDTO) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *EntitlementDTO) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *EntitlementDTO) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *EntitlementDTO) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *EntitlementDTO) SetEndDate(v string) {
	o.EndDate = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *EntitlementDTO) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDTO) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *EntitlementDTO) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *EntitlementDTO) SetHref(v string) {
	o.Href = &v
}

func (o EntitlementDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Consumer) {
		toSerialize["consumer"] = o.Consumer
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableEntitlementDTO struct {
	value *EntitlementDTO
	isSet bool
}

func (v NullableEntitlementDTO) Get() *EntitlementDTO {
	return v.value
}

func (v *NullableEntitlementDTO) Set(val *EntitlementDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDTO(val *EntitlementDTO) *NullableEntitlementDTO {
	return &NullableEntitlementDTO{value: val, isSet: true}
}

func (v NullableEntitlementDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


