# BookmarkContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ImageAssetId** | Pointer to **NullableString** |  | [optional] 
**ScreenshotAssetId** | Pointer to **NullableString** |  | [optional] 
**FullPageArchiveAssetId** | Pointer to **NullableString** |  | [optional] 
**VideoAssetId** | Pointer to **NullableString** |  | [optional] 
**Favicon** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**CrawledAt** | Pointer to **NullableString** |  | [optional] 
**Text** | **string** |  | 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**AssetType** | **string** |  | 
**AssetId** | **string** |  | 
**FileName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBookmarkContent

`func NewBookmarkContent(type_ string, url string, text string, assetType string, assetId string, ) *BookmarkContent`

NewBookmarkContent instantiates a new BookmarkContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkContentWithDefaults

`func NewBookmarkContentWithDefaults() *BookmarkContent`

NewBookmarkContentWithDefaults instantiates a new BookmarkContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BookmarkContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookmarkContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookmarkContent) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *BookmarkContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BookmarkContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BookmarkContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *BookmarkContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarkContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarkContent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarkContent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookmarkContent) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookmarkContent) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BookmarkContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BookmarkContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BookmarkContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BookmarkContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BookmarkContent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BookmarkContent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *BookmarkContent) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BookmarkContent) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BookmarkContent) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BookmarkContent) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *BookmarkContent) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *BookmarkContent) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageAssetId

`func (o *BookmarkContent) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *BookmarkContent) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *BookmarkContent) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *BookmarkContent) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *BookmarkContent) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *BookmarkContent) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetScreenshotAssetId

`func (o *BookmarkContent) GetScreenshotAssetId() string`

GetScreenshotAssetId returns the ScreenshotAssetId field if non-nil, zero value otherwise.

### GetScreenshotAssetIdOk

`func (o *BookmarkContent) GetScreenshotAssetIdOk() (*string, bool)`

GetScreenshotAssetIdOk returns a tuple with the ScreenshotAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotAssetId

`func (o *BookmarkContent) SetScreenshotAssetId(v string)`

SetScreenshotAssetId sets ScreenshotAssetId field to given value.

### HasScreenshotAssetId

`func (o *BookmarkContent) HasScreenshotAssetId() bool`

HasScreenshotAssetId returns a boolean if a field has been set.

### SetScreenshotAssetIdNil

`func (o *BookmarkContent) SetScreenshotAssetIdNil(b bool)`

 SetScreenshotAssetIdNil sets the value for ScreenshotAssetId to be an explicit nil

### UnsetScreenshotAssetId
`func (o *BookmarkContent) UnsetScreenshotAssetId()`

UnsetScreenshotAssetId ensures that no value is present for ScreenshotAssetId, not even an explicit nil
### GetFullPageArchiveAssetId

`func (o *BookmarkContent) GetFullPageArchiveAssetId() string`

GetFullPageArchiveAssetId returns the FullPageArchiveAssetId field if non-nil, zero value otherwise.

### GetFullPageArchiveAssetIdOk

`func (o *BookmarkContent) GetFullPageArchiveAssetIdOk() (*string, bool)`

GetFullPageArchiveAssetIdOk returns a tuple with the FullPageArchiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageArchiveAssetId

`func (o *BookmarkContent) SetFullPageArchiveAssetId(v string)`

SetFullPageArchiveAssetId sets FullPageArchiveAssetId field to given value.

### HasFullPageArchiveAssetId

`func (o *BookmarkContent) HasFullPageArchiveAssetId() bool`

HasFullPageArchiveAssetId returns a boolean if a field has been set.

### SetFullPageArchiveAssetIdNil

`func (o *BookmarkContent) SetFullPageArchiveAssetIdNil(b bool)`

 SetFullPageArchiveAssetIdNil sets the value for FullPageArchiveAssetId to be an explicit nil

### UnsetFullPageArchiveAssetId
`func (o *BookmarkContent) UnsetFullPageArchiveAssetId()`

UnsetFullPageArchiveAssetId ensures that no value is present for FullPageArchiveAssetId, not even an explicit nil
### GetVideoAssetId

`func (o *BookmarkContent) GetVideoAssetId() string`

GetVideoAssetId returns the VideoAssetId field if non-nil, zero value otherwise.

### GetVideoAssetIdOk

`func (o *BookmarkContent) GetVideoAssetIdOk() (*string, bool)`

GetVideoAssetIdOk returns a tuple with the VideoAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAssetId

`func (o *BookmarkContent) SetVideoAssetId(v string)`

SetVideoAssetId sets VideoAssetId field to given value.

### HasVideoAssetId

`func (o *BookmarkContent) HasVideoAssetId() bool`

HasVideoAssetId returns a boolean if a field has been set.

### SetVideoAssetIdNil

`func (o *BookmarkContent) SetVideoAssetIdNil(b bool)`

 SetVideoAssetIdNil sets the value for VideoAssetId to be an explicit nil

### UnsetVideoAssetId
`func (o *BookmarkContent) UnsetVideoAssetId()`

UnsetVideoAssetId ensures that no value is present for VideoAssetId, not even an explicit nil
### GetFavicon

`func (o *BookmarkContent) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BookmarkContent) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BookmarkContent) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BookmarkContent) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *BookmarkContent) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *BookmarkContent) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetHtmlContent

`func (o *BookmarkContent) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *BookmarkContent) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *BookmarkContent) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *BookmarkContent) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *BookmarkContent) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *BookmarkContent) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCrawledAt

`func (o *BookmarkContent) GetCrawledAt() string`

GetCrawledAt returns the CrawledAt field if non-nil, zero value otherwise.

### GetCrawledAtOk

`func (o *BookmarkContent) GetCrawledAtOk() (*string, bool)`

GetCrawledAtOk returns a tuple with the CrawledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawledAt

`func (o *BookmarkContent) SetCrawledAt(v string)`

SetCrawledAt sets CrawledAt field to given value.

### HasCrawledAt

`func (o *BookmarkContent) HasCrawledAt() bool`

HasCrawledAt returns a boolean if a field has been set.

### SetCrawledAtNil

`func (o *BookmarkContent) SetCrawledAtNil(b bool)`

 SetCrawledAtNil sets the value for CrawledAt to be an explicit nil

### UnsetCrawledAt
`func (o *BookmarkContent) UnsetCrawledAt()`

UnsetCrawledAt ensures that no value is present for CrawledAt, not even an explicit nil
### GetText

`func (o *BookmarkContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BookmarkContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BookmarkContent) SetText(v string)`

SetText sets Text field to given value.


### GetSourceUrl

`func (o *BookmarkContent) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *BookmarkContent) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *BookmarkContent) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *BookmarkContent) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *BookmarkContent) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BookmarkContent) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetAssetType

`func (o *BookmarkContent) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *BookmarkContent) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *BookmarkContent) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetAssetId

`func (o *BookmarkContent) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *BookmarkContent) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *BookmarkContent) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetFileName

`func (o *BookmarkContent) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BookmarkContent) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BookmarkContent) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BookmarkContent) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *BookmarkContent) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *BookmarkContent) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


