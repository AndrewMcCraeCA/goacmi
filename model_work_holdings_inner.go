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

// checks if the WorkHoldingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkHoldingsInner{}

// WorkHoldingsInner struct for WorkHoldingsInner
type WorkHoldingsInner struct {
	Name *string `json:"name,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewWorkHoldingsInner instantiates a new WorkHoldingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkHoldingsInner() *WorkHoldingsInner {
	this := WorkHoldingsInner{}
	return &this
}

// NewWorkHoldingsInnerWithDefaults instantiates a new WorkHoldingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkHoldingsInnerWithDefaults() *WorkHoldingsInner {
	this := WorkHoldingsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkHoldingsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkHoldingsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkHoldingsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkHoldingsInner) SetName(v string) {
	o.Name = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *WorkHoldingsInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkHoldingsInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *WorkHoldingsInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *WorkHoldingsInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkHoldingsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkHoldingsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkHoldingsInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkHoldingsInner) SetDescription(v string) {
	o.Description = &v
}

func (o WorkHoldingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkHoldingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableWorkHoldingsInner struct {
	value *WorkHoldingsInner
	isSet bool
}

func (v NullableWorkHoldingsInner) Get() *WorkHoldingsInner {
	return v.value
}

func (v *NullableWorkHoldingsInner) Set(val *WorkHoldingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkHoldingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkHoldingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkHoldingsInner(val *WorkHoldingsInner) *NullableWorkHoldingsInner {
	return &NullableWorkHoldingsInner{value: val, isSet: true}
}

func (v NullableWorkHoldingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkHoldingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

