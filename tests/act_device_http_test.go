package tests

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//type TestSuite struct {
//	suite.Suite
//	deviceManager DeviceManager
//}
//
//func (s *TestSuite) SetupSuite() {
//	cfg, err := config.NewFromYaml("../test_config.yaml")
//	require.NoError(s.T(), err)
//	s.deviceManager = DeviceManager{BaseUrl: cfg.BaseUrl}
//}

func TestSmokeActDeviceHttp(t *testing.T) {

	createRes, err := CreateDeviceHttp("100", "ROUTE")
	require.NoError(t, err)
	t.Logf("Created deviceId: %s", createRes.DeviceId)
	t.Cleanup(func() {
		deleteRes, err := DeleteDeviceHttp(createRes.DeviceId)
		require.NoError(t, err)
		require.True(t, deleteRes.Found)
		t.Logf("Deleted deviceId %s", createRes.DeviceId)
	})

	infoRes, err := DeviceInfoHttp(createRes.DeviceId)
	require.NoError(t, err)
	require.Equal(t, createRes.DeviceId, infoRes.Item.ID)
	t.Logf("Get info for deviceId %s", createRes.DeviceId)

	listRes, err := ListOfDevicesHttp()
	require.NoError(t, err)
	require.Equal(t, createRes.DeviceId, listRes.Items[0].ID)
	t.Logf("Get list info for devices, deviceId %s in list", createRes.DeviceId)

	updateRes, err := UpdateDeviceHttp(createRes.DeviceId, "200", "IOS")
	require.NoError(t, err)
	require.True(t, updateRes.Success)
	t.Logf("Udated userId, platform for deviceId %s", createRes.DeviceId)

	updateInfoRes, err := DeviceInfoHttp(createRes.DeviceId)
	require.NoError(t, err)
	require.Equal(t, "200", updateInfoRes.Item.UserID)
	require.Equal(t, "IOS", updateInfoRes.Item.Platform)
	t.Logf("Get info for updated deviceId %s", createRes.DeviceId)

}

func TestValidateCreateDeviceHttp(t *testing.T) {
	t.Parallel()
	type testcase struct {
		name          string
		userId        string
		platform      string
		expectedError string
	}

	tcs := []testcase{
		{
			name:          "empty_userId",
			userId:        "",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"",
		},
		{
			name:          "empty_platform",
			userId:        "123",
			platform:      "",
			expectedError: "value length must be at least 1 runes",
		},
		{
			name:          "not_valid_userId",
			userId:        "Abcd",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"Abcd\"",
		},
		{
			name:          "userId=0",
			userId:        "0",
			platform:      "ios",
			expectedError: "value must be greater than 0",
		},
		{
			name:          "userId=-1",
			userId:        "-1",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"-1\"",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			createRes, err := CreateDeviceHttp(tc.userId, tc.platform)
			require.NoError(t, err)
			require.Contains(t, createRes.ErrorMessage, tc.expectedError)
		})
	}
}

func TestValidateUpdateDeviceHttp(t *testing.T) {

	createRes, err := CreateDeviceHttp("100", "ROUTE")
	require.NoError(t, err)
	t.Cleanup(func() {
		deleteRes, err := DeleteDeviceHttp(createRes.DeviceId)
		require.NoError(t, err)
		require.True(t, deleteRes.Found)
	})

	type testcase struct {
		name          string
		userId        string
		platform      string
		expectedError string
	}

	tcs := []testcase{
		{
			name:          "empty_userId",
			userId:        "",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"",
		},
		{
			name:          "empty_platform",
			userId:        "123",
			platform:      "",
			expectedError: "value length must be at least 1 runes",
		},
		{
			name:          "not_valid_userId",
			userId:        "Abcd",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"Abcd\"",
		},
		{
			name:          "userId=0",
			userId:        "0",
			platform:      "ios",
			expectedError: "value must be greater than 0",
		},
		{
			name:          "userId=-1",
			userId:        "-1",
			platform:      "ios",
			expectedError: "invalid value for uint64 type: \"-1\"",
		},
	}
	t.Parallel()
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			updateRes, err := UpdateDeviceHttp(createRes.DeviceId, tc.userId, tc.platform)
			require.NoError(t, err)
			require.False(t, updateRes.Success)
		})
	}
}

func TestInfoNonExistentDevicesHttp(t *testing.T) {
	t.Parallel()
	infoRes, err := DeviceInfoHttp("9999999")
	require.NoError(t, err)
	require.Contains(t, infoRes.ErrorMessage, "device not found")
}

func TestUpdateNonExistentDevicesHttp(t *testing.T) {
	t.Parallel()
	deleteRes, err := UpdateDeviceHttp("9999999", "123", "ios")
	require.NoError(t, err)
	require.False(t, deleteRes.Success)
}

func TestDeleteNonExistentDevicesHttp(t *testing.T) {
	t.Parallel()
	deleteRes, err := DeleteDeviceHttp("9999999")
	require.NoError(t, err)
	require.False(t, deleteRes.Found)
}
