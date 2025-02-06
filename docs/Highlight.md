# Highlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookmarkId** | **string** |  | 
**StartOffset** | **float32** |  | 
**EndOffset** | **float32** |  | 
**Color** | Pointer to **string** |  | [optional] [default to "yellow"]
**Text** | **NullableString** |  | 
**Note** | **NullableString** |  | 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewHighlight

`func NewHighlight(bookmarkId string, startOffset float32, endOffset float32, text NullableString, note NullableString, id string, userId string, createdAt string, ) *Highlight`

NewHighlight instantiates a new Highlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighlightWithDefaults

`func NewHighlightWithDefaults() *Highlight`

NewHighlightWithDefaults instantiates a new Highlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarkId

`func (o *Highlight) GetBookmarkId() string`

GetBookmarkId returns the BookmarkId field if non-nil, zero value otherwise.

### GetBookmarkIdOk

`func (o *Highlight) GetBookmarkIdOk() (*string, bool)`

GetBookmarkIdOk returns a tuple with the BookmarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkId

`func (o *Highlight) SetBookmarkId(v string)`

SetBookmarkId sets BookmarkId field to given value.


### GetStartOffset

`func (o *Highlight) GetStartOffset() float32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *Highlight) GetStartOffsetOk() (*float32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *Highlight) SetStartOffset(v float32)`

SetStartOffset sets StartOffset field to given value.


### GetEndOffset

`func (o *Highlight) GetEndOffset() float32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *Highlight) GetEndOffsetOk() (*float32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *Highlight) SetEndOffset(v float32)`

SetEndOffset sets EndOffset field to given value.


### GetColor

`func (o *Highlight) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Highlight) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Highlight) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Highlight) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetText

`func (o *Highlight) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Highlight) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Highlight) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *Highlight) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *Highlight) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetNote

`func (o *Highlight) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Highlight) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Highlight) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *Highlight) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *Highlight) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetId

`func (o *Highlight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Highlight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Highlight) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *Highlight) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Highlight) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Highlight) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCreatedAt

`func (o *Highlight) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Highlight) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Highlight) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


