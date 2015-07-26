package geoip2_test

import (
	"encoding/json"
	"os"

	"github.com/savaki/geoip2"
)

func ExampleCity() {
	userId := os.Getenv("MAXMIND_USER_ID")
	licenseKey := os.Getenv("MAXMIND_LICENSE_KEY")
	api := geoip2.New(userId, licenseKey)
	resp, _ := api.City(nil, "1.2.3.4")
	json.NewEncoder(os.Stdout).Encode(resp)
}
