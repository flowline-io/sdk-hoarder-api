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

// checks if the BookmarkTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkTagsInner{}

// BookmarkTagsInner struct for BookmarkTagsInner
type BookmarkTagsInner struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AttachedBy string `json:"attachedBy"`
}

type _BookmarkTagsInner BookmarkTagsInner

// NewBookmarkTagsInner instantiates a new BookmarkTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkTagsInner(id string, name string, attachedBy string) *BookmarkTagsInner {
	this := BookmarkTagsInner{}
	this.Id = id
	this.Name = name
	this.AttachedBy = attachedBy
	return &this
}

// NewBookmarkTagsInnerWithDefaults instantiates a new BookmarkTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkTagsInnerWithDefaults() *BookmarkTagsInner {
	this := BookmarkTagsInner{}
	return &this
}

// GetId returns the Id field value
func (o *BookmarkTagsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BookmarkTagsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BookmarkTagsInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BookmarkTagsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BookmarkTagsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BookmarkTagsInner) SetName(v string) {
	o.Name = v
}

// GetAttachedBy returns the AttachedBy field value
func (o *BookmarkTagsInner) GetAttachedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttachedBy
}

// GetAttachedByOk returns a tuple with the AttachedBy field value
// and a boolean to check if the value has been set.
func (o *BookmarkTagsInner) GetAttachedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachedBy, true
}

// SetAttachedBy sets field value
func (o *BookmarkTagsInner) SetAttachedBy(v string) {
	o.AttachedBy = v
}

func (o BookmarkTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["attachedBy"] = o.AttachedBy
	return toSerialize, nil
}

func (o *BookmarkTagsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"attachedBy",
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

	varBookmarkTagsInner := _BookmarkTagsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarkTagsInner)

	if err != nil {
		return err
	}

	*o = BookmarkTagsInner(varBookmarkTagsInner)

	return err
}

type NullableBookmarkTagsInner struct {
	value *BookmarkTagsInner
	isSet bool
}

func (v NullableBookmarkTagsInner) Get() *BookmarkTagsInner {
	return v.value
}

func (v *NullableBookmarkTagsInner) Set(val *BookmarkTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkTagsInner(val *BookmarkTagsInner) *NullableBookmarkTagsInner {
	return &NullableBookmarkTagsInner{value: val, isSet: true}
}

func (v NullableBookmarkTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
