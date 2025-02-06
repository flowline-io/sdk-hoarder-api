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

// checks if the BookmarksPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarksPostRequest{}

// BookmarksPostRequest struct for BookmarksPostRequest
type BookmarksPostRequest struct {
	Title NullableString `json:"title,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	Favourited *bool `json:"favourited,omitempty"`
	Note *string `json:"note,omitempty"`
	Summary *string `json:"summary,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	Type string `json:"type"`
	Url string `json:"url"`
	PrecrawledArchiveId *string `json:"precrawledArchiveId,omitempty"`
	Text string `json:"text"`
	SourceUrl *string `json:"sourceUrl,omitempty"`
	AssetType string `json:"assetType"`
	AssetId string `json:"assetId"`
	FileName *string `json:"fileName,omitempty"`
}

type _BookmarksPostRequest BookmarksPostRequest

// NewBookmarksPostRequest instantiates a new BookmarksPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarksPostRequest(type_ string, url string, text string, assetType string, assetId string) *BookmarksPostRequest {
	this := BookmarksPostRequest{}
	this.Type = type_
	this.Url = url
	this.Text = text
	this.AssetType = assetType
	this.AssetId = assetId
	return &this
}

// NewBookmarksPostRequestWithDefaults instantiates a new BookmarksPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarksPostRequestWithDefaults() *BookmarksPostRequest {
	this := BookmarksPostRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookmarksPostRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookmarksPostRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *BookmarksPostRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *BookmarksPostRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *BookmarksPostRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *BookmarksPostRequest) SetArchived(v bool) {
	o.Archived = &v
}

// GetFavourited returns the Favourited field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetFavourited() bool {
	if o == nil || IsNil(o.Favourited) {
		var ret bool
		return ret
	}
	return *o.Favourited
}

// GetFavouritedOk returns a tuple with the Favourited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetFavouritedOk() (*bool, bool) {
	if o == nil || IsNil(o.Favourited) {
		return nil, false
	}
	return o.Favourited, true
}

// HasFavourited returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasFavourited() bool {
	if o != nil && !IsNil(o.Favourited) {
		return true
	}

	return false
}

// SetFavourited gets a reference to the given bool and assigns it to the Favourited field.
func (o *BookmarksPostRequest) SetFavourited(v bool) {
	o.Favourited = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *BookmarksPostRequest) SetNote(v string) {
	o.Note = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *BookmarksPostRequest) SetSummary(v string) {
	o.Summary = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookmarksPostRequest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookmarksPostRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *BookmarksPostRequest) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *BookmarksPostRequest) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *BookmarksPostRequest) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetType returns the Type field value
func (o *BookmarksPostRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BookmarksPostRequest) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *BookmarksPostRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BookmarksPostRequest) SetUrl(v string) {
	o.Url = v
}

// GetPrecrawledArchiveId returns the PrecrawledArchiveId field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetPrecrawledArchiveId() string {
	if o == nil || IsNil(o.PrecrawledArchiveId) {
		var ret string
		return ret
	}
	return *o.PrecrawledArchiveId
}

// GetPrecrawledArchiveIdOk returns a tuple with the PrecrawledArchiveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetPrecrawledArchiveIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrecrawledArchiveId) {
		return nil, false
	}
	return o.PrecrawledArchiveId, true
}

// HasPrecrawledArchiveId returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasPrecrawledArchiveId() bool {
	if o != nil && !IsNil(o.PrecrawledArchiveId) {
		return true
	}

	return false
}

// SetPrecrawledArchiveId gets a reference to the given string and assigns it to the PrecrawledArchiveId field.
func (o *BookmarksPostRequest) SetPrecrawledArchiveId(v string) {
	o.PrecrawledArchiveId = &v
}

// GetText returns the Text field value
func (o *BookmarksPostRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *BookmarksPostRequest) SetText(v string) {
	o.Text = v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl) {
		var ret string
		return ret
	}
	return *o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetSourceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUrl) {
		return nil, false
	}
	return o.SourceUrl, true
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasSourceUrl() bool {
	if o != nil && !IsNil(o.SourceUrl) {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given string and assigns it to the SourceUrl field.
func (o *BookmarksPostRequest) SetSourceUrl(v string) {
	o.SourceUrl = &v
}

// GetAssetType returns the AssetType field value
func (o *BookmarksPostRequest) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetAssetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *BookmarksPostRequest) SetAssetType(v string) {
	o.AssetType = v
}

// GetAssetId returns the AssetId field value
func (o *BookmarksPostRequest) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *BookmarksPostRequest) SetAssetId(v string) {
	o.AssetId = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *BookmarksPostRequest) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarksPostRequest) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *BookmarksPostRequest) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *BookmarksPostRequest) SetFileName(v string) {
	o.FileName = &v
}

func (o BookmarksPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarksPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Favourited) {
		toSerialize["favourited"] = o.Favourited
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	if !IsNil(o.PrecrawledArchiveId) {
		toSerialize["precrawledArchiveId"] = o.PrecrawledArchiveId
	}
	toSerialize["text"] = o.Text
	if !IsNil(o.SourceUrl) {
		toSerialize["sourceUrl"] = o.SourceUrl
	}
	toSerialize["assetType"] = o.AssetType
	toSerialize["assetId"] = o.AssetId
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	return toSerialize, nil
}

func (o *BookmarksPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"url",
		"text",
		"assetType",
		"assetId",
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

	varBookmarksPostRequest := _BookmarksPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarksPostRequest)

	if err != nil {
		return err
	}

	*o = BookmarksPostRequest(varBookmarksPostRequest)

	return err
}

type NullableBookmarksPostRequest struct {
	value *BookmarksPostRequest
	isSet bool
}

func (v NullableBookmarksPostRequest) Get() *BookmarksPostRequest {
	return v.value
}

func (v *NullableBookmarksPostRequest) Set(val *BookmarksPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarksPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarksPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarksPostRequest(val *BookmarksPostRequest) *NullableBookmarksPostRequest {
	return &NullableBookmarksPostRequest{value: val, isSet: true}
}

func (v NullableBookmarksPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarksPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


