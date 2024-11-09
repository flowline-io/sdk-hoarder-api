/*
Hoarder API

The API for the Hoarder app

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BookmarksBookmarkIdTagsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarksBookmarkIdTagsPost200Response{}

// BookmarksBookmarkIdTagsPost200Response struct for BookmarksBookmarkIdTagsPost200Response
type BookmarksBookmarkIdTagsPost200Response struct {
	Attached []string `json:"attached"`
}

type _BookmarksBookmarkIdTagsPost200Response BookmarksBookmarkIdTagsPost200Response

// NewBookmarksBookmarkIdTagsPost200Response instantiates a new BookmarksBookmarkIdTagsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarksBookmarkIdTagsPost200Response(attached []string) *BookmarksBookmarkIdTagsPost200Response {
	this := BookmarksBookmarkIdTagsPost200Response{}
	this.Attached = attached
	return &this
}

// NewBookmarksBookmarkIdTagsPost200ResponseWithDefaults instantiates a new BookmarksBookmarkIdTagsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarksBookmarkIdTagsPost200ResponseWithDefaults() *BookmarksBookmarkIdTagsPost200Response {
	this := BookmarksBookmarkIdTagsPost200Response{}
	return &this
}

// GetAttached returns the Attached field value
func (o *BookmarksBookmarkIdTagsPost200Response) GetAttached() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Attached
}

// GetAttachedOk returns a tuple with the Attached field value
// and a boolean to check if the value has been set.
func (o *BookmarksBookmarkIdTagsPost200Response) GetAttachedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attached, true
}

// SetAttached sets field value
func (o *BookmarksBookmarkIdTagsPost200Response) SetAttached(v []string) {
	o.Attached = v
}

func (o BookmarksBookmarkIdTagsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarksBookmarkIdTagsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attached"] = o.Attached
	return toSerialize, nil
}

func (o *BookmarksBookmarkIdTagsPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attached",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBookmarksBookmarkIdTagsPost200Response := _BookmarksBookmarkIdTagsPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarksBookmarkIdTagsPost200Response)

	if err != nil {
		return err
	}

	*o = BookmarksBookmarkIdTagsPost200Response(varBookmarksBookmarkIdTagsPost200Response)

	return err
}

type NullableBookmarksBookmarkIdTagsPost200Response struct {
	value *BookmarksBookmarkIdTagsPost200Response
	isSet bool
}

func (v NullableBookmarksBookmarkIdTagsPost200Response) Get() *BookmarksBookmarkIdTagsPost200Response {
	return v.value
}

func (v *NullableBookmarksBookmarkIdTagsPost200Response) Set(val *BookmarksBookmarkIdTagsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarksBookmarkIdTagsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarksBookmarkIdTagsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarksBookmarkIdTagsPost200Response(val *BookmarksBookmarkIdTagsPost200Response) *NullableBookmarksBookmarkIdTagsPost200Response {
	return &NullableBookmarksBookmarkIdTagsPost200Response{value: val, isSet: true}
}

func (v NullableBookmarksBookmarkIdTagsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarksBookmarkIdTagsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

