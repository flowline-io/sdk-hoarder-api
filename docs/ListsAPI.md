# \ListsAPI

All URIs are relative to *https://try.hoarder.app/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsGet**](ListsAPI.md#ListsGet) | **Get** /lists | Get all lists
[**ListsListIdBookmarksBookmarkIdDelete**](ListsAPI.md#ListsListIdBookmarksBookmarkIdDelete) | **Delete** /lists/{listId}/bookmarks/{bookmarkId} | Remove a bookmark from a list
[**ListsListIdBookmarksBookmarkIdPut**](ListsAPI.md#ListsListIdBookmarksBookmarkIdPut) | **Put** /lists/{listId}/bookmarks/{bookmarkId} | Add a bookmark to a list
[**ListsListIdBookmarksGet**](ListsAPI.md#ListsListIdBookmarksGet) | **Get** /lists/{listId}/bookmarks | Get a bookmarks in a list
[**ListsListIdDelete**](ListsAPI.md#ListsListIdDelete) | **Delete** /lists/{listId} | Delete a list
[**ListsListIdGet**](ListsAPI.md#ListsListIdGet) | **Get** /lists/{listId} | Get a single list
[**ListsListIdPatch**](ListsAPI.md#ListsListIdPatch) | **Patch** /lists/{listId} | Update a list
[**ListsPost**](ListsAPI.md#ListsPost) | **Post** /lists | Create a new list



## ListsGet

> ListsGet200Response ListsGet(ctx).Execute()

Get all lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsGet`: ListsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListsGetRequest struct via the builder pattern


### Return type

[**ListsGet200Response**](ListsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdBookmarksBookmarkIdDelete

> ListsListIdBookmarksBookmarkIdDelete(ctx, listId, bookmarkId).Execute()

Remove a bookmark from a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 
	bookmarkId := "bookmarkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.ListsListIdBookmarksBookmarkIdDelete(context.Background(), listId, bookmarkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdBookmarksBookmarkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdBookmarksBookmarkIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdBookmarksBookmarkIdPut

> ListsListIdBookmarksBookmarkIdPut(ctx, listId, bookmarkId).Execute()

Add a bookmark to a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 
	bookmarkId := "bookmarkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.ListsListIdBookmarksBookmarkIdPut(context.Background(), listId, bookmarkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdBookmarksBookmarkIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdBookmarksBookmarkIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdBookmarksGet

> PaginatedBookmarks ListsListIdBookmarksGet(ctx, listId).Limit(limit).Cursor(cursor).Execute()

Get a bookmarks in a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsListIdBookmarksGet(context.Background(), listId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdBookmarksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsListIdBookmarksGet`: PaginatedBookmarks
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsListIdBookmarksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **float32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedBookmarks**](PaginatedBookmarks.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdDelete

> ListsListIdDelete(ctx, listId).Execute()

Delete a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.ListsListIdDelete(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdGet

> List ListsListIdGet(ctx, listId).Execute()

Get a single list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsListIdGet(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsListIdGet`: List
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsListIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**List**](List.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsListIdPatch

> List ListsListIdPatch(ctx, listId).ListsListIdPatchRequest(listsListIdPatchRequest).Execute()

Update a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listId := "listId_example" // string | 
	listsListIdPatchRequest := *openapiclient.NewListsListIdPatchRequest() // ListsListIdPatchRequest | The data to update. Only the fields you want to update need to be provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsListIdPatch(context.Background(), listId).ListsListIdPatchRequest(listsListIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsListIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsListIdPatch`: List
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsListIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsListIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listsListIdPatchRequest** | [**ListsListIdPatchRequest**](ListsListIdPatchRequest.md) | The data to update. Only the fields you want to update need to be provided. | 

### Return type

[**List**](List.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsPost

> List ListsPost(ctx).ListsPostRequest(listsPostRequest).Execute()

Create a new list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	listsPostRequest := *openapiclient.NewListsPostRequest("Name_example", "Icon_example") // ListsPostRequest | The list to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.ListsPost(context.Background()).ListsPostRequest(listsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.ListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListsPost`: List
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.ListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listsPostRequest** | [**ListsPostRequest**](ListsPostRequest.md) | The list to create | 

### Return type

[**List**](List.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

