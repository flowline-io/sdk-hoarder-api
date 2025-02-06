# PaginatedHighlights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highlights** | [**[]Highlight**](Highlight.md) |  | 
**NextCursor** | **NullableString** |  | 

## Methods

### NewPaginatedHighlights

`func NewPaginatedHighlights(highlights []Highlight, nextCursor NullableString, ) *PaginatedHighlights`

NewPaginatedHighlights instantiates a new PaginatedHighlights object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedHighlightsWithDefaults

`func NewPaginatedHighlightsWithDefaults() *PaginatedHighlights`

NewPaginatedHighlightsWithDefaults instantiates a new PaginatedHighlights object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlights

`func (o *PaginatedHighlights) GetHighlights() []Highlight`

GetHighlights returns the Highlights field if non-nil, zero value otherwise.

### GetHighlightsOk

`func (o *PaginatedHighlights) GetHighlightsOk() (*[]Highlight, bool)`

GetHighlightsOk returns a tuple with the Highlights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlights

`func (o *PaginatedHighlights) SetHighlights(v []Highlight)`

SetHighlights sets Highlights field to given value.


### GetNextCursor

`func (o *PaginatedHighlights) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginatedHighlights) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginatedHighlights) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *PaginatedHighlights) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *PaginatedHighlights) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


