# BookmarksPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Favourited** | Pointer to **bool** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 
**PrecrawledArchiveId** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 
**SourceUrl** | Pointer to **string** |  | [optional] 
**AssetType** | **string** |  | 
**AssetId** | **string** |  | 
**FileName** | Pointer to **string** |  | [optional] 

## Methods

### NewBookmarksPostRequest

`func NewBookmarksPostRequest(type_ string, url string, text string, assetType string, assetId string, ) *BookmarksPostRequest`

NewBookmarksPostRequest instantiates a new BookmarksPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarksPostRequestWithDefaults

`func NewBookmarksPostRequestWithDefaults() *BookmarksPostRequest`

NewBookmarksPostRequestWithDefaults instantiates a new BookmarksPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BookmarksPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarksPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarksPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarksPostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookmarksPostRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookmarksPostRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetArchived

`func (o *BookmarksPostRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *BookmarksPostRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *BookmarksPostRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *BookmarksPostRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetFavourited

`func (o *BookmarksPostRequest) GetFavourited() bool`

GetFavourited returns the Favourited field if non-nil, zero value otherwise.

### GetFavouritedOk

`func (o *BookmarksPostRequest) GetFavouritedOk() (*bool, bool)`

GetFavouritedOk returns a tuple with the Favourited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourited

`func (o *BookmarksPostRequest) SetFavourited(v bool)`

SetFavourited sets Favourited field to given value.

### HasFavourited

`func (o *BookmarksPostRequest) HasFavourited() bool`

HasFavourited returns a boolean if a field has been set.

### GetNote

`func (o *BookmarksPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BookmarksPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BookmarksPostRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BookmarksPostRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSummary

`func (o *BookmarksPostRequest) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BookmarksPostRequest) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BookmarksPostRequest) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BookmarksPostRequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BookmarksPostRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BookmarksPostRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BookmarksPostRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BookmarksPostRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BookmarksPostRequest) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BookmarksPostRequest) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetType

`func (o *BookmarksPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookmarksPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookmarksPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *BookmarksPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BookmarksPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BookmarksPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPrecrawledArchiveId

`func (o *BookmarksPostRequest) GetPrecrawledArchiveId() string`

GetPrecrawledArchiveId returns the PrecrawledArchiveId field if non-nil, zero value otherwise.

### GetPrecrawledArchiveIdOk

`func (o *BookmarksPostRequest) GetPrecrawledArchiveIdOk() (*string, bool)`

GetPrecrawledArchiveIdOk returns a tuple with the PrecrawledArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecrawledArchiveId

`func (o *BookmarksPostRequest) SetPrecrawledArchiveId(v string)`

SetPrecrawledArchiveId sets PrecrawledArchiveId field to given value.

### HasPrecrawledArchiveId

`func (o *BookmarksPostRequest) HasPrecrawledArchiveId() bool`

HasPrecrawledArchiveId returns a boolean if a field has been set.

### GetText

`func (o *BookmarksPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BookmarksPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BookmarksPostRequest) SetText(v string)`

SetText sets Text field to given value.


### GetSourceUrl

`func (o *BookmarksPostRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BookmarksPostRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BookmarksPostRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BookmarksPostRequest) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetAssetType

`func (o *BookmarksPostRequest) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *BookmarksPostRequest) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *BookmarksPostRequest) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetAssetId

`func (o *BookmarksPostRequest) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *BookmarksPostRequest) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *BookmarksPostRequest) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetFileName

`func (o *BookmarksPostRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BookmarksPostRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BookmarksPostRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BookmarksPostRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


