# \TagsAPI

All URIs are relative to *https://try.hoarder.app/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagsAPI.md#TagsGet) | **Get** /tags | Get all tags
[**TagsTagIdBookmarksGet**](TagsAPI.md#TagsTagIdBookmarksGet) | **Get** /tags/{tagId}/bookmarks | Get a bookmarks with the tag
[**TagsTagIdDelete**](TagsAPI.md#TagsTagIdDelete) | **Delete** /tags/{tagId} | Delete a tag
[**TagsTagIdGet**](TagsAPI.md#TagsTagIdGet) | **Get** /tags/{tagId} | Get a single tag
[**TagsTagIdPatch**](TagsAPI.md#TagsTagIdPatch) | **Patch** /tags/{tagId} | Update a tag



## TagsGet

> TagsGet200Response TagsGet(ctx).Execute()

Get all tags



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
	resp, r, err := apiClient.TagsAPI.TagsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsGet`: TagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.TagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTagsGetRequest struct via the builder pattern


### Return type

[**TagsGet200Response**](TagsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsTagIdBookmarksGet

> PaginatedBookmarks TagsTagIdBookmarksGet(ctx, tagId).Limit(limit).Cursor(cursor).Execute()

Get a bookmarks with the tag



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
	tagId := "tagId_example" // string | 
	limit := float32(8.14) // float32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.TagsTagIdBookmarksGet(context.Background(), tagId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsTagIdBookmarksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsTagIdBookmarksGet`: PaginatedBookmarks
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.TagsTagIdBookmarksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdBookmarksGetRequest struct via the builder pattern


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


## TagsTagIdDelete

> TagsTagIdDelete(ctx, tagId).Execute()

Delete a tag



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
	tagId := "tagId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.TagsTagIdDelete(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsTagIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdDeleteRequest struct via the builder pattern


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


## TagsTagIdGet

> Tag TagsTagIdGet(ctx, tagId).Execute()

Get a single tag



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
	tagId := "tagId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.TagsTagIdGet(context.Background(), tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsTagIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsTagIdGet`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.TagsTagIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagsTagIdPatch

> Tag TagsTagIdPatch(ctx, tagId).TagsTagIdPatchRequest(tagsTagIdPatchRequest).Execute()

Update a tag



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
	tagId := "tagId_example" // string | 
	tagsTagIdPatchRequest := *openapiclient.NewTagsTagIdPatchRequest() // TagsTagIdPatchRequest | The data to update. Only the fields you want to update need to be provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.TagsTagIdPatch(context.Background(), tagId).TagsTagIdPatchRequest(tagsTagIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.TagsTagIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagsTagIdPatch`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.TagsTagIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagsTagIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsTagIdPatchRequest** | [**TagsTagIdPatchRequest**](TagsTagIdPatchRequest.md) | The data to update. Only the fields you want to update need to be provided. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

