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

// checks if the WorkVideosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkVideosInner{}

// WorkVideosInner struct for WorkVideosInner
type WorkVideosInner struct {
	Id *int32 `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Resource *string `json:"resource,omitempty"`
	WebResource *string `json:"web_resource,omitempty"`
	DurationSecs *float32 `json:"duration_secs,omitempty"`
	Subtitles *string `json:"subtitles,omitempty"`
	MasterMetadata *VideoMetadata `json:"master_metadata,omitempty"`
	AccessMetadata *VideoMetadata `json:"access_metadata,omitempty"`
	WebMetadata *VideoMetadata `json:"web_metadata,omitempty"`
	AccessLevel *string `json:"access_level,omitempty"`
	Snapshot *string `json:"snapshot,omitempty"`
	AnimatedSnapshot *string `json:"animated_snapshot,omitempty"`
}

// NewWorkVideosInner instantiates a new WorkVideosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkVideosInner() *WorkVideosInner {
	this := WorkVideosInner{}
	return &this
}

// NewWorkVideosInnerWithDefaults instantiates a new WorkVideosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkVideosInnerWithDefaults() *WorkVideosInner {
	this := WorkVideosInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkVideosInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkVideosInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WorkVideosInner) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkVideosInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkVideosInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkVideosInner) SetTitle(v string) {
	o.Title = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *WorkVideosInner) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *WorkVideosInner) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *WorkVideosInner) SetResource(v string) {
	o.Resource = &v
}

// GetWebResource returns the WebResource field value if set, zero value otherwise.
func (o *WorkVideosInner) GetWebResource() string {
	if o == nil || IsNil(o.WebResource) {
		var ret string
		return ret
	}
	return *o.WebResource
}

// GetWebResourceOk returns a tuple with the WebResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetWebResourceOk() (*string, bool) {
	if o == nil || IsNil(o.WebResource) {
		return nil, false
	}
	return o.WebResource, true
}

// HasWebResource returns a boolean if a field has been set.
func (o *WorkVideosInner) HasWebResource() bool {
	if o != nil && !IsNil(o.WebResource) {
		return true
	}

	return false
}

// SetWebResource gets a reference to the given string and assigns it to the WebResource field.
func (o *WorkVideosInner) SetWebResource(v string) {
	o.WebResource = &v
}

// GetDurationSecs returns the DurationSecs field value if set, zero value otherwise.
func (o *WorkVideosInner) GetDurationSecs() float32 {
	if o == nil || IsNil(o.DurationSecs) {
		var ret float32
		return ret
	}
	return *o.DurationSecs
}

// GetDurationSecsOk returns a tuple with the DurationSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetDurationSecsOk() (*float32, bool) {
	if o == nil || IsNil(o.DurationSecs) {
		return nil, false
	}
	return o.DurationSecs, true
}

// HasDurationSecs returns a boolean if a field has been set.
func (o *WorkVideosInner) HasDurationSecs() bool {
	if o != nil && !IsNil(o.DurationSecs) {
		return true
	}

	return false
}

// SetDurationSecs gets a reference to the given float32 and assigns it to the DurationSecs field.
func (o *WorkVideosInner) SetDurationSecs(v float32) {
	o.DurationSecs = &v
}

// GetSubtitles returns the Subtitles field value if set, zero value otherwise.
func (o *WorkVideosInner) GetSubtitles() string {
	if o == nil || IsNil(o.Subtitles) {
		var ret string
		return ret
	}
	return *o.Subtitles
}

// GetSubtitlesOk returns a tuple with the Subtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetSubtitlesOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitles) {
		return nil, false
	}
	return o.Subtitles, true
}

// HasSubtitles returns a boolean if a field has been set.
func (o *WorkVideosInner) HasSubtitles() bool {
	if o != nil && !IsNil(o.Subtitles) {
		return true
	}

	return false
}

// SetSubtitles gets a reference to the given string and assigns it to the Subtitles field.
func (o *WorkVideosInner) SetSubtitles(v string) {
	o.Subtitles = &v
}

// GetMasterMetadata returns the MasterMetadata field value if set, zero value otherwise.
func (o *WorkVideosInner) GetMasterMetadata() VideoMetadata {
	if o == nil || IsNil(o.MasterMetadata) {
		var ret VideoMetadata
		return ret
	}
	return *o.MasterMetadata
}

// GetMasterMetadataOk returns a tuple with the MasterMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetMasterMetadataOk() (*VideoMetadata, bool) {
	if o == nil || IsNil(o.MasterMetadata) {
		return nil, false
	}
	return o.MasterMetadata, true
}

// HasMasterMetadata returns a boolean if a field has been set.
func (o *WorkVideosInner) HasMasterMetadata() bool {
	if o != nil && !IsNil(o.MasterMetadata) {
		return true
	}

	return false
}

// SetMasterMetadata gets a reference to the given VideoMetadata and assigns it to the MasterMetadata field.
func (o *WorkVideosInner) SetMasterMetadata(v VideoMetadata) {
	o.MasterMetadata = &v
}

// GetAccessMetadata returns the AccessMetadata field value if set, zero value otherwise.
func (o *WorkVideosInner) GetAccessMetadata() VideoMetadata {
	if o == nil || IsNil(o.AccessMetadata) {
		var ret VideoMetadata
		return ret
	}
	return *o.AccessMetadata
}

// GetAccessMetadataOk returns a tuple with the AccessMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetAccessMetadataOk() (*VideoMetadata, bool) {
	if o == nil || IsNil(o.AccessMetadata) {
		return nil, false
	}
	return o.AccessMetadata, true
}

// HasAccessMetadata returns a boolean if a field has been set.
func (o *WorkVideosInner) HasAccessMetadata() bool {
	if o != nil && !IsNil(o.AccessMetadata) {
		return true
	}

	return false
}

// SetAccessMetadata gets a reference to the given VideoMetadata and assigns it to the AccessMetadata field.
func (o *WorkVideosInner) SetAccessMetadata(v VideoMetadata) {
	o.AccessMetadata = &v
}

// GetWebMetadata returns the WebMetadata field value if set, zero value otherwise.
func (o *WorkVideosInner) GetWebMetadata() VideoMetadata {
	if o == nil || IsNil(o.WebMetadata) {
		var ret VideoMetadata
		return ret
	}
	return *o.WebMetadata
}

// GetWebMetadataOk returns a tuple with the WebMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetWebMetadataOk() (*VideoMetadata, bool) {
	if o == nil || IsNil(o.WebMetadata) {
		return nil, false
	}
	return o.WebMetadata, true
}

// HasWebMetadata returns a boolean if a field has been set.
func (o *WorkVideosInner) HasWebMetadata() bool {
	if o != nil && !IsNil(o.WebMetadata) {
		return true
	}

	return false
}

// SetWebMetadata gets a reference to the given VideoMetadata and assigns it to the WebMetadata field.
func (o *WorkVideosInner) SetWebMetadata(v VideoMetadata) {
	o.WebMetadata = &v
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *WorkVideosInner) GetAccessLevel() string {
	if o == nil || IsNil(o.AccessLevel) {
		var ret string
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetAccessLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevel) {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *WorkVideosInner) HasAccessLevel() bool {
	if o != nil && !IsNil(o.AccessLevel) {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given string and assigns it to the AccessLevel field.
func (o *WorkVideosInner) SetAccessLevel(v string) {
	o.AccessLevel = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *WorkVideosInner) GetSnapshot() string {
	if o == nil || IsNil(o.Snapshot) {
		var ret string
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.Snapshot) {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *WorkVideosInner) HasSnapshot() bool {
	if o != nil && !IsNil(o.Snapshot) {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given string and assigns it to the Snapshot field.
func (o *WorkVideosInner) SetSnapshot(v string) {
	o.Snapshot = &v
}

// GetAnimatedSnapshot returns the AnimatedSnapshot field value if set, zero value otherwise.
func (o *WorkVideosInner) GetAnimatedSnapshot() string {
	if o == nil || IsNil(o.AnimatedSnapshot) {
		var ret string
		return ret
	}
	return *o.AnimatedSnapshot
}

// GetAnimatedSnapshotOk returns a tuple with the AnimatedSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkVideosInner) GetAnimatedSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.AnimatedSnapshot) {
		return nil, false
	}
	return o.AnimatedSnapshot, true
}

// HasAnimatedSnapshot returns a boolean if a field has been set.
func (o *WorkVideosInner) HasAnimatedSnapshot() bool {
	if o != nil && !IsNil(o.AnimatedSnapshot) {
		return true
	}

	return false
}

// SetAnimatedSnapshot gets a reference to the given string and assigns it to the AnimatedSnapshot field.
func (o *WorkVideosInner) SetAnimatedSnapshot(v string) {
	o.AnimatedSnapshot = &v
}

func (o WorkVideosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkVideosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.WebResource) {
		toSerialize["web_resource"] = o.WebResource
	}
	if !IsNil(o.DurationSecs) {
		toSerialize["duration_secs"] = o.DurationSecs
	}
	if !IsNil(o.Subtitles) {
		toSerialize["subtitles"] = o.Subtitles
	}
	if !IsNil(o.MasterMetadata) {
		toSerialize["master_metadata"] = o.MasterMetadata
	}
	if !IsNil(o.AccessMetadata) {
		toSerialize["access_metadata"] = o.AccessMetadata
	}
	if !IsNil(o.WebMetadata) {
		toSerialize["web_metadata"] = o.WebMetadata
	}
	if !IsNil(o.AccessLevel) {
		toSerialize["access_level"] = o.AccessLevel
	}
	if !IsNil(o.Snapshot) {
		toSerialize["snapshot"] = o.Snapshot
	}
	if !IsNil(o.AnimatedSnapshot) {
		toSerialize["animated_snapshot"] = o.AnimatedSnapshot
	}
	return toSerialize, nil
}

type NullableWorkVideosInner struct {
	value *WorkVideosInner
	isSet bool
}

func (v NullableWorkVideosInner) Get() *WorkVideosInner {
	return v.value
}

func (v *NullableWorkVideosInner) Set(val *WorkVideosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkVideosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkVideosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkVideosInner(val *WorkVideosInner) *NullableWorkVideosInner {
	return &NullableWorkVideosInner{value: val, isSet: true}
}

func (v NullableWorkVideosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkVideosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

