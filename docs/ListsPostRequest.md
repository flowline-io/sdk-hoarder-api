# ListsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Icon** | **string** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListsPostRequest

`func NewListsPostRequest(name string, icon string, ) *ListsPostRequest`

NewListsPostRequest instantiates a new ListsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListsPostRequestWithDefaults

`func NewListsPostRequestWithDefaults() *ListsPostRequest`

NewListsPostRequestWithDefaults instantiates a new ListsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *ListsPostRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ListsPostRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ListsPostRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetParentId

`func (o *ListsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ListsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ListsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ListsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ListsPostRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ListsPostRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


