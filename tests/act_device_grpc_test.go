package tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSmokeActDeviceApiGrpc(t *testing.T) {

	createRes, err := CreateDeviceGrpc("ROUTE", 123)
	require.NoError(t, err)
	t.Logf("Created deviceId: %d", createRes.DeviceId)
	t.Cleanup(func() {
		deleteRes, err := DeleteDevicesGrpc(createRes.DeviceId)
		require.NoError(t, err)
		require.True(t, deleteRes.Found)
		t.Logf("Deleted deviceId %d", createRes.DeviceId)
	})

	infoRes, err := DeviceInfoGrpc(createRes.DeviceId)
	require.NoError(t, err)
	require.Equal(t, createRes.DeviceId, infoRes.Value.Id)
	t.Logf("Get info for deviceId %d", createRes.DeviceId)

	listRes, err := ListDevicesGrpc(1, 15)
	require.NoError(t, err)
	require.Equal(t, createRes.DeviceId, listRes.Items[0].Id)
	t.Logf("Get list info for devices, deviceId %d in list", createRes.DeviceId)

	updateRes, err := UpdateDeviceGrpc(createRes.DeviceId, 200, "IOS")
	require.NoError(t, err)
	require.True(t, updateRes.Success)
	t.Logf("Udated userId, platform for deviceId %d", createRes.DeviceId)

	updateInfoRes, err := DeviceInfoGrpc(createRes.DeviceId)
	require.NoError(t, err)
	require.Equal(t, uint64(200), updateInfoRes.Value.UserId)
	require.Equal(t, "IOS", updateInfoRes.Value.Platform)
	t.Logf("Get info for updated deviceId %d", createRes.DeviceId)
}
