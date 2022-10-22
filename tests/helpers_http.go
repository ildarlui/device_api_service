package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BaseUrl struct {
	BaseUrl string
}

func GetBaseUrl() string {
	baseUrl := BaseUrl{BaseUrl: "http://localhost:8080"}
	return baseUrl.BaseUrl
}

func CreateDeviceHttp(userId, platform string) (*CreateDeviceResponse, error) {
	path := "/api/v1/devices"
	url := fmt.Sprintf("%s%s", GetBaseUrl(), path)
	method := "POST"

	params := fmt.Sprintf(`{
  "platform": "%s",
  "userId": "%s"
}`, platform, userId)
	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	fmt.Sprintln(string(body))

	data := CreateDeviceResponse{}
	json.Unmarshal(body, &data)
	return &data, err

}

func DeviceInfoHttp(deviceId string) (*DeviceInfoResponse, error) {
	path := "/api/v1/devices/"
	url := fmt.Sprintf("%s%s%s", GetBaseUrl(), path, deviceId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := DeviceInfoResponse{}
	json.Unmarshal(body, &data)
	return &data, err
}

func ListOfDevicesHttp() (*ListOfDevicesResponse, error) {
	path := "/api/v1/devices"
	url := fmt.Sprintf("%s%s?page=1&perPage=50", GetBaseUrl(), path)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := ListOfDevicesResponse{}
	json.Unmarshal(body, &data)
	return &data, err
}

func DeleteDeviceHttp(deviceId string) (*DeleteDeviceResponse, error) {
	path := "/api/v1/devices/"
	url := fmt.Sprintf("%s%s%s", GetBaseUrl(), path, deviceId)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := DeleteDeviceResponse{}
	json.Unmarshal(body, &data)
	return &data, err
}

func UpdateDeviceHttp(deviceId, userId, platform string) (*SuccessDeviceResponse, error) {
	path := "/api/v1/devices/"
	url := fmt.Sprintf("%s%s%s", GetBaseUrl(), path, deviceId)
	method := "PUT"

	params := fmt.Sprintf(`{
  "platform": "%s",
  "userId": "%s"
}`, platform, userId)
	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := SuccessDeviceResponse{}
	json.Unmarshal(body, &data)
	return &data, err
}

func checkDeviceIsInListHttp(items []Item, deviceId string) bool {
	for _, Item := range items {
		if Item.ID == deviceId {
			return true
		}
	}
	return false
}

func checkNotificationInListHttp(list []Notification, notificationId string) bool {
	for _, Notification := range list {
		if Notification.NotificationID == notificationId {
			return true
		}
	}
	return false
}

func GetNotificationHttp(deviceId string) (*GetNotificationResponse, error) {
	baseUrl := "http://localhost:8080/api/v1/notification?deviceId="
	url := fmt.Sprintf("%s%s", baseUrl, deviceId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data := GetNotificationResponse{}
	json.Unmarshal(body, &data)
	return &data, err

}

func SendNotificationHttp(reqData *SendNotificationRequest) (*NotificationResponse, error) {
	path := "/api/v1/notification"
	url := fmt.Sprintf("%s%s", GetBaseUrl(), path)

	method := "POST"

	params := fmt.Sprintf(`{
  "notification": {
    "notificationId": "%s",
    "deviceId": "%s",
    "username": "%s",
    "message": "%s",
    "lang": "%s",
    "notificationStatus": "%s"
  }
}`, reqData.Notification.NotificationID, reqData.Notification.DeviceID, reqData.Notification.Username,
		reqData.Notification.Message, reqData.Notification.Lang, reqData.Notification.NotificationStatus)

	payload := strings.NewReader(params)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := NotificationResponse{}
	json.Unmarshal(body, &data)
	return &data, err

}

func AckNotificationHttp(notificationId string) (*SuccessDeviceResponse, error) {
	path := "/api/v1/notification/ack/"
	url := fmt.Sprintf("%s%s%s", GetBaseUrl(), path, notificationId)
	method := "PUT"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	data := SuccessDeviceResponse{}
	json.Unmarshal(body, &data)
	return &data, err
}
