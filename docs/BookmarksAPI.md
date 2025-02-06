# \BookmarksAPI

All URIs are relative to *https://try.hoarder.app/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookmarksBookmarkIdAssetsAssetIdDelete**](BookmarksAPI.md#BookmarksBookmarkIdAssetsAssetIdDelete) | **Delete** /bookmarks/{bookmarkId}/assets/{assetId} | Detach asset
[**BookmarksBookmarkIdAssetsAssetIdPut**](BookmarksAPI.md#BookmarksBookmarkIdAssetsAssetIdPut) | **Put** /bookmarks/{bookmarkId}/assets/{assetId} | Replace asset
[**BookmarksBookmarkIdAssetsPost**](BookmarksAPI.md#BookmarksBookmarkIdAssetsPost) | **Post** /bookmarks/{bookmarkId}/assets | Attach asset
[**BookmarksBookmarkIdDelete**](BookmarksAPI.md#BookmarksBookmarkIdDelete) | **Delete** /bookmarks/{bookmarkId} | Delete a bookmark
[**BookmarksBookmarkIdGet**](BookmarksAPI.md#BookmarksBookmarkIdGet) | **Get** /bookmarks/{bookmarkId} | Get a single bookmark
[**BookmarksBookmarkIdHighlightsGet**](BookmarksAPI.md#BookmarksBookmarkIdHighlightsGet) | **Get** /bookmarks/{bookmarkId}/highlights | Get highlights of a bookmark
[**BookmarksBookmarkIdPatch**](BookmarksAPI.md#BookmarksBookmarkIdPatch) | **Patch** /bookmarks/{bookmarkId} | Update a bookmark
[**BookmarksBookmarkIdTagsDelete**](BookmarksAPI.md#BookmarksBookmarkIdTagsDelete) | **Delete** /bookmarks/{bookmarkId}/tags | Detach tags from a bookmark
[**BookmarksBookmarkIdTagsPost**](BookmarksAPI.md#BookmarksBookmarkIdTagsPost) | **Post** /bookmarks/{bookmarkId}/tags | Attach tags to a bookmark
[**BookmarksGet**](BookmarksAPI.md#BookmarksGet) | **Get** /bookmarks | Get all bookmarks
[**BookmarksPost**](BookmarksAPI.md#BookmarksPost) | **Post** /bookmarks | Create a new bookmark
[**BookmarksSearchGet**](BookmarksAPI.md#BookmarksSearchGet) | **Get** /bookmarks/search | Search bookmarks



## BookmarksBookmarkIdAssetsAssetIdDelete

> BookmarksBookmarkIdAssetsAssetIdDelete(ctx, bookmarkId, assetId).Execute()

Detach asset



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
	assetId := "assetId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdAssetsAssetIdDelete(context.Background(), bookmarkId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdAssetsAssetIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdAssetsAssetIdDeleteRequest struct via the builder pattern


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


## BookmarksBookmarkIdAssetsAssetIdPut

> BookmarksBookmarkIdAssetsAssetIdPut(ctx, bookmarkId, assetId).BookmarksBookmarkIdAssetsAssetIdPutRequest(bookmarksBookmarkIdAssetsAssetIdPutRequest).Execute()

Replace asset



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
	assetId := "assetId_example" // string | 
	bookmarksBookmarkIdAssetsAssetIdPutRequest := *openapiclient.NewBookmarksBookmarkIdAssetsAssetIdPutRequest("AssetId_example") // BookmarksBookmarkIdAssetsAssetIdPutRequest | The new asset to replace with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdAssetsAssetIdPut(context.Background(), bookmarkId, assetId).BookmarksBookmarkIdAssetsAssetIdPutRequest(bookmarksBookmarkIdAssetsAssetIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdAssetsAssetIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdAssetsAssetIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bookmarksBookmarkIdAssetsAssetIdPutRequest** | [**BookmarksBookmarkIdAssetsAssetIdPutRequest**](BookmarksBookmarkIdAssetsAssetIdPutRequest.md) | The new asset to replace with | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksBookmarkIdAssetsPost

> BookmarksBookmarkIdAssetsPostRequest BookmarksBookmarkIdAssetsPost(ctx, bookmarkId).BookmarksBookmarkIdAssetsPostRequest(bookmarksBookmarkIdAssetsPostRequest).Execute()

Attach asset



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
	bookmarksBookmarkIdAssetsPostRequest := *openapiclient.NewBookmarksBookmarkIdAssetsPostRequest("Id_example", "AssetType_example") // BookmarksBookmarkIdAssetsPostRequest | The asset to attach (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdAssetsPost(context.Background(), bookmarkId).BookmarksBookmarkIdAssetsPostRequest(bookmarksBookmarkIdAssetsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdAssetsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdAssetsPost`: BookmarksBookmarkIdAssetsPostRequest
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdAssetsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdAssetsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bookmarksBookmarkIdAssetsPostRequest** | [**BookmarksBookmarkIdAssetsPostRequest**](BookmarksBookmarkIdAssetsPostRequest.md) | The asset to attach | 

### Return type

[**BookmarksBookmarkIdAssetsPostRequest**](BookmarksBookmarkIdAssetsPostRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## BookmarksBookmarkIdHighlightsGet

> BookmarksBookmarkIdHighlightsGet200Response BookmarksBookmarkIdHighlightsGet(ctx, bookmarkId).Execute()

Get highlights of a bookmark



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
	resp, r, err := apiClient.BookmarksAPI.BookmarksBookmarkIdHighlightsGet(context.Background(), bookmarkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksBookmarkIdHighlightsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksBookmarkIdHighlightsGet`: BookmarksBookmarkIdHighlightsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksBookmarkIdHighlightsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksBookmarkIdHighlightsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookmarksBookmarkIdHighlightsGet200Response**](BookmarksBookmarkIdHighlightsGet200Response.md)

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


## BookmarksSearchGet

> PaginatedBookmarks BookmarksSearchGet(ctx).Q(q).Limit(limit).Cursor(cursor).Execute()

Search bookmarks



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
	q := "q_example" // string | 
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksSearchGet(context.Background()).Q(q).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksSearchGet`: PaginatedBookmarks
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
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

