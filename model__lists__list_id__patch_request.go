/*
Hoarder API

The API for the Hoarder app

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListsListIdPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsListIdPatchRequest{}

// ListsListIdPatchRequest struct for ListsListIdPatchRequest
type ListsListIdPatchRequest struct {
	Name *string `json:"name,omitempty"`
	Icon *string `json:"icon,omitempty"`
	ParentId NullableString `json:"parentId,omitempty"`
	Query *string `json:"query,omitempty"`
}

// NewListsListIdPatchRequest instantiates a new ListsListIdPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsListIdPatchRequest() *ListsListIdPatchRequest {
	this := ListsListIdPatchRequest{}
	return &this
}

// NewListsListIdPatchRequestWithDefaults instantiates a new ListsListIdPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsListIdPatchRequestWithDefaults() *ListsListIdPatchRequest {
	this := ListsListIdPatchRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListsListIdPatchRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsListIdPatchRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListsListIdPatchRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListsListIdPatchRequest) SetName(v string) {
	o.Name = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ListsListIdPatchRequest) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsListIdPatchRequest) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ListsListIdPatchRequest) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ListsListIdPatchRequest) SetIcon(v string) {
	o.Icon = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListsListIdPatchRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListsListIdPatchRequest) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ListsListIdPatchRequest) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *ListsListIdPatchRequest) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *ListsListIdPatchRequest) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *ListsListIdPatchRequest) UnsetParentId() {
	o.ParentId.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ListsListIdPatchRequest) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsListIdPatchRequest) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ListsListIdPatchRequest) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ListsListIdPatchRequest) SetQuery(v string) {
	o.Query = &v
}

func (o ListsListIdPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsListIdPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableListsListIdPatchRequest struct {
	value *ListsListIdPatchRequest
	isSet bool
}

func (v NullableListsListIdPatchRequest) Get() *ListsListIdPatchRequest {
	return v.value
}

func (v *NullableListsListIdPatchRequest) Set(val *ListsListIdPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListsListIdPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListsListIdPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsListIdPatchRequest(val *ListsListIdPatchRequest) *NullableListsListIdPatchRequest {
	return &NullableListsListIdPatchRequest{value: val, isSet: true}
}

func (v NullableListsListIdPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsListIdPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


