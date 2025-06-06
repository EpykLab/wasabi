/*
Wazuh API REST

Testing ExperimentalAPIService

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

func Test_api_ExperimentalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerClearRootcheckDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerClearRootcheckDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerClearSyscheckDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerClearSyscheckDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetCisCatResults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetCisCatResults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetHardwareInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetHardwareInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetHotfixesInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetHotfixesInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetNetworkAddressInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkAddressInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetNetworkInterfaceInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkInterfaceInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetNetworkProtocolInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetNetworkProtocolInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetOsInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetOsInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetPackagesInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetPackagesInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetPortsInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetPortsInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExperimentalAPIService ApiControllersExperimentalControllerGetProcessesInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExperimentalAPI.ApiControllersExperimentalControllerGetProcessesInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
