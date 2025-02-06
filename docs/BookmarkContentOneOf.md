# BookmarkContentOneOf

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
**PrecrawledArchiveAssetId** | Pointer to **NullableString** |  | [optional] 
**VideoAssetId** | Pointer to **NullableString** |  | [optional] 
**Favicon** | Pointer to **NullableString** |  | [optional] 
**HtmlContent** | Pointer to **NullableString** |  | [optional] 
**CrawledAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBookmarkContentOneOf

`func NewBookmarkContentOneOf(type_ string, url string, ) *BookmarkContentOneOf`

NewBookmarkContentOneOf instantiates a new BookmarkContentOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkContentOneOfWithDefaults

`func NewBookmarkContentOneOfWithDefaults() *BookmarkContentOneOf`

NewBookmarkContentOneOfWithDefaults instantiates a new BookmarkContentOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BookmarkContentOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookmarkContentOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookmarkContentOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *BookmarkContentOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BookmarkContentOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BookmarkContentOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *BookmarkContentOneOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarkContentOneOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarkContentOneOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarkContentOneOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookmarkContentOneOf) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookmarkContentOneOf) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BookmarkContentOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BookmarkContentOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BookmarkContentOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BookmarkContentOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BookmarkContentOneOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BookmarkContentOneOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *BookmarkContentOneOf) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BookmarkContentOneOf) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BookmarkContentOneOf) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BookmarkContentOneOf) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *BookmarkContentOneOf) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *BookmarkContentOneOf) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageAssetId

`func (o *BookmarkContentOneOf) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *BookmarkContentOneOf) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *BookmarkContentOneOf) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *BookmarkContentOneOf) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *BookmarkContentOneOf) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *BookmarkContentOneOf) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetScreenshotAssetId

`func (o *BookmarkContentOneOf) GetScreenshotAssetId() string`

GetScreenshotAssetId returns the ScreenshotAssetId field if non-nil, zero value otherwise.

### GetScreenshotAssetIdOk

`func (o *BookmarkContentOneOf) GetScreenshotAssetIdOk() (*string, bool)`

GetScreenshotAssetIdOk returns a tuple with the ScreenshotAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotAssetId

`func (o *BookmarkContentOneOf) SetScreenshotAssetId(v string)`

SetScreenshotAssetId sets ScreenshotAssetId field to given value.

### HasScreenshotAssetId

`func (o *BookmarkContentOneOf) HasScreenshotAssetId() bool`

HasScreenshotAssetId returns a boolean if a field has been set.

### SetScreenshotAssetIdNil

`func (o *BookmarkContentOneOf) SetScreenshotAssetIdNil(b bool)`

 SetScreenshotAssetIdNil sets the value for ScreenshotAssetId to be an explicit nil

### UnsetScreenshotAssetId
`func (o *BookmarkContentOneOf) UnsetScreenshotAssetId()`

UnsetScreenshotAssetId ensures that no value is present for ScreenshotAssetId, not even an explicit nil
### GetFullPageArchiveAssetId

`func (o *BookmarkContentOneOf) GetFullPageArchiveAssetId() string`

GetFullPageArchiveAssetId returns the FullPageArchiveAssetId field if non-nil, zero value otherwise.

### GetFullPageArchiveAssetIdOk

`func (o *BookmarkContentOneOf) GetFullPageArchiveAssetIdOk() (*string, bool)`

GetFullPageArchiveAssetIdOk returns a tuple with the FullPageArchiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageArchiveAssetId

`func (o *BookmarkContentOneOf) SetFullPageArchiveAssetId(v string)`

SetFullPageArchiveAssetId sets FullPageArchiveAssetId field to given value.

### HasFullPageArchiveAssetId

`func (o *BookmarkContentOneOf) HasFullPageArchiveAssetId() bool`

HasFullPageArchiveAssetId returns a boolean if a field has been set.

### SetFullPageArchiveAssetIdNil

`func (o *BookmarkContentOneOf) SetFullPageArchiveAssetIdNil(b bool)`

 SetFullPageArchiveAssetIdNil sets the value for FullPageArchiveAssetId to be an explicit nil

### UnsetFullPageArchiveAssetId
`func (o *BookmarkContentOneOf) UnsetFullPageArchiveAssetId()`

UnsetFullPageArchiveAssetId ensures that no value is present for FullPageArchiveAssetId, not even an explicit nil
### GetPrecrawledArchiveAssetId

`func (o *BookmarkContentOneOf) GetPrecrawledArchiveAssetId() string`

GetPrecrawledArchiveAssetId returns the PrecrawledArchiveAssetId field if non-nil, zero value otherwise.

### GetPrecrawledArchiveAssetIdOk

`func (o *BookmarkContentOneOf) GetPrecrawledArchiveAssetIdOk() (*string, bool)`

GetPrecrawledArchiveAssetIdOk returns a tuple with the PrecrawledArchiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecrawledArchiveAssetId

`func (o *BookmarkContentOneOf) SetPrecrawledArchiveAssetId(v string)`

SetPrecrawledArchiveAssetId sets PrecrawledArchiveAssetId field to given value.

### HasPrecrawledArchiveAssetId

`func (o *BookmarkContentOneOf) HasPrecrawledArchiveAssetId() bool`

HasPrecrawledArchiveAssetId returns a boolean if a field has been set.

### SetPrecrawledArchiveAssetIdNil

`func (o *BookmarkContentOneOf) SetPrecrawledArchiveAssetIdNil(b bool)`

 SetPrecrawledArchiveAssetIdNil sets the value for PrecrawledArchiveAssetId to be an explicit nil

### UnsetPrecrawledArchiveAssetId
`func (o *BookmarkContentOneOf) UnsetPrecrawledArchiveAssetId()`

UnsetPrecrawledArchiveAssetId ensures that no value is present for PrecrawledArchiveAssetId, not even an explicit nil
### GetVideoAssetId

`func (o *BookmarkContentOneOf) GetVideoAssetId() string`

GetVideoAssetId returns the VideoAssetId field if non-nil, zero value otherwise.

### GetVideoAssetIdOk

`func (o *BookmarkContentOneOf) GetVideoAssetIdOk() (*string, bool)`

GetVideoAssetIdOk returns a tuple with the VideoAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAssetId

`func (o *BookmarkContentOneOf) SetVideoAssetId(v string)`

SetVideoAssetId sets VideoAssetId field to given value.

### HasVideoAssetId

`func (o *BookmarkContentOneOf) HasVideoAssetId() bool`

HasVideoAssetId returns a boolean if a field has been set.

### SetVideoAssetIdNil

`func (o *BookmarkContentOneOf) SetVideoAssetIdNil(b bool)`

 SetVideoAssetIdNil sets the value for VideoAssetId to be an explicit nil

### UnsetVideoAssetId
`func (o *BookmarkContentOneOf) UnsetVideoAssetId()`

UnsetVideoAssetId ensures that no value is present for VideoAssetId, not even an explicit nil
### GetFavicon

`func (o *BookmarkContentOneOf) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BookmarkContentOneOf) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BookmarkContentOneOf) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BookmarkContentOneOf) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *BookmarkContentOneOf) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *BookmarkContentOneOf) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetHtmlContent

`func (o *BookmarkContentOneOf) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *BookmarkContentOneOf) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *BookmarkContentOneOf) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *BookmarkContentOneOf) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *BookmarkContentOneOf) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *BookmarkContentOneOf) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCrawledAt

`func (o *BookmarkContentOneOf) GetCrawledAt() string`

GetCrawledAt returns the CrawledAt field if non-nil, zero value otherwise.

### GetCrawledAtOk

`func (o *BookmarkContentOneOf) GetCrawledAtOk() (*string, bool)`

GetCrawledAtOk returns a tuple with the CrawledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawledAt

`func (o *BookmarkContentOneOf) SetCrawledAt(v string)`

SetCrawledAt sets CrawledAt field to given value.

### HasCrawledAt

`func (o *BookmarkContentOneOf) HasCrawledAt() bool`

HasCrawledAt returns a boolean if a field has been set.

### SetCrawledAtNil

`func (o *BookmarkContentOneOf) SetCrawledAtNil(b bool)`

 SetCrawledAtNil sets the value for CrawledAt to be an explicit nil

### UnsetCrawledAt
`func (o *BookmarkContentOneOf) UnsetCrawledAt()`

UnsetCrawledAt ensures that no value is present for CrawledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


