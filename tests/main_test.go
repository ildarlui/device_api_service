package tests

import (
	"fmt"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/tests/config"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("init smoke tests")
	defer fmt.Println("completed smoke tests")

	cfg, _ := config.GetConfig()
	checkUrl := url.URL{
		Scheme: "HTTP",
		Host:   cfg.Host,
		Path:   cfg.Livecheck,
	}

	liveCheckOk := false
	for i := 0; i < 10; i++ {
		res, err := http.Get(checkUrl.String())
		if err == nil && res.StatusCode == http.StatusOK {
			liveCheckOk = true
			break
		}
		log.Printf("service is not running: %v", err)
		log.Printf("Status is not ok: %v", res)
		time.Sleep(5000)
	}
	if !liveCheckOk {
		log.Fatal("service launch problem")
	}
	m.Run()
}
