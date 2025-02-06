# HighlightsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | **string** |  | 
**StartOffset** | **float32** |  | 
**EndOffset** | **float32** |  | 
**Color** | Pointer to **string** |  | [optional] [default to "yellow"]
**Text** | **NullableString** |  | 
**Note** | **NullableString** |  | 

## Methods

### NewHighlightsPostRequest

`func NewHighlightsPostRequest(bookmarkId string, startOffset float32, endOffset float32, text NullableString, note NullableString, ) *HighlightsPostRequest`

NewHighlightsPostRequest instantiates a new HighlightsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightsPostRequestWithDefaults

`func NewHighlightsPostRequestWithDefaults() *HighlightsPostRequest`

NewHighlightsPostRequestWithDefaults instantiates a new HighlightsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *HighlightsPostRequest) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *HighlightsPostRequest) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *HighlightsPostRequest) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.


### GetStartOffset

`func (o *HighlightsPostRequest) GetStartOffset() float32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *HighlightsPostRequest) GetStartOffsetOk() (*float32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *HighlightsPostRequest) SetStartOffset(v float32)`

SetStartOffset sets StartOffset field to given value.


### GetEndOffset

`func (o *HighlightsPostRequest) GetEndOffset() float32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *HighlightsPostRequest) GetEndOffsetOk() (*float32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *HighlightsPostRequest) SetEndOffset(v float32)`

SetEndOffset sets EndOffset field to given value.


### GetColor

`func (o *HighlightsPostRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *HighlightsPostRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *HighlightsPostRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *HighlightsPostRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetText

`func (o *HighlightsPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HighlightsPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HighlightsPostRequest) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *HighlightsPostRequest) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *HighlightsPostRequest) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetNote

`func (o *HighlightsPostRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *HighlightsPostRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *HighlightsPostRequest) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *HighlightsPostRequest) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *HighlightsPostRequest) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


