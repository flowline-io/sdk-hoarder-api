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

// checks if the BookmarkAssetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkAssetsInner{}

// BookmarkAssetsInner struct for BookmarkAssetsInner
type BookmarkAssetsInner struct {
	Id string `json:"id"`
	AssetType string `json:"assetType"`
}

type _BookmarkAssetsInner BookmarkAssetsInner

// NewBookmarkAssetsInner instantiates a new BookmarkAssetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkAssetsInner(id string, assetType string) *BookmarkAssetsInner {
	this := BookmarkAssetsInner{}
	this.Id = id
	this.AssetType = assetType
	return &this
}

// NewBookmarkAssetsInnerWithDefaults instantiates a new BookmarkAssetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkAssetsInnerWithDefaults() *BookmarkAssetsInner {
	this := BookmarkAssetsInner{}
	return &this
}

// GetId returns the Id field value
func (o *BookmarkAssetsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BookmarkAssetsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BookmarkAssetsInner) SetId(v string) {
	o.Id = v
}

// GetAssetType returns the AssetType field value
func (o *BookmarkAssetsInner) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *BookmarkAssetsInner) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *BookmarkAssetsInner) SetAssetType(v string) {
	o.AssetType = v
}

func (o BookmarkAssetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkAssetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["assetType"] = o.AssetType
	return toSerialize, nil
}

func (o *BookmarkAssetsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"assetType",
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

	varBookmarkAssetsInner := _BookmarkAssetsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarkAssetsInner)

	if err != nil {
		return err
	}

	*o = BookmarkAssetsInner(varBookmarkAssetsInner)

	return err
}

type NullableBookmarkAssetsInner struct {
	value *BookmarkAssetsInner
	isSet bool
}

func (v NullableBookmarkAssetsInner) Get() *BookmarkAssetsInner {
	return v.value
}

func (v *NullableBookmarkAssetsInner) Set(val *BookmarkAssetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkAssetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkAssetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkAssetsInner(val *BookmarkAssetsInner) *NullableBookmarkAssetsInner {
	return &NullableBookmarkAssetsInner{value: val, isSet: true}
}

func (v NullableBookmarkAssetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkAssetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

