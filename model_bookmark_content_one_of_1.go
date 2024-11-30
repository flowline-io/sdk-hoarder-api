/*
Hoarder API

The API for the Hoarder app

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BookmarkContentOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkContentOneOf1{}

// BookmarkContentOneOf1 struct for BookmarkContentOneOf1
type BookmarkContentOneOf1 struct {
	Type      string         `json:"type"`
	Text      string         `json:"text"`
	SourceUrl NullableString `json:"sourceUrl,omitempty"`
}

type _BookmarkContentOneOf1 BookmarkContentOneOf1

// NewBookmarkContentOneOf1 instantiates a new BookmarkContentOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkContentOneOf1(type_ string, text string) *BookmarkContentOneOf1 {
	this := BookmarkContentOneOf1{}
	this.Type = type_
	this.Text = text
	return &this
}

// NewBookmarkContentOneOf1WithDefaults instantiates a new BookmarkContentOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkContentOneOf1WithDefaults() *BookmarkContentOneOf1 {
	this := BookmarkContentOneOf1{}
	return &this
}

// GetType returns the Type field value
func (o *BookmarkContentOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BookmarkContentOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BookmarkContentOneOf1) SetType(v string) {
	o.Type = v
}

// GetText returns the Text field value
func (o *BookmarkContentOneOf1) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *BookmarkContentOneOf1) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *BookmarkContentOneOf1) SetText(v string) {
	o.Text = v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookmarkContentOneOf1) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SourceUrl.Get()
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookmarkContentOneOf1) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceUrl.Get(), o.SourceUrl.IsSet()
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *BookmarkContentOneOf1) HasSourceUrl() bool {
	if o != nil && o.SourceUrl.IsSet() {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given NullableString and assigns it to the SourceUrl field.
func (o *BookmarkContentOneOf1) SetSourceUrl(v string) {
	o.SourceUrl.Set(&v)
}

// SetSourceUrlNil sets the value for SourceUrl to be an explicit nil
func (o *BookmarkContentOneOf1) SetSourceUrlNil() {
	o.SourceUrl.Set(nil)
}

// UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
func (o *BookmarkContentOneOf1) UnsetSourceUrl() {
	o.SourceUrl.Unset()
}

func (o BookmarkContentOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkContentOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text"] = o.Text
	if o.SourceUrl.IsSet() {
		toSerialize["sourceUrl"] = o.SourceUrl.Get()
	}
	return toSerialize, nil
}

func (o *BookmarkContentOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"text",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBookmarkContentOneOf1 := _BookmarkContentOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarkContentOneOf1)

	if err != nil {
		return err
	}

	*o = BookmarkContentOneOf1(varBookmarkContentOneOf1)

	return err
}

type NullableBookmarkContentOneOf1 struct {
	value *BookmarkContentOneOf1
	isSet bool
}

func (v NullableBookmarkContentOneOf1) Get() *BookmarkContentOneOf1 {
	return v.value
}

func (v *NullableBookmarkContentOneOf1) Set(val *BookmarkContentOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkContentOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkContentOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkContentOneOf1(val *BookmarkContentOneOf1) *NullableBookmarkContentOneOf1 {
	return &NullableBookmarkContentOneOf1{value: val, isSet: true}
}

func (v NullableBookmarkContentOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkContentOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
