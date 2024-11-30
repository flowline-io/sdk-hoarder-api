# Bookmark

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
**Tags** | [**[]BookmarkTagsInner**](BookmarkTagsInner.md) |  | 
**Content** | [**BookmarkContent**](BookmarkContent.md) |  | 
**Assets** | [**[]BookmarkAssetsInner**](BookmarkAssetsInner.md) |  | 

## Methods

### NewBookmark

`func NewBookmark(id string, createdAt string, archived bool, favourited bool, taggingStatus NullableString, tags []BookmarkTagsInner, content BookmarkContent, assets []BookmarkAssetsInner, ) *Bookmark`

NewBookmark instantiates a new Bookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkWithDefaults

`func NewBookmarkWithDefaults() *Bookmark`

NewBookmarkWithDefaults instantiates a new Bookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookmark) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Bookmark) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bookmark) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bookmark) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetTitle

`func (o *Bookmark) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Bookmark) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Bookmark) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Bookmark) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Bookmark) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Bookmark) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetArchived

`func (o *Bookmark) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Bookmark) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Bookmark) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetFavourited

`func (o *Bookmark) GetFavourited() bool`

GetFavourited returns the Favourited field if non-nil, zero value otherwise.

### GetFavouritedOk

`func (o *Bookmark) GetFavouritedOk() (*bool, bool)`

GetFavouritedOk returns a tuple with the Favourited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourited

`func (o *Bookmark) SetFavourited(v bool)`

SetFavourited sets Favourited field to given value.


### GetTaggingStatus

`func (o *Bookmark) GetTaggingStatus() string`

GetTaggingStatus returns the TaggingStatus field if non-nil, zero value otherwise.

### GetTaggingStatusOk

`func (o *Bookmark) GetTaggingStatusOk() (*string, bool)`

GetTaggingStatusOk returns a tuple with the TaggingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggingStatus

`func (o *Bookmark) SetTaggingStatus(v string)`

SetTaggingStatus sets TaggingStatus field to given value.


### SetTaggingStatusNil

`func (o *Bookmark) SetTaggingStatusNil(b bool)`

 SetTaggingStatusNil sets the value for TaggingStatus to be an explicit nil

### UnsetTaggingStatus
`func (o *Bookmark) UnsetTaggingStatus()`

UnsetTaggingStatus ensures that no value is present for TaggingStatus, not even an explicit nil
### GetNote

`func (o *Bookmark) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Bookmark) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Bookmark) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Bookmark) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *Bookmark) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *Bookmark) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetSummary

`func (o *Bookmark) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Bookmark) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Bookmark) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Bookmark) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *Bookmark) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *Bookmark) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetTags

`func (o *Bookmark) GetTags() []BookmarkTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Bookmark) GetTagsOk() (*[]BookmarkTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Bookmark) SetTags(v []BookmarkTagsInner)`

SetTags sets Tags field to given value.


### GetContent

`func (o *Bookmark) GetContent() BookmarkContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Bookmark) GetContentOk() (*BookmarkContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Bookmark) SetContent(v BookmarkContent)`

SetContent sets Content field to given value.


### GetAssets

`func (o *Bookmark) GetAssets() []BookmarkAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Bookmark) GetAssetsOk() (*[]BookmarkAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Bookmark) SetAssets(v []BookmarkAssetsInner)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


