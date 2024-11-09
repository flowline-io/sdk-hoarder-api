/*
Hoarder API

Testing ListsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/flowline-io/sdk-hoarder-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ListsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ListsAPIService ListListIdPatch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListsAPI.ListListIdPatch(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ListsAPI.ListsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsListIdBookmarksBookmarkIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string
		var bookmarkId string

		httpRes, err := apiClient.ListsAPI.ListsListIdBookmarksBookmarkIdDelete(context.Background(), listId, bookmarkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsListIdBookmarksBookmarkIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string
		var bookmarkId string

		httpRes, err := apiClient.ListsAPI.ListsListIdBookmarksBookmarkIdPut(context.Background(), listId, bookmarkId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsListIdBookmarksGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListsAPI.ListsListIdBookmarksGet(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsListIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string

		httpRes, err := apiClient.ListsAPI.ListsListIdDelete(context.Background(), listId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsListIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var listId string

		resp, httpRes, err := apiClient.ListsAPI.ListsListIdGet(context.Background(), listId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ListsAPIService ListsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ListsAPI.ListsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
