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

// checks if the WorkVideoLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkVideoLinksInner{}

// WorkVideoLinksInner struct for WorkVideoLinksInner
type WorkVideoLinksInner struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewWorkVideoLinksInner instantiates a new WorkVideoLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkVideoLinksInner() *WorkVideoLinksInner {
	this := WorkVideoLinksInner{}
	return &this
}

// NewWorkVideoLinksInnerWithDefaults instantiates a new WorkVideoLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkVideoLinksInnerWithDefaults() *WorkVideoLinksInner {
	this := WorkVideoLinksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkVideoLinksInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideoLinksInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkVideoLinksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WorkVideoLinksInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkVideoLinksInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideoLinksInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkVideoLinksInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkVideoLinksInner) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *WorkVideoLinksInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideoLinksInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *WorkVideoLinksInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *WorkVideoLinksInner) SetUri(v string) {
	o.Uri = &v
}

func (o WorkVideoLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkVideoLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableWorkVideoLinksInner struct {
	value *WorkVideoLinksInner
	isSet bool
}

func (v NullableWorkVideoLinksInner) Get() *WorkVideoLinksInner {
	return v.value
}

func (v *NullableWorkVideoLinksInner) Set(val *WorkVideoLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkVideoLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkVideoLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkVideoLinksInner(val *WorkVideoLinksInner) *NullableWorkVideoLinksInner {
	return &NullableWorkVideoLinksInner{value: val, isSet: true}
}

func (v NullableWorkVideoLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkVideoLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

