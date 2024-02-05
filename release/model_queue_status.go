/*
Candlepin

Candlepin is a subscription management server written in Java. It helps with management of software subscriptions.

API version: 4.4.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package caliri

import (
	"encoding/json"
)

// checks if the QueueStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueStatus{}

// QueueStatus Represents status of the ActiveMQ queues. Used for checking if events are piling up for some reason or being delivered correctly. 
type QueueStatus struct {
	QueueName *string `json:"queueName,omitempty"`
	PendingMessageCount *int64 `json:"pendingMessageCount,omitempty"`
}

// NewQueueStatus instantiates a new QueueStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueStatus() *QueueStatus {
	this := QueueStatus{}
	return &this
}

// NewQueueStatusWithDefaults instantiates a new QueueStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueStatusWithDefaults() *QueueStatus {
	this := QueueStatus{}
	return &this
}

// GetQueueName returns the QueueName field value if set, zero value otherwise.
func (o *QueueStatus) GetQueueName() string {
	if o == nil || IsNil(o.QueueName) {
		var ret string
		return ret
	}
	return *o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatus) GetQueueNameOk() (*string, bool) {
	if o == nil || IsNil(o.QueueName) {
		return nil, false
	}
	return o.QueueName, true
}

// HasQueueName returns a boolean if a field has been set.
func (o *QueueStatus) HasQueueName() bool {
	if o != nil && !IsNil(o.QueueName) {
		return true
	}

	return false
}

// SetQueueName gets a reference to the given string and assigns it to the QueueName field.
func (o *QueueStatus) SetQueueName(v string) {
	o.QueueName = &v
}

// GetPendingMessageCount returns the PendingMessageCount field value if set, zero value otherwise.
func (o *QueueStatus) GetPendingMessageCount() int64 {
	if o == nil || IsNil(o.PendingMessageCount) {
		var ret int64
		return ret
	}
	return *o.PendingMessageCount
}

// GetPendingMessageCountOk returns a tuple with the PendingMessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatus) GetPendingMessageCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PendingMessageCount) {
		return nil, false
	}
	return o.PendingMessageCount, true
}

// HasPendingMessageCount returns a boolean if a field has been set.
func (o *QueueStatus) HasPendingMessageCount() bool {
	if o != nil && !IsNil(o.PendingMessageCount) {
		return true
	}

	return false
}

// SetPendingMessageCount gets a reference to the given int64 and assigns it to the PendingMessageCount field.
func (o *QueueStatus) SetPendingMessageCount(v int64) {
	o.PendingMessageCount = &v
}

func (o QueueStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QueueName) {
		toSerialize["queueName"] = o.QueueName
	}
	if !IsNil(o.PendingMessageCount) {
		toSerialize["pendingMessageCount"] = o.PendingMessageCount
	}
	return toSerialize, nil
}

type NullableQueueStatus struct {
	value *QueueStatus
	isSet bool
}

func (v NullableQueueStatus) Get() *QueueStatus {
	return v.value
}

func (v *NullableQueueStatus) Set(val *QueueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueStatus(val *QueueStatus) *NullableQueueStatus {
	return &NullableQueueStatus{value: val, isSet: true}
}

func (v NullableQueueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

