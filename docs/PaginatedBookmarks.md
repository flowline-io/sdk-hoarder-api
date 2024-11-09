# PaginatedBookmarks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmarks** | [**[]Bookmark**](Bookmark.md) |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewPaginatedBookmarks

`func NewPaginatedBookmarks(bookmarks []Bookmark, nextCursor NullableString, ) *PaginatedBookmarks`

NewPaginatedBookmarks instantiates a new PaginatedBookmarks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBookmarksWithDefaults

`func NewPaginatedBookmarksWithDefaults() *PaginatedBookmarks`

NewPaginatedBookmarksWithDefaults instantiates a new PaginatedBookmarks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarks

`func (o *PaginatedBookmarks) GetBookmarks() []Bookmark`

GetBookmarks returns the Bookmarks field if non-nil, zero value otherwise.

### GetBookmarksOk

`func (o *PaginatedBookmarks) GetBookmarksOk() (*[]Bookmark, bool)`

GetBookmarksOk returns a tuple with the Bookmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarks

`func (o *PaginatedBookmarks) SetBookmarks(v []Bookmark)`

SetBookmarks sets Bookmarks field to given value.


### GetNextCursor

`func (o *PaginatedBookmarks) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatedBookmarks) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatedBookmarks) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *PaginatedBookmarks) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PaginatedBookmarks) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


