# BookmarksBookmarkIdPatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Archived** | **bool** |  | 
**Favourited** | **bool** |  | 
**TaggingStatus** | **NullableString** |  | 
**Note** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBookmarksBookmarkIdPatch200Response

`func NewBookmarksBookmarkIdPatch200Response(id string, createdAt string, archived bool, favourited bool, taggingStatus NullableString, ) *BookmarksBookmarkIdPatch200Response`

NewBookmarksBookmarkIdPatch200Response instantiates a new BookmarksBookmarkIdPatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarksBookmarkIdPatch200ResponseWithDefaults

`func NewBookmarksBookmarkIdPatch200ResponseWithDefaults() *BookmarksBookmarkIdPatch200Response`

NewBookmarksBookmarkIdPatch200ResponseWithDefaults instantiates a new BookmarksBookmarkIdPatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BookmarksBookmarkIdPatch200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BookmarksBookmarkIdPatch200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BookmarksBookmarkIdPatch200Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BookmarksBookmarkIdPatch200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BookmarksBookmarkIdPatch200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BookmarksBookmarkIdPatch200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetTitle

`func (o *BookmarksBookmarkIdPatch200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarksBookmarkIdPatch200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarksBookmarkIdPatch200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarksBookmarkIdPatch200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookmarksBookmarkIdPatch200Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookmarksBookmarkIdPatch200Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetArchived

`func (o *BookmarksBookmarkIdPatch200Response) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BookmarksBookmarkIdPatch200Response) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BookmarksBookmarkIdPatch200Response) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetFavourited

`func (o *BookmarksBookmarkIdPatch200Response) GetFavourited() bool`

GetFavourited returns the Favourited field if non-nil, zero value otherwise.

### GetFavouritedOk

`func (o *BookmarksBookmarkIdPatch200Response) GetFavouritedOk() (*bool, bool)`

GetFavouritedOk returns a tuple with the Favourited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourited

`func (o *BookmarksBookmarkIdPatch200Response) SetFavourited(v bool)`

SetFavourited sets Favourited field to given value.


### GetTaggingStatus

`func (o *BookmarksBookmarkIdPatch200Response) GetTaggingStatus() string`

GetTaggingStatus returns the TaggingStatus field if non-nil, zero value otherwise.

### GetTaggingStatusOk

`func (o *BookmarksBookmarkIdPatch200Response) GetTaggingStatusOk() (*string, bool)`

GetTaggingStatusOk returns a tuple with the TaggingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggingStatus

`func (o *BookmarksBookmarkIdPatch200Response) SetTaggingStatus(v string)`

SetTaggingStatus sets TaggingStatus field to given value.


### SetTaggingStatusNil

`func (o *BookmarksBookmarkIdPatch200Response) SetTaggingStatusNil(b bool)`

 SetTaggingStatusNil sets the value for TaggingStatus to be an explicit nil

### UnsetTaggingStatus
`func (o *BookmarksBookmarkIdPatch200Response) UnsetTaggingStatus()`

UnsetTaggingStatus ensures that no value is present for TaggingStatus, not even an explicit nil
### GetNote

`func (o *BookmarksBookmarkIdPatch200Response) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BookmarksBookmarkIdPatch200Response) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BookmarksBookmarkIdPatch200Response) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BookmarksBookmarkIdPatch200Response) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *BookmarksBookmarkIdPatch200Response) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *BookmarksBookmarkIdPatch200Response) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetSummary

`func (o *BookmarksBookmarkIdPatch200Response) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BookmarksBookmarkIdPatch200Response) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BookmarksBookmarkIdPatch200Response) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BookmarksBookmarkIdPatch200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *BookmarksBookmarkIdPatch200Response) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *BookmarksBookmarkIdPatch200Response) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


