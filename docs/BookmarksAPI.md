# \BookmarksAPI

All URIs are relative to *https://try.hoarder.app/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookmarksBookmarkIdDelete**](BookmarksAPI.md#BookmarksBookmarkIdDelete) | **Delete** /bookmarks/{bookmarkId} | Delete a bookmark
[**BookmarksBookmarkIdGet**](BookmarksAPI.md#BookmarksBookmarkIdGet) | **Get** /bookmarks/{bookmarkId} | Get a single bookmark
[**BookmarksBookmarkIdPatch**](BookmarksAPI.md#BookmarksBookmarkIdPatch) | **Patch** /bookmarks/{bookmarkId} | Update a bookmark
[**BookmarksBookmarkIdTagsDelete**](BookmarksAPI.md#BookmarksBookmarkIdTagsDelete) | **Delete** /bookmarks/{bookmarkId}/tags | Detach tags from a bookmark
[**BookmarksBookmarkIdTagsPost**](BookmarksAPI.md#BookmarksBookmarkIdTagsPost) | **Post** /bookmarks/{bookmarkId}/tags | Attach tags to a bookmark
[**BookmarksGet**](BookmarksAPI.md#BookmarksGet) | **Get** /bookmarks | Get all bookmarks
[**BookmarksPost**](BookmarksAPI.md#BookmarksPost) | **Post** /bookmarks | Create a new bookmark



## BookmarksBookmarkIdDelete

> BookmarksBookmarkIdDelete(ctx, bookmarkId).Execute()

Delete a bookmark



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
	bookmarkId := "bookmarkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdDelete(context.Background(), bookmarkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdDeleteRequest struct via the builder pattern


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


## BookmarksBookmarkIdGet

> Bookmark BookmarksBookmarkIdGet(ctx, bookmarkId).Execute()

Get a single bookmark



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
	bookmarkId := "bookmarkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdGet(context.Background(), bookmarkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdGet`: Bookmark
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Bookmark**](Bookmark.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksBookmarkIdPatch

> BookmarksBookmarkIdPatch200Response BookmarksBookmarkIdPatch(ctx, bookmarkId).BookmarksBookmarkIdPatchRequest(bookmarksBookmarkIdPatchRequest).Execute()

Update a bookmark



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
	bookmarkId := "bookmarkId_example" // string | 
	bookmarksBookmarkIdPatchRequest := *openapiclient.NewBookmarksBookmarkIdPatchRequest() // BookmarksBookmarkIdPatchRequest | The data to update. Only the fields you want to update need to be provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdPatch(context.Background(), bookmarkId).BookmarksBookmarkIdPatchRequest(bookmarksBookmarkIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdPatch`: BookmarksBookmarkIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookmarksBookmarkIdPatchRequest** | [**BookmarksBookmarkIdPatchRequest**](BookmarksBookmarkIdPatchRequest.md) | The data to update. Only the fields you want to update need to be provided. | 

### Return type

[**BookmarksBookmarkIdPatch200Response**](BookmarksBookmarkIdPatch200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksBookmarkIdTagsDelete

> BookmarksBookmarkIdTagsDelete200Response BookmarksBookmarkIdTagsDelete(ctx, bookmarkId).BookmarksBookmarkIdTagsPostRequest(bookmarksBookmarkIdTagsPostRequest).Execute()

Detach tags from a bookmark



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
	bookmarkId := "bookmarkId_example" // string | 
	bookmarksBookmarkIdTagsPostRequest := *openapiclient.NewBookmarksBookmarkIdTagsPostRequest([]openapiclient.BookmarksBookmarkIdTagsPostRequestTagsInner{*openapiclient.NewBookmarksBookmarkIdTagsPostRequestTagsInner()}) // BookmarksBookmarkIdTagsPostRequest | The tags to detach. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdTagsDelete(context.Background(), bookmarkId).BookmarksBookmarkIdTagsPostRequest(bookmarksBookmarkIdTagsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdTagsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdTagsDelete`: BookmarksBookmarkIdTagsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdTagsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdTagsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookmarksBookmarkIdTagsPostRequest** | [**BookmarksBookmarkIdTagsPostRequest**](BookmarksBookmarkIdTagsPostRequest.md) | The tags to detach. | 

### Return type

[**BookmarksBookmarkIdTagsDelete200Response**](BookmarksBookmarkIdTagsDelete200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksBookmarkIdTagsPost

> BookmarksBookmarkIdTagsPost200Response BookmarksBookmarkIdTagsPost(ctx, bookmarkId).BookmarksBookmarkIdTagsPostRequest(bookmarksBookmarkIdTagsPostRequest).Execute()

Attach tags to a bookmark



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
	bookmarkId := "bookmarkId_example" // string | 
	bookmarksBookmarkIdTagsPostRequest := *openapiclient.NewBookmarksBookmarkIdTagsPostRequest([]openapiclient.BookmarksBookmarkIdTagsPostRequestTagsInner{*openapiclient.NewBookmarksBookmarkIdTagsPostRequestTagsInner()}) // BookmarksBookmarkIdTagsPostRequest | The tags to attach. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdTagsPost(context.Background(), bookmarkId).BookmarksBookmarkIdTagsPostRequest(bookmarksBookmarkIdTagsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdTagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdTagsPost`: BookmarksBookmarkIdTagsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdTagsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookmarksBookmarkIdTagsPostRequest** | [**BookmarksBookmarkIdTagsPostRequest**](BookmarksBookmarkIdTagsPostRequest.md) | The tags to attach. | 

### Return type

[**BookmarksBookmarkIdTagsPost200Response**](BookmarksBookmarkIdTagsPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksGet

> PaginatedBookmarks BookmarksGet(ctx).Archived(archived).Favourited(favourited).Limit(limit).Cursor(cursor).Execute()

Get all bookmarks



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
	archived := true // bool |  (optional)
	favourited := true // bool |  (optional)
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksGet(context.Background()).Archived(archived).Favourited(favourited).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksGet`: PaginatedBookmarks
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** |  | 
 **favourited** | **bool** |  | 
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


## BookmarksPost

> Bookmark BookmarksPost(ctx).BookmarksPostRequest(bookmarksPostRequest).Execute()

Create a new bookmark



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
	bookmarksPostRequest := *openapiclient.NewBookmarksPostRequest("Type_example", "Url_example", "Text_example", "AssetType_example", "AssetId_example") // BookmarksPostRequest | The bookmark to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksPost(context.Background()).BookmarksPostRequest(bookmarksPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksPost`: Bookmark
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookmarksPostRequest** | [**BookmarksPostRequest**](BookmarksPostRequest.md) | The bookmark to create | 

### Return type

[**Bookmark**](Bookmark.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

