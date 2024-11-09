/*
Hoarder API

The API for the Hoarder app

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TagsAPIService TagsAPI service
type TagsAPIService service

type ApiTagsGetRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
}

func (r ApiTagsGetRequest) Execute() (*TagsGet200Response, *http.Response, error) {
	return r.ApiService.TagsGetExecute(r)
}

/*
TagsGet Get all tags

Get all tags

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTagsGetRequest
*/
func (a *TagsAPIService) TagsGet(ctx context.Context) ApiTagsGetRequest {
	return ApiTagsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TagsGet200Response
func (a *TagsAPIService) TagsGetExecute(r ApiTagsGetRequest) (*TagsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTagsTagIdBookmarksGetRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	tagId string
	limit *float32
	cursor *string
}

func (r ApiTagsTagIdBookmarksGetRequest) Limit(limit float32) ApiTagsTagIdBookmarksGetRequest {
	r.limit = &limit
	return r
}

func (r ApiTagsTagIdBookmarksGetRequest) Cursor(cursor string) ApiTagsTagIdBookmarksGetRequest {
	r.cursor = &cursor
	return r
}

func (r ApiTagsTagIdBookmarksGetRequest) Execute() (*PaginatedBookmarks, *http.Response, error) {
	return r.ApiService.TagsTagIdBookmarksGetExecute(r)
}

/*
TagsTagIdBookmarksGet Get a bookmarks with the tag

Get the bookmarks with the tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId
 @return ApiTagsTagIdBookmarksGetRequest
*/
func (a *TagsAPIService) TagsTagIdBookmarksGet(ctx context.Context, tagId string) ApiTagsTagIdBookmarksGetRequest {
	return ApiTagsTagIdBookmarksGetRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
	}
}

// Execute executes the request
//  @return PaginatedBookmarks
func (a *TagsAPIService) TagsTagIdBookmarksGetExecute(r ApiTagsTagIdBookmarksGetRequest) (*PaginatedBookmarks, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedBookmarks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsTagIdBookmarksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{tagId}/bookmarks"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", url.PathEscape(parameterValueToString(r.tagId, "tagId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTagsTagIdDeleteRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	tagId string
}

func (r ApiTagsTagIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.TagsTagIdDeleteExecute(r)
}

/*
TagsTagIdDelete Delete a tag

Delete tag by its id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId
 @return ApiTagsTagIdDeleteRequest
*/
func (a *TagsAPIService) TagsTagIdDelete(ctx context.Context, tagId string) ApiTagsTagIdDeleteRequest {
	return ApiTagsTagIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
	}
}

// Execute executes the request
func (a *TagsAPIService) TagsTagIdDeleteExecute(r ApiTagsTagIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsTagIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{tagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", url.PathEscape(parameterValueToString(r.tagId, "tagId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiTagsTagIdGetRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	tagId string
}

func (r ApiTagsTagIdGetRequest) Execute() (*Tag, *http.Response, error) {
	return r.ApiService.TagsTagIdGetExecute(r)
}

/*
TagsTagIdGet Get a single tag

Get tag by its id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId
 @return ApiTagsTagIdGetRequest
*/
func (a *TagsAPIService) TagsTagIdGet(ctx context.Context, tagId string) ApiTagsTagIdGetRequest {
	return ApiTagsTagIdGetRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
	}
}

// Execute executes the request
//  @return Tag
func (a *TagsAPIService) TagsTagIdGetExecute(r ApiTagsTagIdGetRequest) (*Tag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsTagIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{tagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", url.PathEscape(parameterValueToString(r.tagId, "tagId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTagsTagIdPatchRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	tagId string
	tagsTagIdPatchRequest *TagsTagIdPatchRequest
}

// The data to update. Only the fields you want to update need to be provided.
func (r ApiTagsTagIdPatchRequest) TagsTagIdPatchRequest(tagsTagIdPatchRequest TagsTagIdPatchRequest) ApiTagsTagIdPatchRequest {
	r.tagsTagIdPatchRequest = &tagsTagIdPatchRequest
	return r
}

func (r ApiTagsTagIdPatchRequest) Execute() (*Tag, *http.Response, error) {
	return r.ApiService.TagsTagIdPatchExecute(r)
}

/*
TagsTagIdPatch Update a tag

Update tag by its id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tagId
 @return ApiTagsTagIdPatchRequest
*/
func (a *TagsAPIService) TagsTagIdPatch(ctx context.Context, tagId string) ApiTagsTagIdPatchRequest {
	return ApiTagsTagIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		tagId: tagId,
	}
}

// Execute executes the request
//  @return Tag
func (a *TagsAPIService) TagsTagIdPatchExecute(r ApiTagsTagIdPatchRequest) (*Tag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.TagsTagIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{tagId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tagId"+"}", url.PathEscape(parameterValueToString(r.tagId, "tagId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tagsTagIdPatchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
