# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | **string** |  | 
**ParentId** | **NullableString** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "manual"]
**Query** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewList

`func NewList(id string, name string, icon string, parentId NullableString, ) *List`

NewList instantiates a new List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithDefaults

`func NewListWithDefaults() *List`

NewListWithDefaults instantiates a new List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *List) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *List) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *List) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *List) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *List) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *List) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *List) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *List) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *List) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetParentId

`func (o *List) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *List) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *List) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *List) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *List) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetType

`func (o *List) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *List) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *List) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *List) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuery

`func (o *List) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *List) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *List) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *List) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *List) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *List) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


