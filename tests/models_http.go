package tests

import "time"

type DeviceManager struct {
	BaseUrl string
}

type CreateDeviceResponse struct {
	DeviceId     string `json:"deviceId"`
	ErrorMessage string `json:"message,omitempty"`
}

type DeviceInfoResponse struct {
	Item         Item   `json:"value"`
	ErrorMessage string `json:"message,omitempty"`
}

type ListOfDevicesResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID        string    `json:"id"`
	Platform  string    `json:"platform"`
	UserID    string    `json:"userId"`
	EnteredAt time.Time `json:"enteredAt"`
}

type DeleteDeviceResponse struct {
	Found bool `json:"found"`
}

type SuccessDeviceResponse struct {
	Success bool `json:"success"`
}

type UpdateErrorResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details,omitempty"`
}

type GetNotificationResponse struct {
	Notification []Notification `json:"notification"`
}

type SendNotificationRequest struct {
	Notification Notification `json:"notification"`
}

type Notification struct {
	NotificationID     string `json:"notificationId,omitempty"`
	DeviceID           string `json:"deviceId,omitempty"`
	Username           string `json:"username,omitempty"`
	Message            string `json:"message,omitempty"`
	Lang               string `json:"lang,omitempty"`
	NotificationStatus string `json:"notificationStatus,omitempty"`
}

type NotificationResponse struct {
	NotificationId string `json:"notificationId"`
}
