/*
ACMI public API

A public API for ACMI's collection data including Films, TV Shows, Videogames and Art.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goacmi

import (
	"encoding/json"
)

// checks if the WorkStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkStats{}

// WorkStats struct for WorkStats
type WorkStats struct {
	TapCount *int32 `json:"tap_count,omitempty"`
}

// NewWorkStats instantiates a new WorkStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkStats() *WorkStats {
	this := WorkStats{}
	return &this
}

// NewWorkStatsWithDefaults instantiates a new WorkStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkStatsWithDefaults() *WorkStats {
	this := WorkStats{}
	return &this
}

// GetTapCount returns the TapCount field value if set, zero value otherwise.
func (o *WorkStats) GetTapCount() int32 {
	if o == nil || IsNil(o.TapCount) {
		var ret int32
		return ret
	}
	return *o.TapCount
}

// GetTapCountOk returns a tuple with the TapCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkStats) GetTapCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TapCount) {
		return nil, false
	}
	return o.TapCount, true
}

// HasTapCount returns a boolean if a field has been set.
func (o *WorkStats) HasTapCount() bool {
	if o != nil && !IsNil(o.TapCount) {
		return true
	}

	return false
}

// SetTapCount gets a reference to the given int32 and assigns it to the TapCount field.
func (o *WorkStats) SetTapCount(v int32) {
	o.TapCount = &v
}

func (o WorkStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TapCount) {
		toSerialize["tap_count"] = o.TapCount
	}
	return toSerialize, nil
}

type NullableWorkStats struct {
	value *WorkStats
	isSet bool
}

func (v NullableWorkStats) Get() *WorkStats {
	return v.value
}

func (v *NullableWorkStats) Set(val *WorkStats) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkStats) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkStats(val *WorkStats) *NullableWorkStats {
	return &NullableWorkStats{value: val, isSet: true}
}

func (v NullableWorkStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


