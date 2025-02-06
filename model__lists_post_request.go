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

// checks if the ListsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsPostRequest{}

// ListsPostRequest struct for ListsPostRequest
type ListsPostRequest struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Type *string `json:"type,omitempty"`
	Query *string `json:"query,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
}

type _ListsPostRequest ListsPostRequest

// NewListsPostRequest instantiates a new ListsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsPostRequest(name string, icon string) *ListsPostRequest {
	this := ListsPostRequest{}
	this.Name = name
	this.Icon = icon
	var type_ string = "manual"
	this.Type = &type_
	return &this
}

// NewListsPostRequestWithDefaults instantiates a new ListsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsPostRequestWithDefaults() *ListsPostRequest {
	this := ListsPostRequest{}
	var type_ string = "manual"
	this.Type = &type_
	return &this
}

// GetName returns the Name field value
func (o *ListsPostRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListsPostRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListsPostRequest) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value
func (o *ListsPostRequest) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *ListsPostRequest) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *ListsPostRequest) SetIcon(v string) {
	o.Icon = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListsPostRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsPostRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ListsPostRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ListsPostRequest) SetType(v string) {
	o.Type = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ListsPostRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsPostRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ListsPostRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ListsPostRequest) SetQuery(v string) {
	o.Query = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListsPostRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListsPostRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ListsPostRequest) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ListsPostRequest) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ListsPostRequest) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ListsPostRequest) UnsetParentId() {
	o.ParentId.Unset()
}

func (o ListsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["icon"] = o.Icon
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	return toSerialize, nil
}

func (o *ListsPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"icon",
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

	varListsPostRequest := _ListsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListsPostRequest)

	if err != nil {
		return err
	}

	*o = ListsPostRequest(varListsPostRequest)

	return err
}

type NullableListsPostRequest struct {
	value *ListsPostRequest
	isSet bool
}

func (v NullableListsPostRequest) Get() *ListsPostRequest {
	return v.value
}

func (v *NullableListsPostRequest) Set(val *ListsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsPostRequest(val *ListsPostRequest) *NullableListsPostRequest {
	return &NullableListsPostRequest{value: val, isSet: true}
}

func (v NullableListsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


