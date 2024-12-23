/*
Hoarder API

Testing TagsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/flowline-io/sdk-hoarder-api"
)

func Test_openapi_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService TagsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPI.TagsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdBookmarksGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdBookmarksGet(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		httpRes, err := apiClient.TagsAPI.TagsTagIdDelete(context.Background(), tagId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdGet(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService TagsTagIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var tagId string

		resp, httpRes, err := apiClient.TagsAPI.TagsTagIdPatch(context.Background(), tagId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
