package tests

import (
	"context"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
)

type BaseAddr struct {
	addr string
}

func GetBaseAddr() string {
	baseAddr := BaseAddr{addr: "localhost:8082"}
	return baseAddr.addr
}

func CreateDeviceGrpc(platform string, userId uint64) (*act_device_api.CreateDeviceV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActDeviceApiServiceClient(conn)
	req := &act_device_api.CreateDeviceV1Request{Platform: platform, UserId: userId}
	res, err := client.CreateDeviceV1(ctx, req)
	return res, err
}

func DeviceInfoGrpc(deviceId uint64) (*act_device_api.DescribeDeviceV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActDeviceApiServiceClient(conn)
	req := &act_device_api.DescribeDeviceV1Request{DeviceId: deviceId}
	res, err := client.DescribeDeviceV1(ctx, req)
	return res, err
}

func UpdateDeviceGrpc(deviceId, userId uint64, platform string) (*act_device_api.UpdateDeviceV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActDeviceApiServiceClient(conn)
	req := &act_device_api.UpdateDeviceV1Request{DeviceId: deviceId, Platform: platform, UserId: userId}
	res, err := client.UpdateDeviceV1(ctx, req)
	return res, err
}

func ListDevicesGrpc(page, perPage uint64) (*act_device_api.ListDevicesV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActDeviceApiServiceClient(conn)
	req := &act_device_api.ListDevicesV1Request{Page: page, PerPage: perPage}
	res, err := client.ListDevicesV1(ctx, req)
	return res, err
}

func DeleteDevicesGrpc(deviceId uint64) (*act_device_api.RemoveDeviceV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActDeviceApiServiceClient(conn)
	req := &act_device_api.RemoveDeviceV1Request{DeviceId: deviceId}
	res, err := client.RemoveDeviceV1(ctx, req)
	return res, err
}

func GetNotificationGrpc(deviceId uint64) (*act_device_api.GetNotificationV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActNotificationApiServiceClient(conn)
	req := &act_device_api.GetNotificationV1Request{DeviceId: deviceId}
	res, err := client.GetNotification(ctx, req)
	return res, err
}

func SendNotificationGrpc(notification *act_device_api.Notification) (*act_device_api.SendNotificationV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActNotificationApiServiceClient(conn)
	req := &act_device_api.SendNotificationV1Request{Notification: notification}
	res, err := client.SendNotificationV1(ctx, req)
	return res, err
}

func AckNotificationGrpc(notificationId uint64) (*act_device_api.AckNotificationV1Response, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActNotificationApiServiceClient(conn)
	req := &act_device_api.AckNotificationV1Request{NotificationId: notificationId}
	res, err := client.AckNotification(ctx, req)
	return res, err
}

func SubscribeNotification(deviceId uint64) (act_device_api.ActNotificationApiService_SubscribeNotificationClient, error) {
	ctx := context.Background()

	addr := GetBaseAddr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := act_device_api.NewActNotificationApiServiceClient(conn)
	req := &act_device_api.SubscribeNotificationRequest{DeviceId: deviceId}
	res, err := client.SubscribeNotification(ctx, req)

	return res, err
}

func CheckNotificationInListGrpc(list *act_device_api.GetNotificationV1Response, notificationId uint64) bool {
	for _, Notification := range list.Notification {
		if Notification.NotificationId == notificationId {
			return true
		}
	}
	return false
}
