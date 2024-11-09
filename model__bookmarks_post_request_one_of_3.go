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

// checks if the BookmarksPostRequestOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarksPostRequestOneOf3{}

// BookmarksPostRequestOneOf3 struct for BookmarksPostRequestOneOf3
type BookmarksPostRequestOneOf3 struct {
	Type string `json:"type"`
}

type _BookmarksPostRequestOneOf3 BookmarksPostRequestOneOf3

// NewBookmarksPostRequestOneOf3 instantiates a new BookmarksPostRequestOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarksPostRequestOneOf3(type_ string) *BookmarksPostRequestOneOf3 {
	this := BookmarksPostRequestOneOf3{}
	this.Type = type_
	return &this
}

// NewBookmarksPostRequestOneOf3WithDefaults instantiates a new BookmarksPostRequestOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarksPostRequestOneOf3WithDefaults() *BookmarksPostRequestOneOf3 {
	this := BookmarksPostRequestOneOf3{}
	return &this
}

// GetType returns the Type field value
func (o *BookmarksPostRequestOneOf3) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequestOneOf3) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BookmarksPostRequestOneOf3) SetType(v string) {
	o.Type = v
}

func (o BookmarksPostRequestOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarksPostRequestOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *BookmarksPostRequestOneOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varBookmarksPostRequestOneOf3 := _BookmarksPostRequestOneOf3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarksPostRequestOneOf3)

	if err != nil {
		return err
	}

	*o = BookmarksPostRequestOneOf3(varBookmarksPostRequestOneOf3)

	return err
}

type NullableBookmarksPostRequestOneOf3 struct {
	value *BookmarksPostRequestOneOf3
	isSet bool
}

func (v NullableBookmarksPostRequestOneOf3) Get() *BookmarksPostRequestOneOf3 {
	return v.value
}

func (v *NullableBookmarksPostRequestOneOf3) Set(val *BookmarksPostRequestOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarksPostRequestOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarksPostRequestOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarksPostRequestOneOf3(val *BookmarksPostRequestOneOf3) *NullableBookmarksPostRequestOneOf3 {
	return &NullableBookmarksPostRequestOneOf3{value: val, isSet: true}
}

func (v NullableBookmarksPostRequestOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarksPostRequestOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


