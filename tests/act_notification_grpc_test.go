package tests

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestSmokeActNotificationGrpc(t *testing.T) {
	newNotification := act_device_api.Notification{NotificationId: 1, DeviceId: 1, Username: "Ildar", Message: "AAA", Lang: 1, NotificationStatus: 1}
	sendNotificationRes, err := SendNotificationGrpc(&newNotification)
	require.NoError(t, err)
	t.Logf("Send notidication, NotificationId: %d", sendNotificationRes.NotificationId)

	getNotificationRes, err := GetNotificationGrpc(newNotification.NotificationId)
	require.NoError(t, err)
	require.Equal(t, sendNotificationRes.NotificationId, getNotificationRes.Notification[0].NotificationId)
	fmt.Println(getNotificationRes)
	t.Logf("Get notidications for deviceId: %d", newNotification.DeviceId)

	askNotificationRes, err := AckNotificationGrpc(sendNotificationRes.NotificationId)
	require.NoError(t, err)
	require.True(t, askNotificationRes.Success)
	t.Logf("Ask notofocatioId: %d, status: %t", sendNotificationRes.NotificationId, askNotificationRes.Success)

	getNotificationAfterAskRes, err := GetNotificationGrpc(newNotification.DeviceId)
	success := CheckNotificationInListGrpc(getNotificationAfterAskRes, sendNotificationRes.NotificationId)
	require.False(t, success)
	t.Logf("Notification with id %d asked", sendNotificationRes.NotificationId)

}

func TestSubscribeNotification(t *testing.T) {
	var deviceId uint64 = 1

	ctx := context.Background()
	addr := "localhost:8082"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()
	client := act_device_api.NewActNotificationApiServiceClient(conn)
	reqSubscribeNotificat := &act_device_api.SubscribeNotificationRequest{DeviceId: deviceId}
	resSubscribeNotification, err := client.SubscribeNotification(ctx, reqSubscribeNotificat)

	newNotification := act_device_api.Notification{NotificationId: 1, DeviceId: 1, Username: "Ildar", Message: "AAA", Lang: 1, NotificationStatus: 1}
	sendNotificationRes, err := SendNotificationGrpc(&newNotification)
	require.NoError(t, err)
	t.Logf("Send notidication, NotificationId: %d", sendNotificationRes.NotificationId)

	str, _ := resSubscribeNotification.Recv()
	t.Logf("Response SubscribeNotification: %s", str)
	require.Equal(t, sendNotificationRes.NotificationId, str.NotificationId)
	require.NoError(t, err)

}
