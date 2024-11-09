# BookmarksPostRequest

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
### GetDescription

`func (o *BookmarksPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BookmarksPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BookmarksPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BookmarksPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BookmarksPostRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BookmarksPostRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *BookmarksPostRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BookmarksPostRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BookmarksPostRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BookmarksPostRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *BookmarksPostRequest) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *BookmarksPostRequest) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageAssetId

`func (o *BookmarksPostRequest) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *BookmarksPostRequest) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *BookmarksPostRequest) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *BookmarksPostRequest) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *BookmarksPostRequest) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *BookmarksPostRequest) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetScreenshotAssetId

`func (o *BookmarksPostRequest) GetScreenshotAssetId() string`

GetScreenshotAssetId returns the ScreenshotAssetId field if non-nil, zero value otherwise.

### GetScreenshotAssetIdOk

`func (o *BookmarksPostRequest) GetScreenshotAssetIdOk() (*string, bool)`

GetScreenshotAssetIdOk returns a tuple with the ScreenshotAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotAssetId

`func (o *BookmarksPostRequest) SetScreenshotAssetId(v string)`

SetScreenshotAssetId sets ScreenshotAssetId field to given value.

### HasScreenshotAssetId

`func (o *BookmarksPostRequest) HasScreenshotAssetId() bool`

HasScreenshotAssetId returns a boolean if a field has been set.

### SetScreenshotAssetIdNil

`func (o *BookmarksPostRequest) SetScreenshotAssetIdNil(b bool)`

 SetScreenshotAssetIdNil sets the value for ScreenshotAssetId to be an explicit nil

### UnsetScreenshotAssetId
`func (o *BookmarksPostRequest) UnsetScreenshotAssetId()`

UnsetScreenshotAssetId ensures that no value is present for ScreenshotAssetId, not even an explicit nil
### GetFullPageArchiveAssetId

`func (o *BookmarksPostRequest) GetFullPageArchiveAssetId() string`

GetFullPageArchiveAssetId returns the FullPageArchiveAssetId field if non-nil, zero value otherwise.

### GetFullPageArchiveAssetIdOk

`func (o *BookmarksPostRequest) GetFullPageArchiveAssetIdOk() (*string, bool)`

GetFullPageArchiveAssetIdOk returns a tuple with the FullPageArchiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageArchiveAssetId

`func (o *BookmarksPostRequest) SetFullPageArchiveAssetId(v string)`

SetFullPageArchiveAssetId sets FullPageArchiveAssetId field to given value.

### HasFullPageArchiveAssetId

`func (o *BookmarksPostRequest) HasFullPageArchiveAssetId() bool`

HasFullPageArchiveAssetId returns a boolean if a field has been set.

### SetFullPageArchiveAssetIdNil

`func (o *BookmarksPostRequest) SetFullPageArchiveAssetIdNil(b bool)`

 SetFullPageArchiveAssetIdNil sets the value for FullPageArchiveAssetId to be an explicit nil

### UnsetFullPageArchiveAssetId
`func (o *BookmarksPostRequest) UnsetFullPageArchiveAssetId()`

UnsetFullPageArchiveAssetId ensures that no value is present for FullPageArchiveAssetId, not even an explicit nil
### GetVideoAssetId

`func (o *BookmarksPostRequest) GetVideoAssetId() string`

GetVideoAssetId returns the VideoAssetId field if non-nil, zero value otherwise.

### GetVideoAssetIdOk

`func (o *BookmarksPostRequest) GetVideoAssetIdOk() (*string, bool)`

GetVideoAssetIdOk returns a tuple with the VideoAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAssetId

`func (o *BookmarksPostRequest) SetVideoAssetId(v string)`

SetVideoAssetId sets VideoAssetId field to given value.

### HasVideoAssetId

`func (o *BookmarksPostRequest) HasVideoAssetId() bool`

HasVideoAssetId returns a boolean if a field has been set.

### SetVideoAssetIdNil

`func (o *BookmarksPostRequest) SetVideoAssetIdNil(b bool)`

 SetVideoAssetIdNil sets the value for VideoAssetId to be an explicit nil

### UnsetVideoAssetId
`func (o *BookmarksPostRequest) UnsetVideoAssetId()`

UnsetVideoAssetId ensures that no value is present for VideoAssetId, not even an explicit nil
### GetFavicon

`func (o *BookmarksPostRequest) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BookmarksPostRequest) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BookmarksPostRequest) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BookmarksPostRequest) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *BookmarksPostRequest) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *BookmarksPostRequest) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetHtmlContent

`func (o *BookmarksPostRequest) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *BookmarksPostRequest) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *BookmarksPostRequest) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *BookmarksPostRequest) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *BookmarksPostRequest) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *BookmarksPostRequest) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCrawledAt

`func (o *BookmarksPostRequest) GetCrawledAt() string`

GetCrawledAt returns the CrawledAt field if non-nil, zero value otherwise.

### GetCrawledAtOk

`func (o *BookmarksPostRequest) GetCrawledAtOk() (*string, bool)`

GetCrawledAtOk returns a tuple with the CrawledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawledAt

`func (o *BookmarksPostRequest) SetCrawledAt(v string)`

SetCrawledAt sets CrawledAt field to given value.

### HasCrawledAt

`func (o *BookmarksPostRequest) HasCrawledAt() bool`

HasCrawledAt returns a boolean if a field has been set.

### SetCrawledAtNil

`func (o *BookmarksPostRequest) SetCrawledAtNil(b bool)`

 SetCrawledAtNil sets the value for CrawledAt to be an explicit nil

### UnsetCrawledAt
`func (o *BookmarksPostRequest) UnsetCrawledAt()`

UnsetCrawledAt ensures that no value is present for CrawledAt, not even an explicit nil
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

### SetSourceUrlNil

`func (o *BookmarksPostRequest) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *BookmarksPostRequest) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
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

### SetFileNameNil

`func (o *BookmarksPostRequest) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *BookmarksPostRequest) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


