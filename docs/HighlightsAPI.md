# \HighlightsAPI

All URIs are relative to *https://try.hoarder.app/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HighlightsGet**](HighlightsAPI.md#HighlightsGet) | **Get** /highlights | Get all highlights
[**HighlightsHighlightIdDelete**](HighlightsAPI.md#HighlightsHighlightIdDelete) | **Delete** /highlights/{highlightId} | Delete a highlight
[**HighlightsHighlightIdGet**](HighlightsAPI.md#HighlightsHighlightIdGet) | **Get** /highlights/{highlightId} | Get a single highlight
[**HighlightsHighlightIdPatch**](HighlightsAPI.md#HighlightsHighlightIdPatch) | **Patch** /highlights/{highlightId} | Update a highlight
[**HighlightsPost**](HighlightsAPI.md#HighlightsPost) | **Post** /highlights | Create a new highlight



## HighlightsGet

> PaginatedHighlights HighlightsGet(ctx).Limit(limit).Cursor(cursor).Execute()

Get all highlights



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
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.HighlightsGet(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.HighlightsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HighlightsGet`: PaginatedHighlights
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.HighlightsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHighlightsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | 
 **cursor** | **string** |  | 

### Return type

[**PaginatedHighlights**](PaginatedHighlights.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HighlightsHighlightIdDelete

> Highlight HighlightsHighlightIdDelete(ctx, highlightId).Execute()

Delete a highlight



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
	highlightId := "highlightId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.HighlightsHighlightIdDelete(context.Background(), highlightId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.HighlightsHighlightIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HighlightsHighlightIdDelete`: Highlight
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.HighlightsHighlightIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**highlightId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHighlightsHighlightIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Highlight**](Highlight.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HighlightsHighlightIdGet

> Highlight HighlightsHighlightIdGet(ctx, highlightId).Execute()

Get a single highlight



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
	highlightId := "highlightId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.HighlightsHighlightIdGet(context.Background(), highlightId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.HighlightsHighlightIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HighlightsHighlightIdGet`: Highlight
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.HighlightsHighlightIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**highlightId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHighlightsHighlightIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Highlight**](Highlight.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HighlightsHighlightIdPatch

> Highlight HighlightsHighlightIdPatch(ctx, highlightId).HighlightsHighlightIdPatchRequest(highlightsHighlightIdPatchRequest).Execute()

Update a highlight



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
	highlightId := "highlightId_example" // string | 
	highlightsHighlightIdPatchRequest := *openapiclient.NewHighlightsHighlightIdPatchRequest() // HighlightsHighlightIdPatchRequest | The data to update. Only the fields you want to update need to be provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.HighlightsHighlightIdPatch(context.Background(), highlightId).HighlightsHighlightIdPatchRequest(highlightsHighlightIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.HighlightsHighlightIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HighlightsHighlightIdPatch`: Highlight
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.HighlightsHighlightIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**highlightId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHighlightsHighlightIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **highlightsHighlightIdPatchRequest** | [**HighlightsHighlightIdPatchRequest**](HighlightsHighlightIdPatchRequest.md) | The data to update. Only the fields you want to update need to be provided. | 

### Return type

[**Highlight**](Highlight.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HighlightsPost

> Highlight HighlightsPost(ctx).HighlightsPostRequest(highlightsPostRequest).Execute()

Create a new highlight



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
	highlightsPostRequest := *openapiclient.NewHighlightsPostRequest("BookmarkId_example", float32(123), float32(123), "Text_example", "Note_example") // HighlightsPostRequest | The highlight to create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighlightsAPI.HighlightsPost(context.Background()).HighlightsPostRequest(highlightsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighlightsAPI.HighlightsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HighlightsPost`: Highlight
	fmt.Fprintf(os.Stdout, "Response from `HighlightsAPI.HighlightsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHighlightsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **highlightsPostRequest** | [**HighlightsPostRequest**](HighlightsPostRequest.md) | The highlight to create | 

### Return type

[**Highlight**](Highlight.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

