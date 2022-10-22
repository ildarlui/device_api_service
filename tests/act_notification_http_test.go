package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSmokeActNotificationApiService(t *testing.T) {
	newNotification := &SendNotificationRequest{
		Notification{"0", "1", "Ildar", "AAA", "LANG_ENGLISH", "STATUS_CREATED"},
	}
	sendNotificationRes, err := SendNotificationHttp(newNotification)
	require.NoError(t, err)
	t.Logf("Send notidication, NotificationId: %s", sendNotificationRes.NotificationId)

	getNotificationRes, err := GetNotificationHttp(newNotification.Notification.DeviceID)
	require.NoError(t, err)
	require.Equal(t, sendNotificationRes.NotificationId, getNotificationRes.Notification[0].NotificationID)
	fmt.Println(getNotificationRes)
	t.Logf("Get notidications for deviceId: %s", newNotification.Notification.DeviceID)

	askNotificationRes, err := AckNotificationHttp(sendNotificationRes.NotificationId)
	require.NoError(t, err)
	require.True(t, askNotificationRes.Success)
	t.Logf("Ask notofocatioId: %s, status: %t", sendNotificationRes.NotificationId, askNotificationRes.Success)

	getNotificationAfterAskRes, err := GetNotificationHttp(newNotification.Notification.DeviceID)
	success := checkNotificationInListHttp(getNotificationAfterAskRes.Notification, sendNotificationRes.NotificationId)
	require.False(t, success)
	t.Logf("Notification with id %s asked", sendNotificationRes.NotificationId)

}
