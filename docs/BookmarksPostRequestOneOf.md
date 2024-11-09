# BookmarksPostRequestOneOf

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

## Methods

### NewBookmarksPostRequestOneOf

`func NewBookmarksPostRequestOneOf(type_ string, url string, ) *BookmarksPostRequestOneOf`

NewBookmarksPostRequestOneOf instantiates a new BookmarksPostRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarksPostRequestOneOfWithDefaults

`func NewBookmarksPostRequestOneOfWithDefaults() *BookmarksPostRequestOneOf`

NewBookmarksPostRequestOneOfWithDefaults instantiates a new BookmarksPostRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BookmarksPostRequestOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BookmarksPostRequestOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BookmarksPostRequestOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *BookmarksPostRequestOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BookmarksPostRequestOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BookmarksPostRequestOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *BookmarksPostRequestOneOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarksPostRequestOneOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarksPostRequestOneOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarksPostRequestOneOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BookmarksPostRequestOneOf) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BookmarksPostRequestOneOf) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BookmarksPostRequestOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BookmarksPostRequestOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BookmarksPostRequestOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BookmarksPostRequestOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BookmarksPostRequestOneOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BookmarksPostRequestOneOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageUrl

`func (o *BookmarksPostRequestOneOf) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BookmarksPostRequestOneOf) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BookmarksPostRequestOneOf) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BookmarksPostRequestOneOf) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *BookmarksPostRequestOneOf) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *BookmarksPostRequestOneOf) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetImageAssetId

`func (o *BookmarksPostRequestOneOf) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *BookmarksPostRequestOneOf) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *BookmarksPostRequestOneOf) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *BookmarksPostRequestOneOf) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *BookmarksPostRequestOneOf) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *BookmarksPostRequestOneOf) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetScreenshotAssetId

`func (o *BookmarksPostRequestOneOf) GetScreenshotAssetId() string`

GetScreenshotAssetId returns the ScreenshotAssetId field if non-nil, zero value otherwise.

### GetScreenshotAssetIdOk

`func (o *BookmarksPostRequestOneOf) GetScreenshotAssetIdOk() (*string, bool)`

GetScreenshotAssetIdOk returns a tuple with the ScreenshotAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotAssetId

`func (o *BookmarksPostRequestOneOf) SetScreenshotAssetId(v string)`

SetScreenshotAssetId sets ScreenshotAssetId field to given value.

### HasScreenshotAssetId

`func (o *BookmarksPostRequestOneOf) HasScreenshotAssetId() bool`

HasScreenshotAssetId returns a boolean if a field has been set.

### SetScreenshotAssetIdNil

`func (o *BookmarksPostRequestOneOf) SetScreenshotAssetIdNil(b bool)`

 SetScreenshotAssetIdNil sets the value for ScreenshotAssetId to be an explicit nil

### UnsetScreenshotAssetId
`func (o *BookmarksPostRequestOneOf) UnsetScreenshotAssetId()`

UnsetScreenshotAssetId ensures that no value is present for ScreenshotAssetId, not even an explicit nil
### GetFullPageArchiveAssetId

`func (o *BookmarksPostRequestOneOf) GetFullPageArchiveAssetId() string`

GetFullPageArchiveAssetId returns the FullPageArchiveAssetId field if non-nil, zero value otherwise.

### GetFullPageArchiveAssetIdOk

`func (o *BookmarksPostRequestOneOf) GetFullPageArchiveAssetIdOk() (*string, bool)`

GetFullPageArchiveAssetIdOk returns a tuple with the FullPageArchiveAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPageArchiveAssetId

`func (o *BookmarksPostRequestOneOf) SetFullPageArchiveAssetId(v string)`

SetFullPageArchiveAssetId sets FullPageArchiveAssetId field to given value.

### HasFullPageArchiveAssetId

`func (o *BookmarksPostRequestOneOf) HasFullPageArchiveAssetId() bool`

HasFullPageArchiveAssetId returns a boolean if a field has been set.

### SetFullPageArchiveAssetIdNil

`func (o *BookmarksPostRequestOneOf) SetFullPageArchiveAssetIdNil(b bool)`

 SetFullPageArchiveAssetIdNil sets the value for FullPageArchiveAssetId to be an explicit nil

### UnsetFullPageArchiveAssetId
`func (o *BookmarksPostRequestOneOf) UnsetFullPageArchiveAssetId()`

UnsetFullPageArchiveAssetId ensures that no value is present for FullPageArchiveAssetId, not even an explicit nil
### GetVideoAssetId

`func (o *BookmarksPostRequestOneOf) GetVideoAssetId() string`

GetVideoAssetId returns the VideoAssetId field if non-nil, zero value otherwise.

### GetVideoAssetIdOk

`func (o *BookmarksPostRequestOneOf) GetVideoAssetIdOk() (*string, bool)`

GetVideoAssetIdOk returns a tuple with the VideoAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoAssetId

`func (o *BookmarksPostRequestOneOf) SetVideoAssetId(v string)`

SetVideoAssetId sets VideoAssetId field to given value.

### HasVideoAssetId

`func (o *BookmarksPostRequestOneOf) HasVideoAssetId() bool`

HasVideoAssetId returns a boolean if a field has been set.

### SetVideoAssetIdNil

`func (o *BookmarksPostRequestOneOf) SetVideoAssetIdNil(b bool)`

 SetVideoAssetIdNil sets the value for VideoAssetId to be an explicit nil

### UnsetVideoAssetId
`func (o *BookmarksPostRequestOneOf) UnsetVideoAssetId()`

UnsetVideoAssetId ensures that no value is present for VideoAssetId, not even an explicit nil
### GetFavicon

`func (o *BookmarksPostRequestOneOf) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BookmarksPostRequestOneOf) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BookmarksPostRequestOneOf) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BookmarksPostRequestOneOf) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### SetFaviconNil

`func (o *BookmarksPostRequestOneOf) SetFaviconNil(b bool)`

 SetFaviconNil sets the value for Favicon to be an explicit nil

### UnsetFavicon
`func (o *BookmarksPostRequestOneOf) UnsetFavicon()`

UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil
### GetHtmlContent

`func (o *BookmarksPostRequestOneOf) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *BookmarksPostRequestOneOf) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *BookmarksPostRequestOneOf) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *BookmarksPostRequestOneOf) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### SetHtmlContentNil

`func (o *BookmarksPostRequestOneOf) SetHtmlContentNil(b bool)`

 SetHtmlContentNil sets the value for HtmlContent to be an explicit nil

### UnsetHtmlContent
`func (o *BookmarksPostRequestOneOf) UnsetHtmlContent()`

UnsetHtmlContent ensures that no value is present for HtmlContent, not even an explicit nil
### GetCrawledAt

`func (o *BookmarksPostRequestOneOf) GetCrawledAt() string`

GetCrawledAt returns the CrawledAt field if non-nil, zero value otherwise.

### GetCrawledAtOk

`func (o *BookmarksPostRequestOneOf) GetCrawledAtOk() (*string, bool)`

GetCrawledAtOk returns a tuple with the CrawledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawledAt

`func (o *BookmarksPostRequestOneOf) SetCrawledAt(v string)`

SetCrawledAt sets CrawledAt field to given value.

### HasCrawledAt

`func (o *BookmarksPostRequestOneOf) HasCrawledAt() bool`

HasCrawledAt returns a boolean if a field has been set.

### SetCrawledAtNil

`func (o *BookmarksPostRequestOneOf) SetCrawledAtNil(b bool)`

 SetCrawledAtNil sets the value for CrawledAt to be an explicit nil

### UnsetCrawledAt
`func (o *BookmarksPostRequestOneOf) UnsetCrawledAt()`

UnsetCrawledAt ensures that no value is present for CrawledAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


