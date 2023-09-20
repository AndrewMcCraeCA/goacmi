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

// checks if the WorkImagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkImagesInner{}

// WorkImagesInner struct for WorkImagesInner
type WorkImagesInner struct {
	Id *int32 `json:"id,omitempty"`
	ImageFile *string `json:"image_file,omitempty"`
	ImageFileThumbnail *string `json:"image_file_thumbnail,omitempty"`
	ImageFileXs *string `json:"image_file_xs,omitempty"`
	ImageFileS *string `json:"image_file_s,omitempty"`
	ImageFileM *string `json:"image_file_m,omitempty"`
	ImageFileL *string `json:"image_file_l,omitempty"`
	Caption *string `json:"caption,omitempty"`
	CreditLine *string `json:"credit_line,omitempty"`
	AltText *string `json:"alt_text,omitempty"`
	Width *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
	AccessLevel *string `json:"access_level,omitempty"`
}

// NewWorkImagesInner instantiates a new WorkImagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkImagesInner() *WorkImagesInner {
	this := WorkImagesInner{}
	return &this
}

// NewWorkImagesInnerWithDefaults instantiates a new WorkImagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkImagesInnerWithDefaults() *WorkImagesInner {
	this := WorkImagesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkImagesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkImagesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *WorkImagesInner) SetId(v int32) {
	o.Id = &v
}

// GetImageFile returns the ImageFile field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFile() string {
	if o == nil || IsNil(o.ImageFile) {
		var ret string
		return ret
	}
	return *o.ImageFile
}

// GetImageFileOk returns a tuple with the ImageFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFile) {
		return nil, false
	}
	return o.ImageFile, true
}

// HasImageFile returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFile() bool {
	if o != nil && !IsNil(o.ImageFile) {
		return true
	}

	return false
}

// SetImageFile gets a reference to the given string and assigns it to the ImageFile field.
func (o *WorkImagesInner) SetImageFile(v string) {
	o.ImageFile = &v
}

// GetImageFileThumbnail returns the ImageFileThumbnail field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFileThumbnail() string {
	if o == nil || IsNil(o.ImageFileThumbnail) {
		var ret string
		return ret
	}
	return *o.ImageFileThumbnail
}

// GetImageFileThumbnailOk returns a tuple with the ImageFileThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFileThumbnail) {
		return nil, false
	}
	return o.ImageFileThumbnail, true
}

// HasImageFileThumbnail returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFileThumbnail() bool {
	if o != nil && !IsNil(o.ImageFileThumbnail) {
		return true
	}

	return false
}

// SetImageFileThumbnail gets a reference to the given string and assigns it to the ImageFileThumbnail field.
func (o *WorkImagesInner) SetImageFileThumbnail(v string) {
	o.ImageFileThumbnail = &v
}

// GetImageFileXs returns the ImageFileXs field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFileXs() string {
	if o == nil || IsNil(o.ImageFileXs) {
		var ret string
		return ret
	}
	return *o.ImageFileXs
}

// GetImageFileXsOk returns a tuple with the ImageFileXs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileXsOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFileXs) {
		return nil, false
	}
	return o.ImageFileXs, true
}

// HasImageFileXs returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFileXs() bool {
	if o != nil && !IsNil(o.ImageFileXs) {
		return true
	}

	return false
}

// SetImageFileXs gets a reference to the given string and assigns it to the ImageFileXs field.
func (o *WorkImagesInner) SetImageFileXs(v string) {
	o.ImageFileXs = &v
}

// GetImageFileS returns the ImageFileS field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFileS() string {
	if o == nil || IsNil(o.ImageFileS) {
		var ret string
		return ret
	}
	return *o.ImageFileS
}

// GetImageFileSOk returns a tuple with the ImageFileS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileSOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFileS) {
		return nil, false
	}
	return o.ImageFileS, true
}

// HasImageFileS returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFileS() bool {
	if o != nil && !IsNil(o.ImageFileS) {
		return true
	}

	return false
}

// SetImageFileS gets a reference to the given string and assigns it to the ImageFileS field.
func (o *WorkImagesInner) SetImageFileS(v string) {
	o.ImageFileS = &v
}

// GetImageFileM returns the ImageFileM field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFileM() string {
	if o == nil || IsNil(o.ImageFileM) {
		var ret string
		return ret
	}
	return *o.ImageFileM
}

// GetImageFileMOk returns a tuple with the ImageFileM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileMOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFileM) {
		return nil, false
	}
	return o.ImageFileM, true
}

// HasImageFileM returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFileM() bool {
	if o != nil && !IsNil(o.ImageFileM) {
		return true
	}

	return false
}

// SetImageFileM gets a reference to the given string and assigns it to the ImageFileM field.
func (o *WorkImagesInner) SetImageFileM(v string) {
	o.ImageFileM = &v
}

// GetImageFileL returns the ImageFileL field value if set, zero value otherwise.
func (o *WorkImagesInner) GetImageFileL() string {
	if o == nil || IsNil(o.ImageFileL) {
		var ret string
		return ret
	}
	return *o.ImageFileL
}

// GetImageFileLOk returns a tuple with the ImageFileL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetImageFileLOk() (*string, bool) {
	if o == nil || IsNil(o.ImageFileL) {
		return nil, false
	}
	return o.ImageFileL, true
}

// HasImageFileL returns a boolean if a field has been set.
func (o *WorkImagesInner) HasImageFileL() bool {
	if o != nil && !IsNil(o.ImageFileL) {
		return true
	}

	return false
}

// SetImageFileL gets a reference to the given string and assigns it to the ImageFileL field.
func (o *WorkImagesInner) SetImageFileL(v string) {
	o.ImageFileL = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *WorkImagesInner) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *WorkImagesInner) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *WorkImagesInner) SetCaption(v string) {
	o.Caption = &v
}

// GetCreditLine returns the CreditLine field value if set, zero value otherwise.
func (o *WorkImagesInner) GetCreditLine() string {
	if o == nil || IsNil(o.CreditLine) {
		var ret string
		return ret
	}
	return *o.CreditLine
}

// GetCreditLineOk returns a tuple with the CreditLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetCreditLineOk() (*string, bool) {
	if o == nil || IsNil(o.CreditLine) {
		return nil, false
	}
	return o.CreditLine, true
}

// HasCreditLine returns a boolean if a field has been set.
func (o *WorkImagesInner) HasCreditLine() bool {
	if o != nil && !IsNil(o.CreditLine) {
		return true
	}

	return false
}

// SetCreditLine gets a reference to the given string and assigns it to the CreditLine field.
func (o *WorkImagesInner) SetCreditLine(v string) {
	o.CreditLine = &v
}

// GetAltText returns the AltText field value if set, zero value otherwise.
func (o *WorkImagesInner) GetAltText() string {
	if o == nil || IsNil(o.AltText) {
		var ret string
		return ret
	}
	return *o.AltText
}

// GetAltTextOk returns a tuple with the AltText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetAltTextOk() (*string, bool) {
	if o == nil || IsNil(o.AltText) {
		return nil, false
	}
	return o.AltText, true
}

// HasAltText returns a boolean if a field has been set.
func (o *WorkImagesInner) HasAltText() bool {
	if o != nil && !IsNil(o.AltText) {
		return true
	}

	return false
}

// SetAltText gets a reference to the given string and assigns it to the AltText field.
func (o *WorkImagesInner) SetAltText(v string) {
	o.AltText = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *WorkImagesInner) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *WorkImagesInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *WorkImagesInner) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *WorkImagesInner) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *WorkImagesInner) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *WorkImagesInner) SetHeight(v int32) {
	o.Height = &v
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *WorkImagesInner) GetAccessLevel() string {
	if o == nil || IsNil(o.AccessLevel) {
		var ret string
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkImagesInner) GetAccessLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AccessLevel) {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *WorkImagesInner) HasAccessLevel() bool {
	if o != nil && !IsNil(o.AccessLevel) {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given string and assigns it to the AccessLevel field.
func (o *WorkImagesInner) SetAccessLevel(v string) {
	o.AccessLevel = &v
}

func (o WorkImagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkImagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageFile) {
		toSerialize["image_file"] = o.ImageFile
	}
	if !IsNil(o.ImageFileThumbnail) {
		toSerialize["image_file_thumbnail"] = o.ImageFileThumbnail
	}
	if !IsNil(o.ImageFileXs) {
		toSerialize["image_file_xs"] = o.ImageFileXs
	}
	if !IsNil(o.ImageFileS) {
		toSerialize["image_file_s"] = o.ImageFileS
	}
	if !IsNil(o.ImageFileM) {
		toSerialize["image_file_m"] = o.ImageFileM
	}
	if !IsNil(o.ImageFileL) {
		toSerialize["image_file_l"] = o.ImageFileL
	}
	if !IsNil(o.Caption) {
		toSerialize["caption"] = o.Caption
	}
	if !IsNil(o.CreditLine) {
		toSerialize["credit_line"] = o.CreditLine
	}
	if !IsNil(o.AltText) {
		toSerialize["alt_text"] = o.AltText
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.AccessLevel) {
		toSerialize["access_level"] = o.AccessLevel
	}
	return toSerialize, nil
}

type NullableWorkImagesInner struct {
	value *WorkImagesInner
	isSet bool
}

func (v NullableWorkImagesInner) Get() *WorkImagesInner {
	return v.value
}

func (v *NullableWorkImagesInner) Set(val *WorkImagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkImagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkImagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkImagesInner(val *WorkImagesInner) *NullableWorkImagesInner {
	return &NullableWorkImagesInner{value: val, isSet: true}
}

func (v NullableWorkImagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkImagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

