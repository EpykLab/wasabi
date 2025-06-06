/*
Wazuh API REST

Testing LogtestAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api

import (
	"context"
	openapiclient "github.com/EpykLab/wasabi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_api_LogtestAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogtestAPIService ApiControllersLogtestControllerEndLogtestSession", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.LogtestAPI.ApiControllersLogtestControllerEndLogtestSession(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogtestAPIService ApiControllersLogtestControllerRunLogtestTool", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LogtestAPI.ApiControllersLogtestControllerRunLogtestTool(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
