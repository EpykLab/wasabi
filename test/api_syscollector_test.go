/*
Wazuh API REST

Testing SyscollectorAPIService

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

func Test_api_SyscollectorAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetHardwareInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetHardwareInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetHotfixInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetHotfixInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetNetworkAddressInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkAddressInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetNetworkInterfaceInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkInterfaceInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetNetworkProtocolInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetNetworkProtocolInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetOsInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetOsInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetPackagesInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetPackagesInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetPortsInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetPortsInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SyscollectorAPIService ApiControllersSyscollectorControllerGetProcessesInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentId string

		resp, httpRes, err := apiClient.SyscollectorAPI.ApiControllersSyscollectorControllerGetProcessesInfo(context.Background(), agentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
