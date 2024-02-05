/*
Candlepin

Testing EnvironmentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package caliri

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func Test_caliri_EnvironmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentAPIService CreateConsumerInEnvironment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envId string

		resp, httpRes, err := apiClient.EnvironmentAPI.CreateConsumerInEnvironment(context.Background(), envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAPIService DeleteEnvironment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envId string

		httpRes, err := apiClient.EnvironmentAPI.DeleteEnvironment(context.Background(), envId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAPIService DemoteContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envId string

		resp, httpRes, err := apiClient.EnvironmentAPI.DemoteContent(context.Background(), envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAPIService GetEnvironment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envId string

		resp, httpRes, err := apiClient.EnvironmentAPI.GetEnvironment(context.Background(), envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EnvironmentAPIService PromoteContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envId string

		resp, httpRes, err := apiClient.EnvironmentAPI.PromoteContent(context.Background(), envId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}